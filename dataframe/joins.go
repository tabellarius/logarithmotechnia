package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
	"strconv"
	"strings"
)

// InnerJoin makes an inner join with another dataframe
func (df *Dataframe) InnerJoin(with *Dataframe, options ...vector.Option) *Dataframe {
	conf := vector.MergeOptions(options)
	columns := df.determineColumns(conf, with)

	if len(columns) == 0 {
		return df
	}

	rootDfTree := &joinNode{}
	rootWithTree := &joinNode{}
	fillJoinTree(df, rootDfTree, columns)
	fillJoinTree(with, rootWithTree, columns)

	dfTreeKeys := rootDfTree.getKeys()

	dfIndices := make([]int, 0)
	withIndices := make([]int, 0)
	for _, key := range dfTreeKeys {
		indicesForWith := rootWithTree.getIndicesFor(key)
		if indicesForWith == nil {
			continue
		}
		indicesForDf := rootDfTree.getIndicesFor(key)
		for _, idxDf := range indicesForDf {
			for _, idxWith := range indicesForWith {
				dfIndices = append(dfIndices, idxDf)
				withIndices = append(withIndices, idxWith)
			}
		}
	}

	removeColumns := make([]string, len(columns))
	for i, column := range columns {
		removeColumns[i] = "-" + column
	}

	newDf := df.ByIndices(dfIndices)
	newWIth := with.Select(removeColumns).ByIndices(withIndices)

	return newDf.BindColumns(newWIth)
}

// LeftJoin makes a left join with another dataframe.
func (df *Dataframe) LeftJoin(with *Dataframe, options ...vector.Option) *Dataframe {
	conf := vector.MergeOptions(options)
	columns := df.determineColumns(conf, with)

	if len(columns) == 0 {
		return df
	}

	rootDfTree := &joinNode{}
	rootWithTree := &joinNode{}
	fillJoinTree(df, rootDfTree, columns)
	fillJoinTree(with, rootWithTree, columns)

	dfTreeKeys := rootDfTree.getKeys()

	dfIndices := make([]int, 0)
	withIndices := make([]int, 0)
	for _, key := range dfTreeKeys {
		indicesForDf := rootDfTree.getIndicesFor(key)
		indicesForWith := rootWithTree.getIndicesFor(key)
		if indicesForWith == nil {
			for _, idxDf := range indicesForDf {
				dfIndices = append(dfIndices, idxDf)
				withIndices = append(withIndices, 0)
			}
		} else {
			for _, idxDf := range indicesForDf {
				for _, idxWith := range indicesForWith {
					dfIndices = append(dfIndices, idxDf)
					withIndices = append(withIndices, idxWith)
				}
			}
		}
	}

	removeColumns := make([]string, len(columns))
	for i, column := range columns {
		removeColumns[i] = "-" + column
	}

	newDf := df.ByIndices(dfIndices)
	newWIth := with.Select(removeColumns).ByIndices(withIndices)

	return newDf.BindColumns(newWIth)
}

// RightJoin makes a right join with another dataframe
func (df *Dataframe) RightJoin(with *Dataframe, options ...vector.Option) *Dataframe {
	conf := vector.MergeOptions(options)
	columns := df.determineColumns(conf, with)

	if len(columns) == 0 {
		return df
	}

	rootDfTree := &joinNode{}
	rootWithTree := &joinNode{}
	fillJoinTree(df, rootDfTree, columns)
	fillJoinTree(with, rootWithTree, columns)

	treeKeys := rootWithTree.getKeys()

	dfIndices := make([]int, 0)
	withIndices := make([]int, 0)
	for _, key := range treeKeys {
		indicesForDf := rootDfTree.getIndicesFor(key)
		indicesForWith := rootWithTree.getIndicesFor(key)
		if indicesForDf == nil {
			for _, idxWith := range indicesForWith {
				dfIndices = append(dfIndices, 0)
				withIndices = append(withIndices, idxWith)
			}
		} else {
			for _, idxDf := range indicesForDf {
				for _, idxWith := range indicesForWith {
					dfIndices = append(dfIndices, idxDf)
					withIndices = append(withIndices, idxWith)
				}
			}
		}
	}

	removeColumns := make([]string, len(columns))
	for i, column := range columns {
		removeColumns[i] = "-" + column
	}

	newDf := df.Select(removeColumns).ByIndices(dfIndices)
	newWIth := with.ByIndices(withIndices)
	joinedDf := newDf.BindColumns(newWIth)

	selectNames, _ := df.Names().Append(joinedDf.Names()).Unique().Strings()

	return joinedDf.Select(selectNames)
}

// FullJoin makes a full join with another dataframe
func (df *Dataframe) FullJoin(with *Dataframe, options ...vector.Option) *Dataframe {
	conf := vector.MergeOptions(options)
	columns := df.determineColumns(conf, with)

	if len(columns) == 0 {
		return df
	}

	rootDfTree := &joinNode{}
	rootWithTree := &joinNode{}
	fillJoinTree(df, rootDfTree, columns)
	fillJoinTree(with, rootWithTree, columns)

	dfTreeKeys := rootDfTree.getKeys()
	withTreeKeys := rootWithTree.getKeys()

	dfIndices := make([]int, 0)
	withIndices := make([]int, 0)
	for _, key := range dfTreeKeys {
		indicesForDf := rootDfTree.getIndicesFor(key)
		indicesForWith := rootWithTree.getIndicesFor(key)
		if indicesForWith == nil {
			for _, idxDf := range indicesForDf {
				dfIndices = append(dfIndices, idxDf)
				withIndices = append(withIndices, 0)
			}
		} else {
			for _, idxDf := range indicesForDf {
				for _, idxWith := range indicesForWith {
					dfIndices = append(dfIndices, idxDf)
					withIndices = append(withIndices, idxWith)
				}
			}
		}
	}

	for _, key := range withTreeKeys {
		indicesForDf := rootDfTree.getIndicesFor(key)
		if indicesForDf == nil {
			indicesForWith := rootWithTree.getIndicesFor(key)
			for _, idxWith := range indicesForWith {
				dfIndices = append(dfIndices, 0)
				withIndices = append(withIndices, idxWith)
			}
		}
	}

	removeColumns := make([]string, len(columns))
	for i, column := range columns {
		removeColumns[i] = "-" + column
	}

	newDf := df.ByIndices(dfIndices)
	newWIth := with.ByIndices(withIndices)

	coalesceColumns := make([]Column, len(columns))
	for i, column := range columns {
		coalesceColumns[i] = Column{column, newDf.Cn(column).Coalesce(newWIth.Cn(column))}
	}

	return newDf.Mutate(coalesceColumns).BindColumns(newWIth.Select(removeColumns))
}

// SemiJoin makes a semi-join with another dataframe.
func (df *Dataframe) SemiJoin(with *Dataframe, options ...vector.Option) *Dataframe {
	conf := vector.MergeOptions(options)
	columns := df.determineColumns(conf, with)

	if len(columns) == 0 {
		return df
	}

	rootDfTree := &joinNode{}
	rootWithTree := &joinNode{}
	fillJoinTree(df, rootDfTree, columns)
	fillJoinTree(with, rootWithTree, columns)

	dfTreeKeys := rootDfTree.getKeys()

	dfIndices := make([]int, 0)
	for _, key := range dfTreeKeys {
		indicesForWith := rootWithTree.getIndicesFor(key)
		if indicesForWith == nil {
			continue
		}
		indicesForDf := rootDfTree.getIndicesFor(key)
		for _, idxDf := range indicesForDf {
			dfIndices = append(dfIndices, idxDf)
		}
	}

	return df.ByIndices(dfIndices)
}

// AntiJoin makes an anti-join with another dataframe.
func (df *Dataframe) AntiJoin(with *Dataframe, options ...vector.Option) *Dataframe {
	conf := vector.MergeOptions(options)
	columns := df.determineColumns(conf, with)

	if len(columns) == 0 {
		return df
	}

	rootDfTree := &joinNode{}
	rootWithTree := &joinNode{}
	fillJoinTree(df, rootDfTree, columns)
	fillJoinTree(with, rootWithTree, columns)

	dfTreeKeys := rootDfTree.getKeys()

	dfIndices := make([]int, 0)
	for _, key := range dfTreeKeys {
		indicesForWith := rootWithTree.getIndicesFor(key)
		if indicesForWith == nil {
			indicesForDf := rootDfTree.getIndicesFor(key)
			for _, idxDf := range indicesForDf {
				dfIndices = append(dfIndices, idxDf)
			}
		}
	}

	return df.ByIndices(dfIndices)
}

func (df *Dataframe) determineColumns(conf vector.Configuration, src *Dataframe) []string {
	var joinColumns []string

	if conf.HasOption(KeyOptionJoinBy) {
		columns := conf.Value(KeyOptionJoinBy).([]string)
		for _, column := range columns {
			if df.Names().Has(column) && src.Names().Has(column) {
				joinColumns = append(joinColumns, column)
			}
		}
	} else {
		joinColumns = []string{}
		for _, column := range df.columnNames {
			if src.Names().Has(column) {
				joinColumns = append(joinColumns, column)
			}
		}
	}

	return joinColumns
}

type joinNode struct {
	groupVal any
	groupMap map[any]*joinNode
	indices  []int
	values   []any
	keyLen   int
}

func (n *joinNode) getIndicesFor(key []any) []int {
	if len(key) == 0 {
		return nil
	}

	node, ok := n.groupMap[key[0]]
	if !ok {
		return nil
	}

	if len(key) == 1 {
		return node.indices
	}

	return node.getIndicesFor(key[1:])
}

func (n *joinNode) getKeys() [][]any {
	keys := [][]any{}

	if n.keyLen == 0 {
		return keys
	}

	for _, val := range n.values {
		if n.keyLen > 1 {
			subKeys := n.groupMap[val].getKeys()
			for _, subKey := range subKeys {
				key := append([]any{val}, subKey...)
				keys = append(keys, key)
			}
		} else {
			keys = append(keys, []any{val})
		}
	}

	return keys
}

func (n *joinNode) String() string {
	return joinNodeToString(n, 0)
}

func joinNodeToString(node *joinNode, lvl int) string {
	str := strings.Repeat("    ", lvl) + "Group: " + fmt.Sprintf("%v", node.groupVal) + "\n"
	str += strings.Repeat("    ", lvl) + "Values: " + fmt.Sprintf("%v", node.values) + "\n"
	str += strings.Repeat("    ", lvl) + "Values array length: " + strconv.Itoa(len(node.values)) + "\n"
	str += strings.Repeat("    ", lvl) + "Indices: " + fmt.Sprintf("%v", node.indices) + "\n"

	if len(node.values) > 0 {
		for _, value := range node.values {
			str += joinNodeToString(node.groupMap[value], lvl+1) + "\n"
		}
	}

	return str
}

func fillJoinTree(df *Dataframe, node *joinNode, columns []string) {
	if len(columns) == 0 || node == nil {
		return
	}

	isAdditionalColumns := len(columns) > 1

	node.groupMap = map[any]*joinNode{}
	node.keyLen = len(columns)
	column := columns[0]

	groups, values := df.Cn(column).Groups()
	node.values = values
	for i := 0; i < len(values); i++ {
		subNode := &joinNode{}
		subNode.groupVal = values[i]
		if node.indices == nil {
			subNode.indices = groups[i]
		} else {
			subNode.indices = make([]int, len(groups[i]))
			for j, idx := range groups[i] {
				subNode.indices[j] = node.indices[idx-1]
			}
		}

		if isAdditionalColumns {
			fillJoinTree(df.Filter(groups[i]), subNode, columns[1:])
		}

		node.groupMap[values[i]] = subNode
	}

	if len(values) > 0 {
		node.indices = nil
	}
}

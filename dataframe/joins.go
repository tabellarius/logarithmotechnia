package dataframe

import (
	"fmt"
	"logarithmotechnia/vector"
)

func (df *Dataframe) InnerJoin(with *Dataframe, options ...vector.Option) *Dataframe {
	conf := vector.MergeOptions(options)
	columns := df.determineColumns(conf, with)

	if len(columns) == 0 {
		return df
	}

	rootDfTree := &joinNode{}
	//	rootWithTree := &joinNode{}
	fillJoinTree(df, rootDfTree, columns)
	//	fillJoinTree(with, rootWithTree, columns)

	fmt.Println(rootDfTree)

	return nil
}

func (df *Dataframe) determineColumns(conf vector.Configuration, src *Dataframe) []string {
	var joinColumns []string

	if conf.HasOption(vector.KeyOptionJoinBy) {
		columns := conf.Value(vector.KeyOptionJoinBy).([]string)
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
	groupMap map[interface{}]*joinNode
	indices  []int
	values   []interface{}
	keyLen   int
}

func (n *joinNode) getIndicesFor(key []interface{}) []int {
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

func (n *joinNode) getKeys() [][]interface{} {
	keys := [][]interface{}{}

	if n.keyLen == 0 {
		return keys
	}

	for _, val := range n.values {
		if n.keyLen > 1 {
			subKeys := n.groupMap[val].getKeys()
			for _, subKey := range subKeys {
				key := append([]interface{}{val}, subKey...)
				keys = append(keys, key)
			}
		} else {
			keys = append(keys, []interface{}{val})
		}
	}
	fmt.Println("Keys:", keys)
	return keys
}

func fillJoinTree(df *Dataframe, node *joinNode, columns []string) {
	if len(columns) == 0 || node == nil {
		return
	}

	isAdditionalColumns := len(columns) > 1

	node.groupMap = map[interface{}]*joinNode{}
	node.keyLen = len(columns)
	column := columns[0]

	groups, values := df.Cn(column).Groups()
	node.values = values
	for i := 0; i < len(values); i++ {
		subNode := &joinNode{}
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

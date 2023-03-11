# Logarithmotechnia

This project is an implementaton of a dataframe akin to Python's Pandas or R's tibble/dplyr/tidyverse. R's influence is significantly stronger for this project, although I do borrow ideas from Pandas as well. 

Main advantages are decent data organization; full support of NA-values; good extensibility.

Dataframes and vectors (series in Pandas) are immutable.

Supported types are: Integer, Float, Complex, String, Boolean, Time and Any.

Loading from CSV
----------------
```Go
iris, err := dataframe.FromCSVFile("iris.csv")
```
To skip the first line use ```CSVOptionSkipFirstLine(true)```. 
```Go
iris, err := dataframe.FromCSVFile("iris.csv", dataframe.CSVOptionSkipFirstLine(true))
```
If you need to pass options for the new dataframe, use ```CSVOptionDataframeOptions(options...)```.

Loading from SQL
----------------
```Go
db, err := sql.Open("sqlite3", "./test_data/items.sqlite")
if err != nil {
	...
}

tx, err := db.Begin()
if err != nil {
	...
}

df, err := FromSQL(tx, "SELECT * FROM sku", []any{})
```
If you need to pass options for the new dataframe, use ```SQLOptionDataframeOptions(options...)```.

Filtering rows
--------------
Filtering is done with ```df.Filter(whicher)```. Two fundamental whichers are ```[]int``` with elements indices and
```[]bool```. ```Filter()``` filters out elements corresponding to ```false```. In most cases you do not need to pass 
```[]int``` or ```[]bool``` directly as they are returned by many column functions.

**Important:** in Logarithmothechnia first index is **1**! (This betrays **R**'s roots of Logarithmotechnia).

Let's select elements with "setosa" in "species" vector (column).

```Go
filteredDf := iris.Filter(iris.C("species").Eq("setosa"))
```
Or those with sepal length greater than 5. 
```Go
filteredDf := iris.Filter(iris.C("sepal_length").Gt(5))
```
Other comparing functions are: ```Neq()```, ```Lt()```, ```Gte()```, ```Lte()```.

It is possible to filter by using vector ```Which()``` function.

### Filtering by several conditions

What if you need to filter by two conditions at the same time? Here is two ways to do this:

```Go
filteredIris := iris.Filter(iris.C("species").Eq("setosa")).Filter(iris.C("sepal_length").Gte(5))
```

Or

```Go
filteredIris := iris.Filter(vector.And(
    iris.C("species").Eq("setosa"),
    iris.C("sepal_length").Gt(5),
))
```
The second approach is more general. What if you need to select all elements which are either of "setosa" species 
or have sepal length more than 5? It is easy to do by changing ```vector.And``` to ```vector.Or```.
```Go
filteredIris := iris.Filter(vector.Or(
    iris.C("species").Eq("setosa"),
    iris.C("sepal_length").Gt(5), 
))
```

### Filtering by function
It is also possible to filter by passing a function to column's ```Which()```.

```go
filteredIris = iris.Filter(iris.C("sepal_length").Which(
	func(val float64) bool {
		return val >= 5 && val < 7
	},
))
```
Function has to have a signature supported by the vector (column) type.

### Selecting dataframe subset
Select rows from 10th to 20th (including).
```go
subsetIris := iris.FromTo(10, 20)
```

Sorting
-------

To sort a dataframe use ```Arrange()``` function. For example, 

```go
sortedBySepalLength := iris.Arrange("sepal_length")
```
In reverse:

```go
sortedBySepalLengthReverse := iris.Arrange("sepal_length", dataframe.OptionArrangeReverse(true))
```
By two columns:
```go
sortedBySepalLength := iris.Arrange("series", "sepal_length")
```

Adding new columns
------------------

```Mutate()``` allows creates a new data frame with new columns, but preserving all columns of the old one. 
For example, let's add a column which indicate a one of two buckets based on the sepal length.

```go
bucketed := iris.Mutate(dataframe.Column{
	"bucket",
	iris.Cn("sepal_length").Apply(
		func(val float64) int {
			if val < 5 {
				return 1
			}
			
			return 2
		},
	),
})
```
Here you can also an example of vector's ```Apply()``` function which allows to generate a new vector from the other 
one. 

Selecting and dropping columns
------------------------------
```Select()``` function allows selecting and dropping dataframe's columns.

Let's select species and sepal length from iris dataset.

```go
compactIris := iris.Select("species", "sepal_length")
```

Or just drop petal length and petal_width.

```go
compactIris := iris.Select("-petal_length", "-petal_width")
```

It is also possible to use column indices instead of names.

```go
compactIris := iris.Select(5, 1)
```

Changing order of columns
-------------------------
Make "species" column appear before "sepal_length": 
```go
relocated := iris.Relocate("species", dataframe.OptionBeforeColumn("sepal_length"))
```
Or "petal_length" and "petal_width" after "species":
```go
relocated := iris.Relocate("petal_length", "petal_width", dataframe.OptionAfterColumn("species"))
```

Joining dataframes
-----
There are several types of joins available: ```InnerJoin()```, ```LeftJoin()```, ```RightJoin()```, ```FullJoin()```, 
```SemiJoin()``` and ```AntiJoin()```. Last two are from **dplyr**. Here is an example of left join:

```go
joined := employee.LeftJoin(department, OptionJoinBy("DepType"))
```
More examples of the joins can be found in tests.

Converting vectors to slices
----------------------------
Columns (and stand-alone vectors) can be converted to slices. For example:
```go
data, na := iris.Cn("species").Strings()
```
If an element of na is true, it means a corresponding element of the column is NA-value.

Available converting functions are 

* ```Booleans() ([]bool, []bool)```
* ```Integers() ([]int, []bool)```
* ```Floats() ([]float64, []bool)```
* ```Complexes() ([]complex128, []bool)```
* ```Strings() ([]string, []bool)```
* ```Times() ([]time.Time, []bool)```
* ```Anies() ([]any, []bool)```

Converting vectors to other types
---------------------------------
There are also similar functions for converting a vector to other type:

* ```AsInteger(options ...Option) Vector```
* ```AsFloat(options ...Option) Vector```
* ```AsComplex(options ...Option) Vector```
* ```AsBoolean(options ...Option) Vector```
* ```AsString(options ...Option) Vector```
* ```AsTime(options ...Option) Vector```
* ```AsAny(options ...Option) Vector```

Another way is to use ```Apply()``` function as shown before.

Renaming columns
----------------
To rename a column, use a ```Rename()``` function. There are several ways to pass which column to which value you 
would like to rename (check function comment). For example:

```go
renamedIris := iris.Rename([]string{"sepal_width", "s_width"})
```
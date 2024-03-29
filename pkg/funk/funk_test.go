// Package funk
// @author moqi
// On 2022/10/16 11:32:32
// see: https://github.com/thoas/go-funk
// see: https://liuqh.icu/2021/12/25/go/package/32-go-funk/
package funk

import (
	"fmt"
	"github.com/thoas/go-funk"
	"testing"
)

type Student struct {
	Name string
	Age  int
}

type Book struct {
	Id   int
	Name string
}

type User struct {
	Name string
	Home struct {
		City string
	}
}

type Foo struct {
	ID        int
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f Foo) TableName() string {
	return "foo"
}

func TestExist(t *testing.T) {
	// 判断任意类型
	fmt.Println("str->", funk.Contains([]string{"a", "b"}, "a"))
	// int 类型
	fmt.Println("int->", funk.ContainsInt([]int{1, 2}, 1))
}

/**输出
=== RUN   TestExist
str-> true
int-> true
--- PASS: TestExist (0.00s)
PASS
*/

func TestIndexOf(t *testing.T) {
	strArr := []string{"go", "java", "c", "java"}
	// 具体类型
	fmt.Println("c: ", funk.IndexOfString(strArr, "c"))
	// 验证第一次出现位置
	fmt.Println("java: ", funk.IndexOfString(strArr, "java"))
	// 任意类型
	fmt.Println("go: ", funk.IndexOf(strArr, "go"))
	// 不存在时返回-1
	fmt.Println("php: ", funk.IndexOfString(strArr, "php"))
}

/**输出
=== RUN   TestIndexOf
c:  2
java:  1
go:  0
php:  -1
--- PASS: TestIndexOf (0.00s)
PASS
*/

// 查找元素最后一次出现的位置
func TestLastOf(t *testing.T) {
	strArr := []string{"go", "java", "c", "java"}
	// 具体类型
	fmt.Println("c: ", funk.LastIndexOfString(strArr, "c"))
	// 验证第一次出现位置
	fmt.Println("java: ", funk.LastIndexOfString(strArr, "java"))
	// 任意类型
	fmt.Println("go: ", funk.LastIndexOf(strArr, "go"))
	// 不存在时返回-1
	fmt.Println("php: ", funk.LastIndexOf(strArr, "php"))
}

/**输出
=== RUN   TestLastOf
c:  2
java:  3
go:  0
php:  -1
--- PASS: TestLastOf (0.00s)
PASS
*/

func TestEvery(t *testing.T) {
	strArr := []string{"go", "java", "c", "python"}
	fmt.Println("都存在:", funk.Every(strArr, "go", "c"))
	fmt.Println("有一个不存在:", funk.Every(strArr, "php", "go"))
	fmt.Println("都不存在:", funk.Every(strArr, "php", "c++"))
}

/**输出
=== RUN   TestEvery
都存在: true
有一个不存在: false
都不存在: false
--- PASS: TestEvery (0.00s)
PASS
*/

// 批量查找，有一则返回true
func TestSome(t *testing.T) {
	strArr := []string{"go", "java", "c", "python"}
	fmt.Println("都存在:", funk.Some(strArr, "go", "c"))
	fmt.Println("至少一个存在:", funk.Some(strArr, "php", "go"))
	fmt.Println("都不存在:", funk.Some(strArr, "php", "c++"))
}

/***输出
=== RUN   TestSome
都存在: true
至少一个存在: true
都不存在: false
--- PASS: TestSome (0.00s)
PASS
*/

// 获取第一个元素 或 最后一个元素
func TestLastOrFirst(t *testing.T) {
	number := []int{10, 30, 12, 23}
	// 获取第一个元素
	fmt.Println("Head: ", funk.Head(number))
	// 获取最后一个元素
	fmt.Println("Last: ", funk.Last(number))
}

/**输出
=== RUN   TestLastOrFirst
Head:  10
Last:  23
--- PASS: TestLastOrFirst (0.00s)
PASS
*/

// 填充元素
func TestFill(t *testing.T) {
	// 初始化切片
	var data = make([]int, 3)
	fill, _ := funk.Fill(data, 100)
	fmt.Printf("fill: %v \n", fill)
	// 将所有值设置成2
	input := []int{1, 2, 3}
	result, _ := funk.Fill(input, 2)
	fmt.Printf("result: %v \n", result)
	var structData = make([]Student, 2)
	stuInfo, _ := funk.Fill(structData, Student{Name: "张三", Age: 18})
	fmt.Printf("stuInfo: %v \n", stuInfo)
}

/**输出
=== RUN   TestFill
fill: [100 100 100]
result: [2 2 2]
stuInfo: [{张三 18} {张三 18}]
--- PASS: TestFill (0.00s)
PASS
*/

type cus struct {
	Name string
	Age  int
	Home string
}

func TestJoin(t *testing.T) {
	a := []int64{1, 3, 5, 7}
	b := []int64{3, 7}
	// 任意类型切片交集
	join := funk.Join(a, b, funk.InnerJoin)
	fmt.Println("join = ", join)
	// 指定类型取交集
	joinInt64 := funk.JoinInt64(a, b, funk.InnerJoinInt64)
	fmt.Println("joinInt64 = ", joinInt64)
	// 自定义结构体交集
	sliceA := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "南京"},
	}
	sliceB := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "上海"},
	}
	res := funk.Join(sliceA, sliceB, funk.InnerJoin)
	fmt.Println("自定义结构体: ", res)
}

/**输出
=== RUN   TestJoin
join =  [3 7]
joinInt64 =  [3 7]
自定义结构体:  [{张三 20 北京}]
--- PASS: TestJoin (0.00s)
PASS
*/

// 取差集
func TestDiffSlice(t *testing.T) {
	a := []int64{1, 3, 5, 7}
	b := []int64{3, 7, 10}
	// 任意类型切片交集
	join := funk.Join(a, b, funk.OuterJoin)
	fmt.Println("OuterJoin = ", join)
	// 指定类型取交集
	joinInt64 := funk.JoinInt64(a, b, funk.OuterJoinInt64)
	fmt.Println("joinInt64 = ", joinInt64)
	// 自定义结构体交集
	sliceA := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "南京"},
	}
	sliceB := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "上海"},
	}
	res := funk.Join(sliceA, sliceB, funk.OuterJoin)
	fmt.Println("自定义结构体: ", res)
}

/**输出
=== RUN   TestDiffSlice
OuterJoin =  [1 5 10]
joinInt64 =  [1 5 10]
自定义结构体:  [{李四 22 南京} {李四 22 上海}]
--- PASS: TestDiffSlice (0.00s)
PASS
*/

func TestLeftAndRightJoin(t *testing.T) {
	a := []int64{10, 20, 30}
	b := []int64{30, 40}
	// 取出只在a，不在b的元素
	leftJoin := funk.Join(a, b, funk.LeftJoin)
	fmt.Println("只在a切片的元素: ", leftJoin)
	// 取出只在b，不在a的元素
	rightJoin := funk.Join(a, b, funk.RightJoin)
	fmt.Println("只在b切片的元素: ", rightJoin)
}

/**输出
=== RUN   TestLeftAndRightJoin
只在a切片的元素:  [10 20]
只在b切片的元素:  [40]
--- PASS: TestLeftAndRightJoin (0.00s)
PASS
*/

func TestDifferent(t *testing.T) {
	// 处理任意类型
	one := []interface{}{1, "go", 3.2, []int8{10}}
	two := []interface{}{2, "go", 3.2, []int{20}}
	oneRes, twoRes := funk.Difference(one, two)
	fmt.Println("oneRes: ", oneRes)
	fmt.Println("twoRes:  ", twoRes)

	// 只处理具体类型
	str1 := []string{"go", "php", "java"}
	str2 := []string{"c", "python", "java"}
	res, res1 := funk.DifferenceString(str1, str2)
	fmt.Println("res: ", res)
	fmt.Println("res1: ", res1)
}

/**输出
=== RUN   TestDifferent
oneRes:  [1 [10]]
twoRes:   [2 [20]]
res:  [go php]
res1:  [c python]
--- PASS: TestDifferent (0.00s)
PASS
*/

// 遍历切片
func TestForeachSlice(t *testing.T) {
	// 从左边遍历
	var leftRes []int
	funk.ForEach([]int{1, 2, 3, 4}, func(x int) {
		leftRes = append(leftRes, x*2)
	})
	fmt.Println("ForEach:", leftRes)
	// 从右边遍历
	var rightRes []int
	funk.ForEachRight([]int{1, 2, 3, 4}, func(x int) {
		rightRes = append(rightRes, x*2)
	})
	fmt.Println("ForEachRight:", rightRes)
}

/**输出
=== RUN   TestForeachSlice
ForEach: [2 4 6 8]
ForEachRight: [8 6 4 2]
--- PASS: TestForeachSlice (0.00s)
PASS
*/

// 除去第一个 或者 最后一个
func TestDelLastOrFirst(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	// 除去第一个，返回剩余元素
	fmt.Println(funk.Tail(a))
	// 除去最后一个，返回剩余元素
	fmt.Println(funk.Initial(a))
}

/**输出
=== RUN   TestDelLastOrFirst
[2 3 4 5 6 7]
[1 2 3 4 5 6]
--- PASS: TestDelLastOrFirst (0.00s)
PASS
*/

// 判断是否属于子集
func TestSubset(t *testing.T) {
	// 判断基础切片
	fmt.Println("是否属于子集:", funk.Subset([]int{1}, []int{1, 2}))
	// 判断自定义结构体切片
	subStu1 := []Student{{Name: "张三", Age: 18}}
	subStu2 := []Student{{Name: "张三", Age: 22}}
	allStu := []Student{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 22},
	}
	fmt.Println("subStu1: ", funk.Subset(subStu1, allStu))
	fmt.Println("subStu2: ", funk.Subset(subStu2, allStu))
	// 判断空切片是否属于另一切片子集
	fmt.Println("判断空集：", funk.Subset([]int{}, []int{1, 2}))
}

/**输出
=== RUN   TestSubset
是否属于子集: true
subStu1:  true
subStu2:  false
判断空集： true
--- PASS: TestSubset (0.00s)
PASS
*/

// 分组
func TestChunk(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	// 分组(每组2个)
	fmt.Println(funk.Chunk(a, 2))
	// 分组自定义结构体切片 (每组2个)
	stuList := []Student{
		{Name: "张三", Age: 18},
		{Name: "小明", Age: 20},
		{Name: "李四", Age: 22},
		{Name: "赵武", Age: 32},
		{Name: "小英", Age: 19},
	}
	fmt.Println(funk.Chunk(stuList, 2))
}

/**输出
=== RUN   TestChunk
[[1 2] [3 4] [5 6] [7]]
[[{张三 18} {小明 20}] [{李四 22} {赵武 32}] [{小英 19}]]
--- PASS: TestChunk (0.00s)
PASS
*/

// 分组
func TestChunkString(t *testing.T) {
	dataArray := []string{"1", "2", "3", "4", "5"}
	dataFlotArray := funk.ChunkStrings(dataArray, 2)

	for _, innerDataArray := range dataFlotArray {
		fmt.Println("innerDataArray:", innerDataArray)
	}
}

// 结构体转成Map
func TestToMap(t *testing.T) {
	bookList := []Book{
		{Id: 1, Name: "西游记"},
		{Id: 2, Name: "水浒传"},
		{Id: 3, Name: "三国演义"},
	}
	// 转成以Id为Key的Map
	fmt.Println("结果:", funk.ToMap(bookList, "Id"))
}

/**输出
=== RUN   TestToMap
结果: map[1:{1 西游记} 2:{2 水浒传} 3:{3 三国演义}]
--- PASS: TestToMap (0.00s)
PASS
*/

func TestMap(t *testing.T) {
	// 将切片最为map的key
	r := funk.Map([]int{1, 2, 3, 4}, func(x int) (int, string) {
		return x, "go"
	})
	fmt.Println("r=", r)
}

/**输出
=== RUN   TestMap
r= map[1:go 2:go 3:go 4:go]
--- PASS: TestMap (0.00s)
PASS
*/

// 把二维切片转成一维切片
func TestFlatMap(t *testing.T) {
	r := funk.FlatMap([][]int{{1}, {2}, {3}, {4}}, func(x []int) []int {
		return x
	})
	fmt.Printf("%#v\n", r)
}

/**输出
=== RUN   TestFlatMap
[]int{1, 2, 3, 4}
--- PASS: TestFlatMap (0.00s)
PASS
*/

func TestShuffle(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	// 打乱多次
	for i := 1; i <= 3; i++ {
		fmt.Println(fmt.Sprintf("第%v次打乱a", i), funk.Shuffle(a))
	}
}

/**输出
=== RUN   TestShuffle
第1次打乱a [5 4 2 6 7 0 3 1]
第2次打乱a [1 6 7 3 2 5 0 4]
第3次打乱a [5 1 7 6 0 4 2 3]
--- PASS: TestShuffle (0.00s)
PASS
*/

// 反转切片
func TestReverse(t *testing.T) {
	fmt.Println("ReverseInt:", funk.ReverseInt([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println("Reverse,任意类型:", funk.Reverse([]interface{}{1, "2", 3.02, 4, "5", 6}))
	fmt.Println("Reverse,Str:", funk.ReverseStrings([]string{"a", "b", "c", "d"}))
}

/**输出
=== RUN   TestReverse
ReverseInt: [6 5 4 3 2 1]
Reverse,任意类型: [6 5 4 3.02 2 1]
Reverse,Str: [d c b a]
--- PASS: TestReverse (0.00s)
PASS
*/

// 去重
func TestUniq(t *testing.T) {
	a := []int64{1, 2, 3, 4, 3, 2, 1}
	// 过滤整型类型
	fmt.Println("Uniq:", funk.Uniq(a))
	fmt.Println("UniqInt64:", funk.UniqInt64(a))
	b := []string{"php", "go", "c", "go"}
	// 过滤字符串类型
	fmt.Println("UniqString:", funk.UniqString(b))
}

/**输出
=== RUN   TestUniq
Uniq: [1 2 3 4]
UniqInt64: [1 2 3 4]
UniqString: [php go c]
--- PASS: TestUniq (0.00s)
PASS
*/

// 删除指定元素
func TestWithOut(t *testing.T) {
	// 删除具体元素
	b := []int{10, 20, 30, 40}
	without := funk.Without(b, 30)
	fmt.Println("删除30:", without)
	// 删除自定义结构体元素
	stuList := []Student{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 22},
		{Name: "范五", Age: 24},
	}
	res := funk.Without(stuList, Student{Name: "李四", Age: 22})
	fmt.Println("删除李四:", res)
}

/**输出
=== RUN   TestWithOut
删除30: [10 20 40]
删除李四: [{张三 18} {范五 24}]
--- PASS: TestWithOut (0.00s)
PASS
*/

// 获取map中所有的key
func TestMapKeys(t *testing.T) {
	res := map[string]int{
		"张三": 18,
		"李四": 20,
		"赵武": 25,
	}
	keys := funk.Keys(res)
	fmt.Printf("keys: %#v %T \n", keys, keys)
}

/**输出
=== RUN   TestMapKeys
keys: []string{"张三", "李四", "赵武"} []string
--- PASS: TestMapKeys (0.00s)
PASS
*/

// 获取map中的所有values
func TestMapValues(t *testing.T) {
	res := map[string]int{
		"张三": 18,
		"李四": 20,
		"赵武": 25,
	}
	values := funk.Values(res)
	fmt.Printf("values: %#v %T \n", values, values)
}

/**输出
=== RUN   TestMapValues
values: []int{18, 20, 25} []int
--- PASS: TestMapValues (0.00s)
PASS
*/

func TestGetToSlice(t *testing.T) {
	userList := []User{
		{
			Name: "张三",
			Home: struct {
				City string
			}{"北京"},
		},
		{
			Name: "小明",
			Home: struct {
				City string
			}{"南京"},
		},
	}
	// 取一层
	names := funk.Get(userList, "Name")
	fmt.Println("names:", names)
	// 取其他层
	homes := funk.Get(userList, "Home.City")
	fmt.Println("homes:", homes)
}

/**输出
=== RUN   TestGetToSlice
names: [张三 小明]
homes: [北京 南京]
--- PASS: TestGetToSlice (0.00s)
PASS
*/

// 比较
func TestIsEqual(t *testing.T) {
	// 对比字符串
	fmt.Println("对比字符串:", funk.IsEqual("a", "a"))
	// 对比int
	fmt.Println("对比int:", funk.IsEqual(1, 1))
	// 对比float64
	fmt.Println("对比float64:", funk.IsEqual(float64(1), float64(1)))
	// 对比结构体
	stu1 := Student{Name: "张三", Age: 18}
	stu2 := Student{Name: "张三", Age: 18}
	fmt.Println("对比结构体:", funk.IsEqual(stu1, stu2))
}

/**输出
=== RUN   TestIsEqual
对比字符串: true
对比int: true
对比float64: true
对比结构体: true
--- PASS: TestIsEqual (0.00s)
PASS
*/

// 判断类型是否一样
func TestIsType(t *testing.T) {
	var a, b int8 = 1, 2
	fmt.Println("A: ", funk.IsType(a, b))
	c := 3
	d := "3"
	fmt.Println("B:", funk.IsType(c, d))
}

/**输出
=== RUN   TestIsType
A:  true
B: false
--- PASS: TestIsType (0.00s)
PASS
*/

// 判断是否是array|slice
func TestCollect(t *testing.T) {
	// slice
	a := []int{1, 2, 3}
	// 字符串
	b := "str"
	// 自定义结构体切片
	c := []Student{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 18},
	}
	// 数组类型
	d := [2]string{"c", "go"}
	fmt.Println("a:", funk.IsCollection(a))
	fmt.Println("b:", funk.IsCollection(b))
	fmt.Println("c:", funk.IsCollection(c))
	fmt.Println("d:", funk.IsCollection(d))
}

/**输出
=== RUN   TestCollect
a: true
b: false
c: true
d: true
--- PASS: TestCollect (0.00s)
PASS
*/

// 判断是否为空
func TestIsEmpty(t *testing.T) {
	// 空结构体
	fmt.Println("空结构体", funk.IsEmpty([]int{}))
	// 空字符串
	fmt.Println("空字符串:", funk.IsEmpty(""))
	// 判断数字0
	fmt.Println("0:", funk.IsEmpty(0))
	// 判断字符串'0'
	fmt.Println("'0':", funk.IsEmpty("0"))
	// nil
	fmt.Println("nil:", funk.IsEmpty(nil))
}

/**输出
=== RUN   TestIsEmpty
空结构体 true
空字符串: true
0: true
'0': false
nil: true
--- PASS: TestEmpty (0.00s)
PASS
*/

// 数字型转浮点型
func TestToFloat64(t *testing.T) {
	// int to float64
	d1, _ := funk.ToFloat64(10)
	fmt.Printf("d1 = %v %T \n", d1, d1)
	//@会失败
	d2, err := funk.ToFloat64("10")
	fmt.Printf("d2 = %v %v \n", d2, err)
}

/**输出
=== RUN   TestToFloat64
d1 = 10 float64
d2 = 0 false
--- PASS: TestToFloat64 (0.00s)
PASS
*/

// 返回任一类型的类型切片
func TestSliceOf(t *testing.T) {
	a := []int{10, 20, 30}
	// []int 转成 [][]int
	fmt.Printf("%v %T \n", funk.SliceOf(a), funk.SliceOf(a))
	// string 转成 []string
	fmt.Printf("%v %T \n", funk.SliceOf("go"), funk.SliceOf("go"))
}

/**输出
=== RUN   TestSliceOf
[[10 20 30]] [][]int
[go] []string
--- PASS: TestSliceOf (0.00s)
PASS
*/

// 将自定长度字符串生成切片
func TestShard(t *testing.T) {
	token := "Hello,World"
	shard := funk.Shard(token, 1, 5, false)
	shard1 := funk.Shard(token, 1, 5, true)
	fmt.Println("shard: ", shard)
	fmt.Println("shard1: ", shard1)
	shard2 := funk.Shard(token, 2, 5, false)
	shard22 := funk.Shard(token, 2, 5, true)
	fmt.Println("shard2: ", shard2)
	fmt.Println("shard22: ", shard22)
}

/**
=== RUN   TestShard
shard:  [H e l l o Hello,Word]
shard1:  [H e l l o ,Word]
shard2:  [He ll o, Wo rd Hello,Word]
shard22:  [He ll o, Wo rd ]
--- PASS: TestShard (0.00s)
PASS
*/

func TestMax(t *testing.T) {
	// 求最大int
	fmt.Println("MaxInt:", funk.MaxInt([]int{30, 10, 8, 11}))
	// 求最大浮点数
	fmt.Println("MaxFloat64:", funk.MaxFloat64([]float64{10.2, 11.0, 8.03}))
	// 求最大字符串
	fmt.Println("MaxString:", funk.MaxString([]string{"a", "d", "c", "b"}))
}

/**输出
=== RUN   TestMaxInt
MaxInt: 30
MaxFloat64: 11
MaxString: d
--- PASS: TestMaxInt (0.00s)
PASS
*/

func TestMin(t *testing.T) {
	// 求最小int
	fmt.Println("MinInt:", funk.MinInt([]int{30, 10, 8, 11}))
	// 求最小浮点数
	fmt.Println("MinFloat64:", funk.MinFloat64([]float64{10.2, 11.0, 8.03}))
	// 求最小字符串
	fmt.Println("MinString:", funk.MinString([]string{"a", "d", "c", "b"}))
}

/**输出
=== RUN   TestMin
MinInt: 8
MinFloat64: 8.03
MinString: a
--- PASS: TestMin (0.00s)
PASS
*/

// 求和
func TestSum(t *testing.T) {
	// 整型
	a := []int{5, 10, 15, 20}
	fmt.Println("int sum:", funk.Sum(a))
	// 浮点型
	b := []float64{5.11, 2.23, 3.31, 0.32}
	fmt.Println("float64 sum:", funk.Sum(b))
}

/**输出
=== RUN   TestSum
int sum: 50
float64 sum: 10.97
--- PASS: TestSum (0.00s)
PASS
*/

// 求乘积
func TestProduct(t *testing.T) {
	// 整型
	a := []int{2, 3, 4, 5}
	fmt.Println("int Product:", funk.Product(a))
	// 浮点型
	b := []float64{1.1, 1.2, 1.3, 1.4}
	fmt.Println("float64 Product:", funk.Product(b))
}

/**输出
=== RUN   TestProduct
int Product: 120
float64 Product: 2.4024
--- PASS: TestProduct (0.00s)
PASS
*/

// 生成随机数
func TestRandom(t *testing.T) {
	for i := 1; i <= 3; i++ {
		// 生成任意数字类型
		fmt.Println(funk.RandomInt(1, 100))
	}
}

/**输出
=== RUN   TestRandom
24
79
21
--- PASS: TestRandom (0.00s)
PASS
*/

// 生成随机字符串
func TestRandomString(t *testing.T) {
	for i := 1; i <= 3; i++ {
		// 从默认字符串生成
		fmt.Println("从默认字符串生成:", funk.RandomString(i))
		// 从指的字符串生成
		fmt.Println("从指定字符串生成:", funk.RandomString(i, []rune{'您', '好', '北', '京'}))
	}
}

/**输出
=== RUN   TestRandomString
从默认字符串生成: B
从指定字符串生成: 京
从默认字符串生成: Ln
从指定字符串生成: 好北
从默认字符串生成: Dsc
从指定字符串生成: 您北京
--- PASS: TestRandomString (0.00s)
PASS
*/

// 三元运算
//
//goland:noinspection GoBoolExpressions
func TestShortIf(t *testing.T) {
	fmt.Println("10 > 5 :", funk.ShortIf(10 > 5, 10, 5))
	fmt.Println("10.0 == 10 : ", funk.ShortIf(10.0 == 10, "yes", "no"))
	fmt.Println("'a' == 'b' : ", funk.ShortIf('a' == 'b', "equal chars", "unequal chars"))
}

/**输出
=== RUN   TestShortIf
10 > 5 : 10
10.0 == 10 :  yes
'a' == 'b' :  unequal chars
--- PASS: TestShortIf (0.00s)
PASS
*/

func TestIntersect(t *testing.T) {
	f := Foo{
		ID:        1,
		FirstName: "Gilles",
		LastName:  "Fabio",
		Age:       70,
	}

	b := Foo{
		ID:        2,
		FirstName: "Florent",
		LastName:  "Messa",
		Age:       80,
	}

	b2 := Foo{
		ID:        2,
		FirstName: "Florent",
		LastName:  "Messa",
		Age:       80,
	}

	fmt.Println("object set:", funk.ToSet([]Foo{f, b, b2}))
	fmt.Println("int set:", funk.ToSet([4]int{1, 1, 2, 2}))
}

/**输出
=== RUN   TestIntersect
object set: map[{1 Gilles Fabio 70}:{} {2 Florent Messa 80}:{}]
int set: map[1:{} 2:{}]
--- PASS: TestIntersect (0.00s)
PASS
*/

func TestFilter(t *testing.T) {
	result := funk.Filter([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("偶数:", result)
}

/**
=== RUN   TestFilter
偶数: [2 4]
--- PASS: TestFilter (0.00s)
PASS
*/

func TestReduce(t *testing.T) {
	// Using operation runes. '+' and '*' only supported.
	r1 := funk.Reduce([]int{1, 2, 3, 4}, '+', float64(0)) // 10
	r2 := funk.Reduce([]int{1, 2, 3, 4}, '*', 1)          // 24

	// Using accumulator function
	r3 := funk.Reduce([]int{1, 2, 3, 4}, func(acc float64, num int) float64 {
		return acc + float64(num)
	}, float64(0)) // 10

	r4 := funk.Reduce([]int{1, 2, 3, 4}, func(acc string, num int) string {
		return acc + fmt.Sprint(num)
	}, "") // "1234"

	fmt.Println("plus result: ", r1)
	fmt.Println("multiply result: ", r2)
	fmt.Println("plus result: ", r3)
	fmt.Println("string concat result: ", r4)
}

/**
=== RUN   TestReduce
plus result:  10
multiply result:  24
plus result:  10
string concat result:  1234
--- PASS: TestReduce (0.00s)
PASS
*/

func TestFind(t *testing.T) {
	result := funk.Find([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("第一个偶数:", result)
}

/**
=== RUN   TestFind
第一个偶数: 2
--- PASS: TestFind (0.00s)
PASS
*/

func TestMapAllInOne(t *testing.T) {
	r1 := funk.Map([]int{1, 2, 3, 4}, func(x int) int {
		return x * 2
	}) // []int{2, 4, 6, 8}
	fmt.Println("slice -> slice: ", r1)

	r2 := funk.Map([]int{1, 2, 3, 4}, func(x int) string {
		return "Hello"
	}) // []string{"Hello", "Hello", "Hello", "Hello"}
	fmt.Println("slice -> slice: ", r2)

	r3 := funk.Map([]int{1, 2, 3, 4}, func(x int) (int, int) {
		return x, x
	}) // map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	fmt.Println("slice -> map: ", r3)

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	r4 := funk.Map(mapping, func(k int, v string) int {
		return k
	}) // []int{1, 2}
	fmt.Println("map -> slice: ", r4)

	r5 := funk.Map(mapping, func(k int, v string) (string, string) {
		return fmt.Sprintf("%d", k), v
	}) // map[string]string{"1": "Florent", "2": "Gilles"}
	fmt.Println("map -> map: ", r5)
}

/**
=== RUN   TestMapAllInOne
slice -> slice:  [2 4 6 8]
slice -> slice:  [Hello Hello Hello Hello]
slice -> map:  map[1:1 2:2 3:3 4:4]
map -> slice:  [1 2]
map -> map:  map[1:Florent 2:Gilles]
--- PASS: TestMapAllInOne (0.00s)
PASS
*/

func TestFlatMap2(t *testing.T) {
	r1 := funk.FlatMap([][]int{{1, 2}, {3, 4}}, func(x []int) []int {
		return append(x, 0)
	}) // []int{1, 2, 0, 3, 4, 0}
	fmt.Println("map -> slice: ", r1)

	mapping := map[string][]int{
		"Florent": {1, 2},
		"Gilles":  {3, 4},
	}

	r2 := funk.FlatMap(mapping, func(k string, v []int) []int {
		return v
	}) // []int{1, 2, 3, 4}
	fmt.Println("map -> slice: ", r2)
}

func TestGetOrElse(t *testing.T) {
	str := "hello world"
	fmt.Println(funk.GetOrElse(&str, "foobar")) // string{"hello world"}
	fmt.Println(funk.GetOrElse(str, "foobar"))  // string{"hello world"}
	fmt.Println(funk.GetOrElse(nil, "foobar"))  // string{"foobar"}
}

/**
=== RUN   TestGetOrElse
hello world
hello world
foobar
--- PASS: TestGetOrElse (0.00s)
PASS

*/

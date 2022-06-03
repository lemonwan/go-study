package testing

import (
	"fmt"
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestEncode(t *testing.T) {
	s := "hello，世界"
	t.Log(len(s))
	t.Log(utf8.RuneCountInString(s))

	fmt.Println("-------")
	for i, v := range s {
		t.Log(i, v)
		t.Log(string(v))
	}

	bytes := []byte(s)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		ch := fmt.Sprintf("%c", r)
		t.Log(ch)
	}
}

func TestArr(t *testing.T) {
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := new([3]int)
	t.Logf("%T", arr1)
	t.Logf("%T", arr2)
}

type treeNode struct {
	value       string
	left, right *treeNode
}

func (node *treeNode) traverse() {
	if node.left != nil {
		node.left.traverse()
	}
	if node.right != nil {
		node.right.traverse()
	}
	fmt.Println(node.value)
}

func createTreeNode() treeNode {
	root := treeNode{"A", nil, nil}
	root.left = &treeNode{value: "B"}
	root.right = &treeNode{value: "C"}
	root.left.left = &treeNode{value: "D"}
	root.left.right = &treeNode{value: "E"}
	root.right.left = &treeNode{value: "F"}
	root.right.right = &treeNode{value: "G"}
	return root
}

func TestPrintNode(t *testing.T) {
	node := createTreeNode()
	node.traverse()
}

type A struct {
	Face int
}

func (a *A) f() {
	fmt.Println("HI,", a.Face)
}

// 自定义类型不会继承原有类型的方法（自定义类型不会拥有原基础类型所附带的方法），但接口方法或组合类型的内嵌元素则保留原有的方法。
type NA A

// 类型别名
type NB = A

func TestA(t *testing.T) {
	a := A{Face: 10}
	a.f()

	// na := NA{Face: 200}
	// na.f()

	nb := NB{Face: 300}
	nb.f()
}

type Student struct {
	name string "学生的名字"
	Age  int    "学生的年龄"
	Room int    `json:"room,omitempty"`
}

func TestReflect(t *testing.T) {
	s := Student{"wan", 20, 1}
	t.Log(reflect.TypeOf(s).Field(0).Tag)
	t.Log(reflect.TypeOf(s).Field(1).Tag)
	t.Log(reflect.TypeOf(s).Field(2).Tag)
}

type MyInt int

type I interface {
	Print()
	AddAndPrint(int)
	Add()
}

func (m MyInt) Print() {
	fmt.Println(m)
}

func (m MyInt) AddAndPrint(a int) {
	fmt.Println(int(m) + a)
}

func (m *MyInt) Add() {
	fmt.Println(m)
}

func TestSelfType(t *testing.T) {
	a := MyInt(10)
	a.Print()
	a.Add()

	// 选择器
	MyInt.AddAndPrint(a, 5)

	a.AddAndPrint(5)

}

func interHandle(i I) {
	i.AddAndPrint(10)
}

func TestInterface(t *testing.T) {
	i := MyInt(10)
	// 指针对象即可以调用值方法，也可以调用指针方法
	var in I = &i
	interHandle(in)
}

//  类型和作用在它上面定义的方法必须在同一个包中定义
// 以下程序会编译错误
//func (i int) Print() {
//
//}

type People struct {
	Name   string
	gender string
	Age    int
}

type OtherPeople struct {
	People
}

func (p People) PeInfo() {
	fmt.Println("People ", p.Name, ": ", p.Age, "岁, 性别:", p.gender)
}

func (p *People) PeName(name string) {
	fmt.Println("old name:", p.Name)
	p.Name = name
	fmt.Println("new name:", p.Name)
}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("%T\n", a)
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(i, ":", m.Name, m.Type)
	}
}

func TestFuncIn(t *testing.T) {
	p := OtherPeople{
		People{
			Name:   "wan",
			gender: "male",
			Age:    20,
		},
	}
	p.PeInfo()
	p.PeName("lemon")
	methodSet(p)
	methodSet(&p)
}

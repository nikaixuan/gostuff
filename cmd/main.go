package main

import (
	"fmt"
	"goproject/pkg/leetcode/binaryTree"
	"goproject/pkg/leetcode/linkedList"
	"strings"
)

type down struct {
	a *string
	b string
}

type up struct {
	c *string
	d string
}

func main() {
	//a := &card.Card{Color: "red"}
	//b := a
	//a.Number = "123"
	//fmt.Println(a, b)
	//node := datastructure.New(5)
	//node.Print()

	//array := []int{1,2,3,4,5}
	//b := array[0:3]
	//ints := b[0:2]
	//fmt.Println(array, ints)

	//algorithms.EightQueens()

	//str := "123456"
	//if _, err := strconv.Atoi(str); err != nil {
	//	// do stuff, in case str can not be converted to an int
	//}
	//var slice []int // empty slice
	//for _, digit := range str {
	//	slice = append(slice, int(digit)-int('0')) // build up slice
	//}
	//fmt.Println(slice)
	//[["a","b"],["b","c"]]
	//[2.0,3.0]
	//[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]

	//slice := []int{1,2,3}
	//m := map[int]int{1:2}
	//tn := leetcode.ListNode{Val: 1, Next: &leetcode.ListNode{Val: 2}}
	//testFunc(slice, m, tn)
	//fmt.Println(slice, m, tn)
	fmt.Println(strings.Count("123", ""))

}

func TestBasicGolang() {
	d := []int{1, 2, 3}
	f := map[string]int{"a": 1}
	fmt.Println("printing 1")
	fmt.Println(d, f)
	fmt.Println("printing 2")

	fmt.Println("printing 3")
	e := make([]int, 5)
	e1 := make([]int, 0, 5)
	var e2 []int
	fmt.Println(cap(e), len(e))
	fmt.Println(cap(e1), len(e1))
	e = append(e, 1)
	e1 = append(e1, []int{1, 2, 3}...)
	e2 = append(e2, []int{1, 2, 3, 4, 5}...)
	for k, v := range e1 {
		fmt.Println(k, v)
	}
	fmt.Println(cap(e), len(e), e)
	fmt.Println(cap(e1), len(e1), e1)
	fmt.Println(cap(e2), len(e2), e2)
}

func TestRefAndPointer() {
	a := 1
	fmt.Println(a)
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b)
	*b = 8
	fmt.Println(a)
	c := b
	*c = 10

	fmt.Println(a)
	fmt.Println(&a)
	changeNumber(&a)
	fmt.Println(a)
	z := new(string)
	fmt.Println("------------")
	fmt.Printf("type is %T", z)

}

func changeNumber(n *int) {
	fmt.Println(n)
	*n = 20
}

func buildTree() *binaryTree.TreeNode {
	t1 := &binaryTree.TreeNode{Val: 1}
	t2 := &binaryTree.TreeNode{Val: 2}
	t3 := &binaryTree.TreeNode{Val: 3}
	t4 := &binaryTree.TreeNode{Val: 4}
	t5 := &binaryTree.TreeNode{Val: 5}
	t6 := &binaryTree.TreeNode{Val: 6}
	t1.Left = t2
	t1.Right = t5
	t2.Left = t3
	t2.Right = t4
	t5.Right = t6
	return t1
}

func testFunc(slice []int, testMap map[int]int, testStruct linkedList.ListNode) {
	a := slice
	b := testMap
	c := testStruct
	a[0] = 5
	b[5] = 55
	c.Val = 2
	slice = append(slice, 6)
	fmt.Println(slice, testMap, testStruct)
}

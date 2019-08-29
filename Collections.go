package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//数组
func arrayIter() {
	var arr [5]int
	arr1 := [...]int{2, 3, 4, 5, 30}
	fmt.Println(arr, arr1)

	//for i := 0;i < len(arr1);i++  {
	//	fmt.Println(i,arr1[i])
	//}
	//
	//for i,v := range arr1{
	//	fmt.Println("数组：",i,v)
	//}
	changeArray(arr1)
	fmt.Println(arr1)

	changeArr(&arr1)
	fmt.Println(arr1)

}

//函数传数组是传值类型
func changeArray(arr [5]int) {
	arr[0] = 25
}

//必须传指针才能改变原数组数据
func changeArr(arr *[5]int) {
	arr[1] = 41
}

//切片是
func sliceDemo() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println("arr[2:6]", s)
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])

	fmt.Println("Reslice")
	fmt.Println(s)
	s = s[2:]
	fmt.Println(s)
	s = s[:4] //是原slice的view，因此可以取到原slice有但是片段少了的一部分，叫扩展slice
	fmt.Println(s)
}

func updateSlice() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := append(arr, 8)
	s2 := append(s1, 9)
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

func createSlice() {
	var s []int //Zero value for slice is nil
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	pringSlice(s1)

	s2 := make([]int, 10, 32)
	pringSlice(s2)

	//copy
	copy(s2, s1)
	pringSlice(s2)

	//delete
	s2 = append(s2[:3], s2[4:]...)
	pringSlice(s2)

	s2 = s2[1:]
	pringSlice(s2)
	s2 = s2[:len(s2)-1]
	pringSlice(s2)

}

func pringSlice(s []int) {
	fmt.Printf("%v,leng=%d,cap=%d\n", s, len(s), cap(s))
}

func mapDemo() {
	m := map[string]string{
		"name":    "hand",
		"couse":   "golang",
		"site":    "imooc",
		"quality": "good",
	}
	m2 := make(map[string]string) //m2 == empty map

	var m3 map[string]int //m3 == nil

	fmt.Println(m, m2, m3)

	//无序的
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m["name"])

	//判断是否存在
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	//delete
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}

//最长不重复字符子串
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func stringDeal() {
	s := "Yes我爱韩丁"
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println()
	for i, ch := range s { //ch is a rune (int32)
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)
	}

	strings.Compare("abc", "abc")
	//Fields,Split,Join
	//Contains,Index
	//Tolower ToUpper
	//Trim,TrimRight,TrimLeft
}

func main() {
	//3-1
	arrayIter()

	//3-2
	//sliceDemo()
	//updateSlice()
	//createSlice()

	//3-3
	//mapDemo()

	//3-4-5
	//fmt.Println("sub:",lengthOfNonRepeatingSubStr("adfdafasd"))
	//3-6
	stringDeal()

}

package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var aa = 3
var ss = "kkk"

var (
  kk = true
  bb = 3.1
)

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,ss,kk,bb)
	euler()
	triangle()
	consts()
	enums()
}

func enums(){
	const(
		cpp = iota
		_
		java
		python
		golang
	)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb ,pb)
}

func consts(){
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(filename, c)
}

func triangle(){
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue(){
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	var a,b,c,s = 3,4,true, "def"
	fmt.Println(a,b,c,s)
}

func variableShorter() {
	a,b,c,s := 3,4,true,"def"
	b = 5
	fmt.Println(a,b,c,s)
}

func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
}

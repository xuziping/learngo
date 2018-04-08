package main

import (
	"io/ioutil"
	"fmt"
)

func main() {

	const filename = "abc.txt"

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(grade(0), grade(59), grade(99))
}

func grade(score int) string {
	g := ""
  	switch{
	case score < 60:
	    g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
  	}
	return g;
}


package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with Files in golang")
	content := "This need to go in file. "

	file,err := os.Create("./myfile.txt")

	checknilerror(err)

	length, err := io.WriteString(file,content)

	checknilerror(err)

	fmt.Println("Length : ",length)

	defer file.Close()
	Readfile("./myfile.txt")
}

func Readfile(filename string)  {
	databyte, err := ioutil.ReadFile(filename) // it is deprecated 
	// use os.ReadFile
	checknilerror(err)
	fmt.Println("Text data inside the file is:-\n",string(databyte))
}

func checknilerror(err error)  {
	if err != nil {
		panic(err)
	}
}

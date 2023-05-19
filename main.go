package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type ArrayTest struct {
	Id   int
	Name string
}

type CsvTest struct {
	Id   int
	Name string
}

func main() {
	testArray := []ArrayTest{
		{Id: 1, Name: "test1"},
		{Id: 2, Name: "test2"},
		{Id: 3, Name: "test3"},
		{Id: 4, Name: "test4"},
		{Id: 5, Name: "test5"},
	}
	var csvTest []CsvTest
	for _, v := range testArray {
		csvDate := CsvTest(v)

		csvTest = append(csvTest, csvDate)
	}
	fmt.Println(csvTest)
	file, _ := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	gocsv.MarshalFile(csvTest, file)
}

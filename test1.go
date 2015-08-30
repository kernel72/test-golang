package main

import (
	"fmt"
	"github.com/kernel72/test-golang/subpackage"
)

func main() {
	somethingWithSubpackagePackages()
	somethingWithArrays()
	somethingWithMaps()
	somethingWithStructs()
	somethingWithGoRoutines()
}

func somethingWithSubpackagePackages(){
	fmt.Println(subpackage.AnotherExternalFunc())
}

func somethingWithArrays(){
	initial := []int{1,2,3,4,5,6}
	arr := make([]int, len(initial))
	for i, value := range initial{
		arr[i] = value + 15
	}
	fmt.Println(arr)
}

func somethingWithMaps(){
	myMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(myMap)

	mapOfMap := map[string]map[string]string{
		"key1": {
			"subkey1": "value1",
		},
		"key2": {
			"subkey2": "value2",
		},
	}
	fmt.Println(mapOfMap)

	key1 := "key1"

	if _, has := mapOfMap[key1]; has {
		fmt.Printf("%s present\n", key1)
	}
}


type SomeStruct struct {
	fieldInt int
	fieldString string
	fieldMap map[string]string
}

type ChildStruct struct {
	SomeStruct
	childField string
}

func (self *SomeStruct) getFieldMap() map[string]string {
	return self.fieldMap
}

func (self *ChildStruct) getFieldString() string {
	return self.fieldString
}

func somethingWithStructs(){
	str := ChildStruct{SomeStruct{1, "Odin2", map[string]string{"key133": "string"}}, "child Value"}
	fmt.Println(str.getFieldString())
}

func somethingWithGoRoutines(){

	a := []int{1,2,3,4,5,6,7,8,9,10}
	c := make(chan int)

	doPrint := func (routineNum int, c chan int){
		fmt.Printf("Start go routine %d\n", routineNum)

		for i:=0; i< 100; i++ {
			fmt.Println(i)
		}
		c <- routineNum;


	}

	go doPrint(1, c)

	b := make([]int, len(a))

	for i, val := range a {
		b[i] = val + 10
	}
	go doPrint(2, c)
	go doPrint(3, c)

	for i := 0 ; i < 3; i++ {
		chanNum := <- c
		fmt.Printf("Go %d routine complete\n", chanNum)
	}

	fmt.Println("After Routine")


}


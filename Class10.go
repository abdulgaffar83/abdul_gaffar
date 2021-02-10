package main

import "fmt"

func main() {

//long hand way 

//var students [3]string

//fixed length
//same type
//0,1,2,3,4

//students [0]="gaffar"
//students [1]="rakib"
//students [2]="sakil"

//fmt.Println(students[0])

//fmt.Println(students)
//fmt.Println(len(students))

//short hand way

//students := [3]string{"gaffar", "rakib", "sakil"}
students := [...]string{"gaffar", "rakib", "sakil", "bithy"}
//fmt.Println(students)
fmt.Println(students, len(students))

}


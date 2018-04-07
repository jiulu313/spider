package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	goal int //分数
}

//比较两个人的年龄大小，并返回年龄大的那个人，并返回年龄差
func older(p1, p2 Person) (Person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}

	return p2, p2.age - p1.age
}

func main() {

	p1 := Person{name: "tom", age: 23}
	p2 := Person{name: "jim", age: 45}

	fmt.Println(older(p1, p2))

	s1 := Student{Person{"小明",23},  89}
	s2 := Student{Person{name:"小华",age : 36},  89}
	fmt.Println(s1)
	fmt.Println(s2)




}

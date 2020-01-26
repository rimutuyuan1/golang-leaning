package main

import (
	"fmt"
	"strings"
)

//duck typing 描述事物的外部行为而非内部结构
//"像鸭子走路，像鸭子叫，就是鸭子"
//go属于结构化类型系统，类似duck typing

type Say interface {
	Say(url string) string
}

type Eat interface {
	Eat(food string) string
	EatBone(str string) bool
}

type SayAndEat interface {
	Say
	Eat
	run(str string)
}

type Dog struct {
}

func (d *Dog) Say(word string) string {
	return word + "汪汪汪！"
}

func (d Dog) Eat(food string) string {
	return food + "真好吃！"
}

func EatBone(str string) bool {
	if strings.Compare(str, "小狗") != 0 {
		return false
	} else {
		return true
	}
}

func main() {
	//TODO Method call 'sae.EatBone(str)' may lead to nil pointer dereference
	str := "小狗"
	var sae SayAndEat
	bone := sae.EatBone(str)
	fmt.Println(bone)
}

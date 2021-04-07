package main

import "fmt"

/*func main() {
	//声明一个函数 要求输出个人信息
	fmt.Println("输入你的个人信息")
	hangge()
    var age int
	var name string
	fmt.Scan(&name,&age)
	fmt.Println(name,age )

}
func hangge() {
	fmt.Println("你的名字年龄和班级")
}*/

/*func add() {
	var a int
	/*for i:=0;i<100;i++{
		a=a+i
	}*/
/*	fmt.Println(a)
}*/
/*func main() {
add()
}*/



/*func score(n int ){
	if n<60 {
		fmt.Println("不及格，你个废物，dsidkjsakshaklfhkjlhafkjlhjlkdhalfkaf")

	}else{
		fmt.Println("垃圾，才这么点分")
	}
}
func main() {
	var score1 int
	fmt.Scan(&score1)
	score(score1)
}*/

func hanggegge(name string,class string,age int)  {
	fmt.Println(name ,class,age)

}
func main() {
	var name,class string
	var age int
	fmt.Scan(&name,&class,&age)
	hanggegge(name,class,age)
}

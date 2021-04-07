/*package main

import "fmt"

func qwaszx(){
	for i:=1;i<=100;i++{
		if i %3==0||i%7==0{
			fmt.Println(i)
		}
	}
}
func main() {
	//声明一个函数，输出1--100之间能被3或7整除的数据
	qwaszx()
}*/


//-----------------------------------------------------------

/*
//错误示范
package main

import "fmt"

func qwaszx(n int){
	for i:=1;i<=n;i++{
		if i %3==0||i%7==0{
			fmt.Println(i)
		}
	}
}
func main() {
	//声明一个函数，输出1--n之间能被3或7整除的数据
	var x int
	fmt.Scan(x)
	qwaszx(x)
}

*/






/*//反复套娃
package main

import "fmt"

func main() {
	npc1()
	npc2()
}
func npc1() {
	fmt.Println("航")
	npc2()
}
func npc2() {
	fmt.Println("哥")
	npc1()
}*/






/*//循环递增
package main

import "fmt"

func main() {
	for{
		var sum int
		sum=sum
		fmt.Println(sum)
	}
}*/







//函数返回值问题
/*    func 函数名 （形参列表） （返回值列表）{
    	函数体
}*/







/*//练习：声明函数计算1---n之间奇数的和，并把结果返回主函数
package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输入一个大于0 的数值")
	fmt.Println(sum(10))
}
func sum(n int)(int) {
	var sum int
	for i:=1;i<=n;i+=2{
		sum +=i
	}
		return sum
}*/








//声明一个用于实现除法运算的函数，要求判断输入数据是否正确
/*package main

import "fmt"

func main() {
   var a,b float32
   var fuhao string
   for{fmt.Println("请通过键盘输入:(格式为：数字1  enter 符号 enter 数字2 enter)，结尾0表示计算结束，并可以开始新一轮的计算")
       fmt.Scan(&a,&fuhao,&b)
       fmt.Println(jisuanji(a ,fuhao, b))
   }
}
func jisuanji(a float32, fuhao string,b float32)(float32) {     //  func 函数名 （形参列表） （返回值列表）
	//var a, b float32
	//var fuhao string
	//for {
		//fmt.Println("输入:(格式为：数字1  enter 符号 enter 数字2 enter)")
		//fmt.Scan(&a, &fuhao, &b) //输入格式：以空格为分隔符，例如：9 + 9
		switch fuhao {
		case "+":
			fmt.Println(a, fuhao, b, "=", a+b)
		case "-":
			fmt.Println(a, fuhao, b, "=", a-b)
		case "*":
			fmt.Println(a, fuhao, b, "=", a*b)
		case "/":
			for{if b==0{
				fmt.Println("在计算除法时，除数不能为0，请重新输入除数b")
				fmt.Scan(&b)
			}else{
				fmt.Println(a, fuhao, b, "=", a/b)
				break
			}}
		default:
			fmt.Println("error!")

		}
		return 0
	//}
}

*/






//复习  数组和切片
//递归  汉诺塔

/*package main

import "fmt"

func print(n int,x rune,y rune)(){
	fmt.Printf("圆盘 %d 从 %c 到 %c\n",n,x,y)
}

func move(n int,a rune,b rune,c rune)(){
	if n==1{
		print(n,a,c)
	}else {
		move(n-1,a,c,b);
		print(n,a,c);
		move(n-1,b,a,c)
	}
}

func main() {
	var n int;
	fmt.Println("请输入圆盘的个数: ");
	fmt.Scanf("%d",&n);
	move(n,'x','y','z')
}

*/
package main

import "fmt"

func main() {
ex1()
// ex2()

}
func ex1(){
	var n int
	fmt.Println("enter n:")
	fmt.Scan(&n)
	for i:=0;i<=10;i++{
	fmt.Printf("%d*%d=%d ",i,n,n*i)
}	
}
func ex2(){
	for i:=10;i>=1;i--{
		fmt.Println(i)
	}
}
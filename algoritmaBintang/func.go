package algoritmaBintang

import (
	"fmt"
)

func bintang1(number int){
	for i := 1; i <= number; i++ {
		for x := 1; x <=i; x++ { 
			fmt.Print("*")
	}
	fmt.Println("")
  }
}

func bintang2 (number int){
	a:=number
	for i := 1; i <= number; i++ {
		for x := 1; x<=number; x++ {
			if x<a{                  
				fmt.Print(" ")
			}else {
				fmt.Print("*")
			}
		} 
		a=a-1 
		fmt.Println("")   
	}
}

func bintang3(number int) {
	for i := 1; i <= number; i++ {
		for x := number; x >=i; x-- { 
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func bintang4(number int) {
	for i:=1;i<=number;i++{
		for x:=1;x<=number;x++{
			if x>=i{
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

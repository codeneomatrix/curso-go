package main

import "math"
import "fmt"

func main() {
	fmt.Println("hello world")

	fmt.Println("go"+"lang")

	fmt.Println("7.0/3.0", 7.0/3.0)

	fmt.Println(true && false)

	fmt.Println(true || false)

	fmt.Println(!true)

// variables

	var a, b = 1,2
	fmt.Println(a,b)

	f:= "texto"

	fmt.Println(f)

//constantes

	const n= 5000000000
	fmt.Println(int64(n))
	fmt.Println(math.Sin(n))

//control
	i:=1
	for i < 3 {
		fmt.Println(i)
		i=i+1
	}

	
	for j := 7  ; j < 9; j++ {
		fmt.Println(j)		
	}

	for {
		fmt.Println("for")
		break
	}

}

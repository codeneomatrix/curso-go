package main

import "math"
import "fmt"


func plus(a int, b int) int {
    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func vals() (int, int) {
    return 3, 7
}


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

//if

	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

//arrays

    var cadena [5]int
    fmt.Println("emp:", cadena)

    cadena[4] = 100
    fmt.Println("set:", cadena)
    fmt.Println("get:", cadena[4])

    fmt.Println("len:", len(cadena))

    cadena2 := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", cadena2)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

//mapas

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    m2 := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", m2)

//funciones
    res := plus(1, 2)
    fmt.Println("1+2 =", res)
    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)

    va1, va2 := vals()
    fmt.Println(va1)
    fmt.Println(va2)
}

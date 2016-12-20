package main


import "fmt"

func main() {
	slice := make([]int, 2, 5)

	//slice = append(slice, 1, 77, 13, 17)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

//	for i := 0; i < 20; i++ {
//		slice = append(slice, i)
//
//		fmt.Println("Len: ", len(slice), "Cap: ", cap(slice))
//	}

	//sliceTest := slice
	//slice[0] = 10
	//
	//slice = append(slice, 1, 77)
	//
	//fmt.Println(slice)
	//fmt.Println(sliceTest)

	sliceString := []string{
		"Candido",
		"José",
		"de",
		"Souza",
		"Neto",
	}

	fmt.Println(sliceString)
	fmt.Println(sliceString[0]) //Candido
	fmt.Println(sliceString[:2]) // Candido José
	fmt.Println(sliceString[2:5]) // de Souza Neto

}
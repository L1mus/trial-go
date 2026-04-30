package internal

import (
	"fmt"
	"math"
)

func GenerateNumber (n int ,number chan int){
	for i:= range n {
		number <- i+1
	}
}

func EvenNumber(number chan int) {
	for data := range number{
		if data % 2 == 0 {
			number <- data
		}else{
			continue
		}
	}
}

func SquareOfNumber(number chan int)  {
	for data := range number{
		fmt.Println(math.Pow(float64(data),2))
	}
}
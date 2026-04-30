package internal

import (
	"fmt"
	"math"
)

func GenerateNumber (n int ,number chan int){
	for i:= range n {
		number <- i+1
	}
	close(number)
}

func EvenNumber(number chan int,evenNumber chan int) {
	for data := range number{
		if data % 2 == 0 {
			evenNumber <- data
		}else{
			continue
		}
	}
	close(evenNumber)
}

func SquareOfNumber(evenNumber chan int)  {
	for data := range evenNumber{
		fmt.Println(math.Pow(float64(data),2))
	}
}
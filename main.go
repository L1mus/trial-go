package main

import (
	"fmt"
	"sync"

	"github.com/L1mus/trial-go/internal"
)

// func mandi(wg *sync.WaitGroup)  {
// 	defer wg.Done()
// 	fmt.Println("Memulai mandi")
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Selesai mandi")
// }

// func membuatKopi(wg *sync.WaitGroup)  {
// 	defer wg.Done()
// 	fmt.Println("Memulai membuat Kopi")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Selesai membuat Kopi")
// }

// func menyiapkanSarapan(wg *sync.WaitGroup)  {
// 	 defer wg.Done()
// 	fmt.Println("Memulai menyiapkan sarapan")
// 	time.Sleep(10 * time.Second)
// 	fmt.Println("Selesai menyiapkan sarapan")
// }

// func merapikanKamarTidur(wg *sync.WaitGroup)  {
// 	defer wg.Done()
// 	fmt.Println("Memulai merapikan kamar tidur")
// 	time.Sleep(6 * time.Second)
// 	fmt.Println("Selesai merapikan kamar tidur")
// }

func main()  {
	number := make(chan int)
	var wg sync.WaitGroup

	// wg.Add(4)
	// go membuatKopi(&wg)

	// go menyiapkanSarapan(&wg)

	// go mandi(&wg)

	// go merapikanKamarTidur(&wg)
	n := 3
	
	// wg.Wait()
	// println("Berangkat Kerja")
	// internal.Goroutine()
	wg.Go(func() {
			internal.GenerateNumber(n,number)
	})

	wg.Go(func() {
		internal.EvenNumber(number)
	})

	for i :=range number {
		fmt.Println(i)
	}


	// wg.Go(func() {
	// 	internal.SquareOfNumber(number)
	// })

	close(number)

	wg.Wait()



}
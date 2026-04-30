package internal

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)


func sendMessage( messageChan chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan nama anda:")
	scanner.Scan()
	nama := scanner.Text()
	fmt.Print("Masukan pesan anda:")
	scanner.Scan()
	pesan := scanner.Text()
	fmt.Print("Mengirim pesan ....\n")
	time.Sleep(3 * time.Second)
	data := fmt.Sprintf("Nama Pengirim : %s\nPesan : %s \n",nama,pesan)
	fmt.Printf("Pesan terkirim !!! Pesan dikirim ke %s\n",nama)
	messageChan <- data
}

func whiteBoard (messageChan chan string) {
	data := <- messageChan
	fmt.Print(data)
}

func Goroutine()  {
	var wg sync.WaitGroup
	messageChan := make(chan string )
	
	wg.Go(func() {
		whiteBoard(messageChan)
	})

	wg.Go(func ()  {
		sendMessage(messageChan)
	})
	close(messageChan)
	wg.Wait()
}
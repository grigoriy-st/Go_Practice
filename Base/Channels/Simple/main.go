package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func sendData(ch chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Input: ")
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			fmt.Println("Program is exit.")
			close(ch)
			break
		}

		ch <- text

	}
}

func receiveData(ch <-chan string) {
	for line := range ch {
		time.Sleep(3 * time.Second)
		fmt.Println(line)
	}
}

func main() {
	ch := make(chan string, 1) // Создаем канал для строк

	go sendData(ch) // Запускаем горутину для отправки данных
	receiveData(ch) // Получаем данные из канала
}

package main

import (
	"context"
	"fmt"
	"time"
)

////**********************************************************************************************************************////
//// context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)

func readChan(ctx context.Context) {
	for {
		select {
		default:
			time.Sleep(200 * time.Millisecond)
			fmt.Println("I working...")
		case <-ctx.Done(): // Кейс сигнала отмены контекста
			fmt.Println("<<<<I done!>>>>")
			return
		}
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go readChan(ctx)
	time.Sleep(3 * time.Second)
	cancelFunc() // Вызываем функцию отмены контекста
	time.Sleep(1 * time.Second)
	fmt.Println("Program completion")
}

////**********************************************************************************************************************////
//// ch <-chan bool

//func readChan(ch <-chan bool) {
//	for {
//		select {
//		default:
//			time.Sleep(200 * time.Millisecond)
//			fmt.Println("I working...")
//		case <-ch: // Ожидаем сигнал от канала, для выхода из функции
//			fmt.Println("<<<<I done!>>>>")
//			return
//		}
//	}
//}
//
//func main() {
//	stopChan := make(chan bool)
//	go readChan(stopChan)
//	time.Sleep(3 * time.Second)
//	stopChan <- true // Посылаем сигнал в канал
//	close(stopChan)
//	fmt.Println("Program completion")
//}

//**********************************************************************************************************************////
// time.After(int * time.Second)

//func readChan() {
//	to := time.After(3 * time.Second)
//	for {
//		select {
//		case <-to:
//			fmt.Println("<<<<I done!>>>>")
//			return
//		default:
//			time.Sleep(200 * time.Millisecond)
//			fmt.Println("I working...")
//		}
//	}
//}
//
//func main() {
//	go readChan()
//	time.Sleep(4 * time.Second)
//	fmt.Println("Program completion")
//}

////**********************************************************************************************************************////
//// sync.WaitGroup для ожидания завершения главного потока

//func readChan(id int, wg *sync.WaitGroup) {
//	fmt.Println(fmt.Sprintf("id: %d| I working...", id))
//	time.Sleep(2 * time.Second)
//	fmt.Println(fmt.Sprintf("\"id: %d| <<<<I done!>>>>", id))
//	wg.Done()
//}
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(3)
//	for i := 0; i < 3; i++ {
//		go readChan(i, &wg)
//	}
//	wg.Wait() // Главный поток не завершится, пока не будет готова вся группа
//	fmt.Println("Program completion")
//}

////**********************************************************************************************************************////
//// Используем os/signal

//func readChan(sigCh <-chan os.Signal) {
//	for {
//		select {
//		case <-sigCh: // Кейс на сигнал (control + C)
//			fmt.Println("<<<<I done!>>>>")
//			return
//		default:
//			time.Sleep(200 * time.Millisecond)
//			fmt.Println("I working...")
//		}
//	}
//}
//
//func main() {
//	sigCh := make(chan os.Signal, 1)
//	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
//	go readChan(sigCh)
//	<-sigCh // Ожидаем сигнал (control + C)
//	fmt.Println("Program completion")
//}

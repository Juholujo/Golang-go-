package main

import (
	"fmt"
	"unsafe"
)

type hchan struct {
	qcount   uint           // Количество данных в очереди
	dataqsiz uint           // Размер кольцевого буфера
	buf      unsafe.Pointer // Указатель на массив данных
	elemsize uint16         // Размер элемента
	closed   uint32         // Флаг закрытия канала
	elemtype uintptr        // Тип данных
	sendx    uint           // Индекс для записи
	recvx    uint           // Индекс для чтения
	recvq    uintptr        // Очередь ожидающих на чтение
	sendq    uintptr        // Очередь ожидающих на запись
	lock     uintptr        // Мьютекс
}

/*type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	timer    *timer // timer feeding this chan
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}*/

func main() {
	ch := make(chan int, 2)
	ch <- 42

	// Заглянуть в hchan напрямую нельзя, но мы можем представить его так:
	fmt.Println("hchan пример: ")
	fmt.Printf("qcount: 1 (1 элемент в буфере)\ndataqsiz: 2 (ёмкость буфера)\nbuf: [42, _] (кольцевой буфер)\n")
}

/*ch := make(chan int, 4)

// Отправляем данные в канал
ch <- 10 // sendx = 0 → buf[0] = 10, sendx = 1
ch <- 20 // sendx = 1 → buf[1] = 20, sendx = 2
ch <- 30 // sendx = 2 → buf[2] = 30, sendx = 3
ch <- 40 // sendx = 3 → buf[3] = 40, sendx = 0 (оборачивается)

// Читаем данные из канала
fmt.Println(<-ch) // recvx = 0 → buf[0] = 10, recvx = 1
fmt.Println(<-ch) // recvx = 1 → buf[1] = 20, recvx = 2
ch <- 50          // sendx = 0 → buf[0] = 50, sendx = 1 (оборачивается)
fmt.Println(<-ch) // recvx = 2 → buf[2] = 30, recvx = 3
fmt.Println(<-ch) // recvx = 3 → buf[3] = 40, recvx = 0 (оборачивается)
fmt.Println(<-ch) // recvx = 0 → buf[0] = 50, recvx = 1*/

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

func main() {
	ch := make(chan int, 2)
	ch <- 42

	// Заглянуть в hchan напрямую нельзя, но мы можем представить его так:
	fmt.Println("hchan пример: ")
	fmt.Printf("qcount: 1 (1 элемент в буфере)\ndataqsiz: 2 (ёмкость буфера)\nbuf: [42, _] (кольцевой буфер)\n")
}

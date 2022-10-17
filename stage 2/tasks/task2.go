package tasks

import (
	"fmt"
	"runtime"
)

/*
2. Написать программу, которая конкурентно рассчитает значение
квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты
в stdout

Реализация: пул воркеров с учетом Rate-лимитов (гибкая настройка ограничений ресурсов)
Синхронизация через каналы
*/

func Task2() {
	const (
		arrSize    = 5
		goroutines = 3
		quotaLimit = 2
	)

	array := [arrSize]int{2, 4, 6, 8, 10}

	// буферизированный канал - как очередь заданий с квотой на выполнение
	workerInput := make(chan int, quotaLimit)
	// канал для синхронизации. Как вариант альтернативы - использовать sync.Waitgroup
	done := make(chan bool)

	// создаем горутины
	for i := 0; i < goroutines; i++ {
		go worker2(i, workerInput, done)
	}

	// воркеры сами разберутся - задача идет в канал.
	// рандомная горутина принимает к работе
	for _, num := range array {
		workerInput <- num
	}

	// обязательно закрыть канал (пул воркеров) - иначе не дождемся окончания
	// работы воркеров. Это может привести к дедлоку или утечки памяти
	close(workerInput)

	// ожидаем завершения работы горутин
	for i := 0; i < goroutines; i++ {
		<-done
	}
}

func worker2(workerName int, in <-chan int, done chan<- bool) {
	for {
		// num - значение из канала, а more = false, если канал закрыт
		num, more := <-in
		if more {
			// выполнение работы из пула
			fmt.Printf("Goroutine #%d: sqr(%d) = %d \n", workerName, num, num*num)

			/* Готовые к исполнению горутины выполняются в порядке очереди, то есть FIFO
			(First In, First Out). Исполнение горутины прерывается только тогда, когда
			она уже не может выполняться: то есть из-за системного вызова или
			использования синхронизирующих объектов (операции с каналами, мьютексами и
			т.п.). Не существует никаких квантов времени на работу горутины, после
			выполнения которых она бы заново возвращалась в очередь. Чтобы позволить
			планировщику сделать это, нужно самостоятельно вызвать runtime.Gosched().
			На практике это в первую очередь означает, что иногда стоит использовать
			runtime.Gosched(), чтобы несколько долгоживущих горутин не остановили на
			существенное время работу всех других. С другой стороны, такие ситуации
			встречаются на практике довольно редко

			При запуске программы Go без указания переменной среды runtime.GOMAXPROCS(n int),
			Go goroutines запланированы для исполнения в одиночном потоке ОС. Однако,
			чтобы программа казалась многопоточной, планировщик Go должен иногда
			переключать контекст выполнения, поэтому каждая горутина может
			выполнять свою работу.

			В данном случае использование Gosched() позволит (в большинстве случаев)
			выполнять несколько подряд простых задач не одной горутиной, а несколькими
			*/
			runtime.Gosched()
		} else {
			fmt.Printf("Task2. All jobs is done. (Worker #%d) \n", workerName)
			done <- true
			// выход и бесконечного цикла
			return
		}
	}
}

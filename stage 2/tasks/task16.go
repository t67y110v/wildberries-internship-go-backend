package tasks

import (
	"fmt"
)

/*
16.	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func Task16() {
	array := []int{9, 8, 7, 6, 5, 4, 3} //{2, 3, 5, 6, 9, 8, 4} //5, 3, 1, 8, 2, 4, 0, 9}
	fmt.Println("Source array:", array)
	quicksort(array)
	fmt.Println("Sorted array:", array)
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivot := len(a) / 2 // опорный элемент

	// опорный элемент смещаем вправо, чтобы все элементы, мЕньшие чем опорный поместить слева, а бОльше - вправо
	a[pivot], a[right] = a[right], a[pivot]
	// в цикле слева направо сравниваем i-ый элемент с правым (опорным) элементом.
	for i := range a {
		// если элемент >= опорному, по идее мы должны его перенести вправо. Но, поскольку опорный - самый правый, то пока что работаем с мЕньшими элементами, чем опорный
		// находим элемент, мЕньший опорного
		if a[i] < a[right] {
			// если i-ый элемент меньше опорного, кладем его в самую левую позицию. Left - изначально самый левый элемент (с индексом 0).
			// т.е. первый попавшийся элемент, мЕньший опорного кладем в самое лево.
			a[left], a[i] = a[i], a[left]
			// и смещаем индекс left на 1 вправо, чтобы не потерять i-ый.
			left++
		}
	}
	// таким образом: a =  [элементы, мЕньшие опорного] [elem1, elem2, ...элементы, бОльшие опорного] опорный
	//                                                     ^
	//                                                    left - индекс подмножества, бОльшие опорного
	// после цикла left содержит позицию первого элемента (слева направо), который больше, чем опорный
	// поскольку из цикла вышли, то a[left] превысил (либо равен) самому правому (опорному) элементу. Поскольку опорный меньше (либо равен) a[left],
	// меняем их
	a[left], a[right] = a[right], a[left]

	// таким образом: a =  [элементы, мЕньшие опорного] опорный [elem1, elem2, ...элементы, бОльшие опорного]
	//                                                     ^
	//                                                   left - теперь опорный

	// в итоге - [0:left] - элементы, точно меньшие опорного. А справа от позиции left - элементы, бОльшие опорного
	// сортируем отдельно левую часть (не включая опорный элемент)
	quicksort(a[:left])

	// и правую (не включая опорный элемент)
	quicksort(a[left+1:])

	return a
}
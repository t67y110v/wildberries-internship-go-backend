package tasks

import "fmt"

/*
15.	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// использование глобальной переменной - плохая практика:
var justString string

    1 Глобальные переменные в большинстве случаев нарушают инкапсуляцию. К ним открыт неконтролируемый доступ отовсюду.
    2 В большом проекте при обилии глобальных переменных возникает путаница в именах. Глобальную переменную же видно отовсюду, надо, чтобы отовсюду было понятно, зачем она.
    3 Глобальные переменные в большинстве случаев нарушают принцип инверсии зависимостей (или делают возможным его нарушение).
    4 Глобальные переменные ухудшают масштабируемость проекта.
    5 Глобальные переменные ухудшают читаемость кода (в каком-то конкретно взятом месте непонятно, нужна ли какая-то конкретная глобальная переменная, или нет).
    6 Глобальные переменные приводят к трудноуловимым ошибкам. Примеры: нежелательное изменение её значения в другом месте/другим потоком, ошибочное использование
		глобальной переменной для промежуточных вычислений из-за совпадения имен, возвращение функцией неправильного значения при тех же параметрах (оказывается,
		она зависима от глобальной переменной, а ее кто-то поменял).
    7 Глобальные переменные создают большие сложности при использовании модульного тестирования.
    8 Глобальные переменные увеличивают число прямых и косвенных связей в системе, делая её поведение труднопредсказуемым, а её саму - сложной для понимания и развития.

func someFunc() {
	// createHugeString вернет или string или slice длинной 1 << 10( 2^10), рассмотреть их поведение при взятии среза
	v := createHugeString(1 << 10)

	justString = v[:100]	// 2 варианта: строка и срез
}

func main() {
	someFunc()
}
*/

func createHugeString() []byte {
	hugeString := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
	Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
	Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
	Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	// createHugeString возвращает слайс байт (type []byte)
	/*
		В таком случае, нужно обратить внимание:
			1. createHugeString возвращает срез []byte, который указывает на массив. Нарезка среза не делает копию базового массива.
				Поскольку срез ссылается на исходный массив, пока срез хранится, то сборщик мусора не может освободить массив; (!)
				ведь 100 полезных байтов сохраняют все содержимое в памяти.
			2. Нужно убедиться, что в слайсе байт hugeStringByte есть > 100 элементов, поскольку срез берем от 0 до 100 байта.Хотя, если по условию обещают, что строка очень большая, то это не проблема.
			3. Нужно учитывать, что изменение элементов среза изменяет элементы исходного среза.
			4. Если при добавлении элемента в срез длина увеличивается на единицу и тем самым превышает заявленный объем, необходимо
				предоставить новый объем (в этом случае текущий объем обычно удваивается). Что может привести к большим затратам по памяти

		Чтобы решить проблему 1, можно скопировать интересные данные в новый срез перед возвратом:
	*/
	hugeStringByte := []byte(hugeString)
	justStringByte := hugeStringByte[:100]        // оставил срез, для того, чтобы указать диапазон данных, например [23:28]
	newSlice := make([]byte, len(justStringByte)) // создаем пустой срез с заданной длиной
	copy(newSlice, hugeStringByte)                // копируем и возвращаемся из функции
	return newSlice
}

func Task15() {
	// Теперь, когда функция createHugeString() вернула копию полезных данных, то сборщик мусора может освободить массив
	fmt.Printf("Task15: %c\n", createHugeString())
}

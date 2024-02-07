package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		// Инициализируем буферизированный канал. Поскольку мы знаем, что
		// произведём 6 результатов, создадим буферизированный канал с
		// вместимостью 5, чтобы горутина могла завершиться как можно быстрее
		resultStream := make(chan int, 5)

		// Здесь мы запускаем анонимную горутину, которая выполняет запись в
		// resultStream. Следует обратить внимание, что мы изменили способ
		// создания горутин. Теперь горутина инкапсулирована в окружающую
		// функцию
		go func() {
			// Гарантируем, что канал resultStream будет закрыт, как только мы
			// закончим с ним работать. Как владелец канала, мы несём за это
			// ответственность
			defer close(resultStream)

			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()

		// Здесь мы возвращаем канал. Поскольку возвращаемое значение объявлено
		// как канал только для чтения, resultStream будет неявно преобразован в
		// доступный только для чтения для пользователей
		return resultStream
	}

	resultStream := chanOwner()

	// Здесь мы пробегаемся по resultStream. Как пользователя канала, нас
	// беспокоят только блокировка и закрытость каналов
	for result := range resultStream {
		fmt.Printf("Received: %d.\n", result)
	}

	fmt.Println("Done receiving!")
}

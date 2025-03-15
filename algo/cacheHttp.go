// Имеется веб-сервис, который вызывает метод который возвращает значение очень долго, пользователю просто наиболее свежее значение. Реализовать

// Решение: добавил переменную с кэшем. При первой загрузке будет дефолтное значение 0, далее спустя 5 секунд при обновлении страницы будет необходимое значение.
package main

import (
	"fmt"
	"net/http"
	"time"
)

var resultCache int

func LongTask() int {
	time.Sleep(5 * time.Second) // Симуляция задержки
	result := 45                // Обработанный результат
	return result
}

func main() {
	// Определение обработчика
	handler := func(w http.ResponseWriter, r *http.Request) {

		go func() {
			res := LongTask()
			resultCache = res
		}()

		fmt.Fprintf(w, "Result: %d", resultCache)
	}

	// Настройка HTTP-сервера
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

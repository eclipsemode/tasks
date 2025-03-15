// Что выведет код и как поправить?

// В исходной версии код бы имел следующий порядок действий:
// 1) Вывод get interrupted signal и блокировка в select на записи в канал (так как никто не читает небуферизированный канал)
// 2) Вывод get normal signal так как горутина была зараннее запущена и блокировка на записи в канал

// Решение: Либо удалить ticker и убрать его из case. Либо второй вариант, как сделал я, добавил горутину которая будет считывать все прерванные сигналы, и когда первая горутина дойдет до записи, то select считает значение и закончит
// программу. В идеале можно еще добавть доп. канал для inerrupted signal и отправлять туда данные прерванного сигнала, так как хоть в коде и не получается этого добиться, но мне кажется что запись false из первой горутины может попасть в range sync, и тогда
// не отработает fmt.Printf("finish %t", value) эта строка.
package main

import (
	"fmt"
	"time"
)

func main() {
	sync := make(chan bool)

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("get normal signal")
		sync <- false
	}()

	go func() {
		for v := range sync {
			fmt.Println(v)
		}
	}()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("get interrupted signal")
			sync <- true
		case value := <-sync:
			fmt.Printf("finish %t", value)
			return
		}
	}
}

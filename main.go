// Просмотреть код, назвать, что выведется в stdout
package main

import (
	"fmt"
	"time"
)

// https://www.jdoodle.com/execute-go-online
type Agent struct {
	ID      int
	Enabled bool
}

func (a *Agent) Enable() {
	a.Enabled = true
}

type Enabler interface {
	Enable()
}

// 1. Инициализация слайса агентов
// 2. Дополнение слайса сторонними агентами
// 3. Потоковая обработка объектов - активация и отправка на выполнение работы
// 4. Потоковая обработка объектов - сохранение в БД и распечатка резульатов
func main() {
	agents := make([]Agent, 0)
	for i := 0; i < 2; i++ {
		agents = append(agents, Agent{ID: i})
	}

	addThirdPartyAgents(agents)

	pipe := make(chan Enabler, 1)

	go pipeEnableAndSend(pipe, agents)

	for i := 0; i < 2; i++ {
		go func() {
			pipeProcess(pipe)
		}()
	}

	time.Sleep(time.Minute * 1)
}

func addThirdPartyAgents(agents []Agent) {
	thirdParty := []Agent{
		{ID: 4},
		{ID: 5},
	}
	agents = append(agents, thirdParty...)
}

func pipeEnableAndSend(pipe chan Enabler, agents []Agent) {
	for i, a := range agents {
		fmt.Println(i)
		a.Enable()
		pipe <- &a
	}
	close(pipe)
}

func pipeProcess(pipe chan Enabler) {
	for a := range pipe {
		dbWrite(a)
	}
}

var dbWrite = func(a any) {
	fmt.Println(a)
	time.Sleep(time.Second * 1)
}

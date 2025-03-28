package _8_03_2025

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println(Do(context.Background(), []User{{"aaa"}, {"bbb"}, {"ccc"}, {"ddd"}, {"eeee"}}))
}

type User struct {
	Name string
}

func fetchByName(ctx context.Context, userName string) (int, error) {
	// Tyr происходит сетевой поход, который по username возвращает userID

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		// ИМИТАЦИЯ ОШИБКИ - МОЖНО УБРАТЬ
		if userName == "ccc" {
			return 0, errors.New("aaa")
		}
		time.Sleep(10 * time.Millisecond) // ИМИТАЦИЯ СЕТЕВОГО ЗАПРОСА
		return rand.Int() % 100000, nil
	}

}

func Do(ctx context.Context, users []User) (map[string]int, error) {
	collected := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	fetchErr := make(chan error, 1)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, u := range users {
		wg.Add(1)

		go func(usr User) {
			defer wg.Done()

			userID, err := fetchByName(ctx, usr.Name)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}
				fetchErr <- err
				return
			}

			mu.Lock()
			collected[usr.Name] = userID
			mu.Unlock()
		}(u)
	}

	wg.Wait()

	select {
	case err := <-fetchErr:
		close(fetchErr)
		return collected, err
	default:
		return collected, nil
	}
}

//package main
//
//import (
//"context"
//"fmt"
//"math/rand"
//"sync"
//"time"
//)
//
//func main() {
//	fmt.Println(Do(context.Background(), []User{{"aaa"}, {"bbb"}, {"ccc"}, {"ddd"}, {"eeee"}}))
//}
//
//type User struct {
//	Name string
//}
//
//type Collected struct {
//	name string
//	id   int
//	err  error
//}
//
//func fetchByName(ctx context.Context, userName string) (int, error) {
//	// Tyr происходит сетевой поход, который по username возвращает userID
//
//	time.Sleep(10 * time.Millisecond) // ИМИТАЦИЯ СЕТЕВОГО ЗАПРОСА
//	return rand.Int() % 100000, nil
//}
//
//func Do(ctx context.Context, users []User) (map[string]int, error) {
//	collected := make(map[string]int)
//	var wg sync.WaitGroup
//
//	ch := make(chan Collected)
//
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	for _, u := range users {
//		wg.Add(1)
//
//		go func(usr User) {
//			defer wg.Done()
//			userID, err := fetchByName(ctx, usr.Name)
//
//			ch <- Collected{
//				name: usr.Name,
//				id:   userID,
//				err:  err,
//			}
//		}(u)
//	}
//
//	for c := range ch {
//		if c.err != nil {
//			return nil, c.err
//		}
//
//		collected[c.name] = c.id
//	}
//
//	return collected, nil
//}

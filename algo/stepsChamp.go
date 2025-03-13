// Мы в Авито любим проводить соревнования, — недавно мы устроили чемпионат по шагам. И вот настало время подводить итоги!

// Необходимо определить userIds участников, которые прошли наибольшее количество шагов steps за все дни, не пропустив ни одного дня соревнований.

// Пример
// Пример 1
// ввод
// statistics = [
//
//	[{ userId: 1, steps: 1000 }, { userId: 2, steps: 1500 }],
//	[{ userId: 2, steps: 1000 }]
//
// ]
// вывод
// champions = { userIds: [2], steps: 2500 }
package main

import "fmt"

type User struct {
	userId int
	steps  int
}

type Response struct {
	userIds []int
	steps   int
}

func findUser(arr [][]User) Response {

	mUserSteps := make(map[int]int)
	mDays := make(map[int]int)

	for _, day := range arr {
		for _, user := range day {
			mUserSteps[user.userId] += user.steps
			mDays[user.userId]++
		}
	}

	res := Response{}
	maxSteps := 0
	for u, s := range mUserSteps {
		if mDays[u] < len(arr) {
			continue
		}

		if s > maxSteps {
			maxSteps = s

			res = Response{
				userIds: []int{u},
				steps:   s,
			}
		} else if s == maxSteps {
			res.userIds = append(res.userIds, u)
		}
	}

	return res
}

func main() {

	in := [][]User{
		{
			{userId: 1, steps: 1000}, {userId: 2, steps: 1500},
		},
		{
			{userId: 2, steps: 1000},
		},
	}

	res := findUser(in)

	fmt.Println(res)
}

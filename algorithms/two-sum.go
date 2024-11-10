package algorithms

import (
	"time"
)

type HashMap = map[int]int
type Hooks interface {
	Found([]int)
	OnStep(hashMap HashMap, index int)
}

func TwoSum(nums []int, target int, hooks Hooks) {
	go func() { // Run in a goroutine to enable async-like stepping
		hashMap := make(HashMap)

		for i, num := range nums {
			currentNumTarget := target - num
			currentNumTargetIndex, exists := hashMap[currentNumTarget]

			if exists {
				result := []int{currentNumTargetIndex, i}
				hooks.Found(result)
				return
			}

			hashMap[num] = i
			hooks.OnStep(hashMap, i)
			time.Sleep(1 * time.Second) // Sleep for the UI to reflect updates
		}
	}()
}

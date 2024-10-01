package services

import (
	"errors"
	"log"
	"sort"
)

type ArrayService struct {
}

func NewArrayService() *ArrayService {
	//construct repositories if any comes
	return &ArrayService{}
}

// Using two pointer approach to optimize space
func (service *ArrayService) FindPairsWithSum(numbers []int, target int) ([][]int, error) {
	if len(numbers) == 0 {
		return nil, errors.New("the input array is empty")
	}
	sort.Ints(numbers)

	result := [][]int{}
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			result = append(result, []int{left, right})

			// This is to skip duplicates
			for left < right && numbers[left] == numbers[left+1] {
				left++
			}
			for left < right && numbers[right] == numbers[right-1] {
				right--
			}

			left++
			right--
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	if len(result) == 0 {
		log.Println("No pair found that sums up to the target")
	}

	return result, nil
}

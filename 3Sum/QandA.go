package Three_Sum

import (
	"fmt"
	"sort"
)

//https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/776/
//summary of this algorithm, we will reduce 3sum question to 2sum, then maintain deduplicate logic
//take a note if there is no order requirement, we could consider sort or set them first to make location relationship

const target = 0

func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		summary := 0 - nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			fmt.Println(nums[i], nums[left], nums[right])
			fmt.Println(i, left, right)
			if nums[left]+nums[right] == summary {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left = left + 1
				right = right - 1
				for left < right && nums[left] == nums[left-1] {
					left = left + 1
				}
				for left < right && nums[right] == nums[right+1] {
					left = right - 1
				}
			} else if summary > nums[left]+nums[right] {
				left = left + 1
			} else {
				right = right - 1
			}
		}
	}
	return result
}

func Wrong_ThreeSum(nums []int) [][]int {
	result := [][]int{}
	ArrayMap := map[int]map[int]bool{}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == target {
					fmt.Println(nums[i], nums[j], nums[k])
					fmt.Printf("%+v\n", ArrayMap)
					vMap, ok := ArrayMap[nums[i]]
					if ok {
						_, ok := vMap[nums[j]]
						if ok {
							continue
						} else {
							vMap[nums[j]] = true
							ArrayMap[nums[i]] = vMap
						}
					} else {
						vi := map[int]bool{nums[j]: true}
						ArrayMap[nums[i]] = vi
						vi2 := map[int]bool{nums[k]: true}
						ArrayMap[nums[i]] = vi2
						vj := map[int]bool{nums[i]: true}
						ArrayMap[nums[j]] = vj
						vj2 := map[int]bool{nums[k]: true}
						ArrayMap[nums[j]] = vj2
						vk := map[int]bool{nums[i]: true}
						ArrayMap[nums[k]] = vk
						vk2 := map[int]bool{nums[j]: true}
						ArrayMap[nums[k]] = vk2
					}
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}

		}
	}
	return result
}

package Three_Sum

import "fmt"

//todo 2022.8.31-9.1
//https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/776/

const target = 0

func ThreeSum(nums []int) [][]int {
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

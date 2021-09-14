package array

import (
	"sort"
	"strconv"
)

// 179
type mergeStr []string

func largestNumber(nums []int) string {
	str := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		str[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(mergeStr(str))
	if str[0] == "0" {
		return "0"
	}
	res := ""
	for _, s := range str {
		res = res + s
	}
	return res
}

func (m mergeStr) Less(i, j int) bool {
	return m[j]+m[i] < m[i]+m[j]
}

func (m mergeStr) Len() int {
	return len(m)
}

func (m mergeStr) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

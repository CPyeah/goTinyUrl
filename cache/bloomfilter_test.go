package cache

import (
	"fmt"
	"testing"
)

func TestNewRedisBloomFilter(t *testing.T) {
	bloom := NewRedisBloomFilter(Rdb, "test", 10000)

	//bs := []byte("hello")
	//_ = bloom.Put(bs)
	exists, err := bloom.MightContain([]byte("hello"))
	//exists, err = bloom.MightContain([]byte("hello1"))
	fmt.Println(exists, err)
}

func TestGet(t *testing.T) {
	r := maxCoins([]int{3, 1, 5, 8})
	fmt.Println(r)
}

func maxCoins(nums []int) int {
	maxPoint = 0
	poke(nums, 0)
	return maxPoint
}

var maxPoint int

func poke(nums []int, point int) {
	for i, _ := range nums {
		point += getPoint(nums, i)
		if point > maxPoint {
			maxPoint = point
		}
		left := make([]int, 0)
		right := make([]int, 0)
		if i > 0 {
			left = nums[:i-1]
		}
		if i < len(nums)-1 {
			right = nums[i+1:]
		}
		newNums := append(left, right...)
		poke(newNums, point)
	}
}

func getPoint(nums []int, i int) int {
	left := 1
	if i > 0 {
		left = nums[i-1]
	}
	right := 1
	if i < len(nums)-1 {
		right = nums[i+1]
	}
	return left * right * nums[i]
}

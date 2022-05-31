package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

//type general struct {
//	dataType string
//	flags    bool
//	value	float64
//}

type Flags struct {
	Mean   bool
	Median bool
	Mode   bool
	SD     bool
}

type DataMetrics struct {
	Mean   float64
	Median float64
	Mode   int
	SD     float64
	summ   int
}

func (p *DataMetrics) searchMean(nums []int) {
	if len(nums) == 0 {
		return
	}
	p.Mean = float64(p.summ) / float64(len(nums))
	p.Mean = math.Trunc((p.Mean)*100) / 100
}

func (p *DataMetrics) searchMedian(nums []int) {
	if len(nums) == 0 {
		return
	}
	sort.Ints(nums)
	if len(nums)%2 == 1 {
		p.Median = float64(nums[len(nums)/2])
		return
	}
	p.Median = float64(nums[len(nums)/2-1])
}

func ElemExist(mapMode map[int]int, val int) bool {
	for key, _ := range mapMode {
		if mapMode[key] == val {
			fmt.Println(mapMode[key])
			return true
		}
	}
	return false
}

func (p *DataMetrics) searchMode(nums []int) {
	mapMode := map[int]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if ElemExist(mapMode, nums[i]) {
			fmt.Println("YES")
		} else {
			mapMode[nums[i]] = i
		}
	}
	fmt.Println(mapMode)
}

func main() {

	var (
		nums   = make([]int, 0, 10)
		metric DataMetrics
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tmp, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ERROR")
			continue
		}
		metric.summ += tmp
		nums = append(nums, tmp)
	}
	//fmt.Println(nums)
	//metric.searchMean(nums)
	//metric.searchMedian(nums)
	metric.searchMode(nums)
}

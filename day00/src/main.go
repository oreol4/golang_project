package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Flags struct {
	Mean   *bool
	Median *bool
	Mode   *bool
	SD     *bool
}

type DataMetrics struct {
	Mean   float64
	Median float64
	Mode   int
	SD     float64
	summ   int
	ptr    Flags
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
	} else {
		p.Median = float64(nums[len(nums)/2-1])
	}
}

func (p *DataMetrics) searchSD(nums []int) {
	if len(nums) == 0 {
		return
	}
	for i := range nums {
		p.SD += math.Pow(float64(i)-p.Mean, 2)
	}
	p.SD = math.Sqrt(p.SD / float64(len(nums)))
}

func (p *DataMetrics) searchMode(nums []int) {
	if len(nums) == 0 {
		return
	}
	var (
		mapMode = map[int]int{}
		index   int
		tmp     int
	)
	for i := range nums {
		mapMode[nums[i]]++
	}
	tmp = math.MinInt
	for key, val := range mapMode {
		if val > tmp {
			tmp = val
			index = key
		}
	}
	p.Mode = index
}

func (p *DataMetrics) DisplayInfo(f *Flags) {
	if *f.Mean {
		fmt.Println("Mean:", p.Mean)
	}
	if *f.Median {
		fmt.Println("Median:", p.Median)
	}
	if *f.Mode {
		fmt.Println("Mode:", p.Mode)
	}
	if *f.SD {
		fmt.Println("SD:", p.SD)
	}
}
func main() {

	var (
		nums   = make([]int, 0, 10)
		metric DataMetrics
		flags  Flags
	)
	flags.Mean = flag.Bool("mean", false, "This is mean of array")
	flags.Median = flag.Bool("median", false, "This is median of array")
	flags.Mode = flag.Bool("mode", false, "This is mode of array")
	flags.SD = flag.Bool("sd", false, "This is SD of array")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tmp, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ERROR")
			return
		}
		metric.summ += tmp
		nums = append(nums, tmp)
	}
	metric.searchMean(nums)
	metric.searchMedian(nums)
	metric.searchMode(nums)
	metric.searchSD(nums)
	metric.DisplayInfo(&flags)
}

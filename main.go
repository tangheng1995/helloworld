package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world")
	readTxt()
}

func readTxt() {
	f, err := os.Open("1.txt")
	if err != nil {
		fmt.Printf("file read fail ,err:", err)
	}
	defer f.Close()

	m := make(map[string][]int)

	br := bufio.NewReader(f)
	for {
		bl, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		var l string = string(bl)
		prase(l, m)
	}
	sortMap(m)
}

func prase(s string, m map[string][]int) {
	slince := make([]int, 0)
	d := strings.Split(s, " ")
	key := d[0]
	v, _ := strconv.Atoi(d[1])
	if val, ok := m[key]; ok {
		newVal := append(val, v)
		m[key] = newVal
	} else {
		ss := append(slince, v)
		m[key] = ss
	}
}

func sortMap(m map[string][]int) {
	mm := make(map[string]int)
	for k, v := range m {
		n := len(v)
		sum := 0
		for i := 0; i < n; i++ {
			sum += v[i]
		}
		avg := sum / n
		mm[k] = avg
	}
	fmt.Println(mm)
	var sortValus []int
	for _, v := range mm {
		sortValus = append(sortValus, v)
	}
	sort.Ints(sortValus)

	for _, v := range sortValus {
		for kk, vv := range mm {
			if v == vv {
				fmt.Println("Key:", kk, "Avg:", vv)
			}
		}
	}
}

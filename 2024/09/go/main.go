package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"strconv"
)

func main() {
	fs := getFS()
	fs.Defrag()
	fmt.Println(fs.Checksum())

	fs = getFS()
	fs.Defrag2()
	fmt.Println(fs.Checksum())
}

func getFS() FS {
	line := common.GetLines()[0]

	var fs FS
	var index int
	var isFile = true

	for _, item := range line {
		length := common.Must(strconv.Atoi, string(item))

		id := -1
		if isFile {
			id = index
			index++
		}

		fs.Add(id, length)

		isFile = !isFile
	}

	return fs
}

type FS struct {
	data []int
}

func (f *FS) Add(id, length int) {
	f.data = append(f.data, common.Repeat(id, length)...)
}

func (f *FS) Defrag() {
	for i, item := range f.data {
		// gap spotted
		if item == -1 {
			// find last number
			lastIndex := f.LastBlock(i)

			if lastIndex == -1 {
				break
			}
			f.data[i], f.data[lastIndex] = f.data[lastIndex], f.data[i]
		}
	}
}

func (f *FS) Defrag2() {
	end := len(f.data) - 1
	var length int
	current := -1
	lastMoved := -1
	for i := end; i >= 0; i-- {
		digit := f.data[i]

		length++

		if digit != current && current != -1 {
			if lastMoved == -1 || current < lastMoved {
				f.swapNumber(i+1, length)
				lastMoved = current
			}
		}

		if digit != current {
			length = 0
		}

		current = digit
	}
}

func (f *FS) swapNumber(start, length int) {
	gap := f.FindGap(length)
	if gap > start {
		return
	}
	// swap digit because why not
	for g := 0; g < length; g++ {
		f.data[gap+g], f.data[start+g] = f.data[start+g], f.data[gap+g]
	}
}

func (f *FS) FindGap(length int) int {
	var start int
	var current int
	for i, item := range f.data {
		if item == -1 {
			current++
			if current == length {
				return start
			}
			continue
		}

		// pretend to start at the next item
		current = 0
		start = i + 1
	}

	return -1
}

func (f *FS) Checksum() int {
	var result int
	for i, item := range f.data {
		if item == -1 {
			continue
		}
		result += i * item
	}

	return result
}

func (f *FS) LastBlock(after int) int {
	// find last number
	for j := len(f.data) - 1; j > after; j-- {
		if f.data[j] != -1 {
			return j
		}
	}

	return -1
}

func (f *FS) String() string {
	var result string
	for _, b := range f.data {
		if b == -1 {
			result += "."
			continue
		}
		result += strconv.Itoa(b)
	}

	return result
}

//The MIT License (MIT)
//Copyright (c) 2015 Evan Lin
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package kmp

import (
	"fmt"
)

//KMP Return index list if str1 contain str2
func kmp(str1 string, str2 string) []int {
	next := preKMP([]byte(str1))
	i := 0
	j := 0
	m := len(str1)
	n := len(str2)

	x := []byte(str1)
	y := []byte(str2)
	var ret []int

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		fmt.Println(i, j)
		if i >= m {
			ret = append(ret, j-i)
			fmt.Println("find:", j, i)
			i = next[i]
		}
	}

	return ret
}

func preKMP(x []byte) [30]int {
	var i, j int
	length := len(x) - 1
	var kmpNext [30]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}

		i++
		j++

		// fmt.Println(i, j)
		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
		// fmt.Println("2", i, j, length)
	}
	return kmpNext
}

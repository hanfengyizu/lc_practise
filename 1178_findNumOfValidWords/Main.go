package main

import (
	"fmt"
	"math/bits"
)

func main() {
	test1()
	test2()
}

func test2() {
	words := []string{"apple", "pleas", "please"}
	puzzles := []string{"aelwxyz", "aelpxyz", "aelpsxy", "saelpxy", "xaelpsy"}
	fmt.Println(findNumOfValidWords(words, puzzles))
}

func test1() {
	words := []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}
	puzzles := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}
	fmt.Println(findNumOfValidWords(words, puzzles))
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	const puzzleLength = 7
	cnt := map[int]int{}
	for _, s := range words {
		mask := 0
		for _, ch := range s {
			mask |= 1 << (ch - 'a')
		}
		if bits.OnesCount(uint(mask)) <= puzzleLength {
			cnt[mask]++
		}
	}

	ans := make([]int, len(puzzles))
	for i, s := range puzzles {
		first := 1 << (s[0] - 'a')

		// 枚举子集方法一
		//for choose := 0; choose < 1<<(puzzleLength-1); choose++ {
		//    mask := 0
		//    for j := 0; j < puzzleLength-1; j++ {
		//        if choose>>j&1 == 1 {
		//            mask |= 1 << (s[j+1] - 'a')
		//        }
		//    }
		//    ans[i] += cnt[mask|first]
		//}

		// 枚举子集方法二
		mask := 0
		for _, ch := range s[1:] {
			mask |= 1 << (ch - 'a')
		}
		subset := mask
		for {
			ans[i] += cnt[subset|first]
			subset = (subset - 1) & mask
			if subset == mask {
				break
			}
		}
	}
	return ans
}
func findNumOfValidWordsV2(words []string, puzzles []string) []int {
	ans := make([]int, len(puzzles))
	puzzles_Arr := make([]rune, 26)
	notMatch := true
	firstMatch := false
	for index, puzzle := range puzzles {
		puzzles_Arr = make([]rune, 26)
		setByteInArr(puzzles_Arr, puzzle)
		tmpFirst := puzzle[0]
		for _, ws := range words {
			notMatch = true
			firstMatch = false
			for _, q := range ws {
				if !notMatch {
					break
				}
				if !firstMatch && q == rune(tmpFirst) {
					firstMatch = true
				}
				if puzzles_Arr[q-97] != q {
					notMatch = false
					break
				}
			}

			if notMatch && firstMatch {
				ans[index] += 1
			}
		}
	}
	return ans
}

func setByteInArr(arr []rune, puzzle string) {
	for _, e := range puzzle {
		arr[e-97] = e
	}
}

func findNumOfValidWordsV1(words []string, puzzles []string) []int {
	ans := make([]int, len(puzzles))
	puzzles_Arr := make([]rune, 26)
	words_Arr := make([]rune, 26)
	for index, puzzle := range puzzles {
		puzzles_Arr = make([]rune, 26)
		setByteInArr(puzzles_Arr, puzzle)
		// fmt.Println(puzzles_Arr)
		for _, w := range words {
			words_Arr = make([]rune, 26)
			setByteInArr(words_Arr, w)
			notMatch := true

			if words_Arr[puzzle[0]-97] == 0 {
				// fmt.Println(words_Arr)
				// fmt.Printf("words_Arr[%d]=%d\n", puzzle[0]-97, words_Arr[puzzle[0]-97])
				notMatch = false
				continue
			}

			for i, q := range words_Arr {
				if !notMatch {
					break
				}
				if puzzles_Arr[i] != q && q != 0 {
					notMatch = false
					break
				}
			}

			if notMatch {
				ans[index] += 1
			}
		}
	}
	return ans
}

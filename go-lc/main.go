package main

import "fmt"

// func containsDuplicate(nums []int) bool {
// 	//why struct
// }

//Input: strs = ["eat","tea","tan","ate","nat","bat"]
//Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//all are lowercase letters

// func groupAnagrams(strs []string) [][]string {
// 	m := make(map[[26]int][]string)

// 	for _, word := range strs {
// 		var k [26]int //array to store count of each char of a word

// 		for _, char := range word {
// 			k[char-'a']++ //find pos of char from a and increment there. for e in eat thats k[4]++ in 26 int arr
// 		}
// 		//i have k of 1 word with occurences of each letter
// 		//append each k to map with word associated
// 		m[k] = append(m[k], word)
// 	}
// 	fmt.Println(m)
// 	//i have map of all words in strs, to convert into slice
// 	result := [][]string{}
// 	for _, words := range m {
// 		result = append(result, words)
// 	}
// 	fmt.Println(result)
// 	return result
// }

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]
// func topKFrequent(nums []int, k int) []int {
// 	//make a map of occurences
// 	//return array of values >= given k value

// 	m := make(map[int]int)

// 	for _, char := range nums {
// 		m[char]++
// 	}
// 	fmt.Println(m) //counted occurences

// 	n := len(nums)
// 	buckets := make([][]int, n+1) //buckets of length n+1

// 	for num, freq := range m {
// 		buckets[freq] = append(buckets[freq], num) //
// 	}

// 	var result []int
// 	for i := n; i >= 0 && len(result) < k; i-- {
// 		result = append(result, buckets[i]...)
// 	}

// 	fmt.Println(result)

// 	return result[:k]
// }

// func longestConsecutive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	//make a hashset
// 	numSet := make(map[int]bool)
// 	for _, num := range nums {
// 		numSet[num] = true
// 	}

// 	//find the starting index
// 	longestLen := 0
// 	for num := range numSet {
// 		if !numSet[num-1] {
// 			currNum := num //can only be startingIdx the first time
// 			currLen := 1

// 			//check if there are consec. nums from current number
// 			for numSet[currNum+1] {
// 				currNum++
// 				currLen++
// 			}

// 			if currLen > longestLen {
// 				longestLen = currLen
// 			}
// 		}
// 	}
// 	return longestLen
// }

// nums = [a, b, c, d]
// prefix = [1, a, ab, abc]
// suffix = [bcd, cd, d, 1]
// answer = [bcd, acd, abd, abc]
// func productExceptSelf(nums []int) []int {
// 	//using a sum of prefix and suffix product

// 	n := len(nums)
// 	answer := make([]int, n) //array of len n

// 	//calculate prefix products and store in answer
// 	answer[0] = 1 //1st element

// 	for i := 1; i < n; i++ {
// 		answer[i] = answer[i-1] * nums[i-1] //ans[1] = ans[0] * nums[0] = 1 * 1, ans[2] = ans[1] * nums[1] = 2..
// 	}

// 	//calculate suffix and multiply w/ answer
// 	suffix := 1
// 	for i := n - 1; i >= 0; i-- {
// 		answer[i] = answer[i] * suffix //ans[1] = 1 * 1, ans[2] = 2 * 2 = 4
// 		suffix *= nums[i]              //suffix = 1*2 = 2, 2 * 3 = 6
// 	}

// 	return answer
// }

// func isValidSudoku(board [][]byte) bool {
// 	row := make([][9]bool, 9)
// 	col := make([][9]bool, 9)
// 	grid := make([][9]bool, 9) //for 3x3 grid

// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			currVal := board[i][j] // current val of cell
// 			if currVal == '.' {
// 				continue //skip non numbers
// 			}

// 			index := currVal - '1' // ???

// 			//???
// 			if row[i][index] || col[j][index] || grid[(i/3)*3+(j/3)][index] {
// 				return false
// 			}

// 			row[i][index] = true
// 			col[j][index] = true
// 			grid[(i/3)*3+(j/3)][index] = true

// 		}
// 	}
// 	return true
// }

// func isPalindrome(s string) bool {
// 	left := 0
// 	right := len(s) - 1

// 	for left < right {
// 		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
// 			left++
// 			// fmt.Println("left", left)

// 		}
// 		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
// 			right-- // fmt.Println("right", right)

// 		}
// 		fmt.Println("l", unicode.ToLower(rune(s[left])), "r", unicode.ToLower(rune(s[right])))
// 		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
// 			return false
// 		}

//			left++
//			right--
//		}
//		return true
//	}
func threeSum(nums []int) [][]int {
	m := make(map[int]bool)
	// res := []int{}
	// diff := 0

	for _, val := range nums {
		m[val] = true //map[-4:true -1:true 0:true 1:true 2:true]
	}
	//how to add these to get a 3sum

	fmt.Println(m)
	return nil
}

func main() {

	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// groupAnagrams(strs)

	// nums := []int{1, 1, 1, 2, 2, 3}
	// nums := []int{1, 2, 3, 4}

	// productExceptSelf(nums)

	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	// fmt.Println(isPalindrome("race a car"))                     // false
	// fmt.Println(isPalindrome(" "))                              // true
	nums := []int{-1, 0, 1, 2, -1, -4}
	threeSum(nums)
}

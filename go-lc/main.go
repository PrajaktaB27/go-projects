package main

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

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//make a hashset
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	//find the starting index
	longestLen := 0
	for num := range numSet {
		if !numSet[num-1] {
			currNum := num //can only be startingIdx the first time
			currLen := 1

			//check if there are consec. nums from current number
			for numSet[currNum+1] {
				currNum++
				currLen++
			}

			if currLen > longestLen {
				longestLen = currLen
			}
		}
	}
	return longestLen
}

func main() {

	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// groupAnagrams(strs)

	nums := []int{1, 1, 1, 2, 2, 3}
	// topKFrequent(nums, 2)

	longestConsecutive(nums)
}

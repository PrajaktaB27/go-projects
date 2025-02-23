package main

// "slices"
// "time"

// const s string = "hi"

// func add(a int, b int) int {
// 	return a + b
// }

// func addAgain(a, b, c int) int {
// 	return a + b + c
// }

// func checkValue(m map[string]int, key string) (int, error) {
// 	if v1, ok := m[key]; ok {
// 		return v1, nil
// 	} else {
// 		return 0, fmt.Errorf("oop")
// 	}
// }

// func sum(nums ...int) {
// 	fmt.Println(nums, "")
// 	total := 0

// 	for _, nums := range nums {
// 		total += nums
// 	}
// 	fmt.Println(total)
// }

// func intSeq() func() int {
// 	i := 0

// 	return func() int {
// 		i++
// 		return i
// 	}
// }

// func fact(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * fact(n-1)

// 	//n = 3
// 	// n * fact(2) = 3 * 2

// 	//2 * fact(1) = 2 * 1

// 	//1 * fact(0) = 1 * 1

// 	//1
// }

func main() {
	// fmt.Println("Hello, World!")
	// var e bool
	// fmt.Println(e)

	// f := "apple"
	// fmt.Println(f)
	// fmt.Println(s)

	// const n = 5000000
	// const d = 3e20 / n
	// fmt.Println(d)

	// fmt.Println(int64(d))

	// for i := range 2 {
	// 	fmt.Println("range", i+1)
	// }

	// if num := 19; num < 0 {
	// 	fmt.Println(num, "neg")
	// } else if num < 10 {
	// 	fmt.Println(num, "single digit")
	// } else {
	// 	fmt.Println(num, "yayy")
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("yayy weekend")
	// default:
	// 	fmt.Println("nooo weekday")
	// }

	// // the interface is a big no no
	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("bool")
	// 	case int:
	// 		fmt.Println("int")
	// 	default:
	// 		fmt.Println("idk", t)
	// 	}
	// }

	// whatAmI(true)
	// whatAmI(4)
	// whatAmI("who i am")

	// a := [5]int{1, 2, 3, 4, 5}
	// // var b [5]int
	// fmt.Println(a)
	// // fmt.Println(b)
	// fmt.Println(a[4])
	// fmt.Println(len(a))

	// b := [...]int{1, 2: 400, 5, 7}
	// fmt.Println(b)

	// this can be a list of arrays
	// var matrix [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		matrix[i][j] = i + j
	// 	}
	// }

	// fmt.Println("matrix", matrix)

	// SLICES
	// var s []string
	// fmt.Println(s, s == nil, len(s) == 0)

	// s = make([]string, 3)
	// t1 := []string{"h", "a", "h", "a"}
	// fmt.Println(t1)

	// t2 := []string{"h", "b", "h", "a"}
	// if slices.Equal(t1, t2) {
	// 	fmt.Println("equal")
	// } else {
	// 	fmt.Println("false")

	// }

	// s ["" "" ""]
	// fmt.Println(s, len(s), cap(s))

	// s = append(s, "a")
	// s = append(s, "b")
	// s = append(s, "c")

	// fmt.Println(s)

	// c := make([]string, len(s))
	// copy(c, s)

	// fmt.Println(c, cap(s))

	// twoD := make([][]int, 3) //will only initialize 1st slice
	// fmt.Println(twoD)

	// m := make(map[string]int)
	// m["k1"] = 7
	// m["k2"] = 13
	// m["k3"] = 0

	// res, err := checkValue(m, "k4")
	// if err == nil {
	// 	fmt.Println(res)
	// } else {
	// 	fmt.Println(err)
	// }

	// fmt.Println(m)
	// if v1, ok := m["k1"]; ok {
	// 	fmt.Println(v1)
	// }

	// for _, value := range m {
	// 	fmt.Println(value)
	// }

	// v2, ok := m["k2"]
	// if ok {
	// 	fmt.Println(v2)
	// }

	// v3, ok := m["k3"]
	// if ok {
	// 	fmt.Println(v3)
	// } else {
	// 	fmt.Println("oop")
	// }

	// delete(m, "k2")
	// clear(m)
	// fmt.Println(m)

	// s1 := "hello"
	// s2 := "hello"

	// if s1 == s2 {
	// 	fmt.Println(true)
	// }

	// nextInt := intSeq()

	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// fmt.Println(fact(3))

	// var fib func(n int) int

	// fib := func(n int) int {
	// 	f0 := 0
	// 	f1 := 1
	// 	n1 := 0 //fib(n-1)
	// 	n2 := 0 //fib(n-2)
	// 	for i := 0; i < n; i++ {
	// 		fmt.Println("i", i)
	// 		if i == 0 {
	// 			n1 = f0
	// 			n2 = 0
	// 		} else if i == 1 {
	// 			n1 = f0
	// 			n2 = f1
	// 		}
	// 		sum := n1 + n2
	// 		n2 = n1
	// 		n1 = sum
	// 	}
	// 	fmt.Println("n1", n1, "n2", n2)

	// 	return n1 + n2
	// 	// fib(4) + fib(3)
	// 	//fib(4) -> fib(3) + fib(2) //
	// 	//fib(3) -> fib(2) + fib(1) // 1 + 1 = 2
	// 	//fib(2) -> fib(1) + fib(0) // 1 + 0 = 1
	// 	// 1
	// }

	// fmt.Println(fib(5))

	//POINTERS

	// x := 7
	// y := &x //pointer of x
	// fmt.Println(x, y)

	// *y = 8            //* deferences i.e. accesses where the address is pointing to
	// fmt.Println(x, y) //doesnt chance the reference address but changes the value of x which is what y points to

	// toChange := "hello"
	// fmt.Println(toChange)
	// // changeValue(&toChange)
	// changeValue2(toChange)
	// fmt.Println(toChange)

	//STRINGS

	// s := "this is a string"
	// fmt.Println(s)
	// fmt.Println(s[1])         //will give 104
	// fmt.Println(string(s[1])) //will give h

	// var r rune = 't'
	// fmt.Println(r)

}

// in param * stands for wanting to access the pointer and not the actual value
// func changeValue(str *string) {
// 	*str = "changes"
// }

// func changeValue2(str string) string {
// 	str = "changes"
// 	return str
// }

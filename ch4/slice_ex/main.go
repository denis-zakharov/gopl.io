package main

import "fmt"

func reverseArray(s *[10]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateRight(s []int, n int) {
	t := len(s)
	if t == n {
		return
	}
	if n > t {
		n -= t
	}
	var movingPart = make([]int, n)
	copy(movingPart, s[t-n:])
	copy(s[n:], s[:t-n])
	copy(s, movingPart)
}

func removeDuplicates(strings []string) []string {
	if len(strings) < 2 {
		return strings
	}
	cur := strings[0]
	k := 1
	for _, s := range strings[1:] {
		if s != cur {
			strings[k] = s
			cur = s
			k++
		}
	}
	return strings[:k]
}

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5}
	reverseArray(&a)
	fmt.Println(a)

	r := []int{1, 2, 3, 4, 5, 6, 7}
	rotateRight(r, 7)
	fmt.Println(r)
	rotateRight(r, 2)
	fmt.Println(r)
	rotateRight(r, 9)
	fmt.Println(r)

	fmt.Println(removeDuplicates([]string{"a", "a", "b", "c", "c", "c", "d"}))
	fmt.Println(removeDuplicates([]string{"a", "b", "b", "c", "d", "e", "e"}))
	fmt.Println(removeDuplicates([]string{"a", "a"}))
}

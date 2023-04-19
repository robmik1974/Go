package main

func linearSearch(s []int, target int) int {
	/*
		Linear Search algoritms
		Time Complexity:
			- Best case: O(1)
			- Worst case: O(n)
		Space Complexity: O(1)
	*/
	for index, value := range s {
		if value == target {
			return index
		}
	}
	return -1
}

func binarySearch(s []int, target int) int {
	/*
		Linear Search algoritms
		Prerequsite : Array(slice) must be sorted upfront
		Time Complexity:
				- Best case: O(1)
				- Worst case: log(n)
			Space Complexity: O(1)
	*/
	leftPointer := 0
	rightPointer := len(s) - 1
	result := binarySearchHelper(s, target, leftPointer, rightPointer)
	return result
}

func binarySearchHelper(s []int, target int, leftPointer int, rightPointer int) int {
	if leftPointer > rightPointer {
		return -1
	}
	middlePoint := (leftPointer + rightPointer) / 2
	middleElement := s[middlePoint]
	if target == middleElement {
		return middlePoint
	} else if target < middleElement {
		rightPointer = middlePoint - 1
		return binarySearchHelper(s, target, leftPointer, rightPointer)
	} else {
		leftPointer = middlePoint + 1
		return binarySearchHelper(s, target, leftPointer, rightPointer)
	}
}

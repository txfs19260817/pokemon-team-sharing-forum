package util

type m = map[string]int

// Given two maps, recursively merge right into left
// If any right key that already exists in left, add its value to the left key
func UnionKeys(left, right m) {
	for key, rightVal := range right {
		if leftVal, present := left[key]; present {
			//then we don't want to replace it - recurse
			left[key] = leftVal + rightVal
		} else {
			// key not in left so we can just shove it in
			left[key] = rightVal
		}
	}
}

// Given a String Arr
// ["I","am","stupid","and","weak"]

// Make it to by for loop
// ["I","am","smart","and","strong"]
package main

import (
	"fmt"
)

func main() {
	stupidArr := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Println("Before: ", stupidArr)

	for i, c := range stupidArr {
		if c == "stupid" {
			stupidArr[i] = "smart"
		} else if c == "weak" {
			stupidArr[i] = "strong"
		} else {
			continue
		}
	}

	fmt.Println("After: ", stupidArr)
}

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// ! Go doesn't support While Loop
/*
// Mimicking while loop
for condition {
    // Loop body
}


package main

import "fmt"

func main() {
    i := 0
    // While loop equivalent
    for i < 5 {
        fmt.Println(i)
        i++
    }
}

*/

/*
for i := 0; i < 10; i++ {
    // Loop body
}

arr := []int{1, 2, 3, 4, 5}
for index, value := range arr {
    // index is the index of the current element
    // value is the value of the current element
}

m := map[string]int{"a": 1, "b": 2, "c": 3}
for key, value := range m {
    // key is the key of the current element
    // value is the value associated with the key
}

for {
    // Infinite loop body
}


i := 0
for i < 10 {
    // Loop body
    i++
}


for i, j := 0, 0; i < 10 && j < 5; i, j = i+1, j+1 {
    // Loop body
}


for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue
    }
    // Loop body
}


for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    // Loop body
}

*/

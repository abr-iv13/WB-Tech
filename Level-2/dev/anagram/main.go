package main

import (
	"fmt"
	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/dev/anagram/pkg/anagram"
)

func main() {
	arr := []string{"пятак", "пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	r := anagram.FindSetsAnagram(arr)
	fmt.Println(r)

}

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	/*mnemonics := []string{"edge", "now", "blur", "luggage", "motor", "wire", "pistol", "camera", "pole", "piece", "genuine", "decorate"}
	mnemonics = RemoveDuplicates(mnemonics)
	s, _ := json.Marshal(mnemonics)
	result := make([]string, 0, 30)
	_ = json.Unmarshal(s, &result)
	fmt.Println(result)*/
	a := []string{"edge", "now", "blur", "luggage", "motor", "wire", "pistol", "camera", "pole", "piece", "genuine", "decorate"}
	s, _ := json.Marshal([]string{"edge", "now", "blur", "luggage", "motor", "wire", "pistol", "camera", "pole", "piece", "genuine", "decorate"})
	b := string(s)
	sliceResult := make([]string, 0, 12)
	_ = json.Unmarshal([]byte(b), &sliceResult)
	fmt.Println(a)
	fmt.Println(sliceResult)
	fmt.Println("equal? ", reflect.DeepEqual(a, sliceResult))
}

func RemoveDuplicates(strs []string) []string {
	// Use a map to track seen strings.
	seen := make(map[string]struct{}, 48)
	var result []string

	for _, str := range strs {
		// If the string is not in the map, add it to the result and mark it as seen.
		if _, exists := seen[str]; !exists {
			seen[str] = struct{}{}
			result = append(result, str)
		}
	}
	fmt.Printf("result 数量为：%d, result=%v \n", len(result), result)
	return result
}

package helper

import (
	"encoding/json"
	"fmt"
)

func D(v any) {
	// JSON整形
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling document: %v\n", err)
		return
	}

	// 結果をコンソールに表示
	fmt.Printf("Document:\n%s\n", jsonData)
}

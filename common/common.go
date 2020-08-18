package common

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONPretty Красивый вывод
func JSONPretty(br []byte) {
	var responseJSON bytes.Buffer
	json.Indent(&responseJSON, br, "", "\t")
	fmt.Printf("Response: %v \n", string(responseJSON.Bytes()))
}

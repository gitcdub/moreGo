//for json marshal and unmarshal codinig/decoding
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func encodingEx(i interface{}) []byte {
	data, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("failed to encode: %s", data)
	}
	return data
}

func decodingEx(data string) interface{} {
	var i interface{}
	err := json.Unmarshal([]byte(data), &i)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	return i
}
func main() {
	fmt.Printf("\n######## ENCODING USING json.Marshal ############\n")
	fmt.Printf("\nEncoded %d to %s", 100, encodingEx(100))
	fmt.Printf("\nEncoded %f to %s", 1.52, encodingEx(1.52))
	fmt.Printf("\nEncoded %s to %s", "Australia Canada", encodingEx("Australia Canada"))

	strArray := [5]string{"A", "B", "C", "D", "E"}
	encoded := encodingEx(strArray)
	fmt.Printf("\nEncoded %v to %s", strArray, encoded)

	fmt.Printf("\n######## DECODING USING json.Unmarshal ###########\n")
	fmt.Printf("\nDecoded %f from %s", decodingEx("1000"), "1000")
	fmt.Printf("\nDecoded %v from %s", decodingEx(`["Australia"]`), "Australia")
	fmt.Printf("\nDecoded %v from %s", decodingEx(`["A","C","D"]`), `["A", "C", "D"]`)
}

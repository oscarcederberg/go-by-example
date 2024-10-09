package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	boolBytes, _ := json.Marshal(true)
	fmt.Println(string(boolBytes))

	intBytes, _ := json.Marshal(1)
	fmt.Println(string(intBytes))

	floatBytes, _ := json.Marshal(2.34)
	fmt.Println(string(floatBytes))

	stringBytes, _ := json.Marshal("gopher")
	fmt.Println(string(stringBytes))
	
	sliceData := []string{"apple", "peach", "pear"}
	sliceBytes, _ := json.Marshal(sliceData)
	fmt.Println(string(sliceBytes))
	
	mapData := map[string]int{"apple": 5, "lettuce": 7}
	mapBytes, _ := json.Marshal(mapData)
	fmt.Println(string(mapBytes))
	
	response1Data := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	response1Bytes, _ := json.Marshal(response1Data)
	fmt.Println(string(response1Bytes))
	
	response2Data := &response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	response2Bytes, _ := json.Marshal(response2Data)
	fmt.Println(string(response2Bytes))

	bytes := []byte(`{"num": 6.13, "strs": ["a", "b"]}`)
	var data1 map[string]interface{}
	if error := json.Unmarshal(bytes, &data1); error != nil {
		panic(error)
	}

	number := data1["num"].(float64)
	fmt.Println(number)

	strings := data1["strs"].([]interface{})
	string1 := strings[0].(string)
	fmt.Println(string1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	result := response2{}
	json.Unmarshal([]byte(str), &result)
	fmt.Println(result)
	fmt.Println(result.Fruits[0])

	encoder := json.NewEncoder(os.Stdout)
	data2 := map[string]int{"apple": 5, "lettuce": 7}
	encoder.Encode(data2)
}


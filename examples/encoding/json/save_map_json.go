package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	A [2]int
	B map[int]string
}

func main() {
	data := calcOnce()

	fmt.Println(data) // {[101 102] map[1:Hello  2:World.]}
}

func calcOnce() Data {
	const once = "map.json"
	rd, err := ioutil.ReadFile(once)
	if err != nil {
		//calc and save once:
		data := Data{[2]int{103, 104}, map[int]string{13: "Hello ", 14: "World."}}
		buf, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(buf))
		err = ioutil.WriteFile(once, buf, 0666)
		if err != nil {
			panic(err)
		}
		return data
	}
	var d *Data
	err = json.Unmarshal(rd, &d)
	if err != nil {
		panic(err)
	}
	return *d
}

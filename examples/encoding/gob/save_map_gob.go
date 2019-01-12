package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Data struct {
	A [2]int
	B map[int]string
}

func main() {
	data := calcOnce()

	fmt.Println(data) // {[1010 102] map[2:World. 1:Hello ]}
}

func calcOnce() Data {
	const once = "map.bin"
	rd, err := ioutil.ReadFile(once)
	if err != nil {
		//calc and save once:
		data := Data{[2]int{101, 102}, map[int]string{1: "Hello ", 2: "World."}}
		buf := &bytes.Buffer{}
		err = gob.NewEncoder(buf).Encode(data)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(once, buf.Bytes(), 0666)
		if err != nil {
			panic(err)
		}
		return data
	}
	var d Data
	err = gob.NewDecoder(bytes.NewReader(rd)).Decode(&d)
	if err != nil {
		panic(err)
	}
	return d
}

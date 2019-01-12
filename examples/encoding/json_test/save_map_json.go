package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const fileName = "map.json"

// Data type wraps a map
type Data struct {
	Map map[int]string
}

func main() {
	data := loadMap()

	fmt.Println(data)
	// add data
	data.Map[100] = "Apple"

	dumpMap(data)

}

func dumpMap(data Data) {
	// write to file
	buf, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileName, buf, 0666)
	if err != nil {
		panic(err)
	}

}

func loadMap() Data {
	// load from file
	oldData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return Data{map[int]string{}}
	}
	var d *Data
	err = json.Unmarshal(oldData, &d)
	if err != nil {
		panic(err)
	}
	return *d
}

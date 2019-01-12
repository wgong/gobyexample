package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const fileName = "robotnames.json"

// Data type wraps a map
type Data struct {
	Map map[string]string
}

func main() {
	data := loadMap()

	fmt.Println(data)
	// add data
	key, val := "Google", "Alphabet"
	if oldVal, ok := data.Map[key]; ok {
		data.Map[oldVal] = key
	}
	data.Map[key] = val

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
		return Data{map[string]string{}}
	}
	var d *Data
	err = json.Unmarshal(oldData, &d)
	if err != nil {
		panic(err)
	}
	return *d
}

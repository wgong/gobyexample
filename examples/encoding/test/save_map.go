package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const fileNames = "names.json"

type nameMap map[string]string

var mNames nameMap

// Data store
type Data struct {
	mapNames nameMap
}

/*
func init() {
	loadNames()
}
*/

func loadNames() {
	bExist, err := exists(fileNames)
	if err != nil {
		panic(err)
	}
	if !bExist {
		mNames = make(map[string]string)
		return
	}
	buf, err := ioutil.ReadFile(fileNames)
	if err != nil {
		panic(err)
	}
	var d *Data
	if err := json.Unmarshal(buf, &d); err != nil {
		panic(err)
	}
	mNames = (*d).mapNames
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func dumpNames(m nameMap) {
	data := Data{m}
	buf, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	err = ioutil.WriteFile(fileNames, buf, 0666)
	if err != nil {
		panic(err)
	}
}

func main() {
	loadNames()

	mNames["Hello"] = "你好"
	mNames["World"] = "世界"

	fmt.Println(mNames)

	dumpNames(mNames)
}

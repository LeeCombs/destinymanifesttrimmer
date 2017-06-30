package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "os"
)

var (
//
)

func main() {
	fmt.Println("Running trimmer")

	file, _ := ioutil.ReadFile("DestinyManifest.json")

	var manifest map[interface{}]interface{}
	json.Unmarshal(file, &manifest)

	fmt.Println("Type %T", manifest.manifest)
}

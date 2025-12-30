package test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string
var iJsonData map[string]any

const stJsonFileName = "dir.json"

func loadJson() {
	stSeparator = string(filepath.Separator)
	stWorkDir, err := os.Getwd()
	if err != nil {
		panic("Get Work Dir Error" + err.Error())
	}
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]
	fmt.Println(stWorkDir)
	fmt.Println(stRootDir)

	gnJsonBytes, err := os.ReadFile(stWorkDir + stSeparator + stJsonFileName)
	if err != nil {
		panic("Load Json Data Error" + err.Error())
	}
	err = json.Unmarshal(gnJsonBytes, &iJsonData)
	if err != nil {
		panic("Unmarshal Json Data Error" + err.Error())
	}
	fmt.Println(iJsonData)
}

func parseMap(mapData map[string]any, stParentDir string) {
	//fmt.Println("stParentDir", stParentDir)
	for k, v := range mapData {
		//fmt.Println(k, v)
		switch v.(type) {
		case string:
			{
				path, _ := v.(string)
				if path == "" {
					continue
				}
				if stParentDir != "" {
					path = stParentDir + stSeparator + path
					if k == "text" {
						stParentDir = path
					}
				} else {
					stParentDir = path
				}
				createDir(path)
			}
		case []any:
			{
				parseArray(v.([]any), stParentDir)
			}
		}

	}
}

func parseArray(giJsonData []any, stParentDir string) {
	for _, v := range giJsonData {
		mapV, _ := v.(map[string]any)
		parseMap(mapV, stParentDir)
	}
}

func createDir(path string) {
	if path == "" {
		return
	}
	fmt.Println(path)
	err := os.MkdirAll(stRootDir+stSeparator+path, os.ModePerm)
	if err != nil {
		panic("Create Dir Error" + err.Error())
	}

}

func TestGenerateDir01(t *testing.T) {
	fmt.Println("test generate dir 01")
	loadJson()
	parseMap(iJsonData, "")
}

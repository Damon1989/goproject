package test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type Node struct {
	Text     string `json:"text"`
	Children []Node `json:"children"`
}

var stRootDir02 string
var stSeparator02 string
var iRootNode Node
var stJsonFileName02 = "dir.json"

func TestGenerateDir02(t *testing.T) {
	fmt.Println("test generate dir 02")
	loadJson02()
	parseNode(iRootNode, "")

}

func loadJson02() {
	stSeparator02 = string(filepath.Separator)
	stWorkDir, _ := os.Getwd()
	stRootDir02 = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator02)]

	gnJsonFileBytes, _ := os.ReadFile(stWorkDir + stSeparator02 + stJsonFileName02)
	err := json.Unmarshal(gnJsonFileBytes, &iRootNode)
	if err != nil {
		panic("Load Json Data Error" + err.Error())
	}
}

func parseNode(iNode Node, stParentDir string) {
	if iNode.Text != "" {
		createDir02(iNode, stParentDir)
	}
	if stParentDir != "" {
		stParentDir = stParentDir + stSeparator02
	}
	if iNode.Text != "" {
		stParentDir = stParentDir + iNode.Text
	}
	for _, iChildNode := range iNode.Children {
		parseNode(iChildNode, stParentDir)
	}
}

func createDir02(iNode Node, stParentDir string) {
	stDirPath := stRootDir02 + stSeparator02
	if stParentDir != "" {
		stDirPath = stDirPath + stParentDir + stSeparator02
	}
	stDirPath = stDirPath + iNode.Text
	fmt.Println(stDirPath)
	err := os.MkdirAll(stDirPath, os.ModePerm)
	if err != nil {
		panic("Create Dir Error" + err.Error())
	}
}

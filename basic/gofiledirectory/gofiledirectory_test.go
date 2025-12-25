package gofiledirectory

import "testing"

func TestCopyFile(t *testing.T) {
	CopyFile("./a/ta.txt", "./b/test_copy.txt")
}

func TestCopyFileStream(t *testing.T) {
	err := CopyFileStream("./a/ta.txt", "./b/test_copy_stream.txt")
	if err != nil {
		t.Error(err)
	}
}

func TestCreateDir(t *testing.T) {
	CreateDir("./testdir")
}

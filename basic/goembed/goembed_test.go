package goembed

import "testing"

func TestEmbedFile(t *testing.T) {
	EmbedFile()
}

func TestEmbedDir(t *testing.T) {
	EmbedDir()
}

func TestStaticEmbedServer(t *testing.T) {
	StaticEmbedServer()
}

func TestStaticRuntimeServer(t *testing.T) {
	StaticRuntimeServer()
}

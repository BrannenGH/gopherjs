//go:build js
// +build js

package net

import (
	"net"
	"os"
	"syscall"
	"testing"
)

func TestListenUnixSocket(t *testing.T) {
	path := "/tmp/test.sock"

	// Validate unlinked
	err := syscall.Unlink(path)
	if err != nil && !os.IsNotExist(err) {
		t.Fatalf("Failed to unlink file.")
	}

	l, err := net.Listen("unix", path)
	if err != nil {
		t.Fatalf("Failed to listen on socket")
	}

	l.Close()
}

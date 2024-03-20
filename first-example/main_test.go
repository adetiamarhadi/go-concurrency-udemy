package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "hello, universe") {
		t.Error("Expected to find hello universe, but it is not there")
	}

	if !strings.Contains(output, "hello, cosmos") {
		t.Error("Expected to find hello cosmos, but it is not there")
	}

	if !strings.Contains(output, "hello, world") {
		t.Error("Expected to find hello world, but it is not there")
	}
}

func Test_updateMessage(t *testing.T) {
	msg = "hello universe"

	wg.Add(1)
	go updateMessage("hello cosmos")
	wg.Wait()

	if !strings.Contains(msg, "hello cosmos") {
		t.Error("Expected to find hello cosmos, but it is not there")
	}
}

func Test_printMessage(t *testing.T) {
	msg = "hello world"

	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "hello world") {
		t.Error("Expected to find hello world, but it is not there")
	}
}

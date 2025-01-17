package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainDefault(t *testing.T) {
	origOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Args = []string{"lore"}
	os.Stdout = w
	main()
	os.Stdout = origOut
	w.Close()
	out, _ := io.ReadAll(r)
	readableOutput := string(out)
	expected := []string{
		"Welcome to Lore!",
		"monsters",
		"places",
		"npcs",
		"Get information about all available commands.",
	}
	for _, phrase := range expected {
		if !strings.Contains(readableOutput, phrase) {
			t.Errorf("Expected StdOut to contain %s", phrase)
		}
	}
}

func TestMainMonsters(t *testing.T) {

	origOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Args = []string{"lore", "monsters"}
	os.Stdout = w
	main()
	os.Stdout = origOut
	w.Close()
	out, _ := io.ReadAll(r)
	readableOutput := string(out)
	expected := []string{
		"Lore Monsters Help",
		"Monsters subcommands information",
		"*** monsters -i <monster-name>",
	}
	for _, phrase := range expected {
		if !strings.Contains(readableOutput, phrase) {
			t.Errorf("Expected StdOut to contain %s", phrase)
		}
	}
}

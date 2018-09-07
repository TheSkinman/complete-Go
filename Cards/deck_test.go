package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be \"Ace of Spades\", but got \"%v\"", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card to be \"Four of Clubs\", but got \"%v\"", d[len(d)-1])
	}
}

func TestSaveToFileandNewDeckFromFile(t *testing.T) {
	const fileName = "_decktesting"

	os.Remove(fileName)

	d := newDeck()

	d.saveToFile(fileName)

	nd := newDeckFromFile(fileName)

	if len(nd) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(nd))
	}

	if nd[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be \"Ace of Spades\", but got \"%v\"", nd[0])
	}

	if nd[len(nd)-1] != "Four of Clubs" {
		t.Errorf("Expected first card to be \"Four of Clubs\", but got \"%v\"", nd[len(nd)-1])
	}

	os.Remove(fileName)
}

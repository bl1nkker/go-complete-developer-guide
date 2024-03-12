package cards

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	expectedAmount := 44
    d := NewDeck()
	if len(d) != expectedAmount{
		t.Errorf("Expected deck length of %v, but got %v", expectedAmount, len(d))
	}
}

func TestDeal(t *testing.T){
	expectedAmount := 5
	d := NewDeck()
	hand, remainingDeck := Deal(d, expectedAmount)
	if len(hand) != expectedAmount{
		t.Errorf("Expected hand length of %v, but got %v", expectedAmount, len(hand))
	}
	if len(remainingDeck) != len(d) - expectedAmount{
		t.Errorf("Expected remainder length of %v, but got %v", len(d) - expectedAmount, len(remainingDeck))
	}
}

func TestToString(t *testing.T){
	expectedString := "Ace of Spades,Four of Hearts,Six of Clubs"
	d := Deck([]string{"Ace of Spades", "Four of Hearts", "Six of Clubs"})
	result := d.ToString()
	if result != expectedString{
		t.Errorf("Expected %v, but got %v", expectedString, result)
	}
}

func TestSaveToFile(t *testing.T){
	filename := "_decktesting" 
	os.Remove(filename)
	d := NewDeck()
	d.SaveToFile(filename)
	bs, _ := os.ReadFile(filename)
	if string(bs) != d.ToString(){
		t.Errorf("%v is not equal to %v", string(bs), d.ToString())
	}
	os.Remove(filename)
}

func TestNewDeckFromFile(t *testing.T){
	filename := "_decktesting" 
	os.Remove(filename)
	d := NewDeck()
	d.SaveToFile(filename)
	newDeck := NewDeckFromFile(filename)
	if newDeck.ToString() != d.ToString(){
		t.Errorf("%v is not equal to %v", newDeck.ToString(), d.ToString())
	}
	os.Remove(filename)
}
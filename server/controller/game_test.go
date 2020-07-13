package controller

import (
	"github.com/team142/angrychess/model"
	"testing"
)

//All tests that assess for the core functionality
func TestFindSpotEmptyGame(t *testing.T) {
	game := model.Game{Boards: 2}
	expectedFound := true
	found, spot := game.FindSpot()
	if !found {
		t.Errorf("Expected empty game to have found spot %v, got %v", expectedFound, found)
	}
	expectedSpot := 1
	if spot != expectedSpot {
		t.Errorf("Expected empty game to have TEAM spot %v, got %v", expectedSpot, spot)
	}

}

func TestFindSpotGameWithPlayers(t *testing.T) {
	game := &model.Game{Boards: 2, Players: make(map[int]*model.Player)}
	expectedFound := true
	for i := 1; i <= 4; i++ {
		found, spot := game.FindSpot()
		if found != expectedFound {
			t.Errorf("Expected game to have found spot %v, got %v", expectedFound, found)
		}
		expectedSpot := i
		if spot != expectedSpot {
			t.Errorf("Expected game to have TEAM spot %v, got %v", expectedSpot, spot)
		}
		joinGame(game, &model.Player{})
	}
}

func TestFindSpotGameFullPlayers(t *testing.T) {
	game := &model.Game{Boards: 2, Players: make(map[int]*model.Player)}
	for i := 1; i <= 4; i++ {
		joinGame(game, &model.Player{})
	}

	found, _ := game.FindSpot()
	expectedFound := false
	if found != expectedFound {
		t.Errorf("Expected game to have found spot %v, got %v", expectedFound, found)
	}

}

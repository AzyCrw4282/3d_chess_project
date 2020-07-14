package model

import "testing"

//All tests are performed in standard way to check for anomalies
//straight-forward. It can be ehnhacned for futher robust testing
func TestPlayer_SetTeamAndColorNormal(t *testing.T) {
	p := Player{}
	p.SetTeamColorAndBoard(1, 2)

	expectedColor := false
	if p.Color != expectedColor {
		t.Errorf("Expected color %v and got %v", expectedColor, p.Color)
	}

	expectedTeam := 1
	if p.Team != expectedTeam {
		t.Errorf("Expected team %v and got %v", expectedTeam, p.Team)
	}

	expectedBoard := 1
	if p.Board != expectedBoard {
		t.Errorf("Expected board %v and got %v", expectedBoard, p.Board)
	}

}

func TestPlayer_SetTeamAndColorNormal2(t *testing.T) {
	p := Player{}
	p.SetTeamColorAndBoard(3, 2)

	expectedColor := true
	if p.Color != expectedColor {
		t.Errorf("Expected color %v and got %v", expectedColor, p.Color)
	}

	expectedTeam := 2
	if p.Team != expectedTeam {
		t.Errorf("Expected team %v and got %v", expectedTeam, p.Team)
	}

	expectedBoard := 1
	if p.Board != expectedBoard {
		t.Errorf("Expected board %v and got %v", expectedBoard, p.Board)
	}

}

func TestPlayer_SetTeamAndColorAbsurd(t *testing.T) {
	p := Player{}
	p.SetTeamColorAndBoard(7, 4)

	expectedColor := true
	if p.Color != expectedColor {
		t.Errorf("Expected color %v and got %v", expectedColor, p.Color)
	}

	expectedTeam := 2
	if p.Team != expectedTeam {
		t.Errorf("Expected team %v and got %v", expectedTeam, p.Team)
	}

	expectedBoard := 3
	if p.Board != expectedBoard {
		t.Errorf("Expected board %v and got %v", expectedBoard, p.Board)
	}

}

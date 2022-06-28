package onefootball

import "testing"

func TestTeamPlayerFullName(t *testing.T) {
	player := TeamPlayer{
		Firstname: "",
		Lastname:  "",
		Name:      "Najeeb",
	}

	if player.FullName() != "Najeeb" {
		t.Errorf("expected fullname: %s, got %s", "Najeeb", player.FullName())
	}

	player.Firstname = "John"
	if player.FullName() != "John" {
		t.Errorf("expected fullname: %s, got %s", "John", player.FullName())
	}

	player.Firstname = ""
	player.Lastname = "Wick"
	if player.FullName() != "Wick" {
		t.Errorf("expected fullname: %s, got %s", "Wick", player.FullName())
	}

	player.Firstname = "John"
	player.Lastname = "Wick"
	if player.FullName() != "John Wick" {
		t.Errorf("expected fullname: %s, got %s", "John Wick", player.FullName())
	}
}

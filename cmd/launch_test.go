package cmd

import (
	"fmt"
	"testing"
)

func TestLaunch(t *testing.T) {
	expectedString := `As you set course for %v, you and ` +
		`your trusty space crew of %v brave ` +
		`souls launch the S.S. Galactic ` +
		"to the stars!\n" +
		`Zoom!`

	// Generate launch message.
	destination1 := "Mars"
	crew1 := 15
	message1 := launch("Mars", 15)
	if message1 != fmt.Sprintf(expectedString, destination1, crew1) {
		t.Error("The generated launch message was not the correct format.")
	}
}

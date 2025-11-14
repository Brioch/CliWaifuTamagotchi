package utils

import (
	"sync"

	"github.com/rivo/tview"
)

var (
	Happiness        = 1000                 // Initial happiness value
	CurrentBar       = ""                   // Place holder for the current bar state
	happinessMutex   sync.Mutex             // Mutex to protect concurrent access to Happiness
	HappinessBarRef  *tview.TextView        // Link to the happiness bar itself so we dynamically update it
)

// ==============================
// Load all expressions once
// ==============================
var (
	neutral       = LoadASCII("ascii-arts/expressions/neutral")
	neutralBlink  = LoadASCII("ascii-arts/expressions/neutral-blink")
	confused      = LoadASCII("ascii-arts/expressions/confused")
	confusedBlink = LoadASCII("ascii-arts/expressions/confused-blink")
	bored         = LoadASCII("ascii-arts/expressions/bored")
	boredBlink    = LoadASCII("ascii-arts/expressions/bored-blink")
	sad           = LoadASCII("ascii-arts/expressions/sad")
	sadBlink      = LoadASCII("ascii-arts/expressions/sad-blink")
)

// ==============================
// Decrease Happiness
// ==============================
func DecreaseHappiness(n int) {
	happinessMutex.Lock()
	defer happinessMutex.Unlock()

	if Happiness > 0 {
		Happiness -= n
		if Happiness < 0 {
			Happiness = 0
		}
		// Optimize it
		if Happiness%2 == 0{
			updateBar()
		}
	}
}

// ==============================
// Increase Happiness
// ==============================
func IncreaseHappiness(n int) {
	happinessMutex.Lock()
	defer happinessMutex.Unlock()

	if Happiness < 1000 {
		Happiness += n
		if Happiness > 1000 {
			Happiness = 1000
		}
		updateBar()
	}
}

// ==============================
// Returns visual bar string
// ==============================
func GetHappinessBar() {
	switch {
	case Happiness > 900:
		CurrentBar = "██████████"
		SetExpression(neutral, neutralBlink)
	case Happiness > 800:
		CurrentBar =  "█████████░"
		SetExpression(neutral, neutralBlink)
	case Happiness > 700:
		CurrentBar =  "████████░░"
		SetExpression(confused, confusedBlink)
	case Happiness > 600:
		CurrentBar =  "███████░░░"
		SetExpression(confused, confusedBlink)
	case Happiness > 500:
		CurrentBar =  "██████░░░░"
		SetExpression(bored, boredBlink)
	case Happiness > 400:
		CurrentBar =  "█████░░░░░"
		SetExpression(bored, boredBlink)
	case Happiness > 300:
		CurrentBar =  "████░░░░░░"
		SetExpression(bored, boredBlink)
	case Happiness > 200:
		CurrentBar =  "███░░░░░░░"
		SetExpression(sad, sadBlink)
	case Happiness > 100:
		CurrentBar =  "██░░░░░░░░"
		SetExpression(sad, sadBlink)
	case Happiness > 0:
		CurrentBar =  "█░░░░░░░░░"
		SetExpression(sad, sadBlink)
	default:
		CurrentBar =  "░░░░░░░░░░"
		SetExpression(sad, sadBlink)
	}
}

// ==============================
// Internal UI update
// ==============================
func updateBar() {
	if HappinessBarRef != nil && UIEventsChan != nil {
		GetHappinessBar()
		barText := CurrentBar
		UIEventsChan <- func() {
			HappinessBarRef.SetText(barText)
		}
	}
}

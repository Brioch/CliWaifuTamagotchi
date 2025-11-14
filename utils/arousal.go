package utils

import (
	"sync"
	"math/rand"
	"time"

	"github.com/rivo/tview"
)

var (
	Arousal           = 0                    // Initial arousal value
	CurrentArousalBar = ""                   // Place holder for the current arousal bar state
	arousalMutex      sync.Mutex             // Mutex to protect concurrent access to Arousal
	ArousalBarRef     *tview.TextView        // Link to the arousal bar itself so we dynamically update it
	ArousalMessages   []string               // Custom messages for high arousal
	isArousalHigh     bool                   // Flag to track if arousal is already high
)

// ==============================
// Load all expressions once
// ==============================
var (
	happy      = LoadASCII("ascii-arts/expressions/-happy")
	happyBlink = LoadASCII("ascii-arts/expressions/wink")
)

// ==============================
// Decrease Arousal
// ==============================
func DecreaseArousal(n int) {
	arousalMutex.Lock()
	defer arousalMutex.Unlock()

	if Arousal > 0 {
		Arousal -= n
		if Arousal < 0 {
			Arousal = 0
		}
		// Optimize it
		if Arousal%2 == 0 {
			updateArousalBar()
		}
	}
}

// ==============================
// Increase Arousal
// ==============================
func IncreaseArousal(n int) {
	arousalMutex.Lock()
	defer arousalMutex.Unlock()

	if Arousal < 1000 {
		Arousal += n
		if Arousal > 1000 {
			Arousal = 1000
		}
		updateArousalBar()
	}
}

// ==============================
// Returns visual bar string
// ==============================
func GetArousalBar() {
	switch {
	case Arousal > 900:
		CurrentArousalBar = "♥♥♥♥♥♥♥♥♥♥"
		SetExpression(happy, happyBlink)
		if !isArousalHigh && ChatBoxRef != nil && UIEventsChan != nil && len(ArousalMessages) > 0 {
			rand.Seed(time.Now().UnixNano())
			line := ArousalMessages[rand.Intn(len(ArousalMessages))]
			UIEventsChan <- func() {
				ChatBoxRef.SetText("Waifu: " + line)
			}
			isArousalHigh = true
		}
	case Arousal > 800:
		CurrentArousalBar = "♥♥♥♥♥♥♥♥♥♡"
		SetExpression(happy, happyBlink)
		isArousalHigh = false
	case Arousal > 700:
		CurrentArousalBar = "♥♥♥♥♥♥♥♥♡♡"
		isArousalHigh = false
	case Arousal > 600:
		CurrentArousalBar = "♥♥♥♥♥♥♥♡♡♡"
		isArousalHigh = false
	case Arousal > 500:
		CurrentArousalBar = "♥♥♥♥♥♥♡♡♡♡"
		isArousalHigh = false
	case Arousal > 400:
		CurrentArousalBar = "♥♥♥♥♥♡♡♡♡♡"
		isArousalHigh = false
	case Arousal > 300:
		CurrentArousalBar = "♥♥♥♥♡♡♡♡♡♡"
		isArousalHigh = false
	case Arousal > 200:
		CurrentArousalBar = "♥♥♥♡♡♡♡♡♡♡"
		isArousalHigh = false
	case Arousal > 100:
		CurrentArousalBar = "♥♥♡♡♡♡♡♡♡♡"
		isArousalHigh = false
	case Arousal > 0:
		CurrentArousalBar = "♥♡♡♡♡♡♡♡♡♡"
		isArousalHigh = false
	default:
		CurrentArousalBar = "♡♡♡♡♡♡♡♡♡♡"
		isArousalHigh = false
	}
}

// ==============================
// Internal UI update
// ==============================
func updateArousalBar() {
	if ArousalBarRef != nil && UIEventsChan != nil {
		GetArousalBar()
		barText := CurrentArousalBar
		UIEventsChan <- func() {
			ArousalBarRef.SetText(barText)
			// Show/hide H action based on arousal
			if Arousal > 900 {
				if ActionSpaceRef != nil && HSceneActionIndex != -1 {
					found := false
					for i := 0; i < ActionSpaceRef.GetItemCount(); i++ {
						mainText, _ := ActionSpaceRef.GetItemText(i)
						if mainText == "!!!" {
							found = true
							break
						}
					}
					if !found {
						// Re-add the item if it's not there
						ActionSpaceRef.InsertItem(HSceneActionIndex, "!!!", "  Do something naughty.", 'h', HSceneSelectedFunc)
					}
				}
			} else {
				if ActionSpaceRef != nil && HSceneActionIndex != -1 {
					for i := 0; i < ActionSpaceRef.GetItemCount(); i++ {
						mainText, _ := ActionSpaceRef.GetItemText(i)
						if mainText == "!!!" {
							ActionSpaceRef.RemoveItem(i)
							break
						}
					}
				}
			}
		}
	}
}

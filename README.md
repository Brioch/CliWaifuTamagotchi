# 🫂 CliWaifuTamagotchi

Preview:
![Result](result.gif)
---
![Reactions](reactions.jpg)

![Repo size](https://img.shields.io/github/repo-size/HenryLoM/CliWaifuTamagotchi?color=lightgrey)
![Commits](https://img.shields.io/github/commit-activity/t/HenryLoM/CliWaifuTamagotchi/main?color=blue)
![Last commit](https://img.shields.io/github/last-commit/HenryLoM/CliWaifuTamagotchi?color=informational)
![License](https://img.shields.io/github/license/HenryLoM/CliWaifuTamagotchi?color=orange)

## 📑 Table of Contents
- [✨ Overview](#-overview)
- [🎬 Launching Process](#-launching-process)
- [🎨 Customization](#-customization)
- [📂 Project Structure](#-project-structure)
- [⚙️ Core Scripts](#-core-scripts)
    - [launch.go](#launchgo)
    - [utils/app-utils.go](#utilsapp-utilsgo)
    - [utils/commands-utils.go](#utilscommands-utilsgo)
    - [utils/happiness.go](#utilshappinessgo)
    - [utils/palette-utils.go](#utilspalette-utilsgo)
    - [utils/settings.go](#utilssettingsgo)
- [📜 Notes](#-notes)

---

## ✨ Overview
CliWaifuTamagotchi is a **terminal-based companion** that:
- Renders **ASCII face, expressions, and clothes**.
- Provides a small set of **interactions**: Encourage, Dress Up, Quit.
- Uses a **persistent color palette** stored in `~/.config/cliwaifutamagotchi/palette.json`.
- Minimal UI built using **`tview` and `tcell`**.

---

## 🎬 Launching Process

1. **Clone repository**
```bash
git clone https://github.com/HenryLoM/CliWaifuTamagotchi.git
cd CliWaifuTamagotchi
````

2. **Ensure Go 1.20+ is installed**

```bash
go version
```

3. **Build the binary**

```bash
go build -o cliWT
```

4. **Run the app**

```bash
./cliWT
```

- **Or run directly for development**

```bash
go run launch.go
```

> **💡 Notes**
>
> * First run creates `~/.config/cliwaifutamagotchi/palette.json` on its own if missing.
> * On macOS, ensure your terminal supports **true color** for best rendering.

---

## 🎨 Customization

1. Palette<br>
JSON file is in `~/.config/cliwaifutamagotchi/` ; Named `palette.json`<br>
JSON file's structure:
```
{
  "background": "#1e1e2e",
  "foreground": "#cdd6f4",
  "border": "#cba6f7",
  "accent": "#eba0ac",
  "title": "#b4befe"
}
```
> Note: default palette is Catppuchin (Mocha)

2. Settings<br>
JSON file is in `~/.config/cliwaifutamagotchi/` ; Named `settings.json`<br>
JSON file's structure:
```
{
  "name": "Waifu"
  "defaultMessage": "..."
}
```

---

## 📂 Project Structure

```
CliWaifuTamagotchi/
│
├── README.md
├── LICENSE
├── .gitignore
├── result.gif
├── reactions.jpg
├── go.mod
├── go.sum
├── launch.go                       # Main file that launches the project
│
├── ascii-arts/
│   ├── clothes/                    # ASCII bodies
│   └── expressions/                # ASCII heads
│
├── assets/
│   └── words-of-encouragement.txt  # List of lines for the first function
│
└── utils/
    ├── app-utils.go                # Main helpers
    ├── commands-utils.go           # Functions for the Action Space
    ├── happiness.go                # Happiness scoring system
    ├── palette-utils.go            # Handling palette out of the file
    └── settings.go                 # Handling settings out of the file
```

---

## ⚙️ Core Scripts

### **launch.go**

* Loads ASCII **head, blink frames, and body**.
* Displays **actions menu**: Encourage, Dress Up, Quit.
* Handles **user input** (keys and navigation).
* Queues UI updates safely using `app.QueueUpdateDraw` via `UIEventsChan` that keeps UI changes in order.

### **utils/app-utils.go**

* Helper functions for **loading ASCII files**.
* Manages **UI rendering** and **widget updates**.

### **utils/commands-utils.go**

* Implements **interactions logic**:

  * `Encourage`: random encouraging phrase + happy frame.
  * `DressUp`: swaps body/outfit based on selection.
  * `BackgroundMode`: fills the TUI with Waifu, removing all of the odd elements.
* Caches **clothes in memory** to reduce disk reads.

### **utils/happiness.go**

* Handles the bar and changes emotions of the avatar.
* Handles the happiness scores.

### **utils/palette-utils.go**

* Loads palette from `~/.config/cliwaifutamagotchi/palette.json`.
* Creates **default palette** if missing.
* Provides **color application** helpers.

### **utils/settings.go**

* Loads settings from `~/.config/cliwaifutamagotchi/palette.json`.
* Creates **default settings** if missing.

---

## 📜 Notes

* Missing/malformed ASCII files may cause a wrong output; handle carefully if modifying assets inside the structure.
* Future plans:
    * More interactions (feeding, timed events, stats).
    * Save selected outfit and preferences.
    * Unit tests and error handling improvements.
    * Packaging for release binaries.

---

⤴︎ Return to the [📑 Table of Contents](#-table-of-contents) ⤴︎

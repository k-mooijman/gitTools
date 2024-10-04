package lib

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type setting struct {
	Type  string
	Value string
}

type settings struct {
	Info    string `json:"info"`
	Setting map[string]*setting
}

func DoFile() {
	SetSettings()
}

func SetSettings() {

	green := color.New(color.FgGreen).SprintFunc()
	//yellow := color.New(color.FgYellow).SprintFunc()

	settings := settings{}
	content, err := os.ReadFile("settings.json")
	if err != nil {
		settings.Setting = make(map[string]*setting)
		settings.Setting[""] = &setting{Type: "3", Value: "3"}
		fmt.Printf("No File Creating new")

	} else {
		err = json.Unmarshal(content, &settings)
		if err != nil {
			settings.Setting = make(map[string]*setting)
			settings.Setting[""] = &setting{Type: "3", Value: "3"}
			fmt.Printf("No Content Creating new")
		}
	}

	if settings.Setting == nil {
		settings.Setting = make(map[string]*setting)
	}

	settings.Info = "test"
	//fmt.Printf("  Remote: %s -- %s \n", green(settings), yellow(settings))
	settings.Setting["new2"] = &setting{Type: "new", Value: "new"}
	fmt.Printf("  settings: %s \n", green(settings))

	file, _ := json.MarshalIndent(settings, "", " ")

	fmt.Printf("  file: %s  \n", file)

	_ = os.WriteFile("settings.json", file, 0644)

}

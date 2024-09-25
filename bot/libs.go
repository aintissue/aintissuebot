package bot

import "encoding/json"

// Generates link for a project
func generateLink(projectName string) string {
	link := "https://t.me/"

	if conf.Dev {
		link += "Dev"
	}

	link += BotName + "?start=" + projectName

	return link
}

// Pretty print object or variable
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

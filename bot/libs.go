package bot

// Generates link for a project
func generateLink(projectName string) string {
	link := "https://t.me/"

	if conf.Dev {
		link += "Dev"
	}

	link += BotName + "?start=" + projectName

	return link
}

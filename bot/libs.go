package bot

func generateLink(projectName string) string {
	link := "https://t.me/"

	if conf.Dev {
		link += "Dev"
	}

	link += BotName + "?start=" + projectName

	return link
}

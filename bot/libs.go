package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Generates link for a project
func generateLink(projectName string) string {
	projectName = normalizeNs(projectName)

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

func getLoginLink(user *User) string {
	ll := "Finish sign in by clicking the link bellow:\n\nhttps://app.aintissue.com/login/%d/%s"

	resp, err := http.Get(fmt.Sprintf("http://localhost:7331/login/%d", user.TelegramId))
	if err != nil {
		loge(err)
	}
	defer resp.Body.Close()

	var lar *LoginApiResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		loge(err)
	}

	if err := json.Unmarshal(body, lar); err != nil {
		loge(err)
	}

	return fmt.Sprintf(ll, user.TelegramId, lar.SessionId)
}

type LoginApiResponse struct {
	TelegramId int64  `json:"telegram_id"`
	SessionId  string `json:"session_id"`
}

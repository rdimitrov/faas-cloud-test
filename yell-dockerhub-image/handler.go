package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Handle a serverless request
func Handle(req []byte) string {

	var mapRequest map[string]interface{}
	err := json.Unmarshal(req, &mapRequest)

	if err != nil {
		return fmt.Sprintf("Missing or failed request. Error: %s\n", err)
	}

	pusher := mapRequest["push_data"].(map[string]interface{})["pusher"]
	repoName := mapRequest["repository"].(map[string]interface{})["repo_name"]
	tag := mapRequest["push_data"].(map[string]interface{})["tag"]

	ret := fmt.Sprintf("Thanks, %s! You successfully pushed to Docker Hub the following image - %s:%s! %s", pusher, repoName, tag, getEmoticons())

	testReply(ret)

	return fmt.Sprintf(`{"text":"%s"}`, ret)
}

func getEmoticons() string {
	var rStr string
	var pool = []string{":openfaas:", ":whale:", ":thumbsup:", ":wave:", ":sunglasses:", ":ok_hand:", ":chart_with_upwards_trend:", ":sunrise:", ":smiley:", ":smiley_cat:", ":parrot:", ":rocket:", ":100:", ":muscle:", ":signal_strength:", ":man-cartwheeling:"}

	rand.Seed(time.Now().UnixNano())
	randomize := rand.Perm(len(pool))

	for _, i := range randomize[:5] {
		rStr = rStr + pool[i] + " "
	}
	return rStr
}

func testReply(s string) {
	url := "https://webhook.site/aad552dc-fa1d-4464-b39b-0e62ff640532"
	var jsonStr = []byte(fmt.Sprintf(`{"text":"%s"}`, s))
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Do(req)
}

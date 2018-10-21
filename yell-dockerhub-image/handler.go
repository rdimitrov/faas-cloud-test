package function

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

	return fmt.Sprintf("Thanks, %s! You successfully pushed to Docker Hub the following image - %s:%s! %s", pusher, repoName, tag, getEmoticons())
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

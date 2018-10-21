package function

import (
	"encoding/json"
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {

	var mapreq map[string]interface{}
	err := json.Unmarshal(req, &mapreq)

	if err != nil {
		return fmt.Sprintf("Missing or failed request. Error: %s\n", err)
	}

	pusher := mapreq["push_data"].(map[string]interface{})["pusher"]
	repo_name := mapreq["repository"].(map[string]interface{})["repo_name"]
	tag := mapreq["push_data"].(map[string]interface{})["tag"]

	return fmt.Sprintf("Thanks, %s! You successfully pushed to Docker Hub the following image - %s:%s\n", pusher, repo_name, tag)
}

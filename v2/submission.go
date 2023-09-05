package jotform

import (
	"strconv"
	"strings"
)

// DeleteSubmission
// Delete a single submission
// sid (int64): You can get submission IDs when you call /form/{id}/submissions.
// Returns status of request.
func (client jotformAPIClient) DeleteSubmission(sid int64) ([]byte, error) {
	return client.executeHttpRequest("submission/"+strconv.FormatInt(sid, 10), nil, "DELETE")
}

// EditSubmission
// Edit a single submission
// sid (int64): You can get submission IDs when you call /form/{id}/submissions.
// submission (map[string]string): New submission data with question IDs.
// Returns status of request.
func (client jotformAPIClient) EditSubmission(sid int64, submission map[string]string) ([]byte, error) {
	data := make(map[string]string)

	for k := range submission {
		if strings.Contains(k, "_") && k != "created_at" {
			data["submission["+k[0:strings.Index(k, "_")]+"]["+k[strings.Index(k, "_")+1:]+"]"] = submission[k]
		} else {
			data["submission["+k+"]"] = submission[k]
		}
	}

	return client.executeHttpRequest("submission/"+strconv.FormatInt(sid, 10), data, "POST")
}

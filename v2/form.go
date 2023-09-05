package jotform

import (
	"strconv"
	"strings"
)

// GetForm
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// Returns form ID, status, update and creation dates, submission count etc.
func (client jotformAPIClient) GetForm(formID int64) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10), "", "GET")
}

// GetFormQuestions
// Get a list of all questions on a form.
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// Returns question properties of a form.
func (client jotformAPIClient) GetFormQuestions(formID int64) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/questions", "", "GET")
}

// GetFormQuestion
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// qid (int): Identifier for each question on a form. You can get a list of question IDs from /form/{id}/questions.
// Returns question properties like required and validation.
func (client jotformAPIClient) GetFormQuestion(formID int64, qid int) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/question/"+strconv.Itoa(qid), "", "GET")
}

// GetFormSubmission
// List of a form submissions.
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// offset (string): Start of each result set for form list.
// limit (string): Number of results in each result set for form list.
// filter (map[string]string): Filters the query results to fetch a specific form range.
// orderBy (string): Order results by a form field name.
// Returns submissions of a specific form.
func (client jotformAPIClient) GetFormSubmissions(formID int64, offset string, limit string, filter map[string]string, orderBy string) ([]byte, error) {
	var params = createConditions(offset, limit, filter, orderBy)

	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/submissions", params, "GET")
}

// CreateFormSubmission
// Submit data to this form using the API
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// submission (map[string]string): Submission data with question IDs.
// Returns posted submission ID and URL.
func (client jotformAPIClient) CreateFormSubmission(formId int64, submission map[string]string) ([]byte, error) {
	data := make(map[string]string)

	for k := range submission {
		if strings.Contains(k, "_") {
			data["submission["+k[0:strings.Index(k, "_")]+"]["+k[strings.Index(k, "_")+1:]+"]"] = submission[k]
		} else {
			data["submission["+k+"]"] = submission[k]
		}
	}

	return client.executeHttpRequest("form/"+strconv.FormatInt(formId, 10)+"/submissions", data, "POST")
}

// CreateFormSubmissions
// Submit data to this form using the API
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// submission (map[string]string): Submission data with question IDs.
// Returns posted submission ID and URL.
func (client jotformAPIClient) CreateFormSubmissions(formId int64, submission []byte) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formId, 10)+"/submissions", submission, "PUT")
}

// GetFormFiles
// List of files uploaded on a form
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// Returns uploaded file information and URLs on a specific form.
func (client jotformAPIClient) GetFormFiles(formID int64) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/files", "", "GET")
}

// GetFormWebhooks
// Get list of webhooks for a form
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// Returns list of webhooks for a specific form.
func (client jotformAPIClient) GetFormWebhooks(formID int64) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/webhooks", "", "GET")
}

// CreateFormWebhook
// Add a new webhook
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// webhookURL (string): Webhook URL is where form data will be posted when form is submitted.
// Returns list of webhooks for a specific form.
func (client jotformAPIClient) CreateFormWebhook(formId int64, webhookURL string) ([]byte, error) {
	params := map[string]string{
		"webhookURL": webhookURL,
	}

	return client.executeHttpRequest("form/"+strconv.FormatInt(formId, 10)+"/webhooks", params, "POST")
}

// Delete a specific webhook of a form.
// Add a new webhook
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// webhookID (int64): You can get webhook IDs when you call /form/{formID}/webhooks.
// Returns remaining webhook URLs of form.
func (client jotformAPIClient) DeleteFormWebhook(formID int64, webhookID int64) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/webhooks/"+strconv.FormatInt(webhookID, 10), nil, "DELETE")
}

// GetFormProperties
// Get a list of all properties on a form
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// Returns form properties like width, expiration date, style etc.
func (client jotformAPIClient) GetFormProperties(formID int64) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/properties", "", "GET")
}

// GetFormReports
// Get all the reports of a form, such as excel, csv, grid, html, etc.
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// Returns list of all reports in a form, and other details about the reports such as title.
func (client jotformAPIClient) GetFormReports(formID int64) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/reports", "", "GET")
}

// CloneForm
// Clone a single form.
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// Returns status of request.
func (client jotformAPIClient) CloneForm(formID int64) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/clone", nil, "POST")
}

// DeleteFormQuestion
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// qid (int): Identifier for each question on a form. You can get a list of question IDs from /form/{id}/questions.
// Returns status of request.
func (client jotformAPIClient) DeleteFormQuestion(formID int64, qid int) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/question/"+strconv.Itoa(qid), nil, "DELETE")
}

// CreateFormQuestion
// Add new question to specified form.
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// questionProperties (map[string]string): New question properties like type and text.
// Returns properties of new question.
func (client jotformAPIClient) CreateFormQuestion(formID int64, questionProperties map[string]string) ([]byte, error) {
	question := make(map[string]string)

	for k := range questionProperties {
		question["question["+k+"]"] = questionProperties[k]
	}

	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/questions", question, "POST")
}

// CreateFormQuestion
// Add new question to specified form.
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// questions ([]byte): New question properties like type and text.
// Returns properties of new question.
func (client jotformAPIClient) CreateFormQuestions(formID int64, questions []byte) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/questions", questions, "PUT")
}

// EditFormQuestion
// Add or edit a single question properties
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// qid (int): Identifier for each question on a form. You can get a list of question IDs from /form/{id}/questions.
// questionProperties (map[string]string): New question properties like type and text.
// Returns edited property and type of question.
func (client jotformAPIClient) EditFormQuestion(formID int64, qid int, questionProperties map[string]string) ([]byte, error) {
	question := make(map[string]string)

	for k := range questionProperties {
		question["question["+k+"]"] = questionProperties[k]
	}

	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/question/"+strconv.Itoa(qid), question, "POST")
}

// DeleteForm
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// Returns properties of deleted form.
func (client jotformAPIClient) DeleteForm(formID int64) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10), nil, "DELETE")
}

// GetFormProperty
// Get a specific property of the form.]
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// propertyKey (string): You can get property keys when you call /form/{id}/properties.
// Returns given property key value.
func (client jotformAPIClient) GetFormProperty(formID int64, propertyKey string) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/properties/"+propertyKey, "", "POST")
}

// SetFormProperties
// Add or edit properties of a specific form
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// formProperties ([]byte): New properties like label width.
// Returns edited properties.
func (client jotformAPIClient) SetMultipleFormProperties(formID int64, formProperties []byte) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/properties", formProperties, "PUT")
}

// SetFormProperties
// Add or edit properties of a specific form
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// formProperties (map[string]string): New properties like label width.
// Returns edited properties.
func (client jotformAPIClient) SetFormProperties(formID int64, formProperties map[string]string) ([]byte, error) {
	properties := make(map[string]string)

	for k := range formProperties {
		properties["properties["+k+"]"] = formProperties[k]
	}

	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/properties", properties, "POST")
}

package jotform

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const defaultBaseURL = "https://api.jotform.com"
const apiVersion = "v1"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type jotformAPIClient struct {
	apiKey     string
	outputType string
	debugMode  bool
	HttpClient HttpClient
	BaseURL    string
}

type JotformAPIClient interface {
	GetOutputType() string
	SetOutputType(value string)
	GetDebugMode() bool
	SetDebugMode(value bool)

	GetUser() []byte
	GetUsage() []byte
	GetForms(offset string, limit string, filter map[string]string, orderBy string) []byte
	GetSubmissions(offset string, limit string, filter map[string]string, orderBy string) []byte
	GetSubusers() []byte
	GetFolders() []byte
	GetReports() []byte
	GetSettings() []byte
	UpdateSettings(settings map[string]string) []byte
	GetHistory(action string, date string, sortBy string, startDate string, endDate string) []byte
	GetForm(formID int64) []byte
	GetFormQuestions(formID int64) []byte
	GetFormQuestion(formID int64, qid int) []byte
	GetFormSubmissions(formID int64, offset string, limit string, filter map[string]string, orderBy string) []byte
	CreateFormSubmission(formId int64, submission map[string]string) []byte
	CreateFormSubmissions(formId int64, submission []byte) []byte
	GetFormFiles(formID int64) []byte
	GetFormWebhooks(formID int64) []byte
	CreateFormWebhook(formId int64, webhookURL string) []byte
	DeleteFormWebhook(formID int64, webhookID int64) []byte
	GetSubmission(sid int64) []byte
	GetReport(reportID int64) []byte
	GetFolder(folderID string) []byte
	CreateFolder(folderProperties map[string]string) []byte
	DeleteFolder(folderID string) []byte
	UpdateFolder(folderID string, folderProperties []byte) []byte
	AddFormsToFolder(folderID string, formIDs []string) []byte
	AddFormToFolder(folderID string, formID string) []byte
	GetFormProperties(formID int64) []byte
	GetFormReports(formID int64) []byte
	CreateReport(formID int64, report map[string]string) []byte
	GetFormProperty(formID int64, propertyKey string) []byte
	DeleteSubmission(sid int64) []byte
	EditSubmission(sid int64, submission map[string]string) []byte
	CloneForm(formID int64) []byte
	DeleteFormQuestion(formID int64, qid int) []byte
	CreateFormQuestion(formID int64, questionProperties map[string]string) []byte
	CreateFormQuestions(formID int64, questions []byte) []byte
	EditFormQuestion(formID int64, qid int, questionProperties map[string]string) []byte
	SetFormProperties(formID int64, formProperties map[string]string) []byte
	SetMultipleFormProperties(formID int64, formProperties []byte) []byte
	CreateForm(form map[string]interface{}) []byte
	CreateForms(form []byte) []byte
	DeleteForm(formID int64) []byte
	RegisterUser(userDetails map[string]string) []byte
	LoginUser(credentials map[string]string) []byte
	LogoutUser() []byte
	GetPlan(planName string) []byte
	DeleteReport(reportID int64) []byte
}

func NewJotFormAPIClient(apiKey string, outputType string, debugMode bool) *jotformAPIClient {
	client := &jotformAPIClient{
		apiKey:     apiKey,
		outputType: strings.ToLower(outputType),
		debugMode:  debugMode,
		HttpClient: &http.Client{
			Timeout: time.Second * 60,
			Transport: &http.Transport{
				// We occasionally see EOF responses, which may be caused
				// by net/http's default connection reuse.
				DisableKeepAlives: true,
			},
		},
		BaseURL: defaultBaseURL,
	}

	return client
}

func (client jotformAPIClient) GetOutputType() string       { return client.outputType }
func (client *jotformAPIClient) SetOutputType(value string) { client.outputType = value }

func (client jotformAPIClient) GetDebugMode() bool       { return client.debugMode }
func (client *jotformAPIClient) SetDebugMode(value bool) { client.debugMode = value }

func (client jotformAPIClient) debug(str interface{}) {
	if client.debugMode {
		fmt.Println(str)
	}
}

func (client jotformAPIClient) newRequest(requestPath string, params interface{}, method string) *http.Request {
	if client.outputType != "json" {
		requestPath = requestPath + ".xml"
	}

	var path = client.BaseURL + "/" + apiVersion + "/" + requestPath
	client.debug(path)
	client.debug(params)

	var request *http.Request

	if method == "GET" {
		if params != "" {
			data := params.(map[string]string)
			values := make(url.Values)

			for k := range data {
				values.Set(k, data[k])
			}
			path = path + "?" + values.Encode()
		}

		request, _ = http.NewRequest("GET", path, nil)
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else if method == "POST" {
		data := params.(map[string]string)
		values := make(url.Values)

		for k := range data {
			values.Set(k, data[k])
		}

		request, _ = http.NewRequest("POST", path, strings.NewReader(values.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else if method == "DELETE" {
		request, _ = http.NewRequest("DELETE", path, nil)
	} else if method == "PUT" {
		parameters := params.([]byte)
		request, _ = http.NewRequest("PUT", path, bytes.NewBuffer(parameters))
	}

	request.Header.Add("apiKey", client.apiKey)
	return request
}

func (client jotformAPIClient) executeHttpRequest(requestPath string, params interface{}, method string) ([]byte, error) {

	response, err := client.HttpClient.Do(
		client.newRequest(requestPath, params, method),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if client.outputType == "json" {
		var f interface{}
		json.Unmarshal(contents, &f)
		result, ok := f.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unexpected non-json response")
		}
		content, err := json.Marshal(result["content"])

		if err != nil {
			return nil, err
		} else {
			return content, nil
		}
	} else if client.outputType == "xml" {
		var f interface{}
		xml.Unmarshal(contents, &f)
		return contents, nil
	}

	return nil, nil
}

func createConditions(offset string, limit string, filter map[string]string, orderby string) map[string]string {
	args := map[string]interface{}{
		"offset":  offset,
		"limit":   limit,
		"filter":  filter,
		"orderby": orderby,
	}

	params := make(map[string]string)

	for k := range args {
		if k == "filter" {
			filterObj, err := json.Marshal(filter)

			if err == nil {
				params["filter"] = string(filterObj)
			}
		} else {
			params[k] = args[k].(string)
		}
	}
	return params
}

func createHistoryQuery(action string, date string, sortBy string, startDate string, endDate string) map[string]string {
	args := map[string]string{
		"action":    action,
		"date":      date,
		"sortBy":    sortBy,
		"startDate": startDate,
		"endDate":   endDate,
	}

	params := make(map[string]string)

	for k := range args {
		if args[k] != "" {
			params[k] = args[k]
		}
	}
	return params
}

// GetSubmission
// Get submission data
// sid (int64): You can get submission IDs when you call /form/{id}/submissions.
// Returns information and answers of a specific submission.
func (client jotformAPIClient) GetSubmission(sid int64) ([]byte, error) {
	return client.executeHttpRequest("user/submission/"+strconv.FormatInt(sid, 10), "", "GET")
}

// GetFolder
// folderID (int64): You can get a list of folders from /user/folders.
// Returns a list of forms in a folder, and other details about the form such as folder color.
func (client jotformAPIClient) GetFolder(folderID string) ([]byte, error) {
	return client.executeHttpRequest("folder/"+folderID, "", "GET")
}

// GetPlan
// Get details of a plan
// planName (string): Name of the requested plan. FREE, PREMIUM etc.
// Returns details of a plan
func (client jotformAPIClient) GetPlan(planName string) ([]byte, error) {
	return client.executeHttpRequest("system/plan/"+planName, "", "GET")
}

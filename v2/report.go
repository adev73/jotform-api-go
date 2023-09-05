package jotform

import "strconv"

// GetReport
// Get report details
// reportID (int64): You can get a list of reports from /user/reports.
// Returns properties of a speceific report like fields and status.
func (client jotformAPIClient) GetReport(reportID int64) ([]byte, error) {
	return client.executeHttpRequest("user/report/"+strconv.FormatInt(reportID, 10), "", "GET")
}

// CreateReport
// Create new report of a form
// formID (int64): Form ID is the numbers you see on a form URL. You can get form IDs when you call /user/forms.
// report (map[string]string): Report details. List type, title etc.
// Returns report details and URL.
func (client jotformAPIClient) CreateReport(formID int64, report map[string]string) ([]byte, error) {
	return client.executeHttpRequest("form/"+strconv.FormatInt(formID, 10)+"/reports", report, "POST")
}

// DeleteReport
// reportID (int64): You can get a list of reports from /user/reports.
// Returns status of request.
func (client jotformAPIClient) DeleteReport(reportID int64) ([]byte, error) {
	return client.executeHttpRequest("report/"+strconv.FormatInt(reportID, 10), nil, "DELETE")
}

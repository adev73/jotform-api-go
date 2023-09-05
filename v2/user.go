package jotform

// GetUser
// Get user account details for a JotForm user.
// Returns user account type, avatar URL, name, email, website URL and account limits.
func (client jotformAPIClient) GetUser() ([]byte, error) {
	return client.executeHttpRequest("user", "", "GET")
}

// GetForms
// Get a list of forms for this account
// offset (string): Start of each result set for form list.
// limit (string): Number of results in each result set for form list.
// filter (map[string]string): Filters the query results to fetch a specific form range.
// orderBy (string): Order results by a form field name.
// Returns basic details such as title of the form, when it was created, number of new and total submissions.
func (client jotformAPIClient) GetForms(offset string, limit string, filter map[string]string, orderBy string) ([]byte, error) {
	var params = createConditions(offset, limit, filter, orderBy)

	return client.executeHttpRequest("user/forms", params, "GET")
}

// GetSubmissions
// Get a list of submissions for this account
// offset (string): Start of each result set for form list.
// limit (string): Number of results in each result set for form list.
// filter (map[string]string): Filters the query results to fetch a specific form range.
// orderBy (string): Order results by a form field name.
// Returns basic details such as title of the form, when it was created, number of new and total submissions.
func (client jotformAPIClient) GetSubmissions(offset string, limit string, filter map[string]string, orderBy string) ([]byte, error) {
	var params = createConditions(offset, limit, filter, orderBy)

	return client.executeHttpRequest("user/submissions", params, "GET")
}

// GetSubusers
// Get a list of sub users for this account
// Returns list of forms and form folders with access privileges.
func (client jotformAPIClient) GetSubusers() ([]byte, error) {
	return client.executeHttpRequest("user/subusers", "", "GET")
}

// GetFolders
// Get a list of form folders for this account
// Returns name of the folder and owner of the folder for shared folders.
func (client jotformAPIClient) GetFolders() ([]byte, error) {
	return client.executeHttpRequest("user/folders", "", "GET")
}

// GetReports
// List of URLS for reports in this account
// Returns reports for all of the forms. ie. Excel, CSV, printable charts, embeddable HTML tables.
func (client jotformAPIClient) GetReports() ([]byte, error) {
	return client.executeHttpRequest("user/reports", "", "GET")
}

// Update user's settings
// New user setting values with setting keys
// Returns changes on user settings
func (client jotformAPIClient) GetSettings() ([]byte, error) {
	return client.executeHttpRequest("user/settings", "", "GET")
}

// GetSettings
// Get user's settings for this account
// Returns user's time zone and language.
func (client jotformAPIClient) UpdateSettings(settings map[string]string) ([]byte, error) {
	return client.executeHttpRequest("user/settings", settings, "POST")
}

// GetHistory
// Get user activity log
// action (string): Filter results by activity performed. Default is 'all'.
// date (string): Limit results by a date range. If you'd like to limit results by specific dates you can use startDate and endDate fields instead.
// sortBy (string): Lists results by ascending and descending order.
// startDate (string): Limit results to only after a specific date. Format: MM/DD/YYYY.
// endDate (string): Limit results to only before a specific date. Format: MM/DD/YYYY.
// Returns activity log about things like forms created/modified/deleted, account logins and other operations.
func (client jotformAPIClient) GetHistory(action string, date string, sortBy string, startDate string, endDate string) ([]byte, error) {
	var params = createHistoryQuery(action, date, sortBy, startDate, endDate)

	return client.executeHttpRequest("user/history", params, "GET")
}

// RegisterUser
// Register with username, password and email
// userDetails (map[string]string): Username, password and email to register a new user
// Returns new user's details
func (client jotformAPIClient) RegisterUser(userDetails map[string]string) ([]byte, error) {
	return client.executeHttpRequest("user/register", userDetails, "POST")
}

// LoginUser
// Login user with given credentials
// credentials (map[string]string): Username, password, application name and access type of user
// Returns logged in user's settings and app key
func (client jotformAPIClient) LoginUser(credentials map[string]string) ([]byte, error) {
	return client.executeHttpRequest("user/login", credentials, "POST")
}

// LogoutUser
// Logout user
// Returns status of request
func (client jotformAPIClient) LogoutUser() ([]byte, error) {
	return client.executeHttpRequest("user/logout", "", "GET")
}

// CreateForm
// Create a new form
// form ([]byte): Questions, properties and emails of new form.
// Returns new form.
func (client jotformAPIClient) CreateForm(form map[string]interface{}) ([]byte, error) {
	params := make(map[string]string)

	for formKey, formValue := range form {
		if formKey == "properties" {
			properties := formValue

			for properyKey, propertyValue := range properties.(map[string]string) {
				params[formKey+"["+properyKey+"]"] = propertyValue
			}
		} else {
			formItem := formValue

			for formItemKey, formItemValue := range formItem.(map[string]interface{}) {
				item := formItemValue

				for itemKey, itemValue := range item.(map[string]string) {
					params[formKey+"["+formItemKey+"]["+itemKey+"]"] = itemValue
				}
			}
		}
	}

	return client.executeHttpRequest("user/forms", params, "POST")
}

// Create new forms
// Create a new form
// form ([]byte): Questions, properties and emails of forms.
// Returns new forms.
func (client jotformAPIClient) CreateForms(form []byte) ([]byte, error) {
	return client.executeHttpRequest("user/forms", form, "PUT")
}

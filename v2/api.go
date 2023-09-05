package jotform

// GetUsage
// Get number of form submissions received this month
// Returns number of submissions, number of SSL form submissions, payment form submissions and upload space used by user.
func (client jotformAPIClient) GetUsage() ([]byte, error) {
	return client.executeHttpRequest("user/usage", "", "GET")
}

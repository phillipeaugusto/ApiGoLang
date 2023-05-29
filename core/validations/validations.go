package validations

type ValidationData struct {
	Valid   bool
	Field   string
	Message string
}

type ErrorData struct {
	Field   string
	Message string
}

func Validate(validation *[]ValidationData) bool {
	var validate = true
	for _, item := range *validation {
		if item.Valid {
			validate = false
			break
		}
	}
	return validate
}

func Notifications(validation *[]ValidationData) []ErrorData {
	var listErr = new([]ErrorData)
	for _, item := range *validation {
		if item.Valid {
			*listErr = append(*listErr, ErrorData{Field: item.Field, Message: item.Message})
		}
	}
	return *listErr
}

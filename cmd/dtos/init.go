package dtos

type Base struct {
	Success bool     `json:"success"`
	Meta    []string `json:"meta,omitempty"`
}

func Success(result map[string]interface{}) map[string]interface{} {
	result["success"] = true
	return result
}

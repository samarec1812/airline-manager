package http

type createAirlineRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type deleteAirlineRequest struct {
	Code string `json:"code"`
}

type createAccountRequest struct {
	SchemaID int64 `json:"schema_id"`
}

type deleteAccountRequest struct {
	AccountID int64 `json:"account_id"`
}

type createSchemaRequest struct {
	Name string `json:"name"`
}

type deleteSchemaRequest struct {
	SchemaID int64 `json:"schema_id"`
}

func AirlineErrorResponse(err error) map[string]any {
	resp := make(map[string]any)

	resp["data"] = nil
	resp["error"] = err

	return resp
}

func AirlineSuccessResponse(code string) map[string]any {
	resp := make(map[string]any)

	resp["data"] = code
	resp["error"] = nil

	return resp
}

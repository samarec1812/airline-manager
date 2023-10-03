package http

type createAirlineRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type deleteAirlineRequest struct {
	Code string `json:"code"`
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

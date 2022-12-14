package response

type (
	StandardResponse struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
	}
)

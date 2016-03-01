package dealer

type JsonResponse struct {
	id   string                 `json:"id"`
	time int                    `json:"time"`
	data map[string]interface{} `json:"data,omitempty"`
}

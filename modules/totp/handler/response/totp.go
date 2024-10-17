package response

type CreateTotpResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Barcode string `json:"barcode"`
	Url     string `json:"url"`
}

package models

type Totp struct {
	Name        string `json:"name"`
	Generate    bool   `json:"generate"`
	Exported    bool   `json:"exported"`
	KeySize     int    `json:"key_size"`
	Issuer      string `json:"issuer"`
	AccountName string `json:"account_name"`
	Period      int    `json:"period"`
	Algorithm   string `json:"algorithm"`
	Digits      int    `json:"digits"`
	Skew        int    `json:"skew"`
	QrSize      int    `json:"qr_size"`
}

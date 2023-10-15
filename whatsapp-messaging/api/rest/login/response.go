package login

type LoginResponse struct {
	QRCode    string `json:"qr_code"`
	QRTimeout int    `json:"timeout"`
}

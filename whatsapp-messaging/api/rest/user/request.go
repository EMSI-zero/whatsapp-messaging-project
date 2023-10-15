package user

type AddJIDRequest struct{
	UserID int64 `json:"user_id"`
	JID string `json:"jid"`
}
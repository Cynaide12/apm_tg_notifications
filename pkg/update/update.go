package update

type MessageRequest struct {
    UserID []int `json:"user_id"`
	Message string `json:"message"`
}
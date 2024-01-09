package helper

var Body struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

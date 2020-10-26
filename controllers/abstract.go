package controllers

type Response struct {
	Status bool
	Message string
	Data interface{} `json:",omitempty"`
	Error Error
}

type Error struct {
	ErrorCode uint8 `json:",omitempty"`
	ErrorMessage string `json:",omitempty"`
}

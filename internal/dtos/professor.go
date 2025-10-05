package dtos

type ProfessorShortResponse struct{
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Title          string  `json:"title,omitempty"`
	ImageURL       string  `json:"image_url,omitempty"`
	UniRoomAddress string  `json:"uni_room_address,omitempty"`
	// ReviewCount    int     `json:"review_count"`
}

type ProfessorFullResponse struct {
	ProfessorShortResponse
    Phone string `json:"phone"`
    Description string `json:"description"`
    Classes []CourseShortResponse `json:"courses"`
}
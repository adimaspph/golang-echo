package domain

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
	Role  string `json:"role" form:"role" query:"role"`
}

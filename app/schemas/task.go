package schemas

type TaskCreateSchema struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

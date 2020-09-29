package model

type CategoryID int64

type Category struct {
	ID   CategoryID `validate:"required"`
	Name string     `validate:"required"`
}

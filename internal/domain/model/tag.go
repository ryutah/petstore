package model

type TagID uint64

type Tag struct {
	ID   TagID  `validate:"required"`
	Name string `validate:"required"`
}

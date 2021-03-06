package todo

import (
	databases "github.com/pascallin/gin-server/internal/pkg"
)

type Todo struct {
	databases.GormModel
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTodoInput struct {
	Title       string `form:"title" xml:"title" json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTodoInput struct {
	ID          uint64 `uri:"id" binding:"required" json:"id"`
	Title       string `form:"title" xml:"title" json:"title"`
	Description string `form:"title" xml:"title" json:"description"`
}

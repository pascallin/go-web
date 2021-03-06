package task

import (
	"github.com/pascallin/gin-server/internal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Title string             `bson:"title" json:"title"`
}

func (t *Task) New() *Task {
	return &Task{
		ID:    primitive.NewObjectID(),
		Title: t.Title,
	}
}

type CreateTaskInput struct {
	Title string `binding:"required"`
}

type UpdateTaskInput struct {
	Title string `form:"title"`
}

type GetTaskListInput struct {
	internal.Pagination
	Title string `form:"title"`
}

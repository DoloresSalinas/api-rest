package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Titulo        string             `bson:"titulo" json:"titulo"`
	Descripcion   string             `bson:"descripcion" json:"descripcion"`
	FechaInicio   time.Time          `bson:"fecha_inicio" json:"fecha_inicio"`
	FechaDeadline time.Time          `bson:"fecha_deadline" json:"fecha_deadline"`
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type TaskInput struct {
	Titulo        string `json:"titulo"`
	Descripcion   string `json:"descripcion"`
	FechaInicio   string `json:"fecha_inicio"`   // formato: dd/mm/yyyy
	FechaDeadline string `json:"fecha_deadline"` // formato: dd/mm/yyyy
	UserID        string `json:"user_id"`        // string para parsearlo luego a ObjectID
}

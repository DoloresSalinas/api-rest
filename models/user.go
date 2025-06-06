package models

import ( 
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre string `bson:"nombre" json:"nombre"`
	App string `bson:"app" json:"app"`
	Apm string `bson:"apm" json:"apm"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	FechaNacimiento time.Time  `bson:"fecha_nacimiento" json:"fecha_nacimiento"`
	PreguntaSecreta string `bson:"pregunta_secreta" json:"pregunta_secreta"`
	RespuestaSecreta string `bson:"respuesta_secreta" json:"respuesta_secreta"`

}
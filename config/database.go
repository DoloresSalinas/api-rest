package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient   *mongo.Client
	MongoDatabase *mongo.Database
	UsersCol      *mongo.Collection
	TasksCol      *mongo.Collection
)

// ConectarMongo establece la conexión con MongoDB
func ConectarMongo() {
	// Puedes cambiar la URI según tu entorno local/remoto
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	// Opciones del cliente
	clientOpts := options.Client().ApplyURI(mongoURI)

	// Conectar a MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("Error al conectar a MongoDB:", err)
	}

	// Verificar conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("No se pudo hacer ping a MongoDB:", err)
	}

	log.Println("Conectado a MongoDB")

	MongoClient = client
	MongoDatabase = client.Database("api_rest") // Nombre de la base de datos

	// Inicializar colecciones 
	UsersCol = MongoDatabase.Collection("users")
	TasksCol = MongoDatabase.Collection("tasks")

	if UsersCol == nil {
		log.Fatal("UsersCol es nil después de inicializar")
	}
}

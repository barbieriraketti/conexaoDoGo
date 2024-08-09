package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"zuchi/db" // Importa o pacote db

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Obter o cliente MongoDB
	client, err := db.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com o MongoDB estabelecida com sucesso!")

	// Acessando o banco de dados e a coleção
	database := client.Database("zuchi")
	collection := database.Collection("zuchi_task")

	fmt.Printf("Conectado ao banco de dados %s e à coleção %s\n", database.Name(), collection.Name())

	// Criar um documento BSON
	documento := bson.D{
		{Key: "date", Value: primitive.NewDateTimeFromTime(time.Now())},
		{Key: "activities", Value: bson.A{
			bson.D{
				{Key: "description", Value: "Task 1"},
				{Key: "initialTime", Value: "12:00"},
				{Key: "finalTime", Value: "13:00"},
			},
			bson.D{
				{Key: "description", Value: "Task 2"},
				{Key: "initialTime", Value: "13:00"},
				{Key: "finalTime", Value: "14:00"},
			},
		}},
	}

	// Inserir o documento na coleção
	insertResult, err := collection.InsertOne(context.TODO(), documento)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Documento inserido com ID: %v\n", insertResult.InsertedID)

	// Fechar a conexão quando terminar
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Espera para verificar o fechamento da conexão
	time.Sleep(2 * time.Second)
	fmt.Println("Conexão encerrada.")
}

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PatientInfo represents the structure of our resource
type PatientInfo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Age       string             `bson:"age"`
	Situation string             `bson:"situation"`
	Rank      int                `bson:"rank"`
	UserID    int                `bson:"user_id,omitempty"`
	Status    bool               `bson:"status"`
	Telephone string             `bson:"telephone"`
	Location  string             `bson:"location"`
	ChunkID   int                `bson:"chunk_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
}

var client *mongo.Client

func PostPatientInfo(c *gin.Context) {
	var newInfo PatientInfo
	if err := c.BindJSON(&newInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newInfo.ID = primitive.NewObjectID()
	newInfo.CreatedAt = time.Now()
	newInfo.UpdatedAt = time.Now()

	collection := client.Database("yourDatabaseName").Collection("patientInfo")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.InsertOne(ctx, newInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "The new Patient Info is correctly saved", "data": newInfo})
}

func UpdatePatientInfo(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("_id"))
	var updatedInfo PatientInfo
	if err := c.BindJSON(&updatedInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedInfo.UpdatedAt = time.Now()

	collection := client.Database("yourDatabaseName").Collection("patientInfo")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.UpdateByID(ctx, id, bson.M{"$set": updatedInfo})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "The field was successfully updated."})
}

func DeletePatientInfo(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("_id"))
	collection := client.Database("yourDatabaseName").Collection("patientInfo")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "The field was successfully deleted."})
}

func UpdateStatus(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("_id"))
	var statusUpdate struct {
		Status bool `json:"status"`
	}
	if err := c.BindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := client.Database("yourDatabaseName").Collection("patientInfo")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.UpdateByID(ctx, id, bson.M{"$set": bson.M{"status": statusUpdate.Status}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "The field was successfully updated."})
}

func UpdateSituation(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("_id"))
	var situationUpdate struct {
		Situation string `json:"situation"`
	}
	if err := c.BindJSON(&situationUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := client.Database("yourDatabaseName").Collection("patientInfo")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.UpdateByID(ctx, id, bson.M{"$set": bson.M{"situation": situationUpdate.Situation}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "The field was successfully updated."})
}

func UpdateRank(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("_id"))
	var rankUpdate struct {
		Rank int `json:"rank"`
	}
	if err := c.BindJSON(&rankUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := client.Database("yourDatabaseName").Collection("patientInfo")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.UpdateByID(ctx, id, bson.M{"$set": bson.M{"rank": rankUpdate.Rank}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "The field was successfully updated."})
}

func main() {
	// Set up Gin
	r := gin.Default()

	// Connect to MongoDB
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Define routes
	//r.GET("/api/patient-info", GetPatientInfos)
	r.POST("/api/patient-info", PostPatientInfo)
	r.PUT("/api/patient-info/:_id", UpdatePatientInfo)
	r.DELETE("/api/patient-info/:_id", DeletePatientInfo)
	r.PUT("/api/patient-info/status/:_id", UpdateStatus)
	r.PUT("/api/patient-info/situation/:_id", UpdateSituation)
	r.PUT("/api/patient-info/rank/:_id", UpdateRank)

	// Run the server
	r.Run(":8080")
}

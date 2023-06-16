package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ADEXITUM/calorie-counter/pkg/db"
	"github.com/ADEXITUM/calorie-counter/pkg/entities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var entryColletion *mongo.Collection = db.OpenCollection(db.Client, "calories")

// @Summary Get All Entries
// @Tags entries
// @Description Get all previously added entries.
// @ID get-all-entries
// @Produce json
// @Success 200 {object} gin.ResponseWriter
// @Failure 400 {object} http.ProtocolError
// @Failure 404 {object} http.ProtocolError
// @Failure 500 {object} http.ProtocolError
// @Failure default {object} http.ProtocolError
// @Router /entries [get]
func GetEntries(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	var entries []bson.M
	cursor, err := entryColletion.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("routes.GetEntries: #1\nError retrieving all entries\n%s\n\n", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"err": err.Error()})
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		log.Printf("routes.GetEntries: #2\nError iterating cursor\n%s\n\n", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entries)
}

// @Summary Get Entry by ID
// @Tags entries
// @Description Get entry by its ID.
// @ID get-entry-by-id
// @Accept json
// @Produce json
// @Success 200 {object} gin.ResponseWriter
// @Failure 400 {object} http.ProtocolError
// @Failure 404 {object} http.ProtocolError
// @Failure 500 {object} http.ProtocolError
// @Failure default {object} http.ProtocolError
// @Router /entries/{id} [get]
func GetEntryById(c *gin.Context) {
	EntryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(EntryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	var entry bson.M
	if err := entryColletion.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry); err != nil {
		log.Printf("routes.GetEntryById: #1\nError iterating cursor\n\n%s", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entry)
}

// @Summary Add Entry
// @Tags entries
// @Description Add entry of comsumed dish, its macroes and calories.
// @ID add-entry
// @Accept json
// @Produce json
// @Success 200 {object} gin.ResponseWriter
// @Failure 400 {object} http.ProtocolError
// @Failure 404 {object} http.ProtocolError
// @Failure 500 {object} http.ProtocolError
// @Failure default {object} http.ProtocolError
// @Router /entries/create [post]
func AddEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	var entry entities.Entry
	if err := c.BindJSON(&entry); err != nil {
		if err != nil {
			log.Printf("routes.AddEntry: #1\nError creating entry\n%s\n\n", err)
			c.JSON(http.StatusInternalServerError,
				gin.H{"err": err.Error()})
			return
		}
	}

	entry.ID = primitive.NewObjectID()
	result, err := entryColletion.InsertOne(ctx, entry)
	if err != nil {
		log.Printf("routes.AddEntry: #2\nError inserting entry\n%s\n\n", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Update Entry
// @Tags entries
// @Description Update entry with new information.
// @ID update-entry
// @Accept json
// @Produce json
// @Success 200 {object} gin.ResponseWriter
// @Failure 400 {object} http.ProtocolError
// @Failure 404 {object} http.ProtocolError
// @Failure 500 {object} http.ProtocolError
// @Failure default {object} http.ProtocolError
// @Router /entries/update/:id [put]
func UpdateEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	var entry entities.Entry

	if err := c.BindJSON(&entry); err != nil {
		log.Printf("routes.UpdateEntry: #1\nError binding JSON\n%s\n\n", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	result, err := entryColletion.ReplaceOne(ctx,
		bson.M{"_id": docID},
		bson.M{
			"dish":     entry.Dish,
			"macroes":  entry.Macroes,
			"calories": entry.Calories,
		},
	)
	if err != nil {
		log.Printf("routes.UpdateEntry: #2\nError updating entry\n%s\n\n", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Delete All Entries
// @Tags entries
// @Description Delete all entries in a "calories" collection.
// @ID delete-all-entries
// @Produce json
// @Success 200 {object} gin.ResponseWriter
// @Failure 400 {object} http.ProtocolError
// @Failure 404 {object} http.ProtocolError
// @Failure 500 {object} http.ProtocolError
// @Failure default {object} http.ProtocolError
// @Router /entries/delete/all [delete]
func DeleteAll(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	result, err := entryColletion.DeleteMany(ctx, bson.D{})
	if err != nil {
		log.Printf("routes.DeleteAll: #2\nError deleting all items from DB\n%s\n\n", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result.DeletedCount)
}

// @Summary Delete Entry
// @Tags entries
// @Description Delete entry by its ID.
// @ID delete-entry
// @Accept json
// @Produce json
// @Success 200 {object} gin.ResponseWriter
// @Failure 400 {object} http.ProtocolError
// @Failure 404 {object} http.ProtocolError
// @Failure 500 {object} http.ProtocolError
// @Failure default {object} http.ProtocolError
// @Router /entries/delete/{id} [delete]
func DeleteEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")

	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	result, err := entryColletion.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		log.Printf("routes.DeleteEntry: #1\nError deleting item from DB\n%s\n\n", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result.DeletedCount)
}

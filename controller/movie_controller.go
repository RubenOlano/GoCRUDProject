package controller

import (
	"CRUD/config"
	"CRUD/responses"
	"CRUD/schema"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var movieCollection = config.GetCollection("movies")
var validate = validator.New()

func CreateMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var movie schema.Movie
		defer cancel()

		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, responses.MovieResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if err := validate.Struct(movie); err != nil {
			c.JSON(http.StatusBadRequest, responses.MovieResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		new_movie := schema.Movie{
			Title:  movie.Title,
			Year:   movie.Year,
			Poster: movie.Poster,
			Id:     primitive.NewObjectID(),
		}

		result, err := movieCollection.InsertOne(ctx, new_movie)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MovieResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.MovieResponse{Status: http.StatusCreated, Message: "Success", Data: map[string]interface{}{"data": result}})

	}
}

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var movie schema.Movie

		err := movieCollection.FindOne(ctx, primitive.M{}).Decode(&movie)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MovieResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.MovieResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": movie}})
	}
}

func EditMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		movieId := c.Param("id")
		var movie schema.Movie
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(movieId)

		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, responses.MovieResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if err := validate.Struct(&movie); err != nil {
			c.JSON(http.StatusBadRequest, responses.MovieResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		update := bson.M{"Title": movie.Title, "Year": movie.Year, "Poster": movie.Poster}

		result, err := movieCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MovieResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updated_movie schema.Movie

		err = movieCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updated_movie)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MovieResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.MovieResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": updated_movie, "result": result}})
	}
}

func DeleteMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		movieId := c.Param("id")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(movieId)

		result, err := movieCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MovieResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.MovieResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"result": result}})
	}
}

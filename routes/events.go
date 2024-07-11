package routes

import (
	"net/http"
	"strconv"

	"events.com/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventById(int(eventId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.BindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// event.ID = len(models.GetAllEvents()) + 1
	event.UserId = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"status": "Event created successfully", "data": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventById(int(eventId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = context.BindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated := models.UpdateEvent(event)
	if updated != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": updated.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventById(int(eventId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deletedEvent := models.DeleteEvent(int(eventId))

	if deletedEvent != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": deletedEvent.Error()})
		return
	}

	// err = context.BindJSON(&event)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// updated := models.UpdateEvent(event)
	// if updated != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": updated.Error()})
	// 	return
	// }
	context.JSON(http.StatusOK, gin.H{"status": "Event deleted successfully", "data": event})
}

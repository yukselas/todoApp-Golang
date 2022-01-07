package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func getRoot(c *gin.Context) {
	ReplySuccess(c, gin.H{
		"version": VERSION,
	})
}

func getNotes(c *gin.Context) {
	ReplySuccess(c, gin.H{"notes": notes})
}

func getNoteByID(c *gin.Context) {
	id := c.Param("id")
	for _, note := range notes {
		if note.ID == id {
			ReplySuccess(c, gin.H{"note": note})
			return
		}
	}
	ReplyFail(c, gin.H{"error": "not found"})
}

func postNote(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		ReplyFail(c, gin.H{"error": err.Error()})
		return
	}
	if len(jsonData) == 0 {
		ReplyFail(c, gin.H{"error": "empty body"})
		return
	}

	var note Note
	err = json.Unmarshal(jsonData, &note)
	if err != nil {
		ReplyFail(c, gin.H{"error": err.Error()})
		return
	}

	AddNote(note)
	ReplySuccess(c, "ok")
}

func deleteNote(c *gin.Context) {
	id := c.Param("id")
	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			ReplySuccess(c, "ok")
			return
		}
	}
	ReplyFail(c, gin.H{"error": "not found"})
}

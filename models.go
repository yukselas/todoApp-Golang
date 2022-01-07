package main

import (
	"fmt"
	"time"
)

type Note struct {
	ID        string
	Text      string
	Timestamp int64
}

func NewNote(text string) Note {
	id := fmt.Sprintf("%d", len(notes)+1)
	return Note{
		ID:        id,
		Text:      text,
		Timestamp: time.Now().Unix(),
	}
}

func AddNote(note Note) {
	if note.ID == "" {
		note.ID = fmt.Sprintf("%d", len(notes)+1)
	}

	if note.Timestamp == 0 {
		note.Timestamp = time.Now().Unix()
	}

	notes = append(notes, note)
}

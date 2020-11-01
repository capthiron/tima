package note

import (
	"log"
	"time"
)

type Service interface {
	AddNote(note *Note) *Note
	GetNotesForDay(day time.Time) []Note
	UpdateNote(note *Note) *Note
	DeleteNote(note *Note)
}

func NewService() Service {
	return &defaultService{newRepository()}
}

type defaultService struct {
	repo repository
}

func (s defaultService) AddNote(note *Note) *Note {

	note.Start = time.Now()

	note, err := s.repo.addNote(note)
	if err != nil {
		log.Fatal(err)
	}

	return note
}

func (s defaultService) GetNotesForDay(day time.Time) []Note {

	notes, err := s.repo.getNotesForDay(day)
	if err != nil {
		log.Fatal(err)
	}

	return notes
}

func (s defaultService) UpdateNote(note *Note) *Note {

	updatedNote, err := s.repo.updateNote(note)
	if err != nil {
		log.Fatal(err)
	}

	return updatedNote
}

func (s defaultService) DeleteNote(note *Note) {
	err := s.repo.deleteNote(note.ID)
	if err != nil {
		log.Fatal(err)
	}
}

package note

import (
	"fmt"
	"github.com/capthiron/tima/db"
	"gorm.io/gorm"
	"time"
)

type repository interface {
	addNote(note *Note) (*Note, error)
	getNote(id uint) (*Note, error)
	getNotesForDay(day time.Time) ([]Note, error)
	updateNote(note *Note) (*Note, error)
	deleteNote(id uint) error
}

func newRepository() repository {
	dbConn := db.Connect()

	err := dbConn.AutoMigrate(&Note{})
	if err != nil {
		fmt.Println(err)
	}

	return &defaultRepository{dbConn}
}

type defaultRepository struct {
	db *gorm.DB
}

func (r defaultRepository) addNote(note *Note) (*Note, error) {
	err := r.db.Create(&note).Error
	if err != nil {
		err = fmt.Errorf("addNote(%v): %w", note, err)
	}
	return note, err
}

func (r defaultRepository) getNote(id uint) (*Note, error) {
	var note Note
	err := r.db.First(&note, id).Error
	if err != nil {
		err = fmt.Errorf("getNote(%v): %w", id, err)
	}
	return &note, err
}

func (r defaultRepository) getNotesForDay(day time.Time) (notes []Note, err error) {
	err = r.db.Where("strftime('%Y%j', start) = strftime('%Y%j', ?)", day).Find(&notes).Error
	if err != nil {
		err = fmt.Errorf("getNotesForDay(%v): %w", day, err)
	}
	return
}

func (r defaultRepository) updateNote(note *Note) (*Note, error) {
	var noteToUpdate Note
	err := r.db.First(&noteToUpdate, note.ID).Error

	if err != nil {
		err = fmt.Errorf("updateNote(%v) db.First: %w", note, err)
	}

	noteToUpdate.Description = note.Description
	noteToUpdate.Start = note.Start
	noteToUpdate.End = note.End
	noteToUpdate.Status = note.Status

	err = r.db.Save(&noteToUpdate).Error
	if err != nil {
		err = fmt.Errorf("updateNote(%v) db.Save: %w", note, err)
	}

	return &noteToUpdate, err
}

func (r defaultRepository) deleteNote(id uint) (err error) {
	if err = r.db.Delete(&Note{}, id).Error; err != nil {
		err = fmt.Errorf("deleteNote(%v): %w", id, err)
	}
	return
}

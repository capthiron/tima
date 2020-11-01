package note

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var repo repository

func beforeRepoTest() {
	if err := os.Setenv("PROFILE", "test"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repo = newRepository()
}

func TestDefaultRepository_AddNote(t *testing.T) {
	beforeRepoTest()

	addedNote, err := repo.addNote(&Note{})
	if err != nil {
		t.Error("Failed to add a note with error:", err)
	}

	if addedNote == nil {
		t.Error("Added note should not be nil")
	}

	if _, err := repo.getNote(addedNote.ID); err != nil {
		t.Error("Failed to get added note with error:", err)
	}
}

func TestDefaultRepository_GetNotesForDay(t *testing.T) {
	beforeRepoTest()

	timeLayout := "2006-01-02T15:04:05"

	today, err := time.Parse(timeLayout, "1997-06-13T00:00:00")
	if err != nil {
		t.Error(err)
	}

	yesterday, err := time.Parse(timeLayout, "1997-06-12T23:59:59")
	if err != nil {
		t.Error(err)
	}

	tomorrow, err := time.Parse(timeLayout, "1997-06-14T00:00:00")
	if err != nil {
		t.Error(err)
	}

	acceptedNote1 := Note{Description: "Note we want", Start: today}
	acceptedNote2 := Note{Description: "Another Note we want", Start: today}

	deniedNote1 := Note{Description: "Note we dont want", Start: yesterday}
	deniedNote2 := Note{Description: "Another Note we dont want", Start: yesterday}
	deniedNote3 := Note{Description: "A Note from the future", Start: tomorrow}
	deniedNote4 := Note{Description: "Another Note from the future", Start: tomorrow}

	_, _ = repo.addNote(&acceptedNote1)
	_, _ = repo.addNote(&acceptedNote2)
	_, _ = repo.addNote(&deniedNote1)
	_, _ = repo.addNote(&deniedNote2)
	_, _ = repo.addNote(&deniedNote3)
	_, _ = repo.addNote(&deniedNote4)

	notesForToday, err := repo.getNotesForDay(today)
	if err != nil {
		t.Error("Repository returned an error:", err)
	}

	if len(notesForToday) != 2 {
		t.Errorf("Expected exactly 2 notes but got %v", len(notesForToday))
	}

	for _, task := range notesForToday {
		if task.Start != today {
			t.Errorf("Expected note.StartTime to be %v but got %v", today, task.Start)
		}
	}
}

func TestDefaultRepository_UpdateNote(t *testing.T) {
	beforeRepoTest()

	note := Note{
		Description: "node",
		Start:       time.Time{},
		End:         time.Time{},
		Status:      InProgress,
	}
	_, _ = repo.addNote(&note)

	update := note
	update.Description = "note"
	update.Status = Done

	updated, err := repo.updateNote(&update)
	if err != nil {
		t.Error("Failed to update note with error:", err)
	}

	if updated == nil {
		t.Error("updateNote did not return a Note:", err)
	}

	if updated.Description != "note" {
		t.Errorf("updated.Description expected 'note' but got '%v'", updated.Description)
	}

	if updated.Status != Done {
		t.Errorf("updated.Status expected 'Done' but got '%v'", updated.Status)
	}
}

func TestDefaultRepository_DeleteNote(t *testing.T) {
	beforeRepoTest()

	noteToBeDeleted := Note{
		Description: "Should be deleted",
		Start:       time.Time{},
		End:         time.Time{},
		Status:      Done,
	}

	_, _ = repo.addNote(&noteToBeDeleted)

	if err := repo.deleteNote(noteToBeDeleted.ID); err != nil {
		t.Error("Failed to delete note with error:", err)
	}

	if _, err := repo.getNote(noteToBeDeleted.ID); err == nil {
		t.Error("Note should be deleted so an error should be raised")
	}
}

package note

import (
	"testing"
	"time"
)

// mock repository
type mockRepo struct{}

var addNoteMock func(note *Note) (*Note, error)
var getNoteMock func(id uint) (*Note, error)
var getNotesForDayMock func(day time.Time) ([]Note, error)
var updateNoteMock func(note *Note) (*Note, error)
var deleteNoteMock func(id uint) error

func (m mockRepo) addNote(note *Note) (*Note, error) {
	return addNoteMock(note)
}

func (m mockRepo) getNote(id uint) (*Note, error) {
	return getNoteMock(id)
}

func (m mockRepo) getNotesForDay(day time.Time) ([]Note, error) {
	return getNotesForDayMock(day)
}

func (m mockRepo) updateNote(note *Note) (*Note, error) {
	return updateNoteMock(note)
}

func (m mockRepo) deleteNote(id uint) error {
	return deleteNoteMock(id)
}

var service Service
var mockedRepo mockRepo

func beforeServiceTest() {
	mockedRepo = mockRepo{}
	service = defaultService{mockedRepo}
}

func TestDefaultService_AddNote(t *testing.T) {
	beforeServiceTest()

	noteToAdd := Note{
		Description: "new note",
	}

	if noteToAdd.ID != 0 {
		t.Errorf("noteToAdd.ID should be 0 but got %v", noteToAdd.ID)
	}

	noteAdded := noteToAdd
	noteAdded.ID = 1

	addNoteCalled := 0
	addNoteMock = func(note *Note) (*Note, error) {
		addNoteCalled++
		return &noteAdded, nil
	}

	if addNoteCalled != 0 {
		t.Errorf("addNote() should not have been called yet but instead got %v", addNoteCalled)
	}

	noteToAdd = *service.AddNote(&noteToAdd)

	if noteToAdd.ID != 1 {
		t.Errorf("id of noteToAdd should now be 1 but got %v", noteToAdd.ID)
	}

	if addNoteCalled != 1 {
		t.Errorf("addNote() should have been called but instead got %v", addNoteCalled)
	}
}

func TestDefaultService_GetNotesForDay(t *testing.T) {
	beforeServiceTest()

	getNotesForDayCalled := 0
	getNotesForDayMock = func(day time.Time) ([]Note, error) {
		getNotesForDayCalled++
		return []Note{}, nil
	}

	if getNotesForDayCalled != 0 {
		t.Errorf("getNotesForDay() should not have been called yet but got %v", getNotesForDayCalled)
	}

	notesForDay := service.GetNotesForDay(time.Now())

	if notesForDay == nil {
		t.Error("notesForDay should have an empty notes array but got nil")
	}

	if getNotesForDayCalled != 1 {
		t.Errorf("getNotesForDay() should have been called once but got %v", getNotesForDayCalled)
	}
}

func TestDefaultService_UpdateNote(t *testing.T) {
	beforeServiceTest()

	updateNoteCalled := 0
	updateNoteMock = func(note *Note) (*Note, error) {
		updateNoteCalled++
		return note, nil
	}

	if updateNoteCalled != 0 {
		t.Errorf("updateNote() should not have been called yet but got %v", updateNoteCalled)
	}

	updatedNote := service.UpdateNote(&Note{})

	if updatedNote == nil {
		t.Errorf("updatedNote should not be nil")
	}

	if updateNoteCalled != 1 {
		t.Errorf("updateNote() should have been called once but got %v", updateNoteCalled)
	}
}

func TestDefaultService_DeleteNote(t *testing.T) {
	beforeServiceTest()

	deleteNoteCalled := 0
	deleteNoteMock = func(id uint) error {
		deleteNoteCalled++
		return nil
	}

	if deleteNoteCalled != 0 {
		t.Errorf("deleteNote() should not have been called yet but got %v", deleteNoteCalled)
	}

	service.DeleteNote(&Note{})

	if deleteNoteCalled != 1 {
		t.Errorf("deleteNote() should have been called once but got %v", deleteNoteCalled)
	}
}

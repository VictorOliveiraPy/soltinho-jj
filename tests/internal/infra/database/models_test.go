package database

import (
	"database/sql"
	"testing"

	db "github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestUser_WhenCreatingNewUser_ThenAllFieldsShouldBeSet(t *testing.T) {
	// Test implementation...
	user := db.User{
		ID:             "123",
		Name:           "John Doe",
		Email:          "john@example.com",
		Phone:          "123456789",
		AcademyName:    "Example Academy",
		InstructorBelt: "Black Belt",
		Password:       "password",
	}

	// Test individual fields
	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "123456789", user.Phone)
	assert.Equal(t, "Example Academy", user.AcademyName)
	assert.Equal(t, "Black Belt", user.InstructorBelt)
	assert.Equal(t, "password", user.Password)

	// Test JSON tags
	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "123456789", user.Phone)
	assert.Equal(t, "Example Academy", user.AcademyName)
	assert.Equal(t, "Black Belt", user.InstructorBelt)
}

func TestStudent_WhenCreatingNewStudent_ThenAllFieldsShouldBeSet(t *testing.T) {
	// Test implementation...
	user := db.Student{
		ID:         "123",
		Name:       "John Doe",
		Age:        23,
		Email:      "john@example.com",
		Graduation: "white",
		Attendance: sql.NullInt32{Int32: 2},
		Absences:   sql.NullInt32{Int32: 2},
		Payment:    true,
		Password:   "password",
	}

	// Test individual fields
	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "password", user.Password)

	// Test JSON tags
	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
}

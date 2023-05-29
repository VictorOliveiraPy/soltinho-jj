package database

import (
	"testing"

	db "github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestUser_WhenCreatingNewUser_ThenAllFieldsShouldBeSet(t *testing.T) {
	user := db.User{
		ID:       "123",
		Username: "John Doe",
		Email:    "john@example.com",
		Password: "password",
		Active:   true,
		RoleID:   "1",
	}

	// Test individual fields
	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "John Doe", user.Username)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "password", user.Password)
	assert.Equal(t, true, user.Active)

}

func TestStudent_WhenCreatingNewStudent_ThenAllFieldsShouldBeSet(t *testing.T) {
	user := db.Student{
		ID:           "123",
		GymID:        "123",
		Name:         "John Doe",
		Graduation:   "white",
		Active:       true,
		TrainingTime: "1",
	}

	// Test individual fields
	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, true, user.Active)
	assert.Equal(t, "1", user.TrainingTime)
}

func TestGym_WhenCreatingNewSGym_ThenAllFieldsShouldBeSet(t *testing.T) {
	gym := db.Gym{
		ID:       "123",
		GymName:  "academia go",
		TeamName: "GF_TEAM",
		Active:   true,
	}

	// Test individual fields
	assert.Equal(t, "123", gym.ID)
	assert.Equal(t, "academia go", gym.GymName)
	assert.Equal(t, "GF_TEAM", gym.TeamName)
	assert.Equal(t, true, gym.Active)

}

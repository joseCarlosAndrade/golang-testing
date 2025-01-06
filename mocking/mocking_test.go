package mocking

import (
	"testing"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
	"unittest/mocks"
)

func TestGetUser(t *testing.T) {
	mockUserService := new(mocks.DataBaseInterface)

	// set up the expected calls. we're mocking the behavior of the GetUser method
	// so that when it's called with 1, it returns "user1" and nil error
	mockUserService.On("GetUser", 1).Return("user1", nil)


	// using the mocks
	user, err := mockUserService.GetUser(1)
	assert.NoError(t, err)

	assert.Equal(t, "user1", user)

	// assert that the expectations were met
	mockUserService.AssertExpectations(t)
}
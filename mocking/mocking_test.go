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

func TestGetUserTable(t *testing.T) {
	mockGetUser := new(mocks.DataBaseInterface)

	// defining test cases

	tests := []struct {
		name string
		inputId int
		mockReturn string
		mockError error
		expectedUser string
		expectedError error
	}{
		{
			name: "Test 1 - Valid user no error",
			inputId: 1,
			mockReturn: "user1",
			mockError: nil,
			expectedUser: "user1",
			expectedError: nil,
		},
		{
			name: "Test 2 - Error",
			inputId: -999,
			mockReturn: "",
			mockError: assert.AnError,
			expectedUser: "",
			expectedError: assert.AnError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t*testing.T){
			// mocking
			mockGetUser.On("GetUser", tc.inputId).Return(tc.mockReturn, tc.mockError)

			// calling the generated mock
			result, err := mockGetUser.GetUser(tc.inputId)

			if tc.expectedError != nil { // if we expect an error
				assert.Error(t, err) // assert that an error is returned
				assert.Equal(t, tc.expectedError, err) // assert that the error is the expected error
			} else { // if we don't expect an error
				assert.NoError(t, err) // assert that no error is returned
				assert.Equal(t, tc.expectedUser, result) // assert that the result is the expected result
			}
			// assert that the expectations were met
			mockGetUser.AssertExpectations(t)
		})
	}

}
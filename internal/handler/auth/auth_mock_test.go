// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package auth

import (
	"context"
	authModel "github.com/nabiel-syarif/playlist-api/internal/model/auth"
	"sync"
)

// Ensure, that UsecaseMock does implement Usecase.
// If this is not the case, regenerate this file with moq.
var _ Usecase = &UsecaseMock{}

// UsecaseMock is a mock implementation of Usecase.
//
// 	func TestSomethingThatUsesUsecase(t *testing.T) {
//
// 		// make and configure a mocked Usecase
// 		mockedUsecase := &UsecaseMock{
// 			LoginFunc: func(ctx context.Context, email string, password string) (authModel.JwtLoginData, error) {
// 				panic("mock out the Login method")
// 			},
// 			RegisterFunc: func(ctx context.Context, user authModel.UserRegistration) (authModel.UserPublic, error) {
// 				panic("mock out the Register method")
// 			},
// 		}
//
// 		// use mockedUsecase in code that requires Usecase
// 		// and then make assertions.
//
// 	}
type UsecaseMock struct {
	// LoginFunc mocks the Login method.
	LoginFunc func(ctx context.Context, email string, password string) (authModel.JwtLoginData, error)

	// RegisterFunc mocks the Register method.
	RegisterFunc func(ctx context.Context, user authModel.UserRegistration) (authModel.UserPublic, error)

	// calls tracks calls to the methods.
	calls struct {
		// Login holds details about calls to the Login method.
		Login []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Email is the email argument value.
			Email string
			// Password is the password argument value.
			Password string
		}
		// Register holds details about calls to the Register method.
		Register []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// User is the user argument value.
			User authModel.UserRegistration
		}
	}
	lockLogin    sync.RWMutex
	lockRegister sync.RWMutex
}

// Login calls LoginFunc.
func (mock *UsecaseMock) Login(ctx context.Context, email string, password string) (authModel.JwtLoginData, error) {
	if mock.LoginFunc == nil {
		panic("UsecaseMock.LoginFunc: method is nil but Usecase.Login was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Email    string
		Password string
	}{
		Ctx:      ctx,
		Email:    email,
		Password: password,
	}
	mock.lockLogin.Lock()
	mock.calls.Login = append(mock.calls.Login, callInfo)
	mock.lockLogin.Unlock()
	return mock.LoginFunc(ctx, email, password)
}

// LoginCalls gets all the calls that were made to Login.
// Check the length with:
//     len(mockedUsecase.LoginCalls())
func (mock *UsecaseMock) LoginCalls() []struct {
	Ctx      context.Context
	Email    string
	Password string
} {
	var calls []struct {
		Ctx      context.Context
		Email    string
		Password string
	}
	mock.lockLogin.RLock()
	calls = mock.calls.Login
	mock.lockLogin.RUnlock()
	return calls
}

// Register calls RegisterFunc.
func (mock *UsecaseMock) Register(ctx context.Context, user authModel.UserRegistration) (authModel.UserPublic, error) {
	if mock.RegisterFunc == nil {
		panic("UsecaseMock.RegisterFunc: method is nil but Usecase.Register was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		User authModel.UserRegistration
	}{
		Ctx:  ctx,
		User: user,
	}
	mock.lockRegister.Lock()
	mock.calls.Register = append(mock.calls.Register, callInfo)
	mock.lockRegister.Unlock()
	return mock.RegisterFunc(ctx, user)
}

// RegisterCalls gets all the calls that were made to Register.
// Check the length with:
//     len(mockedUsecase.RegisterCalls())
func (mock *UsecaseMock) RegisterCalls() []struct {
	Ctx  context.Context
	User authModel.UserRegistration
} {
	var calls []struct {
		Ctx  context.Context
		User authModel.UserRegistration
	}
	mock.lockRegister.RLock()
	calls = mock.calls.Register
	mock.lockRegister.RUnlock()
	return calls
}

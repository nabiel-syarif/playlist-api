package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	authModel "github.com/nabiel-syarif/playlist-api/internal/model/auth"
	"github.com/stretchr/testify/require"
)

func TestHandler_Register(t *testing.T) {
	type args struct {
		prepareRequest func() *http.Request
	}
	testCases := []struct {
		desc         string
		args         args
		wantResponse map[string]interface{}
		uc           Usecase
	}{
		{
			desc: "case 1 -> success register",
			args: args{
				prepareRequest: func() *http.Request {
					body := &bytes.Buffer{}
					writer := multipart.NewWriter(body)
					fw, _ := writer.CreateFormField("name")
					_, _ = io.Copy(fw, strings.NewReader("Nabiel"))

					fw, _ = writer.CreateFormField("email")
					_, _ = io.Copy(fw, strings.NewReader("nabiel@gmail.com"))

					fw, _ = writer.CreateFormField("password")
					_, _ = io.Copy(fw, strings.NewReader("rahasia"))

					writer.Close()
					req, err := http.NewRequest(http.MethodPost, "/v1/auth/register", bytes.NewReader(body.Bytes()))
					require.NoError(t, err)

					req.Header.Set("Content-Type", writer.FormDataContentType())
					return req
				},
			},
			wantResponse: map[string]interface{}{
				"status": "SUCCESS",
				"data": map[string]interface{}{
					"id":    float64(1),
					"name":  "Nabiel",
					"email": "nabiel@gmail.com",
				},
			},
			uc: &UsecaseMock{
				RegisterFunc: func(ctx context.Context, user authModel.UserRegistration) (authModel.UserPublic, error) {
					return authModel.UserPublic{
						Id:    1,
						Name:  user.Name,
						Email: user.Email,
					}, nil
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			h := New(tC.uc)

			router := chi.NewRouter()
			router.Method(http.MethodPost, "/v1/auth/register", http.HandlerFunc(h.Register))

			recorder := httptest.NewRecorder()
			request := tC.args.prepareRequest()
			router.ServeHTTP(recorder, request)
			var response interface{}
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			require.NoError(t, err, "unmarshall no error")
			require.Equal(t, tC.wantResponse, response)
		})
	}
}

func TestHandler_Login(t *testing.T) {
	type args struct {
		prepareRequest func() *http.Request
	}
	testCases := []struct {
		desc string
		args args
		uc   Usecase
	}{
		{
			desc: "case 1 -> success login",
			args: args{
				prepareRequest: func() *http.Request {
					body := &bytes.Buffer{}
					writer := multipart.NewWriter(body)
					fw, _ := writer.CreateFormField("email")
					_, _ = io.Copy(fw, strings.NewReader("nabiel@gmail.com"))

					fw, _ = writer.CreateFormField("password")
					_, _ = io.Copy(fw, strings.NewReader("rahasia"))

					writer.Close()
					req, err := http.NewRequest(http.MethodPost, "/v1/auth/login", bytes.NewReader(body.Bytes()))
					require.NoError(t, err)

					req.Header.Set("Content-Type", writer.FormDataContentType())
					return req
				},
			},
			uc: &UsecaseMock{
				LoginFunc: func(ctx context.Context, email, password string) (authModel.JwtLoginData, error) {
					return authModel.JwtLoginData{
						Token: "blablabla",
					}, nil
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			h := New(tC.uc)

			router := chi.NewRouter()
			router.Method(http.MethodPost, "/v1/auth/login", http.HandlerFunc(h.Login))

			recorder := httptest.NewRecorder()
			request := tC.args.prepareRequest()
			router.ServeHTTP(recorder, request)
			var response map[string]interface{}
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			require.NoError(t, err, "unmarshall no error")
			require.Contains(t, response, "data")
			require.Contains(t, response["data"], "access_token")
		})
	}
}

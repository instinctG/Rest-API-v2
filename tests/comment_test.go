//go:build e2e
// +build e2e

package tests

import (
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("instinct"))
	if err != nil {
		panic(err)
	}

	return tokenString

}

func TestPostComment(t *testing.T) {
	t.Run("can post comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().SetHeader("Authorization", "Bearer "+createToken()).
			SetBody(`{"slug":"/","author":"instinct","body":"hello world"}`).
			Post("http://localhost:8080/api/v2")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})
}

func TestGetComment(t *testing.T) {
	t.Run("can get comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().Get("http://localhost:8080/api/v2/d9f464d5-d050-49a4-b2a7-f25a37143e2b")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})
}

func TestDeleteComment(t *testing.T) {
	t.Run("delete comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().SetHeader("Authorization", "Bearer "+createToken()).
			Delete("http://localhost:8080/api/v2/1e8abc78-f8cf-43ff-a826-176ede6f871d")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())

	})

}

func TestUpdateComment(t *testing.T) {
	t.Run("update comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().SetHeader("Authorization", "Bearer "+createToken()).
			SetBody(`{"slug":"my girl","author":"instinct","body":"let it happen"}`).
			Put("http://localhost:8080/api/v2/d9f464d5-d050-49a4-b2a7-f25a37143e2b")

		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

}

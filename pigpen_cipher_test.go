package pigpen_cipher

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEncrypt(t *testing.T) {
	imageBytes, err := Encrypt("test", NewOptions().WithFontSize(80))
	assert.Nil(t, err)

	err = os.WriteFile("test.jpg", imageBytes, os.ModePerm)
	assert.Nil(t, err)
}

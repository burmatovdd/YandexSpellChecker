package YandexSpellChecker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithCorrectValues(t *testing.T) {
	spellChecker := newChecker("")
	res, err := spellChecker.Check("Привет")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Empty(t, res.Errors)
	assert.False(t, res.HasError)

}

func TestWithIncorrectValues(t *testing.T) {
	spellChecker := newChecker("")
	res, err := spellChecker.Check("Превет медведь и заиц")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Errors)
	assert.True(t, res.HasError)
	assert.True(t, res.ContainsErrorFor("Превет"), "Превет not found")
	assert.True(t, res.ContainsErrorFor("заиц"), "заиц not found")

}

func TestWithIncorrectURL(t *testing.T) {
	spellChecker := New("localhost:23124")

	res, err := spellChecker.Check("Превет")

	assert.NotNil(t, err)
	assert.NotNil(t, res)

}

func newChecker(url string) *YandexSpellChecker {
	if url != "" {
		return New(url)
	}
	return New("https://speller.yandex.net/services/spellservice.json/checkText?text=")
}

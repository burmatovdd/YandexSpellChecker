package SpellChecker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) GetName() string {
	return "MOQ"
}

func (m *MyMockedObject) Check(value string) (SpellCheckResult, error) {
	args := m.Called(value)
	fmt.Println(args)
	return args.Get(0).(SpellCheckResult), args.Error(1)
}

func TestWithCorrectValues(t *testing.T) {
	testObj := new(MyMockedObject)
	MOQResult := SpellCheckResult{
		HasError:         false,
		Errors:           []SpellCheckError{},
		SpellCheckerName: "123",
	}
	testObj.On("Check", "Привет медведь и заяц").Return(MOQResult, nil)

	spellChecker := newChecker(testObj)
	res, err := spellChecker.Check("Привет медведь и заяц")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Empty(t, res.Errors)
	assert.False(t, res.HasError)
	assert.False(t, res.ContainsErrorFor("Привет"), "Привет not found")
	assert.False(t, res.ContainsErrorFor("заяц"), "заяц not found")
	assert.Equal(t, "MOQ", res.SpellCheckerName)

}

func TestWithIncorrectValues(t *testing.T) {
	testObj := new(MyMockedObject)
	MOQResult := SpellCheckResult{
		HasError: true,
		Errors: []SpellCheckError{
			{
				Word: "Превет",
				Tips: []string{
					"Привет",
				},
			},
			{
				Word: "заиц",
				Tips: []string{
					"заяц",
				},
			},
		},
		SpellCheckerName: "123",
	}
	testObj.On("Check", "Превет медведь и заиц").Return(MOQResult, nil)

	spellChecker := newChecker(testObj)
	res, err := spellChecker.Check("Превет медведь и заиц")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Errors)
	assert.True(t, res.HasError)
	assert.True(t, res.ContainsErrorFor("Превет"), "Превет not found")
	assert.True(t, res.ContainsErrorFor("заиц"), "заиц not found")
	assert.Equal(t, "MOQ", res.SpellCheckerName)

}

func newChecker(checker SpellChecker) *SpellCheckerService {
	return New(checker)
}

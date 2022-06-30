package YandexSpellChecker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"test.com/SpellChecker"
)

type YandexSpellResponse struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	Tips []string `json:"s"`
}

type YandexSpellCheckerConfig struct {
	Url string
}

type YandexSpellChecker struct {
	Config YandexSpellCheckerConfig
}

func New(url string) *YandexSpellChecker {
	return &YandexSpellChecker{
		Config: YandexSpellCheckerConfig{Url: url},
	}
}

func (checker *YandexSpellChecker) Check(value string) (SpellChecker.SpellCheckResult, error) {
	newValue := strings.Join(strings.Split(value, " "), "+")
	response, err := http.Get(checker.Config.Url + newValue)
	if err != nil {
		return SpellChecker.SpellCheckResult{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return SpellChecker.SpellCheckResult{}, err
	}

	dict := []YandexSpellResponse{}
	if err = json.Unmarshal(body, &dict); err != nil {
		return SpellChecker.SpellCheckResult{}, err
	}
	result := SpellChecker.SpellCheckResult{
		HasError: len(dict) > 0,
		Errors:   []SpellChecker.SpellCheckError{},
	}

	for i := 0; i < len(dict); i++ {
		tips := make([]string, len(dict[i].Tips))
		copy(tips, dict[i].Tips)
		result.Errors = append(result.Errors, SpellChecker.SpellCheckError{
			Word: dict[i].Word,
			Tips: tips,
		})
	}

	return result, nil
}

func (checker *YandexSpellChecker) GetName() string {
	return "Yandex"
}

package SpellChecker

import (
	"go.uber.org/zap"
)

type SpellCheckError struct {
	Word string   `json:"word"`
	Tips []string `json:"tips"`
}

type SpellCheckResult struct {
	HasError         bool              `json:"has_error"`
	Errors           []SpellCheckError `json:"errors"`
	SpellCheckerName string            `json:"spell_checker_name"`
}

func (result *SpellCheckResult) ContainsErrorFor(value string) bool {
	for _, n := range result.Errors {
		if value == n.Word {
			return true
		}
	}
	return false
}

type SpellCheckerService struct {
	checker SpellChecker
}

type SpellChecker interface {
	Check(value string) (SpellCheckResult, error)
	GetName() string
}

func New(checker SpellChecker) *SpellCheckerService {
	return &SpellCheckerService{
		checker: checker,
	}
}

// тест который MOQ
func (service *SpellCheckerService) Check(value string) (SpellCheckResult, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	result, err := service.checker.Check(value)

	if err != nil {
		logger.Error(err.Error())
		return result, err
	}

	result.SpellCheckerName = service.checker.GetName()

	logger.Info("the information that came",
		// Structured context as strongly typed Field values.
		zap.String("data", value),
		zap.Any("result", result),
	)
	return result, nil
}

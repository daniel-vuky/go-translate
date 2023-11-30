package translate

import (
	"golang.org/x/text/language"
	"testing"
)

func TestTranslateText(t *testing.T) {
	want := []string{"Hi Daniel."}
	translationText := []string{"Salut Daniel."}
	rootLanguage, rootLanguageError := language.Parse("fr")
	targetLanguage, targetLanguageError := language.Parse("en")
	if rootLanguageError != nil || targetLanguageError != nil {
		t.Errorf("Can not Parse language")
	}
	got, translatedError := TranslateText(translationText, rootLanguage, targetLanguage)
	if translatedError != nil || len(got) == 0 {
		t.Errorf("Something wrong when translated: %v", translatedError)
	}
	if got[0].Text != want[0] {
		t.Errorf("TranslateText() = %v, want: %v", got[0].Text, want[0])
	}
}

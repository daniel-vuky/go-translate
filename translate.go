package translate

import (
	"cloud.google.com/go/translate"
	"context"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func TranslateText(textToTranslate []string, rootLanguage, targetLanguage language.Tag) ([]translate.Translation, error) {
	// TODO: get API from private source
	apiKey := "" // Replace this by Google cloud for testing
	context := context.Background()
	client, clientError := translate.NewClient(context, option.WithAPIKey(apiKey))
	if clientError != nil {
		return []translate.Translation{}, clientError
	}
	defer client.Close()
	translation, translationError := client.Translate(
		context,
		textToTranslate,
		targetLanguage,
		&translate.Options{Source: rootLanguage, Format: translate.Text},
	)
	if translationError != nil {
		return []translate.Translation{}, translationError
	}
	return translation, nil
}

package translator

type Client interface {
	Translate(sourceLang, targetLang, text string) (string, error)
}

type Translator struct {
	sourceLang string
	targetLang string
	client     Client
}

func New(sourceLang, targetLang string, client Client) *Translator {
	return &Translator{
		sourceLang: sourceLang,
		targetLang: targetLang,
		client:     client,
	}
}

func (t *Translator) Translate(text string) (string, error) {
	translated, err := t.client.Translate(t.sourceLang, t.targetLang, text)
	return translated, err
}

package validationexception

type ValidationException struct {
	error error
}

func New(err error) *ValidationException {
	return &ValidationException{error: err}
}

func (e ValidationException) ErrorsTranslation(err error) error {

	return nil
}

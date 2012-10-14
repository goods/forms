package forms

import "errors"

var (
	//Some built in validators that do what you would expect.
	NonemptyValidator ValidatorFunc = nonempty_validator
	DateValidator forms.ValidatorFunc = date_validator
	EmailValidator forms.ValidatorFunc = email_validator
)

func nonempty_validator(in string) (out string, err error) {
	out = in
	if len(in) == 0 {
		err = errors.New("Value must be present")
	}
	return
}

func date_validator(in string) (string, error) {
	_, err := time.Parse(TIME_FORMAT, in)
	if err != nil {
		return in, errors.New("Invalid date")
	}
	return in, nil
}

// vague email validator (better then nothing)
func email_validator(in string) (string, error) {
	ok, err := regexp.MatchString(`<?\S+@\S+?>?`, in)

	if err != nil {
		return in, err
	}
	if !ok {
		return in, errors.New("Invalid e-mail address")
	}
	return in, nil
}

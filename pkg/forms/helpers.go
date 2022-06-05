package forms

import "strings"

type Options []string

func (o *Options) MarshalCSV() (data string, err error) {
	return strings.Join(*o, "/\\"), nil
}

func (o *Options) UnmarshalCSV(data string) (err error) {
	*o = strings.Split(data, "/\\")
	return nil
}

package structs

import "regexp"

type URLFilter struct {
	regexp.Regexp
}

func Compile(expr string) (*URLFilter, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}

	return &URLFilter{*re}, nil
}

func (r *URLFilter) UnmarshalText(expr []byte) error {
	rr, err := Compile(string(expr))
	if err != nil {
		return err
	}

	*r = *rr
	return nil
}

func (r *URLFilter) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func GetRegex(f []URLFilter) []*regexp.Regexp {
	regexArr := []*regexp.Regexp{}
	for _, filter := range f {
		regexArr = append(regexArr, &filter.Regexp)
	}

	return regexArr
}

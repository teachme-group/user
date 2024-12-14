package validate

import "regexp"

var re = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func Email(email string) bool {
	return re.MatchString(email)
}

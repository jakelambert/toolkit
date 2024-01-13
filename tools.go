package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWYZ0123456789_+"

// Tools is the type used to instantiate this module.  Any variable of this type will have access
// to all the methods with the receiver *Tools
type Tools struct {
}

// RandomString returns a string of random characters of length n, using randomStringSource
// as the source for the string
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	y := uint64(len(r))

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		s[i] = r[x%y]
	}

	return string(s)
}

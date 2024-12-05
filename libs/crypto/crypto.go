package crypto

import "golang.org/x/crypto/bcrypt"

type IPasswordEncoder interface {
	Hash(password string) (result string, err error)
	Compare(input, encrypted string) (err error)
}

type passwordEncoder struct{}

func NewPasswordEncoder() *passwordEncoder {
	return &passwordEncoder{}
}

func (p *passwordEncoder) Hash(password string) (result string, err error) {
	cost := bcrypt.DefaultCost
	byteData, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return
	}
	result = string(byteData)
	return
}

func (p *passwordEncoder) Compare(input, encrypted string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(input))
	return
}

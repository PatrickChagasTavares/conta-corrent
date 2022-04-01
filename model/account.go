package model

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	cpfSize = 11
)

var (
	errCPFRequired    = NewError(http.StatusBadRequest, "o cpf é obrigatorio.", nil)
	errCPFSizeInvalid = NewError(http.StatusBadRequest, "o cpf deve ter no 11 caracteres.", nil)
	errCPFInvalid     = NewError(http.StatusBadRequest, "o cpf é inválido.", nil)
	errNameRequired   = NewError(http.StatusBadRequest, "o nome é obrigatorio.", nil)
	errSecretRequired = NewError(http.StatusBadRequest, "a senha é obrigatoria.", nil)

	cpfInvalidKnown = map[string]bool{
		"00000000000": true, "11111111111": true,
		"22222222222": true, "33333333333": true,
		"44444444444": true, "55555555555": true,
		"66666666666": true, "77777777777": true,
		"88888888888": true, "99999999999": true,
	}
	cpfFirstDigitTable  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

type Account struct {
	ID         int       `json:"id,omitempty" db:"id"`
	Name       string    `json:"name" db:"name"`
	CPF        string    `json:"cpf" db:"cpf"`
	SecretHash string    `json:"-" db:"secret_hash"`
	SecretSalt string    `json:"-" db:"secret_salt"`
	Secret     string    `json:"secret,omitempty" db:"-"`
	Balance    float64   `json:"balance,omitempty" db:"balance"`
	CreatedAt  time.Time `json:"created_at,omitempty" db:"created_at"`
}

// validate if account is valid
func (a *Account) Validate() error {
	if a.Name == "" {
		return errNameRequired
	}

	if a.Secret == "" {
		return errSecretRequired
	}

	if a.CPF == "" {
		return errCPFRequired
	}

	if err := a.cpfIsValid(); err != nil {
		return err
	}

	return nil
}

// removeEspecChar remove special characters
func (a *Account) removeSpecialCharacterCPF() {
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")

	a.CPF = regex.ReplaceAllString(a.CPF, "")
}

func (a *Account) cpfIsValid() error {
	a.removeSpecialCharacterCPF()
	if !cpfSizeIsValid(a.CPF) {
		return errCPFSizeInvalid
	}
	if invalidCPFIsKnown(a.CPF) {
		return errCPFInvalid
	}
	if !cpfDigitsValid(a.CPF) {
		return errCPFInvalid
	}
	return nil
}

// cpfSizeIsValid valid if cpf size is valid
func cpfSizeIsValid(cpf string) bool {
	if len(cpf) != cpfSize {
		return false
	}
	return true
}

// InvalidCpfIsKnown valid if cpf is known
func invalidCPFIsKnown(cpf string) bool {

	if cpfInvalidKnown[cpf] {
		return true
	}
	return false
}

//CpfDigitsValid check if the cpf digits are valid
func cpfDigitsValid(cpf string) bool {
	firstPart := cpf[0:9]
	sum := sumDigit(firstPart, cpfFirstDigitTable)

	r1 := sum % cpfSize
	d1 := 0

	if r1 >= 2 {
		d1 = cpfSize - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)

	dsum := sumDigit(secondPart, cpfSecondDigitTable)

	r2 := dsum % cpfSize
	d2 := 0

	if r2 >= 2 {
		d2 = cpfSize - r2
	}

	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)
	return finalPart == cpf
}

// sumDigit sum the digit
func sumDigit(s string, table []int) int {

	if len(s) != len(table) {
		return 0
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}

package api

type MFAService interface {
	Secret() string
	Generate(account, secret string) (string, error)
	Authenticate(token, secret string) (bool, error)
	// Email(address, secret string) error
}
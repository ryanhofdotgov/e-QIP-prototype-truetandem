package mfa

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"net/url"
	"strconv"
	"strings"
	"text/template"

	"github.com/18F/e-QIP-prototype/api"
	"github.com/dgryski/dgoogauth"
	qr "github.com/skip2/go-qrcode"
)

const (
	auth string = "totp"
)

var (
	templateEmail = template.Must(template.New("email").Parse(`# Passcode\n\n{{ . }}`))
)

type MFAService struct {
	Log api.LogService
	Env api.Settings
}

// Secret creates a random secret and then base32 encodes it.
func (service MFAService) Secret() string {
	secret := make([]byte, 6)
	_, err := rand.Read(secret)
	if err != nil {
		return ""
	}

	secret32 := base32.StdEncoding.EncodeToString(secret)
	return secret32
}

// Generate will create a QR code in PNG format which will then
// be base64 encoded so it can traverse the wire to the front end.
func (service MFAService) Generate(account, secret string) (string, error) {
	u, err := url.Parse("otpauth://" + auth)
	if err != nil {
		return "", err
	}

	// Google authenticator (on iOS) does not react well to the base32 right padding
	// of "=" runes. Removing them seems to have no adverse affect for those authenticators
	// currently already working.
	secret = strings.TrimRight(secret, "=")

	u.Path += "/eqip:" + account
	params := &url.Values{}
	params.Add("secret", secret)
	params.Add("issuer", "eqip")
	u.RawQuery = params.Encode()

	png, err := qr.Encode(u.String(), qr.Medium, 256)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(png), nil
}

// Authenticate validates the initial token generated when configuring two-factor
// authentication for the first time.
func (service MFAService) Authenticate(token, secret string) (ok bool, err error) {
	// Get the adjustable window size
	size := 3
	if service.Env.Has(api.WINDOW_SIZE) {
		i, e := strconv.Atoi(service.Env.String(api.WINDOW_SIZE))
		if e == nil {
			size = i
			service.Log.Debug("Setting MFA window size", api.LogFields{"window": i})
		}
	}

	// Create the OTP configuration to use for authenticating the token
	otpc := &dgoogauth.OTPConfig{
		Secret:      secret,
		WindowSize:  size,
		HotpCounter: 0,
		UTC:         true,
	}
	return otpc.Authenticate(token)
}

// // Email delivers code to the specified address.
// func (service MFAService) Email(address, secret string) error {
// 	// Get a valid token for two-factor authentication
// 	code := dgoogauth.ComputeCode(secret, 0)
// 	if code == -1 {
// 		return errors.New("Failed to generate code")
// 	}

// 	// Pull the API key for the mail service
// 	key := service.Env.String("EQIP_SMTP_API_KEY")
// 	if key == "" {
// 		return errors.New("Could not retrieve API key")
// 	}

// 	// Form the mail service
// 	client := mandrill.ClientWithKey(key)
// 	message := &mandrill.Message{
// 		FromEmail: "noreply@mail.gov",
// 		FromName:  "E-QIP",
// 		Subject:   "E-QIP Passcode",
// 	}
// 	message.AddRecipient(address, address, "to")

// 	// The template is stored in markdown format to easily
// 	// transform something human readable to HTML as well
// 	var plain bytes.Buffer
// 	err := templateEmail.Execute(&plain, code)
// 	if err != nil {
// 		return err
// 	}
// 	message.Text = plain.String()
// 	unsafe := blackfriday.MarkdownCommon(plain.Bytes())
// 	message.HTML = string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

// 	// Send the message and return any errors from the service
// 	_, err = client.MessagesSend(message)
// 	return err
// }
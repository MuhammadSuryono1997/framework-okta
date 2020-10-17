package otp

import (
	"encoding/base32"
	"flag"
	"fmt"
	"strings"

	"github.com/MuhammadSuryono1997/framework-okta/utils"
	"github.com/hgfischer/go-otp"
)

var (
	secret   = flag.String("secret", "OTPOktaPOS", "Secret key")
	isBase32 = flag.Bool("base32", true, "If true, the secret is interpreted as a Base32 string")
	length   = flag.Uint("length", 4, "OTP length")
	period   = flag.Uint("period", otp.DefaultPeriod, "Period in seconds")
	counter  = flag.Uint64("counter", 0, "Counter")
)

var totp = &otp.TOTP{
	Secret:         *secret,
	Length:         uint8(*length),
	Period:         uint8(*period),
	IsBase32Secret: true,
}

func GenerateOTP(expired uint8) string {
	flag.Parse()
	key := *secret

	if !*isBase32 {
		key = base32.StdEncoding.EncodeToString([]byte(*secret))
	}

	key = strings.ToUpper(key)
	if !isGoogleAuthenticatorCompatible(key) {
		fmt.Println(string(utils.ColorYellow()), "WARN: Google Authenticator requires 16 chars base32 secret, without padding")
	}

	fmt.Println(string(utils.ColorCyan()), "Secret Base32 Encoded Key: ", key)

	totp.Secret = key
	totp.Period = expired

	return totp.Now().Get()
}

func ValidateOTP(otp string) bool {
	return totp.Now().Verify(otp)
}

func isGoogleAuthenticatorCompatible(base32Secret string) bool {
	cleaned := strings.Replace(base32Secret, "=", "", -1)
	cleaned = strings.Replace(cleaned, " ", "", -1)
	return len(cleaned) == 16
}

func RequestOTP(nohp string) {

}

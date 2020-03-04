package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"server/pkg/setting"
	"time"
)

type reCAPTCHAResponse struct {
	Success        bool      `json:"success"`
	ChallengeTS    time.Time `json:"challenge_ts"`
	Hostname       string    `json:"hostname,omitempty"`
	ApkPackageName string    `json:"apk_package_name,omitempty"`
	Action         string    `json:"action,omitempty"`
	Score          float32   `json:"score,omitempty"`
	ErrorCodes     []string  `json:"error-codes,omitempty"`
}

func ReCaptcha(token string) error {
	// send request to google
	secret := setting.ReCaptchaSecret
	response, err := http.Get("https://recaptcha.net/recaptcha/api/siteverify?secret=" + secret + "&response=" + token)
	if err != nil {
		return err
	}

	// get response
	defer response.Body.Close()
	var data reCAPTCHAResponse
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return err
	}

	// validate response
	if !data.Success {
		return errors.New("reCAPTCHA verification failed. ")
	}
	if data.Score < 0.5 {
		return errors.New(data.Action + " Action: Not a Human. ")
	}

	return nil
}

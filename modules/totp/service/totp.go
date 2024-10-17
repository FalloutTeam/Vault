package service

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/boombuler/barcode/qr"
	"github.com/pquerna/otp/totp"
	"github.com/sirupsen/logrus"
	"image/png"
	"vault/modules/totp/models"
	"vault/modules/totp/repository"
)

type TotpService struct {
	repo *repository.Repository
}

func (t *TotpService) CreateTotp(name string, totpParams models.Totp) (string, string, error) {
	if totpParams.Generate != true {
		logrus.Errorf("modules/totp/service/CreateTotp error: %s", "wrong param value")
		return "", "", errors.New("wrong param value")
	}
	if totpParams.Digits == 0 {
		totpParams.Digits = 6
	}
	if totpParams.Algorithm == "" {
		totpParams.Algorithm = "SHA1"
	}
	url, err := t.generateTotpKey("Google", totpParams.AccountName)
	if err != nil {
		logrus.Errorf("modules/totp/service/CreateTotp error: %s", err.Error())
		return "", "", err
	}

	err = t.repo.Totp.CreateTotp(name, totpParams)

	base64QR, err := t.generateBase64QR(url)
	if err != nil {
		logrus.Errorf("modules/totp/service/CreateTotp error: %s", err.Error())
		return "", "", err
	}

	return base64QR, url, nil
}

func (t *TotpService) generateTotpKey(issuer string, accountName string) (string, error) {
	secret := totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
	}
	key, err := totp.Generate(secret)
	if err != nil {
		logrus.Errorf("modules/totp/service/generateTotpKey error: %s", err.Error())
		return "", err
	}

	return key.URL(), nil
}

func (t *TotpService) generateBase64QR(totpURL string) (string, error) {
	qrCode, err := qr.Encode(totpURL, qr.M, qr.Auto)
	if err != nil {
		logrus.Errorf("modules/totp/service/generateBase64QR error: %s", err.Error())
		return "", err
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, qrCode)
	if err != nil {
		logrus.Errorf("modules/totp/service/generateBase64QR error: %s", err.Error())
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
func NewTotpService(repo *repository.Repository) *TotpService {
	return &TotpService{repo: repo}
}

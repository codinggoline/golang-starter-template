package session

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"golang_starter_template/pkg/jobs/entity"
	"golang_starter_template/pkg/utils"
	"strings"
	"time"
)

var SecretKey = []byte(utils.GetEnv("JWT_SECRET"))

func GenerateToken(userID int, email string, roles []entity.Role) (string, error) {
	// Header
	header := `{ "alg": "HS256", "typ": "JWT" }`

	// Payload
	expiredAt := time.Now().Add(time.Hour * 24).Unix()
	payload := struct {
		UserID int           `json:"user_id"`
		Email  string        `json:"email"`
		Roles  []entity.Role `json:"roles"`
		Exp    int64         `json:"exp"`
	}{
		UserID: userID,
		Email:  email,
		Roles:  roles,
		Exp:    expiredAt,
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Encoder Header and Payload
	encoderHeader := base64.RawURLEncoding.EncodeToString([]byte(header))
	encodePayload := base64.RawURLEncoding.EncodeToString(payloadJSON)

	// Signature
	signatureInput := encoderHeader + "." + encodePayload
	h := hmac.New(sha256.New, SecretKey)
	h.Write([]byte(signatureInput))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("%s.%s.%s", encoderHeader, encodePayload, signature), nil
}

func ValidateToken(token string) (int, []entity.Role, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return 0, nil, errors.New("invalid token format")
	}

	// Check Signature
	signatureInput := parts[0] + "." + parts[1]
	h := hmac.New(sha256.New, SecretKey)
	h.Write([]byte(signatureInput))
	expactedSignature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	if parts[2] != expactedSignature {
		return 0, nil, errors.New("invalid token signature")
	}

	// Decode Payload
	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return 0, nil, err
	}

	var payload struct {
		UserID int           `json:"user_id"`
		Email  string        `json:"email"`
		Roles  []entity.Role `json:"roles"`
		Exp    int64         `json:"exp"`
	}
	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return 0, nil, err
	}

	// Check Expired
	if time.Now().Unix() > payload.Exp {
		return 0, nil, errors.New("token expired")
	}

	return payload.UserID, payload.Roles, nil
}

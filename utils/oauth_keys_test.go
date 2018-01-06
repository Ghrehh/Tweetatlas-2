package utils

import "testing"

func TestOauthKeysParse(t *testing.T) {
	keyData := `{
			"consumer_key":"consumer key",
			"consumer_secret":"consumer secret",
			"access_token":"access token",
			"access_secret":"access secret"
	}`

	keys := parseOauthKeys([]byte(keyData))

	if keys.ConsumerKey != "consumer key" {
		t.Error("expected the correct value")
	}

	if keys.ConsumerSecret != "consumer secret" {
		t.Error("expected the correct value")
	}

	if keys.AccessToken != "access token" {
		t.Error("expected the correct value")
	}

	if keys.AccessSecret != "access secret" {
		t.Error("expected the correct value")
	}
}

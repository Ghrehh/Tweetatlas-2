package utils

import (
	"testing"
)

func TestStreamFilterParse(t *testing.T) {
	paramsData := `{
			"filter": ["foo", "bar"]
	}`

	streamParams := parseStreamParams([]byte(paramsData))
	filter := streamParams.Filter
	expectedFilter := []string{"foo", "bar"}

	if filter[0] != expectedFilter[0] {
		t.Error("expected the correct value")
	}

	if filter[1] != expectedFilter[1] {
		t.Error("expected the correct value")
	}
}

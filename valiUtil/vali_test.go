package valiUtil

import "testing"

func TestValidateData(t *testing.T) {
	t.Parallel()

	rules := []Rules{
		{Mkey: "user_id", Value: "缺少 user_id"},
		{Mkey: "room_id", Value: "缺少 room_id"},
	}

	tests := []struct {
		name    string
		data    map[string]any
		wantOK  bool
		wantMsg string
	}{
		{name: "all present", data: map[string]any{"user_id": 1, "room_id": 2}, wantOK: true},
		{name: "missing key", data: map[string]any{"user_id": 1}, wantMsg: "缺少 room_id"},
		{name: "nil value", data: map[string]any{"user_id": nil, "room_id": 2}, wantMsg: "缺少 user_id"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok, message := ValidateData(tc.data, rules)
			if ok != tc.wantOK || message != tc.wantMsg {
				t.Fatalf("ValidateData() = (%t, %q), want (%t, %q)", ok, message, tc.wantOK, tc.wantMsg)
			}
		})
	}
}

func TestValidatePhone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		value string
		want  bool
	}{
		{value: "13800138000", want: true},
		{value: "010-12345678", want: true},
		{value: "01012345678", want: true},
		{value: "11111111111", want: false},
		{value: "12345", want: false},
	}

	for _, tc := range tests {
		if got := ValidatePhone(tc.value); got != tc.want {
			t.Errorf("ValidatePhone(%q) = %t, want %t", tc.value, got, tc.want)
		}
	}
}

func TestValidateIDCard(t *testing.T) {
	t.Parallel()

	if ok, err := ValidateIdCard("110102197809193026"); !ok || err != nil {
		t.Fatalf("ValidateIdCard(valid) = (%t, %v), want (true, nil)", ok, err)
	}
	if ok, err := ValidateIdCard("123456789012345678"); ok || err == nil {
		t.Fatalf("ValidateIdCard(invalid) = (%t, %v), want (false, error)", ok, err)
	}
	if ok, err := ValidateIdCard("123"); ok || err == nil {
		t.Fatalf("ValidateIdCard(short) = (%t, %v), want (false, error)", ok, err)
	}
}

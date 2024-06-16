package mmphone

import (
	"testing"
)

func TestSanitizePhoneNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" 09 777 123456 ", "09777123456"},
		{"+95 9 777 123456", "09777123456"},
		{"9595 777 123456", "095777123456"},
		{"၀၉ ၇၇၇ ၁၂၃၄၅၆", "09777123456"},
	}

	mm := NewMyanmarPhone()

	for _, test := range tests {
		result := mm.SanitizePhoneNumber(test.input)
		if result != test.expected {
			t.Errorf("SanitizePhoneNumber(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestIsValidMyanmarPhone(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"09777123456", true},
		{"+959777123456", true},
		{"၀၉၇၇၇၁၂၃၄၅၆", true},
		{"123456", false},
		{"+959812345", false},
	}

	mm := NewMyanmarPhone()

	for _, test := range tests {
		result := mm.IsValidMyanmarPhone(test.input)
		if result != test.expected {
			t.Errorf("IsValidMyanmarPhone(%s) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestGetTelecomName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"09977123456", OOREDOO},
		{"+959775541794", ATOM},
		{"၀၉၅၁၂၃၄၅၆", MPT},
		{"09156123456", UNKNOWN},
		{"09678123456", MYTEL},
	}

	mm := NewMyanmarPhone()

	for _, test := range tests {
		result := mm.GetTelecomName(test.input)
		if result != test.expected {
			t.Errorf("GetTelecomName(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestGetPhoneNetworkType(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"09777123456", GSM_TYPE},
		{"+959781234567", GSM_TYPE},
		{"၀၉၅၅၁၂၃၄၆", WCDMA_TYPE},
		{"0949123456", CDMA_450_TYPE},
		{"0937123456", CDMA_800_TYPE},
	}

	mm := NewMyanmarPhone()

	for _, test := range tests {
		result := mm.GetPhoneNetworkType(test.input)
		if result != test.expected {
			t.Errorf("GetPhoneNetworkType(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestConvertBurmeseNumerals(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"၀၉၇၇၇၁၂၃၄၅၆", "09777123456"},
		{"၉၅၉၅၇၇၇၁၂၃၄၅၆", "9595777123456"},
	}

	mm := NewMyanmarPhone()

	for _, test := range tests {
		result := mm.convertBurmeseNumerals(test.input)
		if result != test.expected {
			t.Errorf("convertBurmeseNumerals(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

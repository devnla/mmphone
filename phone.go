// mmphone - pohne.go
// MIT License
//
// Copyright (c) 2024 Nay Lin Aung (https://github.com/devnla)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package mmphone

import (
	"regexp"
	"strings"
)

const (
	OOREDOO = "Ooredoo"
	ATOM    = "ATOM"
	MPT     = "MPT"
	MYTEL   = "MyTel"
	UNKNOWN = "Unknown"

	GSM_TYPE      = "GSM"
	WCDMA_TYPE    = "WCDMA"
	CDMA_450_TYPE = "CDMA 450 MHz"
	CDMA_800_TYPE = "CDMA 800 MHz"
)

var (
	OoredooRegex      = regexp.MustCompile(`^(09|\+?959)9(7|6)\d{7}$`)
	AtomRegex         = regexp.MustCompile(`^(09|\+?959)7(9|8|7)\d{7}$`)
	MptRegex          = regexp.MustCompile(`^(09|\+?959)(5\d{6}|4\d{7,8}|2\d{6,8}|3\d{7,8}|6\d{6}|8\d{6}|7\d{7}|9(0|1|9)\d{5,6})$`)
	MyTelRegex        = regexp.MustCompile(`^(09|\+?959)6([5-9])\d{7}$`)
	MyanmarPhoneRegex = regexp.MustCompile(`^(09|\+?950?9|\+?95950?9)\d{7,9}$`)
)

type MyanmarPhone struct{}

func NewMyanmarPhone() *MyanmarPhone {
	return &MyanmarPhone{}
}

func (mm *MyanmarPhone) IsValidMyanmarPhone(phoneNumber string) bool {
	phoneNumber = mm.SanitizePhoneNumber(phoneNumber)
	return MyanmarPhoneRegex.MatchString(phoneNumber)
}

func (mm *MyanmarPhone) SanitizePhoneNumber(phoneNumber string) string {
	phoneNumber = mm.convertBurmeseNumerals(phoneNumber)
	phoneNumber = strings.TrimSpace(phoneNumber)
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "+", "")

	if matched, _ := regexp.MatchString(`^\+?950?9\d+$`, phoneNumber); matched {
		phoneNumber = strings.Replace(phoneNumber, "959", "09", 1)
		phoneNumber = strings.Replace(phoneNumber, "9509", "09", 1)
	}

	return phoneNumber
}

func (mm *MyanmarPhone) convertBurmeseNumerals(phoneNumber string) string {
	var converted strings.Builder
	for _, r := range phoneNumber {
		if r >= '၀' && r <= '၉' {
			converted.WriteRune('0' + (r - '၀'))
		} else {
			converted.WriteRune(r)
		}
	}
	return converted.String()
}

func (mm *MyanmarPhone) GetTelecomName(phoneNumber string) string {
	if !mm.IsValidMyanmarPhone(phoneNumber) {
		return UNKNOWN
	}

	phoneNumber = mm.SanitizePhoneNumber(phoneNumber)

	switch {
	case OoredooRegex.MatchString(phoneNumber):
		return OOREDOO
	case AtomRegex.MatchString(phoneNumber):
		return ATOM
	case MptRegex.MatchString(phoneNumber):
		return MPT
	case MyTelRegex.MatchString(phoneNumber):
		return MYTEL
	default:
		return UNKNOWN
	}
}

func (mm *MyanmarPhone) GetPhoneNetworkType(phoneNumber string) string {
	if !mm.IsValidMyanmarPhone(phoneNumber) {
		return UNKNOWN
	}

	phoneNumber = mm.SanitizePhoneNumber(phoneNumber)

	wcdmaRe := regexp.MustCompile(`^(09|\+?959)(55\d{5}|25[2-4]\d{6}|26\d{7}|4(4|5|6)\d{7})$`)
	cdma450Re := regexp.MustCompile(`^(09|\+?959)(8\d{6}|6\d{6}|49\d{6})$`)
	cdma800Re := regexp.MustCompile(`^(09|\+?959)(3\d{7}|73\d{6}|91\d{6})$`)

	switch {
	case OoredooRegex.MatchString(phoneNumber), AtomRegex.MatchString(phoneNumber), MyTelRegex.MatchString(phoneNumber):
		return GSM_TYPE
	case MptRegex.MatchString(phoneNumber):
		switch {
		case wcdmaRe.MatchString(phoneNumber):
			return WCDMA_TYPE
		case cdma450Re.MatchString(phoneNumber):
			return CDMA_450_TYPE
		case cdma800Re.MatchString(phoneNumber):
			return CDMA_800_TYPE
		default:
			return GSM_TYPE
		}
	default:
		return UNKNOWN
	}
}

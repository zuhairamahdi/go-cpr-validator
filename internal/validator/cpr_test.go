package validator

import "testing"

func TestValidateCPR_Valid(t *testing.T) {
	valid := "901208302"
	ok, err := ValidateCPR(valid)
	if !ok {
		t.Fatalf("expected valid CPR to pass, got error: %v", err)
	}
}

func TestValidateCPR_RejectsMissingValue(t *testing.T) {
	_, err := ValidateCPR("")
	if err == nil {
		t.Fatal("expected empty input to fail")
	}
}

func TestValidateCPR_RejectsWrongLength(t *testing.T) {
	_, err := ValidateCPR("12345")
	if err == nil {
		t.Fatal("expected wrong-length CPR to fail")
	}
}

func TestValidateCPR_RejectsNonNumericChars(t *testing.T) {
	_, err := ValidateCPR("2001A0103")
	if err == nil {
		t.Fatal("expected non-numeric CPR to fail")
	}
}

func TestValidateCPR_RejectsInvalidDate(t *testing.T) {
	_, err := ValidateCPR("20010230A")
	if err == nil {
		t.Fatal("expected invalid date to fail")
	}
}

func TestValidateCPR_RejectsInvalidChecksum(t *testing.T) {
	_, err := ValidateCPR("200101014")
	if err == nil {
		t.Fatal("expected checksum mismatch to fail")
	}
}

package util

import (
	"testing"
)

func TestStringToDate(t *testing.T) {
	converterTime, err := StringToTime("2025-02-25T18:30:45", LayoutDate)
	if err != nil {
		t.Fail()
		t.Errorf("expected nil, but got error [%s]", err.Error())
	}

	if converterTime.Year() != 2025 {
		t.Errorf("expected 2025, but got %v", converterTime.Year())
	}

	if converterTime.Month().String() != "February" {
		t.Errorf("expected February, but got %v", converterTime.Month().String())
	}

	if converterTime.Day() != 25 {
		t.Errorf("expected 25, but got %v", converterTime.Day())
	}

	if converterTime.Hour() != 18 {
		t.Errorf("expected 18, but got %v", converterTime.Hour())
	}

	if converterTime.Minute() != 30 {
		t.Errorf("expected 30, but got %v", converterTime.Minute())
	}

	if converterTime.Second() != 45 {
		t.Errorf("expected 45, but got %v", converterTime.Second())
	}
}

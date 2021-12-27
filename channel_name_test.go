package main

import (
	"fmt"
	"testing"
	"time"
)

func TestYearPercentage(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		t.Errorf("Unable to load location: %q", err)
	}

	assertPercentage := func(t testing.TB, got, want float64) {
		t.Helper()
		if fmt.Sprintf("%.4f", got) != fmt.Sprintf("%.4f", want) {
			t.Errorf("Bad percentage, got %.4f, want %.4f", got, want)
		}
	}

	t.Run("not special percentage", func(t *testing.T) {
		got := YearPercentage(time.Date(2021, time.December, 27, 1, 20, 0, 0, loc))
		want := .986453
		assertPercentage(t, got, want)
	})

	t.Run("first instant of the year", func(t *testing.T) {
		got := YearPercentage(time.Date(2021, time.January, 1, 0, 0, 0, 0, loc))
		want := .0
		assertPercentage(t, got, want)
	})

	t.Run("non-leap year", func(t *testing.T) {
		got := YearPercentage(time.Date(2021, time.February, 4, 1, 20, 0, 0, loc))
		want := .093300
		assertPercentage(t, got, want)
	})

	t.Run("leap year", func(t *testing.T) {
		got := YearPercentage(time.Date(2016, time.February, 4, 1, 20, 0, 0, loc))
		want := .093050
		assertPercentage(t, got, want)
	})
}

func TestGenerateChannelName(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		t.Errorf("Unable to load location: %q", err)
	}

	got := GenerateChannelName(
		"covid-{begin}➔{end}",
		time.Date(2019, time.November, 16, 0, 0, 0, 0, loc),
		time.Date(2021, time.December, 26, 0, 0, 0, 0, loc),
	)
	expected := "covid-19.87➔21.98"

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}

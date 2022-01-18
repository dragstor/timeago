package timeago

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago/utils"
)

func TestGetOption(test *testing.T) {
	cases := []struct {
		name            string
		date            string
		optionMustBe    string
		hasOptionMustBe bool
	}{
		{"case has online option", "2017-02-01 00:00:00|online", "online", true},
		{"case has random option", "2017-02-01 00:00:00|random", "random", true},
		{"case has online option", "2017-02-01 00:00:00|korotchaeva", "korotchaeva", true},
		{"case don't have option", "2017-02-01 00:00:00", "", false},
	}

	for _, tc := range cases {
		test.Run(tc.name, func(t *testing.T) {
			option, hasOption := getOption(&tc.date)

			if hasOption != tc.hasOptionMustBe || option != tc.optionMustBe {
				t.Errorf("Result of getOption func must return `online` string and `true`, but `%s` string and `%v` returned returned", tc.optionMustBe, tc.hasOptionMustBe)
			}
		})
	}
}

func TestGetWords(t *testing.T) {
	cases := []struct {
		timeKind string
		num      int
		result   string
		lang     string
	}{
		// english
		{"days", 11, "11 days ago", "en"},
		{"days", 21, "21 days ago", "en"},
		{"seconds", 30, "30 seconds ago", "en"},
		{"seconds", 31, "31 seconds ago", "en"},
		{"hours", 10, "10 hours ago", "en"},
		{"years", 2, "2 years ago", "en"},
		// russian
		{"hours", 5, "5 часов назад", "ru"},
		{"days", 11, "11 дней назад", "ru"},
		{"years", 21, "21 год назад", "ru"},
		{"minutes", 59, "59 минут назад", "ru"},
	}

	for _, tc := range cases {
		t.Run(tc.result, func(test *testing.T) {
			Set("language", tc.lang)
			Set("location", "Europe/Kiev")

			if res := getWords(tc.timeKind, tc.num); res != tc.result {
				test.Errorf("Result must be `%s` but got `%s` instead", tc.result, res)
			}
		})
	}
}

func TestGetLastNumber(t *testing.T) {
	cases := []struct {
		number int
		result int
		name   string
	}{
		{0, 0, "must return 0"},
		{1, 1, "must return 1"},
		{9, 9, "must return 9"},
		{20, 0, "must return 0"},
		{253, 3, "must return 3"},
		{23423252, 2, "must return 2"},
		{223424342325, 5, "must return 5"},
		{23423521562348, 8, "must return 8"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			if res := getLastNumber(tc.number); res != tc.result {
				test.Errorf("Result must be %d, but got %d instead", tc.result, res)
			}
		})
	}
}

func TestTakeFunctionCanExceptTimestamp(t *testing.T) {
	cases := []struct {
		timestamp int
		result    string
	}{
		{GetTimestampOfPastDate(time.Minute), "1 minute ago"},
		{GetTimestampOfPastDate(time.Minute * 5), "5 minutes ago"},
		{GetTimestampOfPastDate(time.Hour), "1 hour ago"},
		{GetTimestampOfPastDate(time.Hour * 3), "3 hours ago"},
		{GetTimestampOfPastDate(time.Hour * 24), "1 day ago"},
		{GetTimestampOfPastDate(time.Hour * 48), "2 days ago"},
	}

	Set("language", "en")

	for _, tc := range cases {
		t.Run(tc.result, func(test *testing.T) {
			if res := Take(tc.timestamp); res != tc.result {
				test.Errorf("Result must be %v, but got %v instead", tc.result, res)
			}
		})
	}

}

func TestTakeFunctionCanExceptTimePackage(t *testing.T) {
	cases := []struct {
		time   time.Time
		result string
	}{
		{time.Now().Add(-time.Minute), "1 minute ago"},
		{time.Now().Add(-time.Minute * 2), "2 minutes ago"},
		{time.Now().Add(-time.Minute * 3), "3 minutes ago"},
		{time.Now().Add(-time.Minute * 4), "4 minutes ago"},
		{time.Now().Add(-time.Minute * 5), "5 minutes ago"},
		{time.Now().Add(-time.Minute * 6), "6 minutes ago"},
		{time.Now().Add(-time.Hour * 7), "7 hours ago"},
		{time.Now().Add(-time.Hour * 8), "8 hours ago"},
		{time.Now().Add(-time.Hour * 9), "9 hours ago"},
		{time.Now().Add(-time.Hour * 10), "10 hours ago"},
		{time.Now().Add(-time.Hour * 11), "11 hours ago"},
	}

	Set("language", "en")

	for _, tc := range cases {
		t.Run(tc.result, func(test *testing.T) {
			if res := Take(tc.time); res != tc.result {
				test.Errorf("Result must be %v, but got %v instead", tc.result, res)
			}
		})
	}

}

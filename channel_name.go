package main

import (
	"fmt"
	"strings"
	"time"
)

// YearPercentage computes the percentage, in [0.0; 1.0[, where the given date is
// in its current year.
func YearPercentage(date time.Time) float64 {
	loc := date.Location()

	yearStart := time.Date(date.Year(), time.January, 1, 0, 0, 0, 0, loc)
	yearEnd := time.Date(date.Year(), time.December, 31, 23, 59, 59, 999, loc)

	duration := date.Sub(yearStart)
	yearDuration := yearEnd.Sub(yearStart)

	return float64(duration) / float64(yearDuration)
}

const standardComma = "."
const discordifiedComma = "،"

// DiscordifyFloat converts a float64 to string so that the resulting
// string can be used in a channel name. Discord does not allow comas in
// channel names, so we use an alternative unicode character to represent
// them. We must, too, specify the precision wanted (number of digits after
// the so-called comma), because channel names have limitations.
func DiscordifyFloat(number float64, precision int) string {
	return strings.Replace(
		fmt.Sprintf(fmt.Sprintf("%%.%df", precision), number), standardComma, discordifiedComma, -1,
	)
}

// GenerateChannelName computes a channel name displaying years in decimal form.
// As example, 19.87-to-21.98 where 19/21 are the years, and the decimal parts, where
// we are in the year.
// The first arg is the pattern we should output, where {begin} is replaced by the computed
// decimal year corresponding to the begin argument, and the same for {end}.
func GenerateChannelName(pattern string, begin, end time.Time) string {
	beginPercentage := DiscordifyFloat(float64(begin.Year()-2000)+YearPercentage(begin), 2)
	endPercentage := DiscordifyFloat(float64(end.Year()-2000)+YearPercentage(end), 2)

	r := strings.NewReplacer("{begin}", beginPercentage, "{end}", endPercentage)
	return r.Replace(pattern)
}

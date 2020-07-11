// Package tcnksm is a Go package for showing tcnksm profile.
package tcnksm

import "time"

var now = time.Now

// Name returns tcnksm's real name which most of people dont' know.
func Name() string {
	return "Taichi Nakashima"
}

// Birthday returns tcnksm's birthday.
func Birthday() time.Time {
	return time.Date(1988, time.November, 24, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
}

// Age returns tcnksm's age.
func Age() int {
	return now().Year() - Birthday().Year()
}

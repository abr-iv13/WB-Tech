package time

import "testing"

func TestGetTime(t *testing.T) {
	GetTime("")
	GetTime("0.beevik-ntp.pool.ntp.org")
}

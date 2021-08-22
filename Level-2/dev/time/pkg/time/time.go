package time

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

const baseHost = "0.beevik-ntp.pool.ntp.org"

//GetTime show local and exact time
func GetTime(host string) {
	if host == "" {
		host = baseHost
	}

	response, err := ntp.Query(host)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	localTime := time.Now()
	exactTime := time.Now().Add(response.ClockOffset)

	fmt.Printf("Текущее время: %s\nТочное время: %s\n", localTime, exactTime)
}

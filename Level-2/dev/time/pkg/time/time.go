package time

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

const baseHost = "0.beevik-ntp.pool.ntp.org"

//ShowTime show local and exact time
func ShowTime(host string) {
	if host == "" {
		host = baseHost
	}
	//Получение текущего времени и дополнительный метаднных
	response, err := ntp.Query(host)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	//Локальные время
	localTime := time.Now()
	//Точное время с хоста
	exactTime := time.Now().Add(response.ClockOffset)

	//Вывод результатов
	fmt.Printf("Текущее время: %s\nТочное время: %s\n", localTime, exactTime)
}

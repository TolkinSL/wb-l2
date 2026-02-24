package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const ntpServer = "pool.ntp.org"

func getTime() (time.Time, error) {
	resp, err := ntp.Query(ntpServer)
	if err != nil {
		return time.Time{}, fmt.Errorf("getTime ntp.Query: %w", err)
	}
	return time.Now().Add(resp.ClockOffset), nil
}

func main() {
	t, err := getTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(t.Format(time.RFC3339))
}

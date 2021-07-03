package gigasecond

import (
	"log"
	"time"
)

// Adds 10^9 seconds (a gigasecond, 10^18 nanoseconds) to the supplied time
func AddGigasecond(t time.Time) time.Time {
	gigasecond, err := time.ParseDuration("1000000000s")
	if err!=nil{
		log.Fatal(err)
	}
	return t.Add(gigasecond)
}

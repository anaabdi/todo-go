package helper

import (
	"log"
	"time"
)

func CurrentTimestamp() int64 {
	return time.Now().UTC().Unix()
}

func RFC3339ToTime(value string) time.Time {
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		log.Printf("Invalid date format %s: %s\n", value, err.Error())
		return time.Unix(0, 0)
	}
	return parsed
}

func TimeToRFC3339(value time.Time) string {
	if value.UTC().Unix() == -62135596800 {
		// default value, return empty string
		return ""
	}
	return value.Format(time.RFC3339)
}

func TimeNext(sec int64) time.Time {
	return time.Now().Add(time.Duration(sec) * time.Second)
}

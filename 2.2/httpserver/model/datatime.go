package model

import "time"

func QueryTime() (querytime string) {

	t := time.Now()
	time := t.Format("2006-01-02 15:04:05")

	return time
}

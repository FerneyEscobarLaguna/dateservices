package dateservices

import "time"

const (
	//ApiDbLayout is the date format using to generate dates for DB
	ApiDbLayout = "2006-01-02 15:04:05"
)

//GetNow return current date
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowDBFormat return current date as string
func GetNowDBFormat() string {
	return GetNow().Format(ApiDbLayout)
}

import (
    "time"
    "fmt"
)

func dayOfTheWeek(day int, month int, year int) string {
    givenDate := time.Date(year, time.Month(month), day, 1, 0, 0, 0, time.UTC) 
    return fmt.Sprintf("%s", givenDate.Weekday())
}
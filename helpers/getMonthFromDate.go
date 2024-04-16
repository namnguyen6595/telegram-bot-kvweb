package helpers

import "time"

func GetFirstAndLastDayOfMonth(t time.Time, format string) (string, string) {
	// Ngày đầu tiên của tháng: đặt ngày là 1
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())

	// Ngày cuối cùng của tháng: thêm 1 tháng, trừ đi 1 ngày
	lastDay := firstDay.AddDate(0, 1, -1)

	return firstDay.Format(format), lastDay.Format(format)
}

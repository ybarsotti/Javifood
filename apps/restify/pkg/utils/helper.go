package utils

import (
	"strconv"
	"strings"
)

func SplitStringTimeIntoHourMinute(stringTime string) (int, int) {
	parts := strings.Split(stringTime, ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	return h, m
}

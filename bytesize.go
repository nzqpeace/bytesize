package bytesize

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	B int64 = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

var ByteSizeRegexp = regexp.MustCompile(`(?i)(-?\d+\.?\d*]?)([KMGTP]?B?)`)

func ByteSizeToString(b int64) string {
	switch {
	case b >= PB:
		return fmt.Sprintf("%.2fPB", float64(b/PB))
	case b >= TB:
		return fmt.Sprintf("%.2fTB", float64(b/TB))
	case b >= GB:
		return fmt.Sprintf("%.2fGB", float64(b/GB))
	case b >= MB:
		return fmt.Sprintf("%.2fMB", float64(b/MB))
	case b >= KB:
		return fmt.Sprintf("%.2fKB", float64(b/KB))
	}
	return fmt.Sprintf("%.2fB", float64(b))
}

func Parse(s string) (int64, error) {
	if !ByteSizeRegexp.MatchString(s) {
		return 0, errors.New("invalide byte size format")
	}

	subs := ByteSizeRegexp.FindStringSubmatch(s)
	if len(subs) != 3 {
		return 0, errors.New("invalide format of byte size")
	}

	var size int64
	switch strings.ToUpper(string(subs[2])) {
	case "B", "":
		size = 1
	case "KB", "K":
		size = KB
	case "MB", "M":
		size = MB
	case "GB", "G":
		size = GB
	case "TB", "T":
		size = TB
	case "PB", "P":
		size = PB
	default:
		return 0, errors.New("invalide format of byte size")
	}

	if strings.ContainsAny(subs[1], ".") {
		// float
		digital, err := strconv.ParseFloat(subs[1], 64)
		if err != nil {
			return 0, errors.New("invalide format of byte size")
		}
		size = int64(float64(size) * digital)
	} else {
		// int
		digital, err := strconv.ParseInt(subs[1], 10, 64)
		if err != nil {
			return 0, errors.New("invalide format of byte size")
		}
		size *= digital
	}
	return size, nil
}

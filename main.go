package stack

import (
	"runtime"
	"strings"

	"github.com/ypapax/glog3"
)

func Get() string {
	b := make([]byte, 2048) // adjust buffer size to be larger than expected stack
	n := runtime.Stack(b, false)
	s := string(b[:n])
	return s
}

func GetCut(skip, limit int) string {
	return Cut(Get(), skip, limit)
}

func Gets() string {
	return GetCut(7, 0)
}

func Cut(text string, skip, limit int) string {
	lines := strings.Split(text, "\n")
	if skip > len(lines) {
		glog3.Error("too big linesToCuntFromTop")
		return text
	}
	cut2 := lines[skip:]
	if limit == 0 {
		return strings.Join(cut2, "\n")
	}
	if limit > len(cut2) {
		glog3.Error("too big limit")
		return strings.Join(cut2, "\n")
	}
	cut3 := cut2[:limit]
	return strings.Join(cut3, "\n")
}

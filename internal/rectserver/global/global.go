package global

import "regexp"

var Patterns = map[string]*regexp.Regexp{
	"P1": regexp.MustCompile("test"),
	//"P2": regexp.MustCompile("([一-龥]{2,7})(未|没有)(按时|准时)完成.*报告.*(多少|份|个)+"),
	//"P3": regexp.MustCompile("([一-龥]{2,7})(未|没有)(按时|准时)完成.*报告.*(占比|比例)+"),
	//"P4": regexp.MustCompile("([一-龥]{2,7})(按时|准时)完成.*报告.*"),
	//"P5": regexp.MustCompile("([一-龥]{2,7})(按时|准时)完成.*报告.*(多少|份|个){1}.*"),
	//"P6": regexp.MustCompile("([一-龥]{2,7})(按时|准时)完成.*报告.*(占比|比例)+"),
}

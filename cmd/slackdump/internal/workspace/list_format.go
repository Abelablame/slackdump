package workspace

import (
	"io"
	"strings"
)

const (
	defMark    = "=>"
	timeLayout = "2006-01-02 15:04:05"
)

var hdrItems = []hdrItem{
	{"C", 1},
	{"name", 7},
	{"filename", 12},
	{"modified", 19},
	{"team", 9},
	{"user", 8},
	{"error", 5},
}

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) Write(p []byte) (n int, err error) {
	if ew.err != nil {
		return 0, nil
	}
	n, ew.err = ew.w.Write(p)
	return n, ew.err
}

func (ew *errWriter) Err() error {
	return ew.err
}

func simpleList(m manager, current string, wsps []string) [][]string {
	var rows = make([][]string, 0, len(wsps))
	for _, name := range wsps {
		timestamp := "unknown"
		filename := "-"
		if fi, err := m.FileInfo(name); err == nil {
			timestamp = fi.ModTime().Format(timeLayout)
			filename = fi.Name()
		}
		if name == current {
			name = defMark + " " + name
		} else {
			name = "   " + name
		}
		rows = append(rows, []string{name, filename, timestamp})
	}
	return rows
}

type hdrItem struct {
	name string
	size int
}

func (h *hdrItem) String() string {
	return h.name
}

func (h *hdrItem) Size() int {
	return len(h.String())
}

func (h *hdrItem) Underline(char ...string) string {
	if len(char) == 0 {
		char = []string{"-"}
	}
	return strings.Repeat(char[0], h.Size())
}

// makeHeader creates header, separating columns with tabs and underlining
// them with dashes.
// Example:
//
//	C	name	filename	modified	team	user	error
//	-	----	--------	--------	----	----	-----
func makeHeader(hi ...hdrItem) string {
	var sb strings.Builder
	for i, h := range hi {
		if i > 0 {
			sb.WriteByte('\t')
		}
		sb.WriteString(h.String())
	}
	sb.WriteByte('\n')
	for i, h := range hi {
		if i > 0 {
			sb.WriteByte('\t')
		}
		sb.WriteString(h.Underline())
	}
	return sb.String()
}

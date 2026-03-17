package types

import "bytes"

const (
	ansiReset = "\x1b[0m"
	ansiFile  = "\x1b[1;92m"
	ansiName  = "\x1b[1;97m"
	ansiLabel = "\x1b[90m"
	ansiValue = "\x1b[3m"
)

func (m *Metadata) Format() string {
	if m == nil {
		return "Unknown"
	}

	var b bytes.Buffer

	kind := m.Kind.String()
	typ := m.Type.String()

	b.Grow(len(m.File) + len(kind) + len(typ) + 64)

	if m.File != "" {
		b.WriteString(ansiFile)
		b.WriteString(m.File)
		b.WriteString(ansiReset)
		b.WriteByte('\n')
	}

	if kind != "" {
		b.WriteString("  ")
		b.WriteString(ansiFile)
		b.WriteString("└─ ")
		b.WriteString(ansiReset)
		b.WriteString(ansiName)
		b.WriteString(kind)
		b.WriteString(ansiReset)
	}

	if typ != "" {
		b.WriteByte('\n')
		b.WriteString("     ")
		b.WriteString(ansiLabel)
		b.WriteString("Type: ")
		b.WriteString(ansiReset)
		b.WriteString(ansiValue)
		b.WriteString(typ)
		b.WriteString(ansiReset)
	}

	return b.String()
}

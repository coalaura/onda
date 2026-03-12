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

	b.Grow(len(m.File) + len(m.Name) + len(m.Type) + len(m.Info) + 64)

	if m.File != "" {
		b.WriteString(ansiFile)
		b.WriteString(m.File)
		b.WriteString(ansiReset)
		b.WriteByte('\n')
	}

	if m.Name != "" {
		b.WriteString("  ")
		b.WriteString(ansiFile)
		b.WriteString("⮡ ")
		b.WriteString(ansiReset)
		b.WriteString(ansiName)
		b.WriteString(m.Name)
		b.WriteString(ansiReset)
	}

	if m.Type != "" {
		b.WriteByte('\n')
		b.WriteString("    ")
		b.WriteString(ansiLabel)
		b.WriteString("Type: ")
		b.WriteString(ansiReset)
		b.WriteString(ansiValue)
		b.WriteString(m.Type)
		b.WriteString(ansiReset)
	}

	if m.Info != "" {
		b.WriteByte('\n')
		b.WriteString("    ")
		b.WriteString(ansiLabel)
		b.WriteString("Info: ")
		b.WriteString(ansiReset)
		b.WriteString(ansiValue)
		b.WriteString(m.Info)
		b.WriteString(ansiReset)
	}

	return b.String()
}

package conf

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Load(filename string) ConfigFile {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		c := NewConfigFile()
		c.Error = err
		return c
	}

	return Reader(file)
}

func Reader(reader io.Reader) (c ConfigFile) {
	buf := bufio.NewReader(reader)

	c = NewConfigFile()

	var section, option string
	section = Default_section
	for {
		l, buffer_err := buf.ReadString('\n') // parse line-by-line
		l = strings.TrimSpace(l)

		if buffer_err != nil {
			if buffer_err != io.EOF {
				c.Error = buffer_err
				return
			}

			if len(l) == 0 {
				break
			}
		}

		switch {
		case len(l) == 0: // empty line
			continue

		case l[0] == '#': // comment
			continue

		case l[0] == ';': // comment
			continue

		case l[0] == '[' && l[len(l)-1] == ']': // new section
			option = "" // reset multi-line value
			section = strings.TrimSpace(l[1 : len(l)-1])
			c.addSection(section)

		default: // other alternatives
			i := stripKey(l)
			switch {
			case i > 0: // option and value
				option = strings.TrimSpace(l[0:i])
				option = strings.Replace(option, `\=`, `=`, -1)
				option = strings.Replace(option, `\:`, `:`, -1)

				value := strings.TrimSpace(stripComments(l[i+1:]))
				value = strings.Replace(value, `\#`, `#`, -1)
				c.addOption(section, option, value)

			default:
				continue
			}
		}
		if buffer_err == io.EOF {
			break
		}
	}
	return c
}

func stripKey(l string) (i int) {
	i = strings.IndexAny(l, "=:")
	if i < 1 {
		return i
	}
	ch := l[i-1]
	if ch == '\\' {
		ii := stripKey(l[i+1:])
		if ii == -1 {
			return ii
		} else if ii == 0 {
			return i + 1
		} else {
			return i + ii + 1
		}
	}
	return i
}

func stripComments(l string) string {
	// comments are preceded by space or TAB
	for _, c := range []string{" ;", "\t;", " #", "\t#", " //", "\t//"} {
		if i := strings.Index(l, c); i != -1 {
			l = l[0:i]
		}
	}
	return l
}

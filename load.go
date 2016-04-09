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
                i := strings.IndexAny(l, "=:")
                switch {
                    case i > 0: // option and value
                        i := strings.IndexAny(l, "=:")
                        option = strings.TrimSpace(l[0:i])
                        value := strings.TrimSpace(stripComments(l[i+1:]))
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

func stripComments(l string) string {
    // comments are preceded by space or TAB
    for _, c := range []string{" ;", "\t;", " #", "\t#"} {
        if i := strings.Index(l, c); i != -1 {
            l = l[0:i]
        }
    }
    return l
}

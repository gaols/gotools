package gotools

import "strings"

// Escape the sql to prevent sql injection.
// Simply a copy from https://gist.github.com/siddontang/8875771
func Escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]
		escape = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
			break
		case '\n': /* Must be escaped for logs */
			escape = 'n'
			break
		case '\r':
			escape = 'r'
			break
		case '\\':
			escape = '\\'
			break
		case '\'':
			escape = '\''
			break
		case '"': /* Better safe than sorry */
			escape = '"'
			break
		case '\032': /* This gives problems on Win32 */
			escape = 'Z'
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}

// InParam is a helper method to build in params of sql.
func InParam(ins []string) string {
	if len(ins) <= 0 {
		panic("input params required")
	}
	ins2 := make([]string, 0)
	for _, v := range ins {
		ins2 = append(ins2, "'"+Escape(v)+"'")
	}
	return "(" + strings.Join(ins2, ",") + ")"
}

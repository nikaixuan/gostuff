package stack

import "strings"

// 71
// stack
func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	res := make([]string, 0)
	for _, v := range paths {
		if v != "" {
			switch v {
			case ".":
				continue
			case "..":
				if len(res) > 0 {
					res = res[:len(res)-1]
				}
			default:
				res = append(res, v)
			}
		}
	}
	result := ""
	if len(res) > 0 {
		for i := 0; i < len(res); i++ {
			result = result + "/" + res[i]
		}
	} else {
		result = "/"
	}

	return result
}

package executor

import "strings"

func renderTemplate(value string, replacements map[string]string) string {
	for k, v := range replacements {
		value = strings.ReplaceAll(
			value,
			"{{"+k+"}}",
			v,
		)
	}

	return value
}

func renderArgs(
	args []string,
	replacements map[string]string,
) []string {

	out := make([]string, 0, len(args))

	for _, arg := range args {
		out = append(
			out,
			renderTemplate(arg, replacements),
		)
	}

	return out
}

package validate

import (
	"errors"
	"strings"
)

func Flags(
	flags []string,
	allowlist []string,
) error {

	if len(flags) == 0 {
		return nil
	}

	for _, flag := range flags {

		allowed := false

		for _, allowedFlag := range allowlist {

			if strings.HasSuffix(allowedFlag, "*") {

				prefix := strings.TrimSuffix(
					allowedFlag,
					"*",
				)

				if strings.HasPrefix(
					flag,
					prefix,
				) {
					allowed = true
					break
				}

			} else if flag == allowedFlag {
				allowed = true
				break
			}
		}

		if !allowed {
			return errors.New(
				"flag not allowed: " + flag,
			)
		}
	}

	return nil
}

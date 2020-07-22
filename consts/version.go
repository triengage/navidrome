package consts

import (
	"fmt"
	"strings"
)

var (
	// This will be set in build time. If not, version will be set to "dev"
	gitTag string
	gitSha string
)

// Formats:
// dev
// v0.2.0 (5b84188)
// v0.3.2-SNAPSHOT (715f552)
// master (9ed35cb)
func Version() string {
	if gitSha == "" {
		return "dev"
	}
	gitTag = strings.TrimPrefix(gitTag, "v")
	return fmt.Sprintf("%s (%s)", gitTag, gitSha)
}

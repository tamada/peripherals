package peripherals

import (
	"fmt"
)

const VERSION = "1.0.1"

func Version(appName string) string {
	return fmt.Sprintf("%s version %s", appName, VERSION)
}

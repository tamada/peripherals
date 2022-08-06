package peripherals

import (
	"fmt"
)

const VERSION = "0.9.2"

func Version(appName string) string {
	return fmt.Sprintf("%s version %s", appName, VERSION)
}

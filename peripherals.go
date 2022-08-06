package peripherals

import (
	"fmt"
)

const VERSION = "0.9.0"
const ProductName = "tamada/peripherals"

func Version() string {
	return fmt.Sprintf("%s %s", ProductName, VERSION)
}

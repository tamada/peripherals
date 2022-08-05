package peripherals

import "fmt"

const VERSION = "1.0.0"
const PRODUCT_NAME = "tamada/peripherals"

func Version() string {
	return fmt.Sprintf("%s %s", PRODUCT_NAME, VERSION)
}

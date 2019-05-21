package poker

import (
	"fmt"
	"io"
	"time"
)

// BlindAlerter schedules alerts for blind amounts
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// BlindAlerterFunc allows you to implement BlindAlerter with a function
// NOTE: Any type can implement an interface, not just structs.
//       If you are making a library that exposes an interface with ONE function defined, it is a common idiom to also expose a MyInterfaceFunc type.
//       This type will be a func which will also implement your interface.
//       That way users of your interface have the option to implement your interface with just a function; rather than having to create an empty struct type.
type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	a(duration, amount, to)
}

// Alerter will schedule alerts and print them to os.Stdout
func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}

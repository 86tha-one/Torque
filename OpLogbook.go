package Torque

import "errors"

type    OpLogbook  struct {
	LaunchNotificationChannel chan error
	ShutdownFlag bool
}
func(O *OpLogbook) RecordFLaunch (E error) {
	if (E == nil) { E = errors.New ("-- error not described --") }
	O.LaunchNotificationChannel <- E
}
func(O *OpLogbook) RecordSLaunch () {
	O.LaunchNotificationChannel <- nil
}

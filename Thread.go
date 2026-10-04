package Torque

import "errors"
import "fmt"
import "runtime/debug"
import "sync"

type    Worker  struct {
	uuid    string
	mtxx    *sync.Mutex
	Plss     bool // Plss
	RErr    error // Runtime error
	prgr     func(string,*OpLogbook)
	OpLb    *OpLogbook
	shWh     func()
}
func(O *Worker) Init (UUID string, Program func(string,*OpLogbook), ShutdownWebhook func()) {
	if UUID == "" { panic ("Worker's UUID not provided") }
	if Program == nil { panic ("Worker's program not provided") }
	O.uuid = UUID
	O.mtxx =&sync.Mutex{ }
	O.prgr = Program
	O.shWh = ShutdownWebhook
}
func(O *Worker) Launch() (error) {
	/***1***/
	if O.uuid ==  "" { panic ("Worker was not initialized") }
	if O.prgr == nil { panic ("Worker was not initialized") }
	O.mtxx.Lock (); defer O.mtxx.Unlock ()
	/***2***/
	if O.Plss {
		return errors.New("Worker has active instance")
	}
	O.OpLb= &OpLogbook { LaunchNotificationChannel: make(chan error, 1) }
	O.RErr= nil
	O.Plss= true
	/***3***/
	go   	func () {
		defer   func() {
			BE05 := recover ()
			if BE05 !=  nil {
				CE05 := fmt.Sprintf (`Worker failed [%v]`, string(debug.Stack()))
				O.RErr = errors.New (CE05)
			}
			O.Plss = false
		}()
		O.prgr (O.uuid, O.OpLb)
	}       ( )
return  nil
}
func(O *Worker) Shtdwn() (error) {
	/***1***/
	if O.uuid ==  "" { panic ("Worker was not initialized") }
	if O.prgr == nil { panic ("Worker was not initialized") }
	O.mtxx.Lock (); defer O.mtxx.Unlock ()
	/***2***/
	if O.Plss {
		return errors.New("Worker not running")
	}
	/***3***/
	if O.OpLb.ShutdownFlag { return nil }
	if O.shWh != nil { O.shWh () }
	O.OpLb.ShutdownFlag = true
return  nil
}

package Torque

import "errors"
import "fmt"
import "runtime/debug"
import "sync"
import "time"

type    Worker  struct {
	UUID    string
	mtxx    *sync.Mutex
	Plss     bool // Plss
	RErr    error // Runtime error
	prgr     func(string,*OpLogbook)
	opLb    *OpLogbook
	shWh     func()
}
func(O *Worker) Init (UUID string, Program func(string,*OpLogbook), ShutdownWebhook func()) {
	if UUID == "" { panic ("Worker's UUID not provided") }
	if Program == nil { panic ("Worker's program not provided") }
	O.UUID = UUID
	O.mtxx =&sync.Mutex{ }
	O.prgr = Program
	O.shWh = ShutdownWebhook
}
func(O *Worker) Launch (LaunchGrace time.Duration) (error) {
	/***1***/
	if O.UUID ==  "" { panic ("Worker was not initialized") }
	if O.prgr == nil { panic ("Worker was not initialized") }
	O.mtxx.Lock (); defer O.mtxx.Unlock ()
	/***2***/
	if O.Plss {
		return errors.New("Worker has active instance")
	}
	O.opLb= &OpLogbook { LaunchNotificationChannel: make(chan error, 1) }
	O.RErr= nil
	O.Plss= true
	/***3***/
	AE05 := make (chan bool, 1)
	AE10 := make (chan bool, 1)
	if LaunchGrace != time.Second*0 {
		go func ( ) {
			time.Sleep (LaunchGrace)
			AE10 <- true
		}  (  )
	}
	go func (  )  {
		defer   func() {
			CE05 := recover ()
			if CE05 !=  nil {
				DE05 := fmt.Sprintf (`Worker failed [%v:%v]`, string(debug.Stack()), CE05)
				O.RErr = errors.New (DE05)
			}
			O.Plss = false
			AE05 <- true
		}  (  )
		O.prgr (O.UUID, O.opLb)
	}  (  )
	/***4***/
	select  {
		case _= <- AE05: {
			CE05 := fmt.Sprintf (`Got no launch failed/succeeded signal, and worker has shutdown`)
			if O.RErr != nil { CE05 = fmt.Sprintf (`Got no launch failed/succeeded signal, and worker has shutdown [%s]`, O.RErr) }
			return errors.New( CE05 )
		}
		case CE05 := <- O.opLb.LaunchNotificationChannel: {
			if CE05 !=  nil {
				DE05 := fmt.Sprintf (`Launch failed [%s]`, CE05.Error ())
				return errors.New (DE05)
			}
			return nil
		}
		case _= <- AE10: {
			return errors.New ("TIMEOUT")
		}
	}
}
func(O *Worker) Shtdwn (ShtdwnGrace time.Duration) (error) {
	/***1***/
	if O.UUID ==  "" { panic ("Worker was not initialized") }
	if O.prgr == nil { panic ("Worker was not initialized") }
	O.mtxx.Lock (); defer O.mtxx.Unlock ()
	/***2***/
	if O.Plss ==false{
		return errors.New("Worker not running")
	}
	/***3***/
	if O.opLb.ShutdownFlag { return nil }
	if O.shWh != nil { O.shWh () }
	O.opLb.ShutdownFlag = true
	/***4***/
	AE05 := time.Now ().Add (ShtdwnGrace).Unix ()
	for     {
		if O.Plss == false { return nil }
		if ShtdwnGrace != 0 && (time.Now ().Unix () > AE05) { return errors.New ("TIMEOUT") }
		time.Sleep (time.Millisecond*100)
	}
}

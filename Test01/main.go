package main

import "fmt"
import "github.com/86tha-one/Torque"
import "time"

func    PR01 (I string, L *Torque.OpLogbook) {
	BE05 := fmt.Sprintf ("%s: Running...", I)
	for  {
		fmt.Println (BE05)
		time.Sleep (time.Second*1)
	}
}

func    main () {
	BE05 := &Torque.Torque {}
	BE11 := BE05.Worker ()
	BE12 := BE05.Worker ()
	BE11.Init ("PR01", PR01, nil)
	BE12.Init ("PR02", PR01, nil)
	BE15 := BE05.LaunchSequence ()
	for _ , CE10 := range  BE15 {
		CE15 := CE10.Launch ()
		if CE15 !=  nil { fmt.Println (CE15); return }
		CE20 := <- CE10.OpLb.LaunchNotificationChannel
		if CE20 !=  nil { fmt.Println (CE20); return }
	}
	time.Sleep (time.Second*256)
}

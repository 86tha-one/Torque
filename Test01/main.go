package main

// import "errors"
import "fmt"
import "github.com/86tha-one/Torque"
import "time"

func    SD01 () { fmt.Println ("Shutdown signal received") }
func    PR01 (I string, L *Torque.OpLogbook) {
	L.RecordSLaunch ( )
//	time.Sleep (time.Second*5)
//	if true { panic ("I dont feel like working") }
	BE05 := fmt.Sprintf ("%s: Running...", I)
	for  {
		fmt.Println (BE05)
		if L.ShutdownFlag { return }
		time.Sleep (time.Second*1)
	}
}

func    main () {
	BE05 := &Torque.Torque {}
	BE11 := BE05.Worker ()
	BE12 := BE05.Worker ()
	BE11.Init ("PR01", PR01, SD01)
	BE12.Init ("PR02", PR01, SD01)
	BE15 := BE05.LaunchSequence ()
	BE20 := BE05.ShtdwnSequence ()
	for _ , CE10 := range  BE15 {
		CE15 := CE10.Launch (time.Second*5)
		if  CE15 !=  nil { fmt.Println (CE15); return }
		fmt.Println ("Worker launched")
	}
	/*
	for _ , CE10 := range  BE20 {
		for CE10.Plss { time.Sleep (time.Millisecond*100) }
		fmt.Println (CE10.RErr)
	}*/
	time.Sleep (time.Second*5)
	for _ , CE10 := range  BE20 {
		CE15 := CE10.Shtdwn (time.Second*5)
		if  CE15 !=  nil { fmt.Println (CE15); continue }
		fmt.Println ("Worker shutdown")
	}
}

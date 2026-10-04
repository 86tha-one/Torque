package Torque

type    Torque  struct {
	w []*Worker
}
func(O *Torque) Worker (UUID ... int) (T *Worker) {
	/***1***/
	if O.w  ==  nil { O.w = []*Worker { } }
	if UUID ==  nil || len(UUID) == 0 {
		CE05 := &Worker{}
		O.w = append (O.w, CE05)
		return CE05
	}
	/***2***/
	if UUID[0] < 0 { return nil }
	if len(O.w) < UUID[0]+1 { return nil; }
	return O.w[UUID[0]]
}
func(O *Torque) LaunchSequence () (T []*Worker) {
	if O.w  ==  nil { O.w = []*Worker { } }
	return O.w
}
func(O *Torque) ShtdwnSequence () (T []*Worker) {
	if O.w  ==  nil { O.w = []*Worker { } }
	BE05 := []*Worker {}
	for CE05 := len(O.w); CE05 >= 1; CE05-- { BE05 = append (BE05, O.w[CE05-1]) }
	return BE05
}

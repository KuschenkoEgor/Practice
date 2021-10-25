package main

type st struct {
	On bool
	Ammo, Power int
}

func (s *st) Shoot() bool {
	if s.On == false {
		return false
	}
	if s.Ammo != 0 {
		s.Ammo -=1
		return true
	} else {
		return false
	}
}

func (s *st) RideBike() bool {
	if s.On == false {
		return false
	}
	if s.Power != 0 {
		s.Power -=1
		return true
	} else {
		return false
	}
}
func main()  {
	testStruct := new(st)
	testStruct.Shoot()
	testStruct.RideBike()

}
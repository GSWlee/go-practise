package length

type Meter float64
type Kilometer float64

func MToK( m Meter) Kilometer { return Kilometer(m/1000) }
func KToM(k Kilometer) Meter {return Meter(k*1000)}


package fnTime

import "time"

func TrimYMDH(vs ...time.Time) (t time.Time) {
	t = time.Now()
	if len(vs) == 1 {
		t = vs[0]
	}
	t = t.UTC()
	t = t.Truncate(time.Second)
	t = t.Add(-time.Duration(t.Minute()) * time.Minute)
	t = t.Add(-time.Duration(t.Second()) * time.Second)
	return
}

func TrimYMD(vs ...time.Time) (t time.Time) {
	t = time.Now()
	if len(vs) == 1 {
		t = vs[0]
	}
	t = t.UTC()
	t = t.Truncate(time.Second)
	t = t.Add(-time.Duration(t.Hour()) * time.Hour)
	t = t.Add(-time.Duration(t.Minute()) * time.Minute)
	t = t.Add(-time.Duration(t.Second()) * time.Second)
	return
}

func TrimYM(vs ...time.Time) (t time.Time) {
	t = time.Now()
	if len(vs) == 1 {
		t = vs[0]
	}
	t = t.UTC()
	t = t.Truncate(time.Second)
	t = t.Add(-time.Duration(t.Day()-1) * time.Hour * 24)
	t = t.Add(-time.Duration(t.Hour()) * time.Hour)
	t = t.Add(-time.Duration(t.Minute()) * time.Minute)
	t = t.Add(-time.Duration(t.Second()) * time.Second)

	return
}

func TrimY(vs ...time.Time) (t time.Time) {
	t = time.Now()
	if len(vs) == 1 {
		t = vs[0]
	}
	t = t.UTC()
	t = t.Truncate(time.Second)
	t = t.Add(-time.Duration(t.YearDay()-1) * time.Hour * 24)
	t = t.Add(-time.Duration(t.Day()-1) * time.Hour * 24)
	t = t.Add(-time.Duration(t.Hour()) * time.Hour)
	t = t.Add(-time.Duration(t.Minute()) * time.Minute)
	t = t.Add(-time.Duration(t.Second()) * time.Second)
	return
}

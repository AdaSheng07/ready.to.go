package main

type Assets struct {
	assets []Asset
}

func (a *Assets) OnDuty() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Open()
		}
	}
}

func (a *Assets) OffDuty() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Close()
		}
	}
}

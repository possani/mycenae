package meta

// ComposeID joins a ksid with a tsid
func ComposeID(ksid, tsid string) (ksts string) {
	lid := len(tsid)
	lksid := len(ksid)
	x := make([]byte, lid+lksid+1)
	copy(x, ksid)
	copy(x[lksid:], "|")
	copy(x[lksid+1:], tsid)
	return string(x)
}

package objects

type UIDHandler struct {
	Count int
}

func (u *UIDHandler) getNewUID() int {
	u.Count++
	return u.Count
}

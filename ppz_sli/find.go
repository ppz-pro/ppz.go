package ppz_sli

func (slice Slice[El]) Find3(cb func(El, int, Slice[El]) bool) (El, int, bool) {
	for index, item := range slice {
		if cb(item, index, slice) {
			return item, index, true
		}
	}
	var nil_item El
	return nil_item, 0, false
}

func (slice Slice[El]) Find2(cb func(El, int) bool) (El, int, bool) {
	return slice.Find3(func(el El, index int, _ Slice[El]) bool {
		return cb(el, index)
	})
}

// Find 返回第一个满足某个（或某些）条件的元素，及其下标，及“是否找到”
func (slice Slice[El]) Find(cb func(El) bool) (El, int, bool) {
	return slice.Find3(func(el El, _ int, _ Slice[El]) bool {
		return cb(el)
	})
}

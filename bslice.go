package bslice

func batch1(size, batch int, f func(int, int) bool) {
	i, j := 0, batch
	for i < size {
		if j > size || j < 0 {
			f(i, size)
			break
		}
		if f(i, j) {
			break
		}
		i, j = j, j+batch
	}
}

func batch2(size, batch int, f func(int, int) bool) {
	i, j := size, size+batch
	for i > 0 {
		if j <= 0 {
			f(0, i)
			break
		}
		if f(j, i) {
			break
		}
		i, j = j, j+batch
	}
}

// Batch loop over slice in batches.
// If f return true, loop will break.
// - size is length of slice.
// - batch is length of slice each iteration.
//   If batch is negative number, iterate from end to beginning of slice.
func Batch(size, batch int, f func(int, int) bool) {
	if batch == 0 {
		batch = size
	}
	if batch > 0 {
		batch1(size, batch, f)
	} else {
		batch2(size, batch, f)
	}
}

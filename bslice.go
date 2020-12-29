package bslice

func batch1(tsize, bsize int, f func(int, int) bool) {
	i, j := 0, bsize
	for i < tsize {
		if j > tsize || j < 0 {
			f(i, tsize)
			break
		}
		if f(i, j) {
			break
		}
		i, j = j, j+bsize
	}
}

func batch2(tsize, bsize int, f func(int, int) bool) {
	i, j := tsize, tsize+bsize
	for i > 0 {
		if j <= 0 {
			f(0, i)
			break
		}
		if f(j, i) {
			break
		}
		i, j = j, j+bsize
	}
}

// Batch loops through the target slice in batches and calls the function in each batch.
// tsize is length of target slice.
// bsize is length of slice in each batch.
// If bsize is negative number, iterate from end to beginning of target slice.
// If f return true, loop will break.
func Batch(tsize, bsize int, f func(int, int) bool) {
	if bsize == 0 {
		bsize = tsize
	}
	if bsize > 0 {
		batch1(tsize, bsize, f)
	} else {
		batch2(tsize, bsize, f)
	}
}

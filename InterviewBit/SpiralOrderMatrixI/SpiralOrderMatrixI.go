package SpiralOrderMatrixI

/**
 * @input A : 2D integer array
 *
 * @Output Integer array.
 */
func spiralOrder(A [][]int ) ([]int) {
	t := 0
	b := len(A) - 1
	l := 0
	r := len(A[0]) - 1
	dir := 0

	res := make([]int, 0, len(A)*len(A[0]))
	for ; t <= b && l <= r; {
		switch dir {
		case 0:
			for i := l; i <= r; i++ {
				res = append(res, A[t][i])
			}
			t++
		case 1:
			for i := t; i <= b; i++ {
				res = append(res, A[i][r])
			}
			r--
		case 2:
			for i := r; i >= l; i-- {
				res = append(res, A[b][i])
			}
			b--
		case 3:
			for i := b; i >= t; i-- {
				res = append(res, A[i][l])
			}
			l++
		}
		dir = (dir+1)%4
	}
	return res
}
package main

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {

	l,r := 1,n

	for l <= r {
		m := l + (r-l) / 2
		if isBadVersion(m) {
			r = m-1
		}else {
			l = m+1
		}
	}
	return l
}

func main() {

}

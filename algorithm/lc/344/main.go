package main

func reverseString(s []byte)  {
	for i := 0; i < len(s)/2; i++ {
		s[i],s[len(s) - 1 - i] = s[len(s) - 1 - i],s[i]
	}
}

func reverseString1(s []byte)  {
	l,r := 0,len(s)-1

	for l< r{
		s[l],s[r] = s[r],s[l]
		r--
		l++
	}
}

func main() {

}

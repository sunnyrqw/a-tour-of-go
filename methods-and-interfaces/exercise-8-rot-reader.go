package tour

import "io"

type rot13Reader struct {
	r io.Reader
}

/*
Reader is the interface that wraps the basic Read method.
Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p)) and any error encountered.
About Reader more see : https://pkg.go.dev/io#Reader
*/
func (reader rot13Reader) Read(p []byte) (n int, err error) {
	n, err = reader.r.Read(p)
	for i := 0; i < n; i++ {
		if p[i] >= 'A' && p[i] <= 'z' {
			p[i] += 13
			if p[i] > 'z' {
				p[i] -= 26
			}
		}
	}
	return n, err
}

package tour

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 'A'
	}
	return len(bytes), nil
}

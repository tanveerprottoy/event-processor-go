package file

import "log"

// Progress implements the io.Writer
type Progress struct {
	TotalSize int64
	BytesRead int64
}

func NewProgress(totalSize int64) *Progress {
	return &Progress{TotalSize: totalSize}
}

// print prints out the current progress
func (pr Progress) print() {
	if pr.BytesRead == pr.TotalSize {
		log.Println("DONE!")
		return
	}

	log.Printf("File upload in progress: %d\n", pr.BytesRead)
}

// Write is implemented to satisfy the io.Writer interface
// it calls print to print out the bytes
func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.print()
	return
}

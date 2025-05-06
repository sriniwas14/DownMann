package download

import "io"

type ProgressWriter struct {
	writer io.Writer
	part   *DownloadPart
}

func (p *ProgressWriter) Write(b []byte) (int, error) {
	n, err := p.writer.Write(b)
	p.part.From += int64(n)
	return n, err
}

package awsbilling

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"strings"
)

// Read S3 object content and return all lines. Will skip first line if headline flag is set.
func readObject(objectKey string, content []byte, skipHeadline bool) ([]string, error) {

	lines := []string{}

	if isZipped(objectKey) {
		uzippedContent, err := unzip(content)
		if err != nil {
			return lines, err
		}
		content = uzippedContent
	}

	contentReader := bytes.NewReader(content)
	scanner := bufio.NewScanner(contentReader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if skipHeadline && len(lines) > 1 {
		return lines[1:], scanner.Err()
	}
	return lines, scanner.Err()
}

// isZipped returns true if passed object key ends with "gz".
func isZipped(objectKey string) bool {
	return strings.HasSuffix(objectKey, "gz")
}

// unzip unzip
func unzip(zipedData []byte) ([]byte, error) {

	b := bytes.NewBuffer(zipedData)
	r, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	}

	var unzipedBuffer bytes.Buffer
	_, err = unzipedBuffer.ReadFrom(r)
	return unzipedBuffer.Bytes(), err
}

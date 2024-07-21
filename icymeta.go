// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Package icymeta parses metadata from SHOUTcast audio streams.
package icymeta

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

// GetCurrentStreamTitle opens a stream to the specified url, parses the icy metadata and returns the current stream title.
func GetCurrentStreamTitle(ctx context.Context, url string) (string, error) {
	resp, skip, err := openStream(ctx, url)

	if err != nil {
		return "", err
	}

	skipBuffer := make([]byte, skip)

	reader := bufio.NewReaderSize(resp.Body, skip)

	_, err = io.ReadFull(reader, skipBuffer)

	if err != nil {
		return "", fmt.Errorf("failed to skip bytes: %v", err)
	}

	meta, err := ReadMeta(reader)

	if err != nil {
		return "", err
	}

	return ParseStreamTitle(meta)
}

func openStream(ctx context.Context, url string) (*http.Response, int, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request for %v: %v", url, err)
	}

	req.Header.Add("Icy-MetaData", "1")

	resp, err := client.Do(req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to open %v: %v", url, err)
	}

	metaInt := resp.Header.Get("Icy-Metaint")

	skip, err := strconv.Atoi(metaInt)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to parse Icy-Metaint value '%v': %v", metaInt, err)
	}

	return resp, skip, nil
}

// ReadMeta reads the icy metadata from the reader and returns it as string.
// The bytes following the icy metadata remain untouched in the reader.
func ReadMeta(reader *bufio.Reader) (string, error) {
	length, err := reader.ReadByte()

	if err != nil {
		return "", err
	}

	byteCount := int(length) * 16
	buffer := make([]byte, byteCount)
	_, err = io.ReadFull(reader, buffer)

	if err != nil {
		return "", fmt.Errorf("failed to metadata bytes: %v", err)
	}

	return string(buffer), nil
}

// ParseStreamTitle extracts the stream title from the icy metadata.
func ParseStreamTitle(meta string) (string, error) {
	re := regexp.MustCompile("StreamTitle='([^;]*)';")

	matches := re.FindStringSubmatch(meta)

	if len(matches) < 2 {
		return "", fmt.Errorf("failed to match stream title in \"%v\"", meta)
	}

	return matches[1], nil
}

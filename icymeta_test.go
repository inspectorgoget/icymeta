// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package icymeta

import (
	"bufio"
	"bytes"
	"testing"
)

// TestReadMeta calls ReadMeta with a byte stream, checking for a valid return value.
func TestReadMeta(t *testing.T) {
	const blockCount = 2
	const blockSize = 16

	length := 1 + (blockCount * blockSize) + 3
	buffer := make([]byte, length)
	buffer[0] = blockCount
	buffer[1] = 'f'
	buffer[2] = 'o'
	buffer[3] = 'o'
	buffer[length-3] = 1
	buffer[length-2] = 2
	buffer[length-1] = 3
	reader := bufio.NewReader(bytes.NewReader(buffer))

	title, err := ReadMeta(reader)

	expectedBytes := make([]byte, blockCount*blockSize)
	expectedBytes[0] = 'f'
	expectedBytes[1] = 'o'
	expectedBytes[2] = 'o'
	expected := string(expectedBytes)

	if title != expected || err != nil {
		t.Fatalf("\ngot  %q\nwant %q", title, expected)
	}

	payload := make([]byte, 3)
	_, err = reader.Read(payload)

	if err != nil {
		t.Fatalf("Unexpected error reading payload: %s", err)
	}

	expectedPayload := []byte{1, 2, 3}

	if !bytes.Equal(payload, expectedPayload) {
		t.Fatalf("\ngot  %v\nwant %v", payload, expectedPayload)
	}
}

// TestParseStreamTitle calls ParseStreamTitle with a stream title, checking for a valid return value.
func TestParseStreamTitle(t *testing.T) {
	title, err := ParseStreamTitle("StreamTitle='Beethoven - Moonlight Sonata';")

	if err != nil {
		t.Fatalf("Unexpected error parsing stream title: %s", err)
	}

	expected := "Beethoven - Moonlight Sonata"

	if title != expected {
		t.Fatalf("\ngot  %q\nwant %q", title, expected)
	}
}

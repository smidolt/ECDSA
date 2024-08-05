package internal

import (
    "bytes"
    "compress/gzip"
    "encoding/base64"
    "errors"
    "log"
)

// Statuses stores the statuses as a slice of bytes
type Statuses struct {
    Data []byte
}

// NewStatuses creates a new instance of Statuses
func NewStatuses() *Statuses {
    return &Statuses{Data: make([]byte, 1)}
}

// Set sets the status at the given index
func (s *Statuses) Set(index int, value bool) error {
    byteIndex := index / 8
    bitIndex := uint(index % 8)

    log.Printf("Setting status: index=%d, value=%t, byteIndex=%d, bitIndex=%d", index, value, byteIndex, bitIndex)

    if byteIndex >= len(s.Data) {
        newData := make([]byte, byteIndex+1)
        copy(newData, s.Data)
        s.Data = newData
    }

    if value {
        s.Data[byteIndex] |= 1 << bitIndex
    } else {
        s.Data[byteIndex] &^= 1 << bitIndex
    }

    log.Printf("Updated Data: %v", s.Data)

    return nil
}

// Get gets the status at the given index
func (s *Statuses) Get(index int) (bool, error) {
    byteIndex := index / 8
    bitIndex := uint(index % 8)

    if byteIndex >= len(s.Data) {
        return false, errors.New("index out of range")
    }

    return s.Data[byteIndex]&(1<<bitIndex) != 0, nil
}

// Add adds a new status to the list and returns its index
func (s *Statuses) Add(value bool) int {
    index := len(s.Data) * 8
    _ = s.Set(index, value)
    return index
}

// EncodeBase64 encodes the statuses in gzipped base64 format
func (s *Statuses) EncodeBase64() (string, error) {
    var buf bytes.Buffer
    gzipWriter := gzip.NewWriter(&buf)
    if _, err := gzipWriter.Write(s.Data); err != nil {
        return "", err
    }
    if err := gzipWriter.Close(); err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// DecodeBase64 decodes gzipped base64 format to statuses
func (s *Statuses) DecodeBase64(encoded string) error {
    gzippedData, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        return err
    }

    gzipReader, err := gzip.NewReader(bytes.NewReader(gzippedData))
    if err != nil {
        return err
    }
    defer gzipReader.Close()

    var buf bytes.Buffer
    if _, err := buf.ReadFrom(gzipReader); err != nil {
        return err
    }

    s.Data = buf.Bytes()
    return nil
}

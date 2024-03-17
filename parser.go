package omadm

import (
	"encoding/xml"
	"fmt"
	"io"
)

// ParseSyncMessage parses an OMA DM sync message from an io.Reader.
func ParseSyncMessage(reader io.Reader) (*SyncML, error) {
	var syncMessage SyncML
	decoder := xml.NewDecoder(reader)
	if err := decoder.Decode(&syncMessage); err != nil {
		return nil, fmt.Errorf("error decoding sync message: %v", err)
	}
	return &syncMessage, nil
}

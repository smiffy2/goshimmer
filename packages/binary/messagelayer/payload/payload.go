package payload

import (
	"fmt"

	"github.com/iotaledger/hive.go/marshalutil"
)

const (
	// ObjectName defines the name of the data object.
	ObjectName = "data"

	// MaxMessageSize defines the maximum size of a message.
	MaxMessageSize = 64 * 1024

	// MaxPayloadSize defines the maximum size of a payload.
	// trunkID + branchID + issuerPublicKey + issuingTime + sequenceNumber + nonce + signature
	MaxPayloadSize = MaxMessageSize - 64 - 64 - 32 - 8 - 8 - 8 - 64
)

var (
	// ErrMaxPayloadSizeExceeded is returned if the maximum payload size is exceeded.
	ErrMaxPayloadSizeExceeded = fmt.Errorf("maximum payload size of %d bytes exceeded", MaxPayloadSize)
)

func init() {
	// register the generic unmarshaler
	SetGenericUnmarshalerFactory(GenericPayloadUnmarshalerFactory)
	// register the generic data payload type
	RegisterType(DataType, ObjectName, GenericPayloadUnmarshalerFactory(DataType))
}

// Payload represents some kind of payload of data which only gains meaning by having
// corresponding node logic processing payloads of a given type.
type Payload interface {
	// Type returns the type of the payload.
	Type() Type
	// Bytes returns the payload bytes.
	Bytes() []byte
	// Unmarshal unmarshals the payload from the given bytes.
	Unmarshal(bytes []byte) error
	// String returns a human-friendly representation of the payload.
	String() string
}

// FromBytes unmarshals bytes into a payload.
func FromBytes(bytes []byte) (result Payload, consumedBytes int, err error) {
	// initialize helper
	marshalUtil := marshalutil.New(bytes)

	// calculate result
	payloadType, err := marshalUtil.ReadUint32()
	if err != nil {
		return
	}

	payloadSize, err := marshalUtil.ReadUint32()
	if err != nil {
		return
	}

	if payloadSize > MaxPayloadSize {
		err = fmt.Errorf("%w: %d", ErrMaxPayloadSizeExceeded, payloadSize)
		return
	}

	marshalUtil.ReadSeek(marshalUtil.ReadOffset() - marshalutil.UINT32_SIZE*2)
	payloadBytes, err := marshalUtil.ReadBytes(int(payloadSize) + 8)
	if err != nil {
		return
	}

	readOffset := marshalUtil.ReadOffset()
	result, err = GetUnmarshaler(payloadType)(payloadBytes)
	if err != nil {
		// fallback to the generic unmarshaler if registered type fails to unmarshal
		marshalUtil.ReadSeek(readOffset)
		result, err = GenericPayloadUnmarshalerFactory(payloadType)(payloadBytes)
		if err != nil {
			return
		}
	}

	// return the number of bytes we processed
	consumedBytes = marshalUtil.ReadOffset()
	return
}

// Parse parses a payload by using the given marshal util.
func Parse(marshalUtil *marshalutil.MarshalUtil) (Payload, error) {
	if payload, err := marshalUtil.Parse(func(data []byte) (interface{}, int, error) { return FromBytes(data) }); err != nil {
		return nil, err
	} else {
		return payload.(Payload), nil
	}
}

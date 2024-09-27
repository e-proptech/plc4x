/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// NLM is the corresponding interface of NLM
type NLM interface {
	NLMContract
	NLMRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsNLM is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLM()
	// CreateBuilder creates a NLMBuilder
	CreateNLMBuilder() NLMBuilder
}

// NLMContract provides a set of functions which can be overwritten by a sub struct
type NLMContract interface {
	// GetIsVendorProprietaryMessage returns IsVendorProprietaryMessage (virtual field)
	GetIsVendorProprietaryMessage() bool
	// GetApduLength() returns a parser argument
	GetApduLength() uint16
	// IsNLM is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLM()
	// CreateBuilder creates a NLMBuilder
	CreateNLMBuilder() NLMBuilder
}

// NLMRequirements provides a set of functions which need to be implemented by a sub struct
type NLMRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetIsVendorProprietaryMessage returns IsVendorProprietaryMessage (discriminator field)
	GetIsVendorProprietaryMessage() bool
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
}

// _NLM is the data-structure of this message
type _NLM struct {
	_SubType NLM

	// Arguments.
	ApduLength uint16
}

var _ NLMContract = (*_NLM)(nil)

// NewNLM factory function for _NLM
func NewNLM(apduLength uint16) *_NLM {
	return &_NLM{ApduLength: apduLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMBuilder is a builder for NLM
type NLMBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() NLMBuilder
	// AsNLMWhoIsRouterToNetwork converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMWhoIsRouterToNetwork() interface {
		NLMWhoIsRouterToNetworkBuilder
		Done() NLMBuilder
	}
	// AsNLMIAmRouterToNetwork converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMIAmRouterToNetwork() interface {
		NLMIAmRouterToNetworkBuilder
		Done() NLMBuilder
	}
	// AsNLMICouldBeRouterToNetwork converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMICouldBeRouterToNetwork() interface {
		NLMICouldBeRouterToNetworkBuilder
		Done() NLMBuilder
	}
	// AsNLMRejectMessageToNetwork converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMRejectMessageToNetwork() interface {
		NLMRejectMessageToNetworkBuilder
		Done() NLMBuilder
	}
	// AsNLMRouterBusyToNetwork converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMRouterBusyToNetwork() interface {
		NLMRouterBusyToNetworkBuilder
		Done() NLMBuilder
	}
	// AsNLMRouterAvailableToNetwork converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMRouterAvailableToNetwork() interface {
		NLMRouterAvailableToNetworkBuilder
		Done() NLMBuilder
	}
	// AsNLMInitializeRoutingTable converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMInitializeRoutingTable() interface {
		NLMInitializeRoutingTableBuilder
		Done() NLMBuilder
	}
	// AsNLMInitializeRoutingTableAck converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMInitializeRoutingTableAck() interface {
		NLMInitializeRoutingTableAckBuilder
		Done() NLMBuilder
	}
	// AsNLMEstablishConnectionToNetwork converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMEstablishConnectionToNetwork() interface {
		NLMEstablishConnectionToNetworkBuilder
		Done() NLMBuilder
	}
	// AsNLMDisconnectConnectionToNetwork converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMDisconnectConnectionToNetwork() interface {
		NLMDisconnectConnectionToNetworkBuilder
		Done() NLMBuilder
	}
	// AsNLMChallengeRequest converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMChallengeRequest() interface {
		NLMChallengeRequestBuilder
		Done() NLMBuilder
	}
	// AsNLMSecurityPayload converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMSecurityPayload() interface {
		NLMSecurityPayloadBuilder
		Done() NLMBuilder
	}
	// AsNLMSecurityResponse converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMSecurityResponse() interface {
		NLMSecurityResponseBuilder
		Done() NLMBuilder
	}
	// AsNLMRequestKeyUpdate converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMRequestKeyUpdate() interface {
		NLMRequestKeyUpdateBuilder
		Done() NLMBuilder
	}
	// AsNLMUpdateKeyUpdate converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMUpdateKeyUpdate() interface {
		NLMUpdateKeyUpdateBuilder
		Done() NLMBuilder
	}
	// AsNLMUpdateKeyDistributionKey converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMUpdateKeyDistributionKey() interface {
		NLMUpdateKeyDistributionKeyBuilder
		Done() NLMBuilder
	}
	// AsNLMRequestMasterKey converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMRequestMasterKey() interface {
		NLMRequestMasterKeyBuilder
		Done() NLMBuilder
	}
	// AsNLMSetMasterKey converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMSetMasterKey() interface {
		NLMSetMasterKeyBuilder
		Done() NLMBuilder
	}
	// AsNLMWhatIsNetworkNumber converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMWhatIsNetworkNumber() interface {
		NLMWhatIsNetworkNumberBuilder
		Done() NLMBuilder
	}
	// AsNLMNetworkNumberIs converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMNetworkNumberIs() interface {
		NLMNetworkNumberIsBuilder
		Done() NLMBuilder
	}
	// AsNLMReserved converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMReserved() interface {
		NLMReservedBuilder
		Done() NLMBuilder
	}
	// AsNLMVendorProprietaryMessage converts this build to a subType of NLM. It is always possible to return to current builder using Done()
	AsNLMVendorProprietaryMessage() interface {
		NLMVendorProprietaryMessageBuilder
		Done() NLMBuilder
	}
	// Build builds the NLM or returns an error if something is wrong
	PartialBuild() (NLMContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() NLMContract
	// Build builds the NLM or returns an error if something is wrong
	Build() (NLM, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLM
}

// NewNLMBuilder() creates a NLMBuilder
func NewNLMBuilder() NLMBuilder {
	return &_NLMBuilder{_NLM: new(_NLM)}
}

type _NLMChildBuilder interface {
	utils.Copyable
	setParent(NLMContract)
	buildForNLM() (NLM, error)
}

type _NLMBuilder struct {
	*_NLM

	childBuilder _NLMChildBuilder

	err *utils.MultiError
}

var _ (NLMBuilder) = (*_NLMBuilder)(nil)

func (b *_NLMBuilder) WithMandatoryFields() NLMBuilder {
	return b
}

func (b *_NLMBuilder) PartialBuild() (NLMContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NLM.deepCopy(), nil
}

func (b *_NLMBuilder) PartialMustBuild() NLMContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NLMBuilder) AsNLMWhoIsRouterToNetwork() interface {
	NLMWhoIsRouterToNetworkBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMWhoIsRouterToNetworkBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMWhoIsRouterToNetworkBuilder().(*_NLMWhoIsRouterToNetworkBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMIAmRouterToNetwork() interface {
	NLMIAmRouterToNetworkBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMIAmRouterToNetworkBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMIAmRouterToNetworkBuilder().(*_NLMIAmRouterToNetworkBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMICouldBeRouterToNetwork() interface {
	NLMICouldBeRouterToNetworkBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMICouldBeRouterToNetworkBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMICouldBeRouterToNetworkBuilder().(*_NLMICouldBeRouterToNetworkBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMRejectMessageToNetwork() interface {
	NLMRejectMessageToNetworkBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMRejectMessageToNetworkBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMRejectMessageToNetworkBuilder().(*_NLMRejectMessageToNetworkBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMRouterBusyToNetwork() interface {
	NLMRouterBusyToNetworkBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMRouterBusyToNetworkBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMRouterBusyToNetworkBuilder().(*_NLMRouterBusyToNetworkBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMRouterAvailableToNetwork() interface {
	NLMRouterAvailableToNetworkBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMRouterAvailableToNetworkBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMRouterAvailableToNetworkBuilder().(*_NLMRouterAvailableToNetworkBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMInitializeRoutingTable() interface {
	NLMInitializeRoutingTableBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMInitializeRoutingTableBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMInitializeRoutingTableBuilder().(*_NLMInitializeRoutingTableBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMInitializeRoutingTableAck() interface {
	NLMInitializeRoutingTableAckBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMInitializeRoutingTableAckBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMInitializeRoutingTableAckBuilder().(*_NLMInitializeRoutingTableAckBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMEstablishConnectionToNetwork() interface {
	NLMEstablishConnectionToNetworkBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMEstablishConnectionToNetworkBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMEstablishConnectionToNetworkBuilder().(*_NLMEstablishConnectionToNetworkBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMDisconnectConnectionToNetwork() interface {
	NLMDisconnectConnectionToNetworkBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMDisconnectConnectionToNetworkBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMDisconnectConnectionToNetworkBuilder().(*_NLMDisconnectConnectionToNetworkBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMChallengeRequest() interface {
	NLMChallengeRequestBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMChallengeRequestBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMChallengeRequestBuilder().(*_NLMChallengeRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMSecurityPayload() interface {
	NLMSecurityPayloadBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMSecurityPayloadBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMSecurityPayloadBuilder().(*_NLMSecurityPayloadBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMSecurityResponse() interface {
	NLMSecurityResponseBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMSecurityResponseBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMSecurityResponseBuilder().(*_NLMSecurityResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMRequestKeyUpdate() interface {
	NLMRequestKeyUpdateBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMRequestKeyUpdateBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMRequestKeyUpdateBuilder().(*_NLMRequestKeyUpdateBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMUpdateKeyUpdate() interface {
	NLMUpdateKeyUpdateBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMUpdateKeyUpdateBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMUpdateKeyUpdateBuilder().(*_NLMUpdateKeyUpdateBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMUpdateKeyDistributionKey() interface {
	NLMUpdateKeyDistributionKeyBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMUpdateKeyDistributionKeyBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMUpdateKeyDistributionKeyBuilder().(*_NLMUpdateKeyDistributionKeyBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMRequestMasterKey() interface {
	NLMRequestMasterKeyBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMRequestMasterKeyBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMRequestMasterKeyBuilder().(*_NLMRequestMasterKeyBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMSetMasterKey() interface {
	NLMSetMasterKeyBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMSetMasterKeyBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMSetMasterKeyBuilder().(*_NLMSetMasterKeyBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMWhatIsNetworkNumber() interface {
	NLMWhatIsNetworkNumberBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMWhatIsNetworkNumberBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMWhatIsNetworkNumberBuilder().(*_NLMWhatIsNetworkNumberBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMNetworkNumberIs() interface {
	NLMNetworkNumberIsBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMNetworkNumberIsBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMNetworkNumberIsBuilder().(*_NLMNetworkNumberIsBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMReserved() interface {
	NLMReservedBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMReservedBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMReservedBuilder().(*_NLMReservedBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) AsNLMVendorProprietaryMessage() interface {
	NLMVendorProprietaryMessageBuilder
	Done() NLMBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		NLMVendorProprietaryMessageBuilder
		Done() NLMBuilder
	}); ok {
		return cb
	}
	cb := NewNLMVendorProprietaryMessageBuilder().(*_NLMVendorProprietaryMessageBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NLMBuilder) Build() (NLM, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForNLM()
}

func (b *_NLMBuilder) MustBuild() NLM {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NLMBuilder) DeepCopy() any {
	_copy := b.CreateNLMBuilder().(*_NLMBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_NLMChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNLMBuilder creates a NLMBuilder
func (b *_NLM) CreateNLMBuilder() NLMBuilder {
	if b == nil {
		return NewNLMBuilder()
	}
	return &_NLMBuilder{_NLM: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_NLM) GetIsVendorProprietaryMessage() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMessageType()) >= (128)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLM(structType any) NLM {
	if casted, ok := structType.(NLM); ok {
		return casted
	}
	if casted, ok := structType.(*NLM); ok {
		return *casted
	}
	return nil
}

func (m *_NLM) GetTypeName() string {
	return "NLM"
}

func (m *_NLM) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NLM) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func NLMParse[T NLM](ctx context.Context, theBytes []byte, apduLength uint16) (T, error) {
	return NLMParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMParseWithBufferProducer[T NLM](apduLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := NLMParseWithBuffer[T](ctx, readBuffer, apduLength)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func NLMParseWithBuffer[T NLM](ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (T, error) {
	v, err := (&_NLM{ApduLength: apduLength}).parse(ctx, readBuffer, apduLength)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_NLM) parse(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (__nLM NLM, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLM"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLM")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	messageType, err := ReadDiscriminatorField[uint8](ctx, "messageType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageType' field"))
	}

	isVendorProprietaryMessage, err := ReadVirtualField[bool](ctx, "isVendorProprietaryMessage", (*bool)(nil), bool((messageType) >= (128)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isVendorProprietaryMessage' field"))
	}
	_ = isVendorProprietaryMessage

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child NLM
	switch {
	case messageType == 0x00: // NLMWhoIsRouterToNetwork
		if _child, err = new(_NLMWhoIsRouterToNetwork).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMWhoIsRouterToNetwork for type-switch of NLM")
		}
	case messageType == 0x01: // NLMIAmRouterToNetwork
		if _child, err = new(_NLMIAmRouterToNetwork).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMIAmRouterToNetwork for type-switch of NLM")
		}
	case messageType == 0x02: // NLMICouldBeRouterToNetwork
		if _child, err = new(_NLMICouldBeRouterToNetwork).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMICouldBeRouterToNetwork for type-switch of NLM")
		}
	case messageType == 0x03: // NLMRejectMessageToNetwork
		if _child, err = new(_NLMRejectMessageToNetwork).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMRejectMessageToNetwork for type-switch of NLM")
		}
	case messageType == 0x04: // NLMRouterBusyToNetwork
		if _child, err = new(_NLMRouterBusyToNetwork).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMRouterBusyToNetwork for type-switch of NLM")
		}
	case messageType == 0x05: // NLMRouterAvailableToNetwork
		if _child, err = new(_NLMRouterAvailableToNetwork).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMRouterAvailableToNetwork for type-switch of NLM")
		}
	case messageType == 0x06: // NLMInitializeRoutingTable
		if _child, err = new(_NLMInitializeRoutingTable).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMInitializeRoutingTable for type-switch of NLM")
		}
	case messageType == 0x07: // NLMInitializeRoutingTableAck
		if _child, err = new(_NLMInitializeRoutingTableAck).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMInitializeRoutingTableAck for type-switch of NLM")
		}
	case messageType == 0x08: // NLMEstablishConnectionToNetwork
		if _child, err = new(_NLMEstablishConnectionToNetwork).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMEstablishConnectionToNetwork for type-switch of NLM")
		}
	case messageType == 0x09: // NLMDisconnectConnectionToNetwork
		if _child, err = new(_NLMDisconnectConnectionToNetwork).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMDisconnectConnectionToNetwork for type-switch of NLM")
		}
	case messageType == 0x0A: // NLMChallengeRequest
		if _child, err = new(_NLMChallengeRequest).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMChallengeRequest for type-switch of NLM")
		}
	case messageType == 0x0B: // NLMSecurityPayload
		if _child, err = new(_NLMSecurityPayload).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMSecurityPayload for type-switch of NLM")
		}
	case messageType == 0x0C: // NLMSecurityResponse
		if _child, err = new(_NLMSecurityResponse).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMSecurityResponse for type-switch of NLM")
		}
	case messageType == 0x0D: // NLMRequestKeyUpdate
		if _child, err = new(_NLMRequestKeyUpdate).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMRequestKeyUpdate for type-switch of NLM")
		}
	case messageType == 0x0E: // NLMUpdateKeyUpdate
		if _child, err = new(_NLMUpdateKeyUpdate).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMUpdateKeyUpdate for type-switch of NLM")
		}
	case messageType == 0x0F: // NLMUpdateKeyDistributionKey
		if _child, err = new(_NLMUpdateKeyDistributionKey).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMUpdateKeyDistributionKey for type-switch of NLM")
		}
	case messageType == 0x10: // NLMRequestMasterKey
		if _child, err = new(_NLMRequestMasterKey).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMRequestMasterKey for type-switch of NLM")
		}
	case messageType == 0x11: // NLMSetMasterKey
		if _child, err = new(_NLMSetMasterKey).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMSetMasterKey for type-switch of NLM")
		}
	case messageType == 0x12: // NLMWhatIsNetworkNumber
		if _child, err = new(_NLMWhatIsNetworkNumber).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMWhatIsNetworkNumber for type-switch of NLM")
		}
	case messageType == 0x13: // NLMNetworkNumberIs
		if _child, err = new(_NLMNetworkNumberIs).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMNetworkNumberIs for type-switch of NLM")
		}
	case 0 == 0 && isVendorProprietaryMessage == bool(false): // NLMReserved
		if _child, err = new(_NLMReserved).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMReserved for type-switch of NLM")
		}
	case 0 == 0: // NLMVendorProprietaryMessage
		if _child, err = new(_NLMVendorProprietaryMessage).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMVendorProprietaryMessage for type-switch of NLM")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [messageType=%v, isVendorProprietaryMessage=%v]", messageType, isVendorProprietaryMessage)
	}

	if closeErr := readBuffer.CloseContext("NLM"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLM")
	}

	return _child, nil
}

func (pm *_NLM) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child NLM, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NLM"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NLM")
	}

	if err := WriteDiscriminatorField(ctx, "messageType", m.GetMessageType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageType' field")
	}
	// Virtual field
	isVendorProprietaryMessage := m.GetIsVendorProprietaryMessage()
	_ = isVendorProprietaryMessage
	if _isVendorProprietaryMessageErr := writeBuffer.WriteVirtual(ctx, "isVendorProprietaryMessage", m.GetIsVendorProprietaryMessage()); _isVendorProprietaryMessageErr != nil {
		return errors.Wrap(_isVendorProprietaryMessageErr, "Error serializing 'isVendorProprietaryMessage' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("NLM"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NLM")
	}
	return nil
}

////
// Arguments Getter

func (m *_NLM) GetApduLength() uint16 {
	return m.ApduLength
}

//
////

func (m *_NLM) IsNLM() {}

func (m *_NLM) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLM) deepCopy() *_NLM {
	if m == nil {
		return nil
	}
	_NLMCopy := &_NLM{
		nil, // will be set by child
		m.ApduLength,
	}
	return _NLMCopy
}

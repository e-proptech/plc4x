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

// BACnetServiceAckRequestKey is the corresponding interface of BACnetServiceAckRequestKey
type BACnetServiceAckRequestKey interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetServiceAck
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
	// IsBACnetServiceAckRequestKey is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckRequestKey()
	// CreateBuilder creates a BACnetServiceAckRequestKeyBuilder
	CreateBACnetServiceAckRequestKeyBuilder() BACnetServiceAckRequestKeyBuilder
}

// _BACnetServiceAckRequestKey is the data-structure of this message
type _BACnetServiceAckRequestKey struct {
	BACnetServiceAckContract
	BytesOfRemovedService []byte

	// Arguments.
	ServiceAckPayloadLength uint32
}

var _ BACnetServiceAckRequestKey = (*_BACnetServiceAckRequestKey)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckRequestKey)(nil)

// NewBACnetServiceAckRequestKey factory function for _BACnetServiceAckRequestKey
func NewBACnetServiceAckRequestKey(bytesOfRemovedService []byte, serviceAckPayloadLength uint32, serviceAckLength uint32) *_BACnetServiceAckRequestKey {
	_result := &_BACnetServiceAckRequestKey{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		BytesOfRemovedService:    bytesOfRemovedService,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetServiceAckRequestKeyBuilder is a builder for BACnetServiceAckRequestKey
type BACnetServiceAckRequestKeyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bytesOfRemovedService []byte) BACnetServiceAckRequestKeyBuilder
	// WithBytesOfRemovedService adds BytesOfRemovedService (property field)
	WithBytesOfRemovedService(...byte) BACnetServiceAckRequestKeyBuilder
	// Build builds the BACnetServiceAckRequestKey or returns an error if something is wrong
	Build() (BACnetServiceAckRequestKey, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetServiceAckRequestKey
}

// NewBACnetServiceAckRequestKeyBuilder() creates a BACnetServiceAckRequestKeyBuilder
func NewBACnetServiceAckRequestKeyBuilder() BACnetServiceAckRequestKeyBuilder {
	return &_BACnetServiceAckRequestKeyBuilder{_BACnetServiceAckRequestKey: new(_BACnetServiceAckRequestKey)}
}

type _BACnetServiceAckRequestKeyBuilder struct {
	*_BACnetServiceAckRequestKey

	parentBuilder *_BACnetServiceAckBuilder

	err *utils.MultiError
}

var _ (BACnetServiceAckRequestKeyBuilder) = (*_BACnetServiceAckRequestKeyBuilder)(nil)

func (b *_BACnetServiceAckRequestKeyBuilder) setParent(contract BACnetServiceAckContract) {
	b.BACnetServiceAckContract = contract
}

func (b *_BACnetServiceAckRequestKeyBuilder) WithMandatoryFields(bytesOfRemovedService []byte) BACnetServiceAckRequestKeyBuilder {
	return b.WithBytesOfRemovedService(bytesOfRemovedService...)
}

func (b *_BACnetServiceAckRequestKeyBuilder) WithBytesOfRemovedService(bytesOfRemovedService ...byte) BACnetServiceAckRequestKeyBuilder {
	b.BytesOfRemovedService = bytesOfRemovedService
	return b
}

func (b *_BACnetServiceAckRequestKeyBuilder) Build() (BACnetServiceAckRequestKey, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetServiceAckRequestKey.deepCopy(), nil
}

func (b *_BACnetServiceAckRequestKeyBuilder) MustBuild() BACnetServiceAckRequestKey {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetServiceAckRequestKeyBuilder) Done() BACnetServiceAckBuilder {
	return b.parentBuilder
}

func (b *_BACnetServiceAckRequestKeyBuilder) buildForBACnetServiceAck() (BACnetServiceAck, error) {
	return b.Build()
}

func (b *_BACnetServiceAckRequestKeyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetServiceAckRequestKeyBuilder().(*_BACnetServiceAckRequestKeyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetServiceAckRequestKeyBuilder creates a BACnetServiceAckRequestKeyBuilder
func (b *_BACnetServiceAckRequestKey) CreateBACnetServiceAckRequestKeyBuilder() BACnetServiceAckRequestKeyBuilder {
	if b == nil {
		return NewBACnetServiceAckRequestKeyBuilder()
	}
	return &_BACnetServiceAckRequestKeyBuilder{_BACnetServiceAckRequestKey: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckRequestKey) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REQUEST_KEY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckRequestKey) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckRequestKey) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckRequestKey(structType any) BACnetServiceAckRequestKey {
	if casted, ok := structType.(BACnetServiceAckRequestKey); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckRequestKey); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckRequestKey) GetTypeName() string {
	return "BACnetServiceAckRequestKey"
}

func (m *_BACnetServiceAckRequestKey) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}

func (m *_BACnetServiceAckRequestKey) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckRequestKey) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckPayloadLength uint32, serviceAckLength uint32) (__bACnetServiceAckRequestKey BACnetServiceAckRequestKey, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckRequestKey"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckRequestKey")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bytesOfRemovedService, err := readBuffer.ReadByteArray("bytesOfRemovedService", int(serviceAckPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bytesOfRemovedService' field"))
	}
	m.BytesOfRemovedService = bytesOfRemovedService

	if closeErr := readBuffer.CloseContext("BACnetServiceAckRequestKey"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckRequestKey")
	}

	return m, nil
}

func (m *_BACnetServiceAckRequestKey) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckRequestKey) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckRequestKey"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckRequestKey")
		}

		if err := WriteByteArrayField(ctx, "bytesOfRemovedService", m.GetBytesOfRemovedService(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'bytesOfRemovedService' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckRequestKey"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckRequestKey")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetServiceAckRequestKey) GetServiceAckPayloadLength() uint32 {
	return m.ServiceAckPayloadLength
}

//
////

func (m *_BACnetServiceAckRequestKey) IsBACnetServiceAckRequestKey() {}

func (m *_BACnetServiceAckRequestKey) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServiceAckRequestKey) deepCopy() *_BACnetServiceAckRequestKey {
	if m == nil {
		return nil
	}
	_BACnetServiceAckRequestKeyCopy := &_BACnetServiceAckRequestKey{
		m.BACnetServiceAckContract.(*_BACnetServiceAck).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.BytesOfRemovedService),
		m.ServiceAckPayloadLength,
	}
	m.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = m
	return _BACnetServiceAckRequestKeyCopy
}

func (m *_BACnetServiceAckRequestKey) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

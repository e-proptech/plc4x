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

// BACnetContextTagTime is the corresponding interface of BACnetContextTagTime
type BACnetContextTagTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadTime
	// IsBACnetContextTagTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagTime()
	// CreateBuilder creates a BACnetContextTagTimeBuilder
	CreateBACnetContextTagTimeBuilder() BACnetContextTagTimeBuilder
}

// _BACnetContextTagTime is the data-structure of this message
type _BACnetContextTagTime struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadTime
}

var _ BACnetContextTagTime = (*_BACnetContextTagTime)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagTime)(nil)

// NewBACnetContextTagTime factory function for _BACnetContextTagTime
func NewBACnetContextTagTime(header BACnetTagHeader, payload BACnetTagPayloadTime, tagNumberArgument uint8) *_BACnetContextTagTime {
	if payload == nil {
		panic("payload of type BACnetTagPayloadTime for BACnetContextTagTime must not be nil")
	}
	_result := &_BACnetContextTagTime{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetContextTagTimeBuilder is a builder for BACnetContextTagTime
type BACnetContextTagTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadTime) BACnetContextTagTimeBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadTime) BACnetContextTagTimeBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadTimeBuilder) BACnetTagPayloadTimeBuilder) BACnetContextTagTimeBuilder
	// Build builds the BACnetContextTagTime or returns an error if something is wrong
	Build() (BACnetContextTagTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagTime
}

// NewBACnetContextTagTimeBuilder() creates a BACnetContextTagTimeBuilder
func NewBACnetContextTagTimeBuilder() BACnetContextTagTimeBuilder {
	return &_BACnetContextTagTimeBuilder{_BACnetContextTagTime: new(_BACnetContextTagTime)}
}

type _BACnetContextTagTimeBuilder struct {
	*_BACnetContextTagTime

	parentBuilder *_BACnetContextTagBuilder

	err *utils.MultiError
}

var _ (BACnetContextTagTimeBuilder) = (*_BACnetContextTagTimeBuilder)(nil)

func (b *_BACnetContextTagTimeBuilder) setParent(contract BACnetContextTagContract) {
	b.BACnetContextTagContract = contract
}

func (b *_BACnetContextTagTimeBuilder) WithMandatoryFields(payload BACnetTagPayloadTime) BACnetContextTagTimeBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetContextTagTimeBuilder) WithPayload(payload BACnetTagPayloadTime) BACnetContextTagTimeBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetContextTagTimeBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadTimeBuilder) BACnetTagPayloadTimeBuilder) BACnetContextTagTimeBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadTimeBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetContextTagTimeBuilder) Build() (BACnetContextTagTime, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetContextTagTime.deepCopy(), nil
}

func (b *_BACnetContextTagTimeBuilder) MustBuild() BACnetContextTagTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetContextTagTimeBuilder) Done() BACnetContextTagBuilder {
	return b.parentBuilder
}

func (b *_BACnetContextTagTimeBuilder) buildForBACnetContextTag() (BACnetContextTag, error) {
	return b.Build()
}

func (b *_BACnetContextTagTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetContextTagTimeBuilder().(*_BACnetContextTagTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetContextTagTimeBuilder creates a BACnetContextTagTimeBuilder
func (b *_BACnetContextTagTime) CreateBACnetContextTagTimeBuilder() BACnetContextTagTimeBuilder {
	if b == nil {
		return NewBACnetContextTagTimeBuilder()
	}
	return &_BACnetContextTagTimeBuilder{_BACnetContextTagTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagTime) GetDataType() BACnetDataType {
	return BACnetDataType_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagTime) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagTime) GetPayload() BACnetTagPayloadTime {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTagTime(structType any) BACnetContextTagTime {
	if casted, ok := structType.(BACnetContextTagTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagTime) GetTypeName() string {
	return "BACnetContextTagTime"
}

func (m *_BACnetContextTagTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetContextTagTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagTime BACnetContextTagTime, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadTime](ctx, "payload", ReadComplex[BACnetTagPayloadTime](BACnetTagPayloadTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("BACnetContextTagTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagTime")
	}

	return m, nil
}

func (m *_BACnetContextTagTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagTime")
		}

		if err := WriteSimpleField[BACnetTagPayloadTime](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagTime")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagTime) IsBACnetContextTagTime() {}

func (m *_BACnetContextTagTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTagTime) deepCopy() *_BACnetContextTagTime {
	if m == nil {
		return nil
	}
	_BACnetContextTagTimeCopy := &_BACnetContextTagTime{
		m.BACnetContextTagContract.(*_BACnetContextTag).deepCopy(),
		utils.DeepCopy[BACnetTagPayloadTime](m.Payload),
	}
	m.BACnetContextTagContract.(*_BACnetContextTag)._SubType = m
	return _BACnetContextTagTimeCopy
}

func (m *_BACnetContextTagTime) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}

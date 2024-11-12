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

// BACnetApplicationTagOctetString is the corresponding interface of BACnetApplicationTagOctetString
type BACnetApplicationTagOctetString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadOctetString
	// IsBACnetApplicationTagOctetString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagOctetString()
	// CreateBuilder creates a BACnetApplicationTagOctetStringBuilder
	CreateBACnetApplicationTagOctetStringBuilder() BACnetApplicationTagOctetStringBuilder
}

// _BACnetApplicationTagOctetString is the data-structure of this message
type _BACnetApplicationTagOctetString struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadOctetString
}

var _ BACnetApplicationTagOctetString = (*_BACnetApplicationTagOctetString)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagOctetString)(nil)

// NewBACnetApplicationTagOctetString factory function for _BACnetApplicationTagOctetString
func NewBACnetApplicationTagOctetString(header BACnetTagHeader, payload BACnetTagPayloadOctetString) *_BACnetApplicationTagOctetString {
	if payload == nil {
		panic("payload of type BACnetTagPayloadOctetString for BACnetApplicationTagOctetString must not be nil")
	}
	_result := &_BACnetApplicationTagOctetString{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetApplicationTagOctetStringBuilder is a builder for BACnetApplicationTagOctetString
type BACnetApplicationTagOctetStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadOctetString) BACnetApplicationTagOctetStringBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadOctetString) BACnetApplicationTagOctetStringBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadOctetStringBuilder) BACnetTagPayloadOctetStringBuilder) BACnetApplicationTagOctetStringBuilder
	// Build builds the BACnetApplicationTagOctetString or returns an error if something is wrong
	Build() (BACnetApplicationTagOctetString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetApplicationTagOctetString
}

// NewBACnetApplicationTagOctetStringBuilder() creates a BACnetApplicationTagOctetStringBuilder
func NewBACnetApplicationTagOctetStringBuilder() BACnetApplicationTagOctetStringBuilder {
	return &_BACnetApplicationTagOctetStringBuilder{_BACnetApplicationTagOctetString: new(_BACnetApplicationTagOctetString)}
}

type _BACnetApplicationTagOctetStringBuilder struct {
	*_BACnetApplicationTagOctetString

	parentBuilder *_BACnetApplicationTagBuilder

	err *utils.MultiError
}

var _ (BACnetApplicationTagOctetStringBuilder) = (*_BACnetApplicationTagOctetStringBuilder)(nil)

func (b *_BACnetApplicationTagOctetStringBuilder) setParent(contract BACnetApplicationTagContract) {
	b.BACnetApplicationTagContract = contract
}

func (b *_BACnetApplicationTagOctetStringBuilder) WithMandatoryFields(payload BACnetTagPayloadOctetString) BACnetApplicationTagOctetStringBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetApplicationTagOctetStringBuilder) WithPayload(payload BACnetTagPayloadOctetString) BACnetApplicationTagOctetStringBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetApplicationTagOctetStringBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadOctetStringBuilder) BACnetTagPayloadOctetStringBuilder) BACnetApplicationTagOctetStringBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadOctetStringBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetApplicationTagOctetStringBuilder) Build() (BACnetApplicationTagOctetString, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetApplicationTagOctetString.deepCopy(), nil
}

func (b *_BACnetApplicationTagOctetStringBuilder) MustBuild() BACnetApplicationTagOctetString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetApplicationTagOctetStringBuilder) Done() BACnetApplicationTagBuilder {
	return b.parentBuilder
}

func (b *_BACnetApplicationTagOctetStringBuilder) buildForBACnetApplicationTag() (BACnetApplicationTag, error) {
	return b.Build()
}

func (b *_BACnetApplicationTagOctetStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetApplicationTagOctetStringBuilder().(*_BACnetApplicationTagOctetStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetApplicationTagOctetStringBuilder creates a BACnetApplicationTagOctetStringBuilder
func (b *_BACnetApplicationTagOctetString) CreateBACnetApplicationTagOctetStringBuilder() BACnetApplicationTagOctetStringBuilder {
	if b == nil {
		return NewBACnetApplicationTagOctetStringBuilder()
	}
	return &_BACnetApplicationTagOctetStringBuilder{_BACnetApplicationTagOctetString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagOctetString) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagOctetString) GetPayload() BACnetTagPayloadOctetString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagOctetString(structType any) BACnetApplicationTagOctetString {
	if casted, ok := structType.(BACnetApplicationTagOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagOctetString) GetTypeName() string {
	return "BACnetApplicationTagOctetString"
}

func (m *_BACnetApplicationTagOctetString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetApplicationTagOctetString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagOctetString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag, header BACnetTagHeader) (__bACnetApplicationTagOctetString BACnetApplicationTagOctetString, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadOctetString](ctx, "payload", ReadComplex[BACnetTagPayloadOctetString](BACnetTagPayloadOctetStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagOctetString")
	}

	return m, nil
}

func (m *_BACnetApplicationTagOctetString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagOctetString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagOctetString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagOctetString")
		}

		if err := WriteSimpleField[BACnetTagPayloadOctetString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagOctetString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagOctetString")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagOctetString) IsBACnetApplicationTagOctetString() {}

func (m *_BACnetApplicationTagOctetString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetApplicationTagOctetString) deepCopy() *_BACnetApplicationTagOctetString {
	if m == nil {
		return nil
	}
	_BACnetApplicationTagOctetStringCopy := &_BACnetApplicationTagOctetString{
		m.BACnetApplicationTagContract.(*_BACnetApplicationTag).deepCopy(),
		utils.DeepCopy[BACnetTagPayloadOctetString](m.Payload),
	}
	m.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = m
	return _BACnetApplicationTagOctetStringCopy
}

func (m *_BACnetApplicationTagOctetString) String() string {
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

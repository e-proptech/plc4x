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

// BACnetApplicationTagCharacterString is the corresponding interface of BACnetApplicationTagCharacterString
type BACnetApplicationTagCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadCharacterString
	// GetValue returns Value (virtual field)
	GetValue() string
	// IsBACnetApplicationTagCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagCharacterString()
	// CreateBuilder creates a BACnetApplicationTagCharacterStringBuilder
	CreateBACnetApplicationTagCharacterStringBuilder() BACnetApplicationTagCharacterStringBuilder
}

// _BACnetApplicationTagCharacterString is the data-structure of this message
type _BACnetApplicationTagCharacterString struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadCharacterString
}

var _ BACnetApplicationTagCharacterString = (*_BACnetApplicationTagCharacterString)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagCharacterString)(nil)

// NewBACnetApplicationTagCharacterString factory function for _BACnetApplicationTagCharacterString
func NewBACnetApplicationTagCharacterString(header BACnetTagHeader, payload BACnetTagPayloadCharacterString) *_BACnetApplicationTagCharacterString {
	if payload == nil {
		panic("payload of type BACnetTagPayloadCharacterString for BACnetApplicationTagCharacterString must not be nil")
	}
	_result := &_BACnetApplicationTagCharacterString{
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

// BACnetApplicationTagCharacterStringBuilder is a builder for BACnetApplicationTagCharacterString
type BACnetApplicationTagCharacterStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadCharacterString) BACnetApplicationTagCharacterStringBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadCharacterString) BACnetApplicationTagCharacterStringBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadCharacterStringBuilder) BACnetTagPayloadCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder
	// Build builds the BACnetApplicationTagCharacterString or returns an error if something is wrong
	Build() (BACnetApplicationTagCharacterString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetApplicationTagCharacterString
}

// NewBACnetApplicationTagCharacterStringBuilder() creates a BACnetApplicationTagCharacterStringBuilder
func NewBACnetApplicationTagCharacterStringBuilder() BACnetApplicationTagCharacterStringBuilder {
	return &_BACnetApplicationTagCharacterStringBuilder{_BACnetApplicationTagCharacterString: new(_BACnetApplicationTagCharacterString)}
}

type _BACnetApplicationTagCharacterStringBuilder struct {
	*_BACnetApplicationTagCharacterString

	parentBuilder *_BACnetApplicationTagBuilder

	err *utils.MultiError
}

var _ (BACnetApplicationTagCharacterStringBuilder) = (*_BACnetApplicationTagCharacterStringBuilder)(nil)

func (b *_BACnetApplicationTagCharacterStringBuilder) setParent(contract BACnetApplicationTagContract) {
	b.BACnetApplicationTagContract = contract
}

func (b *_BACnetApplicationTagCharacterStringBuilder) WithMandatoryFields(payload BACnetTagPayloadCharacterString) BACnetApplicationTagCharacterStringBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetApplicationTagCharacterStringBuilder) WithPayload(payload BACnetTagPayloadCharacterString) BACnetApplicationTagCharacterStringBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetApplicationTagCharacterStringBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadCharacterStringBuilder) BACnetTagPayloadCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadCharacterStringBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetApplicationTagCharacterStringBuilder) Build() (BACnetApplicationTagCharacterString, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetApplicationTagCharacterString.deepCopy(), nil
}

func (b *_BACnetApplicationTagCharacterStringBuilder) MustBuild() BACnetApplicationTagCharacterString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetApplicationTagCharacterStringBuilder) Done() BACnetApplicationTagBuilder {
	return b.parentBuilder
}

func (b *_BACnetApplicationTagCharacterStringBuilder) buildForBACnetApplicationTag() (BACnetApplicationTag, error) {
	return b.Build()
}

func (b *_BACnetApplicationTagCharacterStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetApplicationTagCharacterStringBuilder().(*_BACnetApplicationTagCharacterStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetApplicationTagCharacterStringBuilder creates a BACnetApplicationTagCharacterStringBuilder
func (b *_BACnetApplicationTagCharacterString) CreateBACnetApplicationTagCharacterStringBuilder() BACnetApplicationTagCharacterStringBuilder {
	if b == nil {
		return NewBACnetApplicationTagCharacterStringBuilder()
	}
	return &_BACnetApplicationTagCharacterStringBuilder{_BACnetApplicationTagCharacterString: b.deepCopy()}
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

func (m *_BACnetApplicationTagCharacterString) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagCharacterString) GetPayload() BACnetTagPayloadCharacterString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagCharacterString) GetValue() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagCharacterString(structType any) BACnetApplicationTagCharacterString {
	if casted, ok := structType.(BACnetApplicationTagCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagCharacterString) GetTypeName() string {
	return "BACnetApplicationTagCharacterString"
}

func (m *_BACnetApplicationTagCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag, header BACnetTagHeader) (__bACnetApplicationTagCharacterString BACnetApplicationTagCharacterString, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadCharacterString](ctx, "payload", ReadComplex[BACnetTagPayloadCharacterString](BACnetTagPayloadCharacterStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	value, err := ReadVirtualField[string](ctx, "value", (*string)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	_ = value

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagCharacterString")
	}

	return m, nil
}

func (m *_BACnetApplicationTagCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagCharacterString")
		}

		if err := WriteSimpleField[BACnetTagPayloadCharacterString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		value := m.GetValue()
		_ = value
		if _valueErr := writeBuffer.WriteVirtual(ctx, "value", m.GetValue()); _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagCharacterString")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagCharacterString) IsBACnetApplicationTagCharacterString() {}

func (m *_BACnetApplicationTagCharacterString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetApplicationTagCharacterString) deepCopy() *_BACnetApplicationTagCharacterString {
	if m == nil {
		return nil
	}
	_BACnetApplicationTagCharacterStringCopy := &_BACnetApplicationTagCharacterString{
		m.BACnetApplicationTagContract.(*_BACnetApplicationTag).deepCopy(),
		utils.DeepCopy[BACnetTagPayloadCharacterString](m.Payload),
	}
	m.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = m
	return _BACnetApplicationTagCharacterStringCopy
}

func (m *_BACnetApplicationTagCharacterString) String() string {
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

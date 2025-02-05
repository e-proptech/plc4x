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

// BACnetContextTagBoolean is the corresponding interface of BACnetContextTagBoolean
type BACnetContextTagBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetContextTag
	// GetValue returns Value (property field)
	GetValue() uint8
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() bool
	// IsBACnetContextTagBoolean is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagBoolean()
	// CreateBuilder creates a BACnetContextTagBooleanBuilder
	CreateBACnetContextTagBooleanBuilder() BACnetContextTagBooleanBuilder
}

// _BACnetContextTagBoolean is the data-structure of this message
type _BACnetContextTagBoolean struct {
	BACnetContextTagContract
	Value   uint8
	Payload BACnetTagPayloadBoolean
}

var _ BACnetContextTagBoolean = (*_BACnetContextTagBoolean)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagBoolean)(nil)

// NewBACnetContextTagBoolean factory function for _BACnetContextTagBoolean
func NewBACnetContextTagBoolean(header BACnetTagHeader, value uint8, payload BACnetTagPayloadBoolean, tagNumberArgument uint8) *_BACnetContextTagBoolean {
	if payload == nil {
		panic("payload of type BACnetTagPayloadBoolean for BACnetContextTagBoolean must not be nil")
	}
	_result := &_BACnetContextTagBoolean{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Value:                    value,
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetContextTagBooleanBuilder is a builder for BACnetContextTagBoolean
type BACnetContextTagBooleanBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value uint8, payload BACnetTagPayloadBoolean) BACnetContextTagBooleanBuilder
	// WithValue adds Value (property field)
	WithValue(uint8) BACnetContextTagBooleanBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadBoolean) BACnetContextTagBooleanBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadBooleanBuilder) BACnetTagPayloadBooleanBuilder) BACnetContextTagBooleanBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetContextTagBuilder
	// Build builds the BACnetContextTagBoolean or returns an error if something is wrong
	Build() (BACnetContextTagBoolean, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagBoolean
}

// NewBACnetContextTagBooleanBuilder() creates a BACnetContextTagBooleanBuilder
func NewBACnetContextTagBooleanBuilder() BACnetContextTagBooleanBuilder {
	return &_BACnetContextTagBooleanBuilder{_BACnetContextTagBoolean: new(_BACnetContextTagBoolean)}
}

type _BACnetContextTagBooleanBuilder struct {
	*_BACnetContextTagBoolean

	parentBuilder *_BACnetContextTagBuilder

	err *utils.MultiError
}

var _ (BACnetContextTagBooleanBuilder) = (*_BACnetContextTagBooleanBuilder)(nil)

func (b *_BACnetContextTagBooleanBuilder) setParent(contract BACnetContextTagContract) {
	b.BACnetContextTagContract = contract
	contract.(*_BACnetContextTag)._SubType = b._BACnetContextTagBoolean
}

func (b *_BACnetContextTagBooleanBuilder) WithMandatoryFields(value uint8, payload BACnetTagPayloadBoolean) BACnetContextTagBooleanBuilder {
	return b.WithValue(value).WithPayload(payload)
}

func (b *_BACnetContextTagBooleanBuilder) WithValue(value uint8) BACnetContextTagBooleanBuilder {
	b.Value = value
	return b
}

func (b *_BACnetContextTagBooleanBuilder) WithPayload(payload BACnetTagPayloadBoolean) BACnetContextTagBooleanBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetContextTagBooleanBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadBooleanBuilder) BACnetTagPayloadBooleanBuilder) BACnetContextTagBooleanBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadBooleanBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetContextTagBooleanBuilder) Build() (BACnetContextTagBoolean, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetContextTagBoolean.deepCopy(), nil
}

func (b *_BACnetContextTagBooleanBuilder) MustBuild() BACnetContextTagBoolean {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetContextTagBooleanBuilder) Done() BACnetContextTagBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetContextTagBuilder().(*_BACnetContextTagBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetContextTagBooleanBuilder) buildForBACnetContextTag() (BACnetContextTag, error) {
	return b.Build()
}

func (b *_BACnetContextTagBooleanBuilder) DeepCopy() any {
	_copy := b.CreateBACnetContextTagBooleanBuilder().(*_BACnetContextTagBooleanBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetContextTagBooleanBuilder creates a BACnetContextTagBooleanBuilder
func (b *_BACnetContextTagBoolean) CreateBACnetContextTagBooleanBuilder() BACnetContextTagBooleanBuilder {
	if b == nil {
		return NewBACnetContextTagBooleanBuilder()
	}
	return &_BACnetContextTagBooleanBuilder{_BACnetContextTagBoolean: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagBoolean) GetDataType() BACnetDataType {
	return BACnetDataType_BOOLEAN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagBoolean) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagBoolean) GetValue() uint8 {
	return m.Value
}

func (m *_BACnetContextTagBoolean) GetPayload() BACnetTagPayloadBoolean {
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

func (m *_BACnetContextTagBoolean) GetActualValue() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTagBoolean(structType any) BACnetContextTagBoolean {
	if casted, ok := structType.(BACnetContextTagBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagBoolean) GetTypeName() string {
	return "BACnetContextTagBoolean"
}

func (m *_BACnetContextTagBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += 8

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagBoolean) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagBoolean BACnetContextTagBoolean, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((header.GetActualLength()) == (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "length field should be 1"})
	}

	value, err := ReadSimpleField(ctx, "value", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	payload, err := ReadSimpleField[BACnetTagPayloadBoolean](ctx, "payload", ReadComplex[BACnetTagPayloadBoolean](BACnetTagPayloadBooleanParseWithBufferProducer((uint32)(value)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[bool](ctx, "actualValue", (*bool)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagBoolean")
	}

	return m, nil
}

func (m *_BACnetContextTagBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagBoolean")
		}

		if err := WriteSimpleField[uint8](ctx, "value", m.GetValue(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteSimpleField[BACnetTagPayloadBoolean](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagBoolean")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagBoolean) IsBACnetContextTagBoolean() {}

func (m *_BACnetContextTagBoolean) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTagBoolean) deepCopy() *_BACnetContextTagBoolean {
	if m == nil {
		return nil
	}
	_BACnetContextTagBooleanCopy := &_BACnetContextTagBoolean{
		m.BACnetContextTagContract.(*_BACnetContextTag).deepCopy(),
		m.Value,
		utils.DeepCopy[BACnetTagPayloadBoolean](m.Payload),
	}
	_BACnetContextTagBooleanCopy.BACnetContextTagContract.(*_BACnetContextTag)._SubType = m
	return _BACnetContextTagBooleanCopy
}

func (m *_BACnetContextTagBoolean) String() string {
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

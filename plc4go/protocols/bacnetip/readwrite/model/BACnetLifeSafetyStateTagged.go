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

// BACnetLifeSafetyStateTagged is the corresponding interface of BACnetLifeSafetyStateTagged
type BACnetLifeSafetyStateTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLifeSafetyState
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetLifeSafetyStateTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLifeSafetyStateTagged()
	// CreateBuilder creates a BACnetLifeSafetyStateTaggedBuilder
	CreateBACnetLifeSafetyStateTaggedBuilder() BACnetLifeSafetyStateTaggedBuilder
}

// _BACnetLifeSafetyStateTagged is the data-structure of this message
type _BACnetLifeSafetyStateTagged struct {
	Header           BACnetTagHeader
	Value            BACnetLifeSafetyState
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetLifeSafetyStateTagged = (*_BACnetLifeSafetyStateTagged)(nil)

// NewBACnetLifeSafetyStateTagged factory function for _BACnetLifeSafetyStateTagged
func NewBACnetLifeSafetyStateTagged(header BACnetTagHeader, value BACnetLifeSafetyState, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetLifeSafetyStateTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetLifeSafetyStateTagged must not be nil")
	}
	return &_BACnetLifeSafetyStateTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLifeSafetyStateTaggedBuilder is a builder for BACnetLifeSafetyStateTagged
type BACnetLifeSafetyStateTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetLifeSafetyState, proprietaryValue uint32) BACnetLifeSafetyStateTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetLifeSafetyStateTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLifeSafetyStateTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetLifeSafetyState) BACnetLifeSafetyStateTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetLifeSafetyStateTaggedBuilder
	// Build builds the BACnetLifeSafetyStateTagged or returns an error if something is wrong
	Build() (BACnetLifeSafetyStateTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLifeSafetyStateTagged
}

// NewBACnetLifeSafetyStateTaggedBuilder() creates a BACnetLifeSafetyStateTaggedBuilder
func NewBACnetLifeSafetyStateTaggedBuilder() BACnetLifeSafetyStateTaggedBuilder {
	return &_BACnetLifeSafetyStateTaggedBuilder{_BACnetLifeSafetyStateTagged: new(_BACnetLifeSafetyStateTagged)}
}

type _BACnetLifeSafetyStateTaggedBuilder struct {
	*_BACnetLifeSafetyStateTagged

	err *utils.MultiError
}

var _ (BACnetLifeSafetyStateTaggedBuilder) = (*_BACnetLifeSafetyStateTaggedBuilder)(nil)

func (b *_BACnetLifeSafetyStateTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetLifeSafetyState, proprietaryValue uint32) BACnetLifeSafetyStateTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetLifeSafetyStateTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetLifeSafetyStateTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetLifeSafetyStateTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLifeSafetyStateTaggedBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetLifeSafetyStateTaggedBuilder) WithValue(value BACnetLifeSafetyState) BACnetLifeSafetyStateTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetLifeSafetyStateTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetLifeSafetyStateTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetLifeSafetyStateTaggedBuilder) Build() (BACnetLifeSafetyStateTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLifeSafetyStateTagged.deepCopy(), nil
}

func (b *_BACnetLifeSafetyStateTaggedBuilder) MustBuild() BACnetLifeSafetyStateTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLifeSafetyStateTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLifeSafetyStateTaggedBuilder().(*_BACnetLifeSafetyStateTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLifeSafetyStateTaggedBuilder creates a BACnetLifeSafetyStateTaggedBuilder
func (b *_BACnetLifeSafetyStateTagged) CreateBACnetLifeSafetyStateTaggedBuilder() BACnetLifeSafetyStateTaggedBuilder {
	if b == nil {
		return NewBACnetLifeSafetyStateTaggedBuilder()
	}
	return &_BACnetLifeSafetyStateTaggedBuilder{_BACnetLifeSafetyStateTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLifeSafetyStateTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLifeSafetyStateTagged) GetValue() BACnetLifeSafetyState {
	return m.Value
}

func (m *_BACnetLifeSafetyStateTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLifeSafetyStateTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetLifeSafetyState_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLifeSafetyStateTagged(structType any) BACnetLifeSafetyStateTagged {
	if casted, ok := structType.(BACnetLifeSafetyStateTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLifeSafetyStateTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLifeSafetyStateTagged) GetTypeName() string {
	return "BACnetLifeSafetyStateTagged"
}

func (m *_BACnetLifeSafetyStateTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetLifeSafetyStateTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLifeSafetyStateTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLifeSafetyStateTagged, error) {
	return BACnetLifeSafetyStateTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLifeSafetyStateTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLifeSafetyStateTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLifeSafetyStateTagged, error) {
		return BACnetLifeSafetyStateTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetLifeSafetyStateTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLifeSafetyStateTagged, error) {
	v, err := (&_BACnetLifeSafetyStateTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLifeSafetyStateTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetLifeSafetyStateTagged BACnetLifeSafetyStateTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLifeSafetyStateTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLifeSafetyStateTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetLifeSafetyState](ctx, "value", readBuffer, EnsureType[BACnetLifeSafetyState](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetLifeSafetyState_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetLifeSafetyState_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetLifeSafetyStateTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLifeSafetyStateTagged")
	}

	return m, nil
}

func (m *_BACnetLifeSafetyStateTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLifeSafetyStateTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLifeSafetyStateTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLifeSafetyStateTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetLifeSafetyState](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLifeSafetyStateTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLifeSafetyStateTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLifeSafetyStateTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLifeSafetyStateTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLifeSafetyStateTagged) IsBACnetLifeSafetyStateTagged() {}

func (m *_BACnetLifeSafetyStateTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLifeSafetyStateTagged) deepCopy() *_BACnetLifeSafetyStateTagged {
	if m == nil {
		return nil
	}
	_BACnetLifeSafetyStateTaggedCopy := &_BACnetLifeSafetyStateTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetLifeSafetyStateTaggedCopy
}

func (m *_BACnetLifeSafetyStateTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

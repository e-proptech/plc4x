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

// BACnetLiftCarModeTagged is the corresponding interface of BACnetLiftCarModeTagged
type BACnetLiftCarModeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLiftCarMode
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetLiftCarModeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLiftCarModeTagged()
	// CreateBuilder creates a BACnetLiftCarModeTaggedBuilder
	CreateBACnetLiftCarModeTaggedBuilder() BACnetLiftCarModeTaggedBuilder
}

// _BACnetLiftCarModeTagged is the data-structure of this message
type _BACnetLiftCarModeTagged struct {
	Header           BACnetTagHeader
	Value            BACnetLiftCarMode
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetLiftCarModeTagged = (*_BACnetLiftCarModeTagged)(nil)

// NewBACnetLiftCarModeTagged factory function for _BACnetLiftCarModeTagged
func NewBACnetLiftCarModeTagged(header BACnetTagHeader, value BACnetLiftCarMode, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetLiftCarModeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetLiftCarModeTagged must not be nil")
	}
	return &_BACnetLiftCarModeTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLiftCarModeTaggedBuilder is a builder for BACnetLiftCarModeTagged
type BACnetLiftCarModeTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetLiftCarMode, proprietaryValue uint32) BACnetLiftCarModeTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetLiftCarModeTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLiftCarModeTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetLiftCarMode) BACnetLiftCarModeTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetLiftCarModeTaggedBuilder
	// Build builds the BACnetLiftCarModeTagged or returns an error if something is wrong
	Build() (BACnetLiftCarModeTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLiftCarModeTagged
}

// NewBACnetLiftCarModeTaggedBuilder() creates a BACnetLiftCarModeTaggedBuilder
func NewBACnetLiftCarModeTaggedBuilder() BACnetLiftCarModeTaggedBuilder {
	return &_BACnetLiftCarModeTaggedBuilder{_BACnetLiftCarModeTagged: new(_BACnetLiftCarModeTagged)}
}

type _BACnetLiftCarModeTaggedBuilder struct {
	*_BACnetLiftCarModeTagged

	err *utils.MultiError
}

var _ (BACnetLiftCarModeTaggedBuilder) = (*_BACnetLiftCarModeTaggedBuilder)(nil)

func (b *_BACnetLiftCarModeTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetLiftCarMode, proprietaryValue uint32) BACnetLiftCarModeTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetLiftCarModeTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetLiftCarModeTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetLiftCarModeTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLiftCarModeTaggedBuilder {
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

func (b *_BACnetLiftCarModeTaggedBuilder) WithValue(value BACnetLiftCarMode) BACnetLiftCarModeTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetLiftCarModeTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetLiftCarModeTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetLiftCarModeTaggedBuilder) Build() (BACnetLiftCarModeTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLiftCarModeTagged.deepCopy(), nil
}

func (b *_BACnetLiftCarModeTaggedBuilder) MustBuild() BACnetLiftCarModeTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLiftCarModeTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLiftCarModeTaggedBuilder().(*_BACnetLiftCarModeTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLiftCarModeTaggedBuilder creates a BACnetLiftCarModeTaggedBuilder
func (b *_BACnetLiftCarModeTagged) CreateBACnetLiftCarModeTaggedBuilder() BACnetLiftCarModeTaggedBuilder {
	if b == nil {
		return NewBACnetLiftCarModeTaggedBuilder()
	}
	return &_BACnetLiftCarModeTaggedBuilder{_BACnetLiftCarModeTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLiftCarModeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLiftCarModeTagged) GetValue() BACnetLiftCarMode {
	return m.Value
}

func (m *_BACnetLiftCarModeTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetLiftCarModeTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetLiftCarMode_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLiftCarModeTagged(structType any) BACnetLiftCarModeTagged {
	if casted, ok := structType.(BACnetLiftCarModeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLiftCarModeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLiftCarModeTagged) GetTypeName() string {
	return "BACnetLiftCarModeTagged"
}

func (m *_BACnetLiftCarModeTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetLiftCarModeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftCarModeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLiftCarModeTagged, error) {
	return BACnetLiftCarModeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLiftCarModeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarModeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarModeTagged, error) {
		return BACnetLiftCarModeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetLiftCarModeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLiftCarModeTagged, error) {
	v, err := (&_BACnetLiftCarModeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLiftCarModeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetLiftCarModeTagged BACnetLiftCarModeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLiftCarModeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLiftCarModeTagged")
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

	value, err := ReadManualField[BACnetLiftCarMode](ctx, "value", readBuffer, EnsureType[BACnetLiftCarMode](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetLiftCarMode_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetLiftCarMode_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetLiftCarModeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftCarModeTagged")
	}

	return m, nil
}

func (m *_BACnetLiftCarModeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLiftCarModeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLiftCarModeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftCarModeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetLiftCarMode](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetLiftCarModeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLiftCarModeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLiftCarModeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLiftCarModeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLiftCarModeTagged) IsBACnetLiftCarModeTagged() {}

func (m *_BACnetLiftCarModeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLiftCarModeTagged) deepCopy() *_BACnetLiftCarModeTagged {
	if m == nil {
		return nil
	}
	_BACnetLiftCarModeTaggedCopy := &_BACnetLiftCarModeTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetLiftCarModeTaggedCopy
}

func (m *_BACnetLiftCarModeTagged) String() string {
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

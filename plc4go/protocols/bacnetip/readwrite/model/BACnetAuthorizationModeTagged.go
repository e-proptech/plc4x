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

// BACnetAuthorizationModeTagged is the corresponding interface of BACnetAuthorizationModeTagged
type BACnetAuthorizationModeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAuthorizationMode
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetAuthorizationModeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAuthorizationModeTagged()
	// CreateBuilder creates a BACnetAuthorizationModeTaggedBuilder
	CreateBACnetAuthorizationModeTaggedBuilder() BACnetAuthorizationModeTaggedBuilder
}

// _BACnetAuthorizationModeTagged is the data-structure of this message
type _BACnetAuthorizationModeTagged struct {
	Header           BACnetTagHeader
	Value            BACnetAuthorizationMode
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAuthorizationModeTagged = (*_BACnetAuthorizationModeTagged)(nil)

// NewBACnetAuthorizationModeTagged factory function for _BACnetAuthorizationModeTagged
func NewBACnetAuthorizationModeTagged(header BACnetTagHeader, value BACnetAuthorizationMode, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetAuthorizationModeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAuthorizationModeTagged must not be nil")
	}
	return &_BACnetAuthorizationModeTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAuthorizationModeTaggedBuilder is a builder for BACnetAuthorizationModeTagged
type BACnetAuthorizationModeTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAuthorizationMode, proprietaryValue uint32) BACnetAuthorizationModeTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAuthorizationModeTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAuthorizationModeTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAuthorizationMode) BACnetAuthorizationModeTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetAuthorizationModeTaggedBuilder
	// Build builds the BACnetAuthorizationModeTagged or returns an error if something is wrong
	Build() (BACnetAuthorizationModeTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAuthorizationModeTagged
}

// NewBACnetAuthorizationModeTaggedBuilder() creates a BACnetAuthorizationModeTaggedBuilder
func NewBACnetAuthorizationModeTaggedBuilder() BACnetAuthorizationModeTaggedBuilder {
	return &_BACnetAuthorizationModeTaggedBuilder{_BACnetAuthorizationModeTagged: new(_BACnetAuthorizationModeTagged)}
}

type _BACnetAuthorizationModeTaggedBuilder struct {
	*_BACnetAuthorizationModeTagged

	err *utils.MultiError
}

var _ (BACnetAuthorizationModeTaggedBuilder) = (*_BACnetAuthorizationModeTaggedBuilder)(nil)

func (b *_BACnetAuthorizationModeTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAuthorizationMode, proprietaryValue uint32) BACnetAuthorizationModeTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetAuthorizationModeTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAuthorizationModeTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAuthorizationModeTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAuthorizationModeTaggedBuilder {
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

func (b *_BACnetAuthorizationModeTaggedBuilder) WithValue(value BACnetAuthorizationMode) BACnetAuthorizationModeTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAuthorizationModeTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetAuthorizationModeTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetAuthorizationModeTaggedBuilder) Build() (BACnetAuthorizationModeTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAuthorizationModeTagged.deepCopy(), nil
}

func (b *_BACnetAuthorizationModeTaggedBuilder) MustBuild() BACnetAuthorizationModeTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAuthorizationModeTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAuthorizationModeTaggedBuilder().(*_BACnetAuthorizationModeTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAuthorizationModeTaggedBuilder creates a BACnetAuthorizationModeTaggedBuilder
func (b *_BACnetAuthorizationModeTagged) CreateBACnetAuthorizationModeTaggedBuilder() BACnetAuthorizationModeTaggedBuilder {
	if b == nil {
		return NewBACnetAuthorizationModeTaggedBuilder()
	}
	return &_BACnetAuthorizationModeTaggedBuilder{_BACnetAuthorizationModeTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthorizationModeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAuthorizationModeTagged) GetValue() BACnetAuthorizationMode {
	return m.Value
}

func (m *_BACnetAuthorizationModeTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetAuthorizationModeTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAuthorizationMode_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAuthorizationModeTagged(structType any) BACnetAuthorizationModeTagged {
	if casted, ok := structType.(BACnetAuthorizationModeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthorizationModeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthorizationModeTagged) GetTypeName() string {
	return "BACnetAuthorizationModeTagged"
}

func (m *_BACnetAuthorizationModeTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetAuthorizationModeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthorizationModeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAuthorizationModeTagged, error) {
	return BACnetAuthorizationModeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAuthorizationModeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthorizationModeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthorizationModeTagged, error) {
		return BACnetAuthorizationModeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAuthorizationModeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAuthorizationModeTagged, error) {
	v, err := (&_BACnetAuthorizationModeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAuthorizationModeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAuthorizationModeTagged BACnetAuthorizationModeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthorizationModeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthorizationModeTagged")
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

	value, err := ReadManualField[BACnetAuthorizationMode](ctx, "value", readBuffer, EnsureType[BACnetAuthorizationMode](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetAuthorizationMode_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetAuthorizationMode_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetAuthorizationModeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthorizationModeTagged")
	}

	return m, nil
}

func (m *_BACnetAuthorizationModeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthorizationModeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthorizationModeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthorizationModeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAuthorizationMode](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetAuthorizationModeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthorizationModeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAuthorizationModeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAuthorizationModeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAuthorizationModeTagged) IsBACnetAuthorizationModeTagged() {}

func (m *_BACnetAuthorizationModeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAuthorizationModeTagged) deepCopy() *_BACnetAuthorizationModeTagged {
	if m == nil {
		return nil
	}
	_BACnetAuthorizationModeTaggedCopy := &_BACnetAuthorizationModeTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAuthorizationModeTaggedCopy
}

func (m *_BACnetAuthorizationModeTagged) String() string {
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

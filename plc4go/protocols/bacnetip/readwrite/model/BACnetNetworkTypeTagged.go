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

// BACnetNetworkTypeTagged is the corresponding interface of BACnetNetworkTypeTagged
type BACnetNetworkTypeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetNetworkType
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetNetworkTypeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNetworkTypeTagged()
	// CreateBuilder creates a BACnetNetworkTypeTaggedBuilder
	CreateBACnetNetworkTypeTaggedBuilder() BACnetNetworkTypeTaggedBuilder
}

// _BACnetNetworkTypeTagged is the data-structure of this message
type _BACnetNetworkTypeTagged struct {
	Header           BACnetTagHeader
	Value            BACnetNetworkType
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetNetworkTypeTagged = (*_BACnetNetworkTypeTagged)(nil)

// NewBACnetNetworkTypeTagged factory function for _BACnetNetworkTypeTagged
func NewBACnetNetworkTypeTagged(header BACnetTagHeader, value BACnetNetworkType, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetNetworkTypeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetNetworkTypeTagged must not be nil")
	}
	return &_BACnetNetworkTypeTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNetworkTypeTaggedBuilder is a builder for BACnetNetworkTypeTagged
type BACnetNetworkTypeTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetNetworkType, proprietaryValue uint32) BACnetNetworkTypeTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetNetworkTypeTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNetworkTypeTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetNetworkType) BACnetNetworkTypeTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetNetworkTypeTaggedBuilder
	// Build builds the BACnetNetworkTypeTagged or returns an error if something is wrong
	Build() (BACnetNetworkTypeTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNetworkTypeTagged
}

// NewBACnetNetworkTypeTaggedBuilder() creates a BACnetNetworkTypeTaggedBuilder
func NewBACnetNetworkTypeTaggedBuilder() BACnetNetworkTypeTaggedBuilder {
	return &_BACnetNetworkTypeTaggedBuilder{_BACnetNetworkTypeTagged: new(_BACnetNetworkTypeTagged)}
}

type _BACnetNetworkTypeTaggedBuilder struct {
	*_BACnetNetworkTypeTagged

	err *utils.MultiError
}

var _ (BACnetNetworkTypeTaggedBuilder) = (*_BACnetNetworkTypeTaggedBuilder)(nil)

func (b *_BACnetNetworkTypeTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetNetworkType, proprietaryValue uint32) BACnetNetworkTypeTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetNetworkTypeTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetNetworkTypeTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetNetworkTypeTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNetworkTypeTaggedBuilder {
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

func (b *_BACnetNetworkTypeTaggedBuilder) WithValue(value BACnetNetworkType) BACnetNetworkTypeTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetNetworkTypeTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetNetworkTypeTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetNetworkTypeTaggedBuilder) Build() (BACnetNetworkTypeTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNetworkTypeTagged.deepCopy(), nil
}

func (b *_BACnetNetworkTypeTaggedBuilder) MustBuild() BACnetNetworkTypeTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNetworkTypeTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNetworkTypeTaggedBuilder().(*_BACnetNetworkTypeTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNetworkTypeTaggedBuilder creates a BACnetNetworkTypeTaggedBuilder
func (b *_BACnetNetworkTypeTagged) CreateBACnetNetworkTypeTaggedBuilder() BACnetNetworkTypeTaggedBuilder {
	if b == nil {
		return NewBACnetNetworkTypeTaggedBuilder()
	}
	return &_BACnetNetworkTypeTaggedBuilder{_BACnetNetworkTypeTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNetworkTypeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetNetworkTypeTagged) GetValue() BACnetNetworkType {
	return m.Value
}

func (m *_BACnetNetworkTypeTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetNetworkTypeTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetNetworkType_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNetworkTypeTagged(structType any) BACnetNetworkTypeTagged {
	if casted, ok := structType.(BACnetNetworkTypeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNetworkTypeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNetworkTypeTagged) GetTypeName() string {
	return "BACnetNetworkTypeTagged"
}

func (m *_BACnetNetworkTypeTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetNetworkTypeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNetworkTypeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetNetworkTypeTagged, error) {
	return BACnetNetworkTypeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetNetworkTypeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNetworkTypeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNetworkTypeTagged, error) {
		return BACnetNetworkTypeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetNetworkTypeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetNetworkTypeTagged, error) {
	v, err := (&_BACnetNetworkTypeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetNetworkTypeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetNetworkTypeTagged BACnetNetworkTypeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNetworkTypeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNetworkTypeTagged")
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

	value, err := ReadManualField[BACnetNetworkType](ctx, "value", readBuffer, EnsureType[BACnetNetworkType](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetNetworkType_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetNetworkType_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetNetworkTypeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNetworkTypeTagged")
	}

	return m, nil
}

func (m *_BACnetNetworkTypeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNetworkTypeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNetworkTypeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNetworkTypeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetNetworkType](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetNetworkTypeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNetworkTypeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNetworkTypeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetNetworkTypeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetNetworkTypeTagged) IsBACnetNetworkTypeTagged() {}

func (m *_BACnetNetworkTypeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNetworkTypeTagged) deepCopy() *_BACnetNetworkTypeTagged {
	if m == nil {
		return nil
	}
	_BACnetNetworkTypeTaggedCopy := &_BACnetNetworkTypeTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetNetworkTypeTaggedCopy
}

func (m *_BACnetNetworkTypeTagged) String() string {
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

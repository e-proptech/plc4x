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

// BACnetMaintenanceTagged is the corresponding interface of BACnetMaintenanceTagged
type BACnetMaintenanceTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetMaintenance
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetMaintenanceTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetMaintenanceTagged()
	// CreateBuilder creates a BACnetMaintenanceTaggedBuilder
	CreateBACnetMaintenanceTaggedBuilder() BACnetMaintenanceTaggedBuilder
}

// _BACnetMaintenanceTagged is the data-structure of this message
type _BACnetMaintenanceTagged struct {
	Header           BACnetTagHeader
	Value            BACnetMaintenance
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetMaintenanceTagged = (*_BACnetMaintenanceTagged)(nil)

// NewBACnetMaintenanceTagged factory function for _BACnetMaintenanceTagged
func NewBACnetMaintenanceTagged(header BACnetTagHeader, value BACnetMaintenance, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetMaintenanceTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetMaintenanceTagged must not be nil")
	}
	return &_BACnetMaintenanceTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetMaintenanceTaggedBuilder is a builder for BACnetMaintenanceTagged
type BACnetMaintenanceTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetMaintenance, proprietaryValue uint32) BACnetMaintenanceTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetMaintenanceTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetMaintenanceTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetMaintenance) BACnetMaintenanceTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetMaintenanceTaggedBuilder
	// Build builds the BACnetMaintenanceTagged or returns an error if something is wrong
	Build() (BACnetMaintenanceTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetMaintenanceTagged
}

// NewBACnetMaintenanceTaggedBuilder() creates a BACnetMaintenanceTaggedBuilder
func NewBACnetMaintenanceTaggedBuilder() BACnetMaintenanceTaggedBuilder {
	return &_BACnetMaintenanceTaggedBuilder{_BACnetMaintenanceTagged: new(_BACnetMaintenanceTagged)}
}

type _BACnetMaintenanceTaggedBuilder struct {
	*_BACnetMaintenanceTagged

	err *utils.MultiError
}

var _ (BACnetMaintenanceTaggedBuilder) = (*_BACnetMaintenanceTaggedBuilder)(nil)

func (b *_BACnetMaintenanceTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetMaintenance, proprietaryValue uint32) BACnetMaintenanceTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetMaintenanceTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetMaintenanceTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetMaintenanceTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetMaintenanceTaggedBuilder {
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

func (b *_BACnetMaintenanceTaggedBuilder) WithValue(value BACnetMaintenance) BACnetMaintenanceTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetMaintenanceTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetMaintenanceTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetMaintenanceTaggedBuilder) Build() (BACnetMaintenanceTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetMaintenanceTagged.deepCopy(), nil
}

func (b *_BACnetMaintenanceTaggedBuilder) MustBuild() BACnetMaintenanceTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetMaintenanceTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetMaintenanceTaggedBuilder().(*_BACnetMaintenanceTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetMaintenanceTaggedBuilder creates a BACnetMaintenanceTaggedBuilder
func (b *_BACnetMaintenanceTagged) CreateBACnetMaintenanceTaggedBuilder() BACnetMaintenanceTaggedBuilder {
	if b == nil {
		return NewBACnetMaintenanceTaggedBuilder()
	}
	return &_BACnetMaintenanceTaggedBuilder{_BACnetMaintenanceTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetMaintenanceTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetMaintenanceTagged) GetValue() BACnetMaintenance {
	return m.Value
}

func (m *_BACnetMaintenanceTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetMaintenanceTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetMaintenance_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetMaintenanceTagged(structType any) BACnetMaintenanceTagged {
	if casted, ok := structType.(BACnetMaintenanceTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetMaintenanceTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetMaintenanceTagged) GetTypeName() string {
	return "BACnetMaintenanceTagged"
}

func (m *_BACnetMaintenanceTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetMaintenanceTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetMaintenanceTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetMaintenanceTagged, error) {
	return BACnetMaintenanceTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetMaintenanceTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetMaintenanceTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetMaintenanceTagged, error) {
		return BACnetMaintenanceTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetMaintenanceTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetMaintenanceTagged, error) {
	v, err := (&_BACnetMaintenanceTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetMaintenanceTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetMaintenanceTagged BACnetMaintenanceTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetMaintenanceTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetMaintenanceTagged")
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

	value, err := ReadManualField[BACnetMaintenance](ctx, "value", readBuffer, EnsureType[BACnetMaintenance](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetMaintenance_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetMaintenance_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetMaintenanceTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetMaintenanceTagged")
	}

	return m, nil
}

func (m *_BACnetMaintenanceTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetMaintenanceTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetMaintenanceTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetMaintenanceTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetMaintenance](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetMaintenanceTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetMaintenanceTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetMaintenanceTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetMaintenanceTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetMaintenanceTagged) IsBACnetMaintenanceTagged() {}

func (m *_BACnetMaintenanceTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetMaintenanceTagged) deepCopy() *_BACnetMaintenanceTagged {
	if m == nil {
		return nil
	}
	_BACnetMaintenanceTaggedCopy := &_BACnetMaintenanceTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetMaintenanceTaggedCopy
}

func (m *_BACnetMaintenanceTagged) String() string {
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

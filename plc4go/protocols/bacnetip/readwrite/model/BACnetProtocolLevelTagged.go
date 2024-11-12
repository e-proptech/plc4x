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

// BACnetProtocolLevelTagged is the corresponding interface of BACnetProtocolLevelTagged
type BACnetProtocolLevelTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetProtocolLevel
	// IsBACnetProtocolLevelTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetProtocolLevelTagged()
	// CreateBuilder creates a BACnetProtocolLevelTaggedBuilder
	CreateBACnetProtocolLevelTaggedBuilder() BACnetProtocolLevelTaggedBuilder
}

// _BACnetProtocolLevelTagged is the data-structure of this message
type _BACnetProtocolLevelTagged struct {
	Header BACnetTagHeader
	Value  BACnetProtocolLevel

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetProtocolLevelTagged = (*_BACnetProtocolLevelTagged)(nil)

// NewBACnetProtocolLevelTagged factory function for _BACnetProtocolLevelTagged
func NewBACnetProtocolLevelTagged(header BACnetTagHeader, value BACnetProtocolLevel, tagNumber uint8, tagClass TagClass) *_BACnetProtocolLevelTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetProtocolLevelTagged must not be nil")
	}
	return &_BACnetProtocolLevelTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetProtocolLevelTaggedBuilder is a builder for BACnetProtocolLevelTagged
type BACnetProtocolLevelTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetProtocolLevel) BACnetProtocolLevelTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetProtocolLevelTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetProtocolLevelTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetProtocolLevel) BACnetProtocolLevelTaggedBuilder
	// Build builds the BACnetProtocolLevelTagged or returns an error if something is wrong
	Build() (BACnetProtocolLevelTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetProtocolLevelTagged
}

// NewBACnetProtocolLevelTaggedBuilder() creates a BACnetProtocolLevelTaggedBuilder
func NewBACnetProtocolLevelTaggedBuilder() BACnetProtocolLevelTaggedBuilder {
	return &_BACnetProtocolLevelTaggedBuilder{_BACnetProtocolLevelTagged: new(_BACnetProtocolLevelTagged)}
}

type _BACnetProtocolLevelTaggedBuilder struct {
	*_BACnetProtocolLevelTagged

	err *utils.MultiError
}

var _ (BACnetProtocolLevelTaggedBuilder) = (*_BACnetProtocolLevelTaggedBuilder)(nil)

func (b *_BACnetProtocolLevelTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetProtocolLevel) BACnetProtocolLevelTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetProtocolLevelTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetProtocolLevelTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetProtocolLevelTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetProtocolLevelTaggedBuilder {
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

func (b *_BACnetProtocolLevelTaggedBuilder) WithValue(value BACnetProtocolLevel) BACnetProtocolLevelTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetProtocolLevelTaggedBuilder) Build() (BACnetProtocolLevelTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetProtocolLevelTagged.deepCopy(), nil
}

func (b *_BACnetProtocolLevelTaggedBuilder) MustBuild() BACnetProtocolLevelTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetProtocolLevelTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetProtocolLevelTaggedBuilder().(*_BACnetProtocolLevelTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetProtocolLevelTaggedBuilder creates a BACnetProtocolLevelTaggedBuilder
func (b *_BACnetProtocolLevelTagged) CreateBACnetProtocolLevelTaggedBuilder() BACnetProtocolLevelTaggedBuilder {
	if b == nil {
		return NewBACnetProtocolLevelTaggedBuilder()
	}
	return &_BACnetProtocolLevelTaggedBuilder{_BACnetProtocolLevelTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProtocolLevelTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetProtocolLevelTagged) GetValue() BACnetProtocolLevel {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetProtocolLevelTagged(structType any) BACnetProtocolLevelTagged {
	if casted, ok := structType.(BACnetProtocolLevelTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProtocolLevelTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProtocolLevelTagged) GetTypeName() string {
	return "BACnetProtocolLevelTagged"
}

func (m *_BACnetProtocolLevelTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetProtocolLevelTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetProtocolLevelTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetProtocolLevelTagged, error) {
	return BACnetProtocolLevelTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetProtocolLevelTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetProtocolLevelTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetProtocolLevelTagged, error) {
		return BACnetProtocolLevelTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetProtocolLevelTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetProtocolLevelTagged, error) {
	v, err := (&_BACnetProtocolLevelTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetProtocolLevelTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetProtocolLevelTagged BACnetProtocolLevelTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProtocolLevelTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProtocolLevelTagged")
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

	value, err := ReadManualField[BACnetProtocolLevel](ctx, "value", readBuffer, EnsureType[BACnetProtocolLevel](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetProtocolLevel_PHYSICAL)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetProtocolLevelTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProtocolLevelTagged")
	}

	return m, nil
}

func (m *_BACnetProtocolLevelTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProtocolLevelTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetProtocolLevelTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetProtocolLevelTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetProtocolLevel](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetProtocolLevelTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetProtocolLevelTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetProtocolLevelTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetProtocolLevelTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetProtocolLevelTagged) IsBACnetProtocolLevelTagged() {}

func (m *_BACnetProtocolLevelTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetProtocolLevelTagged) deepCopy() *_BACnetProtocolLevelTagged {
	if m == nil {
		return nil
	}
	_BACnetProtocolLevelTaggedCopy := &_BACnetProtocolLevelTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetProtocolLevelTaggedCopy
}

func (m *_BACnetProtocolLevelTagged) String() string {
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

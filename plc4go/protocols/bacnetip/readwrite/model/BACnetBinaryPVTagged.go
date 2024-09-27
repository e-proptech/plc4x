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

// BACnetBinaryPVTagged is the corresponding interface of BACnetBinaryPVTagged
type BACnetBinaryPVTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetBinaryPV
	// IsBACnetBinaryPVTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetBinaryPVTagged()
	// CreateBuilder creates a BACnetBinaryPVTaggedBuilder
	CreateBACnetBinaryPVTaggedBuilder() BACnetBinaryPVTaggedBuilder
}

// _BACnetBinaryPVTagged is the data-structure of this message
type _BACnetBinaryPVTagged struct {
	Header BACnetTagHeader
	Value  BACnetBinaryPV

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetBinaryPVTagged = (*_BACnetBinaryPVTagged)(nil)

// NewBACnetBinaryPVTagged factory function for _BACnetBinaryPVTagged
func NewBACnetBinaryPVTagged(header BACnetTagHeader, value BACnetBinaryPV, tagNumber uint8, tagClass TagClass) *_BACnetBinaryPVTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetBinaryPVTagged must not be nil")
	}
	return &_BACnetBinaryPVTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetBinaryPVTaggedBuilder is a builder for BACnetBinaryPVTagged
type BACnetBinaryPVTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetBinaryPV) BACnetBinaryPVTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetBinaryPVTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetBinaryPVTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetBinaryPV) BACnetBinaryPVTaggedBuilder
	// Build builds the BACnetBinaryPVTagged or returns an error if something is wrong
	Build() (BACnetBinaryPVTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetBinaryPVTagged
}

// NewBACnetBinaryPVTaggedBuilder() creates a BACnetBinaryPVTaggedBuilder
func NewBACnetBinaryPVTaggedBuilder() BACnetBinaryPVTaggedBuilder {
	return &_BACnetBinaryPVTaggedBuilder{_BACnetBinaryPVTagged: new(_BACnetBinaryPVTagged)}
}

type _BACnetBinaryPVTaggedBuilder struct {
	*_BACnetBinaryPVTagged

	err *utils.MultiError
}

var _ (BACnetBinaryPVTaggedBuilder) = (*_BACnetBinaryPVTaggedBuilder)(nil)

func (b *_BACnetBinaryPVTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetBinaryPV) BACnetBinaryPVTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetBinaryPVTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetBinaryPVTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetBinaryPVTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetBinaryPVTaggedBuilder {
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

func (b *_BACnetBinaryPVTaggedBuilder) WithValue(value BACnetBinaryPV) BACnetBinaryPVTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetBinaryPVTaggedBuilder) Build() (BACnetBinaryPVTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetBinaryPVTagged.deepCopy(), nil
}

func (b *_BACnetBinaryPVTaggedBuilder) MustBuild() BACnetBinaryPVTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetBinaryPVTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetBinaryPVTaggedBuilder().(*_BACnetBinaryPVTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetBinaryPVTaggedBuilder creates a BACnetBinaryPVTaggedBuilder
func (b *_BACnetBinaryPVTagged) CreateBACnetBinaryPVTaggedBuilder() BACnetBinaryPVTaggedBuilder {
	if b == nil {
		return NewBACnetBinaryPVTaggedBuilder()
	}
	return &_BACnetBinaryPVTaggedBuilder{_BACnetBinaryPVTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetBinaryPVTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetBinaryPVTagged) GetValue() BACnetBinaryPV {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetBinaryPVTagged(structType any) BACnetBinaryPVTagged {
	if casted, ok := structType.(BACnetBinaryPVTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetBinaryPVTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetBinaryPVTagged) GetTypeName() string {
	return "BACnetBinaryPVTagged"
}

func (m *_BACnetBinaryPVTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetBinaryPVTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetBinaryPVTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetBinaryPVTagged, error) {
	return BACnetBinaryPVTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetBinaryPVTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetBinaryPVTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetBinaryPVTagged, error) {
		return BACnetBinaryPVTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetBinaryPVTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetBinaryPVTagged, error) {
	v, err := (&_BACnetBinaryPVTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetBinaryPVTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetBinaryPVTagged BACnetBinaryPVTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetBinaryPVTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetBinaryPVTagged")
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

	value, err := ReadManualField[BACnetBinaryPV](ctx, "value", readBuffer, EnsureType[BACnetBinaryPV](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetBinaryPV_INACTIVE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetBinaryPVTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetBinaryPVTagged")
	}

	return m, nil
}

func (m *_BACnetBinaryPVTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetBinaryPVTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetBinaryPVTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetBinaryPVTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetBinaryPV](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetBinaryPVTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetBinaryPVTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetBinaryPVTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetBinaryPVTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetBinaryPVTagged) IsBACnetBinaryPVTagged() {}

func (m *_BACnetBinaryPVTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetBinaryPVTagged) deepCopy() *_BACnetBinaryPVTagged {
	if m == nil {
		return nil
	}
	_BACnetBinaryPVTaggedCopy := &_BACnetBinaryPVTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetBinaryPVTaggedCopy
}

func (m *_BACnetBinaryPVTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

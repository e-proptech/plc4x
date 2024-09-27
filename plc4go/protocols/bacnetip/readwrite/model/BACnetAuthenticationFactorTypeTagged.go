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

// BACnetAuthenticationFactorTypeTagged is the corresponding interface of BACnetAuthenticationFactorTypeTagged
type BACnetAuthenticationFactorTypeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAuthenticationFactorType
	// IsBACnetAuthenticationFactorTypeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAuthenticationFactorTypeTagged()
	// CreateBuilder creates a BACnetAuthenticationFactorTypeTaggedBuilder
	CreateBACnetAuthenticationFactorTypeTaggedBuilder() BACnetAuthenticationFactorTypeTaggedBuilder
}

// _BACnetAuthenticationFactorTypeTagged is the data-structure of this message
type _BACnetAuthenticationFactorTypeTagged struct {
	Header BACnetTagHeader
	Value  BACnetAuthenticationFactorType

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAuthenticationFactorTypeTagged = (*_BACnetAuthenticationFactorTypeTagged)(nil)

// NewBACnetAuthenticationFactorTypeTagged factory function for _BACnetAuthenticationFactorTypeTagged
func NewBACnetAuthenticationFactorTypeTagged(header BACnetTagHeader, value BACnetAuthenticationFactorType, tagNumber uint8, tagClass TagClass) *_BACnetAuthenticationFactorTypeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAuthenticationFactorTypeTagged must not be nil")
	}
	return &_BACnetAuthenticationFactorTypeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAuthenticationFactorTypeTaggedBuilder is a builder for BACnetAuthenticationFactorTypeTagged
type BACnetAuthenticationFactorTypeTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAuthenticationFactorType) BACnetAuthenticationFactorTypeTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAuthenticationFactorTypeTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAuthenticationFactorTypeTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAuthenticationFactorType) BACnetAuthenticationFactorTypeTaggedBuilder
	// Build builds the BACnetAuthenticationFactorTypeTagged or returns an error if something is wrong
	Build() (BACnetAuthenticationFactorTypeTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAuthenticationFactorTypeTagged
}

// NewBACnetAuthenticationFactorTypeTaggedBuilder() creates a BACnetAuthenticationFactorTypeTaggedBuilder
func NewBACnetAuthenticationFactorTypeTaggedBuilder() BACnetAuthenticationFactorTypeTaggedBuilder {
	return &_BACnetAuthenticationFactorTypeTaggedBuilder{_BACnetAuthenticationFactorTypeTagged: new(_BACnetAuthenticationFactorTypeTagged)}
}

type _BACnetAuthenticationFactorTypeTaggedBuilder struct {
	*_BACnetAuthenticationFactorTypeTagged

	err *utils.MultiError
}

var _ (BACnetAuthenticationFactorTypeTaggedBuilder) = (*_BACnetAuthenticationFactorTypeTaggedBuilder)(nil)

func (b *_BACnetAuthenticationFactorTypeTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAuthenticationFactorType) BACnetAuthenticationFactorTypeTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetAuthenticationFactorTypeTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAuthenticationFactorTypeTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAuthenticationFactorTypeTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAuthenticationFactorTypeTaggedBuilder {
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

func (b *_BACnetAuthenticationFactorTypeTaggedBuilder) WithValue(value BACnetAuthenticationFactorType) BACnetAuthenticationFactorTypeTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAuthenticationFactorTypeTaggedBuilder) Build() (BACnetAuthenticationFactorTypeTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAuthenticationFactorTypeTagged.deepCopy(), nil
}

func (b *_BACnetAuthenticationFactorTypeTaggedBuilder) MustBuild() BACnetAuthenticationFactorTypeTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAuthenticationFactorTypeTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAuthenticationFactorTypeTaggedBuilder().(*_BACnetAuthenticationFactorTypeTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAuthenticationFactorTypeTaggedBuilder creates a BACnetAuthenticationFactorTypeTaggedBuilder
func (b *_BACnetAuthenticationFactorTypeTagged) CreateBACnetAuthenticationFactorTypeTaggedBuilder() BACnetAuthenticationFactorTypeTaggedBuilder {
	if b == nil {
		return NewBACnetAuthenticationFactorTypeTaggedBuilder()
	}
	return &_BACnetAuthenticationFactorTypeTaggedBuilder{_BACnetAuthenticationFactorTypeTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationFactorTypeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAuthenticationFactorTypeTagged) GetValue() BACnetAuthenticationFactorType {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationFactorTypeTagged(structType any) BACnetAuthenticationFactorTypeTagged {
	if casted, ok := structType.(BACnetAuthenticationFactorTypeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationFactorTypeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationFactorTypeTagged) GetTypeName() string {
	return "BACnetAuthenticationFactorTypeTagged"
}

func (m *_BACnetAuthenticationFactorTypeTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetAuthenticationFactorTypeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationFactorTypeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAuthenticationFactorTypeTagged, error) {
	return BACnetAuthenticationFactorTypeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAuthenticationFactorTypeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactorTypeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactorTypeTagged, error) {
		return BACnetAuthenticationFactorTypeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAuthenticationFactorTypeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAuthenticationFactorTypeTagged, error) {
	v, err := (&_BACnetAuthenticationFactorTypeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAuthenticationFactorTypeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAuthenticationFactorTypeTagged BACnetAuthenticationFactorTypeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationFactorTypeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationFactorTypeTagged")
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

	value, err := ReadManualField[BACnetAuthenticationFactorType](ctx, "value", readBuffer, EnsureType[BACnetAuthenticationFactorType](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetAuthenticationFactorType_UNDEFINED)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationFactorTypeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationFactorTypeTagged")
	}

	return m, nil
}

func (m *_BACnetAuthenticationFactorTypeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationFactorTypeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationFactorTypeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationFactorTypeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAuthenticationFactorType](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationFactorTypeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationFactorTypeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAuthenticationFactorTypeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAuthenticationFactorTypeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAuthenticationFactorTypeTagged) IsBACnetAuthenticationFactorTypeTagged() {}

func (m *_BACnetAuthenticationFactorTypeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAuthenticationFactorTypeTagged) deepCopy() *_BACnetAuthenticationFactorTypeTagged {
	if m == nil {
		return nil
	}
	_BACnetAuthenticationFactorTypeTaggedCopy := &_BACnetAuthenticationFactorTypeTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAuthenticationFactorTypeTaggedCopy
}

func (m *_BACnetAuthenticationFactorTypeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

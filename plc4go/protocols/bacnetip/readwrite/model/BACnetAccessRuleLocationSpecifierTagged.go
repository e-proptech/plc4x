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

// BACnetAccessRuleLocationSpecifierTagged is the corresponding interface of BACnetAccessRuleLocationSpecifierTagged
type BACnetAccessRuleLocationSpecifierTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessRuleLocationSpecifier
	// IsBACnetAccessRuleLocationSpecifierTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccessRuleLocationSpecifierTagged()
	// CreateBuilder creates a BACnetAccessRuleLocationSpecifierTaggedBuilder
	CreateBACnetAccessRuleLocationSpecifierTaggedBuilder() BACnetAccessRuleLocationSpecifierTaggedBuilder
}

// _BACnetAccessRuleLocationSpecifierTagged is the data-structure of this message
type _BACnetAccessRuleLocationSpecifierTagged struct {
	Header BACnetTagHeader
	Value  BACnetAccessRuleLocationSpecifier

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAccessRuleLocationSpecifierTagged = (*_BACnetAccessRuleLocationSpecifierTagged)(nil)

// NewBACnetAccessRuleLocationSpecifierTagged factory function for _BACnetAccessRuleLocationSpecifierTagged
func NewBACnetAccessRuleLocationSpecifierTagged(header BACnetTagHeader, value BACnetAccessRuleLocationSpecifier, tagNumber uint8, tagClass TagClass) *_BACnetAccessRuleLocationSpecifierTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAccessRuleLocationSpecifierTagged must not be nil")
	}
	return &_BACnetAccessRuleLocationSpecifierTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAccessRuleLocationSpecifierTaggedBuilder is a builder for BACnetAccessRuleLocationSpecifierTagged
type BACnetAccessRuleLocationSpecifierTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAccessRuleLocationSpecifier) BACnetAccessRuleLocationSpecifierTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAccessRuleLocationSpecifierTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessRuleLocationSpecifierTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAccessRuleLocationSpecifier) BACnetAccessRuleLocationSpecifierTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetAccessRuleLocationSpecifierTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetAccessRuleLocationSpecifierTaggedBuilder
	// Build builds the BACnetAccessRuleLocationSpecifierTagged or returns an error if something is wrong
	Build() (BACnetAccessRuleLocationSpecifierTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAccessRuleLocationSpecifierTagged
}

// NewBACnetAccessRuleLocationSpecifierTaggedBuilder() creates a BACnetAccessRuleLocationSpecifierTaggedBuilder
func NewBACnetAccessRuleLocationSpecifierTaggedBuilder() BACnetAccessRuleLocationSpecifierTaggedBuilder {
	return &_BACnetAccessRuleLocationSpecifierTaggedBuilder{_BACnetAccessRuleLocationSpecifierTagged: new(_BACnetAccessRuleLocationSpecifierTagged)}
}

type _BACnetAccessRuleLocationSpecifierTaggedBuilder struct {
	*_BACnetAccessRuleLocationSpecifierTagged

	err *utils.MultiError
}

var _ (BACnetAccessRuleLocationSpecifierTaggedBuilder) = (*_BACnetAccessRuleLocationSpecifierTaggedBuilder)(nil)

func (b *_BACnetAccessRuleLocationSpecifierTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAccessRuleLocationSpecifier) BACnetAccessRuleLocationSpecifierTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetAccessRuleLocationSpecifierTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAccessRuleLocationSpecifierTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAccessRuleLocationSpecifierTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessRuleLocationSpecifierTaggedBuilder {
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

func (b *_BACnetAccessRuleLocationSpecifierTaggedBuilder) WithValue(value BACnetAccessRuleLocationSpecifier) BACnetAccessRuleLocationSpecifierTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAccessRuleLocationSpecifierTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetAccessRuleLocationSpecifierTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetAccessRuleLocationSpecifierTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetAccessRuleLocationSpecifierTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetAccessRuleLocationSpecifierTaggedBuilder) Build() (BACnetAccessRuleLocationSpecifierTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAccessRuleLocationSpecifierTagged.deepCopy(), nil
}

func (b *_BACnetAccessRuleLocationSpecifierTaggedBuilder) MustBuild() BACnetAccessRuleLocationSpecifierTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAccessRuleLocationSpecifierTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAccessRuleLocationSpecifierTaggedBuilder().(*_BACnetAccessRuleLocationSpecifierTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAccessRuleLocationSpecifierTaggedBuilder creates a BACnetAccessRuleLocationSpecifierTaggedBuilder
func (b *_BACnetAccessRuleLocationSpecifierTagged) CreateBACnetAccessRuleLocationSpecifierTaggedBuilder() BACnetAccessRuleLocationSpecifierTaggedBuilder {
	if b == nil {
		return NewBACnetAccessRuleLocationSpecifierTaggedBuilder()
	}
	return &_BACnetAccessRuleLocationSpecifierTaggedBuilder{_BACnetAccessRuleLocationSpecifierTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessRuleLocationSpecifierTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessRuleLocationSpecifierTagged) GetValue() BACnetAccessRuleLocationSpecifier {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAccessRuleLocationSpecifierTagged(structType any) BACnetAccessRuleLocationSpecifierTagged {
	if casted, ok := structType.(BACnetAccessRuleLocationSpecifierTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessRuleLocationSpecifierTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessRuleLocationSpecifierTagged) GetTypeName() string {
	return "BACnetAccessRuleLocationSpecifierTagged"
}

func (m *_BACnetAccessRuleLocationSpecifierTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetAccessRuleLocationSpecifierTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessRuleLocationSpecifierTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccessRuleLocationSpecifierTagged, error) {
	return BACnetAccessRuleLocationSpecifierTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccessRuleLocationSpecifierTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessRuleLocationSpecifierTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessRuleLocationSpecifierTagged, error) {
		return BACnetAccessRuleLocationSpecifierTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAccessRuleLocationSpecifierTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessRuleLocationSpecifierTagged, error) {
	v, err := (&_BACnetAccessRuleLocationSpecifierTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccessRuleLocationSpecifierTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAccessRuleLocationSpecifierTagged BACnetAccessRuleLocationSpecifierTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessRuleLocationSpecifierTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessRuleLocationSpecifierTagged")
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

	value, err := ReadManualField[BACnetAccessRuleLocationSpecifier](ctx, "value", readBuffer, EnsureType[BACnetAccessRuleLocationSpecifier](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetAccessRuleLocationSpecifier_SPECIFIED)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetAccessRuleLocationSpecifierTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessRuleLocationSpecifierTagged")
	}

	return m, nil
}

func (m *_BACnetAccessRuleLocationSpecifierTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessRuleLocationSpecifierTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessRuleLocationSpecifierTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessRuleLocationSpecifierTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAccessRuleLocationSpecifier](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessRuleLocationSpecifierTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessRuleLocationSpecifierTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccessRuleLocationSpecifierTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessRuleLocationSpecifierTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccessRuleLocationSpecifierTagged) IsBACnetAccessRuleLocationSpecifierTagged() {}

func (m *_BACnetAccessRuleLocationSpecifierTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAccessRuleLocationSpecifierTagged) deepCopy() *_BACnetAccessRuleLocationSpecifierTagged {
	if m == nil {
		return nil
	}
	_BACnetAccessRuleLocationSpecifierTaggedCopy := &_BACnetAccessRuleLocationSpecifierTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAccessRuleLocationSpecifierTaggedCopy
}

func (m *_BACnetAccessRuleLocationSpecifierTagged) String() string {
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

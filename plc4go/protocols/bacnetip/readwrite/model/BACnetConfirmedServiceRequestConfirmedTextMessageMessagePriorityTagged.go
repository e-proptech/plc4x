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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
type BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority
	// IsBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged()
	// CreateBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder
	CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder
}

// _BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged struct {
	Header BACnetTagHeader
	Value  BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged)(nil)

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged factory function for _BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged(header BACnetTagHeader, value BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority, tagNumber uint8, tagClass TagClass) *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged must not be nil")
	}
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder is a builder for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
type BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder
	// Build builds the BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
}

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder() creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder {
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder{_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged: new(_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged)}
}

type _BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder struct {
	*_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder {
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

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) WithValue(value BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) Build() (BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) MustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder().(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder
func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder()
	}
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedBuilder{_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetValue() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged(structType any) BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, error) {
	return BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, error) {
		return BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, error) {
	v, err := (&_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged")
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

	value, err := ReadManualField[BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority](ctx, "value", readBuffer, EnsureType[BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority_NORMAL)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriority](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) IsBACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged() {
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) deepCopy() *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedCopy := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedCopy
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged) String() string {
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

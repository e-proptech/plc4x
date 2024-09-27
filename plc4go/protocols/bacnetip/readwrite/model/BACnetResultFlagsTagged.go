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

// BACnetResultFlagsTagged is the corresponding interface of BACnetResultFlagsTagged
type BACnetResultFlagsTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetFirstItem returns FirstItem (virtual field)
	GetFirstItem() bool
	// GetLastItem returns LastItem (virtual field)
	GetLastItem() bool
	// GetMoreItems returns MoreItems (virtual field)
	GetMoreItems() bool
	// IsBACnetResultFlagsTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetResultFlagsTagged()
	// CreateBuilder creates a BACnetResultFlagsTaggedBuilder
	CreateBACnetResultFlagsTaggedBuilder() BACnetResultFlagsTaggedBuilder
}

// _BACnetResultFlagsTagged is the data-structure of this message
type _BACnetResultFlagsTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetResultFlagsTagged = (*_BACnetResultFlagsTagged)(nil)

// NewBACnetResultFlagsTagged factory function for _BACnetResultFlagsTagged
func NewBACnetResultFlagsTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetResultFlagsTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetResultFlagsTagged must not be nil")
	}
	if payload == nil {
		panic("payload of type BACnetTagPayloadBitString for BACnetResultFlagsTagged must not be nil")
	}
	return &_BACnetResultFlagsTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetResultFlagsTaggedBuilder is a builder for BACnetResultFlagsTagged
type BACnetResultFlagsTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, payload BACnetTagPayloadBitString) BACnetResultFlagsTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetResultFlagsTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetResultFlagsTaggedBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadBitString) BACnetResultFlagsTaggedBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadBitStringBuilder) BACnetTagPayloadBitStringBuilder) BACnetResultFlagsTaggedBuilder
	// Build builds the BACnetResultFlagsTagged or returns an error if something is wrong
	Build() (BACnetResultFlagsTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetResultFlagsTagged
}

// NewBACnetResultFlagsTaggedBuilder() creates a BACnetResultFlagsTaggedBuilder
func NewBACnetResultFlagsTaggedBuilder() BACnetResultFlagsTaggedBuilder {
	return &_BACnetResultFlagsTaggedBuilder{_BACnetResultFlagsTagged: new(_BACnetResultFlagsTagged)}
}

type _BACnetResultFlagsTaggedBuilder struct {
	*_BACnetResultFlagsTagged

	err *utils.MultiError
}

var _ (BACnetResultFlagsTaggedBuilder) = (*_BACnetResultFlagsTaggedBuilder)(nil)

func (b *_BACnetResultFlagsTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, payload BACnetTagPayloadBitString) BACnetResultFlagsTaggedBuilder {
	return b.WithHeader(header).WithPayload(payload)
}

func (b *_BACnetResultFlagsTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetResultFlagsTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetResultFlagsTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetResultFlagsTaggedBuilder {
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

func (b *_BACnetResultFlagsTaggedBuilder) WithPayload(payload BACnetTagPayloadBitString) BACnetResultFlagsTaggedBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetResultFlagsTaggedBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadBitStringBuilder) BACnetTagPayloadBitStringBuilder) BACnetResultFlagsTaggedBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadBitStringBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadBitStringBuilder failed"))
	}
	return b
}

func (b *_BACnetResultFlagsTaggedBuilder) Build() (BACnetResultFlagsTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetResultFlagsTagged.deepCopy(), nil
}

func (b *_BACnetResultFlagsTaggedBuilder) MustBuild() BACnetResultFlagsTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetResultFlagsTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetResultFlagsTaggedBuilder().(*_BACnetResultFlagsTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetResultFlagsTaggedBuilder creates a BACnetResultFlagsTaggedBuilder
func (b *_BACnetResultFlagsTagged) CreateBACnetResultFlagsTaggedBuilder() BACnetResultFlagsTaggedBuilder {
	if b == nil {
		return NewBACnetResultFlagsTaggedBuilder()
	}
	return &_BACnetResultFlagsTaggedBuilder{_BACnetResultFlagsTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetResultFlagsTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetResultFlagsTagged) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetResultFlagsTagged) GetFirstItem() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (0))), func() any { return bool(m.GetPayload().GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetResultFlagsTagged) GetLastItem() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (1))), func() any { return bool(m.GetPayload().GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetResultFlagsTagged) GetMoreItems() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (2))), func() any { return bool(m.GetPayload().GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetResultFlagsTagged(structType any) BACnetResultFlagsTagged {
	if casted, ok := structType.(BACnetResultFlagsTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetResultFlagsTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetResultFlagsTagged) GetTypeName() string {
	return "BACnetResultFlagsTagged"
}

func (m *_BACnetResultFlagsTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetResultFlagsTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetResultFlagsTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetResultFlagsTagged, error) {
	return BACnetResultFlagsTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetResultFlagsTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetResultFlagsTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetResultFlagsTagged, error) {
		return BACnetResultFlagsTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetResultFlagsTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetResultFlagsTagged, error) {
	v, err := (&_BACnetResultFlagsTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetResultFlagsTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetResultFlagsTagged BACnetResultFlagsTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetResultFlagsTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetResultFlagsTagged")
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

	payload, err := ReadSimpleField[BACnetTagPayloadBitString](ctx, "payload", ReadComplex[BACnetTagPayloadBitString](BACnetTagPayloadBitStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	firstItem, err := ReadVirtualField[bool](ctx, "firstItem", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (0))), func() any { return bool(payload.GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'firstItem' field"))
	}
	_ = firstItem

	lastItem, err := ReadVirtualField[bool](ctx, "lastItem", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (1))), func() any { return bool(payload.GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastItem' field"))
	}
	_ = lastItem

	moreItems, err := ReadVirtualField[bool](ctx, "moreItems", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (2))), func() any { return bool(payload.GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moreItems' field"))
	}
	_ = moreItems

	if closeErr := readBuffer.CloseContext("BACnetResultFlagsTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetResultFlagsTagged")
	}

	return m, nil
}

func (m *_BACnetResultFlagsTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetResultFlagsTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetResultFlagsTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetResultFlagsTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteSimpleField[BACnetTagPayloadBitString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadBitString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}
	// Virtual field
	firstItem := m.GetFirstItem()
	_ = firstItem
	if _firstItemErr := writeBuffer.WriteVirtual(ctx, "firstItem", m.GetFirstItem()); _firstItemErr != nil {
		return errors.Wrap(_firstItemErr, "Error serializing 'firstItem' field")
	}
	// Virtual field
	lastItem := m.GetLastItem()
	_ = lastItem
	if _lastItemErr := writeBuffer.WriteVirtual(ctx, "lastItem", m.GetLastItem()); _lastItemErr != nil {
		return errors.Wrap(_lastItemErr, "Error serializing 'lastItem' field")
	}
	// Virtual field
	moreItems := m.GetMoreItems()
	_ = moreItems
	if _moreItemsErr := writeBuffer.WriteVirtual(ctx, "moreItems", m.GetMoreItems()); _moreItemsErr != nil {
		return errors.Wrap(_moreItemsErr, "Error serializing 'moreItems' field")
	}

	if popErr := writeBuffer.PopContext("BACnetResultFlagsTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetResultFlagsTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetResultFlagsTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetResultFlagsTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetResultFlagsTagged) IsBACnetResultFlagsTagged() {}

func (m *_BACnetResultFlagsTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetResultFlagsTagged) deepCopy() *_BACnetResultFlagsTagged {
	if m == nil {
		return nil
	}
	_BACnetResultFlagsTaggedCopy := &_BACnetResultFlagsTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Payload.DeepCopy().(BACnetTagPayloadBitString),
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetResultFlagsTaggedCopy
}

func (m *_BACnetResultFlagsTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

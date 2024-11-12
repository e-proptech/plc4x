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

// NPDUNetworkPriorityTagged is the corresponding interface of NPDUNetworkPriorityTagged
type NPDUNetworkPriorityTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() NPDUNetworkPriority
	// IsNPDUNetworkPriorityTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNPDUNetworkPriorityTagged()
	// CreateBuilder creates a NPDUNetworkPriorityTaggedBuilder
	CreateNPDUNetworkPriorityTaggedBuilder() NPDUNetworkPriorityTaggedBuilder
}

// _NPDUNetworkPriorityTagged is the data-structure of this message
type _NPDUNetworkPriorityTagged struct {
	Header BACnetTagHeader
	Value  NPDUNetworkPriority

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ NPDUNetworkPriorityTagged = (*_NPDUNetworkPriorityTagged)(nil)

// NewNPDUNetworkPriorityTagged factory function for _NPDUNetworkPriorityTagged
func NewNPDUNetworkPriorityTagged(header BACnetTagHeader, value NPDUNetworkPriority, tagNumber uint8, tagClass TagClass) *_NPDUNetworkPriorityTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for NPDUNetworkPriorityTagged must not be nil")
	}
	return &_NPDUNetworkPriorityTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NPDUNetworkPriorityTaggedBuilder is a builder for NPDUNetworkPriorityTagged
type NPDUNetworkPriorityTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value NPDUNetworkPriority) NPDUNetworkPriorityTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) NPDUNetworkPriorityTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) NPDUNetworkPriorityTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(NPDUNetworkPriority) NPDUNetworkPriorityTaggedBuilder
	// Build builds the NPDUNetworkPriorityTagged or returns an error if something is wrong
	Build() (NPDUNetworkPriorityTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NPDUNetworkPriorityTagged
}

// NewNPDUNetworkPriorityTaggedBuilder() creates a NPDUNetworkPriorityTaggedBuilder
func NewNPDUNetworkPriorityTaggedBuilder() NPDUNetworkPriorityTaggedBuilder {
	return &_NPDUNetworkPriorityTaggedBuilder{_NPDUNetworkPriorityTagged: new(_NPDUNetworkPriorityTagged)}
}

type _NPDUNetworkPriorityTaggedBuilder struct {
	*_NPDUNetworkPriorityTagged

	err *utils.MultiError
}

var _ (NPDUNetworkPriorityTaggedBuilder) = (*_NPDUNetworkPriorityTaggedBuilder)(nil)

func (b *_NPDUNetworkPriorityTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value NPDUNetworkPriority) NPDUNetworkPriorityTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_NPDUNetworkPriorityTaggedBuilder) WithHeader(header BACnetTagHeader) NPDUNetworkPriorityTaggedBuilder {
	b.Header = header
	return b
}

func (b *_NPDUNetworkPriorityTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) NPDUNetworkPriorityTaggedBuilder {
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

func (b *_NPDUNetworkPriorityTaggedBuilder) WithValue(value NPDUNetworkPriority) NPDUNetworkPriorityTaggedBuilder {
	b.Value = value
	return b
}

func (b *_NPDUNetworkPriorityTaggedBuilder) Build() (NPDUNetworkPriorityTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NPDUNetworkPriorityTagged.deepCopy(), nil
}

func (b *_NPDUNetworkPriorityTaggedBuilder) MustBuild() NPDUNetworkPriorityTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NPDUNetworkPriorityTaggedBuilder) DeepCopy() any {
	_copy := b.CreateNPDUNetworkPriorityTaggedBuilder().(*_NPDUNetworkPriorityTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNPDUNetworkPriorityTaggedBuilder creates a NPDUNetworkPriorityTaggedBuilder
func (b *_NPDUNetworkPriorityTagged) CreateNPDUNetworkPriorityTaggedBuilder() NPDUNetworkPriorityTaggedBuilder {
	if b == nil {
		return NewNPDUNetworkPriorityTaggedBuilder()
	}
	return &_NPDUNetworkPriorityTaggedBuilder{_NPDUNetworkPriorityTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NPDUNetworkPriorityTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_NPDUNetworkPriorityTagged) GetValue() NPDUNetworkPriority {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNPDUNetworkPriorityTagged(structType any) NPDUNetworkPriorityTagged {
	if casted, ok := structType.(NPDUNetworkPriorityTagged); ok {
		return casted
	}
	if casted, ok := structType.(*NPDUNetworkPriorityTagged); ok {
		return *casted
	}
	return nil
}

func (m *_NPDUNetworkPriorityTagged) GetTypeName() string {
	return "NPDUNetworkPriorityTagged"
}

func (m *_NPDUNetworkPriorityTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_NPDUNetworkPriorityTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NPDUNetworkPriorityTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (NPDUNetworkPriorityTagged, error) {
	return NPDUNetworkPriorityTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func NPDUNetworkPriorityTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (NPDUNetworkPriorityTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (NPDUNetworkPriorityTagged, error) {
		return NPDUNetworkPriorityTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func NPDUNetworkPriorityTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (NPDUNetworkPriorityTagged, error) {
	v, err := (&_NPDUNetworkPriorityTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_NPDUNetworkPriorityTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__nPDUNetworkPriorityTagged NPDUNetworkPriorityTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NPDUNetworkPriorityTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NPDUNetworkPriorityTagged")
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

	value, err := ReadManualField[NPDUNetworkPriority](ctx, "value", readBuffer, EnsureType[NPDUNetworkPriority](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), NPDUNetworkPriority_LIFE_SAVETY_MESSAGE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("NPDUNetworkPriorityTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NPDUNetworkPriorityTagged")
	}

	return m, nil
}

func (m *_NPDUNetworkPriorityTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NPDUNetworkPriorityTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NPDUNetworkPriorityTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NPDUNetworkPriorityTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[NPDUNetworkPriority](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("NPDUNetworkPriorityTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NPDUNetworkPriorityTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_NPDUNetworkPriorityTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_NPDUNetworkPriorityTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_NPDUNetworkPriorityTagged) IsNPDUNetworkPriorityTagged() {}

func (m *_NPDUNetworkPriorityTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NPDUNetworkPriorityTagged) deepCopy() *_NPDUNetworkPriorityTagged {
	if m == nil {
		return nil
	}
	_NPDUNetworkPriorityTaggedCopy := &_NPDUNetworkPriorityTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _NPDUNetworkPriorityTaggedCopy
}

func (m *_NPDUNetworkPriorityTagged) String() string {
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

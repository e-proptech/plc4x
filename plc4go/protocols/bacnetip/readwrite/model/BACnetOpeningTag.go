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

// BACnetOpeningTag is the corresponding interface of BACnetOpeningTag
type BACnetOpeningTag interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// IsBACnetOpeningTag is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetOpeningTag()
	// CreateBuilder creates a BACnetOpeningTagBuilder
	CreateBACnetOpeningTagBuilder() BACnetOpeningTagBuilder
}

// _BACnetOpeningTag is the data-structure of this message
type _BACnetOpeningTag struct {
	Header BACnetTagHeader

	// Arguments.
	TagNumberArgument uint8
}

var _ BACnetOpeningTag = (*_BACnetOpeningTag)(nil)

// NewBACnetOpeningTag factory function for _BACnetOpeningTag
func NewBACnetOpeningTag(header BACnetTagHeader, tagNumberArgument uint8) *_BACnetOpeningTag {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetOpeningTag must not be nil")
	}
	return &_BACnetOpeningTag{Header: header, TagNumberArgument: tagNumberArgument}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetOpeningTagBuilder is a builder for BACnetOpeningTag
type BACnetOpeningTagBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader) BACnetOpeningTagBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetOpeningTagBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetOpeningTagBuilder
	// Build builds the BACnetOpeningTag or returns an error if something is wrong
	Build() (BACnetOpeningTag, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetOpeningTag
}

// NewBACnetOpeningTagBuilder() creates a BACnetOpeningTagBuilder
func NewBACnetOpeningTagBuilder() BACnetOpeningTagBuilder {
	return &_BACnetOpeningTagBuilder{_BACnetOpeningTag: new(_BACnetOpeningTag)}
}

type _BACnetOpeningTagBuilder struct {
	*_BACnetOpeningTag

	err *utils.MultiError
}

var _ (BACnetOpeningTagBuilder) = (*_BACnetOpeningTagBuilder)(nil)

func (b *_BACnetOpeningTagBuilder) WithMandatoryFields(header BACnetTagHeader) BACnetOpeningTagBuilder {
	return b.WithHeader(header)
}

func (b *_BACnetOpeningTagBuilder) WithHeader(header BACnetTagHeader) BACnetOpeningTagBuilder {
	b.Header = header
	return b
}

func (b *_BACnetOpeningTagBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetOpeningTagBuilder {
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

func (b *_BACnetOpeningTagBuilder) Build() (BACnetOpeningTag, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetOpeningTag.deepCopy(), nil
}

func (b *_BACnetOpeningTagBuilder) MustBuild() BACnetOpeningTag {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetOpeningTagBuilder) DeepCopy() any {
	_copy := b.CreateBACnetOpeningTagBuilder().(*_BACnetOpeningTagBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetOpeningTagBuilder creates a BACnetOpeningTagBuilder
func (b *_BACnetOpeningTag) CreateBACnetOpeningTagBuilder() BACnetOpeningTagBuilder {
	if b == nil {
		return NewBACnetOpeningTagBuilder()
	}
	return &_BACnetOpeningTagBuilder{_BACnetOpeningTag: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOpeningTag) GetHeader() BACnetTagHeader {
	return m.Header
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetOpeningTag(structType any) BACnetOpeningTag {
	if casted, ok := structType.(BACnetOpeningTag); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOpeningTag); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOpeningTag) GetTypeName() string {
	return "BACnetOpeningTag"
}

func (m *_BACnetOpeningTag) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOpeningTag) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetOpeningTagParse(ctx context.Context, theBytes []byte, tagNumberArgument uint8) (BACnetOpeningTag, error) {
	return BACnetOpeningTagParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumberArgument)
}

func BACnetOpeningTagParseWithBufferProducer(tagNumberArgument uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetOpeningTag, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetOpeningTag, error) {
		return BACnetOpeningTagParseWithBuffer(ctx, readBuffer, tagNumberArgument)
	}
}

func BACnetOpeningTagParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumberArgument uint8) (BACnetOpeningTag, error) {
	v, err := (&_BACnetOpeningTag{TagNumberArgument: tagNumberArgument}).parse(ctx, readBuffer, tagNumberArgument)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetOpeningTag) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumberArgument uint8) (__bACnetOpeningTag BACnetOpeningTag, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOpeningTag")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetActualTagNumber()) == (tagNumberArgument))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	// Validation
	if !(bool((header.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "should be a context tag"})
	}

	// Validation
	if !(bool((header.GetLengthValueType()) == (6))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "opening tag should have a value of 6"})
	}

	if closeErr := readBuffer.CloseContext("BACnetOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOpeningTag")
	}

	return m, nil
}

func (m *_BACnetOpeningTag) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOpeningTag) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetOpeningTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetOpeningTag")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if popErr := writeBuffer.PopContext("BACnetOpeningTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetOpeningTag")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetOpeningTag) GetTagNumberArgument() uint8 {
	return m.TagNumberArgument
}

//
////

func (m *_BACnetOpeningTag) IsBACnetOpeningTag() {}

func (m *_BACnetOpeningTag) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetOpeningTag) deepCopy() *_BACnetOpeningTag {
	if m == nil {
		return nil
	}
	_BACnetOpeningTagCopy := &_BACnetOpeningTag{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.TagNumberArgument,
	}
	return _BACnetOpeningTagCopy
}

func (m *_BACnetOpeningTag) String() string {
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

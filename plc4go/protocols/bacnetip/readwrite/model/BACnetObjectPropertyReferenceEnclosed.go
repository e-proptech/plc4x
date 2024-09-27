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

// BACnetObjectPropertyReferenceEnclosed is the corresponding interface of BACnetObjectPropertyReferenceEnclosed
type BACnetObjectPropertyReferenceEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetObjectPropertyReference returns ObjectPropertyReference (property field)
	GetObjectPropertyReference() BACnetObjectPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetObjectPropertyReferenceEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetObjectPropertyReferenceEnclosed()
	// CreateBuilder creates a BACnetObjectPropertyReferenceEnclosedBuilder
	CreateBACnetObjectPropertyReferenceEnclosedBuilder() BACnetObjectPropertyReferenceEnclosedBuilder
}

// _BACnetObjectPropertyReferenceEnclosed is the data-structure of this message
type _BACnetObjectPropertyReferenceEnclosed struct {
	OpeningTag              BACnetOpeningTag
	ObjectPropertyReference BACnetObjectPropertyReference
	ClosingTag              BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetObjectPropertyReferenceEnclosed = (*_BACnetObjectPropertyReferenceEnclosed)(nil)

// NewBACnetObjectPropertyReferenceEnclosed factory function for _BACnetObjectPropertyReferenceEnclosed
func NewBACnetObjectPropertyReferenceEnclosed(openingTag BACnetOpeningTag, objectPropertyReference BACnetObjectPropertyReference, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetObjectPropertyReferenceEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetObjectPropertyReferenceEnclosed must not be nil")
	}
	if objectPropertyReference == nil {
		panic("objectPropertyReference of type BACnetObjectPropertyReference for BACnetObjectPropertyReferenceEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetObjectPropertyReferenceEnclosed must not be nil")
	}
	return &_BACnetObjectPropertyReferenceEnclosed{OpeningTag: openingTag, ObjectPropertyReference: objectPropertyReference, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetObjectPropertyReferenceEnclosedBuilder is a builder for BACnetObjectPropertyReferenceEnclosed
type BACnetObjectPropertyReferenceEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, objectPropertyReference BACnetObjectPropertyReference, closingTag BACnetClosingTag) BACnetObjectPropertyReferenceEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetObjectPropertyReferenceEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetObjectPropertyReferenceEnclosedBuilder
	// WithObjectPropertyReference adds ObjectPropertyReference (property field)
	WithObjectPropertyReference(BACnetObjectPropertyReference) BACnetObjectPropertyReferenceEnclosedBuilder
	// WithObjectPropertyReferenceBuilder adds ObjectPropertyReference (property field) which is build by the builder
	WithObjectPropertyReferenceBuilder(func(BACnetObjectPropertyReferenceBuilder) BACnetObjectPropertyReferenceBuilder) BACnetObjectPropertyReferenceEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetObjectPropertyReferenceEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetObjectPropertyReferenceEnclosedBuilder
	// Build builds the BACnetObjectPropertyReferenceEnclosed or returns an error if something is wrong
	Build() (BACnetObjectPropertyReferenceEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetObjectPropertyReferenceEnclosed
}

// NewBACnetObjectPropertyReferenceEnclosedBuilder() creates a BACnetObjectPropertyReferenceEnclosedBuilder
func NewBACnetObjectPropertyReferenceEnclosedBuilder() BACnetObjectPropertyReferenceEnclosedBuilder {
	return &_BACnetObjectPropertyReferenceEnclosedBuilder{_BACnetObjectPropertyReferenceEnclosed: new(_BACnetObjectPropertyReferenceEnclosed)}
}

type _BACnetObjectPropertyReferenceEnclosedBuilder struct {
	*_BACnetObjectPropertyReferenceEnclosed

	err *utils.MultiError
}

var _ (BACnetObjectPropertyReferenceEnclosedBuilder) = (*_BACnetObjectPropertyReferenceEnclosedBuilder)(nil)

func (b *_BACnetObjectPropertyReferenceEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, objectPropertyReference BACnetObjectPropertyReference, closingTag BACnetClosingTag) BACnetObjectPropertyReferenceEnclosedBuilder {
	return b.WithOpeningTag(openingTag).WithObjectPropertyReference(objectPropertyReference).WithClosingTag(closingTag)
}

func (b *_BACnetObjectPropertyReferenceEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetObjectPropertyReferenceEnclosedBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetObjectPropertyReferenceEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetObjectPropertyReferenceEnclosedBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetObjectPropertyReferenceEnclosedBuilder) WithObjectPropertyReference(objectPropertyReference BACnetObjectPropertyReference) BACnetObjectPropertyReferenceEnclosedBuilder {
	b.ObjectPropertyReference = objectPropertyReference
	return b
}

func (b *_BACnetObjectPropertyReferenceEnclosedBuilder) WithObjectPropertyReferenceBuilder(builderSupplier func(BACnetObjectPropertyReferenceBuilder) BACnetObjectPropertyReferenceBuilder) BACnetObjectPropertyReferenceEnclosedBuilder {
	builder := builderSupplier(b.ObjectPropertyReference.CreateBACnetObjectPropertyReferenceBuilder())
	var err error
	b.ObjectPropertyReference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetObjectPropertyReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetObjectPropertyReferenceEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetObjectPropertyReferenceEnclosedBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetObjectPropertyReferenceEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetObjectPropertyReferenceEnclosedBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetObjectPropertyReferenceEnclosedBuilder) Build() (BACnetObjectPropertyReferenceEnclosed, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.ObjectPropertyReference == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectPropertyReference' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetObjectPropertyReferenceEnclosed.deepCopy(), nil
}

func (b *_BACnetObjectPropertyReferenceEnclosedBuilder) MustBuild() BACnetObjectPropertyReferenceEnclosed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetObjectPropertyReferenceEnclosedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetObjectPropertyReferenceEnclosedBuilder().(*_BACnetObjectPropertyReferenceEnclosedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetObjectPropertyReferenceEnclosedBuilder creates a BACnetObjectPropertyReferenceEnclosedBuilder
func (b *_BACnetObjectPropertyReferenceEnclosed) CreateBACnetObjectPropertyReferenceEnclosedBuilder() BACnetObjectPropertyReferenceEnclosedBuilder {
	if b == nil {
		return NewBACnetObjectPropertyReferenceEnclosedBuilder()
	}
	return &_BACnetObjectPropertyReferenceEnclosedBuilder{_BACnetObjectPropertyReferenceEnclosed: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetObjectPropertyReferenceEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetObjectPropertyReference() BACnetObjectPropertyReference {
	return m.ObjectPropertyReference
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetObjectPropertyReferenceEnclosed(structType any) BACnetObjectPropertyReferenceEnclosed {
	if casted, ok := structType.(BACnetObjectPropertyReferenceEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetObjectPropertyReferenceEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetTypeName() string {
	return "BACnetObjectPropertyReferenceEnclosed"
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (objectPropertyReference)
	lengthInBits += m.ObjectPropertyReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetObjectPropertyReferenceEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetObjectPropertyReferenceEnclosed, error) {
	return BACnetObjectPropertyReferenceEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetObjectPropertyReferenceEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetObjectPropertyReferenceEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetObjectPropertyReferenceEnclosed, error) {
		return BACnetObjectPropertyReferenceEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetObjectPropertyReferenceEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetObjectPropertyReferenceEnclosed, error) {
	v, err := (&_BACnetObjectPropertyReferenceEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetObjectPropertyReferenceEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetObjectPropertyReferenceEnclosed BACnetObjectPropertyReferenceEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetObjectPropertyReferenceEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetObjectPropertyReferenceEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	objectPropertyReference, err := ReadSimpleField[BACnetObjectPropertyReference](ctx, "objectPropertyReference", ReadComplex[BACnetObjectPropertyReference](BACnetObjectPropertyReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectPropertyReference' field"))
	}
	m.ObjectPropertyReference = objectPropertyReference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetObjectPropertyReferenceEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetObjectPropertyReferenceEnclosed")
	}

	return m, nil
}

func (m *_BACnetObjectPropertyReferenceEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetObjectPropertyReferenceEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetObjectPropertyReferenceEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetObjectPropertyReferenceEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetObjectPropertyReference](ctx, "objectPropertyReference", m.GetObjectPropertyReference(), WriteComplex[BACnetObjectPropertyReference](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectPropertyReference' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetObjectPropertyReferenceEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetObjectPropertyReferenceEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetObjectPropertyReferenceEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetObjectPropertyReferenceEnclosed) IsBACnetObjectPropertyReferenceEnclosed() {}

func (m *_BACnetObjectPropertyReferenceEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetObjectPropertyReferenceEnclosed) deepCopy() *_BACnetObjectPropertyReferenceEnclosed {
	if m == nil {
		return nil
	}
	_BACnetObjectPropertyReferenceEnclosedCopy := &_BACnetObjectPropertyReferenceEnclosed{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.ObjectPropertyReference.DeepCopy().(BACnetObjectPropertyReference),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetObjectPropertyReferenceEnclosedCopy
}

func (m *_BACnetObjectPropertyReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

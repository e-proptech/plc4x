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

// BACnetSpecialEventListOfTimeValues is the corresponding interface of BACnetSpecialEventListOfTimeValues
type BACnetSpecialEventListOfTimeValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfTimeValues returns ListOfTimeValues (property field)
	GetListOfTimeValues() []BACnetTimeValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetSpecialEventListOfTimeValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetSpecialEventListOfTimeValues()
	// CreateBuilder creates a BACnetSpecialEventListOfTimeValuesBuilder
	CreateBACnetSpecialEventListOfTimeValuesBuilder() BACnetSpecialEventListOfTimeValuesBuilder
}

// _BACnetSpecialEventListOfTimeValues is the data-structure of this message
type _BACnetSpecialEventListOfTimeValues struct {
	OpeningTag       BACnetOpeningTag
	ListOfTimeValues []BACnetTimeValue
	ClosingTag       BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetSpecialEventListOfTimeValues = (*_BACnetSpecialEventListOfTimeValues)(nil)

// NewBACnetSpecialEventListOfTimeValues factory function for _BACnetSpecialEventListOfTimeValues
func NewBACnetSpecialEventListOfTimeValues(openingTag BACnetOpeningTag, listOfTimeValues []BACnetTimeValue, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetSpecialEventListOfTimeValues {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetSpecialEventListOfTimeValues must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetSpecialEventListOfTimeValues must not be nil")
	}
	return &_BACnetSpecialEventListOfTimeValues{OpeningTag: openingTag, ListOfTimeValues: listOfTimeValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetSpecialEventListOfTimeValuesBuilder is a builder for BACnetSpecialEventListOfTimeValues
type BACnetSpecialEventListOfTimeValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listOfTimeValues []BACnetTimeValue, closingTag BACnetClosingTag) BACnetSpecialEventListOfTimeValuesBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetSpecialEventListOfTimeValuesBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetSpecialEventListOfTimeValuesBuilder
	// WithListOfTimeValues adds ListOfTimeValues (property field)
	WithListOfTimeValues(...BACnetTimeValue) BACnetSpecialEventListOfTimeValuesBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetSpecialEventListOfTimeValuesBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetSpecialEventListOfTimeValuesBuilder
	// Build builds the BACnetSpecialEventListOfTimeValues or returns an error if something is wrong
	Build() (BACnetSpecialEventListOfTimeValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetSpecialEventListOfTimeValues
}

// NewBACnetSpecialEventListOfTimeValuesBuilder() creates a BACnetSpecialEventListOfTimeValuesBuilder
func NewBACnetSpecialEventListOfTimeValuesBuilder() BACnetSpecialEventListOfTimeValuesBuilder {
	return &_BACnetSpecialEventListOfTimeValuesBuilder{_BACnetSpecialEventListOfTimeValues: new(_BACnetSpecialEventListOfTimeValues)}
}

type _BACnetSpecialEventListOfTimeValuesBuilder struct {
	*_BACnetSpecialEventListOfTimeValues

	err *utils.MultiError
}

var _ (BACnetSpecialEventListOfTimeValuesBuilder) = (*_BACnetSpecialEventListOfTimeValuesBuilder)(nil)

func (b *_BACnetSpecialEventListOfTimeValuesBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listOfTimeValues []BACnetTimeValue, closingTag BACnetClosingTag) BACnetSpecialEventListOfTimeValuesBuilder {
	return b.WithOpeningTag(openingTag).WithListOfTimeValues(listOfTimeValues...).WithClosingTag(closingTag)
}

func (b *_BACnetSpecialEventListOfTimeValuesBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetSpecialEventListOfTimeValuesBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetSpecialEventListOfTimeValuesBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetSpecialEventListOfTimeValuesBuilder {
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

func (b *_BACnetSpecialEventListOfTimeValuesBuilder) WithListOfTimeValues(listOfTimeValues ...BACnetTimeValue) BACnetSpecialEventListOfTimeValuesBuilder {
	b.ListOfTimeValues = listOfTimeValues
	return b
}

func (b *_BACnetSpecialEventListOfTimeValuesBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetSpecialEventListOfTimeValuesBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetSpecialEventListOfTimeValuesBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetSpecialEventListOfTimeValuesBuilder {
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

func (b *_BACnetSpecialEventListOfTimeValuesBuilder) Build() (BACnetSpecialEventListOfTimeValues, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
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
	return b._BACnetSpecialEventListOfTimeValues.deepCopy(), nil
}

func (b *_BACnetSpecialEventListOfTimeValuesBuilder) MustBuild() BACnetSpecialEventListOfTimeValues {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetSpecialEventListOfTimeValuesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetSpecialEventListOfTimeValuesBuilder().(*_BACnetSpecialEventListOfTimeValuesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetSpecialEventListOfTimeValuesBuilder creates a BACnetSpecialEventListOfTimeValuesBuilder
func (b *_BACnetSpecialEventListOfTimeValues) CreateBACnetSpecialEventListOfTimeValuesBuilder() BACnetSpecialEventListOfTimeValuesBuilder {
	if b == nil {
		return NewBACnetSpecialEventListOfTimeValuesBuilder()
	}
	return &_BACnetSpecialEventListOfTimeValuesBuilder{_BACnetSpecialEventListOfTimeValues: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSpecialEventListOfTimeValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetSpecialEventListOfTimeValues) GetListOfTimeValues() []BACnetTimeValue {
	return m.ListOfTimeValues
}

func (m *_BACnetSpecialEventListOfTimeValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetSpecialEventListOfTimeValues(structType any) BACnetSpecialEventListOfTimeValues {
	if casted, ok := structType.(BACnetSpecialEventListOfTimeValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSpecialEventListOfTimeValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSpecialEventListOfTimeValues) GetTypeName() string {
	return "BACnetSpecialEventListOfTimeValues"
}

func (m *_BACnetSpecialEventListOfTimeValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfTimeValues) > 0 {
		for _, element := range m.ListOfTimeValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSpecialEventListOfTimeValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSpecialEventListOfTimeValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetSpecialEventListOfTimeValues, error) {
	return BACnetSpecialEventListOfTimeValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetSpecialEventListOfTimeValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSpecialEventListOfTimeValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSpecialEventListOfTimeValues, error) {
		return BACnetSpecialEventListOfTimeValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetSpecialEventListOfTimeValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetSpecialEventListOfTimeValues, error) {
	v, err := (&_BACnetSpecialEventListOfTimeValues{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetSpecialEventListOfTimeValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetSpecialEventListOfTimeValues BACnetSpecialEventListOfTimeValues, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSpecialEventListOfTimeValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEventListOfTimeValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfTimeValues, err := ReadTerminatedArrayField[BACnetTimeValue](ctx, "listOfTimeValues", ReadComplex[BACnetTimeValue](BACnetTimeValueParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfTimeValues' field"))
	}
	m.ListOfTimeValues = listOfTimeValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetSpecialEventListOfTimeValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEventListOfTimeValues")
	}

	return m, nil
}

func (m *_BACnetSpecialEventListOfTimeValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSpecialEventListOfTimeValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetSpecialEventListOfTimeValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEventListOfTimeValues")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfTimeValues", m.GetListOfTimeValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfTimeValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSpecialEventListOfTimeValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSpecialEventListOfTimeValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetSpecialEventListOfTimeValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetSpecialEventListOfTimeValues) IsBACnetSpecialEventListOfTimeValues() {}

func (m *_BACnetSpecialEventListOfTimeValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetSpecialEventListOfTimeValues) deepCopy() *_BACnetSpecialEventListOfTimeValues {
	if m == nil {
		return nil
	}
	_BACnetSpecialEventListOfTimeValuesCopy := &_BACnetSpecialEventListOfTimeValues{
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopySlice[BACnetTimeValue, BACnetTimeValue](m.ListOfTimeValues),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
		m.TagNumber,
	}
	return _BACnetSpecialEventListOfTimeValuesCopy
}

func (m *_BACnetSpecialEventListOfTimeValues) String() string {
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

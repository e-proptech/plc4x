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

// BACnetFaultParameterFaultCharacterString is the corresponding interface of BACnetFaultParameterFaultCharacterString
type BACnetFaultParameterFaultCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() BACnetFaultParameterFaultCharacterStringListOfFaultValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetFaultParameterFaultCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultCharacterString()
	// CreateBuilder creates a BACnetFaultParameterFaultCharacterStringBuilder
	CreateBACnetFaultParameterFaultCharacterStringBuilder() BACnetFaultParameterFaultCharacterStringBuilder
}

// _BACnetFaultParameterFaultCharacterString is the data-structure of this message
type _BACnetFaultParameterFaultCharacterString struct {
	BACnetFaultParameterContract
	OpeningTag        BACnetOpeningTag
	ListOfFaultValues BACnetFaultParameterFaultCharacterStringListOfFaultValues
	ClosingTag        BACnetClosingTag
}

var _ BACnetFaultParameterFaultCharacterString = (*_BACnetFaultParameterFaultCharacterString)(nil)
var _ BACnetFaultParameterRequirements = (*_BACnetFaultParameterFaultCharacterString)(nil)

// NewBACnetFaultParameterFaultCharacterString factory function for _BACnetFaultParameterFaultCharacterString
func NewBACnetFaultParameterFaultCharacterString(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultCharacterStringListOfFaultValues, closingTag BACnetClosingTag) *_BACnetFaultParameterFaultCharacterString {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetFaultParameterFaultCharacterString must not be nil")
	}
	if listOfFaultValues == nil {
		panic("listOfFaultValues of type BACnetFaultParameterFaultCharacterStringListOfFaultValues for BACnetFaultParameterFaultCharacterString must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetFaultParameterFaultCharacterString must not be nil")
	}
	_result := &_BACnetFaultParameterFaultCharacterString{
		BACnetFaultParameterContract: NewBACnetFaultParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		ListOfFaultValues:            listOfFaultValues,
		ClosingTag:                   closingTag,
	}
	_result.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultCharacterStringBuilder is a builder for BACnetFaultParameterFaultCharacterString
type BACnetFaultParameterFaultCharacterStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultCharacterStringListOfFaultValues, closingTag BACnetClosingTag) BACnetFaultParameterFaultCharacterStringBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetFaultParameterFaultCharacterStringBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetFaultParameterFaultCharacterStringBuilder
	// WithListOfFaultValues adds ListOfFaultValues (property field)
	WithListOfFaultValues(BACnetFaultParameterFaultCharacterStringListOfFaultValues) BACnetFaultParameterFaultCharacterStringBuilder
	// WithListOfFaultValuesBuilder adds ListOfFaultValues (property field) which is build by the builder
	WithListOfFaultValuesBuilder(func(BACnetFaultParameterFaultCharacterStringListOfFaultValuesBuilder) BACnetFaultParameterFaultCharacterStringListOfFaultValuesBuilder) BACnetFaultParameterFaultCharacterStringBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetFaultParameterFaultCharacterStringBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetFaultParameterFaultCharacterStringBuilder
	// Build builds the BACnetFaultParameterFaultCharacterString or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultCharacterString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultCharacterString
}

// NewBACnetFaultParameterFaultCharacterStringBuilder() creates a BACnetFaultParameterFaultCharacterStringBuilder
func NewBACnetFaultParameterFaultCharacterStringBuilder() BACnetFaultParameterFaultCharacterStringBuilder {
	return &_BACnetFaultParameterFaultCharacterStringBuilder{_BACnetFaultParameterFaultCharacterString: new(_BACnetFaultParameterFaultCharacterString)}
}

type _BACnetFaultParameterFaultCharacterStringBuilder struct {
	*_BACnetFaultParameterFaultCharacterString

	parentBuilder *_BACnetFaultParameterBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultCharacterStringBuilder) = (*_BACnetFaultParameterFaultCharacterStringBuilder)(nil)

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) setParent(contract BACnetFaultParameterContract) {
	b.BACnetFaultParameterContract = contract
}

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultCharacterStringListOfFaultValues, closingTag BACnetClosingTag) BACnetFaultParameterFaultCharacterStringBuilder {
	return b.WithOpeningTag(openingTag).WithListOfFaultValues(listOfFaultValues).WithClosingTag(closingTag)
}

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetFaultParameterFaultCharacterStringBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetFaultParameterFaultCharacterStringBuilder {
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

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) WithListOfFaultValues(listOfFaultValues BACnetFaultParameterFaultCharacterStringListOfFaultValues) BACnetFaultParameterFaultCharacterStringBuilder {
	b.ListOfFaultValues = listOfFaultValues
	return b
}

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) WithListOfFaultValuesBuilder(builderSupplier func(BACnetFaultParameterFaultCharacterStringListOfFaultValuesBuilder) BACnetFaultParameterFaultCharacterStringListOfFaultValuesBuilder) BACnetFaultParameterFaultCharacterStringBuilder {
	builder := builderSupplier(b.ListOfFaultValues.CreateBACnetFaultParameterFaultCharacterStringListOfFaultValuesBuilder())
	var err error
	b.ListOfFaultValues, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetFaultParameterFaultCharacterStringListOfFaultValuesBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetFaultParameterFaultCharacterStringBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetFaultParameterFaultCharacterStringBuilder {
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

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) Build() (BACnetFaultParameterFaultCharacterString, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.ListOfFaultValues == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'listOfFaultValues' not set"))
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
	return b._BACnetFaultParameterFaultCharacterString.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) MustBuild() BACnetFaultParameterFaultCharacterString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetFaultParameterFaultCharacterStringBuilder) Done() BACnetFaultParameterBuilder {
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) buildForBACnetFaultParameter() (BACnetFaultParameter, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultCharacterStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultCharacterStringBuilder().(*_BACnetFaultParameterFaultCharacterStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultCharacterStringBuilder creates a BACnetFaultParameterFaultCharacterStringBuilder
func (b *_BACnetFaultParameterFaultCharacterString) CreateBACnetFaultParameterFaultCharacterStringBuilder() BACnetFaultParameterFaultCharacterStringBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultCharacterStringBuilder()
	}
	return &_BACnetFaultParameterFaultCharacterStringBuilder{_BACnetFaultParameterFaultCharacterString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultCharacterString) GetParent() BACnetFaultParameterContract {
	return m.BACnetFaultParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultCharacterString) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultCharacterString) GetListOfFaultValues() BACnetFaultParameterFaultCharacterStringListOfFaultValues {
	return m.ListOfFaultValues
}

func (m *_BACnetFaultParameterFaultCharacterString) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultCharacterString(structType any) BACnetFaultParameterFaultCharacterString {
	if casted, ok := structType.(BACnetFaultParameterFaultCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultCharacterString) GetTypeName() string {
	return "BACnetFaultParameterFaultCharacterString"
}

func (m *_BACnetFaultParameterFaultCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterContract.(*_BACnetFaultParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (listOfFaultValues)
	lengthInBits += m.ListOfFaultValues.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameter) (__bACnetFaultParameterFaultCharacterString BACnetFaultParameterFaultCharacterString, err error) {
	m.BACnetFaultParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfFaultValues, err := ReadSimpleField[BACnetFaultParameterFaultCharacterStringListOfFaultValues](ctx, "listOfFaultValues", ReadComplex[BACnetFaultParameterFaultCharacterStringListOfFaultValues](BACnetFaultParameterFaultCharacterStringListOfFaultValuesParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfFaultValues' field"))
	}
	m.ListOfFaultValues = listOfFaultValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultCharacterString")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultCharacterString")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetFaultParameterFaultCharacterStringListOfFaultValues](ctx, "listOfFaultValues", m.GetListOfFaultValues(), WriteComplex[BACnetFaultParameterFaultCharacterStringListOfFaultValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfFaultValues' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultCharacterString")
		}
		return nil
	}
	return m.BACnetFaultParameterContract.(*_BACnetFaultParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultCharacterString) IsBACnetFaultParameterFaultCharacterString() {}

func (m *_BACnetFaultParameterFaultCharacterString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultCharacterString) deepCopy() *_BACnetFaultParameterFaultCharacterString {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultCharacterStringCopy := &_BACnetFaultParameterFaultCharacterString{
		m.BACnetFaultParameterContract.(*_BACnetFaultParameter).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetFaultParameterFaultCharacterStringListOfFaultValues](m.ListOfFaultValues),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
	}
	m.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = m
	return _BACnetFaultParameterFaultCharacterStringCopy
}

func (m *_BACnetFaultParameterFaultCharacterString) String() string {
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

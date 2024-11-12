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

// BACnetFaultParameterFaultState is the corresponding interface of BACnetFaultParameterFaultState
type BACnetFaultParameterFaultState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() BACnetFaultParameterFaultStateListOfFaultValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetFaultParameterFaultState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultState()
	// CreateBuilder creates a BACnetFaultParameterFaultStateBuilder
	CreateBACnetFaultParameterFaultStateBuilder() BACnetFaultParameterFaultStateBuilder
}

// _BACnetFaultParameterFaultState is the data-structure of this message
type _BACnetFaultParameterFaultState struct {
	BACnetFaultParameterContract
	OpeningTag        BACnetOpeningTag
	ListOfFaultValues BACnetFaultParameterFaultStateListOfFaultValues
	ClosingTag        BACnetClosingTag
}

var _ BACnetFaultParameterFaultState = (*_BACnetFaultParameterFaultState)(nil)
var _ BACnetFaultParameterRequirements = (*_BACnetFaultParameterFaultState)(nil)

// NewBACnetFaultParameterFaultState factory function for _BACnetFaultParameterFaultState
func NewBACnetFaultParameterFaultState(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultStateListOfFaultValues, closingTag BACnetClosingTag) *_BACnetFaultParameterFaultState {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetFaultParameterFaultState must not be nil")
	}
	if listOfFaultValues == nil {
		panic("listOfFaultValues of type BACnetFaultParameterFaultStateListOfFaultValues for BACnetFaultParameterFaultState must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetFaultParameterFaultState must not be nil")
	}
	_result := &_BACnetFaultParameterFaultState{
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

// BACnetFaultParameterFaultStateBuilder is a builder for BACnetFaultParameterFaultState
type BACnetFaultParameterFaultStateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultStateListOfFaultValues, closingTag BACnetClosingTag) BACnetFaultParameterFaultStateBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetFaultParameterFaultStateBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetFaultParameterFaultStateBuilder
	// WithListOfFaultValues adds ListOfFaultValues (property field)
	WithListOfFaultValues(BACnetFaultParameterFaultStateListOfFaultValues) BACnetFaultParameterFaultStateBuilder
	// WithListOfFaultValuesBuilder adds ListOfFaultValues (property field) which is build by the builder
	WithListOfFaultValuesBuilder(func(BACnetFaultParameterFaultStateListOfFaultValuesBuilder) BACnetFaultParameterFaultStateListOfFaultValuesBuilder) BACnetFaultParameterFaultStateBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetFaultParameterFaultStateBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetFaultParameterFaultStateBuilder
	// Build builds the BACnetFaultParameterFaultState or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultState, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultState
}

// NewBACnetFaultParameterFaultStateBuilder() creates a BACnetFaultParameterFaultStateBuilder
func NewBACnetFaultParameterFaultStateBuilder() BACnetFaultParameterFaultStateBuilder {
	return &_BACnetFaultParameterFaultStateBuilder{_BACnetFaultParameterFaultState: new(_BACnetFaultParameterFaultState)}
}

type _BACnetFaultParameterFaultStateBuilder struct {
	*_BACnetFaultParameterFaultState

	parentBuilder *_BACnetFaultParameterBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultStateBuilder) = (*_BACnetFaultParameterFaultStateBuilder)(nil)

func (b *_BACnetFaultParameterFaultStateBuilder) setParent(contract BACnetFaultParameterContract) {
	b.BACnetFaultParameterContract = contract
}

func (b *_BACnetFaultParameterFaultStateBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listOfFaultValues BACnetFaultParameterFaultStateListOfFaultValues, closingTag BACnetClosingTag) BACnetFaultParameterFaultStateBuilder {
	return b.WithOpeningTag(openingTag).WithListOfFaultValues(listOfFaultValues).WithClosingTag(closingTag)
}

func (b *_BACnetFaultParameterFaultStateBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetFaultParameterFaultStateBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetFaultParameterFaultStateBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetFaultParameterFaultStateBuilder {
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

func (b *_BACnetFaultParameterFaultStateBuilder) WithListOfFaultValues(listOfFaultValues BACnetFaultParameterFaultStateListOfFaultValues) BACnetFaultParameterFaultStateBuilder {
	b.ListOfFaultValues = listOfFaultValues
	return b
}

func (b *_BACnetFaultParameterFaultStateBuilder) WithListOfFaultValuesBuilder(builderSupplier func(BACnetFaultParameterFaultStateListOfFaultValuesBuilder) BACnetFaultParameterFaultStateListOfFaultValuesBuilder) BACnetFaultParameterFaultStateBuilder {
	builder := builderSupplier(b.ListOfFaultValues.CreateBACnetFaultParameterFaultStateListOfFaultValuesBuilder())
	var err error
	b.ListOfFaultValues, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetFaultParameterFaultStateListOfFaultValuesBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultStateBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetFaultParameterFaultStateBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetFaultParameterFaultStateBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetFaultParameterFaultStateBuilder {
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

func (b *_BACnetFaultParameterFaultStateBuilder) Build() (BACnetFaultParameterFaultState, error) {
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
	return b._BACnetFaultParameterFaultState.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultStateBuilder) MustBuild() BACnetFaultParameterFaultState {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetFaultParameterFaultStateBuilder) Done() BACnetFaultParameterBuilder {
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultStateBuilder) buildForBACnetFaultParameter() (BACnetFaultParameter, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultStateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultStateBuilder().(*_BACnetFaultParameterFaultStateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultStateBuilder creates a BACnetFaultParameterFaultStateBuilder
func (b *_BACnetFaultParameterFaultState) CreateBACnetFaultParameterFaultStateBuilder() BACnetFaultParameterFaultStateBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultStateBuilder()
	}
	return &_BACnetFaultParameterFaultStateBuilder{_BACnetFaultParameterFaultState: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultState) GetParent() BACnetFaultParameterContract {
	return m.BACnetFaultParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultState) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultState) GetListOfFaultValues() BACnetFaultParameterFaultStateListOfFaultValues {
	return m.ListOfFaultValues
}

func (m *_BACnetFaultParameterFaultState) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultState(structType any) BACnetFaultParameterFaultState {
	if casted, ok := structType.(BACnetFaultParameterFaultState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultState) GetTypeName() string {
	return "BACnetFaultParameterFaultState"
}

func (m *_BACnetFaultParameterFaultState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterContract.(*_BACnetFaultParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (listOfFaultValues)
	lengthInBits += m.ListOfFaultValues.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameter) (__bACnetFaultParameterFaultState BACnetFaultParameterFaultState, err error) {
	m.BACnetFaultParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfFaultValues, err := ReadSimpleField[BACnetFaultParameterFaultStateListOfFaultValues](ctx, "listOfFaultValues", ReadComplex[BACnetFaultParameterFaultStateListOfFaultValues](BACnetFaultParameterFaultStateListOfFaultValuesParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfFaultValues' field"))
	}
	m.ListOfFaultValues = listOfFaultValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultState")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultState")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetFaultParameterFaultStateListOfFaultValues](ctx, "listOfFaultValues", m.GetListOfFaultValues(), WriteComplex[BACnetFaultParameterFaultStateListOfFaultValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfFaultValues' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultState")
		}
		return nil
	}
	return m.BACnetFaultParameterContract.(*_BACnetFaultParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultState) IsBACnetFaultParameterFaultState() {}

func (m *_BACnetFaultParameterFaultState) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultState) deepCopy() *_BACnetFaultParameterFaultState {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultStateCopy := &_BACnetFaultParameterFaultState{
		m.BACnetFaultParameterContract.(*_BACnetFaultParameter).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetFaultParameterFaultStateListOfFaultValues](m.ListOfFaultValues),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
	}
	m.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = m
	return _BACnetFaultParameterFaultStateCopy
}

func (m *_BACnetFaultParameterFaultState) String() string {
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

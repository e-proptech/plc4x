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

// BACnetNotificationParametersComplexEventType is the corresponding interface of BACnetNotificationParametersComplexEventType
type BACnetNotificationParametersComplexEventType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() BACnetPropertyValues
	// IsBACnetNotificationParametersComplexEventType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersComplexEventType()
	// CreateBuilder creates a BACnetNotificationParametersComplexEventTypeBuilder
	CreateBACnetNotificationParametersComplexEventTypeBuilder() BACnetNotificationParametersComplexEventTypeBuilder
}

// _BACnetNotificationParametersComplexEventType is the data-structure of this message
type _BACnetNotificationParametersComplexEventType struct {
	BACnetNotificationParametersContract
	ListOfValues BACnetPropertyValues
}

var _ BACnetNotificationParametersComplexEventType = (*_BACnetNotificationParametersComplexEventType)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersComplexEventType)(nil)

// NewBACnetNotificationParametersComplexEventType factory function for _BACnetNotificationParametersComplexEventType
func NewBACnetNotificationParametersComplexEventType(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, listOfValues BACnetPropertyValues, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersComplexEventType {
	if listOfValues == nil {
		panic("listOfValues of type BACnetPropertyValues for BACnetNotificationParametersComplexEventType must not be nil")
	}
	_result := &_BACnetNotificationParametersComplexEventType{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		ListOfValues:                         listOfValues,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersComplexEventTypeBuilder is a builder for BACnetNotificationParametersComplexEventType
type BACnetNotificationParametersComplexEventTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(listOfValues BACnetPropertyValues) BACnetNotificationParametersComplexEventTypeBuilder
	// WithListOfValues adds ListOfValues (property field)
	WithListOfValues(BACnetPropertyValues) BACnetNotificationParametersComplexEventTypeBuilder
	// WithListOfValuesBuilder adds ListOfValues (property field) which is build by the builder
	WithListOfValuesBuilder(func(BACnetPropertyValuesBuilder) BACnetPropertyValuesBuilder) BACnetNotificationParametersComplexEventTypeBuilder
	// Build builds the BACnetNotificationParametersComplexEventType or returns an error if something is wrong
	Build() (BACnetNotificationParametersComplexEventType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersComplexEventType
}

// NewBACnetNotificationParametersComplexEventTypeBuilder() creates a BACnetNotificationParametersComplexEventTypeBuilder
func NewBACnetNotificationParametersComplexEventTypeBuilder() BACnetNotificationParametersComplexEventTypeBuilder {
	return &_BACnetNotificationParametersComplexEventTypeBuilder{_BACnetNotificationParametersComplexEventType: new(_BACnetNotificationParametersComplexEventType)}
}

type _BACnetNotificationParametersComplexEventTypeBuilder struct {
	*_BACnetNotificationParametersComplexEventType

	parentBuilder *_BACnetNotificationParametersBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersComplexEventTypeBuilder) = (*_BACnetNotificationParametersComplexEventTypeBuilder)(nil)

func (b *_BACnetNotificationParametersComplexEventTypeBuilder) setParent(contract BACnetNotificationParametersContract) {
	b.BACnetNotificationParametersContract = contract
}

func (b *_BACnetNotificationParametersComplexEventTypeBuilder) WithMandatoryFields(listOfValues BACnetPropertyValues) BACnetNotificationParametersComplexEventTypeBuilder {
	return b.WithListOfValues(listOfValues)
}

func (b *_BACnetNotificationParametersComplexEventTypeBuilder) WithListOfValues(listOfValues BACnetPropertyValues) BACnetNotificationParametersComplexEventTypeBuilder {
	b.ListOfValues = listOfValues
	return b
}

func (b *_BACnetNotificationParametersComplexEventTypeBuilder) WithListOfValuesBuilder(builderSupplier func(BACnetPropertyValuesBuilder) BACnetPropertyValuesBuilder) BACnetNotificationParametersComplexEventTypeBuilder {
	builder := builderSupplier(b.ListOfValues.CreateBACnetPropertyValuesBuilder())
	var err error
	b.ListOfValues, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPropertyValuesBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersComplexEventTypeBuilder) Build() (BACnetNotificationParametersComplexEventType, error) {
	if b.ListOfValues == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'listOfValues' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersComplexEventType.deepCopy(), nil
}

func (b *_BACnetNotificationParametersComplexEventTypeBuilder) MustBuild() BACnetNotificationParametersComplexEventType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetNotificationParametersComplexEventTypeBuilder) Done() BACnetNotificationParametersBuilder {
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersComplexEventTypeBuilder) buildForBACnetNotificationParameters() (BACnetNotificationParameters, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersComplexEventTypeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersComplexEventTypeBuilder().(*_BACnetNotificationParametersComplexEventTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersComplexEventTypeBuilder creates a BACnetNotificationParametersComplexEventTypeBuilder
func (b *_BACnetNotificationParametersComplexEventType) CreateBACnetNotificationParametersComplexEventTypeBuilder() BACnetNotificationParametersComplexEventTypeBuilder {
	if b == nil {
		return NewBACnetNotificationParametersComplexEventTypeBuilder()
	}
	return &_BACnetNotificationParametersComplexEventTypeBuilder{_BACnetNotificationParametersComplexEventType: b.deepCopy()}
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

func (m *_BACnetNotificationParametersComplexEventType) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersComplexEventType) GetListOfValues() BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersComplexEventType(structType any) BACnetNotificationParametersComplexEventType {
	if casted, ok := structType.(BACnetNotificationParametersComplexEventType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersComplexEventType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersComplexEventType) GetTypeName() string {
	return "BACnetNotificationParametersComplexEventType"
}

func (m *_BACnetNotificationParametersComplexEventType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersComplexEventType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersComplexEventType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersComplexEventType BACnetNotificationParametersComplexEventType, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersComplexEventType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersComplexEventType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	listOfValues, err := ReadSimpleField[BACnetPropertyValues](ctx, "listOfValues", ReadComplex[BACnetPropertyValues](BACnetPropertyValuesParseWithBufferProducer((uint8)(peekedTagNumber), (BACnetObjectType)(objectTypeArgument)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfValues' field"))
	}
	m.ListOfValues = listOfValues

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersComplexEventType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersComplexEventType")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersComplexEventType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersComplexEventType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersComplexEventType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersComplexEventType")
		}

		if err := WriteSimpleField[BACnetPropertyValues](ctx, "listOfValues", m.GetListOfValues(), WriteComplex[BACnetPropertyValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersComplexEventType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersComplexEventType")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersComplexEventType) IsBACnetNotificationParametersComplexEventType() {
}

func (m *_BACnetNotificationParametersComplexEventType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersComplexEventType) deepCopy() *_BACnetNotificationParametersComplexEventType {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersComplexEventTypeCopy := &_BACnetNotificationParametersComplexEventType{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		m.ListOfValues.DeepCopy().(BACnetPropertyValues),
	}
	m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersComplexEventTypeCopy
}

func (m *_BACnetNotificationParametersComplexEventType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

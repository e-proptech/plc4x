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

// BACnetConstructedDataPriorityArray is the corresponding interface of BACnetConstructedDataPriorityArray
type BACnetConstructedDataPriorityArray interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPriorityArray returns PriorityArray (property field)
	GetPriorityArray() BACnetPriorityArray
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetPriorityArray
	// IsBACnetConstructedDataPriorityArray is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPriorityArray()
	// CreateBuilder creates a BACnetConstructedDataPriorityArrayBuilder
	CreateBACnetConstructedDataPriorityArrayBuilder() BACnetConstructedDataPriorityArrayBuilder
}

// _BACnetConstructedDataPriorityArray is the data-structure of this message
type _BACnetConstructedDataPriorityArray struct {
	BACnetConstructedDataContract
	PriorityArray BACnetPriorityArray
}

var _ BACnetConstructedDataPriorityArray = (*_BACnetConstructedDataPriorityArray)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPriorityArray)(nil)

// NewBACnetConstructedDataPriorityArray factory function for _BACnetConstructedDataPriorityArray
func NewBACnetConstructedDataPriorityArray(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, priorityArray BACnetPriorityArray, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPriorityArray {
	if priorityArray == nil {
		panic("priorityArray of type BACnetPriorityArray for BACnetConstructedDataPriorityArray must not be nil")
	}
	_result := &_BACnetConstructedDataPriorityArray{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PriorityArray:                 priorityArray,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPriorityArrayBuilder is a builder for BACnetConstructedDataPriorityArray
type BACnetConstructedDataPriorityArrayBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(priorityArray BACnetPriorityArray) BACnetConstructedDataPriorityArrayBuilder
	// WithPriorityArray adds PriorityArray (property field)
	WithPriorityArray(BACnetPriorityArray) BACnetConstructedDataPriorityArrayBuilder
	// WithPriorityArrayBuilder adds PriorityArray (property field) which is build by the builder
	WithPriorityArrayBuilder(func(BACnetPriorityArrayBuilder) BACnetPriorityArrayBuilder) BACnetConstructedDataPriorityArrayBuilder
	// Build builds the BACnetConstructedDataPriorityArray or returns an error if something is wrong
	Build() (BACnetConstructedDataPriorityArray, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPriorityArray
}

// NewBACnetConstructedDataPriorityArrayBuilder() creates a BACnetConstructedDataPriorityArrayBuilder
func NewBACnetConstructedDataPriorityArrayBuilder() BACnetConstructedDataPriorityArrayBuilder {
	return &_BACnetConstructedDataPriorityArrayBuilder{_BACnetConstructedDataPriorityArray: new(_BACnetConstructedDataPriorityArray)}
}

type _BACnetConstructedDataPriorityArrayBuilder struct {
	*_BACnetConstructedDataPriorityArray

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPriorityArrayBuilder) = (*_BACnetConstructedDataPriorityArrayBuilder)(nil)

func (b *_BACnetConstructedDataPriorityArrayBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataPriorityArrayBuilder) WithMandatoryFields(priorityArray BACnetPriorityArray) BACnetConstructedDataPriorityArrayBuilder {
	return b.WithPriorityArray(priorityArray)
}

func (b *_BACnetConstructedDataPriorityArrayBuilder) WithPriorityArray(priorityArray BACnetPriorityArray) BACnetConstructedDataPriorityArrayBuilder {
	b.PriorityArray = priorityArray
	return b
}

func (b *_BACnetConstructedDataPriorityArrayBuilder) WithPriorityArrayBuilder(builderSupplier func(BACnetPriorityArrayBuilder) BACnetPriorityArrayBuilder) BACnetConstructedDataPriorityArrayBuilder {
	builder := builderSupplier(b.PriorityArray.CreateBACnetPriorityArrayBuilder())
	var err error
	b.PriorityArray, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPriorityArrayBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataPriorityArrayBuilder) Build() (BACnetConstructedDataPriorityArray, error) {
	if b.PriorityArray == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'priorityArray' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPriorityArray.deepCopy(), nil
}

func (b *_BACnetConstructedDataPriorityArrayBuilder) MustBuild() BACnetConstructedDataPriorityArray {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataPriorityArrayBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPriorityArrayBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPriorityArrayBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPriorityArrayBuilder().(*_BACnetConstructedDataPriorityArrayBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPriorityArrayBuilder creates a BACnetConstructedDataPriorityArrayBuilder
func (b *_BACnetConstructedDataPriorityArray) CreateBACnetConstructedDataPriorityArrayBuilder() BACnetConstructedDataPriorityArrayBuilder {
	if b == nil {
		return NewBACnetConstructedDataPriorityArrayBuilder()
	}
	return &_BACnetConstructedDataPriorityArrayBuilder{_BACnetConstructedDataPriorityArray: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPriorityArray) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPriorityArray) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRIORITY_ARRAY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPriorityArray) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPriorityArray) GetPriorityArray() BACnetPriorityArray {
	return m.PriorityArray
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPriorityArray) GetActualValue() BACnetPriorityArray {
	ctx := context.Background()
	_ = ctx
	return CastBACnetPriorityArray(m.GetPriorityArray())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPriorityArray(structType any) BACnetConstructedDataPriorityArray {
	if casted, ok := structType.(BACnetConstructedDataPriorityArray); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPriorityArray); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPriorityArray) GetTypeName() string {
	return "BACnetConstructedDataPriorityArray"
}

func (m *_BACnetConstructedDataPriorityArray) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (priorityArray)
	lengthInBits += m.PriorityArray.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPriorityArray) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPriorityArray) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPriorityArray BACnetConstructedDataPriorityArray, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPriorityArray"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPriorityArray")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	priorityArray, err := ReadSimpleField[BACnetPriorityArray](ctx, "priorityArray", ReadComplex[BACnetPriorityArray](BACnetPriorityArrayParseWithBufferProducer((BACnetObjectType)(objectTypeArgument), (uint8)(tagNumber), (BACnetTagPayloadUnsignedInteger)(arrayIndexArgument)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priorityArray' field"))
	}
	m.PriorityArray = priorityArray

	actualValue, err := ReadVirtualField[BACnetPriorityArray](ctx, "actualValue", (*BACnetPriorityArray)(nil), priorityArray)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPriorityArray"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPriorityArray")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPriorityArray) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPriorityArray) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPriorityArray"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPriorityArray")
		}

		if err := WriteSimpleField[BACnetPriorityArray](ctx, "priorityArray", m.GetPriorityArray(), WriteComplex[BACnetPriorityArray](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'priorityArray' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPriorityArray"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPriorityArray")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPriorityArray) IsBACnetConstructedDataPriorityArray() {}

func (m *_BACnetConstructedDataPriorityArray) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPriorityArray) deepCopy() *_BACnetConstructedDataPriorityArray {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPriorityArrayCopy := &_BACnetConstructedDataPriorityArray{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetPriorityArray](m.PriorityArray),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPriorityArrayCopy
}

func (m *_BACnetConstructedDataPriorityArray) String() string {
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

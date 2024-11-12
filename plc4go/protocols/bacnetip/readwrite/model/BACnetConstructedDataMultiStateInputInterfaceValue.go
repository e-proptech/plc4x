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

// BACnetConstructedDataMultiStateInputInterfaceValue is the corresponding interface of BACnetConstructedDataMultiStateInputInterfaceValue
type BACnetConstructedDataMultiStateInputInterfaceValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() BACnetOptionalBinaryPV
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalBinaryPV
	// IsBACnetConstructedDataMultiStateInputInterfaceValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMultiStateInputInterfaceValue()
	// CreateBuilder creates a BACnetConstructedDataMultiStateInputInterfaceValueBuilder
	CreateBACnetConstructedDataMultiStateInputInterfaceValueBuilder() BACnetConstructedDataMultiStateInputInterfaceValueBuilder
}

// _BACnetConstructedDataMultiStateInputInterfaceValue is the data-structure of this message
type _BACnetConstructedDataMultiStateInputInterfaceValue struct {
	BACnetConstructedDataContract
	InterfaceValue BACnetOptionalBinaryPV
}

var _ BACnetConstructedDataMultiStateInputInterfaceValue = (*_BACnetConstructedDataMultiStateInputInterfaceValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMultiStateInputInterfaceValue)(nil)

// NewBACnetConstructedDataMultiStateInputInterfaceValue factory function for _BACnetConstructedDataMultiStateInputInterfaceValue
func NewBACnetConstructedDataMultiStateInputInterfaceValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, interfaceValue BACnetOptionalBinaryPV, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMultiStateInputInterfaceValue {
	if interfaceValue == nil {
		panic("interfaceValue of type BACnetOptionalBinaryPV for BACnetConstructedDataMultiStateInputInterfaceValue must not be nil")
	}
	_result := &_BACnetConstructedDataMultiStateInputInterfaceValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InterfaceValue:                interfaceValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMultiStateInputInterfaceValueBuilder is a builder for BACnetConstructedDataMultiStateInputInterfaceValue
type BACnetConstructedDataMultiStateInputInterfaceValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(interfaceValue BACnetOptionalBinaryPV) BACnetConstructedDataMultiStateInputInterfaceValueBuilder
	// WithInterfaceValue adds InterfaceValue (property field)
	WithInterfaceValue(BACnetOptionalBinaryPV) BACnetConstructedDataMultiStateInputInterfaceValueBuilder
	// WithInterfaceValueBuilder adds InterfaceValue (property field) which is build by the builder
	WithInterfaceValueBuilder(func(BACnetOptionalBinaryPVBuilder) BACnetOptionalBinaryPVBuilder) BACnetConstructedDataMultiStateInputInterfaceValueBuilder
	// Build builds the BACnetConstructedDataMultiStateInputInterfaceValue or returns an error if something is wrong
	Build() (BACnetConstructedDataMultiStateInputInterfaceValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMultiStateInputInterfaceValue
}

// NewBACnetConstructedDataMultiStateInputInterfaceValueBuilder() creates a BACnetConstructedDataMultiStateInputInterfaceValueBuilder
func NewBACnetConstructedDataMultiStateInputInterfaceValueBuilder() BACnetConstructedDataMultiStateInputInterfaceValueBuilder {
	return &_BACnetConstructedDataMultiStateInputInterfaceValueBuilder{_BACnetConstructedDataMultiStateInputInterfaceValue: new(_BACnetConstructedDataMultiStateInputInterfaceValue)}
}

type _BACnetConstructedDataMultiStateInputInterfaceValueBuilder struct {
	*_BACnetConstructedDataMultiStateInputInterfaceValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMultiStateInputInterfaceValueBuilder) = (*_BACnetConstructedDataMultiStateInputInterfaceValueBuilder)(nil)

func (b *_BACnetConstructedDataMultiStateInputInterfaceValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataMultiStateInputInterfaceValueBuilder) WithMandatoryFields(interfaceValue BACnetOptionalBinaryPV) BACnetConstructedDataMultiStateInputInterfaceValueBuilder {
	return b.WithInterfaceValue(interfaceValue)
}

func (b *_BACnetConstructedDataMultiStateInputInterfaceValueBuilder) WithInterfaceValue(interfaceValue BACnetOptionalBinaryPV) BACnetConstructedDataMultiStateInputInterfaceValueBuilder {
	b.InterfaceValue = interfaceValue
	return b
}

func (b *_BACnetConstructedDataMultiStateInputInterfaceValueBuilder) WithInterfaceValueBuilder(builderSupplier func(BACnetOptionalBinaryPVBuilder) BACnetOptionalBinaryPVBuilder) BACnetConstructedDataMultiStateInputInterfaceValueBuilder {
	builder := builderSupplier(b.InterfaceValue.CreateBACnetOptionalBinaryPVBuilder())
	var err error
	b.InterfaceValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOptionalBinaryPVBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMultiStateInputInterfaceValueBuilder) Build() (BACnetConstructedDataMultiStateInputInterfaceValue, error) {
	if b.InterfaceValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'interfaceValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMultiStateInputInterfaceValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataMultiStateInputInterfaceValueBuilder) MustBuild() BACnetConstructedDataMultiStateInputInterfaceValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataMultiStateInputInterfaceValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMultiStateInputInterfaceValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMultiStateInputInterfaceValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMultiStateInputInterfaceValueBuilder().(*_BACnetConstructedDataMultiStateInputInterfaceValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMultiStateInputInterfaceValueBuilder creates a BACnetConstructedDataMultiStateInputInterfaceValueBuilder
func (b *_BACnetConstructedDataMultiStateInputInterfaceValue) CreateBACnetConstructedDataMultiStateInputInterfaceValueBuilder() BACnetConstructedDataMultiStateInputInterfaceValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataMultiStateInputInterfaceValueBuilder()
	}
	return &_BACnetConstructedDataMultiStateInputInterfaceValueBuilder{_BACnetConstructedDataMultiStateInputInterfaceValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_INPUT
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetInterfaceValue() BACnetOptionalBinaryPV {
	return m.InterfaceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetActualValue() BACnetOptionalBinaryPV {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalBinaryPV(m.GetInterfaceValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMultiStateInputInterfaceValue(structType any) BACnetConstructedDataMultiStateInputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataMultiStateInputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateInputInterfaceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataMultiStateInputInterfaceValue"
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMultiStateInputInterfaceValue BACnetConstructedDataMultiStateInputInterfaceValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateInputInterfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMultiStateInputInterfaceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceValue, err := ReadSimpleField[BACnetOptionalBinaryPV](ctx, "interfaceValue", ReadComplex[BACnetOptionalBinaryPV](BACnetOptionalBinaryPVParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceValue' field"))
	}
	m.InterfaceValue = interfaceValue

	actualValue, err := ReadVirtualField[BACnetOptionalBinaryPV](ctx, "actualValue", (*BACnetOptionalBinaryPV)(nil), interfaceValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateInputInterfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMultiStateInputInterfaceValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateInputInterfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMultiStateInputInterfaceValue")
		}

		if err := WriteSimpleField[BACnetOptionalBinaryPV](ctx, "interfaceValue", m.GetInterfaceValue(), WriteComplex[BACnetOptionalBinaryPV](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateInputInterfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMultiStateInputInterfaceValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) IsBACnetConstructedDataMultiStateInputInterfaceValue() {
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) deepCopy() *_BACnetConstructedDataMultiStateInputInterfaceValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMultiStateInputInterfaceValueCopy := &_BACnetConstructedDataMultiStateInputInterfaceValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetOptionalBinaryPV](m.InterfaceValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMultiStateInputInterfaceValueCopy
}

func (m *_BACnetConstructedDataMultiStateInputInterfaceValue) String() string {
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

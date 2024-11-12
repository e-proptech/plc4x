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

// BACnetConstructedDataIntegralConstantUnits is the corresponding interface of BACnetConstructedDataIntegralConstantUnits
type BACnetConstructedDataIntegralConstantUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetUnits returns Units (property field)
	GetUnits() BACnetEngineeringUnitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEngineeringUnitsTagged
	// IsBACnetConstructedDataIntegralConstantUnits is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIntegralConstantUnits()
	// CreateBuilder creates a BACnetConstructedDataIntegralConstantUnitsBuilder
	CreateBACnetConstructedDataIntegralConstantUnitsBuilder() BACnetConstructedDataIntegralConstantUnitsBuilder
}

// _BACnetConstructedDataIntegralConstantUnits is the data-structure of this message
type _BACnetConstructedDataIntegralConstantUnits struct {
	BACnetConstructedDataContract
	Units BACnetEngineeringUnitsTagged
}

var _ BACnetConstructedDataIntegralConstantUnits = (*_BACnetConstructedDataIntegralConstantUnits)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIntegralConstantUnits)(nil)

// NewBACnetConstructedDataIntegralConstantUnits factory function for _BACnetConstructedDataIntegralConstantUnits
func NewBACnetConstructedDataIntegralConstantUnits(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, units BACnetEngineeringUnitsTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegralConstantUnits {
	if units == nil {
		panic("units of type BACnetEngineeringUnitsTagged for BACnetConstructedDataIntegralConstantUnits must not be nil")
	}
	_result := &_BACnetConstructedDataIntegralConstantUnits{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Units:                         units,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIntegralConstantUnitsBuilder is a builder for BACnetConstructedDataIntegralConstantUnits
type BACnetConstructedDataIntegralConstantUnitsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(units BACnetEngineeringUnitsTagged) BACnetConstructedDataIntegralConstantUnitsBuilder
	// WithUnits adds Units (property field)
	WithUnits(BACnetEngineeringUnitsTagged) BACnetConstructedDataIntegralConstantUnitsBuilder
	// WithUnitsBuilder adds Units (property field) which is build by the builder
	WithUnitsBuilder(func(BACnetEngineeringUnitsTaggedBuilder) BACnetEngineeringUnitsTaggedBuilder) BACnetConstructedDataIntegralConstantUnitsBuilder
	// Build builds the BACnetConstructedDataIntegralConstantUnits or returns an error if something is wrong
	Build() (BACnetConstructedDataIntegralConstantUnits, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIntegralConstantUnits
}

// NewBACnetConstructedDataIntegralConstantUnitsBuilder() creates a BACnetConstructedDataIntegralConstantUnitsBuilder
func NewBACnetConstructedDataIntegralConstantUnitsBuilder() BACnetConstructedDataIntegralConstantUnitsBuilder {
	return &_BACnetConstructedDataIntegralConstantUnitsBuilder{_BACnetConstructedDataIntegralConstantUnits: new(_BACnetConstructedDataIntegralConstantUnits)}
}

type _BACnetConstructedDataIntegralConstantUnitsBuilder struct {
	*_BACnetConstructedDataIntegralConstantUnits

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIntegralConstantUnitsBuilder) = (*_BACnetConstructedDataIntegralConstantUnitsBuilder)(nil)

func (b *_BACnetConstructedDataIntegralConstantUnitsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIntegralConstantUnitsBuilder) WithMandatoryFields(units BACnetEngineeringUnitsTagged) BACnetConstructedDataIntegralConstantUnitsBuilder {
	return b.WithUnits(units)
}

func (b *_BACnetConstructedDataIntegralConstantUnitsBuilder) WithUnits(units BACnetEngineeringUnitsTagged) BACnetConstructedDataIntegralConstantUnitsBuilder {
	b.Units = units
	return b
}

func (b *_BACnetConstructedDataIntegralConstantUnitsBuilder) WithUnitsBuilder(builderSupplier func(BACnetEngineeringUnitsTaggedBuilder) BACnetEngineeringUnitsTaggedBuilder) BACnetConstructedDataIntegralConstantUnitsBuilder {
	builder := builderSupplier(b.Units.CreateBACnetEngineeringUnitsTaggedBuilder())
	var err error
	b.Units, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEngineeringUnitsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIntegralConstantUnitsBuilder) Build() (BACnetConstructedDataIntegralConstantUnits, error) {
	if b.Units == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'units' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIntegralConstantUnits.deepCopy(), nil
}

func (b *_BACnetConstructedDataIntegralConstantUnitsBuilder) MustBuild() BACnetConstructedDataIntegralConstantUnits {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataIntegralConstantUnitsBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIntegralConstantUnitsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIntegralConstantUnitsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIntegralConstantUnitsBuilder().(*_BACnetConstructedDataIntegralConstantUnitsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIntegralConstantUnitsBuilder creates a BACnetConstructedDataIntegralConstantUnitsBuilder
func (b *_BACnetConstructedDataIntegralConstantUnits) CreateBACnetConstructedDataIntegralConstantUnitsBuilder() BACnetConstructedDataIntegralConstantUnitsBuilder {
	if b == nil {
		return NewBACnetConstructedDataIntegralConstantUnitsBuilder()
	}
	return &_BACnetConstructedDataIntegralConstantUnitsBuilder{_BACnetConstructedDataIntegralConstantUnits: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegralConstantUnits) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIntegralConstantUnits) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTEGRAL_CONSTANT_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegralConstantUnits) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegralConstantUnits) GetUnits() BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegralConstantUnits) GetActualValue() BACnetEngineeringUnitsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEngineeringUnitsTagged(m.GetUnits())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegralConstantUnits(structType any) BACnetConstructedDataIntegralConstantUnits {
	if casted, ok := structType.(BACnetConstructedDataIntegralConstantUnits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegralConstantUnits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegralConstantUnits) GetTypeName() string {
	return "BACnetConstructedDataIntegralConstantUnits"
}

func (m *_BACnetConstructedDataIntegralConstantUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegralConstantUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIntegralConstantUnits) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIntegralConstantUnits BACnetConstructedDataIntegralConstantUnits, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegralConstantUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegralConstantUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	units, err := ReadSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", ReadComplex[BACnetEngineeringUnitsTagged](BACnetEngineeringUnitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'units' field"))
	}
	m.Units = units

	actualValue, err := ReadVirtualField[BACnetEngineeringUnitsTagged](ctx, "actualValue", (*BACnetEngineeringUnitsTagged)(nil), units)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegralConstantUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegralConstantUnits")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIntegralConstantUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegralConstantUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegralConstantUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegralConstantUnits")
		}

		if err := WriteSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", m.GetUnits(), WriteComplex[BACnetEngineeringUnitsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'units' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegralConstantUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegralConstantUnits")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegralConstantUnits) IsBACnetConstructedDataIntegralConstantUnits() {
}

func (m *_BACnetConstructedDataIntegralConstantUnits) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIntegralConstantUnits) deepCopy() *_BACnetConstructedDataIntegralConstantUnits {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIntegralConstantUnitsCopy := &_BACnetConstructedDataIntegralConstantUnits{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetEngineeringUnitsTagged](m.Units),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIntegralConstantUnitsCopy
}

func (m *_BACnetConstructedDataIntegralConstantUnits) String() string {
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

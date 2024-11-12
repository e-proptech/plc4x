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

// BACnetConstructedDataZoneTo is the corresponding interface of BACnetConstructedDataZoneTo
type BACnetConstructedDataZoneTo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetZoneTo returns ZoneTo (property field)
	GetZoneTo() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
	// IsBACnetConstructedDataZoneTo is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataZoneTo()
	// CreateBuilder creates a BACnetConstructedDataZoneToBuilder
	CreateBACnetConstructedDataZoneToBuilder() BACnetConstructedDataZoneToBuilder
}

// _BACnetConstructedDataZoneTo is the data-structure of this message
type _BACnetConstructedDataZoneTo struct {
	BACnetConstructedDataContract
	ZoneTo BACnetDeviceObjectReference
}

var _ BACnetConstructedDataZoneTo = (*_BACnetConstructedDataZoneTo)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataZoneTo)(nil)

// NewBACnetConstructedDataZoneTo factory function for _BACnetConstructedDataZoneTo
func NewBACnetConstructedDataZoneTo(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, zoneTo BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataZoneTo {
	if zoneTo == nil {
		panic("zoneTo of type BACnetDeviceObjectReference for BACnetConstructedDataZoneTo must not be nil")
	}
	_result := &_BACnetConstructedDataZoneTo{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ZoneTo:                        zoneTo,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataZoneToBuilder is a builder for BACnetConstructedDataZoneTo
type BACnetConstructedDataZoneToBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneTo BACnetDeviceObjectReference) BACnetConstructedDataZoneToBuilder
	// WithZoneTo adds ZoneTo (property field)
	WithZoneTo(BACnetDeviceObjectReference) BACnetConstructedDataZoneToBuilder
	// WithZoneToBuilder adds ZoneTo (property field) which is build by the builder
	WithZoneToBuilder(func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataZoneToBuilder
	// Build builds the BACnetConstructedDataZoneTo or returns an error if something is wrong
	Build() (BACnetConstructedDataZoneTo, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataZoneTo
}

// NewBACnetConstructedDataZoneToBuilder() creates a BACnetConstructedDataZoneToBuilder
func NewBACnetConstructedDataZoneToBuilder() BACnetConstructedDataZoneToBuilder {
	return &_BACnetConstructedDataZoneToBuilder{_BACnetConstructedDataZoneTo: new(_BACnetConstructedDataZoneTo)}
}

type _BACnetConstructedDataZoneToBuilder struct {
	*_BACnetConstructedDataZoneTo

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataZoneToBuilder) = (*_BACnetConstructedDataZoneToBuilder)(nil)

func (b *_BACnetConstructedDataZoneToBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataZoneToBuilder) WithMandatoryFields(zoneTo BACnetDeviceObjectReference) BACnetConstructedDataZoneToBuilder {
	return b.WithZoneTo(zoneTo)
}

func (b *_BACnetConstructedDataZoneToBuilder) WithZoneTo(zoneTo BACnetDeviceObjectReference) BACnetConstructedDataZoneToBuilder {
	b.ZoneTo = zoneTo
	return b
}

func (b *_BACnetConstructedDataZoneToBuilder) WithZoneToBuilder(builderSupplier func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataZoneToBuilder {
	builder := builderSupplier(b.ZoneTo.CreateBACnetDeviceObjectReferenceBuilder())
	var err error
	b.ZoneTo, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataZoneToBuilder) Build() (BACnetConstructedDataZoneTo, error) {
	if b.ZoneTo == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'zoneTo' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataZoneTo.deepCopy(), nil
}

func (b *_BACnetConstructedDataZoneToBuilder) MustBuild() BACnetConstructedDataZoneTo {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataZoneToBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataZoneToBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataZoneToBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataZoneToBuilder().(*_BACnetConstructedDataZoneToBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataZoneToBuilder creates a BACnetConstructedDataZoneToBuilder
func (b *_BACnetConstructedDataZoneTo) CreateBACnetConstructedDataZoneToBuilder() BACnetConstructedDataZoneToBuilder {
	if b == nil {
		return NewBACnetConstructedDataZoneToBuilder()
	}
	return &_BACnetConstructedDataZoneToBuilder{_BACnetConstructedDataZoneTo: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataZoneTo) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataZoneTo) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ZONE_TO
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataZoneTo) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataZoneTo) GetZoneTo() BACnetDeviceObjectReference {
	return m.ZoneTo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataZoneTo) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetZoneTo())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataZoneTo(structType any) BACnetConstructedDataZoneTo {
	if casted, ok := structType.(BACnetConstructedDataZoneTo); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataZoneTo); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataZoneTo) GetTypeName() string {
	return "BACnetConstructedDataZoneTo"
}

func (m *_BACnetConstructedDataZoneTo) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (zoneTo)
	lengthInBits += m.ZoneTo.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataZoneTo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataZoneTo) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataZoneTo BACnetConstructedDataZoneTo, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataZoneTo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataZoneTo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneTo, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "zoneTo", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneTo' field"))
	}
	m.ZoneTo = zoneTo

	actualValue, err := ReadVirtualField[BACnetDeviceObjectReference](ctx, "actualValue", (*BACnetDeviceObjectReference)(nil), zoneTo)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataZoneTo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataZoneTo")
	}

	return m, nil
}

func (m *_BACnetConstructedDataZoneTo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataZoneTo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataZoneTo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataZoneTo")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReference](ctx, "zoneTo", m.GetZoneTo(), WriteComplex[BACnetDeviceObjectReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneTo' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataZoneTo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataZoneTo")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataZoneTo) IsBACnetConstructedDataZoneTo() {}

func (m *_BACnetConstructedDataZoneTo) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataZoneTo) deepCopy() *_BACnetConstructedDataZoneTo {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataZoneToCopy := &_BACnetConstructedDataZoneTo{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDeviceObjectReference](m.ZoneTo),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataZoneToCopy
}

func (m *_BACnetConstructedDataZoneTo) String() string {
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

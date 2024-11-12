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

// BACnetConstructedDataBelongsTo is the corresponding interface of BACnetConstructedDataBelongsTo
type BACnetConstructedDataBelongsTo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetBelongsTo returns BelongsTo (property field)
	GetBelongsTo() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
	// IsBACnetConstructedDataBelongsTo is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBelongsTo()
	// CreateBuilder creates a BACnetConstructedDataBelongsToBuilder
	CreateBACnetConstructedDataBelongsToBuilder() BACnetConstructedDataBelongsToBuilder
}

// _BACnetConstructedDataBelongsTo is the data-structure of this message
type _BACnetConstructedDataBelongsTo struct {
	BACnetConstructedDataContract
	BelongsTo BACnetDeviceObjectReference
}

var _ BACnetConstructedDataBelongsTo = (*_BACnetConstructedDataBelongsTo)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBelongsTo)(nil)

// NewBACnetConstructedDataBelongsTo factory function for _BACnetConstructedDataBelongsTo
func NewBACnetConstructedDataBelongsTo(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, belongsTo BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBelongsTo {
	if belongsTo == nil {
		panic("belongsTo of type BACnetDeviceObjectReference for BACnetConstructedDataBelongsTo must not be nil")
	}
	_result := &_BACnetConstructedDataBelongsTo{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BelongsTo:                     belongsTo,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBelongsToBuilder is a builder for BACnetConstructedDataBelongsTo
type BACnetConstructedDataBelongsToBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(belongsTo BACnetDeviceObjectReference) BACnetConstructedDataBelongsToBuilder
	// WithBelongsTo adds BelongsTo (property field)
	WithBelongsTo(BACnetDeviceObjectReference) BACnetConstructedDataBelongsToBuilder
	// WithBelongsToBuilder adds BelongsTo (property field) which is build by the builder
	WithBelongsToBuilder(func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataBelongsToBuilder
	// Build builds the BACnetConstructedDataBelongsTo or returns an error if something is wrong
	Build() (BACnetConstructedDataBelongsTo, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBelongsTo
}

// NewBACnetConstructedDataBelongsToBuilder() creates a BACnetConstructedDataBelongsToBuilder
func NewBACnetConstructedDataBelongsToBuilder() BACnetConstructedDataBelongsToBuilder {
	return &_BACnetConstructedDataBelongsToBuilder{_BACnetConstructedDataBelongsTo: new(_BACnetConstructedDataBelongsTo)}
}

type _BACnetConstructedDataBelongsToBuilder struct {
	*_BACnetConstructedDataBelongsTo

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBelongsToBuilder) = (*_BACnetConstructedDataBelongsToBuilder)(nil)

func (b *_BACnetConstructedDataBelongsToBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataBelongsToBuilder) WithMandatoryFields(belongsTo BACnetDeviceObjectReference) BACnetConstructedDataBelongsToBuilder {
	return b.WithBelongsTo(belongsTo)
}

func (b *_BACnetConstructedDataBelongsToBuilder) WithBelongsTo(belongsTo BACnetDeviceObjectReference) BACnetConstructedDataBelongsToBuilder {
	b.BelongsTo = belongsTo
	return b
}

func (b *_BACnetConstructedDataBelongsToBuilder) WithBelongsToBuilder(builderSupplier func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataBelongsToBuilder {
	builder := builderSupplier(b.BelongsTo.CreateBACnetDeviceObjectReferenceBuilder())
	var err error
	b.BelongsTo, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBelongsToBuilder) Build() (BACnetConstructedDataBelongsTo, error) {
	if b.BelongsTo == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'belongsTo' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBelongsTo.deepCopy(), nil
}

func (b *_BACnetConstructedDataBelongsToBuilder) MustBuild() BACnetConstructedDataBelongsTo {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataBelongsToBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBelongsToBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBelongsToBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBelongsToBuilder().(*_BACnetConstructedDataBelongsToBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBelongsToBuilder creates a BACnetConstructedDataBelongsToBuilder
func (b *_BACnetConstructedDataBelongsTo) CreateBACnetConstructedDataBelongsToBuilder() BACnetConstructedDataBelongsToBuilder {
	if b == nil {
		return NewBACnetConstructedDataBelongsToBuilder()
	}
	return &_BACnetConstructedDataBelongsToBuilder{_BACnetConstructedDataBelongsTo: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBelongsTo) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBelongsTo) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BELONGS_TO
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBelongsTo) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBelongsTo) GetBelongsTo() BACnetDeviceObjectReference {
	return m.BelongsTo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBelongsTo) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetBelongsTo())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBelongsTo(structType any) BACnetConstructedDataBelongsTo {
	if casted, ok := structType.(BACnetConstructedDataBelongsTo); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBelongsTo); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBelongsTo) GetTypeName() string {
	return "BACnetConstructedDataBelongsTo"
}

func (m *_BACnetConstructedDataBelongsTo) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (belongsTo)
	lengthInBits += m.BelongsTo.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBelongsTo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBelongsTo) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBelongsTo BACnetConstructedDataBelongsTo, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBelongsTo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBelongsTo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	belongsTo, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "belongsTo", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'belongsTo' field"))
	}
	m.BelongsTo = belongsTo

	actualValue, err := ReadVirtualField[BACnetDeviceObjectReference](ctx, "actualValue", (*BACnetDeviceObjectReference)(nil), belongsTo)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBelongsTo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBelongsTo")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBelongsTo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBelongsTo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBelongsTo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBelongsTo")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReference](ctx, "belongsTo", m.GetBelongsTo(), WriteComplex[BACnetDeviceObjectReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'belongsTo' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBelongsTo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBelongsTo")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBelongsTo) IsBACnetConstructedDataBelongsTo() {}

func (m *_BACnetConstructedDataBelongsTo) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBelongsTo) deepCopy() *_BACnetConstructedDataBelongsTo {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBelongsToCopy := &_BACnetConstructedDataBelongsTo{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDeviceObjectReference](m.BelongsTo),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBelongsToCopy
}

func (m *_BACnetConstructedDataBelongsTo) String() string {
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

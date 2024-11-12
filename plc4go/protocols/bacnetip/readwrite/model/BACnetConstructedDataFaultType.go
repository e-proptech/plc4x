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

// BACnetConstructedDataFaultType is the corresponding interface of BACnetConstructedDataFaultType
type BACnetConstructedDataFaultType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFaultType returns FaultType (property field)
	GetFaultType() BACnetFaultTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetFaultTypeTagged
	// IsBACnetConstructedDataFaultType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFaultType()
	// CreateBuilder creates a BACnetConstructedDataFaultTypeBuilder
	CreateBACnetConstructedDataFaultTypeBuilder() BACnetConstructedDataFaultTypeBuilder
}

// _BACnetConstructedDataFaultType is the data-structure of this message
type _BACnetConstructedDataFaultType struct {
	BACnetConstructedDataContract
	FaultType BACnetFaultTypeTagged
}

var _ BACnetConstructedDataFaultType = (*_BACnetConstructedDataFaultType)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFaultType)(nil)

// NewBACnetConstructedDataFaultType factory function for _BACnetConstructedDataFaultType
func NewBACnetConstructedDataFaultType(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, faultType BACnetFaultTypeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFaultType {
	if faultType == nil {
		panic("faultType of type BACnetFaultTypeTagged for BACnetConstructedDataFaultType must not be nil")
	}
	_result := &_BACnetConstructedDataFaultType{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultType:                     faultType,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataFaultTypeBuilder is a builder for BACnetConstructedDataFaultType
type BACnetConstructedDataFaultTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(faultType BACnetFaultTypeTagged) BACnetConstructedDataFaultTypeBuilder
	// WithFaultType adds FaultType (property field)
	WithFaultType(BACnetFaultTypeTagged) BACnetConstructedDataFaultTypeBuilder
	// WithFaultTypeBuilder adds FaultType (property field) which is build by the builder
	WithFaultTypeBuilder(func(BACnetFaultTypeTaggedBuilder) BACnetFaultTypeTaggedBuilder) BACnetConstructedDataFaultTypeBuilder
	// Build builds the BACnetConstructedDataFaultType or returns an error if something is wrong
	Build() (BACnetConstructedDataFaultType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataFaultType
}

// NewBACnetConstructedDataFaultTypeBuilder() creates a BACnetConstructedDataFaultTypeBuilder
func NewBACnetConstructedDataFaultTypeBuilder() BACnetConstructedDataFaultTypeBuilder {
	return &_BACnetConstructedDataFaultTypeBuilder{_BACnetConstructedDataFaultType: new(_BACnetConstructedDataFaultType)}
}

type _BACnetConstructedDataFaultTypeBuilder struct {
	*_BACnetConstructedDataFaultType

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataFaultTypeBuilder) = (*_BACnetConstructedDataFaultTypeBuilder)(nil)

func (b *_BACnetConstructedDataFaultTypeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataFaultTypeBuilder) WithMandatoryFields(faultType BACnetFaultTypeTagged) BACnetConstructedDataFaultTypeBuilder {
	return b.WithFaultType(faultType)
}

func (b *_BACnetConstructedDataFaultTypeBuilder) WithFaultType(faultType BACnetFaultTypeTagged) BACnetConstructedDataFaultTypeBuilder {
	b.FaultType = faultType
	return b
}

func (b *_BACnetConstructedDataFaultTypeBuilder) WithFaultTypeBuilder(builderSupplier func(BACnetFaultTypeTaggedBuilder) BACnetFaultTypeTaggedBuilder) BACnetConstructedDataFaultTypeBuilder {
	builder := builderSupplier(b.FaultType.CreateBACnetFaultTypeTaggedBuilder())
	var err error
	b.FaultType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetFaultTypeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataFaultTypeBuilder) Build() (BACnetConstructedDataFaultType, error) {
	if b.FaultType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'faultType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataFaultType.deepCopy(), nil
}

func (b *_BACnetConstructedDataFaultTypeBuilder) MustBuild() BACnetConstructedDataFaultType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataFaultTypeBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataFaultTypeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataFaultTypeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataFaultTypeBuilder().(*_BACnetConstructedDataFaultTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataFaultTypeBuilder creates a BACnetConstructedDataFaultTypeBuilder
func (b *_BACnetConstructedDataFaultType) CreateBACnetConstructedDataFaultTypeBuilder() BACnetConstructedDataFaultTypeBuilder {
	if b == nil {
		return NewBACnetConstructedDataFaultTypeBuilder()
	}
	return &_BACnetConstructedDataFaultTypeBuilder{_BACnetConstructedDataFaultType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFaultType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFaultType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFaultType) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFaultType) GetFaultType() BACnetFaultTypeTagged {
	return m.FaultType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFaultType) GetActualValue() BACnetFaultTypeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetFaultTypeTagged(m.GetFaultType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFaultType(structType any) BACnetConstructedDataFaultType {
	if casted, ok := structType.(BACnetConstructedDataFaultType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFaultType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFaultType) GetTypeName() string {
	return "BACnetConstructedDataFaultType"
}

func (m *_BACnetConstructedDataFaultType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (faultType)
	lengthInBits += m.FaultType.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFaultType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFaultType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFaultType BACnetConstructedDataFaultType, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFaultType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFaultType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultType, err := ReadSimpleField[BACnetFaultTypeTagged](ctx, "faultType", ReadComplex[BACnetFaultTypeTagged](BACnetFaultTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultType' field"))
	}
	m.FaultType = faultType

	actualValue, err := ReadVirtualField[BACnetFaultTypeTagged](ctx, "actualValue", (*BACnetFaultTypeTagged)(nil), faultType)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFaultType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFaultType")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFaultType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFaultType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFaultType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFaultType")
		}

		if err := WriteSimpleField[BACnetFaultTypeTagged](ctx, "faultType", m.GetFaultType(), WriteComplex[BACnetFaultTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'faultType' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFaultType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFaultType")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFaultType) IsBACnetConstructedDataFaultType() {}

func (m *_BACnetConstructedDataFaultType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataFaultType) deepCopy() *_BACnetConstructedDataFaultType {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataFaultTypeCopy := &_BACnetConstructedDataFaultType{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetFaultTypeTagged](m.FaultType),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataFaultTypeCopy
}

func (m *_BACnetConstructedDataFaultType) String() string {
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

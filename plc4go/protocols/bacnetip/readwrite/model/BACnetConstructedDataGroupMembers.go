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

// BACnetConstructedDataGroupMembers is the corresponding interface of BACnetConstructedDataGroupMembers
type BACnetConstructedDataGroupMembers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetGroupMembers returns GroupMembers (property field)
	GetGroupMembers() []BACnetApplicationTagObjectIdentifier
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataGroupMembers is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataGroupMembers()
	// CreateBuilder creates a BACnetConstructedDataGroupMembersBuilder
	CreateBACnetConstructedDataGroupMembersBuilder() BACnetConstructedDataGroupMembersBuilder
}

// _BACnetConstructedDataGroupMembers is the data-structure of this message
type _BACnetConstructedDataGroupMembers struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	GroupMembers         []BACnetApplicationTagObjectIdentifier
}

var _ BACnetConstructedDataGroupMembers = (*_BACnetConstructedDataGroupMembers)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataGroupMembers)(nil)

// NewBACnetConstructedDataGroupMembers factory function for _BACnetConstructedDataGroupMembers
func NewBACnetConstructedDataGroupMembers(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, groupMembers []BACnetApplicationTagObjectIdentifier, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGroupMembers {
	_result := &_BACnetConstructedDataGroupMembers{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		GroupMembers:                  groupMembers,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataGroupMembersBuilder is a builder for BACnetConstructedDataGroupMembers
type BACnetConstructedDataGroupMembersBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(groupMembers []BACnetApplicationTagObjectIdentifier) BACnetConstructedDataGroupMembersBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGroupMembersBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataGroupMembersBuilder
	// WithGroupMembers adds GroupMembers (property field)
	WithGroupMembers(...BACnetApplicationTagObjectIdentifier) BACnetConstructedDataGroupMembersBuilder
	// Build builds the BACnetConstructedDataGroupMembers or returns an error if something is wrong
	Build() (BACnetConstructedDataGroupMembers, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataGroupMembers
}

// NewBACnetConstructedDataGroupMembersBuilder() creates a BACnetConstructedDataGroupMembersBuilder
func NewBACnetConstructedDataGroupMembersBuilder() BACnetConstructedDataGroupMembersBuilder {
	return &_BACnetConstructedDataGroupMembersBuilder{_BACnetConstructedDataGroupMembers: new(_BACnetConstructedDataGroupMembers)}
}

type _BACnetConstructedDataGroupMembersBuilder struct {
	*_BACnetConstructedDataGroupMembers

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataGroupMembersBuilder) = (*_BACnetConstructedDataGroupMembersBuilder)(nil)

func (b *_BACnetConstructedDataGroupMembersBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataGroupMembersBuilder) WithMandatoryFields(groupMembers []BACnetApplicationTagObjectIdentifier) BACnetConstructedDataGroupMembersBuilder {
	return b.WithGroupMembers(groupMembers...)
}

func (b *_BACnetConstructedDataGroupMembersBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataGroupMembersBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataGroupMembersBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataGroupMembersBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataGroupMembersBuilder) WithGroupMembers(groupMembers ...BACnetApplicationTagObjectIdentifier) BACnetConstructedDataGroupMembersBuilder {
	b.GroupMembers = groupMembers
	return b
}

func (b *_BACnetConstructedDataGroupMembersBuilder) Build() (BACnetConstructedDataGroupMembers, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataGroupMembers.deepCopy(), nil
}

func (b *_BACnetConstructedDataGroupMembersBuilder) MustBuild() BACnetConstructedDataGroupMembers {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataGroupMembersBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataGroupMembersBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataGroupMembersBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataGroupMembersBuilder().(*_BACnetConstructedDataGroupMembersBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataGroupMembersBuilder creates a BACnetConstructedDataGroupMembersBuilder
func (b *_BACnetConstructedDataGroupMembers) CreateBACnetConstructedDataGroupMembersBuilder() BACnetConstructedDataGroupMembersBuilder {
	if b == nil {
		return NewBACnetConstructedDataGroupMembersBuilder()
	}
	return &_BACnetConstructedDataGroupMembersBuilder{_BACnetConstructedDataGroupMembers: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGroupMembers) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataGroupMembers) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GROUP_MEMBERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGroupMembers) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGroupMembers) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataGroupMembers) GetGroupMembers() []BACnetApplicationTagObjectIdentifier {
	return m.GroupMembers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGroupMembers) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGroupMembers(structType any) BACnetConstructedDataGroupMembers {
	if casted, ok := structType.(BACnetConstructedDataGroupMembers); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupMembers); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGroupMembers) GetTypeName() string {
	return "BACnetConstructedDataGroupMembers"
}

func (m *_BACnetConstructedDataGroupMembers) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.GroupMembers) > 0 {
		for _, element := range m.GroupMembers {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataGroupMembers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataGroupMembers) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataGroupMembers BACnetConstructedDataGroupMembers, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupMembers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGroupMembers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	groupMembers, err := ReadTerminatedArrayField[BACnetApplicationTagObjectIdentifier](ctx, "groupMembers", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'groupMembers' field"))
	}
	m.GroupMembers = groupMembers

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupMembers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGroupMembers")
	}

	return m, nil
}

func (m *_BACnetConstructedDataGroupMembers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGroupMembers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupMembers"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGroupMembers")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "groupMembers", m.GetGroupMembers(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'groupMembers' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupMembers"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGroupMembers")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGroupMembers) IsBACnetConstructedDataGroupMembers() {}

func (m *_BACnetConstructedDataGroupMembers) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataGroupMembers) deepCopy() *_BACnetConstructedDataGroupMembers {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataGroupMembersCopy := &_BACnetConstructedDataGroupMembers{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetApplicationTagObjectIdentifier, BACnetApplicationTagObjectIdentifier](m.GroupMembers),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataGroupMembersCopy
}

func (m *_BACnetConstructedDataGroupMembers) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

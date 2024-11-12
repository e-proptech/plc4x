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

// BACnetConstructedDataNetworkPortMaxInfoFrames is the corresponding interface of BACnetConstructedDataNetworkPortMaxInfoFrames
type BACnetConstructedDataNetworkPortMaxInfoFrames interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxInfoFrames returns MaxInfoFrames (property field)
	GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataNetworkPortMaxInfoFrames is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNetworkPortMaxInfoFrames()
	// CreateBuilder creates a BACnetConstructedDataNetworkPortMaxInfoFramesBuilder
	CreateBACnetConstructedDataNetworkPortMaxInfoFramesBuilder() BACnetConstructedDataNetworkPortMaxInfoFramesBuilder
}

// _BACnetConstructedDataNetworkPortMaxInfoFrames is the data-structure of this message
type _BACnetConstructedDataNetworkPortMaxInfoFrames struct {
	BACnetConstructedDataContract
	MaxInfoFrames BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataNetworkPortMaxInfoFrames = (*_BACnetConstructedDataNetworkPortMaxInfoFrames)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNetworkPortMaxInfoFrames)(nil)

// NewBACnetConstructedDataNetworkPortMaxInfoFrames factory function for _BACnetConstructedDataNetworkPortMaxInfoFrames
func NewBACnetConstructedDataNetworkPortMaxInfoFrames(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxInfoFrames BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNetworkPortMaxInfoFrames {
	if maxInfoFrames == nil {
		panic("maxInfoFrames of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataNetworkPortMaxInfoFrames must not be nil")
	}
	_result := &_BACnetConstructedDataNetworkPortMaxInfoFrames{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxInfoFrames:                 maxInfoFrames,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataNetworkPortMaxInfoFramesBuilder is a builder for BACnetConstructedDataNetworkPortMaxInfoFrames
type BACnetConstructedDataNetworkPortMaxInfoFramesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxInfoFrames BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNetworkPortMaxInfoFramesBuilder
	// WithMaxInfoFrames adds MaxInfoFrames (property field)
	WithMaxInfoFrames(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNetworkPortMaxInfoFramesBuilder
	// WithMaxInfoFramesBuilder adds MaxInfoFrames (property field) which is build by the builder
	WithMaxInfoFramesBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataNetworkPortMaxInfoFramesBuilder
	// Build builds the BACnetConstructedDataNetworkPortMaxInfoFrames or returns an error if something is wrong
	Build() (BACnetConstructedDataNetworkPortMaxInfoFrames, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataNetworkPortMaxInfoFrames
}

// NewBACnetConstructedDataNetworkPortMaxInfoFramesBuilder() creates a BACnetConstructedDataNetworkPortMaxInfoFramesBuilder
func NewBACnetConstructedDataNetworkPortMaxInfoFramesBuilder() BACnetConstructedDataNetworkPortMaxInfoFramesBuilder {
	return &_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder{_BACnetConstructedDataNetworkPortMaxInfoFrames: new(_BACnetConstructedDataNetworkPortMaxInfoFrames)}
}

type _BACnetConstructedDataNetworkPortMaxInfoFramesBuilder struct {
	*_BACnetConstructedDataNetworkPortMaxInfoFrames

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataNetworkPortMaxInfoFramesBuilder) = (*_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder)(nil)

func (b *_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder) WithMandatoryFields(maxInfoFrames BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNetworkPortMaxInfoFramesBuilder {
	return b.WithMaxInfoFrames(maxInfoFrames)
}

func (b *_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder) WithMaxInfoFrames(maxInfoFrames BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNetworkPortMaxInfoFramesBuilder {
	b.MaxInfoFrames = maxInfoFrames
	return b
}

func (b *_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder) WithMaxInfoFramesBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataNetworkPortMaxInfoFramesBuilder {
	builder := builderSupplier(b.MaxInfoFrames.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.MaxInfoFrames, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder) Build() (BACnetConstructedDataNetworkPortMaxInfoFrames, error) {
	if b.MaxInfoFrames == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maxInfoFrames' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataNetworkPortMaxInfoFrames.deepCopy(), nil
}

func (b *_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder) MustBuild() BACnetConstructedDataNetworkPortMaxInfoFrames {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataNetworkPortMaxInfoFramesBuilder().(*_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataNetworkPortMaxInfoFramesBuilder creates a BACnetConstructedDataNetworkPortMaxInfoFramesBuilder
func (b *_BACnetConstructedDataNetworkPortMaxInfoFrames) CreateBACnetConstructedDataNetworkPortMaxInfoFramesBuilder() BACnetConstructedDataNetworkPortMaxInfoFramesBuilder {
	if b == nil {
		return NewBACnetConstructedDataNetworkPortMaxInfoFramesBuilder()
	}
	return &_BACnetConstructedDataNetworkPortMaxInfoFramesBuilder{_BACnetConstructedDataNetworkPortMaxInfoFrames: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_NETWORK_PORT
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_INFO_FRAMES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger {
	return m.MaxInfoFrames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxInfoFrames())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNetworkPortMaxInfoFrames(structType any) BACnetConstructedDataNetworkPortMaxInfoFrames {
	if casted, ok := structType.(BACnetConstructedDataNetworkPortMaxInfoFrames); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkPortMaxInfoFrames); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetTypeName() string {
	return "BACnetConstructedDataNetworkPortMaxInfoFrames"
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxInfoFrames)
	lengthInBits += m.MaxInfoFrames.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNetworkPortMaxInfoFrames BACnetConstructedDataNetworkPortMaxInfoFrames, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkPortMaxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkPortMaxInfoFrames")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxInfoFrames, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxInfoFrames", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxInfoFrames' field"))
	}
	m.MaxInfoFrames = maxInfoFrames

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), maxInfoFrames)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkPortMaxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkPortMaxInfoFrames")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkPortMaxInfoFrames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkPortMaxInfoFrames")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxInfoFrames", m.GetMaxInfoFrames(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxInfoFrames' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkPortMaxInfoFrames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkPortMaxInfoFrames")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) IsBACnetConstructedDataNetworkPortMaxInfoFrames() {
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) deepCopy() *_BACnetConstructedDataNetworkPortMaxInfoFrames {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataNetworkPortMaxInfoFramesCopy := &_BACnetConstructedDataNetworkPortMaxInfoFrames{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.MaxInfoFrames),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataNetworkPortMaxInfoFramesCopy
}

func (m *_BACnetConstructedDataNetworkPortMaxInfoFrames) String() string {
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

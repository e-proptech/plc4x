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

// BACnetConstructedDataIntegerValueResolution is the corresponding interface of BACnetConstructedDataIntegerValueResolution
type BACnetConstructedDataIntegerValueResolution interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetResolution returns Resolution (property field)
	GetResolution() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataIntegerValueResolution is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIntegerValueResolution()
	// CreateBuilder creates a BACnetConstructedDataIntegerValueResolutionBuilder
	CreateBACnetConstructedDataIntegerValueResolutionBuilder() BACnetConstructedDataIntegerValueResolutionBuilder
}

// _BACnetConstructedDataIntegerValueResolution is the data-structure of this message
type _BACnetConstructedDataIntegerValueResolution struct {
	BACnetConstructedDataContract
	Resolution BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataIntegerValueResolution = (*_BACnetConstructedDataIntegerValueResolution)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIntegerValueResolution)(nil)

// NewBACnetConstructedDataIntegerValueResolution factory function for _BACnetConstructedDataIntegerValueResolution
func NewBACnetConstructedDataIntegerValueResolution(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, resolution BACnetApplicationTagSignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueResolution {
	if resolution == nil {
		panic("resolution of type BACnetApplicationTagSignedInteger for BACnetConstructedDataIntegerValueResolution must not be nil")
	}
	_result := &_BACnetConstructedDataIntegerValueResolution{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Resolution:                    resolution,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIntegerValueResolutionBuilder is a builder for BACnetConstructedDataIntegerValueResolution
type BACnetConstructedDataIntegerValueResolutionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(resolution BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueResolutionBuilder
	// WithResolution adds Resolution (property field)
	WithResolution(BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueResolutionBuilder
	// WithResolutionBuilder adds Resolution (property field) which is build by the builder
	WithResolutionBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataIntegerValueResolutionBuilder
	// Build builds the BACnetConstructedDataIntegerValueResolution or returns an error if something is wrong
	Build() (BACnetConstructedDataIntegerValueResolution, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIntegerValueResolution
}

// NewBACnetConstructedDataIntegerValueResolutionBuilder() creates a BACnetConstructedDataIntegerValueResolutionBuilder
func NewBACnetConstructedDataIntegerValueResolutionBuilder() BACnetConstructedDataIntegerValueResolutionBuilder {
	return &_BACnetConstructedDataIntegerValueResolutionBuilder{_BACnetConstructedDataIntegerValueResolution: new(_BACnetConstructedDataIntegerValueResolution)}
}

type _BACnetConstructedDataIntegerValueResolutionBuilder struct {
	*_BACnetConstructedDataIntegerValueResolution

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIntegerValueResolutionBuilder) = (*_BACnetConstructedDataIntegerValueResolutionBuilder)(nil)

func (b *_BACnetConstructedDataIntegerValueResolutionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIntegerValueResolutionBuilder) WithMandatoryFields(resolution BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueResolutionBuilder {
	return b.WithResolution(resolution)
}

func (b *_BACnetConstructedDataIntegerValueResolutionBuilder) WithResolution(resolution BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueResolutionBuilder {
	b.Resolution = resolution
	return b
}

func (b *_BACnetConstructedDataIntegerValueResolutionBuilder) WithResolutionBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataIntegerValueResolutionBuilder {
	builder := builderSupplier(b.Resolution.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.Resolution, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIntegerValueResolutionBuilder) Build() (BACnetConstructedDataIntegerValueResolution, error) {
	if b.Resolution == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'resolution' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIntegerValueResolution.deepCopy(), nil
}

func (b *_BACnetConstructedDataIntegerValueResolutionBuilder) MustBuild() BACnetConstructedDataIntegerValueResolution {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataIntegerValueResolutionBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIntegerValueResolutionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIntegerValueResolutionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIntegerValueResolutionBuilder().(*_BACnetConstructedDataIntegerValueResolutionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIntegerValueResolutionBuilder creates a BACnetConstructedDataIntegerValueResolutionBuilder
func (b *_BACnetConstructedDataIntegerValueResolution) CreateBACnetConstructedDataIntegerValueResolutionBuilder() BACnetConstructedDataIntegerValueResolutionBuilder {
	if b == nil {
		return NewBACnetConstructedDataIntegerValueResolutionBuilder()
	}
	return &_BACnetConstructedDataIntegerValueResolutionBuilder{_BACnetConstructedDataIntegerValueResolution: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueResolution) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueResolution) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESOLUTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueResolution) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueResolution) GetResolution() BACnetApplicationTagSignedInteger {
	return m.Resolution
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueResolution) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetResolution())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueResolution(structType any) BACnetConstructedDataIntegerValueResolution {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueResolution); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueResolution); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueResolution) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueResolution"
}

func (m *_BACnetConstructedDataIntegerValueResolution) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (resolution)
	lengthInBits += m.Resolution.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueResolution) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIntegerValueResolution) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIntegerValueResolution BACnetConstructedDataIntegerValueResolution, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueResolution"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueResolution")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	resolution, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "resolution", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resolution' field"))
	}
	m.Resolution = resolution

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), resolution)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueResolution"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueResolution")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIntegerValueResolution) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegerValueResolution) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueResolution"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueResolution")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "resolution", m.GetResolution(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'resolution' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueResolution"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueResolution")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueResolution) IsBACnetConstructedDataIntegerValueResolution() {
}

func (m *_BACnetConstructedDataIntegerValueResolution) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIntegerValueResolution) deepCopy() *_BACnetConstructedDataIntegerValueResolution {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIntegerValueResolutionCopy := &_BACnetConstructedDataIntegerValueResolution{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.Resolution),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIntegerValueResolutionCopy
}

func (m *_BACnetConstructedDataIntegerValueResolution) String() string {
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

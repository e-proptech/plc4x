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

// BACnetConstructedDataLargeAnalogValueDeadband is the corresponding interface of BACnetConstructedDataLargeAnalogValueDeadband
type BACnetConstructedDataLargeAnalogValueDeadband interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
	// IsBACnetConstructedDataLargeAnalogValueDeadband is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLargeAnalogValueDeadband()
	// CreateBuilder creates a BACnetConstructedDataLargeAnalogValueDeadbandBuilder
	CreateBACnetConstructedDataLargeAnalogValueDeadbandBuilder() BACnetConstructedDataLargeAnalogValueDeadbandBuilder
}

// _BACnetConstructedDataLargeAnalogValueDeadband is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueDeadband struct {
	BACnetConstructedDataContract
	Deadband BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValueDeadband = (*_BACnetConstructedDataLargeAnalogValueDeadband)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLargeAnalogValueDeadband)(nil)

// NewBACnetConstructedDataLargeAnalogValueDeadband factory function for _BACnetConstructedDataLargeAnalogValueDeadband
func NewBACnetConstructedDataLargeAnalogValueDeadband(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, deadband BACnetApplicationTagDouble, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueDeadband {
	if deadband == nil {
		panic("deadband of type BACnetApplicationTagDouble for BACnetConstructedDataLargeAnalogValueDeadband must not be nil")
	}
	_result := &_BACnetConstructedDataLargeAnalogValueDeadband{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Deadband:                      deadband,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLargeAnalogValueDeadbandBuilder is a builder for BACnetConstructedDataLargeAnalogValueDeadband
type BACnetConstructedDataLargeAnalogValueDeadbandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(deadband BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueDeadbandBuilder
	// WithDeadband adds Deadband (property field)
	WithDeadband(BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueDeadbandBuilder
	// WithDeadbandBuilder adds Deadband (property field) which is build by the builder
	WithDeadbandBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueDeadbandBuilder
	// Build builds the BACnetConstructedDataLargeAnalogValueDeadband or returns an error if something is wrong
	Build() (BACnetConstructedDataLargeAnalogValueDeadband, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLargeAnalogValueDeadband
}

// NewBACnetConstructedDataLargeAnalogValueDeadbandBuilder() creates a BACnetConstructedDataLargeAnalogValueDeadbandBuilder
func NewBACnetConstructedDataLargeAnalogValueDeadbandBuilder() BACnetConstructedDataLargeAnalogValueDeadbandBuilder {
	return &_BACnetConstructedDataLargeAnalogValueDeadbandBuilder{_BACnetConstructedDataLargeAnalogValueDeadband: new(_BACnetConstructedDataLargeAnalogValueDeadband)}
}

type _BACnetConstructedDataLargeAnalogValueDeadbandBuilder struct {
	*_BACnetConstructedDataLargeAnalogValueDeadband

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLargeAnalogValueDeadbandBuilder) = (*_BACnetConstructedDataLargeAnalogValueDeadbandBuilder)(nil)

func (b *_BACnetConstructedDataLargeAnalogValueDeadbandBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLargeAnalogValueDeadbandBuilder) WithMandatoryFields(deadband BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueDeadbandBuilder {
	return b.WithDeadband(deadband)
}

func (b *_BACnetConstructedDataLargeAnalogValueDeadbandBuilder) WithDeadband(deadband BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueDeadbandBuilder {
	b.Deadband = deadband
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueDeadbandBuilder) WithDeadbandBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueDeadbandBuilder {
	builder := builderSupplier(b.Deadband.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	b.Deadband, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueDeadbandBuilder) Build() (BACnetConstructedDataLargeAnalogValueDeadband, error) {
	if b.Deadband == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deadband' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLargeAnalogValueDeadband.deepCopy(), nil
}

func (b *_BACnetConstructedDataLargeAnalogValueDeadbandBuilder) MustBuild() BACnetConstructedDataLargeAnalogValueDeadband {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLargeAnalogValueDeadbandBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLargeAnalogValueDeadbandBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLargeAnalogValueDeadbandBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLargeAnalogValueDeadbandBuilder().(*_BACnetConstructedDataLargeAnalogValueDeadbandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLargeAnalogValueDeadbandBuilder creates a BACnetConstructedDataLargeAnalogValueDeadbandBuilder
func (b *_BACnetConstructedDataLargeAnalogValueDeadband) CreateBACnetConstructedDataLargeAnalogValueDeadbandBuilder() BACnetConstructedDataLargeAnalogValueDeadbandBuilder {
	if b == nil {
		return NewBACnetConstructedDataLargeAnalogValueDeadbandBuilder()
	}
	return &_BACnetConstructedDataLargeAnalogValueDeadbandBuilder{_BACnetConstructedDataLargeAnalogValueDeadband: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEADBAND
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetDeadband() BACnetApplicationTagDouble {
	return m.Deadband
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetDeadband())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueDeadband(structType any) BACnetConstructedDataLargeAnalogValueDeadband {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueDeadband); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueDeadband); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueDeadband"
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLargeAnalogValueDeadband BACnetConstructedDataLargeAnalogValueDeadband, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueDeadband"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueDeadband")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deadband, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "deadband", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}
	m.Deadband = deadband

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), deadband)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueDeadband"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueDeadband")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueDeadband"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueDeadband")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueDeadband"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueDeadband")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) IsBACnetConstructedDataLargeAnalogValueDeadband() {
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) deepCopy() *_BACnetConstructedDataLargeAnalogValueDeadband {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLargeAnalogValueDeadbandCopy := &_BACnetConstructedDataLargeAnalogValueDeadband{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Deadband.DeepCopy().(BACnetApplicationTagDouble),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLargeAnalogValueDeadbandCopy
}

func (m *_BACnetConstructedDataLargeAnalogValueDeadband) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

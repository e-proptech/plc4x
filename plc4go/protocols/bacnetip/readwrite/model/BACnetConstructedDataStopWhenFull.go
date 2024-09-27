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

// BACnetConstructedDataStopWhenFull is the corresponding interface of BACnetConstructedDataStopWhenFull
type BACnetConstructedDataStopWhenFull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetStopWhenFull returns StopWhenFull (property field)
	GetStopWhenFull() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataStopWhenFull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataStopWhenFull()
	// CreateBuilder creates a BACnetConstructedDataStopWhenFullBuilder
	CreateBACnetConstructedDataStopWhenFullBuilder() BACnetConstructedDataStopWhenFullBuilder
}

// _BACnetConstructedDataStopWhenFull is the data-structure of this message
type _BACnetConstructedDataStopWhenFull struct {
	BACnetConstructedDataContract
	StopWhenFull BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataStopWhenFull = (*_BACnetConstructedDataStopWhenFull)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataStopWhenFull)(nil)

// NewBACnetConstructedDataStopWhenFull factory function for _BACnetConstructedDataStopWhenFull
func NewBACnetConstructedDataStopWhenFull(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, stopWhenFull BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStopWhenFull {
	if stopWhenFull == nil {
		panic("stopWhenFull of type BACnetApplicationTagBoolean for BACnetConstructedDataStopWhenFull must not be nil")
	}
	_result := &_BACnetConstructedDataStopWhenFull{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		StopWhenFull:                  stopWhenFull,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataStopWhenFullBuilder is a builder for BACnetConstructedDataStopWhenFull
type BACnetConstructedDataStopWhenFullBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(stopWhenFull BACnetApplicationTagBoolean) BACnetConstructedDataStopWhenFullBuilder
	// WithStopWhenFull adds StopWhenFull (property field)
	WithStopWhenFull(BACnetApplicationTagBoolean) BACnetConstructedDataStopWhenFullBuilder
	// WithStopWhenFullBuilder adds StopWhenFull (property field) which is build by the builder
	WithStopWhenFullBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataStopWhenFullBuilder
	// Build builds the BACnetConstructedDataStopWhenFull or returns an error if something is wrong
	Build() (BACnetConstructedDataStopWhenFull, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataStopWhenFull
}

// NewBACnetConstructedDataStopWhenFullBuilder() creates a BACnetConstructedDataStopWhenFullBuilder
func NewBACnetConstructedDataStopWhenFullBuilder() BACnetConstructedDataStopWhenFullBuilder {
	return &_BACnetConstructedDataStopWhenFullBuilder{_BACnetConstructedDataStopWhenFull: new(_BACnetConstructedDataStopWhenFull)}
}

type _BACnetConstructedDataStopWhenFullBuilder struct {
	*_BACnetConstructedDataStopWhenFull

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataStopWhenFullBuilder) = (*_BACnetConstructedDataStopWhenFullBuilder)(nil)

func (b *_BACnetConstructedDataStopWhenFullBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataStopWhenFullBuilder) WithMandatoryFields(stopWhenFull BACnetApplicationTagBoolean) BACnetConstructedDataStopWhenFullBuilder {
	return b.WithStopWhenFull(stopWhenFull)
}

func (b *_BACnetConstructedDataStopWhenFullBuilder) WithStopWhenFull(stopWhenFull BACnetApplicationTagBoolean) BACnetConstructedDataStopWhenFullBuilder {
	b.StopWhenFull = stopWhenFull
	return b
}

func (b *_BACnetConstructedDataStopWhenFullBuilder) WithStopWhenFullBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataStopWhenFullBuilder {
	builder := builderSupplier(b.StopWhenFull.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.StopWhenFull, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataStopWhenFullBuilder) Build() (BACnetConstructedDataStopWhenFull, error) {
	if b.StopWhenFull == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'stopWhenFull' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataStopWhenFull.deepCopy(), nil
}

func (b *_BACnetConstructedDataStopWhenFullBuilder) MustBuild() BACnetConstructedDataStopWhenFull {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataStopWhenFullBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataStopWhenFullBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataStopWhenFullBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataStopWhenFullBuilder().(*_BACnetConstructedDataStopWhenFullBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataStopWhenFullBuilder creates a BACnetConstructedDataStopWhenFullBuilder
func (b *_BACnetConstructedDataStopWhenFull) CreateBACnetConstructedDataStopWhenFullBuilder() BACnetConstructedDataStopWhenFullBuilder {
	if b == nil {
		return NewBACnetConstructedDataStopWhenFullBuilder()
	}
	return &_BACnetConstructedDataStopWhenFullBuilder{_BACnetConstructedDataStopWhenFull: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStopWhenFull) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STOP_WHEN_FULL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetStopWhenFull() BACnetApplicationTagBoolean {
	return m.StopWhenFull
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStopWhenFull) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetStopWhenFull())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStopWhenFull(structType any) BACnetConstructedDataStopWhenFull {
	if casted, ok := structType.(BACnetConstructedDataStopWhenFull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStopWhenFull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStopWhenFull) GetTypeName() string {
	return "BACnetConstructedDataStopWhenFull"
}

func (m *_BACnetConstructedDataStopWhenFull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (stopWhenFull)
	lengthInBits += m.StopWhenFull.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStopWhenFull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStopWhenFull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStopWhenFull BACnetConstructedDataStopWhenFull, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStopWhenFull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStopWhenFull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	stopWhenFull, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "stopWhenFull", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stopWhenFull' field"))
	}
	m.StopWhenFull = stopWhenFull

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), stopWhenFull)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStopWhenFull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStopWhenFull")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStopWhenFull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStopWhenFull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStopWhenFull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStopWhenFull")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "stopWhenFull", m.GetStopWhenFull(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'stopWhenFull' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStopWhenFull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStopWhenFull")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStopWhenFull) IsBACnetConstructedDataStopWhenFull() {}

func (m *_BACnetConstructedDataStopWhenFull) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataStopWhenFull) deepCopy() *_BACnetConstructedDataStopWhenFull {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataStopWhenFullCopy := &_BACnetConstructedDataStopWhenFull{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.StopWhenFull.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataStopWhenFullCopy
}

func (m *_BACnetConstructedDataStopWhenFull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

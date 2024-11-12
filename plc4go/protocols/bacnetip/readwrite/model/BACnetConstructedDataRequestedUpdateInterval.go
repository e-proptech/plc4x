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

// BACnetConstructedDataRequestedUpdateInterval is the corresponding interface of BACnetConstructedDataRequestedUpdateInterval
type BACnetConstructedDataRequestedUpdateInterval interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRequestedUpdateInterval returns RequestedUpdateInterval (property field)
	GetRequestedUpdateInterval() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataRequestedUpdateInterval is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataRequestedUpdateInterval()
	// CreateBuilder creates a BACnetConstructedDataRequestedUpdateIntervalBuilder
	CreateBACnetConstructedDataRequestedUpdateIntervalBuilder() BACnetConstructedDataRequestedUpdateIntervalBuilder
}

// _BACnetConstructedDataRequestedUpdateInterval is the data-structure of this message
type _BACnetConstructedDataRequestedUpdateInterval struct {
	BACnetConstructedDataContract
	RequestedUpdateInterval BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataRequestedUpdateInterval = (*_BACnetConstructedDataRequestedUpdateInterval)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataRequestedUpdateInterval)(nil)

// NewBACnetConstructedDataRequestedUpdateInterval factory function for _BACnetConstructedDataRequestedUpdateInterval
func NewBACnetConstructedDataRequestedUpdateInterval(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, requestedUpdateInterval BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRequestedUpdateInterval {
	if requestedUpdateInterval == nil {
		panic("requestedUpdateInterval of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataRequestedUpdateInterval must not be nil")
	}
	_result := &_BACnetConstructedDataRequestedUpdateInterval{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RequestedUpdateInterval:       requestedUpdateInterval,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataRequestedUpdateIntervalBuilder is a builder for BACnetConstructedDataRequestedUpdateInterval
type BACnetConstructedDataRequestedUpdateIntervalBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestedUpdateInterval BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRequestedUpdateIntervalBuilder
	// WithRequestedUpdateInterval adds RequestedUpdateInterval (property field)
	WithRequestedUpdateInterval(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRequestedUpdateIntervalBuilder
	// WithRequestedUpdateIntervalBuilder adds RequestedUpdateInterval (property field) which is build by the builder
	WithRequestedUpdateIntervalBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataRequestedUpdateIntervalBuilder
	// Build builds the BACnetConstructedDataRequestedUpdateInterval or returns an error if something is wrong
	Build() (BACnetConstructedDataRequestedUpdateInterval, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataRequestedUpdateInterval
}

// NewBACnetConstructedDataRequestedUpdateIntervalBuilder() creates a BACnetConstructedDataRequestedUpdateIntervalBuilder
func NewBACnetConstructedDataRequestedUpdateIntervalBuilder() BACnetConstructedDataRequestedUpdateIntervalBuilder {
	return &_BACnetConstructedDataRequestedUpdateIntervalBuilder{_BACnetConstructedDataRequestedUpdateInterval: new(_BACnetConstructedDataRequestedUpdateInterval)}
}

type _BACnetConstructedDataRequestedUpdateIntervalBuilder struct {
	*_BACnetConstructedDataRequestedUpdateInterval

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataRequestedUpdateIntervalBuilder) = (*_BACnetConstructedDataRequestedUpdateIntervalBuilder)(nil)

func (b *_BACnetConstructedDataRequestedUpdateIntervalBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataRequestedUpdateIntervalBuilder) WithMandatoryFields(requestedUpdateInterval BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRequestedUpdateIntervalBuilder {
	return b.WithRequestedUpdateInterval(requestedUpdateInterval)
}

func (b *_BACnetConstructedDataRequestedUpdateIntervalBuilder) WithRequestedUpdateInterval(requestedUpdateInterval BACnetApplicationTagUnsignedInteger) BACnetConstructedDataRequestedUpdateIntervalBuilder {
	b.RequestedUpdateInterval = requestedUpdateInterval
	return b
}

func (b *_BACnetConstructedDataRequestedUpdateIntervalBuilder) WithRequestedUpdateIntervalBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataRequestedUpdateIntervalBuilder {
	builder := builderSupplier(b.RequestedUpdateInterval.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.RequestedUpdateInterval, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataRequestedUpdateIntervalBuilder) Build() (BACnetConstructedDataRequestedUpdateInterval, error) {
	if b.RequestedUpdateInterval == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestedUpdateInterval' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataRequestedUpdateInterval.deepCopy(), nil
}

func (b *_BACnetConstructedDataRequestedUpdateIntervalBuilder) MustBuild() BACnetConstructedDataRequestedUpdateInterval {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataRequestedUpdateIntervalBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataRequestedUpdateIntervalBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataRequestedUpdateIntervalBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataRequestedUpdateIntervalBuilder().(*_BACnetConstructedDataRequestedUpdateIntervalBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataRequestedUpdateIntervalBuilder creates a BACnetConstructedDataRequestedUpdateIntervalBuilder
func (b *_BACnetConstructedDataRequestedUpdateInterval) CreateBACnetConstructedDataRequestedUpdateIntervalBuilder() BACnetConstructedDataRequestedUpdateIntervalBuilder {
	if b == nil {
		return NewBACnetConstructedDataRequestedUpdateIntervalBuilder()
	}
	return &_BACnetConstructedDataRequestedUpdateIntervalBuilder{_BACnetConstructedDataRequestedUpdateInterval: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REQUESTED_UPDATE_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetRequestedUpdateInterval() BACnetApplicationTagUnsignedInteger {
	return m.RequestedUpdateInterval
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetRequestedUpdateInterval())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRequestedUpdateInterval(structType any) BACnetConstructedDataRequestedUpdateInterval {
	if casted, ok := structType.(BACnetConstructedDataRequestedUpdateInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRequestedUpdateInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetTypeName() string {
	return "BACnetConstructedDataRequestedUpdateInterval"
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (requestedUpdateInterval)
	lengthInBits += m.RequestedUpdateInterval.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataRequestedUpdateInterval BACnetConstructedDataRequestedUpdateInterval, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRequestedUpdateInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRequestedUpdateInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestedUpdateInterval, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "requestedUpdateInterval", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedUpdateInterval' field"))
	}
	m.RequestedUpdateInterval = requestedUpdateInterval

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), requestedUpdateInterval)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRequestedUpdateInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRequestedUpdateInterval")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRequestedUpdateInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRequestedUpdateInterval")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "requestedUpdateInterval", m.GetRequestedUpdateInterval(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedUpdateInterval' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRequestedUpdateInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRequestedUpdateInterval")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) IsBACnetConstructedDataRequestedUpdateInterval() {
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) deepCopy() *_BACnetConstructedDataRequestedUpdateInterval {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataRequestedUpdateIntervalCopy := &_BACnetConstructedDataRequestedUpdateInterval{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.RequestedUpdateInterval),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataRequestedUpdateIntervalCopy
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) String() string {
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

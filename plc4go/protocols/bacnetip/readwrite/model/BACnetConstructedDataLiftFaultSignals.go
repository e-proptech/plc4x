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

// BACnetConstructedDataLiftFaultSignals is the corresponding interface of BACnetConstructedDataLiftFaultSignals
type BACnetConstructedDataLiftFaultSignals interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFaultSignals returns FaultSignals (property field)
	GetFaultSignals() []BACnetLiftFaultTagged
	// IsBACnetConstructedDataLiftFaultSignals is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLiftFaultSignals()
	// CreateBuilder creates a BACnetConstructedDataLiftFaultSignalsBuilder
	CreateBACnetConstructedDataLiftFaultSignalsBuilder() BACnetConstructedDataLiftFaultSignalsBuilder
}

// _BACnetConstructedDataLiftFaultSignals is the data-structure of this message
type _BACnetConstructedDataLiftFaultSignals struct {
	BACnetConstructedDataContract
	FaultSignals []BACnetLiftFaultTagged
}

var _ BACnetConstructedDataLiftFaultSignals = (*_BACnetConstructedDataLiftFaultSignals)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLiftFaultSignals)(nil)

// NewBACnetConstructedDataLiftFaultSignals factory function for _BACnetConstructedDataLiftFaultSignals
func NewBACnetConstructedDataLiftFaultSignals(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, faultSignals []BACnetLiftFaultTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLiftFaultSignals {
	_result := &_BACnetConstructedDataLiftFaultSignals{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultSignals:                  faultSignals,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLiftFaultSignalsBuilder is a builder for BACnetConstructedDataLiftFaultSignals
type BACnetConstructedDataLiftFaultSignalsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(faultSignals []BACnetLiftFaultTagged) BACnetConstructedDataLiftFaultSignalsBuilder
	// WithFaultSignals adds FaultSignals (property field)
	WithFaultSignals(...BACnetLiftFaultTagged) BACnetConstructedDataLiftFaultSignalsBuilder
	// Build builds the BACnetConstructedDataLiftFaultSignals or returns an error if something is wrong
	Build() (BACnetConstructedDataLiftFaultSignals, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLiftFaultSignals
}

// NewBACnetConstructedDataLiftFaultSignalsBuilder() creates a BACnetConstructedDataLiftFaultSignalsBuilder
func NewBACnetConstructedDataLiftFaultSignalsBuilder() BACnetConstructedDataLiftFaultSignalsBuilder {
	return &_BACnetConstructedDataLiftFaultSignalsBuilder{_BACnetConstructedDataLiftFaultSignals: new(_BACnetConstructedDataLiftFaultSignals)}
}

type _BACnetConstructedDataLiftFaultSignalsBuilder struct {
	*_BACnetConstructedDataLiftFaultSignals

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLiftFaultSignalsBuilder) = (*_BACnetConstructedDataLiftFaultSignalsBuilder)(nil)

func (b *_BACnetConstructedDataLiftFaultSignalsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLiftFaultSignalsBuilder) WithMandatoryFields(faultSignals []BACnetLiftFaultTagged) BACnetConstructedDataLiftFaultSignalsBuilder {
	return b.WithFaultSignals(faultSignals...)
}

func (b *_BACnetConstructedDataLiftFaultSignalsBuilder) WithFaultSignals(faultSignals ...BACnetLiftFaultTagged) BACnetConstructedDataLiftFaultSignalsBuilder {
	b.FaultSignals = faultSignals
	return b
}

func (b *_BACnetConstructedDataLiftFaultSignalsBuilder) Build() (BACnetConstructedDataLiftFaultSignals, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLiftFaultSignals.deepCopy(), nil
}

func (b *_BACnetConstructedDataLiftFaultSignalsBuilder) MustBuild() BACnetConstructedDataLiftFaultSignals {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLiftFaultSignalsBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLiftFaultSignalsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLiftFaultSignalsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLiftFaultSignalsBuilder().(*_BACnetConstructedDataLiftFaultSignalsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLiftFaultSignalsBuilder creates a BACnetConstructedDataLiftFaultSignalsBuilder
func (b *_BACnetConstructedDataLiftFaultSignals) CreateBACnetConstructedDataLiftFaultSignalsBuilder() BACnetConstructedDataLiftFaultSignalsBuilder {
	if b == nil {
		return NewBACnetConstructedDataLiftFaultSignalsBuilder()
	}
	return &_BACnetConstructedDataLiftFaultSignalsBuilder{_BACnetConstructedDataLiftFaultSignals: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLiftFaultSignals) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFT
}

func (m *_BACnetConstructedDataLiftFaultSignals) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_SIGNALS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLiftFaultSignals) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLiftFaultSignals) GetFaultSignals() []BACnetLiftFaultTagged {
	return m.FaultSignals
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLiftFaultSignals(structType any) BACnetConstructedDataLiftFaultSignals {
	if casted, ok := structType.(BACnetConstructedDataLiftFaultSignals); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLiftFaultSignals); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLiftFaultSignals) GetTypeName() string {
	return "BACnetConstructedDataLiftFaultSignals"
}

func (m *_BACnetConstructedDataLiftFaultSignals) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.FaultSignals) > 0 {
		for _, element := range m.FaultSignals {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataLiftFaultSignals) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLiftFaultSignals) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLiftFaultSignals BACnetConstructedDataLiftFaultSignals, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLiftFaultSignals"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLiftFaultSignals")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultSignals, err := ReadTerminatedArrayField[BACnetLiftFaultTagged](ctx, "faultSignals", ReadComplex[BACnetLiftFaultTagged](BACnetLiftFaultTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultSignals' field"))
	}
	m.FaultSignals = faultSignals

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLiftFaultSignals"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLiftFaultSignals")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLiftFaultSignals) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLiftFaultSignals) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLiftFaultSignals"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLiftFaultSignals")
		}

		if err := WriteComplexTypeArrayField(ctx, "faultSignals", m.GetFaultSignals(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'faultSignals' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLiftFaultSignals"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLiftFaultSignals")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLiftFaultSignals) IsBACnetConstructedDataLiftFaultSignals() {}

func (m *_BACnetConstructedDataLiftFaultSignals) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLiftFaultSignals) deepCopy() *_BACnetConstructedDataLiftFaultSignals {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLiftFaultSignalsCopy := &_BACnetConstructedDataLiftFaultSignals{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetLiftFaultTagged, BACnetLiftFaultTagged](m.FaultSignals),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLiftFaultSignalsCopy
}

func (m *_BACnetConstructedDataLiftFaultSignals) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

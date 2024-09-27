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

// BACnetPropertyStatesUnits is the corresponding interface of BACnetPropertyStatesUnits
type BACnetPropertyStatesUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetUnits returns Units (property field)
	GetUnits() BACnetEngineeringUnitsTagged
	// IsBACnetPropertyStatesUnits is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesUnits()
	// CreateBuilder creates a BACnetPropertyStatesUnitsBuilder
	CreateBACnetPropertyStatesUnitsBuilder() BACnetPropertyStatesUnitsBuilder
}

// _BACnetPropertyStatesUnits is the data-structure of this message
type _BACnetPropertyStatesUnits struct {
	BACnetPropertyStatesContract
	Units BACnetEngineeringUnitsTagged
}

var _ BACnetPropertyStatesUnits = (*_BACnetPropertyStatesUnits)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesUnits)(nil)

// NewBACnetPropertyStatesUnits factory function for _BACnetPropertyStatesUnits
func NewBACnetPropertyStatesUnits(peekedTagHeader BACnetTagHeader, units BACnetEngineeringUnitsTagged) *_BACnetPropertyStatesUnits {
	if units == nil {
		panic("units of type BACnetEngineeringUnitsTagged for BACnetPropertyStatesUnits must not be nil")
	}
	_result := &_BACnetPropertyStatesUnits{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		Units:                        units,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesUnitsBuilder is a builder for BACnetPropertyStatesUnits
type BACnetPropertyStatesUnitsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(units BACnetEngineeringUnitsTagged) BACnetPropertyStatesUnitsBuilder
	// WithUnits adds Units (property field)
	WithUnits(BACnetEngineeringUnitsTagged) BACnetPropertyStatesUnitsBuilder
	// WithUnitsBuilder adds Units (property field) which is build by the builder
	WithUnitsBuilder(func(BACnetEngineeringUnitsTaggedBuilder) BACnetEngineeringUnitsTaggedBuilder) BACnetPropertyStatesUnitsBuilder
	// Build builds the BACnetPropertyStatesUnits or returns an error if something is wrong
	Build() (BACnetPropertyStatesUnits, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesUnits
}

// NewBACnetPropertyStatesUnitsBuilder() creates a BACnetPropertyStatesUnitsBuilder
func NewBACnetPropertyStatesUnitsBuilder() BACnetPropertyStatesUnitsBuilder {
	return &_BACnetPropertyStatesUnitsBuilder{_BACnetPropertyStatesUnits: new(_BACnetPropertyStatesUnits)}
}

type _BACnetPropertyStatesUnitsBuilder struct {
	*_BACnetPropertyStatesUnits

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesUnitsBuilder) = (*_BACnetPropertyStatesUnitsBuilder)(nil)

func (b *_BACnetPropertyStatesUnitsBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesUnitsBuilder) WithMandatoryFields(units BACnetEngineeringUnitsTagged) BACnetPropertyStatesUnitsBuilder {
	return b.WithUnits(units)
}

func (b *_BACnetPropertyStatesUnitsBuilder) WithUnits(units BACnetEngineeringUnitsTagged) BACnetPropertyStatesUnitsBuilder {
	b.Units = units
	return b
}

func (b *_BACnetPropertyStatesUnitsBuilder) WithUnitsBuilder(builderSupplier func(BACnetEngineeringUnitsTaggedBuilder) BACnetEngineeringUnitsTaggedBuilder) BACnetPropertyStatesUnitsBuilder {
	builder := builderSupplier(b.Units.CreateBACnetEngineeringUnitsTaggedBuilder())
	var err error
	b.Units, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEngineeringUnitsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesUnitsBuilder) Build() (BACnetPropertyStatesUnits, error) {
	if b.Units == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'units' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesUnits.deepCopy(), nil
}

func (b *_BACnetPropertyStatesUnitsBuilder) MustBuild() BACnetPropertyStatesUnits {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesUnitsBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesUnitsBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesUnitsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesUnitsBuilder().(*_BACnetPropertyStatesUnitsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesUnitsBuilder creates a BACnetPropertyStatesUnitsBuilder
func (b *_BACnetPropertyStatesUnits) CreateBACnetPropertyStatesUnitsBuilder() BACnetPropertyStatesUnitsBuilder {
	if b == nil {
		return NewBACnetPropertyStatesUnitsBuilder()
	}
	return &_BACnetPropertyStatesUnitsBuilder{_BACnetPropertyStatesUnits: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesUnits) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesUnits) GetUnits() BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesUnits(structType any) BACnetPropertyStatesUnits {
	if casted, ok := structType.(BACnetPropertyStatesUnits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesUnits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesUnits) GetTypeName() string {
	return "BACnetPropertyStatesUnits"
}

func (m *_BACnetPropertyStatesUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesUnits) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesUnits BACnetPropertyStatesUnits, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	units, err := ReadSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", ReadComplex[BACnetEngineeringUnitsTagged](BACnetEngineeringUnitsTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'units' field"))
	}
	m.Units = units

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesUnits")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesUnits")
		}

		if err := WriteSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", m.GetUnits(), WriteComplex[BACnetEngineeringUnitsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'units' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesUnits")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesUnits) IsBACnetPropertyStatesUnits() {}

func (m *_BACnetPropertyStatesUnits) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesUnits) deepCopy() *_BACnetPropertyStatesUnits {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesUnitsCopy := &_BACnetPropertyStatesUnits{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.Units.DeepCopy().(BACnetEngineeringUnitsTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesUnitsCopy
}

func (m *_BACnetPropertyStatesUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

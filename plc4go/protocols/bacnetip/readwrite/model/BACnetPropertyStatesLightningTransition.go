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

// BACnetPropertyStatesLightningTransition is the corresponding interface of BACnetPropertyStatesLightningTransition
type BACnetPropertyStatesLightningTransition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetLightningTransition returns LightningTransition (property field)
	GetLightningTransition() BACnetLightingTransitionTagged
	// IsBACnetPropertyStatesLightningTransition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLightningTransition()
	// CreateBuilder creates a BACnetPropertyStatesLightningTransitionBuilder
	CreateBACnetPropertyStatesLightningTransitionBuilder() BACnetPropertyStatesLightningTransitionBuilder
}

// _BACnetPropertyStatesLightningTransition is the data-structure of this message
type _BACnetPropertyStatesLightningTransition struct {
	BACnetPropertyStatesContract
	LightningTransition BACnetLightingTransitionTagged
}

var _ BACnetPropertyStatesLightningTransition = (*_BACnetPropertyStatesLightningTransition)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLightningTransition)(nil)

// NewBACnetPropertyStatesLightningTransition factory function for _BACnetPropertyStatesLightningTransition
func NewBACnetPropertyStatesLightningTransition(peekedTagHeader BACnetTagHeader, lightningTransition BACnetLightingTransitionTagged) *_BACnetPropertyStatesLightningTransition {
	if lightningTransition == nil {
		panic("lightningTransition of type BACnetLightingTransitionTagged for BACnetPropertyStatesLightningTransition must not be nil")
	}
	_result := &_BACnetPropertyStatesLightningTransition{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LightningTransition:          lightningTransition,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesLightningTransitionBuilder is a builder for BACnetPropertyStatesLightningTransition
type BACnetPropertyStatesLightningTransitionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lightningTransition BACnetLightingTransitionTagged) BACnetPropertyStatesLightningTransitionBuilder
	// WithLightningTransition adds LightningTransition (property field)
	WithLightningTransition(BACnetLightingTransitionTagged) BACnetPropertyStatesLightningTransitionBuilder
	// WithLightningTransitionBuilder adds LightningTransition (property field) which is build by the builder
	WithLightningTransitionBuilder(func(BACnetLightingTransitionTaggedBuilder) BACnetLightingTransitionTaggedBuilder) BACnetPropertyStatesLightningTransitionBuilder
	// Build builds the BACnetPropertyStatesLightningTransition or returns an error if something is wrong
	Build() (BACnetPropertyStatesLightningTransition, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesLightningTransition
}

// NewBACnetPropertyStatesLightningTransitionBuilder() creates a BACnetPropertyStatesLightningTransitionBuilder
func NewBACnetPropertyStatesLightningTransitionBuilder() BACnetPropertyStatesLightningTransitionBuilder {
	return &_BACnetPropertyStatesLightningTransitionBuilder{_BACnetPropertyStatesLightningTransition: new(_BACnetPropertyStatesLightningTransition)}
}

type _BACnetPropertyStatesLightningTransitionBuilder struct {
	*_BACnetPropertyStatesLightningTransition

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesLightningTransitionBuilder) = (*_BACnetPropertyStatesLightningTransitionBuilder)(nil)

func (b *_BACnetPropertyStatesLightningTransitionBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesLightningTransitionBuilder) WithMandatoryFields(lightningTransition BACnetLightingTransitionTagged) BACnetPropertyStatesLightningTransitionBuilder {
	return b.WithLightningTransition(lightningTransition)
}

func (b *_BACnetPropertyStatesLightningTransitionBuilder) WithLightningTransition(lightningTransition BACnetLightingTransitionTagged) BACnetPropertyStatesLightningTransitionBuilder {
	b.LightningTransition = lightningTransition
	return b
}

func (b *_BACnetPropertyStatesLightningTransitionBuilder) WithLightningTransitionBuilder(builderSupplier func(BACnetLightingTransitionTaggedBuilder) BACnetLightingTransitionTaggedBuilder) BACnetPropertyStatesLightningTransitionBuilder {
	builder := builderSupplier(b.LightningTransition.CreateBACnetLightingTransitionTaggedBuilder())
	var err error
	b.LightningTransition, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLightingTransitionTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesLightningTransitionBuilder) Build() (BACnetPropertyStatesLightningTransition, error) {
	if b.LightningTransition == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lightningTransition' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesLightningTransition.deepCopy(), nil
}

func (b *_BACnetPropertyStatesLightningTransitionBuilder) MustBuild() BACnetPropertyStatesLightningTransition {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesLightningTransitionBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesLightningTransitionBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesLightningTransitionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesLightningTransitionBuilder().(*_BACnetPropertyStatesLightningTransitionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesLightningTransitionBuilder creates a BACnetPropertyStatesLightningTransitionBuilder
func (b *_BACnetPropertyStatesLightningTransition) CreateBACnetPropertyStatesLightningTransitionBuilder() BACnetPropertyStatesLightningTransitionBuilder {
	if b == nil {
		return NewBACnetPropertyStatesLightningTransitionBuilder()
	}
	return &_BACnetPropertyStatesLightningTransitionBuilder{_BACnetPropertyStatesLightningTransition: b.deepCopy()}
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

func (m *_BACnetPropertyStatesLightningTransition) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLightningTransition) GetLightningTransition() BACnetLightingTransitionTagged {
	return m.LightningTransition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLightningTransition(structType any) BACnetPropertyStatesLightningTransition {
	if casted, ok := structType.(BACnetPropertyStatesLightningTransition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLightningTransition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLightningTransition) GetTypeName() string {
	return "BACnetPropertyStatesLightningTransition"
}

func (m *_BACnetPropertyStatesLightningTransition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lightningTransition)
	lengthInBits += m.LightningTransition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLightningTransition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLightningTransition) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLightningTransition BACnetPropertyStatesLightningTransition, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLightningTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLightningTransition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightningTransition, err := ReadSimpleField[BACnetLightingTransitionTagged](ctx, "lightningTransition", ReadComplex[BACnetLightingTransitionTagged](BACnetLightingTransitionTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightningTransition' field"))
	}
	m.LightningTransition = lightningTransition

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLightningTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLightningTransition")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLightningTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLightningTransition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLightningTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLightningTransition")
		}

		if err := WriteSimpleField[BACnetLightingTransitionTagged](ctx, "lightningTransition", m.GetLightningTransition(), WriteComplex[BACnetLightingTransitionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lightningTransition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLightningTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLightningTransition")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLightningTransition) IsBACnetPropertyStatesLightningTransition() {}

func (m *_BACnetPropertyStatesLightningTransition) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesLightningTransition) deepCopy() *_BACnetPropertyStatesLightningTransition {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesLightningTransitionCopy := &_BACnetPropertyStatesLightningTransition{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetLightingTransitionTagged](m.LightningTransition),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesLightningTransitionCopy
}

func (m *_BACnetPropertyStatesLightningTransition) String() string {
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

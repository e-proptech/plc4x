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

// BACnetPropertyStatesLifeSafetyMode is the corresponding interface of BACnetPropertyStatesLifeSafetyMode
type BACnetPropertyStatesLifeSafetyMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetLifeSafetyMode returns LifeSafetyMode (property field)
	GetLifeSafetyMode() BACnetLifeSafetyModeTagged
	// IsBACnetPropertyStatesLifeSafetyMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLifeSafetyMode()
	// CreateBuilder creates a BACnetPropertyStatesLifeSafetyModeBuilder
	CreateBACnetPropertyStatesLifeSafetyModeBuilder() BACnetPropertyStatesLifeSafetyModeBuilder
}

// _BACnetPropertyStatesLifeSafetyMode is the data-structure of this message
type _BACnetPropertyStatesLifeSafetyMode struct {
	BACnetPropertyStatesContract
	LifeSafetyMode BACnetLifeSafetyModeTagged
}

var _ BACnetPropertyStatesLifeSafetyMode = (*_BACnetPropertyStatesLifeSafetyMode)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLifeSafetyMode)(nil)

// NewBACnetPropertyStatesLifeSafetyMode factory function for _BACnetPropertyStatesLifeSafetyMode
func NewBACnetPropertyStatesLifeSafetyMode(peekedTagHeader BACnetTagHeader, lifeSafetyMode BACnetLifeSafetyModeTagged) *_BACnetPropertyStatesLifeSafetyMode {
	if lifeSafetyMode == nil {
		panic("lifeSafetyMode of type BACnetLifeSafetyModeTagged for BACnetPropertyStatesLifeSafetyMode must not be nil")
	}
	_result := &_BACnetPropertyStatesLifeSafetyMode{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LifeSafetyMode:               lifeSafetyMode,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesLifeSafetyModeBuilder is a builder for BACnetPropertyStatesLifeSafetyMode
type BACnetPropertyStatesLifeSafetyModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lifeSafetyMode BACnetLifeSafetyModeTagged) BACnetPropertyStatesLifeSafetyModeBuilder
	// WithLifeSafetyMode adds LifeSafetyMode (property field)
	WithLifeSafetyMode(BACnetLifeSafetyModeTagged) BACnetPropertyStatesLifeSafetyModeBuilder
	// WithLifeSafetyModeBuilder adds LifeSafetyMode (property field) which is build by the builder
	WithLifeSafetyModeBuilder(func(BACnetLifeSafetyModeTaggedBuilder) BACnetLifeSafetyModeTaggedBuilder) BACnetPropertyStatesLifeSafetyModeBuilder
	// Build builds the BACnetPropertyStatesLifeSafetyMode or returns an error if something is wrong
	Build() (BACnetPropertyStatesLifeSafetyMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesLifeSafetyMode
}

// NewBACnetPropertyStatesLifeSafetyModeBuilder() creates a BACnetPropertyStatesLifeSafetyModeBuilder
func NewBACnetPropertyStatesLifeSafetyModeBuilder() BACnetPropertyStatesLifeSafetyModeBuilder {
	return &_BACnetPropertyStatesLifeSafetyModeBuilder{_BACnetPropertyStatesLifeSafetyMode: new(_BACnetPropertyStatesLifeSafetyMode)}
}

type _BACnetPropertyStatesLifeSafetyModeBuilder struct {
	*_BACnetPropertyStatesLifeSafetyMode

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesLifeSafetyModeBuilder) = (*_BACnetPropertyStatesLifeSafetyModeBuilder)(nil)

func (b *_BACnetPropertyStatesLifeSafetyModeBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesLifeSafetyModeBuilder) WithMandatoryFields(lifeSafetyMode BACnetLifeSafetyModeTagged) BACnetPropertyStatesLifeSafetyModeBuilder {
	return b.WithLifeSafetyMode(lifeSafetyMode)
}

func (b *_BACnetPropertyStatesLifeSafetyModeBuilder) WithLifeSafetyMode(lifeSafetyMode BACnetLifeSafetyModeTagged) BACnetPropertyStatesLifeSafetyModeBuilder {
	b.LifeSafetyMode = lifeSafetyMode
	return b
}

func (b *_BACnetPropertyStatesLifeSafetyModeBuilder) WithLifeSafetyModeBuilder(builderSupplier func(BACnetLifeSafetyModeTaggedBuilder) BACnetLifeSafetyModeTaggedBuilder) BACnetPropertyStatesLifeSafetyModeBuilder {
	builder := builderSupplier(b.LifeSafetyMode.CreateBACnetLifeSafetyModeTaggedBuilder())
	var err error
	b.LifeSafetyMode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLifeSafetyModeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesLifeSafetyModeBuilder) Build() (BACnetPropertyStatesLifeSafetyMode, error) {
	if b.LifeSafetyMode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lifeSafetyMode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesLifeSafetyMode.deepCopy(), nil
}

func (b *_BACnetPropertyStatesLifeSafetyModeBuilder) MustBuild() BACnetPropertyStatesLifeSafetyMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesLifeSafetyModeBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesLifeSafetyModeBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesLifeSafetyModeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesLifeSafetyModeBuilder().(*_BACnetPropertyStatesLifeSafetyModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesLifeSafetyModeBuilder creates a BACnetPropertyStatesLifeSafetyModeBuilder
func (b *_BACnetPropertyStatesLifeSafetyMode) CreateBACnetPropertyStatesLifeSafetyModeBuilder() BACnetPropertyStatesLifeSafetyModeBuilder {
	if b == nil {
		return NewBACnetPropertyStatesLifeSafetyModeBuilder()
	}
	return &_BACnetPropertyStatesLifeSafetyModeBuilder{_BACnetPropertyStatesLifeSafetyMode: b.deepCopy()}
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

func (m *_BACnetPropertyStatesLifeSafetyMode) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLifeSafetyMode) GetLifeSafetyMode() BACnetLifeSafetyModeTagged {
	return m.LifeSafetyMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLifeSafetyMode(structType any) BACnetPropertyStatesLifeSafetyMode {
	if casted, ok := structType.(BACnetPropertyStatesLifeSafetyMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLifeSafetyMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetTypeName() string {
	return "BACnetPropertyStatesLifeSafetyMode"
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lifeSafetyMode)
	lengthInBits += m.LifeSafetyMode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLifeSafetyMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLifeSafetyMode BACnetPropertyStatesLifeSafetyMode, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLifeSafetyMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLifeSafetyMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lifeSafetyMode, err := ReadSimpleField[BACnetLifeSafetyModeTagged](ctx, "lifeSafetyMode", ReadComplex[BACnetLifeSafetyModeTagged](BACnetLifeSafetyModeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifeSafetyMode' field"))
	}
	m.LifeSafetyMode = lifeSafetyMode

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLifeSafetyMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLifeSafetyMode")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLifeSafetyMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLifeSafetyMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLifeSafetyMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLifeSafetyMode")
		}

		if err := WriteSimpleField[BACnetLifeSafetyModeTagged](ctx, "lifeSafetyMode", m.GetLifeSafetyMode(), WriteComplex[BACnetLifeSafetyModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lifeSafetyMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLifeSafetyMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLifeSafetyMode")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLifeSafetyMode) IsBACnetPropertyStatesLifeSafetyMode() {}

func (m *_BACnetPropertyStatesLifeSafetyMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesLifeSafetyMode) deepCopy() *_BACnetPropertyStatesLifeSafetyMode {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesLifeSafetyModeCopy := &_BACnetPropertyStatesLifeSafetyMode{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetLifeSafetyModeTagged](m.LifeSafetyMode),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesLifeSafetyModeCopy
}

func (m *_BACnetPropertyStatesLifeSafetyMode) String() string {
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

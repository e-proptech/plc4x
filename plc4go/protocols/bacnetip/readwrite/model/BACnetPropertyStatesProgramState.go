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

// BACnetPropertyStatesProgramState is the corresponding interface of BACnetPropertyStatesProgramState
type BACnetPropertyStatesProgramState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetProgramState returns ProgramState (property field)
	GetProgramState() BACnetProgramStateTagged
	// IsBACnetPropertyStatesProgramState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesProgramState()
	// CreateBuilder creates a BACnetPropertyStatesProgramStateBuilder
	CreateBACnetPropertyStatesProgramStateBuilder() BACnetPropertyStatesProgramStateBuilder
}

// _BACnetPropertyStatesProgramState is the data-structure of this message
type _BACnetPropertyStatesProgramState struct {
	BACnetPropertyStatesContract
	ProgramState BACnetProgramStateTagged
}

var _ BACnetPropertyStatesProgramState = (*_BACnetPropertyStatesProgramState)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesProgramState)(nil)

// NewBACnetPropertyStatesProgramState factory function for _BACnetPropertyStatesProgramState
func NewBACnetPropertyStatesProgramState(peekedTagHeader BACnetTagHeader, programState BACnetProgramStateTagged) *_BACnetPropertyStatesProgramState {
	if programState == nil {
		panic("programState of type BACnetProgramStateTagged for BACnetPropertyStatesProgramState must not be nil")
	}
	_result := &_BACnetPropertyStatesProgramState{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		ProgramState:                 programState,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesProgramStateBuilder is a builder for BACnetPropertyStatesProgramState
type BACnetPropertyStatesProgramStateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(programState BACnetProgramStateTagged) BACnetPropertyStatesProgramStateBuilder
	// WithProgramState adds ProgramState (property field)
	WithProgramState(BACnetProgramStateTagged) BACnetPropertyStatesProgramStateBuilder
	// WithProgramStateBuilder adds ProgramState (property field) which is build by the builder
	WithProgramStateBuilder(func(BACnetProgramStateTaggedBuilder) BACnetProgramStateTaggedBuilder) BACnetPropertyStatesProgramStateBuilder
	// Build builds the BACnetPropertyStatesProgramState or returns an error if something is wrong
	Build() (BACnetPropertyStatesProgramState, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesProgramState
}

// NewBACnetPropertyStatesProgramStateBuilder() creates a BACnetPropertyStatesProgramStateBuilder
func NewBACnetPropertyStatesProgramStateBuilder() BACnetPropertyStatesProgramStateBuilder {
	return &_BACnetPropertyStatesProgramStateBuilder{_BACnetPropertyStatesProgramState: new(_BACnetPropertyStatesProgramState)}
}

type _BACnetPropertyStatesProgramStateBuilder struct {
	*_BACnetPropertyStatesProgramState

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesProgramStateBuilder) = (*_BACnetPropertyStatesProgramStateBuilder)(nil)

func (b *_BACnetPropertyStatesProgramStateBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesProgramStateBuilder) WithMandatoryFields(programState BACnetProgramStateTagged) BACnetPropertyStatesProgramStateBuilder {
	return b.WithProgramState(programState)
}

func (b *_BACnetPropertyStatesProgramStateBuilder) WithProgramState(programState BACnetProgramStateTagged) BACnetPropertyStatesProgramStateBuilder {
	b.ProgramState = programState
	return b
}

func (b *_BACnetPropertyStatesProgramStateBuilder) WithProgramStateBuilder(builderSupplier func(BACnetProgramStateTaggedBuilder) BACnetProgramStateTaggedBuilder) BACnetPropertyStatesProgramStateBuilder {
	builder := builderSupplier(b.ProgramState.CreateBACnetProgramStateTaggedBuilder())
	var err error
	b.ProgramState, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetProgramStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesProgramStateBuilder) Build() (BACnetPropertyStatesProgramState, error) {
	if b.ProgramState == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'programState' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesProgramState.deepCopy(), nil
}

func (b *_BACnetPropertyStatesProgramStateBuilder) MustBuild() BACnetPropertyStatesProgramState {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesProgramStateBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesProgramStateBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesProgramStateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesProgramStateBuilder().(*_BACnetPropertyStatesProgramStateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesProgramStateBuilder creates a BACnetPropertyStatesProgramStateBuilder
func (b *_BACnetPropertyStatesProgramState) CreateBACnetPropertyStatesProgramStateBuilder() BACnetPropertyStatesProgramStateBuilder {
	if b == nil {
		return NewBACnetPropertyStatesProgramStateBuilder()
	}
	return &_BACnetPropertyStatesProgramStateBuilder{_BACnetPropertyStatesProgramState: b.deepCopy()}
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

func (m *_BACnetPropertyStatesProgramState) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesProgramState) GetProgramState() BACnetProgramStateTagged {
	return m.ProgramState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesProgramState(structType any) BACnetPropertyStatesProgramState {
	if casted, ok := structType.(BACnetPropertyStatesProgramState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesProgramState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesProgramState) GetTypeName() string {
	return "BACnetPropertyStatesProgramState"
}

func (m *_BACnetPropertyStatesProgramState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (programState)
	lengthInBits += m.ProgramState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesProgramState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesProgramState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesProgramState BACnetPropertyStatesProgramState, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesProgramState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesProgramState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	programState, err := ReadSimpleField[BACnetProgramStateTagged](ctx, "programState", ReadComplex[BACnetProgramStateTagged](BACnetProgramStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'programState' field"))
	}
	m.ProgramState = programState

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesProgramState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesProgramState")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesProgramState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesProgramState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesProgramState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesProgramState")
		}

		if err := WriteSimpleField[BACnetProgramStateTagged](ctx, "programState", m.GetProgramState(), WriteComplex[BACnetProgramStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'programState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesProgramState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesProgramState")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesProgramState) IsBACnetPropertyStatesProgramState() {}

func (m *_BACnetPropertyStatesProgramState) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesProgramState) deepCopy() *_BACnetPropertyStatesProgramState {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesProgramStateCopy := &_BACnetPropertyStatesProgramState{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetProgramStateTagged](m.ProgramState),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesProgramStateCopy
}

func (m *_BACnetPropertyStatesProgramState) String() string {
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

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

// S7ParameterModeTransition is the corresponding interface of S7ParameterModeTransition
type S7ParameterModeTransition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7Parameter
	// GetMethod returns Method (property field)
	GetMethod() uint8
	// GetCpuFunctionType returns CpuFunctionType (property field)
	GetCpuFunctionType() uint8
	// GetCpuFunctionGroup returns CpuFunctionGroup (property field)
	GetCpuFunctionGroup() uint8
	// GetCurrentMode returns CurrentMode (property field)
	GetCurrentMode() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
	// IsS7ParameterModeTransition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7ParameterModeTransition()
	// CreateBuilder creates a S7ParameterModeTransitionBuilder
	CreateS7ParameterModeTransitionBuilder() S7ParameterModeTransitionBuilder
}

// _S7ParameterModeTransition is the data-structure of this message
type _S7ParameterModeTransition struct {
	S7ParameterContract
	Method           uint8
	CpuFunctionType  uint8
	CpuFunctionGroup uint8
	CurrentMode      uint8
	SequenceNumber   uint8
	// Reserved Fields
	reservedField0 *uint16
}

var _ S7ParameterModeTransition = (*_S7ParameterModeTransition)(nil)
var _ S7ParameterRequirements = (*_S7ParameterModeTransition)(nil)

// NewS7ParameterModeTransition factory function for _S7ParameterModeTransition
func NewS7ParameterModeTransition(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, currentMode uint8, sequenceNumber uint8) *_S7ParameterModeTransition {
	_result := &_S7ParameterModeTransition{
		S7ParameterContract: NewS7Parameter(),
		Method:              method,
		CpuFunctionType:     cpuFunctionType,
		CpuFunctionGroup:    cpuFunctionGroup,
		CurrentMode:         currentMode,
		SequenceNumber:      sequenceNumber,
	}
	_result.S7ParameterContract.(*_S7Parameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7ParameterModeTransitionBuilder is a builder for S7ParameterModeTransition
type S7ParameterModeTransitionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, currentMode uint8, sequenceNumber uint8) S7ParameterModeTransitionBuilder
	// WithMethod adds Method (property field)
	WithMethod(uint8) S7ParameterModeTransitionBuilder
	// WithCpuFunctionType adds CpuFunctionType (property field)
	WithCpuFunctionType(uint8) S7ParameterModeTransitionBuilder
	// WithCpuFunctionGroup adds CpuFunctionGroup (property field)
	WithCpuFunctionGroup(uint8) S7ParameterModeTransitionBuilder
	// WithCurrentMode adds CurrentMode (property field)
	WithCurrentMode(uint8) S7ParameterModeTransitionBuilder
	// WithSequenceNumber adds SequenceNumber (property field)
	WithSequenceNumber(uint8) S7ParameterModeTransitionBuilder
	// Build builds the S7ParameterModeTransition or returns an error if something is wrong
	Build() (S7ParameterModeTransition, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7ParameterModeTransition
}

// NewS7ParameterModeTransitionBuilder() creates a S7ParameterModeTransitionBuilder
func NewS7ParameterModeTransitionBuilder() S7ParameterModeTransitionBuilder {
	return &_S7ParameterModeTransitionBuilder{_S7ParameterModeTransition: new(_S7ParameterModeTransition)}
}

type _S7ParameterModeTransitionBuilder struct {
	*_S7ParameterModeTransition

	parentBuilder *_S7ParameterBuilder

	err *utils.MultiError
}

var _ (S7ParameterModeTransitionBuilder) = (*_S7ParameterModeTransitionBuilder)(nil)

func (b *_S7ParameterModeTransitionBuilder) setParent(contract S7ParameterContract) {
	b.S7ParameterContract = contract
}

func (b *_S7ParameterModeTransitionBuilder) WithMandatoryFields(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, currentMode uint8, sequenceNumber uint8) S7ParameterModeTransitionBuilder {
	return b.WithMethod(method).WithCpuFunctionType(cpuFunctionType).WithCpuFunctionGroup(cpuFunctionGroup).WithCurrentMode(currentMode).WithSequenceNumber(sequenceNumber)
}

func (b *_S7ParameterModeTransitionBuilder) WithMethod(method uint8) S7ParameterModeTransitionBuilder {
	b.Method = method
	return b
}

func (b *_S7ParameterModeTransitionBuilder) WithCpuFunctionType(cpuFunctionType uint8) S7ParameterModeTransitionBuilder {
	b.CpuFunctionType = cpuFunctionType
	return b
}

func (b *_S7ParameterModeTransitionBuilder) WithCpuFunctionGroup(cpuFunctionGroup uint8) S7ParameterModeTransitionBuilder {
	b.CpuFunctionGroup = cpuFunctionGroup
	return b
}

func (b *_S7ParameterModeTransitionBuilder) WithCurrentMode(currentMode uint8) S7ParameterModeTransitionBuilder {
	b.CurrentMode = currentMode
	return b
}

func (b *_S7ParameterModeTransitionBuilder) WithSequenceNumber(sequenceNumber uint8) S7ParameterModeTransitionBuilder {
	b.SequenceNumber = sequenceNumber
	return b
}

func (b *_S7ParameterModeTransitionBuilder) Build() (S7ParameterModeTransition, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7ParameterModeTransition.deepCopy(), nil
}

func (b *_S7ParameterModeTransitionBuilder) MustBuild() S7ParameterModeTransition {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_S7ParameterModeTransitionBuilder) Done() S7ParameterBuilder {
	return b.parentBuilder
}

func (b *_S7ParameterModeTransitionBuilder) buildForS7Parameter() (S7Parameter, error) {
	return b.Build()
}

func (b *_S7ParameterModeTransitionBuilder) DeepCopy() any {
	_copy := b.CreateS7ParameterModeTransitionBuilder().(*_S7ParameterModeTransitionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7ParameterModeTransitionBuilder creates a S7ParameterModeTransitionBuilder
func (b *_S7ParameterModeTransition) CreateS7ParameterModeTransitionBuilder() S7ParameterModeTransitionBuilder {
	if b == nil {
		return NewS7ParameterModeTransitionBuilder()
	}
	return &_S7ParameterModeTransitionBuilder{_S7ParameterModeTransition: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterModeTransition) GetParameterType() uint8 {
	return 0x01
}

func (m *_S7ParameterModeTransition) GetMessageType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterModeTransition) GetParent() S7ParameterContract {
	return m.S7ParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterModeTransition) GetMethod() uint8 {
	return m.Method
}

func (m *_S7ParameterModeTransition) GetCpuFunctionType() uint8 {
	return m.CpuFunctionType
}

func (m *_S7ParameterModeTransition) GetCpuFunctionGroup() uint8 {
	return m.CpuFunctionGroup
}

func (m *_S7ParameterModeTransition) GetCurrentMode() uint8 {
	return m.CurrentMode
}

func (m *_S7ParameterModeTransition) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7ParameterModeTransition(structType any) S7ParameterModeTransition {
	if casted, ok := structType.(S7ParameterModeTransition); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterModeTransition); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterModeTransition) GetTypeName() string {
	return "S7ParameterModeTransition"
}

func (m *_S7ParameterModeTransition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7ParameterContract.(*_S7Parameter).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (method)
	lengthInBits += 8

	// Simple field (cpuFunctionType)
	lengthInBits += 4

	// Simple field (cpuFunctionGroup)
	lengthInBits += 4

	// Simple field (currentMode)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7ParameterModeTransition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7ParameterModeTransition) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Parameter, messageType uint8) (__s7ParameterModeTransition S7ParameterModeTransition, err error) {
	m.S7ParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterModeTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterModeTransition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0010))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	itemLength, err := ReadImplicitField[uint8](ctx, "itemLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemLength' field"))
	}
	_ = itemLength

	method, err := ReadSimpleField(ctx, "method", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'method' field"))
	}
	m.Method = method

	cpuFunctionType, err := ReadSimpleField(ctx, "cpuFunctionType", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cpuFunctionType' field"))
	}
	m.CpuFunctionType = cpuFunctionType

	cpuFunctionGroup, err := ReadSimpleField(ctx, "cpuFunctionGroup", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cpuFunctionGroup' field"))
	}
	m.CpuFunctionGroup = cpuFunctionGroup

	currentMode, err := ReadSimpleField(ctx, "currentMode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentMode' field"))
	}
	m.CurrentMode = currentMode

	sequenceNumber, err := ReadSimpleField(ctx, "sequenceNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	if closeErr := readBuffer.CloseContext("S7ParameterModeTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterModeTransition")
	}

	return m, nil
}

func (m *_S7ParameterModeTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterModeTransition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterModeTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterModeTransition")
		}

		if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0010), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}
		itemLength := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8(uint8(2)))
		if err := WriteImplicitField(ctx, "itemLength", itemLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemLength' field")
		}

		if err := WriteSimpleField[uint8](ctx, "method", m.GetMethod(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'method' field")
		}

		if err := WriteSimpleField[uint8](ctx, "cpuFunctionType", m.GetCpuFunctionType(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'cpuFunctionType' field")
		}

		if err := WriteSimpleField[uint8](ctx, "cpuFunctionGroup", m.GetCpuFunctionGroup(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'cpuFunctionGroup' field")
		}

		if err := WriteSimpleField[uint8](ctx, "currentMode", m.GetCurrentMode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentMode' field")
		}

		if err := WriteSimpleField[uint8](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterModeTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterModeTransition")
		}
		return nil
	}
	return m.S7ParameterContract.(*_S7Parameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterModeTransition) IsS7ParameterModeTransition() {}

func (m *_S7ParameterModeTransition) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7ParameterModeTransition) deepCopy() *_S7ParameterModeTransition {
	if m == nil {
		return nil
	}
	_S7ParameterModeTransitionCopy := &_S7ParameterModeTransition{
		m.S7ParameterContract.(*_S7Parameter).deepCopy(),
		m.Method,
		m.CpuFunctionType,
		m.CpuFunctionGroup,
		m.CurrentMode,
		m.SequenceNumber,
		m.reservedField0,
	}
	m.S7ParameterContract.(*_S7Parameter)._SubType = m
	return _S7ParameterModeTransitionCopy
}

func (m *_S7ParameterModeTransition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

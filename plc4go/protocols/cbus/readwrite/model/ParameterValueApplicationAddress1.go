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

// ParameterValueApplicationAddress1 is the corresponding interface of ParameterValueApplicationAddress1
type ParameterValueApplicationAddress1 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() ApplicationAddress1
	// GetData returns Data (property field)
	GetData() []byte
	// IsParameterValueApplicationAddress1 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValueApplicationAddress1()
	// CreateBuilder creates a ParameterValueApplicationAddress1Builder
	CreateParameterValueApplicationAddress1Builder() ParameterValueApplicationAddress1Builder
}

// _ParameterValueApplicationAddress1 is the data-structure of this message
type _ParameterValueApplicationAddress1 struct {
	ParameterValueContract
	Value ApplicationAddress1
	Data  []byte
}

var _ ParameterValueApplicationAddress1 = (*_ParameterValueApplicationAddress1)(nil)
var _ ParameterValueRequirements = (*_ParameterValueApplicationAddress1)(nil)

// NewParameterValueApplicationAddress1 factory function for _ParameterValueApplicationAddress1
func NewParameterValueApplicationAddress1(value ApplicationAddress1, data []byte, numBytes uint8) *_ParameterValueApplicationAddress1 {
	if value == nil {
		panic("value of type ApplicationAddress1 for ParameterValueApplicationAddress1 must not be nil")
	}
	_result := &_ParameterValueApplicationAddress1{
		ParameterValueContract: NewParameterValue(numBytes),
		Value:                  value,
		Data:                   data,
	}
	_result.ParameterValueContract.(*_ParameterValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ParameterValueApplicationAddress1Builder is a builder for ParameterValueApplicationAddress1
type ParameterValueApplicationAddress1Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value ApplicationAddress1, data []byte) ParameterValueApplicationAddress1Builder
	// WithValue adds Value (property field)
	WithValue(ApplicationAddress1) ParameterValueApplicationAddress1Builder
	// WithValueBuilder adds Value (property field) which is build by the builder
	WithValueBuilder(func(ApplicationAddress1Builder) ApplicationAddress1Builder) ParameterValueApplicationAddress1Builder
	// WithData adds Data (property field)
	WithData(...byte) ParameterValueApplicationAddress1Builder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ParameterValueBuilder
	// Build builds the ParameterValueApplicationAddress1 or returns an error if something is wrong
	Build() (ParameterValueApplicationAddress1, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ParameterValueApplicationAddress1
}

// NewParameterValueApplicationAddress1Builder() creates a ParameterValueApplicationAddress1Builder
func NewParameterValueApplicationAddress1Builder() ParameterValueApplicationAddress1Builder {
	return &_ParameterValueApplicationAddress1Builder{_ParameterValueApplicationAddress1: new(_ParameterValueApplicationAddress1)}
}

type _ParameterValueApplicationAddress1Builder struct {
	*_ParameterValueApplicationAddress1

	parentBuilder *_ParameterValueBuilder

	err *utils.MultiError
}

var _ (ParameterValueApplicationAddress1Builder) = (*_ParameterValueApplicationAddress1Builder)(nil)

func (b *_ParameterValueApplicationAddress1Builder) setParent(contract ParameterValueContract) {
	b.ParameterValueContract = contract
	contract.(*_ParameterValue)._SubType = b._ParameterValueApplicationAddress1
}

func (b *_ParameterValueApplicationAddress1Builder) WithMandatoryFields(value ApplicationAddress1, data []byte) ParameterValueApplicationAddress1Builder {
	return b.WithValue(value).WithData(data...)
}

func (b *_ParameterValueApplicationAddress1Builder) WithValue(value ApplicationAddress1) ParameterValueApplicationAddress1Builder {
	b.Value = value
	return b
}

func (b *_ParameterValueApplicationAddress1Builder) WithValueBuilder(builderSupplier func(ApplicationAddress1Builder) ApplicationAddress1Builder) ParameterValueApplicationAddress1Builder {
	builder := builderSupplier(b.Value.CreateApplicationAddress1Builder())
	var err error
	b.Value, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ApplicationAddress1Builder failed"))
	}
	return b
}

func (b *_ParameterValueApplicationAddress1Builder) WithData(data ...byte) ParameterValueApplicationAddress1Builder {
	b.Data = data
	return b
}

func (b *_ParameterValueApplicationAddress1Builder) Build() (ParameterValueApplicationAddress1, error) {
	if b.Value == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ParameterValueApplicationAddress1.deepCopy(), nil
}

func (b *_ParameterValueApplicationAddress1Builder) MustBuild() ParameterValueApplicationAddress1 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ParameterValueApplicationAddress1Builder) Done() ParameterValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewParameterValueBuilder().(*_ParameterValueBuilder)
	}
	return b.parentBuilder
}

func (b *_ParameterValueApplicationAddress1Builder) buildForParameterValue() (ParameterValue, error) {
	return b.Build()
}

func (b *_ParameterValueApplicationAddress1Builder) DeepCopy() any {
	_copy := b.CreateParameterValueApplicationAddress1Builder().(*_ParameterValueApplicationAddress1Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateParameterValueApplicationAddress1Builder creates a ParameterValueApplicationAddress1Builder
func (b *_ParameterValueApplicationAddress1) CreateParameterValueApplicationAddress1Builder() ParameterValueApplicationAddress1Builder {
	if b == nil {
		return NewParameterValueApplicationAddress1Builder()
	}
	return &_ParameterValueApplicationAddress1Builder{_ParameterValueApplicationAddress1: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueApplicationAddress1) GetParameterType() ParameterType {
	return ParameterType_APPLICATION_ADDRESS_1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueApplicationAddress1) GetParent() ParameterValueContract {
	return m.ParameterValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueApplicationAddress1) GetValue() ApplicationAddress1 {
	return m.Value
}

func (m *_ParameterValueApplicationAddress1) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastParameterValueApplicationAddress1(structType any) ParameterValueApplicationAddress1 {
	if casted, ok := structType.(ParameterValueApplicationAddress1); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueApplicationAddress1); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueApplicationAddress1) GetTypeName() string {
	return "ParameterValueApplicationAddress1"
}

func (m *_ParameterValueApplicationAddress1) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ParameterValueContract.(*_ParameterValue).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ParameterValueApplicationAddress1) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ParameterValueApplicationAddress1) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ParameterValue, parameterType ParameterType, numBytes uint8) (__parameterValueApplicationAddress1 ParameterValueApplicationAddress1, err error) {
	m.ParameterValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueApplicationAddress1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueApplicationAddress1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) >= (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "ApplicationAddress1 has exactly one byte"})
	}

	value, err := ReadSimpleField[ApplicationAddress1](ctx, "value", ReadComplex[ApplicationAddress1](ApplicationAddress1ParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	data, err := readBuffer.ReadByteArray("data", int(int32(numBytes)-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ParameterValueApplicationAddress1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueApplicationAddress1")
	}

	return m, nil
}

func (m *_ParameterValueApplicationAddress1) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueApplicationAddress1) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueApplicationAddress1"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueApplicationAddress1")
		}

		if err := WriteSimpleField[ApplicationAddress1](ctx, "value", m.GetValue(), WriteComplex[ApplicationAddress1](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueApplicationAddress1"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueApplicationAddress1")
		}
		return nil
	}
	return m.ParameterValueContract.(*_ParameterValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueApplicationAddress1) IsParameterValueApplicationAddress1() {}

func (m *_ParameterValueApplicationAddress1) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ParameterValueApplicationAddress1) deepCopy() *_ParameterValueApplicationAddress1 {
	if m == nil {
		return nil
	}
	_ParameterValueApplicationAddress1Copy := &_ParameterValueApplicationAddress1{
		m.ParameterValueContract.(*_ParameterValue).deepCopy(),
		utils.DeepCopy[ApplicationAddress1](m.Value),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	_ParameterValueApplicationAddress1Copy.ParameterValueContract.(*_ParameterValue)._SubType = m
	return _ParameterValueApplicationAddress1Copy
}

func (m *_ParameterValueApplicationAddress1) String() string {
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

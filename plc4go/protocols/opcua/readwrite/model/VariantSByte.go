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

// VariantSByte is the corresponding interface of VariantSByte
type VariantSByte interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []byte
	// IsVariantSByte is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantSByte()
	// CreateBuilder creates a VariantSByteBuilder
	CreateVariantSByteBuilder() VariantSByteBuilder
}

// _VariantSByte is the data-structure of this message
type _VariantSByte struct {
	VariantContract
	ArrayLength *int32
	Value       []byte
}

var _ VariantSByte = (*_VariantSByte)(nil)
var _ VariantRequirements = (*_VariantSByte)(nil)

// NewVariantSByte factory function for _VariantSByte
func NewVariantSByte(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []byte) *_VariantSByte {
	_result := &_VariantSByte{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VariantSByteBuilder is a builder for VariantSByte
type VariantSByteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []byte) VariantSByteBuilder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantSByteBuilder
	// WithValue adds Value (property field)
	WithValue(...byte) VariantSByteBuilder
	// Build builds the VariantSByte or returns an error if something is wrong
	Build() (VariantSByte, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantSByte
}

// NewVariantSByteBuilder() creates a VariantSByteBuilder
func NewVariantSByteBuilder() VariantSByteBuilder {
	return &_VariantSByteBuilder{_VariantSByte: new(_VariantSByte)}
}

type _VariantSByteBuilder struct {
	*_VariantSByte

	parentBuilder *_VariantBuilder

	err *utils.MultiError
}

var _ (VariantSByteBuilder) = (*_VariantSByteBuilder)(nil)

func (b *_VariantSByteBuilder) setParent(contract VariantContract) {
	b.VariantContract = contract
}

func (b *_VariantSByteBuilder) WithMandatoryFields(value []byte) VariantSByteBuilder {
	return b.WithValue(value...)
}

func (b *_VariantSByteBuilder) WithOptionalArrayLength(arrayLength int32) VariantSByteBuilder {
	b.ArrayLength = &arrayLength
	return b
}

func (b *_VariantSByteBuilder) WithValue(value ...byte) VariantSByteBuilder {
	b.Value = value
	return b
}

func (b *_VariantSByteBuilder) Build() (VariantSByte, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._VariantSByte.deepCopy(), nil
}

func (b *_VariantSByteBuilder) MustBuild() VariantSByte {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_VariantSByteBuilder) Done() VariantBuilder {
	return b.parentBuilder
}

func (b *_VariantSByteBuilder) buildForVariant() (Variant, error) {
	return b.Build()
}

func (b *_VariantSByteBuilder) DeepCopy() any {
	_copy := b.CreateVariantSByteBuilder().(*_VariantSByteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateVariantSByteBuilder creates a VariantSByteBuilder
func (b *_VariantSByte) CreateVariantSByteBuilder() VariantSByteBuilder {
	if b == nil {
		return NewVariantSByteBuilder()
	}
	return &_VariantSByteBuilder{_VariantSByte: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantSByte) GetVariantType() uint8 {
	return uint8(2)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantSByte) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantSByte) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantSByte) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantSByte(structType any) VariantSByte {
	if casted, ok := structType.(VariantSByte); ok {
		return casted
	}
	if casted, ok := structType.(*VariantSByte); ok {
		return *casted
	}
	return nil
}

func (m *_VariantSByte) GetTypeName() string {
	return "VariantSByte"
}

func (m *_VariantSByte) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantSByte) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantSByte) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantSByte VariantSByte, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantSByte"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantSByte")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := readBuffer.ReadByteArray("value", int(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantSByte"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantSByte")
	}

	return m, nil
}

func (m *_VariantSByte) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantSByte) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantSByte"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantSByte")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteByteArrayField(ctx, "value", m.GetValue(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantSByte"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantSByte")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantSByte) IsVariantSByte() {}

func (m *_VariantSByte) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantSByte) deepCopy() *_VariantSByte {
	if m == nil {
		return nil
	}
	_VariantSByteCopy := &_VariantSByte{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[byte, byte](m.Value),
	}
	m.VariantContract.(*_Variant)._SubType = m
	return _VariantSByteCopy
}

func (m *_VariantSByte) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

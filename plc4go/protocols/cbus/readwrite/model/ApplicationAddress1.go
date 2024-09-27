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

// ApplicationAddress1 is the corresponding interface of ApplicationAddress1
type ApplicationAddress1 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetAddress returns Address (property field)
	GetAddress() byte
	// GetIsWildcard returns IsWildcard (virtual field)
	GetIsWildcard() bool
	// IsApplicationAddress1 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApplicationAddress1()
	// CreateBuilder creates a ApplicationAddress1Builder
	CreateApplicationAddress1Builder() ApplicationAddress1Builder
}

// _ApplicationAddress1 is the data-structure of this message
type _ApplicationAddress1 struct {
	Address byte
}

var _ ApplicationAddress1 = (*_ApplicationAddress1)(nil)

// NewApplicationAddress1 factory function for _ApplicationAddress1
func NewApplicationAddress1(address byte) *_ApplicationAddress1 {
	return &_ApplicationAddress1{Address: address}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApplicationAddress1Builder is a builder for ApplicationAddress1
type ApplicationAddress1Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(address byte) ApplicationAddress1Builder
	// WithAddress adds Address (property field)
	WithAddress(byte) ApplicationAddress1Builder
	// Build builds the ApplicationAddress1 or returns an error if something is wrong
	Build() (ApplicationAddress1, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApplicationAddress1
}

// NewApplicationAddress1Builder() creates a ApplicationAddress1Builder
func NewApplicationAddress1Builder() ApplicationAddress1Builder {
	return &_ApplicationAddress1Builder{_ApplicationAddress1: new(_ApplicationAddress1)}
}

type _ApplicationAddress1Builder struct {
	*_ApplicationAddress1

	err *utils.MultiError
}

var _ (ApplicationAddress1Builder) = (*_ApplicationAddress1Builder)(nil)

func (b *_ApplicationAddress1Builder) WithMandatoryFields(address byte) ApplicationAddress1Builder {
	return b.WithAddress(address)
}

func (b *_ApplicationAddress1Builder) WithAddress(address byte) ApplicationAddress1Builder {
	b.Address = address
	return b
}

func (b *_ApplicationAddress1Builder) Build() (ApplicationAddress1, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApplicationAddress1.deepCopy(), nil
}

func (b *_ApplicationAddress1Builder) MustBuild() ApplicationAddress1 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApplicationAddress1Builder) DeepCopy() any {
	_copy := b.CreateApplicationAddress1Builder().(*_ApplicationAddress1Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApplicationAddress1Builder creates a ApplicationAddress1Builder
func (b *_ApplicationAddress1) CreateApplicationAddress1Builder() ApplicationAddress1Builder {
	if b == nil {
		return NewApplicationAddress1Builder()
	}
	return &_ApplicationAddress1Builder{_ApplicationAddress1: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApplicationAddress1) GetAddress() byte {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ApplicationAddress1) GetIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetAddress()) == (0xFF)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApplicationAddress1(structType any) ApplicationAddress1 {
	if casted, ok := structType.(ApplicationAddress1); ok {
		return casted
	}
	if casted, ok := structType.(*ApplicationAddress1); ok {
		return *casted
	}
	return nil
}

func (m *_ApplicationAddress1) GetTypeName() string {
	return "ApplicationAddress1"
}

func (m *_ApplicationAddress1) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (address)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ApplicationAddress1) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApplicationAddress1Parse(ctx context.Context, theBytes []byte) (ApplicationAddress1, error) {
	return ApplicationAddress1ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ApplicationAddress1ParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ApplicationAddress1, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ApplicationAddress1, error) {
		return ApplicationAddress1ParseWithBuffer(ctx, readBuffer)
	}
}

func ApplicationAddress1ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ApplicationAddress1, error) {
	v, err := (&_ApplicationAddress1{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ApplicationAddress1) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__applicationAddress1 ApplicationAddress1, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApplicationAddress1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApplicationAddress1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	isWildcard, err := ReadVirtualField[bool](ctx, "isWildcard", (*bool)(nil), bool((address) == (0xFF)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isWildcard' field"))
	}
	_ = isWildcard

	if closeErr := readBuffer.CloseContext("ApplicationAddress1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApplicationAddress1")
	}

	return m, nil
}

func (m *_ApplicationAddress1) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApplicationAddress1) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ApplicationAddress1"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApplicationAddress1")
	}

	if err := WriteSimpleField[byte](ctx, "address", m.GetAddress(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'address' field")
	}
	// Virtual field
	isWildcard := m.GetIsWildcard()
	_ = isWildcard
	if _isWildcardErr := writeBuffer.WriteVirtual(ctx, "isWildcard", m.GetIsWildcard()); _isWildcardErr != nil {
		return errors.Wrap(_isWildcardErr, "Error serializing 'isWildcard' field")
	}

	if popErr := writeBuffer.PopContext("ApplicationAddress1"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApplicationAddress1")
	}
	return nil
}

func (m *_ApplicationAddress1) IsApplicationAddress1() {}

func (m *_ApplicationAddress1) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApplicationAddress1) deepCopy() *_ApplicationAddress1 {
	if m == nil {
		return nil
	}
	_ApplicationAddress1Copy := &_ApplicationAddress1{
		m.Address,
	}
	return _ApplicationAddress1Copy
}

func (m *_ApplicationAddress1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

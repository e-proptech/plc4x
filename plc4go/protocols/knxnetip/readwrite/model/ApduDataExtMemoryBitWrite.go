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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtMemoryBitWrite is the corresponding interface of ApduDataExtMemoryBitWrite
type ApduDataExtMemoryBitWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtMemoryBitWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtMemoryBitWrite()
	// CreateBuilder creates a ApduDataExtMemoryBitWriteBuilder
	CreateApduDataExtMemoryBitWriteBuilder() ApduDataExtMemoryBitWriteBuilder
}

// _ApduDataExtMemoryBitWrite is the data-structure of this message
type _ApduDataExtMemoryBitWrite struct {
	ApduDataExtContract
}

var _ ApduDataExtMemoryBitWrite = (*_ApduDataExtMemoryBitWrite)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtMemoryBitWrite)(nil)

// NewApduDataExtMemoryBitWrite factory function for _ApduDataExtMemoryBitWrite
func NewApduDataExtMemoryBitWrite(length uint8) *_ApduDataExtMemoryBitWrite {
	_result := &_ApduDataExtMemoryBitWrite{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtMemoryBitWriteBuilder is a builder for ApduDataExtMemoryBitWrite
type ApduDataExtMemoryBitWriteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtMemoryBitWriteBuilder
	// Build builds the ApduDataExtMemoryBitWrite or returns an error if something is wrong
	Build() (ApduDataExtMemoryBitWrite, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtMemoryBitWrite
}

// NewApduDataExtMemoryBitWriteBuilder() creates a ApduDataExtMemoryBitWriteBuilder
func NewApduDataExtMemoryBitWriteBuilder() ApduDataExtMemoryBitWriteBuilder {
	return &_ApduDataExtMemoryBitWriteBuilder{_ApduDataExtMemoryBitWrite: new(_ApduDataExtMemoryBitWrite)}
}

type _ApduDataExtMemoryBitWriteBuilder struct {
	*_ApduDataExtMemoryBitWrite

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtMemoryBitWriteBuilder) = (*_ApduDataExtMemoryBitWriteBuilder)(nil)

func (b *_ApduDataExtMemoryBitWriteBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
}

func (b *_ApduDataExtMemoryBitWriteBuilder) WithMandatoryFields() ApduDataExtMemoryBitWriteBuilder {
	return b
}

func (b *_ApduDataExtMemoryBitWriteBuilder) Build() (ApduDataExtMemoryBitWrite, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtMemoryBitWrite.deepCopy(), nil
}

func (b *_ApduDataExtMemoryBitWriteBuilder) MustBuild() ApduDataExtMemoryBitWrite {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ApduDataExtMemoryBitWriteBuilder) Done() ApduDataExtBuilder {
	return b.parentBuilder
}

func (b *_ApduDataExtMemoryBitWriteBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtMemoryBitWriteBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtMemoryBitWriteBuilder().(*_ApduDataExtMemoryBitWriteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtMemoryBitWriteBuilder creates a ApduDataExtMemoryBitWriteBuilder
func (b *_ApduDataExtMemoryBitWrite) CreateApduDataExtMemoryBitWriteBuilder() ApduDataExtMemoryBitWriteBuilder {
	if b == nil {
		return NewApduDataExtMemoryBitWriteBuilder()
	}
	return &_ApduDataExtMemoryBitWriteBuilder{_ApduDataExtMemoryBitWrite: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtMemoryBitWrite) GetExtApciType() uint8 {
	return 0x10
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtMemoryBitWrite) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtMemoryBitWrite(structType any) ApduDataExtMemoryBitWrite {
	if casted, ok := structType.(ApduDataExtMemoryBitWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtMemoryBitWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtMemoryBitWrite) GetTypeName() string {
	return "ApduDataExtMemoryBitWrite"
}

func (m *_ApduDataExtMemoryBitWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtMemoryBitWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtMemoryBitWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtMemoryBitWrite ApduDataExtMemoryBitWrite, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtMemoryBitWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtMemoryBitWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtMemoryBitWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtMemoryBitWrite")
	}

	return m, nil
}

func (m *_ApduDataExtMemoryBitWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtMemoryBitWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtMemoryBitWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtMemoryBitWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtMemoryBitWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtMemoryBitWrite")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtMemoryBitWrite) IsApduDataExtMemoryBitWrite() {}

func (m *_ApduDataExtMemoryBitWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtMemoryBitWrite) deepCopy() *_ApduDataExtMemoryBitWrite {
	if m == nil {
		return nil
	}
	_ApduDataExtMemoryBitWriteCopy := &_ApduDataExtMemoryBitWrite{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtMemoryBitWriteCopy
}

func (m *_ApduDataExtMemoryBitWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

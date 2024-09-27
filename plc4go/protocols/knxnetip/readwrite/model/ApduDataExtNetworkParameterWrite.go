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

// ApduDataExtNetworkParameterWrite is the corresponding interface of ApduDataExtNetworkParameterWrite
type ApduDataExtNetworkParameterWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtNetworkParameterWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtNetworkParameterWrite()
	// CreateBuilder creates a ApduDataExtNetworkParameterWriteBuilder
	CreateApduDataExtNetworkParameterWriteBuilder() ApduDataExtNetworkParameterWriteBuilder
}

// _ApduDataExtNetworkParameterWrite is the data-structure of this message
type _ApduDataExtNetworkParameterWrite struct {
	ApduDataExtContract
}

var _ ApduDataExtNetworkParameterWrite = (*_ApduDataExtNetworkParameterWrite)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtNetworkParameterWrite)(nil)

// NewApduDataExtNetworkParameterWrite factory function for _ApduDataExtNetworkParameterWrite
func NewApduDataExtNetworkParameterWrite(length uint8) *_ApduDataExtNetworkParameterWrite {
	_result := &_ApduDataExtNetworkParameterWrite{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtNetworkParameterWriteBuilder is a builder for ApduDataExtNetworkParameterWrite
type ApduDataExtNetworkParameterWriteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtNetworkParameterWriteBuilder
	// Build builds the ApduDataExtNetworkParameterWrite or returns an error if something is wrong
	Build() (ApduDataExtNetworkParameterWrite, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtNetworkParameterWrite
}

// NewApduDataExtNetworkParameterWriteBuilder() creates a ApduDataExtNetworkParameterWriteBuilder
func NewApduDataExtNetworkParameterWriteBuilder() ApduDataExtNetworkParameterWriteBuilder {
	return &_ApduDataExtNetworkParameterWriteBuilder{_ApduDataExtNetworkParameterWrite: new(_ApduDataExtNetworkParameterWrite)}
}

type _ApduDataExtNetworkParameterWriteBuilder struct {
	*_ApduDataExtNetworkParameterWrite

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtNetworkParameterWriteBuilder) = (*_ApduDataExtNetworkParameterWriteBuilder)(nil)

func (b *_ApduDataExtNetworkParameterWriteBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
}

func (b *_ApduDataExtNetworkParameterWriteBuilder) WithMandatoryFields() ApduDataExtNetworkParameterWriteBuilder {
	return b
}

func (b *_ApduDataExtNetworkParameterWriteBuilder) Build() (ApduDataExtNetworkParameterWrite, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtNetworkParameterWrite.deepCopy(), nil
}

func (b *_ApduDataExtNetworkParameterWriteBuilder) MustBuild() ApduDataExtNetworkParameterWrite {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ApduDataExtNetworkParameterWriteBuilder) Done() ApduDataExtBuilder {
	return b.parentBuilder
}

func (b *_ApduDataExtNetworkParameterWriteBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtNetworkParameterWriteBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtNetworkParameterWriteBuilder().(*_ApduDataExtNetworkParameterWriteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtNetworkParameterWriteBuilder creates a ApduDataExtNetworkParameterWriteBuilder
func (b *_ApduDataExtNetworkParameterWrite) CreateApduDataExtNetworkParameterWriteBuilder() ApduDataExtNetworkParameterWriteBuilder {
	if b == nil {
		return NewApduDataExtNetworkParameterWriteBuilder()
	}
	return &_ApduDataExtNetworkParameterWriteBuilder{_ApduDataExtNetworkParameterWrite: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtNetworkParameterWrite) GetExtApciType() uint8 {
	return 0x24
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtNetworkParameterWrite) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtNetworkParameterWrite(structType any) ApduDataExtNetworkParameterWrite {
	if casted, ok := structType.(ApduDataExtNetworkParameterWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtNetworkParameterWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtNetworkParameterWrite) GetTypeName() string {
	return "ApduDataExtNetworkParameterWrite"
}

func (m *_ApduDataExtNetworkParameterWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtNetworkParameterWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtNetworkParameterWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtNetworkParameterWrite ApduDataExtNetworkParameterWrite, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtNetworkParameterWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtNetworkParameterWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtNetworkParameterWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtNetworkParameterWrite")
	}

	return m, nil
}

func (m *_ApduDataExtNetworkParameterWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtNetworkParameterWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtNetworkParameterWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtNetworkParameterWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtNetworkParameterWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtNetworkParameterWrite")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtNetworkParameterWrite) IsApduDataExtNetworkParameterWrite() {}

func (m *_ApduDataExtNetworkParameterWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtNetworkParameterWrite) deepCopy() *_ApduDataExtNetworkParameterWrite {
	if m == nil {
		return nil
	}
	_ApduDataExtNetworkParameterWriteCopy := &_ApduDataExtNetworkParameterWrite{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtNetworkParameterWriteCopy
}

func (m *_ApduDataExtNetworkParameterWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

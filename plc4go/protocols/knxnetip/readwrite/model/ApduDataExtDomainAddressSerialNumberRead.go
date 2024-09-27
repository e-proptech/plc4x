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

// ApduDataExtDomainAddressSerialNumberRead is the corresponding interface of ApduDataExtDomainAddressSerialNumberRead
type ApduDataExtDomainAddressSerialNumberRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtDomainAddressSerialNumberRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtDomainAddressSerialNumberRead()
	// CreateBuilder creates a ApduDataExtDomainAddressSerialNumberReadBuilder
	CreateApduDataExtDomainAddressSerialNumberReadBuilder() ApduDataExtDomainAddressSerialNumberReadBuilder
}

// _ApduDataExtDomainAddressSerialNumberRead is the data-structure of this message
type _ApduDataExtDomainAddressSerialNumberRead struct {
	ApduDataExtContract
}

var _ ApduDataExtDomainAddressSerialNumberRead = (*_ApduDataExtDomainAddressSerialNumberRead)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtDomainAddressSerialNumberRead)(nil)

// NewApduDataExtDomainAddressSerialNumberRead factory function for _ApduDataExtDomainAddressSerialNumberRead
func NewApduDataExtDomainAddressSerialNumberRead(length uint8) *_ApduDataExtDomainAddressSerialNumberRead {
	_result := &_ApduDataExtDomainAddressSerialNumberRead{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtDomainAddressSerialNumberReadBuilder is a builder for ApduDataExtDomainAddressSerialNumberRead
type ApduDataExtDomainAddressSerialNumberReadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtDomainAddressSerialNumberReadBuilder
	// Build builds the ApduDataExtDomainAddressSerialNumberRead or returns an error if something is wrong
	Build() (ApduDataExtDomainAddressSerialNumberRead, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtDomainAddressSerialNumberRead
}

// NewApduDataExtDomainAddressSerialNumberReadBuilder() creates a ApduDataExtDomainAddressSerialNumberReadBuilder
func NewApduDataExtDomainAddressSerialNumberReadBuilder() ApduDataExtDomainAddressSerialNumberReadBuilder {
	return &_ApduDataExtDomainAddressSerialNumberReadBuilder{_ApduDataExtDomainAddressSerialNumberRead: new(_ApduDataExtDomainAddressSerialNumberRead)}
}

type _ApduDataExtDomainAddressSerialNumberReadBuilder struct {
	*_ApduDataExtDomainAddressSerialNumberRead

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtDomainAddressSerialNumberReadBuilder) = (*_ApduDataExtDomainAddressSerialNumberReadBuilder)(nil)

func (b *_ApduDataExtDomainAddressSerialNumberReadBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
}

func (b *_ApduDataExtDomainAddressSerialNumberReadBuilder) WithMandatoryFields() ApduDataExtDomainAddressSerialNumberReadBuilder {
	return b
}

func (b *_ApduDataExtDomainAddressSerialNumberReadBuilder) Build() (ApduDataExtDomainAddressSerialNumberRead, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtDomainAddressSerialNumberRead.deepCopy(), nil
}

func (b *_ApduDataExtDomainAddressSerialNumberReadBuilder) MustBuild() ApduDataExtDomainAddressSerialNumberRead {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ApduDataExtDomainAddressSerialNumberReadBuilder) Done() ApduDataExtBuilder {
	return b.parentBuilder
}

func (b *_ApduDataExtDomainAddressSerialNumberReadBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtDomainAddressSerialNumberReadBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtDomainAddressSerialNumberReadBuilder().(*_ApduDataExtDomainAddressSerialNumberReadBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtDomainAddressSerialNumberReadBuilder creates a ApduDataExtDomainAddressSerialNumberReadBuilder
func (b *_ApduDataExtDomainAddressSerialNumberRead) CreateApduDataExtDomainAddressSerialNumberReadBuilder() ApduDataExtDomainAddressSerialNumberReadBuilder {
	if b == nil {
		return NewApduDataExtDomainAddressSerialNumberReadBuilder()
	}
	return &_ApduDataExtDomainAddressSerialNumberReadBuilder{_ApduDataExtDomainAddressSerialNumberRead: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressSerialNumberRead) GetExtApciType() uint8 {
	return 0x2C
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressSerialNumberRead) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressSerialNumberRead(structType any) ApduDataExtDomainAddressSerialNumberRead {
	if casted, ok := structType.(ApduDataExtDomainAddressSerialNumberRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressSerialNumberRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) GetTypeName() string {
	return "ApduDataExtDomainAddressSerialNumberRead"
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtDomainAddressSerialNumberRead ApduDataExtDomainAddressSerialNumberRead, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressSerialNumberRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressSerialNumberRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressSerialNumberRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressSerialNumberRead")
	}

	return m, nil
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressSerialNumberRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressSerialNumberRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressSerialNumberRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressSerialNumberRead")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) IsApduDataExtDomainAddressSerialNumberRead() {}

func (m *_ApduDataExtDomainAddressSerialNumberRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) deepCopy() *_ApduDataExtDomainAddressSerialNumberRead {
	if m == nil {
		return nil
	}
	_ApduDataExtDomainAddressSerialNumberReadCopy := &_ApduDataExtDomainAddressSerialNumberRead{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtDomainAddressSerialNumberReadCopy
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

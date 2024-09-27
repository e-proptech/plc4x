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

// COTPParameterChecksum is the corresponding interface of COTPParameterChecksum
type COTPParameterChecksum interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	COTPParameter
	// GetCrc returns Crc (property field)
	GetCrc() uint8
	// IsCOTPParameterChecksum is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPParameterChecksum()
	// CreateBuilder creates a COTPParameterChecksumBuilder
	CreateCOTPParameterChecksumBuilder() COTPParameterChecksumBuilder
}

// _COTPParameterChecksum is the data-structure of this message
type _COTPParameterChecksum struct {
	COTPParameterContract
	Crc uint8
}

var _ COTPParameterChecksum = (*_COTPParameterChecksum)(nil)
var _ COTPParameterRequirements = (*_COTPParameterChecksum)(nil)

// NewCOTPParameterChecksum factory function for _COTPParameterChecksum
func NewCOTPParameterChecksum(crc uint8, rest uint8) *_COTPParameterChecksum {
	_result := &_COTPParameterChecksum{
		COTPParameterContract: NewCOTPParameter(rest),
		Crc:                   crc,
	}
	_result.COTPParameterContract.(*_COTPParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// COTPParameterChecksumBuilder is a builder for COTPParameterChecksum
type COTPParameterChecksumBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(crc uint8) COTPParameterChecksumBuilder
	// WithCrc adds Crc (property field)
	WithCrc(uint8) COTPParameterChecksumBuilder
	// Build builds the COTPParameterChecksum or returns an error if something is wrong
	Build() (COTPParameterChecksum, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() COTPParameterChecksum
}

// NewCOTPParameterChecksumBuilder() creates a COTPParameterChecksumBuilder
func NewCOTPParameterChecksumBuilder() COTPParameterChecksumBuilder {
	return &_COTPParameterChecksumBuilder{_COTPParameterChecksum: new(_COTPParameterChecksum)}
}

type _COTPParameterChecksumBuilder struct {
	*_COTPParameterChecksum

	parentBuilder *_COTPParameterBuilder

	err *utils.MultiError
}

var _ (COTPParameterChecksumBuilder) = (*_COTPParameterChecksumBuilder)(nil)

func (b *_COTPParameterChecksumBuilder) setParent(contract COTPParameterContract) {
	b.COTPParameterContract = contract
}

func (b *_COTPParameterChecksumBuilder) WithMandatoryFields(crc uint8) COTPParameterChecksumBuilder {
	return b.WithCrc(crc)
}

func (b *_COTPParameterChecksumBuilder) WithCrc(crc uint8) COTPParameterChecksumBuilder {
	b.Crc = crc
	return b
}

func (b *_COTPParameterChecksumBuilder) Build() (COTPParameterChecksum, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._COTPParameterChecksum.deepCopy(), nil
}

func (b *_COTPParameterChecksumBuilder) MustBuild() COTPParameterChecksum {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_COTPParameterChecksumBuilder) Done() COTPParameterBuilder {
	return b.parentBuilder
}

func (b *_COTPParameterChecksumBuilder) buildForCOTPParameter() (COTPParameter, error) {
	return b.Build()
}

func (b *_COTPParameterChecksumBuilder) DeepCopy() any {
	_copy := b.CreateCOTPParameterChecksumBuilder().(*_COTPParameterChecksumBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCOTPParameterChecksumBuilder creates a COTPParameterChecksumBuilder
func (b *_COTPParameterChecksum) CreateCOTPParameterChecksumBuilder() COTPParameterChecksumBuilder {
	if b == nil {
		return NewCOTPParameterChecksumBuilder()
	}
	return &_COTPParameterChecksumBuilder{_COTPParameterChecksum: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPParameterChecksum) GetParameterType() uint8 {
	return 0xC3
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPParameterChecksum) GetParent() COTPParameterContract {
	return m.COTPParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPParameterChecksum) GetCrc() uint8 {
	return m.Crc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCOTPParameterChecksum(structType any) COTPParameterChecksum {
	if casted, ok := structType.(COTPParameterChecksum); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameterChecksum); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameterChecksum) GetTypeName() string {
	return "COTPParameterChecksum"
}

func (m *_COTPParameterChecksum) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.COTPParameterContract.(*_COTPParameter).getLengthInBits(ctx))

	// Simple field (crc)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPParameterChecksum) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_COTPParameterChecksum) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_COTPParameter, rest uint8) (__cOTPParameterChecksum COTPParameterChecksum, err error) {
	m.COTPParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameterChecksum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameterChecksum")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	crc, err := ReadSimpleField(ctx, "crc", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'crc' field"))
	}
	m.Crc = crc

	if closeErr := readBuffer.CloseContext("COTPParameterChecksum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameterChecksum")
	}

	return m, nil
}

func (m *_COTPParameterChecksum) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPParameterChecksum) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterChecksum"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPParameterChecksum")
		}

		if err := WriteSimpleField[uint8](ctx, "crc", m.GetCrc(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterChecksum"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPParameterChecksum")
		}
		return nil
	}
	return m.COTPParameterContract.(*_COTPParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPParameterChecksum) IsCOTPParameterChecksum() {}

func (m *_COTPParameterChecksum) DeepCopy() any {
	return m.deepCopy()
}

func (m *_COTPParameterChecksum) deepCopy() *_COTPParameterChecksum {
	if m == nil {
		return nil
	}
	_COTPParameterChecksumCopy := &_COTPParameterChecksum{
		m.COTPParameterContract.(*_COTPParameter).deepCopy(),
		m.Crc,
	}
	m.COTPParameterContract.(*_COTPParameter)._SubType = m
	return _COTPParameterChecksumCopy
}

func (m *_COTPParameterChecksum) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCRegisterForeignDevice is the corresponding interface of BVLCRegisterForeignDevice
type BVLCRegisterForeignDevice interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BVLC
	// GetTtl returns Ttl (property field)
	GetTtl() uint16
	// IsBVLCRegisterForeignDevice is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLCRegisterForeignDevice()
	// CreateBuilder creates a BVLCRegisterForeignDeviceBuilder
	CreateBVLCRegisterForeignDeviceBuilder() BVLCRegisterForeignDeviceBuilder
}

// _BVLCRegisterForeignDevice is the data-structure of this message
type _BVLCRegisterForeignDevice struct {
	BVLCContract
	Ttl uint16
}

var _ BVLCRegisterForeignDevice = (*_BVLCRegisterForeignDevice)(nil)
var _ BVLCRequirements = (*_BVLCRegisterForeignDevice)(nil)

// NewBVLCRegisterForeignDevice factory function for _BVLCRegisterForeignDevice
func NewBVLCRegisterForeignDevice(ttl uint16) *_BVLCRegisterForeignDevice {
	_result := &_BVLCRegisterForeignDevice{
		BVLCContract: NewBVLC(),
		Ttl:          ttl,
	}
	_result.BVLCContract.(*_BVLC)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BVLCRegisterForeignDeviceBuilder is a builder for BVLCRegisterForeignDevice
type BVLCRegisterForeignDeviceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ttl uint16) BVLCRegisterForeignDeviceBuilder
	// WithTtl adds Ttl (property field)
	WithTtl(uint16) BVLCRegisterForeignDeviceBuilder
	// Build builds the BVLCRegisterForeignDevice or returns an error if something is wrong
	Build() (BVLCRegisterForeignDevice, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BVLCRegisterForeignDevice
}

// NewBVLCRegisterForeignDeviceBuilder() creates a BVLCRegisterForeignDeviceBuilder
func NewBVLCRegisterForeignDeviceBuilder() BVLCRegisterForeignDeviceBuilder {
	return &_BVLCRegisterForeignDeviceBuilder{_BVLCRegisterForeignDevice: new(_BVLCRegisterForeignDevice)}
}

type _BVLCRegisterForeignDeviceBuilder struct {
	*_BVLCRegisterForeignDevice

	parentBuilder *_BVLCBuilder

	err *utils.MultiError
}

var _ (BVLCRegisterForeignDeviceBuilder) = (*_BVLCRegisterForeignDeviceBuilder)(nil)

func (b *_BVLCRegisterForeignDeviceBuilder) setParent(contract BVLCContract) {
	b.BVLCContract = contract
}

func (b *_BVLCRegisterForeignDeviceBuilder) WithMandatoryFields(ttl uint16) BVLCRegisterForeignDeviceBuilder {
	return b.WithTtl(ttl)
}

func (b *_BVLCRegisterForeignDeviceBuilder) WithTtl(ttl uint16) BVLCRegisterForeignDeviceBuilder {
	b.Ttl = ttl
	return b
}

func (b *_BVLCRegisterForeignDeviceBuilder) Build() (BVLCRegisterForeignDevice, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BVLCRegisterForeignDevice.deepCopy(), nil
}

func (b *_BVLCRegisterForeignDeviceBuilder) MustBuild() BVLCRegisterForeignDevice {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BVLCRegisterForeignDeviceBuilder) Done() BVLCBuilder {
	return b.parentBuilder
}

func (b *_BVLCRegisterForeignDeviceBuilder) buildForBVLC() (BVLC, error) {
	return b.Build()
}

func (b *_BVLCRegisterForeignDeviceBuilder) DeepCopy() any {
	_copy := b.CreateBVLCRegisterForeignDeviceBuilder().(*_BVLCRegisterForeignDeviceBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBVLCRegisterForeignDeviceBuilder creates a BVLCRegisterForeignDeviceBuilder
func (b *_BVLCRegisterForeignDevice) CreateBVLCRegisterForeignDeviceBuilder() BVLCRegisterForeignDeviceBuilder {
	if b == nil {
		return NewBVLCRegisterForeignDeviceBuilder()
	}
	return &_BVLCRegisterForeignDeviceBuilder{_BVLCRegisterForeignDevice: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCRegisterForeignDevice) GetBvlcFunction() uint8 {
	return 0x05
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCRegisterForeignDevice) GetParent() BVLCContract {
	return m.BVLCContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCRegisterForeignDevice) GetTtl() uint16 {
	return m.Ttl
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBVLCRegisterForeignDevice(structType any) BVLCRegisterForeignDevice {
	if casted, ok := structType.(BVLCRegisterForeignDevice); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCRegisterForeignDevice); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCRegisterForeignDevice) GetTypeName() string {
	return "BVLCRegisterForeignDevice"
}

func (m *_BVLCRegisterForeignDevice) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	// Simple field (ttl)
	lengthInBits += 16

	return lengthInBits
}

func (m *_BVLCRegisterForeignDevice) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCRegisterForeignDevice) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC) (__bVLCRegisterForeignDevice BVLCRegisterForeignDevice, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCRegisterForeignDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCRegisterForeignDevice")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ttl, err := ReadSimpleField(ctx, "ttl", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ttl' field"))
	}
	m.Ttl = ttl

	if closeErr := readBuffer.CloseContext("BVLCRegisterForeignDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCRegisterForeignDevice")
	}

	return m, nil
}

func (m *_BVLCRegisterForeignDevice) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCRegisterForeignDevice) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCRegisterForeignDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCRegisterForeignDevice")
		}

		if err := WriteSimpleField[uint16](ctx, "ttl", m.GetTtl(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'ttl' field")
		}

		if popErr := writeBuffer.PopContext("BVLCRegisterForeignDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCRegisterForeignDevice")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BVLCRegisterForeignDevice) IsBVLCRegisterForeignDevice() {}

func (m *_BVLCRegisterForeignDevice) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BVLCRegisterForeignDevice) deepCopy() *_BVLCRegisterForeignDevice {
	if m == nil {
		return nil
	}
	_BVLCRegisterForeignDeviceCopy := &_BVLCRegisterForeignDevice{
		m.BVLCContract.(*_BVLC).deepCopy(),
		m.Ttl,
	}
	m.BVLCContract.(*_BVLC)._SubType = m
	return _BVLCRegisterForeignDeviceCopy
}

func (m *_BVLCRegisterForeignDevice) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

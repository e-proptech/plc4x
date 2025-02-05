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

// ModbusAsciiADU is the corresponding interface of ModbusAsciiADU
type ModbusAsciiADU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusADU
	// GetAddress returns Address (property field)
	GetAddress() uint8
	// GetPdu returns Pdu (property field)
	GetPdu() ModbusPDU
	// IsModbusAsciiADU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusAsciiADU()
	// CreateBuilder creates a ModbusAsciiADUBuilder
	CreateModbusAsciiADUBuilder() ModbusAsciiADUBuilder
}

// _ModbusAsciiADU is the data-structure of this message
type _ModbusAsciiADU struct {
	ModbusADUContract
	Address uint8
	Pdu     ModbusPDU
}

var _ ModbusAsciiADU = (*_ModbusAsciiADU)(nil)
var _ ModbusADURequirements = (*_ModbusAsciiADU)(nil)

// NewModbusAsciiADU factory function for _ModbusAsciiADU
func NewModbusAsciiADU(address uint8, pdu ModbusPDU, response bool) *_ModbusAsciiADU {
	if pdu == nil {
		panic("pdu of type ModbusPDU for ModbusAsciiADU must not be nil")
	}
	_result := &_ModbusAsciiADU{
		ModbusADUContract: NewModbusADU(response),
		Address:           address,
		Pdu:               pdu,
	}
	_result.ModbusADUContract.(*_ModbusADU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusAsciiADUBuilder is a builder for ModbusAsciiADU
type ModbusAsciiADUBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(address uint8, pdu ModbusPDU) ModbusAsciiADUBuilder
	// WithAddress adds Address (property field)
	WithAddress(uint8) ModbusAsciiADUBuilder
	// WithPdu adds Pdu (property field)
	WithPdu(ModbusPDU) ModbusAsciiADUBuilder
	// WithPduBuilder adds Pdu (property field) which is build by the builder
	WithPduBuilder(func(ModbusPDUBuilder) ModbusPDUBuilder) ModbusAsciiADUBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusADUBuilder
	// Build builds the ModbusAsciiADU or returns an error if something is wrong
	Build() (ModbusAsciiADU, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusAsciiADU
}

// NewModbusAsciiADUBuilder() creates a ModbusAsciiADUBuilder
func NewModbusAsciiADUBuilder() ModbusAsciiADUBuilder {
	return &_ModbusAsciiADUBuilder{_ModbusAsciiADU: new(_ModbusAsciiADU)}
}

type _ModbusAsciiADUBuilder struct {
	*_ModbusAsciiADU

	parentBuilder *_ModbusADUBuilder

	err *utils.MultiError
}

var _ (ModbusAsciiADUBuilder) = (*_ModbusAsciiADUBuilder)(nil)

func (b *_ModbusAsciiADUBuilder) setParent(contract ModbusADUContract) {
	b.ModbusADUContract = contract
	contract.(*_ModbusADU)._SubType = b._ModbusAsciiADU
}

func (b *_ModbusAsciiADUBuilder) WithMandatoryFields(address uint8, pdu ModbusPDU) ModbusAsciiADUBuilder {
	return b.WithAddress(address).WithPdu(pdu)
}

func (b *_ModbusAsciiADUBuilder) WithAddress(address uint8) ModbusAsciiADUBuilder {
	b.Address = address
	return b
}

func (b *_ModbusAsciiADUBuilder) WithPdu(pdu ModbusPDU) ModbusAsciiADUBuilder {
	b.Pdu = pdu
	return b
}

func (b *_ModbusAsciiADUBuilder) WithPduBuilder(builderSupplier func(ModbusPDUBuilder) ModbusPDUBuilder) ModbusAsciiADUBuilder {
	builder := builderSupplier(b.Pdu.CreateModbusPDUBuilder())
	var err error
	b.Pdu, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ModbusPDUBuilder failed"))
	}
	return b
}

func (b *_ModbusAsciiADUBuilder) Build() (ModbusAsciiADU, error) {
	if b.Pdu == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'pdu' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusAsciiADU.deepCopy(), nil
}

func (b *_ModbusAsciiADUBuilder) MustBuild() ModbusAsciiADU {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusAsciiADUBuilder) Done() ModbusADUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusADUBuilder().(*_ModbusADUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusAsciiADUBuilder) buildForModbusADU() (ModbusADU, error) {
	return b.Build()
}

func (b *_ModbusAsciiADUBuilder) DeepCopy() any {
	_copy := b.CreateModbusAsciiADUBuilder().(*_ModbusAsciiADUBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusAsciiADUBuilder creates a ModbusAsciiADUBuilder
func (b *_ModbusAsciiADU) CreateModbusAsciiADUBuilder() ModbusAsciiADUBuilder {
	if b == nil {
		return NewModbusAsciiADUBuilder()
	}
	return &_ModbusAsciiADUBuilder{_ModbusAsciiADU: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusAsciiADU) GetDriverType() DriverType {
	return DriverType_MODBUS_ASCII
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusAsciiADU) GetParent() ModbusADUContract {
	return m.ModbusADUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusAsciiADU) GetAddress() uint8 {
	return m.Address
}

func (m *_ModbusAsciiADU) GetPdu() ModbusPDU {
	return m.Pdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusAsciiADU(structType any) ModbusAsciiADU {
	if casted, ok := structType.(ModbusAsciiADU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusAsciiADU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusAsciiADU) GetTypeName() string {
	return "ModbusAsciiADU"
}

func (m *_ModbusAsciiADU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusADUContract.(*_ModbusADU).getLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.GetLengthInBits(ctx)

	// Checksum Field (checksum)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModbusAsciiADU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusAsciiADU) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusADU, driverType DriverType, response bool) (__modbusAsciiADU ModbusAsciiADU, err error) {
	m.ModbusADUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusAsciiADU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusAsciiADU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	pdu, err := ReadSimpleField[ModbusPDU](ctx, "pdu", ReadComplex[ModbusPDU](ModbusPDUParseWithBufferProducer[ModbusPDU]((bool)(response)), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pdu' field"))
	}
	m.Pdu = pdu

	crc, err := ReadChecksumField[uint8](ctx, "crc", ReadUnsignedByte(readBuffer, uint8(8)), AsciiLrcCheck(ctx, address, pdu), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'crc' field"))
	}
	_ = crc

	if closeErr := readBuffer.CloseContext("ModbusAsciiADU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusAsciiADU")
	}

	return m, nil
}

func (m *_ModbusAsciiADU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusAsciiADU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusAsciiADU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusAsciiADU")
		}

		if err := WriteSimpleField[uint8](ctx, "address", m.GetAddress(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if err := WriteSimpleField[ModbusPDU](ctx, "pdu", m.GetPdu(), WriteComplex[ModbusPDU](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'pdu' field")
		}

		if err := WriteChecksumField[uint8](ctx, "crc", AsciiLrcCheck(ctx, m.GetAddress(), m.GetPdu()), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("ModbusAsciiADU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusAsciiADU")
		}
		return nil
	}
	return m.ModbusADUContract.(*_ModbusADU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusAsciiADU) IsModbusAsciiADU() {}

func (m *_ModbusAsciiADU) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusAsciiADU) deepCopy() *_ModbusAsciiADU {
	if m == nil {
		return nil
	}
	_ModbusAsciiADUCopy := &_ModbusAsciiADU{
		m.ModbusADUContract.(*_ModbusADU).deepCopy(),
		m.Address,
		utils.DeepCopy[ModbusPDU](m.Pdu),
	}
	_ModbusAsciiADUCopy.ModbusADUContract.(*_ModbusADU)._SubType = m
	return _ModbusAsciiADUCopy
}

func (m *_ModbusAsciiADU) String() string {
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

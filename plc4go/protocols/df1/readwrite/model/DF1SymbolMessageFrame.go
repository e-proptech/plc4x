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

// Constant values.
const DF1SymbolMessageFrame_MESSAGEEND uint8 = 0x10
const DF1SymbolMessageFrame_ENDTRANSACTION uint8 = 0x03

// DF1SymbolMessageFrame is the corresponding interface of DF1SymbolMessageFrame
type DF1SymbolMessageFrame interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	DF1Symbol
	// GetDestinationAddress returns DestinationAddress (property field)
	GetDestinationAddress() uint8
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() uint8
	// GetCommand returns Command (property field)
	GetCommand() DF1Command
	// IsDF1SymbolMessageFrame is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1SymbolMessageFrame()
	// CreateBuilder creates a DF1SymbolMessageFrameBuilder
	CreateDF1SymbolMessageFrameBuilder() DF1SymbolMessageFrameBuilder
}

// _DF1SymbolMessageFrame is the data-structure of this message
type _DF1SymbolMessageFrame struct {
	DF1SymbolContract
	DestinationAddress uint8
	SourceAddress      uint8
	Command            DF1Command
}

var _ DF1SymbolMessageFrame = (*_DF1SymbolMessageFrame)(nil)
var _ DF1SymbolRequirements = (*_DF1SymbolMessageFrame)(nil)

// NewDF1SymbolMessageFrame factory function for _DF1SymbolMessageFrame
func NewDF1SymbolMessageFrame(destinationAddress uint8, sourceAddress uint8, command DF1Command) *_DF1SymbolMessageFrame {
	if command == nil {
		panic("command of type DF1Command for DF1SymbolMessageFrame must not be nil")
	}
	_result := &_DF1SymbolMessageFrame{
		DF1SymbolContract:  NewDF1Symbol(),
		DestinationAddress: destinationAddress,
		SourceAddress:      sourceAddress,
		Command:            command,
	}
	_result.DF1SymbolContract.(*_DF1Symbol)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DF1SymbolMessageFrameBuilder is a builder for DF1SymbolMessageFrame
type DF1SymbolMessageFrameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationAddress uint8, sourceAddress uint8, command DF1Command) DF1SymbolMessageFrameBuilder
	// WithDestinationAddress adds DestinationAddress (property field)
	WithDestinationAddress(uint8) DF1SymbolMessageFrameBuilder
	// WithSourceAddress adds SourceAddress (property field)
	WithSourceAddress(uint8) DF1SymbolMessageFrameBuilder
	// WithCommand adds Command (property field)
	WithCommand(DF1Command) DF1SymbolMessageFrameBuilder
	// WithCommandBuilder adds Command (property field) which is build by the builder
	WithCommandBuilder(func(DF1CommandBuilder) DF1CommandBuilder) DF1SymbolMessageFrameBuilder
	// Build builds the DF1SymbolMessageFrame or returns an error if something is wrong
	Build() (DF1SymbolMessageFrame, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DF1SymbolMessageFrame
}

// NewDF1SymbolMessageFrameBuilder() creates a DF1SymbolMessageFrameBuilder
func NewDF1SymbolMessageFrameBuilder() DF1SymbolMessageFrameBuilder {
	return &_DF1SymbolMessageFrameBuilder{_DF1SymbolMessageFrame: new(_DF1SymbolMessageFrame)}
}

type _DF1SymbolMessageFrameBuilder struct {
	*_DF1SymbolMessageFrame

	parentBuilder *_DF1SymbolBuilder

	err *utils.MultiError
}

var _ (DF1SymbolMessageFrameBuilder) = (*_DF1SymbolMessageFrameBuilder)(nil)

func (b *_DF1SymbolMessageFrameBuilder) setParent(contract DF1SymbolContract) {
	b.DF1SymbolContract = contract
}

func (b *_DF1SymbolMessageFrameBuilder) WithMandatoryFields(destinationAddress uint8, sourceAddress uint8, command DF1Command) DF1SymbolMessageFrameBuilder {
	return b.WithDestinationAddress(destinationAddress).WithSourceAddress(sourceAddress).WithCommand(command)
}

func (b *_DF1SymbolMessageFrameBuilder) WithDestinationAddress(destinationAddress uint8) DF1SymbolMessageFrameBuilder {
	b.DestinationAddress = destinationAddress
	return b
}

func (b *_DF1SymbolMessageFrameBuilder) WithSourceAddress(sourceAddress uint8) DF1SymbolMessageFrameBuilder {
	b.SourceAddress = sourceAddress
	return b
}

func (b *_DF1SymbolMessageFrameBuilder) WithCommand(command DF1Command) DF1SymbolMessageFrameBuilder {
	b.Command = command
	return b
}

func (b *_DF1SymbolMessageFrameBuilder) WithCommandBuilder(builderSupplier func(DF1CommandBuilder) DF1CommandBuilder) DF1SymbolMessageFrameBuilder {
	builder := builderSupplier(b.Command.CreateDF1CommandBuilder())
	var err error
	b.Command, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DF1CommandBuilder failed"))
	}
	return b
}

func (b *_DF1SymbolMessageFrameBuilder) Build() (DF1SymbolMessageFrame, error) {
	if b.Command == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'command' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DF1SymbolMessageFrame.deepCopy(), nil
}

func (b *_DF1SymbolMessageFrameBuilder) MustBuild() DF1SymbolMessageFrame {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_DF1SymbolMessageFrameBuilder) Done() DF1SymbolBuilder {
	return b.parentBuilder
}

func (b *_DF1SymbolMessageFrameBuilder) buildForDF1Symbol() (DF1Symbol, error) {
	return b.Build()
}

func (b *_DF1SymbolMessageFrameBuilder) DeepCopy() any {
	_copy := b.CreateDF1SymbolMessageFrameBuilder().(*_DF1SymbolMessageFrameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDF1SymbolMessageFrameBuilder creates a DF1SymbolMessageFrameBuilder
func (b *_DF1SymbolMessageFrame) CreateDF1SymbolMessageFrameBuilder() DF1SymbolMessageFrameBuilder {
	if b == nil {
		return NewDF1SymbolMessageFrameBuilder()
	}
	return &_DF1SymbolMessageFrameBuilder{_DF1SymbolMessageFrame: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1SymbolMessageFrame) GetSymbolType() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1SymbolMessageFrame) GetParent() DF1SymbolContract {
	return m.DF1SymbolContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1SymbolMessageFrame) GetDestinationAddress() uint8 {
	return m.DestinationAddress
}

func (m *_DF1SymbolMessageFrame) GetSourceAddress() uint8 {
	return m.SourceAddress
}

func (m *_DF1SymbolMessageFrame) GetCommand() DF1Command {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_DF1SymbolMessageFrame) GetMessageEnd() uint8 {
	return DF1SymbolMessageFrame_MESSAGEEND
}

func (m *_DF1SymbolMessageFrame) GetEndTransaction() uint8 {
	return DF1SymbolMessageFrame_ENDTRANSACTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDF1SymbolMessageFrame(structType any) DF1SymbolMessageFrame {
	if casted, ok := structType.(DF1SymbolMessageFrame); ok {
		return casted
	}
	if casted, ok := structType.(*DF1SymbolMessageFrame); ok {
		return *casted
	}
	return nil
}

func (m *_DF1SymbolMessageFrame) GetTypeName() string {
	return "DF1SymbolMessageFrame"
}

func (m *_DF1SymbolMessageFrame) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DF1SymbolContract.(*_DF1Symbol).getLengthInBits(ctx))

	// Simple field (destinationAddress)
	lengthInBits += 8

	// Simple field (sourceAddress)
	lengthInBits += 8

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	// Const Field (messageEnd)
	lengthInBits += 8

	// Const Field (endTransaction)
	lengthInBits += 8

	// Checksum Field (checksum)
	lengthInBits += 16

	return lengthInBits
}

func (m *_DF1SymbolMessageFrame) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DF1SymbolMessageFrame) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DF1Symbol) (__dF1SymbolMessageFrame DF1SymbolMessageFrame, err error) {
	m.DF1SymbolContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1SymbolMessageFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1SymbolMessageFrame")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationAddress, err := ReadSimpleField(ctx, "destinationAddress", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationAddress' field"))
	}
	m.DestinationAddress = destinationAddress

	sourceAddress, err := ReadSimpleField(ctx, "sourceAddress", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceAddress' field"))
	}
	m.SourceAddress = sourceAddress

	command, err := ReadSimpleField[DF1Command](ctx, "command", ReadComplex[DF1Command](DF1CommandParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}
	m.Command = command

	messageEnd, err := ReadConstField[uint8](ctx, "messageEnd", ReadUnsignedByte(readBuffer, uint8(8)), DF1SymbolMessageFrame_MESSAGEEND, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageEnd' field"))
	}
	_ = messageEnd

	endTransaction, err := ReadConstField[uint8](ctx, "endTransaction", ReadUnsignedByte(readBuffer, uint8(8)), DF1SymbolMessageFrame_ENDTRANSACTION, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endTransaction' field"))
	}
	_ = endTransaction

	crc, err := ReadChecksumField[uint16](ctx, "crc", ReadUnsignedShort(readBuffer, uint8(16)), CrcCheck(ctx, destinationAddress, sourceAddress, command), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'crc' field"))
	}
	_ = crc

	if closeErr := readBuffer.CloseContext("DF1SymbolMessageFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1SymbolMessageFrame")
	}

	return m, nil
}

func (m *_DF1SymbolMessageFrame) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1SymbolMessageFrame) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1SymbolMessageFrame"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1SymbolMessageFrame")
		}

		if err := WriteSimpleField[uint8](ctx, "destinationAddress", m.GetDestinationAddress(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationAddress' field")
		}

		if err := WriteSimpleField[uint8](ctx, "sourceAddress", m.GetSourceAddress(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'sourceAddress' field")
		}

		if err := WriteSimpleField[DF1Command](ctx, "command", m.GetCommand(), WriteComplex[DF1Command](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'command' field")
		}

		if err := WriteConstField(ctx, "messageEnd", DF1SymbolMessageFrame_MESSAGEEND, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'messageEnd' field")
		}

		if err := WriteConstField(ctx, "endTransaction", DF1SymbolMessageFrame_ENDTRANSACTION, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'endTransaction' field")
		}

		if err := WriteChecksumField[uint16](ctx, "crc", CrcCheck(ctx, m.GetDestinationAddress(), m.GetSourceAddress(), m.GetCommand()), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("DF1SymbolMessageFrame"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1SymbolMessageFrame")
		}
		return nil
	}
	return m.DF1SymbolContract.(*_DF1Symbol).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1SymbolMessageFrame) IsDF1SymbolMessageFrame() {}

func (m *_DF1SymbolMessageFrame) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DF1SymbolMessageFrame) deepCopy() *_DF1SymbolMessageFrame {
	if m == nil {
		return nil
	}
	_DF1SymbolMessageFrameCopy := &_DF1SymbolMessageFrame{
		m.DF1SymbolContract.(*_DF1Symbol).deepCopy(),
		m.DestinationAddress,
		m.SourceAddress,
		m.Command.DeepCopy().(DF1Command),
	}
	m.DF1SymbolContract.(*_DF1Symbol)._SubType = m
	return _DF1SymbolMessageFrameCopy
}

func (m *_DF1SymbolMessageFrame) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

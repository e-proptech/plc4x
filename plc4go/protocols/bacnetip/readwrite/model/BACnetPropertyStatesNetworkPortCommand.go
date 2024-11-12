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

// BACnetPropertyStatesNetworkPortCommand is the corresponding interface of BACnetPropertyStatesNetworkPortCommand
type BACnetPropertyStatesNetworkPortCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetNetworkPortCommand returns NetworkPortCommand (property field)
	GetNetworkPortCommand() BACnetNetworkPortCommandTagged
	// IsBACnetPropertyStatesNetworkPortCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesNetworkPortCommand()
	// CreateBuilder creates a BACnetPropertyStatesNetworkPortCommandBuilder
	CreateBACnetPropertyStatesNetworkPortCommandBuilder() BACnetPropertyStatesNetworkPortCommandBuilder
}

// _BACnetPropertyStatesNetworkPortCommand is the data-structure of this message
type _BACnetPropertyStatesNetworkPortCommand struct {
	BACnetPropertyStatesContract
	NetworkPortCommand BACnetNetworkPortCommandTagged
}

var _ BACnetPropertyStatesNetworkPortCommand = (*_BACnetPropertyStatesNetworkPortCommand)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesNetworkPortCommand)(nil)

// NewBACnetPropertyStatesNetworkPortCommand factory function for _BACnetPropertyStatesNetworkPortCommand
func NewBACnetPropertyStatesNetworkPortCommand(peekedTagHeader BACnetTagHeader, networkPortCommand BACnetNetworkPortCommandTagged) *_BACnetPropertyStatesNetworkPortCommand {
	if networkPortCommand == nil {
		panic("networkPortCommand of type BACnetNetworkPortCommandTagged for BACnetPropertyStatesNetworkPortCommand must not be nil")
	}
	_result := &_BACnetPropertyStatesNetworkPortCommand{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		NetworkPortCommand:           networkPortCommand,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesNetworkPortCommandBuilder is a builder for BACnetPropertyStatesNetworkPortCommand
type BACnetPropertyStatesNetworkPortCommandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(networkPortCommand BACnetNetworkPortCommandTagged) BACnetPropertyStatesNetworkPortCommandBuilder
	// WithNetworkPortCommand adds NetworkPortCommand (property field)
	WithNetworkPortCommand(BACnetNetworkPortCommandTagged) BACnetPropertyStatesNetworkPortCommandBuilder
	// WithNetworkPortCommandBuilder adds NetworkPortCommand (property field) which is build by the builder
	WithNetworkPortCommandBuilder(func(BACnetNetworkPortCommandTaggedBuilder) BACnetNetworkPortCommandTaggedBuilder) BACnetPropertyStatesNetworkPortCommandBuilder
	// Build builds the BACnetPropertyStatesNetworkPortCommand or returns an error if something is wrong
	Build() (BACnetPropertyStatesNetworkPortCommand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesNetworkPortCommand
}

// NewBACnetPropertyStatesNetworkPortCommandBuilder() creates a BACnetPropertyStatesNetworkPortCommandBuilder
func NewBACnetPropertyStatesNetworkPortCommandBuilder() BACnetPropertyStatesNetworkPortCommandBuilder {
	return &_BACnetPropertyStatesNetworkPortCommandBuilder{_BACnetPropertyStatesNetworkPortCommand: new(_BACnetPropertyStatesNetworkPortCommand)}
}

type _BACnetPropertyStatesNetworkPortCommandBuilder struct {
	*_BACnetPropertyStatesNetworkPortCommand

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesNetworkPortCommandBuilder) = (*_BACnetPropertyStatesNetworkPortCommandBuilder)(nil)

func (b *_BACnetPropertyStatesNetworkPortCommandBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesNetworkPortCommandBuilder) WithMandatoryFields(networkPortCommand BACnetNetworkPortCommandTagged) BACnetPropertyStatesNetworkPortCommandBuilder {
	return b.WithNetworkPortCommand(networkPortCommand)
}

func (b *_BACnetPropertyStatesNetworkPortCommandBuilder) WithNetworkPortCommand(networkPortCommand BACnetNetworkPortCommandTagged) BACnetPropertyStatesNetworkPortCommandBuilder {
	b.NetworkPortCommand = networkPortCommand
	return b
}

func (b *_BACnetPropertyStatesNetworkPortCommandBuilder) WithNetworkPortCommandBuilder(builderSupplier func(BACnetNetworkPortCommandTaggedBuilder) BACnetNetworkPortCommandTaggedBuilder) BACnetPropertyStatesNetworkPortCommandBuilder {
	builder := builderSupplier(b.NetworkPortCommand.CreateBACnetNetworkPortCommandTaggedBuilder())
	var err error
	b.NetworkPortCommand, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetNetworkPortCommandTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesNetworkPortCommandBuilder) Build() (BACnetPropertyStatesNetworkPortCommand, error) {
	if b.NetworkPortCommand == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'networkPortCommand' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesNetworkPortCommand.deepCopy(), nil
}

func (b *_BACnetPropertyStatesNetworkPortCommandBuilder) MustBuild() BACnetPropertyStatesNetworkPortCommand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesNetworkPortCommandBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesNetworkPortCommandBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesNetworkPortCommandBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesNetworkPortCommandBuilder().(*_BACnetPropertyStatesNetworkPortCommandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesNetworkPortCommandBuilder creates a BACnetPropertyStatesNetworkPortCommandBuilder
func (b *_BACnetPropertyStatesNetworkPortCommand) CreateBACnetPropertyStatesNetworkPortCommandBuilder() BACnetPropertyStatesNetworkPortCommandBuilder {
	if b == nil {
		return NewBACnetPropertyStatesNetworkPortCommandBuilder()
	}
	return &_BACnetPropertyStatesNetworkPortCommandBuilder{_BACnetPropertyStatesNetworkPortCommand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesNetworkPortCommand) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesNetworkPortCommand) GetNetworkPortCommand() BACnetNetworkPortCommandTagged {
	return m.NetworkPortCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesNetworkPortCommand(structType any) BACnetPropertyStatesNetworkPortCommand {
	if casted, ok := structType.(BACnetPropertyStatesNetworkPortCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNetworkPortCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesNetworkPortCommand) GetTypeName() string {
	return "BACnetPropertyStatesNetworkPortCommand"
}

func (m *_BACnetPropertyStatesNetworkPortCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (networkPortCommand)
	lengthInBits += m.NetworkPortCommand.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesNetworkPortCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesNetworkPortCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesNetworkPortCommand BACnetPropertyStatesNetworkPortCommand, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNetworkPortCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesNetworkPortCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkPortCommand, err := ReadSimpleField[BACnetNetworkPortCommandTagged](ctx, "networkPortCommand", ReadComplex[BACnetNetworkPortCommandTagged](BACnetNetworkPortCommandTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkPortCommand' field"))
	}
	m.NetworkPortCommand = networkPortCommand

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNetworkPortCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesNetworkPortCommand")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesNetworkPortCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesNetworkPortCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNetworkPortCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesNetworkPortCommand")
		}

		if err := WriteSimpleField[BACnetNetworkPortCommandTagged](ctx, "networkPortCommand", m.GetNetworkPortCommand(), WriteComplex[BACnetNetworkPortCommandTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkPortCommand' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNetworkPortCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesNetworkPortCommand")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesNetworkPortCommand) IsBACnetPropertyStatesNetworkPortCommand() {}

func (m *_BACnetPropertyStatesNetworkPortCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesNetworkPortCommand) deepCopy() *_BACnetPropertyStatesNetworkPortCommand {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesNetworkPortCommandCopy := &_BACnetPropertyStatesNetworkPortCommand{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetNetworkPortCommandTagged](m.NetworkPortCommand),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesNetworkPortCommandCopy
}

func (m *_BACnetPropertyStatesNetworkPortCommand) String() string {
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

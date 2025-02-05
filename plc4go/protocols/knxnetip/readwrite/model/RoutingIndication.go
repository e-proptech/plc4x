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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// RoutingIndication is the corresponding interface of RoutingIndication
type RoutingIndication interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// IsRoutingIndication is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRoutingIndication()
	// CreateBuilder creates a RoutingIndicationBuilder
	CreateRoutingIndicationBuilder() RoutingIndicationBuilder
}

// _RoutingIndication is the data-structure of this message
type _RoutingIndication struct {
	KnxNetIpMessageContract
}

var _ RoutingIndication = (*_RoutingIndication)(nil)
var _ KnxNetIpMessageRequirements = (*_RoutingIndication)(nil)

// NewRoutingIndication factory function for _RoutingIndication
func NewRoutingIndication() *_RoutingIndication {
	_result := &_RoutingIndication{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RoutingIndicationBuilder is a builder for RoutingIndication
type RoutingIndicationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() RoutingIndicationBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() KnxNetIpMessageBuilder
	// Build builds the RoutingIndication or returns an error if something is wrong
	Build() (RoutingIndication, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RoutingIndication
}

// NewRoutingIndicationBuilder() creates a RoutingIndicationBuilder
func NewRoutingIndicationBuilder() RoutingIndicationBuilder {
	return &_RoutingIndicationBuilder{_RoutingIndication: new(_RoutingIndication)}
}

type _RoutingIndicationBuilder struct {
	*_RoutingIndication

	parentBuilder *_KnxNetIpMessageBuilder

	err *utils.MultiError
}

var _ (RoutingIndicationBuilder) = (*_RoutingIndicationBuilder)(nil)

func (b *_RoutingIndicationBuilder) setParent(contract KnxNetIpMessageContract) {
	b.KnxNetIpMessageContract = contract
	contract.(*_KnxNetIpMessage)._SubType = b._RoutingIndication
}

func (b *_RoutingIndicationBuilder) WithMandatoryFields() RoutingIndicationBuilder {
	return b
}

func (b *_RoutingIndicationBuilder) Build() (RoutingIndication, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RoutingIndication.deepCopy(), nil
}

func (b *_RoutingIndicationBuilder) MustBuild() RoutingIndication {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_RoutingIndicationBuilder) Done() KnxNetIpMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewKnxNetIpMessageBuilder().(*_KnxNetIpMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_RoutingIndicationBuilder) buildForKnxNetIpMessage() (KnxNetIpMessage, error) {
	return b.Build()
}

func (b *_RoutingIndicationBuilder) DeepCopy() any {
	_copy := b.CreateRoutingIndicationBuilder().(*_RoutingIndicationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRoutingIndicationBuilder creates a RoutingIndicationBuilder
func (b *_RoutingIndication) CreateRoutingIndicationBuilder() RoutingIndicationBuilder {
	if b == nil {
		return NewRoutingIndicationBuilder()
	}
	return &_RoutingIndicationBuilder{_RoutingIndication: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RoutingIndication) GetMsgType() uint16 {
	return 0x0530
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RoutingIndication) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

// Deprecated: use the interface for direct cast
func CastRoutingIndication(structType any) RoutingIndication {
	if casted, ok := structType.(RoutingIndication); ok {
		return casted
	}
	if casted, ok := structType.(*RoutingIndication); ok {
		return *casted
	}
	return nil
}

func (m *_RoutingIndication) GetTypeName() string {
	return "RoutingIndication"
}

func (m *_RoutingIndication) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_RoutingIndication) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RoutingIndication) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__routingIndication RoutingIndication, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RoutingIndication"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RoutingIndication")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("RoutingIndication"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RoutingIndication")
	}

	return m, nil
}

func (m *_RoutingIndication) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RoutingIndication) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RoutingIndication"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RoutingIndication")
		}

		if popErr := writeBuffer.PopContext("RoutingIndication"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RoutingIndication")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RoutingIndication) IsRoutingIndication() {}

func (m *_RoutingIndication) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RoutingIndication) deepCopy() *_RoutingIndication {
	if m == nil {
		return nil
	}
	_RoutingIndicationCopy := &_RoutingIndication{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
	}
	_RoutingIndicationCopy.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _RoutingIndicationCopy
}

func (m *_RoutingIndication) String() string {
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

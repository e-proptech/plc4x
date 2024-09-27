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

// S7PayloadNotify is the corresponding interface of S7PayloadNotify
type S7PayloadNotify interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// GetAlarmMessage returns AlarmMessage (property field)
	GetAlarmMessage() AlarmMessagePushType
	// IsS7PayloadNotify is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadNotify()
	// CreateBuilder creates a S7PayloadNotifyBuilder
	CreateS7PayloadNotifyBuilder() S7PayloadNotifyBuilder
}

// _S7PayloadNotify is the data-structure of this message
type _S7PayloadNotify struct {
	S7PayloadUserDataItemContract
	AlarmMessage AlarmMessagePushType
}

var _ S7PayloadNotify = (*_S7PayloadNotify)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadNotify)(nil)

// NewS7PayloadNotify factory function for _S7PayloadNotify
func NewS7PayloadNotify(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16, alarmMessage AlarmMessagePushType) *_S7PayloadNotify {
	if alarmMessage == nil {
		panic("alarmMessage of type AlarmMessagePushType for S7PayloadNotify must not be nil")
	}
	_result := &_S7PayloadNotify{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		AlarmMessage:                  alarmMessage,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadNotifyBuilder is a builder for S7PayloadNotify
type S7PayloadNotifyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(alarmMessage AlarmMessagePushType) S7PayloadNotifyBuilder
	// WithAlarmMessage adds AlarmMessage (property field)
	WithAlarmMessage(AlarmMessagePushType) S7PayloadNotifyBuilder
	// WithAlarmMessageBuilder adds AlarmMessage (property field) which is build by the builder
	WithAlarmMessageBuilder(func(AlarmMessagePushTypeBuilder) AlarmMessagePushTypeBuilder) S7PayloadNotifyBuilder
	// Build builds the S7PayloadNotify or returns an error if something is wrong
	Build() (S7PayloadNotify, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadNotify
}

// NewS7PayloadNotifyBuilder() creates a S7PayloadNotifyBuilder
func NewS7PayloadNotifyBuilder() S7PayloadNotifyBuilder {
	return &_S7PayloadNotifyBuilder{_S7PayloadNotify: new(_S7PayloadNotify)}
}

type _S7PayloadNotifyBuilder struct {
	*_S7PayloadNotify

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadNotifyBuilder) = (*_S7PayloadNotifyBuilder)(nil)

func (b *_S7PayloadNotifyBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
}

func (b *_S7PayloadNotifyBuilder) WithMandatoryFields(alarmMessage AlarmMessagePushType) S7PayloadNotifyBuilder {
	return b.WithAlarmMessage(alarmMessage)
}

func (b *_S7PayloadNotifyBuilder) WithAlarmMessage(alarmMessage AlarmMessagePushType) S7PayloadNotifyBuilder {
	b.AlarmMessage = alarmMessage
	return b
}

func (b *_S7PayloadNotifyBuilder) WithAlarmMessageBuilder(builderSupplier func(AlarmMessagePushTypeBuilder) AlarmMessagePushTypeBuilder) S7PayloadNotifyBuilder {
	builder := builderSupplier(b.AlarmMessage.CreateAlarmMessagePushTypeBuilder())
	var err error
	b.AlarmMessage, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "AlarmMessagePushTypeBuilder failed"))
	}
	return b
}

func (b *_S7PayloadNotifyBuilder) Build() (S7PayloadNotify, error) {
	if b.AlarmMessage == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'alarmMessage' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadNotify.deepCopy(), nil
}

func (b *_S7PayloadNotifyBuilder) MustBuild() S7PayloadNotify {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_S7PayloadNotifyBuilder) Done() S7PayloadUserDataItemBuilder {
	return b.parentBuilder
}

func (b *_S7PayloadNotifyBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadNotifyBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadNotifyBuilder().(*_S7PayloadNotifyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadNotifyBuilder creates a S7PayloadNotifyBuilder
func (b *_S7PayloadNotify) CreateS7PayloadNotifyBuilder() S7PayloadNotifyBuilder {
	if b == nil {
		return NewS7PayloadNotifyBuilder()
	}
	return &_S7PayloadNotifyBuilder{_S7PayloadNotify: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadNotify) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadNotify) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *_S7PayloadNotify) GetCpuSubfunction() uint8 {
	return 0x06
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadNotify) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadNotify) GetAlarmMessage() AlarmMessagePushType {
	return m.AlarmMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadNotify(structType any) S7PayloadNotify {
	if casted, ok := structType.(S7PayloadNotify); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadNotify); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadNotify) GetTypeName() string {
	return "S7PayloadNotify"
}

func (m *_S7PayloadNotify) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (alarmMessage)
	lengthInBits += m.AlarmMessage.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7PayloadNotify) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadNotify) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadNotify S7PayloadNotify, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadNotify"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadNotify")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	alarmMessage, err := ReadSimpleField[AlarmMessagePushType](ctx, "alarmMessage", ReadComplex[AlarmMessagePushType](AlarmMessagePushTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmMessage' field"))
	}
	m.AlarmMessage = alarmMessage

	if closeErr := readBuffer.CloseContext("S7PayloadNotify"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadNotify")
	}

	return m, nil
}

func (m *_S7PayloadNotify) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadNotify) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadNotify"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadNotify")
		}

		if err := WriteSimpleField[AlarmMessagePushType](ctx, "alarmMessage", m.GetAlarmMessage(), WriteComplex[AlarmMessagePushType](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmMessage' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadNotify"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadNotify")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadNotify) IsS7PayloadNotify() {}

func (m *_S7PayloadNotify) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadNotify) deepCopy() *_S7PayloadNotify {
	if m == nil {
		return nil
	}
	_S7PayloadNotifyCopy := &_S7PayloadNotify{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
		m.AlarmMessage.DeepCopy().(AlarmMessagePushType),
	}
	m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadNotifyCopy
}

func (m *_S7PayloadNotify) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

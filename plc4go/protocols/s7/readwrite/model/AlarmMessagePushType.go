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

// AlarmMessagePushType is the corresponding interface of AlarmMessagePushType
type AlarmMessagePushType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetTimeStamp returns TimeStamp (property field)
	GetTimeStamp() DateAndTime
	// GetFunctionId returns FunctionId (property field)
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects (property field)
	GetNumberOfObjects() uint8
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []AlarmMessageObjectPushType
	// IsAlarmMessagePushType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAlarmMessagePushType()
	// CreateBuilder creates a AlarmMessagePushTypeBuilder
	CreateAlarmMessagePushTypeBuilder() AlarmMessagePushTypeBuilder
}

// _AlarmMessagePushType is the data-structure of this message
type _AlarmMessagePushType struct {
	TimeStamp       DateAndTime
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []AlarmMessageObjectPushType
}

var _ AlarmMessagePushType = (*_AlarmMessagePushType)(nil)

// NewAlarmMessagePushType factory function for _AlarmMessagePushType
func NewAlarmMessagePushType(timeStamp DateAndTime, functionId uint8, numberOfObjects uint8, messageObjects []AlarmMessageObjectPushType) *_AlarmMessagePushType {
	if timeStamp == nil {
		panic("timeStamp of type DateAndTime for AlarmMessagePushType must not be nil")
	}
	return &_AlarmMessagePushType{TimeStamp: timeStamp, FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AlarmMessagePushTypeBuilder is a builder for AlarmMessagePushType
type AlarmMessagePushTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeStamp DateAndTime, functionId uint8, numberOfObjects uint8, messageObjects []AlarmMessageObjectPushType) AlarmMessagePushTypeBuilder
	// WithTimeStamp adds TimeStamp (property field)
	WithTimeStamp(DateAndTime) AlarmMessagePushTypeBuilder
	// WithTimeStampBuilder adds TimeStamp (property field) which is build by the builder
	WithTimeStampBuilder(func(DateAndTimeBuilder) DateAndTimeBuilder) AlarmMessagePushTypeBuilder
	// WithFunctionId adds FunctionId (property field)
	WithFunctionId(uint8) AlarmMessagePushTypeBuilder
	// WithNumberOfObjects adds NumberOfObjects (property field)
	WithNumberOfObjects(uint8) AlarmMessagePushTypeBuilder
	// WithMessageObjects adds MessageObjects (property field)
	WithMessageObjects(...AlarmMessageObjectPushType) AlarmMessagePushTypeBuilder
	// Build builds the AlarmMessagePushType or returns an error if something is wrong
	Build() (AlarmMessagePushType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AlarmMessagePushType
}

// NewAlarmMessagePushTypeBuilder() creates a AlarmMessagePushTypeBuilder
func NewAlarmMessagePushTypeBuilder() AlarmMessagePushTypeBuilder {
	return &_AlarmMessagePushTypeBuilder{_AlarmMessagePushType: new(_AlarmMessagePushType)}
}

type _AlarmMessagePushTypeBuilder struct {
	*_AlarmMessagePushType

	err *utils.MultiError
}

var _ (AlarmMessagePushTypeBuilder) = (*_AlarmMessagePushTypeBuilder)(nil)

func (b *_AlarmMessagePushTypeBuilder) WithMandatoryFields(timeStamp DateAndTime, functionId uint8, numberOfObjects uint8, messageObjects []AlarmMessageObjectPushType) AlarmMessagePushTypeBuilder {
	return b.WithTimeStamp(timeStamp).WithFunctionId(functionId).WithNumberOfObjects(numberOfObjects).WithMessageObjects(messageObjects...)
}

func (b *_AlarmMessagePushTypeBuilder) WithTimeStamp(timeStamp DateAndTime) AlarmMessagePushTypeBuilder {
	b.TimeStamp = timeStamp
	return b
}

func (b *_AlarmMessagePushTypeBuilder) WithTimeStampBuilder(builderSupplier func(DateAndTimeBuilder) DateAndTimeBuilder) AlarmMessagePushTypeBuilder {
	builder := builderSupplier(b.TimeStamp.CreateDateAndTimeBuilder())
	var err error
	b.TimeStamp, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DateAndTimeBuilder failed"))
	}
	return b
}

func (b *_AlarmMessagePushTypeBuilder) WithFunctionId(functionId uint8) AlarmMessagePushTypeBuilder {
	b.FunctionId = functionId
	return b
}

func (b *_AlarmMessagePushTypeBuilder) WithNumberOfObjects(numberOfObjects uint8) AlarmMessagePushTypeBuilder {
	b.NumberOfObjects = numberOfObjects
	return b
}

func (b *_AlarmMessagePushTypeBuilder) WithMessageObjects(messageObjects ...AlarmMessageObjectPushType) AlarmMessagePushTypeBuilder {
	b.MessageObjects = messageObjects
	return b
}

func (b *_AlarmMessagePushTypeBuilder) Build() (AlarmMessagePushType, error) {
	if b.TimeStamp == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeStamp' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AlarmMessagePushType.deepCopy(), nil
}

func (b *_AlarmMessagePushTypeBuilder) MustBuild() AlarmMessagePushType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AlarmMessagePushTypeBuilder) DeepCopy() any {
	_copy := b.CreateAlarmMessagePushTypeBuilder().(*_AlarmMessagePushTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAlarmMessagePushTypeBuilder creates a AlarmMessagePushTypeBuilder
func (b *_AlarmMessagePushType) CreateAlarmMessagePushTypeBuilder() AlarmMessagePushTypeBuilder {
	if b == nil {
		return NewAlarmMessagePushTypeBuilder()
	}
	return &_AlarmMessagePushTypeBuilder{_AlarmMessagePushType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessagePushType) GetTimeStamp() DateAndTime {
	return m.TimeStamp
}

func (m *_AlarmMessagePushType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *_AlarmMessagePushType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *_AlarmMessagePushType) GetMessageObjects() []AlarmMessageObjectPushType {
	return m.MessageObjects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAlarmMessagePushType(structType any) AlarmMessagePushType {
	if casted, ok := structType.(AlarmMessagePushType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessagePushType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessagePushType) GetTypeName() string {
	return "AlarmMessagePushType"
}

func (m *_AlarmMessagePushType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits(ctx)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for _curItem, element := range m.MessageObjects {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MessageObjects), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AlarmMessagePushType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessagePushTypeParse(ctx context.Context, theBytes []byte) (AlarmMessagePushType, error) {
	return AlarmMessagePushTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessagePushTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessagePushType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessagePushType, error) {
		return AlarmMessagePushTypeParseWithBuffer(ctx, readBuffer)
	}
}

func AlarmMessagePushTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessagePushType, error) {
	v, err := (&_AlarmMessagePushType{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AlarmMessagePushType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__alarmMessagePushType AlarmMessagePushType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessagePushType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessagePushType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeStamp, err := ReadSimpleField[DateAndTime](ctx, "timeStamp", ReadComplex[DateAndTime](DateAndTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeStamp' field"))
	}
	m.TimeStamp = timeStamp

	functionId, err := ReadSimpleField(ctx, "functionId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionId' field"))
	}
	m.FunctionId = functionId

	numberOfObjects, err := ReadSimpleField(ctx, "numberOfObjects", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfObjects' field"))
	}
	m.NumberOfObjects = numberOfObjects

	messageObjects, err := ReadCountArrayField[AlarmMessageObjectPushType](ctx, "messageObjects", ReadComplex[AlarmMessageObjectPushType](AlarmMessageObjectPushTypeParseWithBuffer, readBuffer), uint64(numberOfObjects))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageObjects' field"))
	}
	m.MessageObjects = messageObjects

	if closeErr := readBuffer.CloseContext("AlarmMessagePushType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessagePushType")
	}

	return m, nil
}

func (m *_AlarmMessagePushType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessagePushType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessagePushType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessagePushType")
	}

	if err := WriteSimpleField[DateAndTime](ctx, "timeStamp", m.GetTimeStamp(), WriteComplex[DateAndTime](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeStamp' field")
	}

	if err := WriteSimpleField[uint8](ctx, "functionId", m.GetFunctionId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'functionId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "numberOfObjects", m.GetNumberOfObjects(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'numberOfObjects' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "messageObjects", m.GetMessageObjects(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'messageObjects' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessagePushType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessagePushType")
	}
	return nil
}

func (m *_AlarmMessagePushType) IsAlarmMessagePushType() {}

func (m *_AlarmMessagePushType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AlarmMessagePushType) deepCopy() *_AlarmMessagePushType {
	if m == nil {
		return nil
	}
	_AlarmMessagePushTypeCopy := &_AlarmMessagePushType{
		m.TimeStamp.DeepCopy().(DateAndTime),
		m.FunctionId,
		m.NumberOfObjects,
		utils.DeepCopySlice[AlarmMessageObjectPushType, AlarmMessageObjectPushType](m.MessageObjects),
	}
	return _AlarmMessagePushTypeCopy
}

func (m *_AlarmMessagePushType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

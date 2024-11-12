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

// Constant values.
const AlarmMessageObjectQueryType_VARIABLESPEC uint8 = 0x12

// AlarmMessageObjectQueryType is the corresponding interface of AlarmMessageObjectQueryType
type AlarmMessageObjectQueryType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetLengthDataset returns LengthDataset (property field)
	GetLengthDataset() uint8
	// GetEventState returns EventState (property field)
	GetEventState() State
	// GetAckStateGoing returns AckStateGoing (property field)
	GetAckStateGoing() State
	// GetAckStateComing returns AckStateComing (property field)
	GetAckStateComing() State
	// GetTimeComing returns TimeComing (property field)
	GetTimeComing() DateAndTime
	// GetValueComing returns ValueComing (property field)
	GetValueComing() AssociatedValueType
	// GetTimeGoing returns TimeGoing (property field)
	GetTimeGoing() DateAndTime
	// GetValueGoing returns ValueGoing (property field)
	GetValueGoing() AssociatedValueType
	// IsAlarmMessageObjectQueryType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAlarmMessageObjectQueryType()
	// CreateBuilder creates a AlarmMessageObjectQueryTypeBuilder
	CreateAlarmMessageObjectQueryTypeBuilder() AlarmMessageObjectQueryTypeBuilder
}

// _AlarmMessageObjectQueryType is the data-structure of this message
type _AlarmMessageObjectQueryType struct {
	LengthDataset  uint8
	EventState     State
	AckStateGoing  State
	AckStateComing State
	TimeComing     DateAndTime
	ValueComing    AssociatedValueType
	TimeGoing      DateAndTime
	ValueGoing     AssociatedValueType
	// Reserved Fields
	reservedField0 *uint16
}

var _ AlarmMessageObjectQueryType = (*_AlarmMessageObjectQueryType)(nil)

// NewAlarmMessageObjectQueryType factory function for _AlarmMessageObjectQueryType
func NewAlarmMessageObjectQueryType(lengthDataset uint8, eventState State, ackStateGoing State, ackStateComing State, timeComing DateAndTime, valueComing AssociatedValueType, timeGoing DateAndTime, valueGoing AssociatedValueType) *_AlarmMessageObjectQueryType {
	if eventState == nil {
		panic("eventState of type State for AlarmMessageObjectQueryType must not be nil")
	}
	if ackStateGoing == nil {
		panic("ackStateGoing of type State for AlarmMessageObjectQueryType must not be nil")
	}
	if ackStateComing == nil {
		panic("ackStateComing of type State for AlarmMessageObjectQueryType must not be nil")
	}
	if timeComing == nil {
		panic("timeComing of type DateAndTime for AlarmMessageObjectQueryType must not be nil")
	}
	if valueComing == nil {
		panic("valueComing of type AssociatedValueType for AlarmMessageObjectQueryType must not be nil")
	}
	if timeGoing == nil {
		panic("timeGoing of type DateAndTime for AlarmMessageObjectQueryType must not be nil")
	}
	if valueGoing == nil {
		panic("valueGoing of type AssociatedValueType for AlarmMessageObjectQueryType must not be nil")
	}
	return &_AlarmMessageObjectQueryType{LengthDataset: lengthDataset, EventState: eventState, AckStateGoing: ackStateGoing, AckStateComing: ackStateComing, TimeComing: timeComing, ValueComing: valueComing, TimeGoing: timeGoing, ValueGoing: valueGoing}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AlarmMessageObjectQueryTypeBuilder is a builder for AlarmMessageObjectQueryType
type AlarmMessageObjectQueryTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lengthDataset uint8, eventState State, ackStateGoing State, ackStateComing State, timeComing DateAndTime, valueComing AssociatedValueType, timeGoing DateAndTime, valueGoing AssociatedValueType) AlarmMessageObjectQueryTypeBuilder
	// WithLengthDataset adds LengthDataset (property field)
	WithLengthDataset(uint8) AlarmMessageObjectQueryTypeBuilder
	// WithEventState adds EventState (property field)
	WithEventState(State) AlarmMessageObjectQueryTypeBuilder
	// WithEventStateBuilder adds EventState (property field) which is build by the builder
	WithEventStateBuilder(func(StateBuilder) StateBuilder) AlarmMessageObjectQueryTypeBuilder
	// WithAckStateGoing adds AckStateGoing (property field)
	WithAckStateGoing(State) AlarmMessageObjectQueryTypeBuilder
	// WithAckStateGoingBuilder adds AckStateGoing (property field) which is build by the builder
	WithAckStateGoingBuilder(func(StateBuilder) StateBuilder) AlarmMessageObjectQueryTypeBuilder
	// WithAckStateComing adds AckStateComing (property field)
	WithAckStateComing(State) AlarmMessageObjectQueryTypeBuilder
	// WithAckStateComingBuilder adds AckStateComing (property field) which is build by the builder
	WithAckStateComingBuilder(func(StateBuilder) StateBuilder) AlarmMessageObjectQueryTypeBuilder
	// WithTimeComing adds TimeComing (property field)
	WithTimeComing(DateAndTime) AlarmMessageObjectQueryTypeBuilder
	// WithTimeComingBuilder adds TimeComing (property field) which is build by the builder
	WithTimeComingBuilder(func(DateAndTimeBuilder) DateAndTimeBuilder) AlarmMessageObjectQueryTypeBuilder
	// WithValueComing adds ValueComing (property field)
	WithValueComing(AssociatedValueType) AlarmMessageObjectQueryTypeBuilder
	// WithValueComingBuilder adds ValueComing (property field) which is build by the builder
	WithValueComingBuilder(func(AssociatedValueTypeBuilder) AssociatedValueTypeBuilder) AlarmMessageObjectQueryTypeBuilder
	// WithTimeGoing adds TimeGoing (property field)
	WithTimeGoing(DateAndTime) AlarmMessageObjectQueryTypeBuilder
	// WithTimeGoingBuilder adds TimeGoing (property field) which is build by the builder
	WithTimeGoingBuilder(func(DateAndTimeBuilder) DateAndTimeBuilder) AlarmMessageObjectQueryTypeBuilder
	// WithValueGoing adds ValueGoing (property field)
	WithValueGoing(AssociatedValueType) AlarmMessageObjectQueryTypeBuilder
	// WithValueGoingBuilder adds ValueGoing (property field) which is build by the builder
	WithValueGoingBuilder(func(AssociatedValueTypeBuilder) AssociatedValueTypeBuilder) AlarmMessageObjectQueryTypeBuilder
	// Build builds the AlarmMessageObjectQueryType or returns an error if something is wrong
	Build() (AlarmMessageObjectQueryType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AlarmMessageObjectQueryType
}

// NewAlarmMessageObjectQueryTypeBuilder() creates a AlarmMessageObjectQueryTypeBuilder
func NewAlarmMessageObjectQueryTypeBuilder() AlarmMessageObjectQueryTypeBuilder {
	return &_AlarmMessageObjectQueryTypeBuilder{_AlarmMessageObjectQueryType: new(_AlarmMessageObjectQueryType)}
}

type _AlarmMessageObjectQueryTypeBuilder struct {
	*_AlarmMessageObjectQueryType

	err *utils.MultiError
}

var _ (AlarmMessageObjectQueryTypeBuilder) = (*_AlarmMessageObjectQueryTypeBuilder)(nil)

func (b *_AlarmMessageObjectQueryTypeBuilder) WithMandatoryFields(lengthDataset uint8, eventState State, ackStateGoing State, ackStateComing State, timeComing DateAndTime, valueComing AssociatedValueType, timeGoing DateAndTime, valueGoing AssociatedValueType) AlarmMessageObjectQueryTypeBuilder {
	return b.WithLengthDataset(lengthDataset).WithEventState(eventState).WithAckStateGoing(ackStateGoing).WithAckStateComing(ackStateComing).WithTimeComing(timeComing).WithValueComing(valueComing).WithTimeGoing(timeGoing).WithValueGoing(valueGoing)
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithLengthDataset(lengthDataset uint8) AlarmMessageObjectQueryTypeBuilder {
	b.LengthDataset = lengthDataset
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithEventState(eventState State) AlarmMessageObjectQueryTypeBuilder {
	b.EventState = eventState
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithEventStateBuilder(builderSupplier func(StateBuilder) StateBuilder) AlarmMessageObjectQueryTypeBuilder {
	builder := builderSupplier(b.EventState.CreateStateBuilder())
	var err error
	b.EventState, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StateBuilder failed"))
	}
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithAckStateGoing(ackStateGoing State) AlarmMessageObjectQueryTypeBuilder {
	b.AckStateGoing = ackStateGoing
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithAckStateGoingBuilder(builderSupplier func(StateBuilder) StateBuilder) AlarmMessageObjectQueryTypeBuilder {
	builder := builderSupplier(b.AckStateGoing.CreateStateBuilder())
	var err error
	b.AckStateGoing, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StateBuilder failed"))
	}
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithAckStateComing(ackStateComing State) AlarmMessageObjectQueryTypeBuilder {
	b.AckStateComing = ackStateComing
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithAckStateComingBuilder(builderSupplier func(StateBuilder) StateBuilder) AlarmMessageObjectQueryTypeBuilder {
	builder := builderSupplier(b.AckStateComing.CreateStateBuilder())
	var err error
	b.AckStateComing, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StateBuilder failed"))
	}
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithTimeComing(timeComing DateAndTime) AlarmMessageObjectQueryTypeBuilder {
	b.TimeComing = timeComing
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithTimeComingBuilder(builderSupplier func(DateAndTimeBuilder) DateAndTimeBuilder) AlarmMessageObjectQueryTypeBuilder {
	builder := builderSupplier(b.TimeComing.CreateDateAndTimeBuilder())
	var err error
	b.TimeComing, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DateAndTimeBuilder failed"))
	}
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithValueComing(valueComing AssociatedValueType) AlarmMessageObjectQueryTypeBuilder {
	b.ValueComing = valueComing
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithValueComingBuilder(builderSupplier func(AssociatedValueTypeBuilder) AssociatedValueTypeBuilder) AlarmMessageObjectQueryTypeBuilder {
	builder := builderSupplier(b.ValueComing.CreateAssociatedValueTypeBuilder())
	var err error
	b.ValueComing, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "AssociatedValueTypeBuilder failed"))
	}
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithTimeGoing(timeGoing DateAndTime) AlarmMessageObjectQueryTypeBuilder {
	b.TimeGoing = timeGoing
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithTimeGoingBuilder(builderSupplier func(DateAndTimeBuilder) DateAndTimeBuilder) AlarmMessageObjectQueryTypeBuilder {
	builder := builderSupplier(b.TimeGoing.CreateDateAndTimeBuilder())
	var err error
	b.TimeGoing, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DateAndTimeBuilder failed"))
	}
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithValueGoing(valueGoing AssociatedValueType) AlarmMessageObjectQueryTypeBuilder {
	b.ValueGoing = valueGoing
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) WithValueGoingBuilder(builderSupplier func(AssociatedValueTypeBuilder) AssociatedValueTypeBuilder) AlarmMessageObjectQueryTypeBuilder {
	builder := builderSupplier(b.ValueGoing.CreateAssociatedValueTypeBuilder())
	var err error
	b.ValueGoing, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "AssociatedValueTypeBuilder failed"))
	}
	return b
}

func (b *_AlarmMessageObjectQueryTypeBuilder) Build() (AlarmMessageObjectQueryType, error) {
	if b.EventState == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'eventState' not set"))
	}
	if b.AckStateGoing == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ackStateGoing' not set"))
	}
	if b.AckStateComing == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ackStateComing' not set"))
	}
	if b.TimeComing == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeComing' not set"))
	}
	if b.ValueComing == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'valueComing' not set"))
	}
	if b.TimeGoing == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeGoing' not set"))
	}
	if b.ValueGoing == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'valueGoing' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AlarmMessageObjectQueryType.deepCopy(), nil
}

func (b *_AlarmMessageObjectQueryTypeBuilder) MustBuild() AlarmMessageObjectQueryType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AlarmMessageObjectQueryTypeBuilder) DeepCopy() any {
	_copy := b.CreateAlarmMessageObjectQueryTypeBuilder().(*_AlarmMessageObjectQueryTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAlarmMessageObjectQueryTypeBuilder creates a AlarmMessageObjectQueryTypeBuilder
func (b *_AlarmMessageObjectQueryType) CreateAlarmMessageObjectQueryTypeBuilder() AlarmMessageObjectQueryTypeBuilder {
	if b == nil {
		return NewAlarmMessageObjectQueryTypeBuilder()
	}
	return &_AlarmMessageObjectQueryTypeBuilder{_AlarmMessageObjectQueryType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageObjectQueryType) GetLengthDataset() uint8 {
	return m.LengthDataset
}

func (m *_AlarmMessageObjectQueryType) GetEventState() State {
	return m.EventState
}

func (m *_AlarmMessageObjectQueryType) GetAckStateGoing() State {
	return m.AckStateGoing
}

func (m *_AlarmMessageObjectQueryType) GetAckStateComing() State {
	return m.AckStateComing
}

func (m *_AlarmMessageObjectQueryType) GetTimeComing() DateAndTime {
	return m.TimeComing
}

func (m *_AlarmMessageObjectQueryType) GetValueComing() AssociatedValueType {
	return m.ValueComing
}

func (m *_AlarmMessageObjectQueryType) GetTimeGoing() DateAndTime {
	return m.TimeGoing
}

func (m *_AlarmMessageObjectQueryType) GetValueGoing() AssociatedValueType {
	return m.ValueGoing
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AlarmMessageObjectQueryType) GetVariableSpec() uint8 {
	return AlarmMessageObjectQueryType_VARIABLESPEC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAlarmMessageObjectQueryType(structType any) AlarmMessageObjectQueryType {
	if casted, ok := structType.(AlarmMessageObjectQueryType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageObjectQueryType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageObjectQueryType) GetTypeName() string {
	return "AlarmMessageObjectQueryType"
}

func (m *_AlarmMessageObjectQueryType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (lengthDataset)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Const Field (variableSpec)
	lengthInBits += 8

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits(ctx)

	// Simple field (ackStateGoing)
	lengthInBits += m.AckStateGoing.GetLengthInBits(ctx)

	// Simple field (ackStateComing)
	lengthInBits += m.AckStateComing.GetLengthInBits(ctx)

	// Simple field (timeComing)
	lengthInBits += m.TimeComing.GetLengthInBits(ctx)

	// Simple field (valueComing)
	lengthInBits += m.ValueComing.GetLengthInBits(ctx)

	// Simple field (timeGoing)
	lengthInBits += m.TimeGoing.GetLengthInBits(ctx)

	// Simple field (valueGoing)
	lengthInBits += m.ValueGoing.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AlarmMessageObjectQueryType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageObjectQueryTypeParse(ctx context.Context, theBytes []byte) (AlarmMessageObjectQueryType, error) {
	return AlarmMessageObjectQueryTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageObjectQueryTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectQueryType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectQueryType, error) {
		return AlarmMessageObjectQueryTypeParseWithBuffer(ctx, readBuffer)
	}
}

func AlarmMessageObjectQueryTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectQueryType, error) {
	v, err := (&_AlarmMessageObjectQueryType{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_AlarmMessageObjectQueryType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__alarmMessageObjectQueryType AlarmMessageObjectQueryType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageObjectQueryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageObjectQueryType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lengthDataset, err := ReadSimpleField(ctx, "lengthDataset", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lengthDataset' field"))
	}
	m.LengthDataset = lengthDataset

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	variableSpec, err := ReadConstField[uint8](ctx, "variableSpec", ReadUnsignedByte(readBuffer, uint8(8)), AlarmMessageObjectQueryType_VARIABLESPEC)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'variableSpec' field"))
	}
	_ = variableSpec

	eventState, err := ReadSimpleField[State](ctx, "eventState", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventState' field"))
	}
	m.EventState = eventState

	ackStateGoing, err := ReadSimpleField[State](ctx, "ackStateGoing", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackStateGoing' field"))
	}
	m.AckStateGoing = ackStateGoing

	ackStateComing, err := ReadSimpleField[State](ctx, "ackStateComing", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackStateComing' field"))
	}
	m.AckStateComing = ackStateComing

	timeComing, err := ReadSimpleField[DateAndTime](ctx, "timeComing", ReadComplex[DateAndTime](DateAndTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeComing' field"))
	}
	m.TimeComing = timeComing

	valueComing, err := ReadSimpleField[AssociatedValueType](ctx, "valueComing", ReadComplex[AssociatedValueType](AssociatedValueTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueComing' field"))
	}
	m.ValueComing = valueComing

	timeGoing, err := ReadSimpleField[DateAndTime](ctx, "timeGoing", ReadComplex[DateAndTime](DateAndTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeGoing' field"))
	}
	m.TimeGoing = timeGoing

	valueGoing, err := ReadSimpleField[AssociatedValueType](ctx, "valueGoing", ReadComplex[AssociatedValueType](AssociatedValueTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueGoing' field"))
	}
	m.ValueGoing = valueGoing

	if closeErr := readBuffer.CloseContext("AlarmMessageObjectQueryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageObjectQueryType")
	}

	return m, nil
}

func (m *_AlarmMessageObjectQueryType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageObjectQueryType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessageObjectQueryType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageObjectQueryType")
	}

	if err := WriteSimpleField[uint8](ctx, "lengthDataset", m.GetLengthDataset(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'lengthDataset' field")
	}

	if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0000), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteConstField(ctx, "variableSpec", AlarmMessageObjectQueryType_VARIABLESPEC, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'variableSpec' field")
	}

	if err := WriteSimpleField[State](ctx, "eventState", m.GetEventState(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventState' field")
	}

	if err := WriteSimpleField[State](ctx, "ackStateGoing", m.GetAckStateGoing(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'ackStateGoing' field")
	}

	if err := WriteSimpleField[State](ctx, "ackStateComing", m.GetAckStateComing(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'ackStateComing' field")
	}

	if err := WriteSimpleField[DateAndTime](ctx, "timeComing", m.GetTimeComing(), WriteComplex[DateAndTime](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeComing' field")
	}

	if err := WriteSimpleField[AssociatedValueType](ctx, "valueComing", m.GetValueComing(), WriteComplex[AssociatedValueType](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'valueComing' field")
	}

	if err := WriteSimpleField[DateAndTime](ctx, "timeGoing", m.GetTimeGoing(), WriteComplex[DateAndTime](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeGoing' field")
	}

	if err := WriteSimpleField[AssociatedValueType](ctx, "valueGoing", m.GetValueGoing(), WriteComplex[AssociatedValueType](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'valueGoing' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageObjectQueryType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageObjectQueryType")
	}
	return nil
}

func (m *_AlarmMessageObjectQueryType) IsAlarmMessageObjectQueryType() {}

func (m *_AlarmMessageObjectQueryType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AlarmMessageObjectQueryType) deepCopy() *_AlarmMessageObjectQueryType {
	if m == nil {
		return nil
	}
	_AlarmMessageObjectQueryTypeCopy := &_AlarmMessageObjectQueryType{
		m.LengthDataset,
		utils.DeepCopy[State](m.EventState),
		utils.DeepCopy[State](m.AckStateGoing),
		utils.DeepCopy[State](m.AckStateComing),
		utils.DeepCopy[DateAndTime](m.TimeComing),
		utils.DeepCopy[AssociatedValueType](m.ValueComing),
		utils.DeepCopy[DateAndTime](m.TimeGoing),
		utils.DeepCopy[AssociatedValueType](m.ValueGoing),
		m.reservedField0,
	}
	return _AlarmMessageObjectQueryTypeCopy
}

func (m *_AlarmMessageObjectQueryType) String() string {
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

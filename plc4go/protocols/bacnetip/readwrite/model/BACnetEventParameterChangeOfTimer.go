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

// BACnetEventParameterChangeOfTimer is the corresponding interface of BACnetEventParameterChangeOfTimer
type BACnetEventParameterChangeOfTimer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() BACnetEventParameterChangeOfTimerAlarmValue
	// GetUpdateTimeReference returns UpdateTimeReference (property field)
	GetUpdateTimeReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfTimer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfTimer()
	// CreateBuilder creates a BACnetEventParameterChangeOfTimerBuilder
	CreateBACnetEventParameterChangeOfTimerBuilder() BACnetEventParameterChangeOfTimerBuilder
}

// _BACnetEventParameterChangeOfTimer is the data-structure of this message
type _BACnetEventParameterChangeOfTimer struct {
	BACnetEventParameterContract
	OpeningTag          BACnetOpeningTag
	TimeDelay           BACnetContextTagUnsignedInteger
	AlarmValues         BACnetEventParameterChangeOfTimerAlarmValue
	UpdateTimeReference BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag          BACnetClosingTag
}

var _ BACnetEventParameterChangeOfTimer = (*_BACnetEventParameterChangeOfTimer)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterChangeOfTimer)(nil)

// NewBACnetEventParameterChangeOfTimer factory function for _BACnetEventParameterChangeOfTimer
func NewBACnetEventParameterChangeOfTimer(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, alarmValues BACnetEventParameterChangeOfTimerAlarmValue, updateTimeReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) *_BACnetEventParameterChangeOfTimer {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterChangeOfTimer must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterChangeOfTimer must not be nil")
	}
	if alarmValues == nil {
		panic("alarmValues of type BACnetEventParameterChangeOfTimerAlarmValue for BACnetEventParameterChangeOfTimer must not be nil")
	}
	if updateTimeReference == nil {
		panic("updateTimeReference of type BACnetDeviceObjectPropertyReferenceEnclosed for BACnetEventParameterChangeOfTimer must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterChangeOfTimer must not be nil")
	}
	_result := &_BACnetEventParameterChangeOfTimer{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		AlarmValues:                  alarmValues,
		UpdateTimeReference:          updateTimeReference,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterChangeOfTimerBuilder is a builder for BACnetEventParameterChangeOfTimer
type BACnetEventParameterChangeOfTimerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, alarmValues BACnetEventParameterChangeOfTimerAlarmValue, updateTimeReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) BACnetEventParameterChangeOfTimerBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterChangeOfTimerBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfTimerBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetContextTagUnsignedInteger) BACnetEventParameterChangeOfTimerBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterChangeOfTimerBuilder
	// WithAlarmValues adds AlarmValues (property field)
	WithAlarmValues(BACnetEventParameterChangeOfTimerAlarmValue) BACnetEventParameterChangeOfTimerBuilder
	// WithAlarmValuesBuilder adds AlarmValues (property field) which is build by the builder
	WithAlarmValuesBuilder(func(BACnetEventParameterChangeOfTimerAlarmValueBuilder) BACnetEventParameterChangeOfTimerAlarmValueBuilder) BACnetEventParameterChangeOfTimerBuilder
	// WithUpdateTimeReference adds UpdateTimeReference (property field)
	WithUpdateTimeReference(BACnetDeviceObjectPropertyReferenceEnclosed) BACnetEventParameterChangeOfTimerBuilder
	// WithUpdateTimeReferenceBuilder adds UpdateTimeReference (property field) which is build by the builder
	WithUpdateTimeReferenceBuilder(func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetEventParameterChangeOfTimerBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterChangeOfTimerBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfTimerBuilder
	// Build builds the BACnetEventParameterChangeOfTimer or returns an error if something is wrong
	Build() (BACnetEventParameterChangeOfTimer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterChangeOfTimer
}

// NewBACnetEventParameterChangeOfTimerBuilder() creates a BACnetEventParameterChangeOfTimerBuilder
func NewBACnetEventParameterChangeOfTimerBuilder() BACnetEventParameterChangeOfTimerBuilder {
	return &_BACnetEventParameterChangeOfTimerBuilder{_BACnetEventParameterChangeOfTimer: new(_BACnetEventParameterChangeOfTimer)}
}

type _BACnetEventParameterChangeOfTimerBuilder struct {
	*_BACnetEventParameterChangeOfTimer

	parentBuilder *_BACnetEventParameterBuilder

	err *utils.MultiError
}

var _ (BACnetEventParameterChangeOfTimerBuilder) = (*_BACnetEventParameterChangeOfTimerBuilder)(nil)

func (b *_BACnetEventParameterChangeOfTimerBuilder) setParent(contract BACnetEventParameterContract) {
	b.BACnetEventParameterContract = contract
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, alarmValues BACnetEventParameterChangeOfTimerAlarmValue, updateTimeReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) BACnetEventParameterChangeOfTimerBuilder {
	return b.WithOpeningTag(openingTag).WithTimeDelay(timeDelay).WithAlarmValues(alarmValues).WithUpdateTimeReference(updateTimeReference).WithClosingTag(closingTag)
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterChangeOfTimerBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfTimerBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithTimeDelay(timeDelay BACnetContextTagUnsignedInteger) BACnetEventParameterChangeOfTimerBuilder {
	b.TimeDelay = timeDelay
	return b
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterChangeOfTimerBuilder {
	builder := builderSupplier(b.TimeDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeDelay, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithAlarmValues(alarmValues BACnetEventParameterChangeOfTimerAlarmValue) BACnetEventParameterChangeOfTimerBuilder {
	b.AlarmValues = alarmValues
	return b
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithAlarmValuesBuilder(builderSupplier func(BACnetEventParameterChangeOfTimerAlarmValueBuilder) BACnetEventParameterChangeOfTimerAlarmValueBuilder) BACnetEventParameterChangeOfTimerBuilder {
	builder := builderSupplier(b.AlarmValues.CreateBACnetEventParameterChangeOfTimerAlarmValueBuilder())
	var err error
	b.AlarmValues, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEventParameterChangeOfTimerAlarmValueBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithUpdateTimeReference(updateTimeReference BACnetDeviceObjectPropertyReferenceEnclosed) BACnetEventParameterChangeOfTimerBuilder {
	b.UpdateTimeReference = updateTimeReference
	return b
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithUpdateTimeReferenceBuilder(builderSupplier func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetEventParameterChangeOfTimerBuilder {
	builder := builderSupplier(b.UpdateTimeReference.CreateBACnetDeviceObjectPropertyReferenceEnclosedBuilder())
	var err error
	b.UpdateTimeReference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectPropertyReferenceEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterChangeOfTimerBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfTimerBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) Build() (BACnetEventParameterChangeOfTimer, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.TimeDelay == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeDelay' not set"))
	}
	if b.AlarmValues == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'alarmValues' not set"))
	}
	if b.UpdateTimeReference == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'updateTimeReference' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetEventParameterChangeOfTimer.deepCopy(), nil
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) MustBuild() BACnetEventParameterChangeOfTimer {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetEventParameterChangeOfTimerBuilder) Done() BACnetEventParameterBuilder {
	return b.parentBuilder
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) buildForBACnetEventParameter() (BACnetEventParameter, error) {
	return b.Build()
}

func (b *_BACnetEventParameterChangeOfTimerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterChangeOfTimerBuilder().(*_BACnetEventParameterChangeOfTimerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterChangeOfTimerBuilder creates a BACnetEventParameterChangeOfTimerBuilder
func (b *_BACnetEventParameterChangeOfTimer) CreateBACnetEventParameterChangeOfTimerBuilder() BACnetEventParameterChangeOfTimerBuilder {
	if b == nil {
		return NewBACnetEventParameterChangeOfTimerBuilder()
	}
	return &_BACnetEventParameterChangeOfTimerBuilder{_BACnetEventParameterChangeOfTimer: b.deepCopy()}
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

func (m *_BACnetEventParameterChangeOfTimer) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfTimer) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfTimer) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfTimer) GetAlarmValues() BACnetEventParameterChangeOfTimerAlarmValue {
	return m.AlarmValues
}

func (m *_BACnetEventParameterChangeOfTimer) GetUpdateTimeReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.UpdateTimeReference
}

func (m *_BACnetEventParameterChangeOfTimer) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfTimer(structType any) BACnetEventParameterChangeOfTimer {
	if casted, ok := structType.(BACnetEventParameterChangeOfTimer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfTimer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfTimer) GetTypeName() string {
	return "BACnetEventParameterChangeOfTimer"
}

func (m *_BACnetEventParameterChangeOfTimer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (alarmValues)
	lengthInBits += m.AlarmValues.GetLengthInBits(ctx)

	// Simple field (updateTimeReference)
	lengthInBits += m.UpdateTimeReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfTimer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterChangeOfTimer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterChangeOfTimer BACnetEventParameterChangeOfTimer, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfTimer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfTimer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(22))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	alarmValues, err := ReadSimpleField[BACnetEventParameterChangeOfTimerAlarmValue](ctx, "alarmValues", ReadComplex[BACnetEventParameterChangeOfTimerAlarmValue](BACnetEventParameterChangeOfTimerAlarmValueParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmValues' field"))
	}
	m.AlarmValues = alarmValues

	updateTimeReference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "updateTimeReference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'updateTimeReference' field"))
	}
	m.UpdateTimeReference = updateTimeReference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(22))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfTimer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfTimer")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfTimer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfTimer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfTimer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfTimer")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetEventParameterChangeOfTimerAlarmValue](ctx, "alarmValues", m.GetAlarmValues(), WriteComplex[BACnetEventParameterChangeOfTimerAlarmValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmValues' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "updateTimeReference", m.GetUpdateTimeReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'updateTimeReference' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfTimer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfTimer")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfTimer) IsBACnetEventParameterChangeOfTimer() {}

func (m *_BACnetEventParameterChangeOfTimer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterChangeOfTimer) deepCopy() *_BACnetEventParameterChangeOfTimer {
	if m == nil {
		return nil
	}
	_BACnetEventParameterChangeOfTimerCopy := &_BACnetEventParameterChangeOfTimer{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.TimeDelay.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.AlarmValues.DeepCopy().(BACnetEventParameterChangeOfTimerAlarmValue),
		m.UpdateTimeReference.DeepCopy().(BACnetDeviceObjectPropertyReferenceEnclosed),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterChangeOfTimerCopy
}

func (m *_BACnetEventParameterChangeOfTimer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

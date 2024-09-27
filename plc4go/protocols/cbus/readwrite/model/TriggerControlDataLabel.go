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

// TriggerControlDataLabel is the corresponding interface of TriggerControlDataLabel
type TriggerControlDataLabel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	TriggerControlData
	// GetTriggerControlOptions returns TriggerControlOptions (property field)
	GetTriggerControlOptions() TriggerControlLabelOptions
	// GetActionSelector returns ActionSelector (property field)
	GetActionSelector() byte
	// GetLanguage returns Language (property field)
	GetLanguage() *Language
	// GetData returns Data (property field)
	GetData() []byte
	// IsTriggerControlDataLabel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTriggerControlDataLabel()
	// CreateBuilder creates a TriggerControlDataLabelBuilder
	CreateTriggerControlDataLabelBuilder() TriggerControlDataLabelBuilder
}

// _TriggerControlDataLabel is the data-structure of this message
type _TriggerControlDataLabel struct {
	TriggerControlDataContract
	TriggerControlOptions TriggerControlLabelOptions
	ActionSelector        byte
	Language              *Language
	Data                  []byte
}

var _ TriggerControlDataLabel = (*_TriggerControlDataLabel)(nil)
var _ TriggerControlDataRequirements = (*_TriggerControlDataLabel)(nil)

// NewTriggerControlDataLabel factory function for _TriggerControlDataLabel
func NewTriggerControlDataLabel(commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte, triggerControlOptions TriggerControlLabelOptions, actionSelector byte, language *Language, data []byte) *_TriggerControlDataLabel {
	if triggerControlOptions == nil {
		panic("triggerControlOptions of type TriggerControlLabelOptions for TriggerControlDataLabel must not be nil")
	}
	_result := &_TriggerControlDataLabel{
		TriggerControlDataContract: NewTriggerControlData(commandTypeContainer, triggerGroup),
		TriggerControlOptions:      triggerControlOptions,
		ActionSelector:             actionSelector,
		Language:                   language,
		Data:                       data,
	}
	_result.TriggerControlDataContract.(*_TriggerControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TriggerControlDataLabelBuilder is a builder for TriggerControlDataLabel
type TriggerControlDataLabelBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(triggerControlOptions TriggerControlLabelOptions, actionSelector byte, data []byte) TriggerControlDataLabelBuilder
	// WithTriggerControlOptions adds TriggerControlOptions (property field)
	WithTriggerControlOptions(TriggerControlLabelOptions) TriggerControlDataLabelBuilder
	// WithTriggerControlOptionsBuilder adds TriggerControlOptions (property field) which is build by the builder
	WithTriggerControlOptionsBuilder(func(TriggerControlLabelOptionsBuilder) TriggerControlLabelOptionsBuilder) TriggerControlDataLabelBuilder
	// WithActionSelector adds ActionSelector (property field)
	WithActionSelector(byte) TriggerControlDataLabelBuilder
	// WithLanguage adds Language (property field)
	WithOptionalLanguage(Language) TriggerControlDataLabelBuilder
	// WithData adds Data (property field)
	WithData(...byte) TriggerControlDataLabelBuilder
	// Build builds the TriggerControlDataLabel or returns an error if something is wrong
	Build() (TriggerControlDataLabel, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TriggerControlDataLabel
}

// NewTriggerControlDataLabelBuilder() creates a TriggerControlDataLabelBuilder
func NewTriggerControlDataLabelBuilder() TriggerControlDataLabelBuilder {
	return &_TriggerControlDataLabelBuilder{_TriggerControlDataLabel: new(_TriggerControlDataLabel)}
}

type _TriggerControlDataLabelBuilder struct {
	*_TriggerControlDataLabel

	parentBuilder *_TriggerControlDataBuilder

	err *utils.MultiError
}

var _ (TriggerControlDataLabelBuilder) = (*_TriggerControlDataLabelBuilder)(nil)

func (b *_TriggerControlDataLabelBuilder) setParent(contract TriggerControlDataContract) {
	b.TriggerControlDataContract = contract
}

func (b *_TriggerControlDataLabelBuilder) WithMandatoryFields(triggerControlOptions TriggerControlLabelOptions, actionSelector byte, data []byte) TriggerControlDataLabelBuilder {
	return b.WithTriggerControlOptions(triggerControlOptions).WithActionSelector(actionSelector).WithData(data...)
}

func (b *_TriggerControlDataLabelBuilder) WithTriggerControlOptions(triggerControlOptions TriggerControlLabelOptions) TriggerControlDataLabelBuilder {
	b.TriggerControlOptions = triggerControlOptions
	return b
}

func (b *_TriggerControlDataLabelBuilder) WithTriggerControlOptionsBuilder(builderSupplier func(TriggerControlLabelOptionsBuilder) TriggerControlLabelOptionsBuilder) TriggerControlDataLabelBuilder {
	builder := builderSupplier(b.TriggerControlOptions.CreateTriggerControlLabelOptionsBuilder())
	var err error
	b.TriggerControlOptions, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "TriggerControlLabelOptionsBuilder failed"))
	}
	return b
}

func (b *_TriggerControlDataLabelBuilder) WithActionSelector(actionSelector byte) TriggerControlDataLabelBuilder {
	b.ActionSelector = actionSelector
	return b
}

func (b *_TriggerControlDataLabelBuilder) WithOptionalLanguage(language Language) TriggerControlDataLabelBuilder {
	b.Language = &language
	return b
}

func (b *_TriggerControlDataLabelBuilder) WithData(data ...byte) TriggerControlDataLabelBuilder {
	b.Data = data
	return b
}

func (b *_TriggerControlDataLabelBuilder) Build() (TriggerControlDataLabel, error) {
	if b.TriggerControlOptions == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'triggerControlOptions' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TriggerControlDataLabel.deepCopy(), nil
}

func (b *_TriggerControlDataLabelBuilder) MustBuild() TriggerControlDataLabel {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_TriggerControlDataLabelBuilder) Done() TriggerControlDataBuilder {
	return b.parentBuilder
}

func (b *_TriggerControlDataLabelBuilder) buildForTriggerControlData() (TriggerControlData, error) {
	return b.Build()
}

func (b *_TriggerControlDataLabelBuilder) DeepCopy() any {
	_copy := b.CreateTriggerControlDataLabelBuilder().(*_TriggerControlDataLabelBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTriggerControlDataLabelBuilder creates a TriggerControlDataLabelBuilder
func (b *_TriggerControlDataLabel) CreateTriggerControlDataLabelBuilder() TriggerControlDataLabelBuilder {
	if b == nil {
		return NewTriggerControlDataLabelBuilder()
	}
	return &_TriggerControlDataLabelBuilder{_TriggerControlDataLabel: b.deepCopy()}
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

func (m *_TriggerControlDataLabel) GetParent() TriggerControlDataContract {
	return m.TriggerControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TriggerControlDataLabel) GetTriggerControlOptions() TriggerControlLabelOptions {
	return m.TriggerControlOptions
}

func (m *_TriggerControlDataLabel) GetActionSelector() byte {
	return m.ActionSelector
}

func (m *_TriggerControlDataLabel) GetLanguage() *Language {
	return m.Language
}

func (m *_TriggerControlDataLabel) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTriggerControlDataLabel(structType any) TriggerControlDataLabel {
	if casted, ok := structType.(TriggerControlDataLabel); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlDataLabel); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlDataLabel) GetTypeName() string {
	return "TriggerControlDataLabel"
}

func (m *_TriggerControlDataLabel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TriggerControlDataContract.(*_TriggerControlData).getLengthInBits(ctx))

	// Simple field (triggerControlOptions)
	lengthInBits += m.TriggerControlOptions.GetLengthInBits(ctx)

	// Simple field (actionSelector)
	lengthInBits += 8

	// Optional Field (language)
	if m.Language != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_TriggerControlDataLabel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TriggerControlDataLabel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TriggerControlData, commandTypeContainer TriggerControlCommandTypeContainer) (__triggerControlDataLabel TriggerControlDataLabel, err error) {
	m.TriggerControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TriggerControlDataLabel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlDataLabel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	triggerControlOptions, err := ReadSimpleField[TriggerControlLabelOptions](ctx, "triggerControlOptions", ReadComplex[TriggerControlLabelOptions](TriggerControlLabelOptionsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'triggerControlOptions' field"))
	}
	m.TriggerControlOptions = triggerControlOptions

	actionSelector, err := ReadSimpleField(ctx, "actionSelector", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actionSelector' field"))
	}
	m.ActionSelector = actionSelector

	var language *Language
	language, err = ReadOptionalField[Language](ctx, "language", ReadEnum(LanguageByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool((triggerControlOptions.GetLabelType()) != (TriggerControlLabelType_LOAD_DYNAMIC_ICON)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'language' field"))
	}
	m.Language = language

	data, err := readBuffer.ReadByteArray("data", int((int32(commandTypeContainer.NumBytes()) - int32((utils.InlineIf((bool((triggerControlOptions.GetLabelType()) != (TriggerControlLabelType_LOAD_DYNAMIC_ICON))), func() any { return int32((int32(4))) }, func() any { return int32((int32(3))) }).(int32))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("TriggerControlDataLabel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlDataLabel")
	}

	return m, nil
}

func (m *_TriggerControlDataLabel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TriggerControlDataLabel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TriggerControlDataLabel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TriggerControlDataLabel")
		}

		if err := WriteSimpleField[TriggerControlLabelOptions](ctx, "triggerControlOptions", m.GetTriggerControlOptions(), WriteComplex[TriggerControlLabelOptions](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'triggerControlOptions' field")
		}

		if err := WriteSimpleField[byte](ctx, "actionSelector", m.GetActionSelector(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'actionSelector' field")
		}

		if err := WriteOptionalEnumField[Language](ctx, "language", "Language", m.GetLanguage(), WriteEnum[Language, uint8](Language.GetValue, Language.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool((m.GetTriggerControlOptions().GetLabelType()) != (TriggerControlLabelType_LOAD_DYNAMIC_ICON))); err != nil {
			return errors.Wrap(err, "Error serializing 'language' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("TriggerControlDataLabel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TriggerControlDataLabel")
		}
		return nil
	}
	return m.TriggerControlDataContract.(*_TriggerControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TriggerControlDataLabel) IsTriggerControlDataLabel() {}

func (m *_TriggerControlDataLabel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TriggerControlDataLabel) deepCopy() *_TriggerControlDataLabel {
	if m == nil {
		return nil
	}
	_TriggerControlDataLabelCopy := &_TriggerControlDataLabel{
		m.TriggerControlDataContract.(*_TriggerControlData).deepCopy(),
		m.TriggerControlOptions.DeepCopy().(TriggerControlLabelOptions),
		m.ActionSelector,
		utils.CopyPtr[Language](m.Language),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.TriggerControlDataContract.(*_TriggerControlData)._SubType = m
	return _TriggerControlDataLabelCopy
}

func (m *_TriggerControlDataLabel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

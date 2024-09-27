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

// HVACTemperature is the corresponding interface of HVACTemperature
type HVACTemperature interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetTemperatureValue returns TemperatureValue (property field)
	GetTemperatureValue() int16
	// GetTemperatureInCelcius returns TemperatureInCelcius (virtual field)
	GetTemperatureInCelcius() float32
	// IsHVACTemperature is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHVACTemperature()
	// CreateBuilder creates a HVACTemperatureBuilder
	CreateHVACTemperatureBuilder() HVACTemperatureBuilder
}

// _HVACTemperature is the data-structure of this message
type _HVACTemperature struct {
	TemperatureValue int16
}

var _ HVACTemperature = (*_HVACTemperature)(nil)

// NewHVACTemperature factory function for _HVACTemperature
func NewHVACTemperature(temperatureValue int16) *_HVACTemperature {
	return &_HVACTemperature{TemperatureValue: temperatureValue}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HVACTemperatureBuilder is a builder for HVACTemperature
type HVACTemperatureBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(temperatureValue int16) HVACTemperatureBuilder
	// WithTemperatureValue adds TemperatureValue (property field)
	WithTemperatureValue(int16) HVACTemperatureBuilder
	// Build builds the HVACTemperature or returns an error if something is wrong
	Build() (HVACTemperature, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HVACTemperature
}

// NewHVACTemperatureBuilder() creates a HVACTemperatureBuilder
func NewHVACTemperatureBuilder() HVACTemperatureBuilder {
	return &_HVACTemperatureBuilder{_HVACTemperature: new(_HVACTemperature)}
}

type _HVACTemperatureBuilder struct {
	*_HVACTemperature

	err *utils.MultiError
}

var _ (HVACTemperatureBuilder) = (*_HVACTemperatureBuilder)(nil)

func (b *_HVACTemperatureBuilder) WithMandatoryFields(temperatureValue int16) HVACTemperatureBuilder {
	return b.WithTemperatureValue(temperatureValue)
}

func (b *_HVACTemperatureBuilder) WithTemperatureValue(temperatureValue int16) HVACTemperatureBuilder {
	b.TemperatureValue = temperatureValue
	return b
}

func (b *_HVACTemperatureBuilder) Build() (HVACTemperature, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HVACTemperature.deepCopy(), nil
}

func (b *_HVACTemperatureBuilder) MustBuild() HVACTemperature {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HVACTemperatureBuilder) DeepCopy() any {
	_copy := b.CreateHVACTemperatureBuilder().(*_HVACTemperatureBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHVACTemperatureBuilder creates a HVACTemperatureBuilder
func (b *_HVACTemperature) CreateHVACTemperatureBuilder() HVACTemperatureBuilder {
	if b == nil {
		return NewHVACTemperatureBuilder()
	}
	return &_HVACTemperatureBuilder{_HVACTemperature: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACTemperature) GetTemperatureValue() int16 {
	return m.TemperatureValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACTemperature) GetTemperatureInCelcius() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(float32(m.GetTemperatureValue()) / float32(float32(256)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHVACTemperature(structType any) HVACTemperature {
	if casted, ok := structType.(HVACTemperature); ok {
		return casted
	}
	if casted, ok := structType.(*HVACTemperature); ok {
		return *casted
	}
	return nil
}

func (m *_HVACTemperature) GetTypeName() string {
	return "HVACTemperature"
}

func (m *_HVACTemperature) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (temperatureValue)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_HVACTemperature) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACTemperatureParse(ctx context.Context, theBytes []byte) (HVACTemperature, error) {
	return HVACTemperatureParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACTemperatureParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACTemperature, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACTemperature, error) {
		return HVACTemperatureParseWithBuffer(ctx, readBuffer)
	}
}

func HVACTemperatureParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACTemperature, error) {
	v, err := (&_HVACTemperature{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_HVACTemperature) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__hVACTemperature HVACTemperature, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACTemperature"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACTemperature")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	temperatureValue, err := ReadSimpleField(ctx, "temperatureValue", ReadSignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'temperatureValue' field"))
	}
	m.TemperatureValue = temperatureValue

	temperatureInCelcius, err := ReadVirtualField[float32](ctx, "temperatureInCelcius", (*float32)(nil), float32(temperatureValue)/float32(float32(256)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'temperatureInCelcius' field"))
	}
	_ = temperatureInCelcius

	if closeErr := readBuffer.CloseContext("HVACTemperature"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACTemperature")
	}

	return m, nil
}

func (m *_HVACTemperature) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACTemperature) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HVACTemperature"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACTemperature")
	}

	if err := WriteSimpleField[int16](ctx, "temperatureValue", m.GetTemperatureValue(), WriteSignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'temperatureValue' field")
	}
	// Virtual field
	temperatureInCelcius := m.GetTemperatureInCelcius()
	_ = temperatureInCelcius
	if _temperatureInCelciusErr := writeBuffer.WriteVirtual(ctx, "temperatureInCelcius", m.GetTemperatureInCelcius()); _temperatureInCelciusErr != nil {
		return errors.Wrap(_temperatureInCelciusErr, "Error serializing 'temperatureInCelcius' field")
	}

	if popErr := writeBuffer.PopContext("HVACTemperature"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACTemperature")
	}
	return nil
}

func (m *_HVACTemperature) IsHVACTemperature() {}

func (m *_HVACTemperature) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HVACTemperature) deepCopy() *_HVACTemperature {
	if m == nil {
		return nil
	}
	_HVACTemperatureCopy := &_HVACTemperature{
		m.TemperatureValue,
	}
	return _HVACTemperatureCopy
}

func (m *_HVACTemperature) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

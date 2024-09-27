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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ParameterValue is the corresponding interface of ParameterValue
type ParameterValue interface {
	ParameterValueContract
	ParameterValueRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsParameterValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValue()
	// CreateBuilder creates a ParameterValueBuilder
	CreateParameterValueBuilder() ParameterValueBuilder
}

// ParameterValueContract provides a set of functions which can be overwritten by a sub struct
type ParameterValueContract interface {
	// GetNumBytes() returns a parser argument
	GetNumBytes() uint8
	// IsParameterValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValue()
	// CreateBuilder creates a ParameterValueBuilder
	CreateParameterValueBuilder() ParameterValueBuilder
}

// ParameterValueRequirements provides a set of functions which need to be implemented by a sub struct
type ParameterValueRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetParameterType returns ParameterType (discriminator field)
	GetParameterType() ParameterType
}

// _ParameterValue is the data-structure of this message
type _ParameterValue struct {
	_SubType ParameterValue

	// Arguments.
	NumBytes uint8
}

var _ ParameterValueContract = (*_ParameterValue)(nil)

// NewParameterValue factory function for _ParameterValue
func NewParameterValue(numBytes uint8) *_ParameterValue {
	return &_ParameterValue{NumBytes: numBytes}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ParameterValueBuilder is a builder for ParameterValue
type ParameterValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ParameterValueBuilder
	// AsParameterValueApplicationAddress1 converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueApplicationAddress1() interface {
		ParameterValueApplicationAddress1Builder
		Done() ParameterValueBuilder
	}
	// AsParameterValueApplicationAddress2 converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueApplicationAddress2() interface {
		ParameterValueApplicationAddress2Builder
		Done() ParameterValueBuilder
	}
	// AsParameterValueInterfaceOptions1 converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueInterfaceOptions1() interface {
		ParameterValueInterfaceOptions1Builder
		Done() ParameterValueBuilder
	}
	// AsParameterValueBaudRateSelector converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueBaudRateSelector() interface {
		ParameterValueBaudRateSelectorBuilder
		Done() ParameterValueBuilder
	}
	// AsParameterValueInterfaceOptions2 converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueInterfaceOptions2() interface {
		ParameterValueInterfaceOptions2Builder
		Done() ParameterValueBuilder
	}
	// AsParameterValueInterfaceOptions1PowerUpSettings converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueInterfaceOptions1PowerUpSettings() interface {
		ParameterValueInterfaceOptions1PowerUpSettingsBuilder
		Done() ParameterValueBuilder
	}
	// AsParameterValueInterfaceOptions3 converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueInterfaceOptions3() interface {
		ParameterValueInterfaceOptions3Builder
		Done() ParameterValueBuilder
	}
	// AsParameterValueCustomManufacturer converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueCustomManufacturer() interface {
		ParameterValueCustomManufacturerBuilder
		Done() ParameterValueBuilder
	}
	// AsParameterValueSerialNumber converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueSerialNumber() interface {
		ParameterValueSerialNumberBuilder
		Done() ParameterValueBuilder
	}
	// AsParameterValueCustomTypes converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueCustomTypes() interface {
		ParameterValueCustomTypesBuilder
		Done() ParameterValueBuilder
	}
	// AsParameterValueRaw converts this build to a subType of ParameterValue. It is always possible to return to current builder using Done()
	AsParameterValueRaw() interface {
		ParameterValueRawBuilder
		Done() ParameterValueBuilder
	}
	// Build builds the ParameterValue or returns an error if something is wrong
	PartialBuild() (ParameterValueContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() ParameterValueContract
	// Build builds the ParameterValue or returns an error if something is wrong
	Build() (ParameterValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ParameterValue
}

// NewParameterValueBuilder() creates a ParameterValueBuilder
func NewParameterValueBuilder() ParameterValueBuilder {
	return &_ParameterValueBuilder{_ParameterValue: new(_ParameterValue)}
}

type _ParameterValueChildBuilder interface {
	utils.Copyable
	setParent(ParameterValueContract)
	buildForParameterValue() (ParameterValue, error)
}

type _ParameterValueBuilder struct {
	*_ParameterValue

	childBuilder _ParameterValueChildBuilder

	err *utils.MultiError
}

var _ (ParameterValueBuilder) = (*_ParameterValueBuilder)(nil)

func (b *_ParameterValueBuilder) WithMandatoryFields() ParameterValueBuilder {
	return b
}

func (b *_ParameterValueBuilder) PartialBuild() (ParameterValueContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ParameterValue.deepCopy(), nil
}

func (b *_ParameterValueBuilder) PartialMustBuild() ParameterValueContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ParameterValueBuilder) AsParameterValueApplicationAddress1() interface {
	ParameterValueApplicationAddress1Builder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueApplicationAddress1Builder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueApplicationAddress1Builder().(*_ParameterValueApplicationAddress1Builder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) AsParameterValueApplicationAddress2() interface {
	ParameterValueApplicationAddress2Builder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueApplicationAddress2Builder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueApplicationAddress2Builder().(*_ParameterValueApplicationAddress2Builder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) AsParameterValueInterfaceOptions1() interface {
	ParameterValueInterfaceOptions1Builder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueInterfaceOptions1Builder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueInterfaceOptions1Builder().(*_ParameterValueInterfaceOptions1Builder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) AsParameterValueBaudRateSelector() interface {
	ParameterValueBaudRateSelectorBuilder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueBaudRateSelectorBuilder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueBaudRateSelectorBuilder().(*_ParameterValueBaudRateSelectorBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) AsParameterValueInterfaceOptions2() interface {
	ParameterValueInterfaceOptions2Builder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueInterfaceOptions2Builder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueInterfaceOptions2Builder().(*_ParameterValueInterfaceOptions2Builder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) AsParameterValueInterfaceOptions1PowerUpSettings() interface {
	ParameterValueInterfaceOptions1PowerUpSettingsBuilder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueInterfaceOptions1PowerUpSettingsBuilder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueInterfaceOptions1PowerUpSettingsBuilder().(*_ParameterValueInterfaceOptions1PowerUpSettingsBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) AsParameterValueInterfaceOptions3() interface {
	ParameterValueInterfaceOptions3Builder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueInterfaceOptions3Builder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueInterfaceOptions3Builder().(*_ParameterValueInterfaceOptions3Builder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) AsParameterValueCustomManufacturer() interface {
	ParameterValueCustomManufacturerBuilder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueCustomManufacturerBuilder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueCustomManufacturerBuilder().(*_ParameterValueCustomManufacturerBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) AsParameterValueSerialNumber() interface {
	ParameterValueSerialNumberBuilder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueSerialNumberBuilder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueSerialNumberBuilder().(*_ParameterValueSerialNumberBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) AsParameterValueCustomTypes() interface {
	ParameterValueCustomTypesBuilder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueCustomTypesBuilder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueCustomTypesBuilder().(*_ParameterValueCustomTypesBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) AsParameterValueRaw() interface {
	ParameterValueRawBuilder
	Done() ParameterValueBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ParameterValueRawBuilder
		Done() ParameterValueBuilder
	}); ok {
		return cb
	}
	cb := NewParameterValueRawBuilder().(*_ParameterValueRawBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ParameterValueBuilder) Build() (ParameterValue, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForParameterValue()
}

func (b *_ParameterValueBuilder) MustBuild() ParameterValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ParameterValueBuilder) DeepCopy() any {
	_copy := b.CreateParameterValueBuilder().(*_ParameterValueBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_ParameterValueChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateParameterValueBuilder creates a ParameterValueBuilder
func (b *_ParameterValue) CreateParameterValueBuilder() ParameterValueBuilder {
	if b == nil {
		return NewParameterValueBuilder()
	}
	return &_ParameterValueBuilder{_ParameterValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastParameterValue(structType any) ParameterValue {
	if casted, ok := structType.(ParameterValue); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValue); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValue) GetTypeName() string {
	return "ParameterValue"
}

func (m *_ParameterValue) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ParameterValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ParameterValueParse[T ParameterValue](ctx context.Context, theBytes []byte, parameterType ParameterType, numBytes uint8) (T, error) {
	return ParameterValueParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), parameterType, numBytes)
}

func ParameterValueParseWithBufferProducer[T ParameterValue](parameterType ParameterType, numBytes uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ParameterValueParseWithBuffer[T](ctx, readBuffer, parameterType, numBytes)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ParameterValueParseWithBuffer[T ParameterValue](ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (T, error) {
	v, err := (&_ParameterValue{NumBytes: numBytes}).parse(ctx, readBuffer, parameterType, numBytes)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_ParameterValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (__parameterValue ParameterValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ParameterValue
	switch {
	case parameterType == ParameterType_APPLICATION_ADDRESS_1: // ParameterValueApplicationAddress1
		if _child, err = new(_ParameterValueApplicationAddress1).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueApplicationAddress1 for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_APPLICATION_ADDRESS_2: // ParameterValueApplicationAddress2
		if _child, err = new(_ParameterValueApplicationAddress2).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueApplicationAddress2 for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_INTERFACE_OPTIONS_1: // ParameterValueInterfaceOptions1
		if _child, err = new(_ParameterValueInterfaceOptions1).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueInterfaceOptions1 for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_BAUD_RATE_SELECTOR: // ParameterValueBaudRateSelector
		if _child, err = new(_ParameterValueBaudRateSelector).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueBaudRateSelector for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_INTERFACE_OPTIONS_2: // ParameterValueInterfaceOptions2
		if _child, err = new(_ParameterValueInterfaceOptions2).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueInterfaceOptions2 for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_INTERFACE_OPTIONS_1_POWER_UP_SETTINGS: // ParameterValueInterfaceOptions1PowerUpSettings
		if _child, err = new(_ParameterValueInterfaceOptions1PowerUpSettings).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueInterfaceOptions1PowerUpSettings for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_INTERFACE_OPTIONS_3: // ParameterValueInterfaceOptions3
		if _child, err = new(_ParameterValueInterfaceOptions3).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueInterfaceOptions3 for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_CUSTOM_MANUFACTURER: // ParameterValueCustomManufacturer
		if _child, err = new(_ParameterValueCustomManufacturer).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueCustomManufacturer for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_SERIAL_NUMBER: // ParameterValueSerialNumber
		if _child, err = new(_ParameterValueSerialNumber).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueSerialNumber for type-switch of ParameterValue")
		}
	case parameterType == ParameterType_CUSTOM_TYPE: // ParameterValueCustomTypes
		if _child, err = new(_ParameterValueCustomTypes).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueCustomTypes for type-switch of ParameterValue")
		}
	case 0 == 0: // ParameterValueRaw
		if _child, err = new(_ParameterValueRaw).parse(ctx, readBuffer, m, parameterType, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ParameterValueRaw for type-switch of ParameterValue")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [parameterType=%v]", parameterType)
	}

	if closeErr := readBuffer.CloseContext("ParameterValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValue")
	}

	return _child, nil
}

func (pm *_ParameterValue) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ParameterValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ParameterValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ParameterValue")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ParameterValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ParameterValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_ParameterValue) GetNumBytes() uint8 {
	return m.NumBytes
}

//
////

func (m *_ParameterValue) IsParameterValue() {}

func (m *_ParameterValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ParameterValue) deepCopy() *_ParameterValue {
	if m == nil {
		return nil
	}
	_ParameterValueCopy := &_ParameterValue{
		nil, // will be set by child
		m.NumBytes,
	}
	return _ParameterValueCopy
}

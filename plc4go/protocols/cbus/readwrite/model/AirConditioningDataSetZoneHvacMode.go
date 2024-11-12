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

// AirConditioningDataSetZoneHvacMode is the corresponding interface of AirConditioningDataSetZoneHvacMode
type AirConditioningDataSetZoneHvacMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACModeAndFlags
	// GetHvacType returns HvacType (property field)
	GetHvacType() HVACType
	// GetLevel returns Level (property field)
	GetLevel() HVACTemperature
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
	// GetAuxLevel returns AuxLevel (property field)
	GetAuxLevel() HVACAuxiliaryLevel
	// IsAirConditioningDataSetZoneHvacMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataSetZoneHvacMode()
	// CreateBuilder creates a AirConditioningDataSetZoneHvacModeBuilder
	CreateAirConditioningDataSetZoneHvacModeBuilder() AirConditioningDataSetZoneHvacModeBuilder
}

// _AirConditioningDataSetZoneHvacMode is the data-structure of this message
type _AirConditioningDataSetZoneHvacMode struct {
	AirConditioningDataContract
	ZoneGroup        byte
	ZoneList         HVACZoneList
	HvacModeAndFlags HVACModeAndFlags
	HvacType         HVACType
	Level            HVACTemperature
	RawLevel         HVACRawLevels
	AuxLevel         HVACAuxiliaryLevel
}

var _ AirConditioningDataSetZoneHvacMode = (*_AirConditioningDataSetZoneHvacMode)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataSetZoneHvacMode)(nil)

// NewAirConditioningDataSetZoneHvacMode factory function for _AirConditioningDataSetZoneHvacMode
func NewAirConditioningDataSetZoneHvacMode(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte, zoneList HVACZoneList, hvacModeAndFlags HVACModeAndFlags, hvacType HVACType, level HVACTemperature, rawLevel HVACRawLevels, auxLevel HVACAuxiliaryLevel) *_AirConditioningDataSetZoneHvacMode {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataSetZoneHvacMode must not be nil")
	}
	if hvacModeAndFlags == nil {
		panic("hvacModeAndFlags of type HVACModeAndFlags for AirConditioningDataSetZoneHvacMode must not be nil")
	}
	_result := &_AirConditioningDataSetZoneHvacMode{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		HvacModeAndFlags:            hvacModeAndFlags,
		HvacType:                    hvacType,
		Level:                       level,
		RawLevel:                    rawLevel,
		AuxLevel:                    auxLevel,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AirConditioningDataSetZoneHvacModeBuilder is a builder for AirConditioningDataSetZoneHvacMode
type AirConditioningDataSetZoneHvacModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, hvacModeAndFlags HVACModeAndFlags, hvacType HVACType) AirConditioningDataSetZoneHvacModeBuilder
	// WithZoneGroup adds ZoneGroup (property field)
	WithZoneGroup(byte) AirConditioningDataSetZoneHvacModeBuilder
	// WithZoneList adds ZoneList (property field)
	WithZoneList(HVACZoneList) AirConditioningDataSetZoneHvacModeBuilder
	// WithZoneListBuilder adds ZoneList (property field) which is build by the builder
	WithZoneListBuilder(func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataSetZoneHvacModeBuilder
	// WithHvacModeAndFlags adds HvacModeAndFlags (property field)
	WithHvacModeAndFlags(HVACModeAndFlags) AirConditioningDataSetZoneHvacModeBuilder
	// WithHvacModeAndFlagsBuilder adds HvacModeAndFlags (property field) which is build by the builder
	WithHvacModeAndFlagsBuilder(func(HVACModeAndFlagsBuilder) HVACModeAndFlagsBuilder) AirConditioningDataSetZoneHvacModeBuilder
	// WithHvacType adds HvacType (property field)
	WithHvacType(HVACType) AirConditioningDataSetZoneHvacModeBuilder
	// WithLevel adds Level (property field)
	WithOptionalLevel(HVACTemperature) AirConditioningDataSetZoneHvacModeBuilder
	// WithOptionalLevelBuilder adds Level (property field) which is build by the builder
	WithOptionalLevelBuilder(func(HVACTemperatureBuilder) HVACTemperatureBuilder) AirConditioningDataSetZoneHvacModeBuilder
	// WithRawLevel adds RawLevel (property field)
	WithOptionalRawLevel(HVACRawLevels) AirConditioningDataSetZoneHvacModeBuilder
	// WithOptionalRawLevelBuilder adds RawLevel (property field) which is build by the builder
	WithOptionalRawLevelBuilder(func(HVACRawLevelsBuilder) HVACRawLevelsBuilder) AirConditioningDataSetZoneHvacModeBuilder
	// WithAuxLevel adds AuxLevel (property field)
	WithOptionalAuxLevel(HVACAuxiliaryLevel) AirConditioningDataSetZoneHvacModeBuilder
	// WithOptionalAuxLevelBuilder adds AuxLevel (property field) which is build by the builder
	WithOptionalAuxLevelBuilder(func(HVACAuxiliaryLevelBuilder) HVACAuxiliaryLevelBuilder) AirConditioningDataSetZoneHvacModeBuilder
	// Build builds the AirConditioningDataSetZoneHvacMode or returns an error if something is wrong
	Build() (AirConditioningDataSetZoneHvacMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AirConditioningDataSetZoneHvacMode
}

// NewAirConditioningDataSetZoneHvacModeBuilder() creates a AirConditioningDataSetZoneHvacModeBuilder
func NewAirConditioningDataSetZoneHvacModeBuilder() AirConditioningDataSetZoneHvacModeBuilder {
	return &_AirConditioningDataSetZoneHvacModeBuilder{_AirConditioningDataSetZoneHvacMode: new(_AirConditioningDataSetZoneHvacMode)}
}

type _AirConditioningDataSetZoneHvacModeBuilder struct {
	*_AirConditioningDataSetZoneHvacMode

	parentBuilder *_AirConditioningDataBuilder

	err *utils.MultiError
}

var _ (AirConditioningDataSetZoneHvacModeBuilder) = (*_AirConditioningDataSetZoneHvacModeBuilder)(nil)

func (b *_AirConditioningDataSetZoneHvacModeBuilder) setParent(contract AirConditioningDataContract) {
	b.AirConditioningDataContract = contract
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, hvacModeAndFlags HVACModeAndFlags, hvacType HVACType) AirConditioningDataSetZoneHvacModeBuilder {
	return b.WithZoneGroup(zoneGroup).WithZoneList(zoneList).WithHvacModeAndFlags(hvacModeAndFlags).WithHvacType(hvacType)
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithZoneGroup(zoneGroup byte) AirConditioningDataSetZoneHvacModeBuilder {
	b.ZoneGroup = zoneGroup
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithZoneList(zoneList HVACZoneList) AirConditioningDataSetZoneHvacModeBuilder {
	b.ZoneList = zoneList
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithZoneListBuilder(builderSupplier func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataSetZoneHvacModeBuilder {
	builder := builderSupplier(b.ZoneList.CreateHVACZoneListBuilder())
	var err error
	b.ZoneList, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACZoneListBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithHvacModeAndFlags(hvacModeAndFlags HVACModeAndFlags) AirConditioningDataSetZoneHvacModeBuilder {
	b.HvacModeAndFlags = hvacModeAndFlags
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithHvacModeAndFlagsBuilder(builderSupplier func(HVACModeAndFlagsBuilder) HVACModeAndFlagsBuilder) AirConditioningDataSetZoneHvacModeBuilder {
	builder := builderSupplier(b.HvacModeAndFlags.CreateHVACModeAndFlagsBuilder())
	var err error
	b.HvacModeAndFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACModeAndFlagsBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithHvacType(hvacType HVACType) AirConditioningDataSetZoneHvacModeBuilder {
	b.HvacType = hvacType
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithOptionalLevel(level HVACTemperature) AirConditioningDataSetZoneHvacModeBuilder {
	b.Level = level
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithOptionalLevelBuilder(builderSupplier func(HVACTemperatureBuilder) HVACTemperatureBuilder) AirConditioningDataSetZoneHvacModeBuilder {
	builder := builderSupplier(b.Level.CreateHVACTemperatureBuilder())
	var err error
	b.Level, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACTemperatureBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithOptionalRawLevel(rawLevel HVACRawLevels) AirConditioningDataSetZoneHvacModeBuilder {
	b.RawLevel = rawLevel
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithOptionalRawLevelBuilder(builderSupplier func(HVACRawLevelsBuilder) HVACRawLevelsBuilder) AirConditioningDataSetZoneHvacModeBuilder {
	builder := builderSupplier(b.RawLevel.CreateHVACRawLevelsBuilder())
	var err error
	b.RawLevel, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACRawLevelsBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithOptionalAuxLevel(auxLevel HVACAuxiliaryLevel) AirConditioningDataSetZoneHvacModeBuilder {
	b.AuxLevel = auxLevel
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) WithOptionalAuxLevelBuilder(builderSupplier func(HVACAuxiliaryLevelBuilder) HVACAuxiliaryLevelBuilder) AirConditioningDataSetZoneHvacModeBuilder {
	builder := builderSupplier(b.AuxLevel.CreateHVACAuxiliaryLevelBuilder())
	var err error
	b.AuxLevel, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACAuxiliaryLevelBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) Build() (AirConditioningDataSetZoneHvacMode, error) {
	if b.ZoneList == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'zoneList' not set"))
	}
	if b.HvacModeAndFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'hvacModeAndFlags' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AirConditioningDataSetZoneHvacMode.deepCopy(), nil
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) MustBuild() AirConditioningDataSetZoneHvacMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AirConditioningDataSetZoneHvacModeBuilder) Done() AirConditioningDataBuilder {
	return b.parentBuilder
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) buildForAirConditioningData() (AirConditioningData, error) {
	return b.Build()
}

func (b *_AirConditioningDataSetZoneHvacModeBuilder) DeepCopy() any {
	_copy := b.CreateAirConditioningDataSetZoneHvacModeBuilder().(*_AirConditioningDataSetZoneHvacModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAirConditioningDataSetZoneHvacModeBuilder creates a AirConditioningDataSetZoneHvacModeBuilder
func (b *_AirConditioningDataSetZoneHvacMode) CreateAirConditioningDataSetZoneHvacModeBuilder() AirConditioningDataSetZoneHvacModeBuilder {
	if b == nil {
		return NewAirConditioningDataSetZoneHvacModeBuilder()
	}
	return &_AirConditioningDataSetZoneHvacModeBuilder{_AirConditioningDataSetZoneHvacMode: b.deepCopy()}
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

func (m *_AirConditioningDataSetZoneHvacMode) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetZoneHvacMode) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetZoneHvacMode) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetZoneHvacMode) GetHvacModeAndFlags() HVACModeAndFlags {
	return m.HvacModeAndFlags
}

func (m *_AirConditioningDataSetZoneHvacMode) GetHvacType() HVACType {
	return m.HvacType
}

func (m *_AirConditioningDataSetZoneHvacMode) GetLevel() HVACTemperature {
	return m.Level
}

func (m *_AirConditioningDataSetZoneHvacMode) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

func (m *_AirConditioningDataSetZoneHvacMode) GetAuxLevel() HVACAuxiliaryLevel {
	return m.AuxLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetZoneHvacMode(structType any) AirConditioningDataSetZoneHvacMode {
	if casted, ok := structType.(AirConditioningDataSetZoneHvacMode); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetZoneHvacMode); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetZoneHvacMode) GetTypeName() string {
	return "AirConditioningDataSetZoneHvacMode"
}

func (m *_AirConditioningDataSetZoneHvacMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	// Simple field (hvacType)
	lengthInBits += 8

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits(ctx)
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits(ctx)
	}

	// Optional Field (auxLevel)
	if m.AuxLevel != nil {
		lengthInBits += m.AuxLevel.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_AirConditioningDataSetZoneHvacMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetZoneHvacMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetZoneHvacMode AirConditioningDataSetZoneHvacMode, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetZoneHvacMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetZoneHvacMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}
	m.ZoneList = zoneList

	hvacModeAndFlags, err := ReadSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", ReadComplex[HVACModeAndFlags](HVACModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacModeAndFlags' field"))
	}
	m.HvacModeAndFlags = hvacModeAndFlags

	hvacType, err := ReadEnumField[HVACType](ctx, "hvacType", "HVACType", ReadEnum(HVACTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacType' field"))
	}
	m.HvacType = hvacType

	var level HVACTemperature
	_level, err := ReadOptionalField[HVACTemperature](ctx, "level", ReadComplex[HVACTemperature](HVACTemperatureParseWithBuffer, readBuffer), hvacModeAndFlags.GetIsLevelTemperature())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	if _level != nil {
		level = *_level
		m.Level = level
	}

	var rawLevel HVACRawLevels
	_rawLevel, err := ReadOptionalField[HVACRawLevels](ctx, "rawLevel", ReadComplex[HVACRawLevels](HVACRawLevelsParseWithBuffer, readBuffer), hvacModeAndFlags.GetIsLevelRaw())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rawLevel' field"))
	}
	if _rawLevel != nil {
		rawLevel = *_rawLevel
		m.RawLevel = rawLevel
	}

	var auxLevel HVACAuxiliaryLevel
	_auxLevel, err := ReadOptionalField[HVACAuxiliaryLevel](ctx, "auxLevel", ReadComplex[HVACAuxiliaryLevel](HVACAuxiliaryLevelParseWithBuffer, readBuffer), hvacModeAndFlags.GetIsAuxLevelUsed())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'auxLevel' field"))
	}
	if _auxLevel != nil {
		auxLevel = *_auxLevel
		m.AuxLevel = auxLevel
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetZoneHvacMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetZoneHvacMode")
	}

	return m, nil
}

func (m *_AirConditioningDataSetZoneHvacMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetZoneHvacMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetZoneHvacMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetZoneHvacMode")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", m.GetHvacModeAndFlags(), WriteComplex[HVACModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacModeAndFlags' field")
		}

		if err := WriteSimpleEnumField[HVACType](ctx, "hvacType", "HVACType", m.GetHvacType(), WriteEnum[HVACType, uint8](HVACType.GetValue, HVACType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacType' field")
		}

		if err := WriteOptionalField[HVACTemperature](ctx, "level", GetRef(m.GetLevel()), WriteComplex[HVACTemperature](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if err := WriteOptionalField[HVACRawLevels](ctx, "rawLevel", GetRef(m.GetRawLevel()), WriteComplex[HVACRawLevels](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'rawLevel' field")
		}

		if err := WriteOptionalField[HVACAuxiliaryLevel](ctx, "auxLevel", GetRef(m.GetAuxLevel()), WriteComplex[HVACAuxiliaryLevel](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'auxLevel' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetZoneHvacMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetZoneHvacMode")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetZoneHvacMode) IsAirConditioningDataSetZoneHvacMode() {}

func (m *_AirConditioningDataSetZoneHvacMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataSetZoneHvacMode) deepCopy() *_AirConditioningDataSetZoneHvacMode {
	if m == nil {
		return nil
	}
	_AirConditioningDataSetZoneHvacModeCopy := &_AirConditioningDataSetZoneHvacMode{
		m.AirConditioningDataContract.(*_AirConditioningData).deepCopy(),
		m.ZoneGroup,
		utils.DeepCopy[HVACZoneList](m.ZoneList),
		utils.DeepCopy[HVACModeAndFlags](m.HvacModeAndFlags),
		m.HvacType,
		utils.DeepCopy[HVACTemperature](m.Level),
		utils.DeepCopy[HVACRawLevels](m.RawLevel),
		utils.DeepCopy[HVACAuxiliaryLevel](m.AuxLevel),
	}
	m.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataSetZoneHvacModeCopy
}

func (m *_AirConditioningDataSetZoneHvacMode) String() string {
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

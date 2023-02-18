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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AirConditioningDataZoneHumidityPlantStatus is the corresponding interface of AirConditioningDataZoneHumidityPlantStatus
type AirConditioningDataZoneHumidityPlantStatus interface {
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHumidityType returns HumidityType (property field)
	GetHumidityType() HVACHumidityType
	// GetHumidityStatus returns HumidityStatus (property field)
	GetHumidityStatus() HVACHumidityStatusFlags
	// GetHumidityErrorCode returns HumidityErrorCode (property field)
	GetHumidityErrorCode() HVACHumidityError
}

// AirConditioningDataZoneHumidityPlantStatusExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataZoneHumidityPlantStatus.
// This is useful for switch cases.
type AirConditioningDataZoneHumidityPlantStatusExactly interface {
	AirConditioningDataZoneHumidityPlantStatus
	isAirConditioningDataZoneHumidityPlantStatus() bool
}

// _AirConditioningDataZoneHumidityPlantStatus is the data-structure of this message
type _AirConditioningDataZoneHumidityPlantStatus struct {
	*_AirConditioningData
	ZoneGroup         byte
	ZoneList          HVACZoneList
	HumidityType      HVACHumidityType
	HumidityStatus    HVACHumidityStatusFlags
	HumidityErrorCode HVACHumidityError
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataZoneHumidityPlantStatus) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetHumidityType() HVACHumidityType {
	return m.HumidityType
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetHumidityStatus() HVACHumidityStatusFlags {
	return m.HumidityStatus
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetHumidityErrorCode() HVACHumidityError {
	return m.HumidityErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataZoneHumidityPlantStatus factory function for _AirConditioningDataZoneHumidityPlantStatus
func NewAirConditioningDataZoneHumidityPlantStatus(zoneGroup byte, zoneList HVACZoneList, humidityType HVACHumidityType, humidityStatus HVACHumidityStatusFlags, humidityErrorCode HVACHumidityError, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataZoneHumidityPlantStatus {
	_result := &_AirConditioningDataZoneHumidityPlantStatus{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		HumidityType:         humidityType,
		HumidityStatus:       humidityStatus,
		HumidityErrorCode:    humidityErrorCode,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataZoneHumidityPlantStatus(structType interface{}) AirConditioningDataZoneHumidityPlantStatus {
	if casted, ok := structType.(AirConditioningDataZoneHumidityPlantStatus); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataZoneHumidityPlantStatus); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetTypeName() string {
	return "AirConditioningDataZoneHumidityPlantStatus"
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (humidityType)
	lengthInBits += 8

	// Simple field (humidityStatus)
	lengthInBits += m.HumidityStatus.GetLengthInBits(ctx)

	// Simple field (humidityErrorCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningDataZoneHumidityPlantStatusParse(theBytes []byte) (AirConditioningDataZoneHumidityPlantStatus, error) {
	return AirConditioningDataZoneHumidityPlantStatusParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningDataZoneHumidityPlantStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataZoneHumidityPlantStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataZoneHumidityPlantStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataZoneHumidityPlantStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
	_zoneGroup, _zoneGroupErr := readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataZoneHumidityPlantStatus")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
	_zoneList, _zoneListErr := HVACZoneListParseWithBuffer(ctx, readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataZoneHumidityPlantStatus")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (humidityType)
	if pullErr := readBuffer.PullContext("humidityType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for humidityType")
	}
	_humidityType, _humidityTypeErr := HVACHumidityTypeParseWithBuffer(ctx, readBuffer)
	if _humidityTypeErr != nil {
		return nil, errors.Wrap(_humidityTypeErr, "Error parsing 'humidityType' field of AirConditioningDataZoneHumidityPlantStatus")
	}
	humidityType := _humidityType
	if closeErr := readBuffer.CloseContext("humidityType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for humidityType")
	}

	// Simple Field (humidityStatus)
	if pullErr := readBuffer.PullContext("humidityStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for humidityStatus")
	}
	_humidityStatus, _humidityStatusErr := HVACHumidityStatusFlagsParseWithBuffer(ctx, readBuffer)
	if _humidityStatusErr != nil {
		return nil, errors.Wrap(_humidityStatusErr, "Error parsing 'humidityStatus' field of AirConditioningDataZoneHumidityPlantStatus")
	}
	humidityStatus := _humidityStatus.(HVACHumidityStatusFlags)
	if closeErr := readBuffer.CloseContext("humidityStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for humidityStatus")
	}

	// Simple Field (humidityErrorCode)
	if pullErr := readBuffer.PullContext("humidityErrorCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for humidityErrorCode")
	}
	_humidityErrorCode, _humidityErrorCodeErr := HVACHumidityErrorParseWithBuffer(ctx, readBuffer)
	if _humidityErrorCodeErr != nil {
		return nil, errors.Wrap(_humidityErrorCodeErr, "Error parsing 'humidityErrorCode' field of AirConditioningDataZoneHumidityPlantStatus")
	}
	humidityErrorCode := _humidityErrorCode
	if closeErr := readBuffer.CloseContext("humidityErrorCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for humidityErrorCode")
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataZoneHumidityPlantStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataZoneHumidityPlantStatus")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataZoneHumidityPlantStatus{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		HumidityType:         humidityType,
		HumidityStatus:       humidityStatus,
		HumidityErrorCode:    humidityErrorCode,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataZoneHumidityPlantStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataZoneHumidityPlantStatus")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		// Simple Field (zoneList)
		if pushErr := writeBuffer.PushContext("zoneList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneList")
		}
		_zoneListErr := writeBuffer.WriteSerializable(ctx, m.GetZoneList())
		if popErr := writeBuffer.PopContext("zoneList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneList")
		}
		if _zoneListErr != nil {
			return errors.Wrap(_zoneListErr, "Error serializing 'zoneList' field")
		}

		// Simple Field (humidityType)
		if pushErr := writeBuffer.PushContext("humidityType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for humidityType")
		}
		_humidityTypeErr := writeBuffer.WriteSerializable(ctx, m.GetHumidityType())
		if popErr := writeBuffer.PopContext("humidityType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for humidityType")
		}
		if _humidityTypeErr != nil {
			return errors.Wrap(_humidityTypeErr, "Error serializing 'humidityType' field")
		}

		// Simple Field (humidityStatus)
		if pushErr := writeBuffer.PushContext("humidityStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for humidityStatus")
		}
		_humidityStatusErr := writeBuffer.WriteSerializable(ctx, m.GetHumidityStatus())
		if popErr := writeBuffer.PopContext("humidityStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for humidityStatus")
		}
		if _humidityStatusErr != nil {
			return errors.Wrap(_humidityStatusErr, "Error serializing 'humidityStatus' field")
		}

		// Simple Field (humidityErrorCode)
		if pushErr := writeBuffer.PushContext("humidityErrorCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for humidityErrorCode")
		}
		_humidityErrorCodeErr := writeBuffer.WriteSerializable(ctx, m.GetHumidityErrorCode())
		if popErr := writeBuffer.PopContext("humidityErrorCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for humidityErrorCode")
		}
		if _humidityErrorCodeErr != nil {
			return errors.Wrap(_humidityErrorCodeErr, "Error serializing 'humidityErrorCode' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataZoneHumidityPlantStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataZoneHumidityPlantStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) isAirConditioningDataZoneHumidityPlantStatus() bool {
	return true
}

func (m *_AirConditioningDataZoneHumidityPlantStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

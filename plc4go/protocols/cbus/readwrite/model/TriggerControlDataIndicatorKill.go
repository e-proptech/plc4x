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

// TriggerControlDataIndicatorKill is the corresponding interface of TriggerControlDataIndicatorKill
type TriggerControlDataIndicatorKill interface {
	utils.LengthAware
	utils.Serializable
	TriggerControlData
}

// TriggerControlDataIndicatorKillExactly can be used when we want exactly this type and not a type which fulfills TriggerControlDataIndicatorKill.
// This is useful for switch cases.
type TriggerControlDataIndicatorKillExactly interface {
	TriggerControlDataIndicatorKill
	isTriggerControlDataIndicatorKill() bool
}

// _TriggerControlDataIndicatorKill is the data-structure of this message
type _TriggerControlDataIndicatorKill struct {
	*_TriggerControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TriggerControlDataIndicatorKill) InitializeParent(parent TriggerControlData, commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.TriggerGroup = triggerGroup
}

func (m *_TriggerControlDataIndicatorKill) GetParent() TriggerControlData {
	return m._TriggerControlData
}

// NewTriggerControlDataIndicatorKill factory function for _TriggerControlDataIndicatorKill
func NewTriggerControlDataIndicatorKill(commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte) *_TriggerControlDataIndicatorKill {
	_result := &_TriggerControlDataIndicatorKill{
		_TriggerControlData: NewTriggerControlData(commandTypeContainer, triggerGroup),
	}
	_result._TriggerControlData._TriggerControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTriggerControlDataIndicatorKill(structType interface{}) TriggerControlDataIndicatorKill {
	if casted, ok := structType.(TriggerControlDataIndicatorKill); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlDataIndicatorKill); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlDataIndicatorKill) GetTypeName() string {
	return "TriggerControlDataIndicatorKill"
}

func (m *_TriggerControlDataIndicatorKill) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_TriggerControlDataIndicatorKill) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TriggerControlDataIndicatorKillParse(theBytes []byte) (TriggerControlDataIndicatorKill, error) {
	return TriggerControlDataIndicatorKillParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func TriggerControlDataIndicatorKillParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlDataIndicatorKill, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TriggerControlDataIndicatorKill"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlDataIndicatorKill")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TriggerControlDataIndicatorKill"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlDataIndicatorKill")
	}

	// Create a partially initialized instance
	_child := &_TriggerControlDataIndicatorKill{
		_TriggerControlData: &_TriggerControlData{},
	}
	_child._TriggerControlData._TriggerControlDataChildRequirements = _child
	return _child, nil
}

func (m *_TriggerControlDataIndicatorKill) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TriggerControlDataIndicatorKill) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TriggerControlDataIndicatorKill"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TriggerControlDataIndicatorKill")
		}

		if popErr := writeBuffer.PopContext("TriggerControlDataIndicatorKill"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TriggerControlDataIndicatorKill")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TriggerControlDataIndicatorKill) isTriggerControlDataIndicatorKill() bool {
	return true
}

func (m *_TriggerControlDataIndicatorKill) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

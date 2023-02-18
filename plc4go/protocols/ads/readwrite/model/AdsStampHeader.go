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
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsStampHeader is the corresponding interface of AdsStampHeader
type AdsStampHeader interface {
	utils.LengthAware
	utils.Serializable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() uint64
	// GetSamples returns Samples (property field)
	GetSamples() uint32
	// GetAdsNotificationSamples returns AdsNotificationSamples (property field)
	GetAdsNotificationSamples() []AdsNotificationSample
}

// AdsStampHeaderExactly can be used when we want exactly this type and not a type which fulfills AdsStampHeader.
// This is useful for switch cases.
type AdsStampHeaderExactly interface {
	AdsStampHeader
	isAdsStampHeader() bool
}

// _AdsStampHeader is the data-structure of this message
type _AdsStampHeader struct {
	Timestamp              uint64
	Samples                uint32
	AdsNotificationSamples []AdsNotificationSample
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsStampHeader) GetTimestamp() uint64 {
	return m.Timestamp
}

func (m *_AdsStampHeader) GetSamples() uint32 {
	return m.Samples
}

func (m *_AdsStampHeader) GetAdsNotificationSamples() []AdsNotificationSample {
	return m.AdsNotificationSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsStampHeader factory function for _AdsStampHeader
func NewAdsStampHeader(timestamp uint64, samples uint32, adsNotificationSamples []AdsNotificationSample) *_AdsStampHeader {
	return &_AdsStampHeader{Timestamp: timestamp, Samples: samples, AdsNotificationSamples: adsNotificationSamples}
}

// Deprecated: use the interface for direct cast
func CastAdsStampHeader(structType interface{}) AdsStampHeader {
	if casted, ok := structType.(AdsStampHeader); ok {
		return casted
	}
	if casted, ok := structType.(*AdsStampHeader); ok {
		return *casted
	}
	return nil
}

func (m *_AdsStampHeader) GetTypeName() string {
	return "AdsStampHeader"
}

func (m *_AdsStampHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += 64

	// Simple field (samples)
	lengthInBits += 32

	// Array field
	if len(m.AdsNotificationSamples) > 0 {
		for _curItem, element := range m.AdsNotificationSamples {
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.AdsNotificationSamples), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AdsStampHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsStampHeaderParse(theBytes []byte) (AdsStampHeader, error) {
	return AdsStampHeaderParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AdsStampHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsStampHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsStampHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsStampHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timestamp)
	_timestamp, _timestampErr := readBuffer.ReadUint64("timestamp", 64)
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field of AdsStampHeader")
	}
	timestamp := _timestamp

	// Simple Field (samples)
	_samples, _samplesErr := readBuffer.ReadUint32("samples", 32)
	if _samplesErr != nil {
		return nil, errors.Wrap(_samplesErr, "Error parsing 'samples' field of AdsStampHeader")
	}
	samples := _samples

	// Array field (adsNotificationSamples)
	if pullErr := readBuffer.PullContext("adsNotificationSamples", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for adsNotificationSamples")
	}
	// Count array
	adsNotificationSamples := make([]AdsNotificationSample, samples)
	// This happens when the size is set conditional to 0
	if len(adsNotificationSamples) == 0 {
		adsNotificationSamples = nil
	}
	{
		_numItems := uint16(samples)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := spiContext.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := AdsNotificationSampleParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'adsNotificationSamples' field of AdsStampHeader")
			}
			adsNotificationSamples[_curItem] = _item.(AdsNotificationSample)
		}
	}
	if closeErr := readBuffer.CloseContext("adsNotificationSamples", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for adsNotificationSamples")
	}

	if closeErr := readBuffer.CloseContext("AdsStampHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsStampHeader")
	}

	// Create the instance
	return &_AdsStampHeader{
		Timestamp:              timestamp,
		Samples:                samples,
		AdsNotificationSamples: adsNotificationSamples,
	}, nil
}

func (m *_AdsStampHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsStampHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AdsStampHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsStampHeader")
	}

	// Simple Field (timestamp)
	timestamp := uint64(m.GetTimestamp())
	_timestampErr := writeBuffer.WriteUint64("timestamp", 64, (timestamp))
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (samples)
	samples := uint32(m.GetSamples())
	_samplesErr := writeBuffer.WriteUint32("samples", 32, (samples))
	if _samplesErr != nil {
		return errors.Wrap(_samplesErr, "Error serializing 'samples' field")
	}

	// Array Field (adsNotificationSamples)
	if pushErr := writeBuffer.PushContext("adsNotificationSamples", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for adsNotificationSamples")
	}
	for _curItem, _element := range m.GetAdsNotificationSamples() {
		_ = _curItem
		arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetAdsNotificationSamples()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'adsNotificationSamples' field")
		}
	}
	if popErr := writeBuffer.PopContext("adsNotificationSamples", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for adsNotificationSamples")
	}

	if popErr := writeBuffer.PopContext("AdsStampHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsStampHeader")
	}
	return nil
}

func (m *_AdsStampHeader) isAdsStampHeader() bool {
	return true
}

func (m *_AdsStampHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

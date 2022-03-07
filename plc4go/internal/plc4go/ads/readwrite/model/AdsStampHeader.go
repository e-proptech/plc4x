/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type AdsStampHeader struct {
	Timestamp              uint64
	Samples                uint32
	AdsNotificationSamples []*AdsNotificationSample
}

// The corresponding interface
type IAdsStampHeader interface {
	// GetTimestamp returns Timestamp
	GetTimestamp() uint64
	// GetSamples returns Samples
	GetSamples() uint32
	// GetAdsNotificationSamples returns AdsNotificationSamples
	GetAdsNotificationSamples() []*AdsNotificationSample
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *AdsStampHeader) GetTimestamp() uint64 {
	return m.Timestamp
}

func (m *AdsStampHeader) GetSamples() uint32 {
	return m.Samples
}

func (m *AdsStampHeader) GetAdsNotificationSamples() []*AdsNotificationSample {
	return m.AdsNotificationSamples
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAdsStampHeader factory function for AdsStampHeader
func NewAdsStampHeader(timestamp uint64, samples uint32, adsNotificationSamples []*AdsNotificationSample) *AdsStampHeader {
	return &AdsStampHeader{Timestamp: timestamp, Samples: samples, AdsNotificationSamples: adsNotificationSamples}
}

func CastAdsStampHeader(structType interface{}) *AdsStampHeader {
	if casted, ok := structType.(AdsStampHeader); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsStampHeader); ok {
		return casted
	}
	return nil
}

func (m *AdsStampHeader) GetTypeName() string {
	return "AdsStampHeader"
}

func (m *AdsStampHeader) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsStampHeader) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += 64

	// Simple field (samples)
	lengthInBits += 32

	// Array field
	if len(m.AdsNotificationSamples) > 0 {
		for i, element := range m.AdsNotificationSamples {
			last := i == len(m.AdsNotificationSamples)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *AdsStampHeader) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsStampHeaderParse(readBuffer utils.ReadBuffer) (*AdsStampHeader, error) {
	if pullErr := readBuffer.PullContext("AdsStampHeader"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (timestamp)
	_timestamp, _timestampErr := readBuffer.ReadUint64("timestamp", 64)
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field")
	}
	timestamp := _timestamp

	// Simple Field (samples)
	_samples, _samplesErr := readBuffer.ReadUint32("samples", 32)
	if _samplesErr != nil {
		return nil, errors.Wrap(_samplesErr, "Error parsing 'samples' field")
	}
	samples := _samples

	// Array field (adsNotificationSamples)
	if pullErr := readBuffer.PullContext("adsNotificationSamples", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	adsNotificationSamples := make([]*AdsNotificationSample, samples)
	{
		for curItem := uint16(0); curItem < uint16(samples); curItem++ {
			_item, _err := AdsNotificationSampleParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'adsNotificationSamples' field")
			}
			adsNotificationSamples[curItem] = CastAdsNotificationSample(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("adsNotificationSamples", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AdsStampHeader"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAdsStampHeader(timestamp, samples, adsNotificationSamples), nil
}

func (m *AdsStampHeader) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AdsStampHeader"); pushErr != nil {
		return pushErr
	}

	// Simple Field (timestamp)
	timestamp := uint64(m.Timestamp)
	_timestampErr := writeBuffer.WriteUint64("timestamp", 64, (timestamp))
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (samples)
	samples := uint32(m.Samples)
	_samplesErr := writeBuffer.WriteUint32("samples", 32, (samples))
	if _samplesErr != nil {
		return errors.Wrap(_samplesErr, "Error serializing 'samples' field")
	}

	// Array Field (adsNotificationSamples)
	if m.AdsNotificationSamples != nil {
		if pushErr := writeBuffer.PushContext("adsNotificationSamples", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.AdsNotificationSamples {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'adsNotificationSamples' field")
			}
		}
		if popErr := writeBuffer.PopContext("adsNotificationSamples", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("AdsStampHeader"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AdsStampHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

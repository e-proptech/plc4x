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

// AdsReadStateRequest is the corresponding interface of AdsReadStateRequest
type AdsReadStateRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// IsAdsReadStateRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsReadStateRequest()
	// CreateBuilder creates a AdsReadStateRequestBuilder
	CreateAdsReadStateRequestBuilder() AdsReadStateRequestBuilder
}

// _AdsReadStateRequest is the data-structure of this message
type _AdsReadStateRequest struct {
	AmsPacketContract
}

var _ AdsReadStateRequest = (*_AdsReadStateRequest)(nil)
var _ AmsPacketRequirements = (*_AdsReadStateRequest)(nil)

// NewAdsReadStateRequest factory function for _AdsReadStateRequest
func NewAdsReadStateRequest(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsReadStateRequest {
	_result := &_AdsReadStateRequest{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsReadStateRequestBuilder is a builder for AdsReadStateRequest
type AdsReadStateRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() AdsReadStateRequestBuilder
	// Build builds the AdsReadStateRequest or returns an error if something is wrong
	Build() (AdsReadStateRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsReadStateRequest
}

// NewAdsReadStateRequestBuilder() creates a AdsReadStateRequestBuilder
func NewAdsReadStateRequestBuilder() AdsReadStateRequestBuilder {
	return &_AdsReadStateRequestBuilder{_AdsReadStateRequest: new(_AdsReadStateRequest)}
}

type _AdsReadStateRequestBuilder struct {
	*_AdsReadStateRequest

	parentBuilder *_AmsPacketBuilder

	err *utils.MultiError
}

var _ (AdsReadStateRequestBuilder) = (*_AdsReadStateRequestBuilder)(nil)

func (b *_AdsReadStateRequestBuilder) setParent(contract AmsPacketContract) {
	b.AmsPacketContract = contract
}

func (b *_AdsReadStateRequestBuilder) WithMandatoryFields() AdsReadStateRequestBuilder {
	return b
}

func (b *_AdsReadStateRequestBuilder) Build() (AdsReadStateRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsReadStateRequest.deepCopy(), nil
}

func (b *_AdsReadStateRequestBuilder) MustBuild() AdsReadStateRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AdsReadStateRequestBuilder) Done() AmsPacketBuilder {
	return b.parentBuilder
}

func (b *_AdsReadStateRequestBuilder) buildForAmsPacket() (AmsPacket, error) {
	return b.Build()
}

func (b *_AdsReadStateRequestBuilder) DeepCopy() any {
	_copy := b.CreateAdsReadStateRequestBuilder().(*_AdsReadStateRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsReadStateRequestBuilder creates a AdsReadStateRequestBuilder
func (b *_AdsReadStateRequest) CreateAdsReadStateRequestBuilder() AdsReadStateRequestBuilder {
	if b == nil {
		return NewAdsReadStateRequestBuilder()
	}
	return &_AdsReadStateRequestBuilder{_AdsReadStateRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadStateRequest) GetCommandId() CommandId {
	return CommandId_ADS_READ_STATE
}

func (m *_AdsReadStateRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadStateRequest) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

// Deprecated: use the interface for direct cast
func CastAdsReadStateRequest(structType any) AdsReadStateRequest {
	if casted, ok := structType.(AdsReadStateRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadStateRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadStateRequest) GetTypeName() string {
	return "AdsReadStateRequest"
}

func (m *_AdsReadStateRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AdsReadStateRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsReadStateRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsReadStateRequest AdsReadStateRequest, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadStateRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadStateRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsReadStateRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadStateRequest")
	}

	return m, nil
}

func (m *_AdsReadStateRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsReadStateRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadStateRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadStateRequest")
		}

		if popErr := writeBuffer.PopContext("AdsReadStateRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadStateRequest")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsReadStateRequest) IsAdsReadStateRequest() {}

func (m *_AdsReadStateRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsReadStateRequest) deepCopy() *_AdsReadStateRequest {
	if m == nil {
		return nil
	}
	_AdsReadStateRequestCopy := &_AdsReadStateRequest{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
	}
	m.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsReadStateRequestCopy
}

func (m *_AdsReadStateRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

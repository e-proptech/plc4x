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

// GetAttributeSingleRequest is the corresponding interface of GetAttributeSingleRequest
type GetAttributeSingleRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// IsGetAttributeSingleRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGetAttributeSingleRequest()
	// CreateBuilder creates a GetAttributeSingleRequestBuilder
	CreateGetAttributeSingleRequestBuilder() GetAttributeSingleRequestBuilder
}

// _GetAttributeSingleRequest is the data-structure of this message
type _GetAttributeSingleRequest struct {
	CipServiceContract
}

var _ GetAttributeSingleRequest = (*_GetAttributeSingleRequest)(nil)
var _ CipServiceRequirements = (*_GetAttributeSingleRequest)(nil)

// NewGetAttributeSingleRequest factory function for _GetAttributeSingleRequest
func NewGetAttributeSingleRequest(serviceLen uint16) *_GetAttributeSingleRequest {
	_result := &_GetAttributeSingleRequest{
		CipServiceContract: NewCipService(serviceLen),
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GetAttributeSingleRequestBuilder is a builder for GetAttributeSingleRequest
type GetAttributeSingleRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() GetAttributeSingleRequestBuilder
	// Build builds the GetAttributeSingleRequest or returns an error if something is wrong
	Build() (GetAttributeSingleRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() GetAttributeSingleRequest
}

// NewGetAttributeSingleRequestBuilder() creates a GetAttributeSingleRequestBuilder
func NewGetAttributeSingleRequestBuilder() GetAttributeSingleRequestBuilder {
	return &_GetAttributeSingleRequestBuilder{_GetAttributeSingleRequest: new(_GetAttributeSingleRequest)}
}

type _GetAttributeSingleRequestBuilder struct {
	*_GetAttributeSingleRequest

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (GetAttributeSingleRequestBuilder) = (*_GetAttributeSingleRequestBuilder)(nil)

func (b *_GetAttributeSingleRequestBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
}

func (b *_GetAttributeSingleRequestBuilder) WithMandatoryFields() GetAttributeSingleRequestBuilder {
	return b
}

func (b *_GetAttributeSingleRequestBuilder) Build() (GetAttributeSingleRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._GetAttributeSingleRequest.deepCopy(), nil
}

func (b *_GetAttributeSingleRequestBuilder) MustBuild() GetAttributeSingleRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_GetAttributeSingleRequestBuilder) Done() CipServiceBuilder {
	return b.parentBuilder
}

func (b *_GetAttributeSingleRequestBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_GetAttributeSingleRequestBuilder) DeepCopy() any {
	_copy := b.CreateGetAttributeSingleRequestBuilder().(*_GetAttributeSingleRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateGetAttributeSingleRequestBuilder creates a GetAttributeSingleRequestBuilder
func (b *_GetAttributeSingleRequest) CreateGetAttributeSingleRequestBuilder() GetAttributeSingleRequestBuilder {
	if b == nil {
		return NewGetAttributeSingleRequestBuilder()
	}
	return &_GetAttributeSingleRequestBuilder{_GetAttributeSingleRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeSingleRequest) GetService() uint8 {
	return 0x0E
}

func (m *_GetAttributeSingleRequest) GetResponse() bool {
	return bool(false)
}

func (m *_GetAttributeSingleRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeSingleRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

// Deprecated: use the interface for direct cast
func CastGetAttributeSingleRequest(structType any) GetAttributeSingleRequest {
	if casted, ok := structType.(GetAttributeSingleRequest); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeSingleRequest); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeSingleRequest) GetTypeName() string {
	return "GetAttributeSingleRequest"
}

func (m *_GetAttributeSingleRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_GetAttributeSingleRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GetAttributeSingleRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__getAttributeSingleRequest GetAttributeSingleRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeSingleRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeSingleRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("GetAttributeSingleRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeSingleRequest")
	}

	return m, nil
}

func (m *_GetAttributeSingleRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeSingleRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeSingleRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeSingleRequest")
		}

		if popErr := writeBuffer.PopContext("GetAttributeSingleRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeSingleRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeSingleRequest) IsGetAttributeSingleRequest() {}

func (m *_GetAttributeSingleRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GetAttributeSingleRequest) deepCopy() *_GetAttributeSingleRequest {
	if m == nil {
		return nil
	}
	_GetAttributeSingleRequestCopy := &_GetAttributeSingleRequest{
		m.CipServiceContract.(*_CipService).deepCopy(),
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _GetAttributeSingleRequestCopy
}

func (m *_GetAttributeSingleRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

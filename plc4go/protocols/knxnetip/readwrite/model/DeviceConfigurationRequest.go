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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DeviceConfigurationRequest is the corresponding interface of DeviceConfigurationRequest
type DeviceConfigurationRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetDeviceConfigurationRequestDataBlock returns DeviceConfigurationRequestDataBlock (property field)
	GetDeviceConfigurationRequestDataBlock() DeviceConfigurationRequestDataBlock
	// GetCemi returns Cemi (property field)
	GetCemi() CEMI
	// IsDeviceConfigurationRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeviceConfigurationRequest()
	// CreateBuilder creates a DeviceConfigurationRequestBuilder
	CreateDeviceConfigurationRequestBuilder() DeviceConfigurationRequestBuilder
}

// _DeviceConfigurationRequest is the data-structure of this message
type _DeviceConfigurationRequest struct {
	KnxNetIpMessageContract
	DeviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock
	Cemi                                CEMI

	// Arguments.
	TotalLength uint16
}

var _ DeviceConfigurationRequest = (*_DeviceConfigurationRequest)(nil)
var _ KnxNetIpMessageRequirements = (*_DeviceConfigurationRequest)(nil)

// NewDeviceConfigurationRequest factory function for _DeviceConfigurationRequest
func NewDeviceConfigurationRequest(deviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock, cemi CEMI, totalLength uint16) *_DeviceConfigurationRequest {
	if deviceConfigurationRequestDataBlock == nil {
		panic("deviceConfigurationRequestDataBlock of type DeviceConfigurationRequestDataBlock for DeviceConfigurationRequest must not be nil")
	}
	if cemi == nil {
		panic("cemi of type CEMI for DeviceConfigurationRequest must not be nil")
	}
	_result := &_DeviceConfigurationRequest{
		KnxNetIpMessageContract:             NewKnxNetIpMessage(),
		DeviceConfigurationRequestDataBlock: deviceConfigurationRequestDataBlock,
		Cemi:                                cemi,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeviceConfigurationRequestBuilder is a builder for DeviceConfigurationRequest
type DeviceConfigurationRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(deviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock, cemi CEMI) DeviceConfigurationRequestBuilder
	// WithDeviceConfigurationRequestDataBlock adds DeviceConfigurationRequestDataBlock (property field)
	WithDeviceConfigurationRequestDataBlock(DeviceConfigurationRequestDataBlock) DeviceConfigurationRequestBuilder
	// WithDeviceConfigurationRequestDataBlockBuilder adds DeviceConfigurationRequestDataBlock (property field) which is build by the builder
	WithDeviceConfigurationRequestDataBlockBuilder(func(DeviceConfigurationRequestDataBlockBuilder) DeviceConfigurationRequestDataBlockBuilder) DeviceConfigurationRequestBuilder
	// WithCemi adds Cemi (property field)
	WithCemi(CEMI) DeviceConfigurationRequestBuilder
	// WithCemiBuilder adds Cemi (property field) which is build by the builder
	WithCemiBuilder(func(CEMIBuilder) CEMIBuilder) DeviceConfigurationRequestBuilder
	// WithArgTotalLength sets a parser argument
	WithArgTotalLength(uint16) DeviceConfigurationRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() KnxNetIpMessageBuilder
	// Build builds the DeviceConfigurationRequest or returns an error if something is wrong
	Build() (DeviceConfigurationRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeviceConfigurationRequest
}

// NewDeviceConfigurationRequestBuilder() creates a DeviceConfigurationRequestBuilder
func NewDeviceConfigurationRequestBuilder() DeviceConfigurationRequestBuilder {
	return &_DeviceConfigurationRequestBuilder{_DeviceConfigurationRequest: new(_DeviceConfigurationRequest)}
}

type _DeviceConfigurationRequestBuilder struct {
	*_DeviceConfigurationRequest

	parentBuilder *_KnxNetIpMessageBuilder

	err *utils.MultiError
}

var _ (DeviceConfigurationRequestBuilder) = (*_DeviceConfigurationRequestBuilder)(nil)

func (b *_DeviceConfigurationRequestBuilder) setParent(contract KnxNetIpMessageContract) {
	b.KnxNetIpMessageContract = contract
	contract.(*_KnxNetIpMessage)._SubType = b._DeviceConfigurationRequest
}

func (b *_DeviceConfigurationRequestBuilder) WithMandatoryFields(deviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock, cemi CEMI) DeviceConfigurationRequestBuilder {
	return b.WithDeviceConfigurationRequestDataBlock(deviceConfigurationRequestDataBlock).WithCemi(cemi)
}

func (b *_DeviceConfigurationRequestBuilder) WithDeviceConfigurationRequestDataBlock(deviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock) DeviceConfigurationRequestBuilder {
	b.DeviceConfigurationRequestDataBlock = deviceConfigurationRequestDataBlock
	return b
}

func (b *_DeviceConfigurationRequestBuilder) WithDeviceConfigurationRequestDataBlockBuilder(builderSupplier func(DeviceConfigurationRequestDataBlockBuilder) DeviceConfigurationRequestDataBlockBuilder) DeviceConfigurationRequestBuilder {
	builder := builderSupplier(b.DeviceConfigurationRequestDataBlock.CreateDeviceConfigurationRequestDataBlockBuilder())
	var err error
	b.DeviceConfigurationRequestDataBlock, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DeviceConfigurationRequestDataBlockBuilder failed"))
	}
	return b
}

func (b *_DeviceConfigurationRequestBuilder) WithCemi(cemi CEMI) DeviceConfigurationRequestBuilder {
	b.Cemi = cemi
	return b
}

func (b *_DeviceConfigurationRequestBuilder) WithCemiBuilder(builderSupplier func(CEMIBuilder) CEMIBuilder) DeviceConfigurationRequestBuilder {
	builder := builderSupplier(b.Cemi.CreateCEMIBuilder())
	var err error
	b.Cemi, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CEMIBuilder failed"))
	}
	return b
}

func (b *_DeviceConfigurationRequestBuilder) WithArgTotalLength(totalLength uint16) DeviceConfigurationRequestBuilder {
	b.TotalLength = totalLength
	return b
}

func (b *_DeviceConfigurationRequestBuilder) Build() (DeviceConfigurationRequest, error) {
	if b.DeviceConfigurationRequestDataBlock == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deviceConfigurationRequestDataBlock' not set"))
	}
	if b.Cemi == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'cemi' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DeviceConfigurationRequest.deepCopy(), nil
}

func (b *_DeviceConfigurationRequestBuilder) MustBuild() DeviceConfigurationRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DeviceConfigurationRequestBuilder) Done() KnxNetIpMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewKnxNetIpMessageBuilder().(*_KnxNetIpMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_DeviceConfigurationRequestBuilder) buildForKnxNetIpMessage() (KnxNetIpMessage, error) {
	return b.Build()
}

func (b *_DeviceConfigurationRequestBuilder) DeepCopy() any {
	_copy := b.CreateDeviceConfigurationRequestBuilder().(*_DeviceConfigurationRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDeviceConfigurationRequestBuilder creates a DeviceConfigurationRequestBuilder
func (b *_DeviceConfigurationRequest) CreateDeviceConfigurationRequestBuilder() DeviceConfigurationRequestBuilder {
	if b == nil {
		return NewDeviceConfigurationRequestBuilder()
	}
	return &_DeviceConfigurationRequestBuilder{_DeviceConfigurationRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeviceConfigurationRequest) GetMsgType() uint16 {
	return 0x0310
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeviceConfigurationRequest) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationRequest) GetDeviceConfigurationRequestDataBlock() DeviceConfigurationRequestDataBlock {
	return m.DeviceConfigurationRequestDataBlock
}

func (m *_DeviceConfigurationRequest) GetCemi() CEMI {
	return m.Cemi
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationRequest(structType any) DeviceConfigurationRequest {
	if casted, ok := structType.(DeviceConfigurationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationRequest) GetTypeName() string {
	return "DeviceConfigurationRequest"
}

func (m *_DeviceConfigurationRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (deviceConfigurationRequestDataBlock)
	lengthInBits += m.DeviceConfigurationRequestDataBlock.GetLengthInBits(ctx)

	// Simple field (cemi)
	lengthInBits += m.Cemi.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DeviceConfigurationRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeviceConfigurationRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage, totalLength uint16) (__deviceConfigurationRequest DeviceConfigurationRequest, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceConfigurationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceConfigurationRequestDataBlock, err := ReadSimpleField[DeviceConfigurationRequestDataBlock](ctx, "deviceConfigurationRequestDataBlock", ReadComplex[DeviceConfigurationRequestDataBlock](DeviceConfigurationRequestDataBlockParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceConfigurationRequestDataBlock' field"))
	}
	m.DeviceConfigurationRequestDataBlock = deviceConfigurationRequestDataBlock

	cemi, err := ReadSimpleField[CEMI](ctx, "cemi", ReadComplex[CEMI](CEMIParseWithBufferProducer[CEMI]((uint16)(uint16(totalLength)-uint16((uint16(uint16(6))+uint16(deviceConfigurationRequestDataBlock.GetLengthInBytes(ctx)))))), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cemi' field"))
	}
	m.Cemi = cemi

	if closeErr := readBuffer.CloseContext("DeviceConfigurationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationRequest")
	}

	return m, nil
}

func (m *_DeviceConfigurationRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeviceConfigurationRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeviceConfigurationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationRequest")
		}

		if err := WriteSimpleField[DeviceConfigurationRequestDataBlock](ctx, "deviceConfigurationRequestDataBlock", m.GetDeviceConfigurationRequestDataBlock(), WriteComplex[DeviceConfigurationRequestDataBlock](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceConfigurationRequestDataBlock' field")
		}

		if err := WriteSimpleField[CEMI](ctx, "cemi", m.GetCemi(), WriteComplex[CEMI](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'cemi' field")
		}

		if popErr := writeBuffer.PopContext("DeviceConfigurationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeviceConfigurationRequest")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_DeviceConfigurationRequest) GetTotalLength() uint16 {
	return m.TotalLength
}

//
////

func (m *_DeviceConfigurationRequest) IsDeviceConfigurationRequest() {}

func (m *_DeviceConfigurationRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeviceConfigurationRequest) deepCopy() *_DeviceConfigurationRequest {
	if m == nil {
		return nil
	}
	_DeviceConfigurationRequestCopy := &_DeviceConfigurationRequest{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		utils.DeepCopy[DeviceConfigurationRequestDataBlock](m.DeviceConfigurationRequestDataBlock),
		utils.DeepCopy[CEMI](m.Cemi),
		m.TotalLength,
	}
	_DeviceConfigurationRequestCopy.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _DeviceConfigurationRequestCopy
}

func (m *_DeviceConfigurationRequest) String() string {
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

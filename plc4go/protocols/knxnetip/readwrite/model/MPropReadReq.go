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

// MPropReadReq is the corresponding interface of MPropReadReq
type MPropReadReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// GetInterfaceObjectType returns InterfaceObjectType (property field)
	GetInterfaceObjectType() uint16
	// GetObjectInstance returns ObjectInstance (property field)
	GetObjectInstance() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetNumberOfElements returns NumberOfElements (property field)
	GetNumberOfElements() uint8
	// GetStartIndex returns StartIndex (property field)
	GetStartIndex() uint16
	// IsMPropReadReq is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMPropReadReq()
	// CreateBuilder creates a MPropReadReqBuilder
	CreateMPropReadReqBuilder() MPropReadReqBuilder
}

// _MPropReadReq is the data-structure of this message
type _MPropReadReq struct {
	CEMIContract
	InterfaceObjectType uint16
	ObjectInstance      uint8
	PropertyId          uint8
	NumberOfElements    uint8
	StartIndex          uint16
}

var _ MPropReadReq = (*_MPropReadReq)(nil)
var _ CEMIRequirements = (*_MPropReadReq)(nil)

// NewMPropReadReq factory function for _MPropReadReq
func NewMPropReadReq(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16, size uint16) *_MPropReadReq {
	_result := &_MPropReadReq{
		CEMIContract:        NewCEMI(size),
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MPropReadReqBuilder is a builder for MPropReadReq
type MPropReadReqBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16) MPropReadReqBuilder
	// WithInterfaceObjectType adds InterfaceObjectType (property field)
	WithInterfaceObjectType(uint16) MPropReadReqBuilder
	// WithObjectInstance adds ObjectInstance (property field)
	WithObjectInstance(uint8) MPropReadReqBuilder
	// WithPropertyId adds PropertyId (property field)
	WithPropertyId(uint8) MPropReadReqBuilder
	// WithNumberOfElements adds NumberOfElements (property field)
	WithNumberOfElements(uint8) MPropReadReqBuilder
	// WithStartIndex adds StartIndex (property field)
	WithStartIndex(uint16) MPropReadReqBuilder
	// Build builds the MPropReadReq or returns an error if something is wrong
	Build() (MPropReadReq, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MPropReadReq
}

// NewMPropReadReqBuilder() creates a MPropReadReqBuilder
func NewMPropReadReqBuilder() MPropReadReqBuilder {
	return &_MPropReadReqBuilder{_MPropReadReq: new(_MPropReadReq)}
}

type _MPropReadReqBuilder struct {
	*_MPropReadReq

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (MPropReadReqBuilder) = (*_MPropReadReqBuilder)(nil)

func (b *_MPropReadReqBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
}

func (b *_MPropReadReqBuilder) WithMandatoryFields(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16) MPropReadReqBuilder {
	return b.WithInterfaceObjectType(interfaceObjectType).WithObjectInstance(objectInstance).WithPropertyId(propertyId).WithNumberOfElements(numberOfElements).WithStartIndex(startIndex)
}

func (b *_MPropReadReqBuilder) WithInterfaceObjectType(interfaceObjectType uint16) MPropReadReqBuilder {
	b.InterfaceObjectType = interfaceObjectType
	return b
}

func (b *_MPropReadReqBuilder) WithObjectInstance(objectInstance uint8) MPropReadReqBuilder {
	b.ObjectInstance = objectInstance
	return b
}

func (b *_MPropReadReqBuilder) WithPropertyId(propertyId uint8) MPropReadReqBuilder {
	b.PropertyId = propertyId
	return b
}

func (b *_MPropReadReqBuilder) WithNumberOfElements(numberOfElements uint8) MPropReadReqBuilder {
	b.NumberOfElements = numberOfElements
	return b
}

func (b *_MPropReadReqBuilder) WithStartIndex(startIndex uint16) MPropReadReqBuilder {
	b.StartIndex = startIndex
	return b
}

func (b *_MPropReadReqBuilder) Build() (MPropReadReq, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MPropReadReq.deepCopy(), nil
}

func (b *_MPropReadReqBuilder) MustBuild() MPropReadReq {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MPropReadReqBuilder) Done() CEMIBuilder {
	return b.parentBuilder
}

func (b *_MPropReadReqBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_MPropReadReqBuilder) DeepCopy() any {
	_copy := b.CreateMPropReadReqBuilder().(*_MPropReadReqBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMPropReadReqBuilder creates a MPropReadReqBuilder
func (b *_MPropReadReq) CreateMPropReadReqBuilder() MPropReadReqBuilder {
	if b == nil {
		return NewMPropReadReqBuilder()
	}
	return &_MPropReadReqBuilder{_MPropReadReq: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropReadReq) GetMessageCode() uint8 {
	return 0xFC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropReadReq) GetParent() CEMIContract {
	return m.CEMIContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MPropReadReq) GetInterfaceObjectType() uint16 {
	return m.InterfaceObjectType
}

func (m *_MPropReadReq) GetObjectInstance() uint8 {
	return m.ObjectInstance
}

func (m *_MPropReadReq) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_MPropReadReq) GetNumberOfElements() uint8 {
	return m.NumberOfElements
}

func (m *_MPropReadReq) GetStartIndex() uint16 {
	return m.StartIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMPropReadReq(structType any) MPropReadReq {
	if casted, ok := structType.(MPropReadReq); ok {
		return casted
	}
	if casted, ok := structType.(*MPropReadReq); ok {
		return *casted
	}
	return nil
}

func (m *_MPropReadReq) GetTypeName() string {
	return "MPropReadReq"
}

func (m *_MPropReadReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	// Simple field (interfaceObjectType)
	lengthInBits += 16

	// Simple field (objectInstance)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (numberOfElements)
	lengthInBits += 4

	// Simple field (startIndex)
	lengthInBits += 12

	return lengthInBits
}

func (m *_MPropReadReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MPropReadReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mPropReadReq MPropReadReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropReadReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropReadReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceObjectType, err := ReadSimpleField(ctx, "interfaceObjectType", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceObjectType' field"))
	}
	m.InterfaceObjectType = interfaceObjectType

	objectInstance, err := ReadSimpleField(ctx, "objectInstance", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectInstance' field"))
	}
	m.ObjectInstance = objectInstance

	propertyId, err := ReadSimpleField(ctx, "propertyId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyId' field"))
	}
	m.PropertyId = propertyId

	numberOfElements, err := ReadSimpleField(ctx, "numberOfElements", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfElements' field"))
	}
	m.NumberOfElements = numberOfElements

	startIndex, err := ReadSimpleField(ctx, "startIndex", ReadUnsignedShort(readBuffer, uint8(12)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startIndex' field"))
	}
	m.StartIndex = startIndex

	if closeErr := readBuffer.CloseContext("MPropReadReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropReadReq")
	}

	return m, nil
}

func (m *_MPropReadReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropReadReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropReadReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropReadReq")
		}

		if err := WriteSimpleField[uint16](ctx, "interfaceObjectType", m.GetInterfaceObjectType(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceObjectType' field")
		}

		if err := WriteSimpleField[uint8](ctx, "objectInstance", m.GetObjectInstance(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectInstance' field")
		}

		if err := WriteSimpleField[uint8](ctx, "propertyId", m.GetPropertyId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "numberOfElements", m.GetNumberOfElements(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfElements' field")
		}

		if err := WriteSimpleField[uint16](ctx, "startIndex", m.GetStartIndex(), WriteUnsignedShort(writeBuffer, 12)); err != nil {
			return errors.Wrap(err, "Error serializing 'startIndex' field")
		}

		if popErr := writeBuffer.PopContext("MPropReadReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropReadReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MPropReadReq) IsMPropReadReq() {}

func (m *_MPropReadReq) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MPropReadReq) deepCopy() *_MPropReadReq {
	if m == nil {
		return nil
	}
	_MPropReadReqCopy := &_MPropReadReq{
		m.CEMIContract.(*_CEMI).deepCopy(),
		m.InterfaceObjectType,
		m.ObjectInstance,
		m.PropertyId,
		m.NumberOfElements,
		m.StartIndex,
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _MPropReadReqCopy
}

func (m *_MPropReadReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

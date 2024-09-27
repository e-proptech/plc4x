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

// SzlId is the corresponding interface of SzlId
type SzlId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetTypeClass returns TypeClass (property field)
	GetTypeClass() SzlModuleTypeClass
	// GetSublistExtract returns SublistExtract (property field)
	GetSublistExtract() uint8
	// GetSublistList returns SublistList (property field)
	GetSublistList() SzlSublist
	// IsSzlId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSzlId()
	// CreateBuilder creates a SzlIdBuilder
	CreateSzlIdBuilder() SzlIdBuilder
}

// _SzlId is the data-structure of this message
type _SzlId struct {
	TypeClass      SzlModuleTypeClass
	SublistExtract uint8
	SublistList    SzlSublist
}

var _ SzlId = (*_SzlId)(nil)

// NewSzlId factory function for _SzlId
func NewSzlId(typeClass SzlModuleTypeClass, sublistExtract uint8, sublistList SzlSublist) *_SzlId {
	return &_SzlId{TypeClass: typeClass, SublistExtract: sublistExtract, SublistList: sublistList}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SzlIdBuilder is a builder for SzlId
type SzlIdBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(typeClass SzlModuleTypeClass, sublistExtract uint8, sublistList SzlSublist) SzlIdBuilder
	// WithTypeClass adds TypeClass (property field)
	WithTypeClass(SzlModuleTypeClass) SzlIdBuilder
	// WithSublistExtract adds SublistExtract (property field)
	WithSublistExtract(uint8) SzlIdBuilder
	// WithSublistList adds SublistList (property field)
	WithSublistList(SzlSublist) SzlIdBuilder
	// Build builds the SzlId or returns an error if something is wrong
	Build() (SzlId, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SzlId
}

// NewSzlIdBuilder() creates a SzlIdBuilder
func NewSzlIdBuilder() SzlIdBuilder {
	return &_SzlIdBuilder{_SzlId: new(_SzlId)}
}

type _SzlIdBuilder struct {
	*_SzlId

	err *utils.MultiError
}

var _ (SzlIdBuilder) = (*_SzlIdBuilder)(nil)

func (b *_SzlIdBuilder) WithMandatoryFields(typeClass SzlModuleTypeClass, sublistExtract uint8, sublistList SzlSublist) SzlIdBuilder {
	return b.WithTypeClass(typeClass).WithSublistExtract(sublistExtract).WithSublistList(sublistList)
}

func (b *_SzlIdBuilder) WithTypeClass(typeClass SzlModuleTypeClass) SzlIdBuilder {
	b.TypeClass = typeClass
	return b
}

func (b *_SzlIdBuilder) WithSublistExtract(sublistExtract uint8) SzlIdBuilder {
	b.SublistExtract = sublistExtract
	return b
}

func (b *_SzlIdBuilder) WithSublistList(sublistList SzlSublist) SzlIdBuilder {
	b.SublistList = sublistList
	return b
}

func (b *_SzlIdBuilder) Build() (SzlId, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SzlId.deepCopy(), nil
}

func (b *_SzlIdBuilder) MustBuild() SzlId {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SzlIdBuilder) DeepCopy() any {
	_copy := b.CreateSzlIdBuilder().(*_SzlIdBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSzlIdBuilder creates a SzlIdBuilder
func (b *_SzlId) CreateSzlIdBuilder() SzlIdBuilder {
	if b == nil {
		return NewSzlIdBuilder()
	}
	return &_SzlIdBuilder{_SzlId: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SzlId) GetTypeClass() SzlModuleTypeClass {
	return m.TypeClass
}

func (m *_SzlId) GetSublistExtract() uint8 {
	return m.SublistExtract
}

func (m *_SzlId) GetSublistList() SzlSublist {
	return m.SublistList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSzlId(structType any) SzlId {
	if casted, ok := structType.(SzlId); ok {
		return casted
	}
	if casted, ok := structType.(*SzlId); ok {
		return *casted
	}
	return nil
}

func (m *_SzlId) GetTypeName() string {
	return "SzlId"
}

func (m *_SzlId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (typeClass)
	lengthInBits += 4

	// Simple field (sublistExtract)
	lengthInBits += 4

	// Simple field (sublistList)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SzlId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SzlIdParse(ctx context.Context, theBytes []byte) (SzlId, error) {
	return SzlIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SzlIdParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SzlId, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SzlId, error) {
		return SzlIdParseWithBuffer(ctx, readBuffer)
	}
}

func SzlIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SzlId, error) {
	v, err := (&_SzlId{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_SzlId) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__szlId SzlId, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SzlId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SzlId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	typeClass, err := ReadEnumField[SzlModuleTypeClass](ctx, "typeClass", "SzlModuleTypeClass", ReadEnum(SzlModuleTypeClassByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeClass' field"))
	}
	m.TypeClass = typeClass

	sublistExtract, err := ReadSimpleField(ctx, "sublistExtract", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sublistExtract' field"))
	}
	m.SublistExtract = sublistExtract

	sublistList, err := ReadEnumField[SzlSublist](ctx, "sublistList", "SzlSublist", ReadEnum(SzlSublistByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sublistList' field"))
	}
	m.SublistList = sublistList

	if closeErr := readBuffer.CloseContext("SzlId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SzlId")
	}

	return m, nil
}

func (m *_SzlId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SzlId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SzlId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SzlId")
	}

	if err := WriteSimpleEnumField[SzlModuleTypeClass](ctx, "typeClass", "SzlModuleTypeClass", m.GetTypeClass(), WriteEnum[SzlModuleTypeClass, uint8](SzlModuleTypeClass.GetValue, SzlModuleTypeClass.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 4))); err != nil {
		return errors.Wrap(err, "Error serializing 'typeClass' field")
	}

	if err := WriteSimpleField[uint8](ctx, "sublistExtract", m.GetSublistExtract(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
		return errors.Wrap(err, "Error serializing 'sublistExtract' field")
	}

	if err := WriteSimpleEnumField[SzlSublist](ctx, "sublistList", "SzlSublist", m.GetSublistList(), WriteEnum[SzlSublist, uint8](SzlSublist.GetValue, SzlSublist.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'sublistList' field")
	}

	if popErr := writeBuffer.PopContext("SzlId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SzlId")
	}
	return nil
}

func (m *_SzlId) IsSzlId() {}

func (m *_SzlId) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SzlId) deepCopy() *_SzlId {
	if m == nil {
		return nil
	}
	_SzlIdCopy := &_SzlId{
		m.TypeClass,
		m.SublistExtract,
		m.SublistList,
	}
	return _SzlIdCopy
}

func (m *_SzlId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

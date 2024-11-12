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

// SimpleTypeDescription is the corresponding interface of SimpleTypeDescription
type SimpleTypeDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetDataTypeId returns DataTypeId (property field)
	GetDataTypeId() NodeId
	// GetName returns Name (property field)
	GetName() QualifiedName
	// GetBaseDataType returns BaseDataType (property field)
	GetBaseDataType() NodeId
	// GetBuiltInType returns BuiltInType (property field)
	GetBuiltInType() uint8
	// IsSimpleTypeDescription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSimpleTypeDescription()
	// CreateBuilder creates a SimpleTypeDescriptionBuilder
	CreateSimpleTypeDescriptionBuilder() SimpleTypeDescriptionBuilder
}

// _SimpleTypeDescription is the data-structure of this message
type _SimpleTypeDescription struct {
	ExtensionObjectDefinitionContract
	DataTypeId   NodeId
	Name         QualifiedName
	BaseDataType NodeId
	BuiltInType  uint8
}

var _ SimpleTypeDescription = (*_SimpleTypeDescription)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SimpleTypeDescription)(nil)

// NewSimpleTypeDescription factory function for _SimpleTypeDescription
func NewSimpleTypeDescription(dataTypeId NodeId, name QualifiedName, baseDataType NodeId, builtInType uint8) *_SimpleTypeDescription {
	if dataTypeId == nil {
		panic("dataTypeId of type NodeId for SimpleTypeDescription must not be nil")
	}
	if name == nil {
		panic("name of type QualifiedName for SimpleTypeDescription must not be nil")
	}
	if baseDataType == nil {
		panic("baseDataType of type NodeId for SimpleTypeDescription must not be nil")
	}
	_result := &_SimpleTypeDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		DataTypeId:                        dataTypeId,
		Name:                              name,
		BaseDataType:                      baseDataType,
		BuiltInType:                       builtInType,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SimpleTypeDescriptionBuilder is a builder for SimpleTypeDescription
type SimpleTypeDescriptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dataTypeId NodeId, name QualifiedName, baseDataType NodeId, builtInType uint8) SimpleTypeDescriptionBuilder
	// WithDataTypeId adds DataTypeId (property field)
	WithDataTypeId(NodeId) SimpleTypeDescriptionBuilder
	// WithDataTypeIdBuilder adds DataTypeId (property field) which is build by the builder
	WithDataTypeIdBuilder(func(NodeIdBuilder) NodeIdBuilder) SimpleTypeDescriptionBuilder
	// WithName adds Name (property field)
	WithName(QualifiedName) SimpleTypeDescriptionBuilder
	// WithNameBuilder adds Name (property field) which is build by the builder
	WithNameBuilder(func(QualifiedNameBuilder) QualifiedNameBuilder) SimpleTypeDescriptionBuilder
	// WithBaseDataType adds BaseDataType (property field)
	WithBaseDataType(NodeId) SimpleTypeDescriptionBuilder
	// WithBaseDataTypeBuilder adds BaseDataType (property field) which is build by the builder
	WithBaseDataTypeBuilder(func(NodeIdBuilder) NodeIdBuilder) SimpleTypeDescriptionBuilder
	// WithBuiltInType adds BuiltInType (property field)
	WithBuiltInType(uint8) SimpleTypeDescriptionBuilder
	// Build builds the SimpleTypeDescription or returns an error if something is wrong
	Build() (SimpleTypeDescription, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SimpleTypeDescription
}

// NewSimpleTypeDescriptionBuilder() creates a SimpleTypeDescriptionBuilder
func NewSimpleTypeDescriptionBuilder() SimpleTypeDescriptionBuilder {
	return &_SimpleTypeDescriptionBuilder{_SimpleTypeDescription: new(_SimpleTypeDescription)}
}

type _SimpleTypeDescriptionBuilder struct {
	*_SimpleTypeDescription

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (SimpleTypeDescriptionBuilder) = (*_SimpleTypeDescriptionBuilder)(nil)

func (b *_SimpleTypeDescriptionBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_SimpleTypeDescriptionBuilder) WithMandatoryFields(dataTypeId NodeId, name QualifiedName, baseDataType NodeId, builtInType uint8) SimpleTypeDescriptionBuilder {
	return b.WithDataTypeId(dataTypeId).WithName(name).WithBaseDataType(baseDataType).WithBuiltInType(builtInType)
}

func (b *_SimpleTypeDescriptionBuilder) WithDataTypeId(dataTypeId NodeId) SimpleTypeDescriptionBuilder {
	b.DataTypeId = dataTypeId
	return b
}

func (b *_SimpleTypeDescriptionBuilder) WithDataTypeIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) SimpleTypeDescriptionBuilder {
	builder := builderSupplier(b.DataTypeId.CreateNodeIdBuilder())
	var err error
	b.DataTypeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_SimpleTypeDescriptionBuilder) WithName(name QualifiedName) SimpleTypeDescriptionBuilder {
	b.Name = name
	return b
}

func (b *_SimpleTypeDescriptionBuilder) WithNameBuilder(builderSupplier func(QualifiedNameBuilder) QualifiedNameBuilder) SimpleTypeDescriptionBuilder {
	builder := builderSupplier(b.Name.CreateQualifiedNameBuilder())
	var err error
	b.Name, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "QualifiedNameBuilder failed"))
	}
	return b
}

func (b *_SimpleTypeDescriptionBuilder) WithBaseDataType(baseDataType NodeId) SimpleTypeDescriptionBuilder {
	b.BaseDataType = baseDataType
	return b
}

func (b *_SimpleTypeDescriptionBuilder) WithBaseDataTypeBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) SimpleTypeDescriptionBuilder {
	builder := builderSupplier(b.BaseDataType.CreateNodeIdBuilder())
	var err error
	b.BaseDataType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_SimpleTypeDescriptionBuilder) WithBuiltInType(builtInType uint8) SimpleTypeDescriptionBuilder {
	b.BuiltInType = builtInType
	return b
}

func (b *_SimpleTypeDescriptionBuilder) Build() (SimpleTypeDescription, error) {
	if b.DataTypeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dataTypeId' not set"))
	}
	if b.Name == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'name' not set"))
	}
	if b.BaseDataType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'baseDataType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SimpleTypeDescription.deepCopy(), nil
}

func (b *_SimpleTypeDescriptionBuilder) MustBuild() SimpleTypeDescription {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SimpleTypeDescriptionBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_SimpleTypeDescriptionBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_SimpleTypeDescriptionBuilder) DeepCopy() any {
	_copy := b.CreateSimpleTypeDescriptionBuilder().(*_SimpleTypeDescriptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSimpleTypeDescriptionBuilder creates a SimpleTypeDescriptionBuilder
func (b *_SimpleTypeDescription) CreateSimpleTypeDescriptionBuilder() SimpleTypeDescriptionBuilder {
	if b == nil {
		return NewSimpleTypeDescriptionBuilder()
	}
	return &_SimpleTypeDescriptionBuilder{_SimpleTypeDescription: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SimpleTypeDescription) GetExtensionId() int32 {
	return int32(15007)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SimpleTypeDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SimpleTypeDescription) GetDataTypeId() NodeId {
	return m.DataTypeId
}

func (m *_SimpleTypeDescription) GetName() QualifiedName {
	return m.Name
}

func (m *_SimpleTypeDescription) GetBaseDataType() NodeId {
	return m.BaseDataType
}

func (m *_SimpleTypeDescription) GetBuiltInType() uint8 {
	return m.BuiltInType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSimpleTypeDescription(structType any) SimpleTypeDescription {
	if casted, ok := structType.(SimpleTypeDescription); ok {
		return casted
	}
	if casted, ok := structType.(*SimpleTypeDescription); ok {
		return *casted
	}
	return nil
}

func (m *_SimpleTypeDescription) GetTypeName() string {
	return "SimpleTypeDescription"
}

func (m *_SimpleTypeDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (dataTypeId)
	lengthInBits += m.DataTypeId.GetLengthInBits(ctx)

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Simple field (baseDataType)
	lengthInBits += m.BaseDataType.GetLengthInBits(ctx)

	// Simple field (builtInType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SimpleTypeDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SimpleTypeDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__simpleTypeDescription SimpleTypeDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SimpleTypeDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SimpleTypeDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataTypeId, err := ReadSimpleField[NodeId](ctx, "dataTypeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeId' field"))
	}
	m.DataTypeId = dataTypeId

	name, err := ReadSimpleField[QualifiedName](ctx, "name", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	baseDataType, err := ReadSimpleField[NodeId](ctx, "baseDataType", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'baseDataType' field"))
	}
	m.BaseDataType = baseDataType

	builtInType, err := ReadSimpleField(ctx, "builtInType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'builtInType' field"))
	}
	m.BuiltInType = builtInType

	if closeErr := readBuffer.CloseContext("SimpleTypeDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SimpleTypeDescription")
	}

	return m, nil
}

func (m *_SimpleTypeDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SimpleTypeDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SimpleTypeDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SimpleTypeDescription")
		}

		if err := WriteSimpleField[NodeId](ctx, "dataTypeId", m.GetDataTypeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataTypeId' field")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "name", m.GetName(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "baseDataType", m.GetBaseDataType(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'baseDataType' field")
		}

		if err := WriteSimpleField[uint8](ctx, "builtInType", m.GetBuiltInType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'builtInType' field")
		}

		if popErr := writeBuffer.PopContext("SimpleTypeDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SimpleTypeDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SimpleTypeDescription) IsSimpleTypeDescription() {}

func (m *_SimpleTypeDescription) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SimpleTypeDescription) deepCopy() *_SimpleTypeDescription {
	if m == nil {
		return nil
	}
	_SimpleTypeDescriptionCopy := &_SimpleTypeDescription{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[NodeId](m.DataTypeId),
		utils.DeepCopy[QualifiedName](m.Name),
		utils.DeepCopy[NodeId](m.BaseDataType),
		m.BuiltInType,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _SimpleTypeDescriptionCopy
}

func (m *_SimpleTypeDescription) String() string {
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

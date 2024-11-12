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

// SemanticChangeStructureDataType is the corresponding interface of SemanticChangeStructureDataType
type SemanticChangeStructureDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetAffected returns Affected (property field)
	GetAffected() NodeId
	// GetAffectedType returns AffectedType (property field)
	GetAffectedType() NodeId
	// IsSemanticChangeStructureDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSemanticChangeStructureDataType()
	// CreateBuilder creates a SemanticChangeStructureDataTypeBuilder
	CreateSemanticChangeStructureDataTypeBuilder() SemanticChangeStructureDataTypeBuilder
}

// _SemanticChangeStructureDataType is the data-structure of this message
type _SemanticChangeStructureDataType struct {
	ExtensionObjectDefinitionContract
	Affected     NodeId
	AffectedType NodeId
}

var _ SemanticChangeStructureDataType = (*_SemanticChangeStructureDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SemanticChangeStructureDataType)(nil)

// NewSemanticChangeStructureDataType factory function for _SemanticChangeStructureDataType
func NewSemanticChangeStructureDataType(affected NodeId, affectedType NodeId) *_SemanticChangeStructureDataType {
	if affected == nil {
		panic("affected of type NodeId for SemanticChangeStructureDataType must not be nil")
	}
	if affectedType == nil {
		panic("affectedType of type NodeId for SemanticChangeStructureDataType must not be nil")
	}
	_result := &_SemanticChangeStructureDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Affected:                          affected,
		AffectedType:                      affectedType,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SemanticChangeStructureDataTypeBuilder is a builder for SemanticChangeStructureDataType
type SemanticChangeStructureDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(affected NodeId, affectedType NodeId) SemanticChangeStructureDataTypeBuilder
	// WithAffected adds Affected (property field)
	WithAffected(NodeId) SemanticChangeStructureDataTypeBuilder
	// WithAffectedBuilder adds Affected (property field) which is build by the builder
	WithAffectedBuilder(func(NodeIdBuilder) NodeIdBuilder) SemanticChangeStructureDataTypeBuilder
	// WithAffectedType adds AffectedType (property field)
	WithAffectedType(NodeId) SemanticChangeStructureDataTypeBuilder
	// WithAffectedTypeBuilder adds AffectedType (property field) which is build by the builder
	WithAffectedTypeBuilder(func(NodeIdBuilder) NodeIdBuilder) SemanticChangeStructureDataTypeBuilder
	// Build builds the SemanticChangeStructureDataType or returns an error if something is wrong
	Build() (SemanticChangeStructureDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SemanticChangeStructureDataType
}

// NewSemanticChangeStructureDataTypeBuilder() creates a SemanticChangeStructureDataTypeBuilder
func NewSemanticChangeStructureDataTypeBuilder() SemanticChangeStructureDataTypeBuilder {
	return &_SemanticChangeStructureDataTypeBuilder{_SemanticChangeStructureDataType: new(_SemanticChangeStructureDataType)}
}

type _SemanticChangeStructureDataTypeBuilder struct {
	*_SemanticChangeStructureDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (SemanticChangeStructureDataTypeBuilder) = (*_SemanticChangeStructureDataTypeBuilder)(nil)

func (b *_SemanticChangeStructureDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_SemanticChangeStructureDataTypeBuilder) WithMandatoryFields(affected NodeId, affectedType NodeId) SemanticChangeStructureDataTypeBuilder {
	return b.WithAffected(affected).WithAffectedType(affectedType)
}

func (b *_SemanticChangeStructureDataTypeBuilder) WithAffected(affected NodeId) SemanticChangeStructureDataTypeBuilder {
	b.Affected = affected
	return b
}

func (b *_SemanticChangeStructureDataTypeBuilder) WithAffectedBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) SemanticChangeStructureDataTypeBuilder {
	builder := builderSupplier(b.Affected.CreateNodeIdBuilder())
	var err error
	b.Affected, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_SemanticChangeStructureDataTypeBuilder) WithAffectedType(affectedType NodeId) SemanticChangeStructureDataTypeBuilder {
	b.AffectedType = affectedType
	return b
}

func (b *_SemanticChangeStructureDataTypeBuilder) WithAffectedTypeBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) SemanticChangeStructureDataTypeBuilder {
	builder := builderSupplier(b.AffectedType.CreateNodeIdBuilder())
	var err error
	b.AffectedType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_SemanticChangeStructureDataTypeBuilder) Build() (SemanticChangeStructureDataType, error) {
	if b.Affected == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'affected' not set"))
	}
	if b.AffectedType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'affectedType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SemanticChangeStructureDataType.deepCopy(), nil
}

func (b *_SemanticChangeStructureDataTypeBuilder) MustBuild() SemanticChangeStructureDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SemanticChangeStructureDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_SemanticChangeStructureDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_SemanticChangeStructureDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateSemanticChangeStructureDataTypeBuilder().(*_SemanticChangeStructureDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSemanticChangeStructureDataTypeBuilder creates a SemanticChangeStructureDataTypeBuilder
func (b *_SemanticChangeStructureDataType) CreateSemanticChangeStructureDataTypeBuilder() SemanticChangeStructureDataTypeBuilder {
	if b == nil {
		return NewSemanticChangeStructureDataTypeBuilder()
	}
	return &_SemanticChangeStructureDataTypeBuilder{_SemanticChangeStructureDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SemanticChangeStructureDataType) GetExtensionId() int32 {
	return int32(899)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SemanticChangeStructureDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SemanticChangeStructureDataType) GetAffected() NodeId {
	return m.Affected
}

func (m *_SemanticChangeStructureDataType) GetAffectedType() NodeId {
	return m.AffectedType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSemanticChangeStructureDataType(structType any) SemanticChangeStructureDataType {
	if casted, ok := structType.(SemanticChangeStructureDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SemanticChangeStructureDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SemanticChangeStructureDataType) GetTypeName() string {
	return "SemanticChangeStructureDataType"
}

func (m *_SemanticChangeStructureDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (affected)
	lengthInBits += m.Affected.GetLengthInBits(ctx)

	// Simple field (affectedType)
	lengthInBits += m.AffectedType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SemanticChangeStructureDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SemanticChangeStructureDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__semanticChangeStructureDataType SemanticChangeStructureDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SemanticChangeStructureDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SemanticChangeStructureDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	affected, err := ReadSimpleField[NodeId](ctx, "affected", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'affected' field"))
	}
	m.Affected = affected

	affectedType, err := ReadSimpleField[NodeId](ctx, "affectedType", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'affectedType' field"))
	}
	m.AffectedType = affectedType

	if closeErr := readBuffer.CloseContext("SemanticChangeStructureDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SemanticChangeStructureDataType")
	}

	return m, nil
}

func (m *_SemanticChangeStructureDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SemanticChangeStructureDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SemanticChangeStructureDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SemanticChangeStructureDataType")
		}

		if err := WriteSimpleField[NodeId](ctx, "affected", m.GetAffected(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'affected' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "affectedType", m.GetAffectedType(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'affectedType' field")
		}

		if popErr := writeBuffer.PopContext("SemanticChangeStructureDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SemanticChangeStructureDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SemanticChangeStructureDataType) IsSemanticChangeStructureDataType() {}

func (m *_SemanticChangeStructureDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SemanticChangeStructureDataType) deepCopy() *_SemanticChangeStructureDataType {
	if m == nil {
		return nil
	}
	_SemanticChangeStructureDataTypeCopy := &_SemanticChangeStructureDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[NodeId](m.Affected),
		utils.DeepCopy[NodeId](m.AffectedType),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _SemanticChangeStructureDataTypeCopy
}

func (m *_SemanticChangeStructureDataType) String() string {
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

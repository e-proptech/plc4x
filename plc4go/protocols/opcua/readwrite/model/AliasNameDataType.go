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

// AliasNameDataType is the corresponding interface of AliasNameDataType
type AliasNameDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetAliasName returns AliasName (property field)
	GetAliasName() QualifiedName
	// GetReferencedNodes returns ReferencedNodes (property field)
	GetReferencedNodes() []ExpandedNodeId
	// IsAliasNameDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAliasNameDataType()
	// CreateBuilder creates a AliasNameDataTypeBuilder
	CreateAliasNameDataTypeBuilder() AliasNameDataTypeBuilder
}

// _AliasNameDataType is the data-structure of this message
type _AliasNameDataType struct {
	ExtensionObjectDefinitionContract
	AliasName       QualifiedName
	ReferencedNodes []ExpandedNodeId
}

var _ AliasNameDataType = (*_AliasNameDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AliasNameDataType)(nil)

// NewAliasNameDataType factory function for _AliasNameDataType
func NewAliasNameDataType(aliasName QualifiedName, referencedNodes []ExpandedNodeId) *_AliasNameDataType {
	if aliasName == nil {
		panic("aliasName of type QualifiedName for AliasNameDataType must not be nil")
	}
	_result := &_AliasNameDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		AliasName:                         aliasName,
		ReferencedNodes:                   referencedNodes,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AliasNameDataTypeBuilder is a builder for AliasNameDataType
type AliasNameDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(aliasName QualifiedName, referencedNodes []ExpandedNodeId) AliasNameDataTypeBuilder
	// WithAliasName adds AliasName (property field)
	WithAliasName(QualifiedName) AliasNameDataTypeBuilder
	// WithAliasNameBuilder adds AliasName (property field) which is build by the builder
	WithAliasNameBuilder(func(QualifiedNameBuilder) QualifiedNameBuilder) AliasNameDataTypeBuilder
	// WithReferencedNodes adds ReferencedNodes (property field)
	WithReferencedNodes(...ExpandedNodeId) AliasNameDataTypeBuilder
	// Build builds the AliasNameDataType or returns an error if something is wrong
	Build() (AliasNameDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AliasNameDataType
}

// NewAliasNameDataTypeBuilder() creates a AliasNameDataTypeBuilder
func NewAliasNameDataTypeBuilder() AliasNameDataTypeBuilder {
	return &_AliasNameDataTypeBuilder{_AliasNameDataType: new(_AliasNameDataType)}
}

type _AliasNameDataTypeBuilder struct {
	*_AliasNameDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (AliasNameDataTypeBuilder) = (*_AliasNameDataTypeBuilder)(nil)

func (b *_AliasNameDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_AliasNameDataTypeBuilder) WithMandatoryFields(aliasName QualifiedName, referencedNodes []ExpandedNodeId) AliasNameDataTypeBuilder {
	return b.WithAliasName(aliasName).WithReferencedNodes(referencedNodes...)
}

func (b *_AliasNameDataTypeBuilder) WithAliasName(aliasName QualifiedName) AliasNameDataTypeBuilder {
	b.AliasName = aliasName
	return b
}

func (b *_AliasNameDataTypeBuilder) WithAliasNameBuilder(builderSupplier func(QualifiedNameBuilder) QualifiedNameBuilder) AliasNameDataTypeBuilder {
	builder := builderSupplier(b.AliasName.CreateQualifiedNameBuilder())
	var err error
	b.AliasName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "QualifiedNameBuilder failed"))
	}
	return b
}

func (b *_AliasNameDataTypeBuilder) WithReferencedNodes(referencedNodes ...ExpandedNodeId) AliasNameDataTypeBuilder {
	b.ReferencedNodes = referencedNodes
	return b
}

func (b *_AliasNameDataTypeBuilder) Build() (AliasNameDataType, error) {
	if b.AliasName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'aliasName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AliasNameDataType.deepCopy(), nil
}

func (b *_AliasNameDataTypeBuilder) MustBuild() AliasNameDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AliasNameDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_AliasNameDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_AliasNameDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateAliasNameDataTypeBuilder().(*_AliasNameDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAliasNameDataTypeBuilder creates a AliasNameDataTypeBuilder
func (b *_AliasNameDataType) CreateAliasNameDataTypeBuilder() AliasNameDataTypeBuilder {
	if b == nil {
		return NewAliasNameDataTypeBuilder()
	}
	return &_AliasNameDataTypeBuilder{_AliasNameDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AliasNameDataType) GetExtensionId() int32 {
	return int32(23470)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AliasNameDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AliasNameDataType) GetAliasName() QualifiedName {
	return m.AliasName
}

func (m *_AliasNameDataType) GetReferencedNodes() []ExpandedNodeId {
	return m.ReferencedNodes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAliasNameDataType(structType any) AliasNameDataType {
	if casted, ok := structType.(AliasNameDataType); ok {
		return casted
	}
	if casted, ok := structType.(*AliasNameDataType); ok {
		return *casted
	}
	return nil
}

func (m *_AliasNameDataType) GetTypeName() string {
	return "AliasNameDataType"
}

func (m *_AliasNameDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (aliasName)
	lengthInBits += m.AliasName.GetLengthInBits(ctx)

	// Implicit Field (noOfReferencedNodes)
	lengthInBits += 32

	// Array field
	if len(m.ReferencedNodes) > 0 {
		for _curItem, element := range m.ReferencedNodes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReferencedNodes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AliasNameDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AliasNameDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__aliasNameDataType AliasNameDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AliasNameDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AliasNameDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	aliasName, err := ReadSimpleField[QualifiedName](ctx, "aliasName", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'aliasName' field"))
	}
	m.AliasName = aliasName

	noOfReferencedNodes, err := ReadImplicitField[int32](ctx, "noOfReferencedNodes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfReferencedNodes' field"))
	}
	_ = noOfReferencedNodes

	referencedNodes, err := ReadCountArrayField[ExpandedNodeId](ctx, "referencedNodes", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer), uint64(noOfReferencedNodes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencedNodes' field"))
	}
	m.ReferencedNodes = referencedNodes

	if closeErr := readBuffer.CloseContext("AliasNameDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AliasNameDataType")
	}

	return m, nil
}

func (m *_AliasNameDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AliasNameDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AliasNameDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AliasNameDataType")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "aliasName", m.GetAliasName(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'aliasName' field")
		}
		noOfReferencedNodes := int32(utils.InlineIf(bool((m.GetReferencedNodes()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetReferencedNodes()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfReferencedNodes", noOfReferencedNodes, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfReferencedNodes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "referencedNodes", m.GetReferencedNodes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'referencedNodes' field")
		}

		if popErr := writeBuffer.PopContext("AliasNameDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AliasNameDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AliasNameDataType) IsAliasNameDataType() {}

func (m *_AliasNameDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AliasNameDataType) deepCopy() *_AliasNameDataType {
	if m == nil {
		return nil
	}
	_AliasNameDataTypeCopy := &_AliasNameDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[QualifiedName](m.AliasName),
		utils.DeepCopySlice[ExpandedNodeId, ExpandedNodeId](m.ReferencedNodes),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _AliasNameDataTypeCopy
}

func (m *_AliasNameDataType) String() string {
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

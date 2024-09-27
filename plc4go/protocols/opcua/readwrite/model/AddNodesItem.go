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

// AddNodesItem is the corresponding interface of AddNodesItem
type AddNodesItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetParentNodeId returns ParentNodeId (property field)
	GetParentNodeId() ExpandedNodeId
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetRequestedNewNodeId returns RequestedNewNodeId (property field)
	GetRequestedNewNodeId() ExpandedNodeId
	// GetBrowseName returns BrowseName (property field)
	GetBrowseName() QualifiedName
	// GetNodeClass returns NodeClass (property field)
	GetNodeClass() NodeClass
	// GetNodeAttributes returns NodeAttributes (property field)
	GetNodeAttributes() ExtensionObject
	// GetTypeDefinition returns TypeDefinition (property field)
	GetTypeDefinition() ExpandedNodeId
	// IsAddNodesItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAddNodesItem()
	// CreateBuilder creates a AddNodesItemBuilder
	CreateAddNodesItemBuilder() AddNodesItemBuilder
}

// _AddNodesItem is the data-structure of this message
type _AddNodesItem struct {
	ExtensionObjectDefinitionContract
	ParentNodeId       ExpandedNodeId
	ReferenceTypeId    NodeId
	RequestedNewNodeId ExpandedNodeId
	BrowseName         QualifiedName
	NodeClass          NodeClass
	NodeAttributes     ExtensionObject
	TypeDefinition     ExpandedNodeId
}

var _ AddNodesItem = (*_AddNodesItem)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AddNodesItem)(nil)

// NewAddNodesItem factory function for _AddNodesItem
func NewAddNodesItem(parentNodeId ExpandedNodeId, referenceTypeId NodeId, requestedNewNodeId ExpandedNodeId, browseName QualifiedName, nodeClass NodeClass, nodeAttributes ExtensionObject, typeDefinition ExpandedNodeId) *_AddNodesItem {
	if parentNodeId == nil {
		panic("parentNodeId of type ExpandedNodeId for AddNodesItem must not be nil")
	}
	if referenceTypeId == nil {
		panic("referenceTypeId of type NodeId for AddNodesItem must not be nil")
	}
	if requestedNewNodeId == nil {
		panic("requestedNewNodeId of type ExpandedNodeId for AddNodesItem must not be nil")
	}
	if browseName == nil {
		panic("browseName of type QualifiedName for AddNodesItem must not be nil")
	}
	if nodeAttributes == nil {
		panic("nodeAttributes of type ExtensionObject for AddNodesItem must not be nil")
	}
	if typeDefinition == nil {
		panic("typeDefinition of type ExpandedNodeId for AddNodesItem must not be nil")
	}
	_result := &_AddNodesItem{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ParentNodeId:                      parentNodeId,
		ReferenceTypeId:                   referenceTypeId,
		RequestedNewNodeId:                requestedNewNodeId,
		BrowseName:                        browseName,
		NodeClass:                         nodeClass,
		NodeAttributes:                    nodeAttributes,
		TypeDefinition:                    typeDefinition,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AddNodesItemBuilder is a builder for AddNodesItem
type AddNodesItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(parentNodeId ExpandedNodeId, referenceTypeId NodeId, requestedNewNodeId ExpandedNodeId, browseName QualifiedName, nodeClass NodeClass, nodeAttributes ExtensionObject, typeDefinition ExpandedNodeId) AddNodesItemBuilder
	// WithParentNodeId adds ParentNodeId (property field)
	WithParentNodeId(ExpandedNodeId) AddNodesItemBuilder
	// WithParentNodeIdBuilder adds ParentNodeId (property field) which is build by the builder
	WithParentNodeIdBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) AddNodesItemBuilder
	// WithReferenceTypeId adds ReferenceTypeId (property field)
	WithReferenceTypeId(NodeId) AddNodesItemBuilder
	// WithReferenceTypeIdBuilder adds ReferenceTypeId (property field) which is build by the builder
	WithReferenceTypeIdBuilder(func(NodeIdBuilder) NodeIdBuilder) AddNodesItemBuilder
	// WithRequestedNewNodeId adds RequestedNewNodeId (property field)
	WithRequestedNewNodeId(ExpandedNodeId) AddNodesItemBuilder
	// WithRequestedNewNodeIdBuilder adds RequestedNewNodeId (property field) which is build by the builder
	WithRequestedNewNodeIdBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) AddNodesItemBuilder
	// WithBrowseName adds BrowseName (property field)
	WithBrowseName(QualifiedName) AddNodesItemBuilder
	// WithBrowseNameBuilder adds BrowseName (property field) which is build by the builder
	WithBrowseNameBuilder(func(QualifiedNameBuilder) QualifiedNameBuilder) AddNodesItemBuilder
	// WithNodeClass adds NodeClass (property field)
	WithNodeClass(NodeClass) AddNodesItemBuilder
	// WithNodeAttributes adds NodeAttributes (property field)
	WithNodeAttributes(ExtensionObject) AddNodesItemBuilder
	// WithNodeAttributesBuilder adds NodeAttributes (property field) which is build by the builder
	WithNodeAttributesBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) AddNodesItemBuilder
	// WithTypeDefinition adds TypeDefinition (property field)
	WithTypeDefinition(ExpandedNodeId) AddNodesItemBuilder
	// WithTypeDefinitionBuilder adds TypeDefinition (property field) which is build by the builder
	WithTypeDefinitionBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) AddNodesItemBuilder
	// Build builds the AddNodesItem or returns an error if something is wrong
	Build() (AddNodesItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AddNodesItem
}

// NewAddNodesItemBuilder() creates a AddNodesItemBuilder
func NewAddNodesItemBuilder() AddNodesItemBuilder {
	return &_AddNodesItemBuilder{_AddNodesItem: new(_AddNodesItem)}
}

type _AddNodesItemBuilder struct {
	*_AddNodesItem

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (AddNodesItemBuilder) = (*_AddNodesItemBuilder)(nil)

func (b *_AddNodesItemBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_AddNodesItemBuilder) WithMandatoryFields(parentNodeId ExpandedNodeId, referenceTypeId NodeId, requestedNewNodeId ExpandedNodeId, browseName QualifiedName, nodeClass NodeClass, nodeAttributes ExtensionObject, typeDefinition ExpandedNodeId) AddNodesItemBuilder {
	return b.WithParentNodeId(parentNodeId).WithReferenceTypeId(referenceTypeId).WithRequestedNewNodeId(requestedNewNodeId).WithBrowseName(browseName).WithNodeClass(nodeClass).WithNodeAttributes(nodeAttributes).WithTypeDefinition(typeDefinition)
}

func (b *_AddNodesItemBuilder) WithParentNodeId(parentNodeId ExpandedNodeId) AddNodesItemBuilder {
	b.ParentNodeId = parentNodeId
	return b
}

func (b *_AddNodesItemBuilder) WithParentNodeIdBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) AddNodesItemBuilder {
	builder := builderSupplier(b.ParentNodeId.CreateExpandedNodeIdBuilder())
	var err error
	b.ParentNodeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return b
}

func (b *_AddNodesItemBuilder) WithReferenceTypeId(referenceTypeId NodeId) AddNodesItemBuilder {
	b.ReferenceTypeId = referenceTypeId
	return b
}

func (b *_AddNodesItemBuilder) WithReferenceTypeIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) AddNodesItemBuilder {
	builder := builderSupplier(b.ReferenceTypeId.CreateNodeIdBuilder())
	var err error
	b.ReferenceTypeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_AddNodesItemBuilder) WithRequestedNewNodeId(requestedNewNodeId ExpandedNodeId) AddNodesItemBuilder {
	b.RequestedNewNodeId = requestedNewNodeId
	return b
}

func (b *_AddNodesItemBuilder) WithRequestedNewNodeIdBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) AddNodesItemBuilder {
	builder := builderSupplier(b.RequestedNewNodeId.CreateExpandedNodeIdBuilder())
	var err error
	b.RequestedNewNodeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return b
}

func (b *_AddNodesItemBuilder) WithBrowseName(browseName QualifiedName) AddNodesItemBuilder {
	b.BrowseName = browseName
	return b
}

func (b *_AddNodesItemBuilder) WithBrowseNameBuilder(builderSupplier func(QualifiedNameBuilder) QualifiedNameBuilder) AddNodesItemBuilder {
	builder := builderSupplier(b.BrowseName.CreateQualifiedNameBuilder())
	var err error
	b.BrowseName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "QualifiedNameBuilder failed"))
	}
	return b
}

func (b *_AddNodesItemBuilder) WithNodeClass(nodeClass NodeClass) AddNodesItemBuilder {
	b.NodeClass = nodeClass
	return b
}

func (b *_AddNodesItemBuilder) WithNodeAttributes(nodeAttributes ExtensionObject) AddNodesItemBuilder {
	b.NodeAttributes = nodeAttributes
	return b
}

func (b *_AddNodesItemBuilder) WithNodeAttributesBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) AddNodesItemBuilder {
	builder := builderSupplier(b.NodeAttributes.CreateExtensionObjectBuilder())
	var err error
	b.NodeAttributes, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_AddNodesItemBuilder) WithTypeDefinition(typeDefinition ExpandedNodeId) AddNodesItemBuilder {
	b.TypeDefinition = typeDefinition
	return b
}

func (b *_AddNodesItemBuilder) WithTypeDefinitionBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) AddNodesItemBuilder {
	builder := builderSupplier(b.TypeDefinition.CreateExpandedNodeIdBuilder())
	var err error
	b.TypeDefinition, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return b
}

func (b *_AddNodesItemBuilder) Build() (AddNodesItem, error) {
	if b.ParentNodeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'parentNodeId' not set"))
	}
	if b.ReferenceTypeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'referenceTypeId' not set"))
	}
	if b.RequestedNewNodeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestedNewNodeId' not set"))
	}
	if b.BrowseName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'browseName' not set"))
	}
	if b.NodeAttributes == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nodeAttributes' not set"))
	}
	if b.TypeDefinition == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'typeDefinition' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AddNodesItem.deepCopy(), nil
}

func (b *_AddNodesItemBuilder) MustBuild() AddNodesItem {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AddNodesItemBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_AddNodesItemBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_AddNodesItemBuilder) DeepCopy() any {
	_copy := b.CreateAddNodesItemBuilder().(*_AddNodesItemBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAddNodesItemBuilder creates a AddNodesItemBuilder
func (b *_AddNodesItem) CreateAddNodesItemBuilder() AddNodesItemBuilder {
	if b == nil {
		return NewAddNodesItemBuilder()
	}
	return &_AddNodesItemBuilder{_AddNodesItem: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AddNodesItem) GetIdentifier() string {
	return "378"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AddNodesItem) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AddNodesItem) GetParentNodeId() ExpandedNodeId {
	return m.ParentNodeId
}

func (m *_AddNodesItem) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_AddNodesItem) GetRequestedNewNodeId() ExpandedNodeId {
	return m.RequestedNewNodeId
}

func (m *_AddNodesItem) GetBrowseName() QualifiedName {
	return m.BrowseName
}

func (m *_AddNodesItem) GetNodeClass() NodeClass {
	return m.NodeClass
}

func (m *_AddNodesItem) GetNodeAttributes() ExtensionObject {
	return m.NodeAttributes
}

func (m *_AddNodesItem) GetTypeDefinition() ExpandedNodeId {
	return m.TypeDefinition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAddNodesItem(structType any) AddNodesItem {
	if casted, ok := structType.(AddNodesItem); ok {
		return casted
	}
	if casted, ok := structType.(*AddNodesItem); ok {
		return *casted
	}
	return nil
}

func (m *_AddNodesItem) GetTypeName() string {
	return "AddNodesItem"
}

func (m *_AddNodesItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (parentNodeId)
	lengthInBits += m.ParentNodeId.GetLengthInBits(ctx)

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Simple field (requestedNewNodeId)
	lengthInBits += m.RequestedNewNodeId.GetLengthInBits(ctx)

	// Simple field (browseName)
	lengthInBits += m.BrowseName.GetLengthInBits(ctx)

	// Simple field (nodeClass)
	lengthInBits += 32

	// Simple field (nodeAttributes)
	lengthInBits += m.NodeAttributes.GetLengthInBits(ctx)

	// Simple field (typeDefinition)
	lengthInBits += m.TypeDefinition.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AddNodesItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AddNodesItem) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__addNodesItem AddNodesItem, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AddNodesItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AddNodesItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	parentNodeId, err := ReadSimpleField[ExpandedNodeId](ctx, "parentNodeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parentNodeId' field"))
	}
	m.ParentNodeId = parentNodeId

	referenceTypeId, err := ReadSimpleField[NodeId](ctx, "referenceTypeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceTypeId' field"))
	}
	m.ReferenceTypeId = referenceTypeId

	requestedNewNodeId, err := ReadSimpleField[ExpandedNodeId](ctx, "requestedNewNodeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedNewNodeId' field"))
	}
	m.RequestedNewNodeId = requestedNewNodeId

	browseName, err := ReadSimpleField[QualifiedName](ctx, "browseName", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'browseName' field"))
	}
	m.BrowseName = browseName

	nodeClass, err := ReadEnumField[NodeClass](ctx, "nodeClass", "NodeClass", ReadEnum(NodeClassByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeClass' field"))
	}
	m.NodeClass = nodeClass

	nodeAttributes, err := ReadSimpleField[ExtensionObject](ctx, "nodeAttributes", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeAttributes' field"))
	}
	m.NodeAttributes = nodeAttributes

	typeDefinition, err := ReadSimpleField[ExpandedNodeId](ctx, "typeDefinition", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeDefinition' field"))
	}
	m.TypeDefinition = typeDefinition

	if closeErr := readBuffer.CloseContext("AddNodesItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AddNodesItem")
	}

	return m, nil
}

func (m *_AddNodesItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AddNodesItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AddNodesItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AddNodesItem")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "parentNodeId", m.GetParentNodeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parentNodeId' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "referenceTypeId", m.GetReferenceTypeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceTypeId' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "requestedNewNodeId", m.GetRequestedNewNodeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedNewNodeId' field")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "browseName", m.GetBrowseName(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'browseName' field")
		}

		if err := WriteSimpleEnumField[NodeClass](ctx, "nodeClass", "NodeClass", m.GetNodeClass(), WriteEnum[NodeClass, uint32](NodeClass.GetValue, NodeClass.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeClass' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "nodeAttributes", m.GetNodeAttributes(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeAttributes' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "typeDefinition", m.GetTypeDefinition(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'typeDefinition' field")
		}

		if popErr := writeBuffer.PopContext("AddNodesItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AddNodesItem")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AddNodesItem) IsAddNodesItem() {}

func (m *_AddNodesItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AddNodesItem) deepCopy() *_AddNodesItem {
	if m == nil {
		return nil
	}
	_AddNodesItemCopy := &_AddNodesItem{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ParentNodeId.DeepCopy().(ExpandedNodeId),
		m.ReferenceTypeId.DeepCopy().(NodeId),
		m.RequestedNewNodeId.DeepCopy().(ExpandedNodeId),
		m.BrowseName.DeepCopy().(QualifiedName),
		m.NodeClass,
		m.NodeAttributes.DeepCopy().(ExtensionObject),
		m.TypeDefinition.DeepCopy().(ExpandedNodeId),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _AddNodesItemCopy
}

func (m *_AddNodesItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

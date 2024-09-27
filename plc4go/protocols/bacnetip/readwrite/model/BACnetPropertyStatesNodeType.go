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

// BACnetPropertyStatesNodeType is the corresponding interface of BACnetPropertyStatesNodeType
type BACnetPropertyStatesNodeType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetNodeType returns NodeType (property field)
	GetNodeType() BACnetNodeTypeTagged
	// IsBACnetPropertyStatesNodeType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesNodeType()
	// CreateBuilder creates a BACnetPropertyStatesNodeTypeBuilder
	CreateBACnetPropertyStatesNodeTypeBuilder() BACnetPropertyStatesNodeTypeBuilder
}

// _BACnetPropertyStatesNodeType is the data-structure of this message
type _BACnetPropertyStatesNodeType struct {
	BACnetPropertyStatesContract
	NodeType BACnetNodeTypeTagged
}

var _ BACnetPropertyStatesNodeType = (*_BACnetPropertyStatesNodeType)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesNodeType)(nil)

// NewBACnetPropertyStatesNodeType factory function for _BACnetPropertyStatesNodeType
func NewBACnetPropertyStatesNodeType(peekedTagHeader BACnetTagHeader, nodeType BACnetNodeTypeTagged) *_BACnetPropertyStatesNodeType {
	if nodeType == nil {
		panic("nodeType of type BACnetNodeTypeTagged for BACnetPropertyStatesNodeType must not be nil")
	}
	_result := &_BACnetPropertyStatesNodeType{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		NodeType:                     nodeType,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesNodeTypeBuilder is a builder for BACnetPropertyStatesNodeType
type BACnetPropertyStatesNodeTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nodeType BACnetNodeTypeTagged) BACnetPropertyStatesNodeTypeBuilder
	// WithNodeType adds NodeType (property field)
	WithNodeType(BACnetNodeTypeTagged) BACnetPropertyStatesNodeTypeBuilder
	// WithNodeTypeBuilder adds NodeType (property field) which is build by the builder
	WithNodeTypeBuilder(func(BACnetNodeTypeTaggedBuilder) BACnetNodeTypeTaggedBuilder) BACnetPropertyStatesNodeTypeBuilder
	// Build builds the BACnetPropertyStatesNodeType or returns an error if something is wrong
	Build() (BACnetPropertyStatesNodeType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesNodeType
}

// NewBACnetPropertyStatesNodeTypeBuilder() creates a BACnetPropertyStatesNodeTypeBuilder
func NewBACnetPropertyStatesNodeTypeBuilder() BACnetPropertyStatesNodeTypeBuilder {
	return &_BACnetPropertyStatesNodeTypeBuilder{_BACnetPropertyStatesNodeType: new(_BACnetPropertyStatesNodeType)}
}

type _BACnetPropertyStatesNodeTypeBuilder struct {
	*_BACnetPropertyStatesNodeType

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesNodeTypeBuilder) = (*_BACnetPropertyStatesNodeTypeBuilder)(nil)

func (b *_BACnetPropertyStatesNodeTypeBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesNodeTypeBuilder) WithMandatoryFields(nodeType BACnetNodeTypeTagged) BACnetPropertyStatesNodeTypeBuilder {
	return b.WithNodeType(nodeType)
}

func (b *_BACnetPropertyStatesNodeTypeBuilder) WithNodeType(nodeType BACnetNodeTypeTagged) BACnetPropertyStatesNodeTypeBuilder {
	b.NodeType = nodeType
	return b
}

func (b *_BACnetPropertyStatesNodeTypeBuilder) WithNodeTypeBuilder(builderSupplier func(BACnetNodeTypeTaggedBuilder) BACnetNodeTypeTaggedBuilder) BACnetPropertyStatesNodeTypeBuilder {
	builder := builderSupplier(b.NodeType.CreateBACnetNodeTypeTaggedBuilder())
	var err error
	b.NodeType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetNodeTypeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesNodeTypeBuilder) Build() (BACnetPropertyStatesNodeType, error) {
	if b.NodeType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nodeType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesNodeType.deepCopy(), nil
}

func (b *_BACnetPropertyStatesNodeTypeBuilder) MustBuild() BACnetPropertyStatesNodeType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesNodeTypeBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesNodeTypeBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesNodeTypeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesNodeTypeBuilder().(*_BACnetPropertyStatesNodeTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesNodeTypeBuilder creates a BACnetPropertyStatesNodeTypeBuilder
func (b *_BACnetPropertyStatesNodeType) CreateBACnetPropertyStatesNodeTypeBuilder() BACnetPropertyStatesNodeTypeBuilder {
	if b == nil {
		return NewBACnetPropertyStatesNodeTypeBuilder()
	}
	return &_BACnetPropertyStatesNodeTypeBuilder{_BACnetPropertyStatesNodeType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesNodeType) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesNodeType) GetNodeType() BACnetNodeTypeTagged {
	return m.NodeType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesNodeType(structType any) BACnetPropertyStatesNodeType {
	if casted, ok := structType.(BACnetPropertyStatesNodeType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNodeType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesNodeType) GetTypeName() string {
	return "BACnetPropertyStatesNodeType"
}

func (m *_BACnetPropertyStatesNodeType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (nodeType)
	lengthInBits += m.NodeType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesNodeType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesNodeType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesNodeType BACnetPropertyStatesNodeType, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNodeType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesNodeType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeType, err := ReadSimpleField[BACnetNodeTypeTagged](ctx, "nodeType", ReadComplex[BACnetNodeTypeTagged](BACnetNodeTypeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeType' field"))
	}
	m.NodeType = nodeType

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNodeType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesNodeType")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesNodeType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesNodeType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNodeType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesNodeType")
		}

		if err := WriteSimpleField[BACnetNodeTypeTagged](ctx, "nodeType", m.GetNodeType(), WriteComplex[BACnetNodeTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNodeType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesNodeType")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesNodeType) IsBACnetPropertyStatesNodeType() {}

func (m *_BACnetPropertyStatesNodeType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesNodeType) deepCopy() *_BACnetPropertyStatesNodeType {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesNodeTypeCopy := &_BACnetPropertyStatesNodeType{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.NodeType.DeepCopy().(BACnetNodeTypeTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesNodeTypeCopy
}

func (m *_BACnetPropertyStatesNodeType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

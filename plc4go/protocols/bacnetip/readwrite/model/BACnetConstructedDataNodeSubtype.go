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

// BACnetConstructedDataNodeSubtype is the corresponding interface of BACnetConstructedDataNodeSubtype
type BACnetConstructedDataNodeSubtype interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNodeSubType returns NodeSubType (property field)
	GetNodeSubType() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataNodeSubtype is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNodeSubtype()
	// CreateBuilder creates a BACnetConstructedDataNodeSubtypeBuilder
	CreateBACnetConstructedDataNodeSubtypeBuilder() BACnetConstructedDataNodeSubtypeBuilder
}

// _BACnetConstructedDataNodeSubtype is the data-structure of this message
type _BACnetConstructedDataNodeSubtype struct {
	BACnetConstructedDataContract
	NodeSubType BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataNodeSubtype = (*_BACnetConstructedDataNodeSubtype)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNodeSubtype)(nil)

// NewBACnetConstructedDataNodeSubtype factory function for _BACnetConstructedDataNodeSubtype
func NewBACnetConstructedDataNodeSubtype(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, nodeSubType BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNodeSubtype {
	if nodeSubType == nil {
		panic("nodeSubType of type BACnetApplicationTagCharacterString for BACnetConstructedDataNodeSubtype must not be nil")
	}
	_result := &_BACnetConstructedDataNodeSubtype{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NodeSubType:                   nodeSubType,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataNodeSubtypeBuilder is a builder for BACnetConstructedDataNodeSubtype
type BACnetConstructedDataNodeSubtypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nodeSubType BACnetApplicationTagCharacterString) BACnetConstructedDataNodeSubtypeBuilder
	// WithNodeSubType adds NodeSubType (property field)
	WithNodeSubType(BACnetApplicationTagCharacterString) BACnetConstructedDataNodeSubtypeBuilder
	// WithNodeSubTypeBuilder adds NodeSubType (property field) which is build by the builder
	WithNodeSubTypeBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataNodeSubtypeBuilder
	// Build builds the BACnetConstructedDataNodeSubtype or returns an error if something is wrong
	Build() (BACnetConstructedDataNodeSubtype, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataNodeSubtype
}

// NewBACnetConstructedDataNodeSubtypeBuilder() creates a BACnetConstructedDataNodeSubtypeBuilder
func NewBACnetConstructedDataNodeSubtypeBuilder() BACnetConstructedDataNodeSubtypeBuilder {
	return &_BACnetConstructedDataNodeSubtypeBuilder{_BACnetConstructedDataNodeSubtype: new(_BACnetConstructedDataNodeSubtype)}
}

type _BACnetConstructedDataNodeSubtypeBuilder struct {
	*_BACnetConstructedDataNodeSubtype

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataNodeSubtypeBuilder) = (*_BACnetConstructedDataNodeSubtypeBuilder)(nil)

func (b *_BACnetConstructedDataNodeSubtypeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataNodeSubtypeBuilder) WithMandatoryFields(nodeSubType BACnetApplicationTagCharacterString) BACnetConstructedDataNodeSubtypeBuilder {
	return b.WithNodeSubType(nodeSubType)
}

func (b *_BACnetConstructedDataNodeSubtypeBuilder) WithNodeSubType(nodeSubType BACnetApplicationTagCharacterString) BACnetConstructedDataNodeSubtypeBuilder {
	b.NodeSubType = nodeSubType
	return b
}

func (b *_BACnetConstructedDataNodeSubtypeBuilder) WithNodeSubTypeBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataNodeSubtypeBuilder {
	builder := builderSupplier(b.NodeSubType.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.NodeSubType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataNodeSubtypeBuilder) Build() (BACnetConstructedDataNodeSubtype, error) {
	if b.NodeSubType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nodeSubType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataNodeSubtype.deepCopy(), nil
}

func (b *_BACnetConstructedDataNodeSubtypeBuilder) MustBuild() BACnetConstructedDataNodeSubtype {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataNodeSubtypeBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataNodeSubtypeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataNodeSubtypeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataNodeSubtypeBuilder().(*_BACnetConstructedDataNodeSubtypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataNodeSubtypeBuilder creates a BACnetConstructedDataNodeSubtypeBuilder
func (b *_BACnetConstructedDataNodeSubtype) CreateBACnetConstructedDataNodeSubtypeBuilder() BACnetConstructedDataNodeSubtypeBuilder {
	if b == nil {
		return NewBACnetConstructedDataNodeSubtypeBuilder()
	}
	return &_BACnetConstructedDataNodeSubtypeBuilder{_BACnetConstructedDataNodeSubtype: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNodeSubtype) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNodeSubtype) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NODE_SUBTYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNodeSubtype) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNodeSubtype) GetNodeSubType() BACnetApplicationTagCharacterString {
	return m.NodeSubType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNodeSubtype) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetNodeSubType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNodeSubtype(structType any) BACnetConstructedDataNodeSubtype {
	if casted, ok := structType.(BACnetConstructedDataNodeSubtype); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNodeSubtype); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNodeSubtype) GetTypeName() string {
	return "BACnetConstructedDataNodeSubtype"
}

func (m *_BACnetConstructedDataNodeSubtype) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (nodeSubType)
	lengthInBits += m.NodeSubType.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNodeSubtype) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNodeSubtype) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNodeSubtype BACnetConstructedDataNodeSubtype, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNodeSubtype"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNodeSubtype")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeSubType, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "nodeSubType", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeSubType' field"))
	}
	m.NodeSubType = nodeSubType

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), nodeSubType)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNodeSubtype"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNodeSubtype")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNodeSubtype) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNodeSubtype) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNodeSubtype"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNodeSubtype")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "nodeSubType", m.GetNodeSubType(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeSubType' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNodeSubtype"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNodeSubtype")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNodeSubtype) IsBACnetConstructedDataNodeSubtype() {}

func (m *_BACnetConstructedDataNodeSubtype) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataNodeSubtype) deepCopy() *_BACnetConstructedDataNodeSubtype {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataNodeSubtypeCopy := &_BACnetConstructedDataNodeSubtype{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.NodeSubType),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataNodeSubtypeCopy
}

func (m *_BACnetConstructedDataNodeSubtype) String() string {
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

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

// BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier is the corresponding interface of BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier
type BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetUnconfirmedServiceRequestWhoHasObject
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// IsBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier()
	// CreateBuilder creates a BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder
	CreateBACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder() BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder
}

// _BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier is the data-structure of this message
type _BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier struct {
	BACnetUnconfirmedServiceRequestWhoHasObjectContract
	ObjectIdentifier BACnetContextTagObjectIdentifier
}

var _ BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier = (*_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier)(nil)
var _ BACnetUnconfirmedServiceRequestWhoHasObjectRequirements = (*_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier)(nil)

// NewBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier factory function for _BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier
func NewBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier(peekedTagHeader BACnetTagHeader, objectIdentifier BACnetContextTagObjectIdentifier) *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetContextTagObjectIdentifier for BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier must not be nil")
	}
	_result := &_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier{
		BACnetUnconfirmedServiceRequestWhoHasObjectContract: NewBACnetUnconfirmedServiceRequestWhoHasObject(peekedTagHeader),
		ObjectIdentifier: objectIdentifier,
	}
	_result.BACnetUnconfirmedServiceRequestWhoHasObjectContract.(*_BACnetUnconfirmedServiceRequestWhoHasObject)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder is a builder for BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier
type BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier) BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder
	// Build builds the BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier or returns an error if something is wrong
	Build() (BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier
}

// NewBACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder() creates a BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder
func NewBACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder() BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder {
	return &_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder{_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier: new(_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier)}
}

type _BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder struct {
	*_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier

	parentBuilder *_BACnetUnconfirmedServiceRequestWhoHasObjectBuilder

	err *utils.MultiError
}

var _ (BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder) = (*_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder)(nil)

func (b *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder) setParent(contract BACnetUnconfirmedServiceRequestWhoHasObjectContract) {
	b.BACnetUnconfirmedServiceRequestWhoHasObjectContract = contract
}

func (b *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder) WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier) BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder {
	return b.WithObjectIdentifier(objectIdentifier)
}

func (b *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder) WithObjectIdentifier(objectIdentifier BACnetContextTagObjectIdentifier) BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder {
	b.ObjectIdentifier = objectIdentifier
	return b
}

func (b *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder) WithObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder {
	builder := builderSupplier(b.ObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.ObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder) Build() (BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier, error) {
	if b.ObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier.deepCopy(), nil
}

func (b *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder) MustBuild() BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder) Done() BACnetUnconfirmedServiceRequestWhoHasObjectBuilder {
	return b.parentBuilder
}

func (b *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder) buildForBACnetUnconfirmedServiceRequestWhoHasObject() (BACnetUnconfirmedServiceRequestWhoHasObject, error) {
	return b.Build()
}

func (b *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder) DeepCopy() any {
	_copy := b.CreateBACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder().(*_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder creates a BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder
func (b *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) CreateBACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder() BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder {
	if b == nil {
		return NewBACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder()
	}
	return &_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierBuilder{_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier: b.deepCopy()}
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

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetParent() BACnetUnconfirmedServiceRequestWhoHasObjectContract {
	return m.BACnetUnconfirmedServiceRequestWhoHasObjectContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier(structType any) BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier"
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestWhoHasObjectContract.(*_BACnetUnconfirmedServiceRequestWhoHasObject).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequestWhoHasObject) (__bACnetUnconfirmedServiceRequestWhoHasObjectIdentifier BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier, err error) {
	m.BACnetUnconfirmedServiceRequestWhoHasObjectContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestWhoHasObjectContract.(*_BACnetUnconfirmedServiceRequestWhoHasObject).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) IsBACnetUnconfirmedServiceRequestWhoHasObjectIdentifier() {
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) deepCopy() *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier {
	if m == nil {
		return nil
	}
	_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierCopy := &_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier{
		m.BACnetUnconfirmedServiceRequestWhoHasObjectContract.(*_BACnetUnconfirmedServiceRequestWhoHasObject).deepCopy(),
		utils.DeepCopy[BACnetContextTagObjectIdentifier](m.ObjectIdentifier),
	}
	m.BACnetUnconfirmedServiceRequestWhoHasObjectContract.(*_BACnetUnconfirmedServiceRequestWhoHasObject)._SubType = m
	return _BACnetUnconfirmedServiceRequestWhoHasObjectIdentifierCopy
}

func (m *_BACnetUnconfirmedServiceRequestWhoHasObjectIdentifier) String() string {
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

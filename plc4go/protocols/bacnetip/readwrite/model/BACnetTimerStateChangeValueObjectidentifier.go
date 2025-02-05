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

// BACnetTimerStateChangeValueObjectidentifier is the corresponding interface of BACnetTimerStateChangeValueObjectidentifier
type BACnetTimerStateChangeValueObjectidentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimerStateChangeValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier
	// IsBACnetTimerStateChangeValueObjectidentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueObjectidentifier()
	// CreateBuilder creates a BACnetTimerStateChangeValueObjectidentifierBuilder
	CreateBACnetTimerStateChangeValueObjectidentifierBuilder() BACnetTimerStateChangeValueObjectidentifierBuilder
}

// _BACnetTimerStateChangeValueObjectidentifier is the data-structure of this message
type _BACnetTimerStateChangeValueObjectidentifier struct {
	BACnetTimerStateChangeValueContract
	ObjectidentifierValue BACnetApplicationTagObjectIdentifier
}

var _ BACnetTimerStateChangeValueObjectidentifier = (*_BACnetTimerStateChangeValueObjectidentifier)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueObjectidentifier)(nil)

// NewBACnetTimerStateChangeValueObjectidentifier factory function for _BACnetTimerStateChangeValueObjectidentifier
func NewBACnetTimerStateChangeValueObjectidentifier(peekedTagHeader BACnetTagHeader, objectidentifierValue BACnetApplicationTagObjectIdentifier, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueObjectidentifier {
	if objectidentifierValue == nil {
		panic("objectidentifierValue of type BACnetApplicationTagObjectIdentifier for BACnetTimerStateChangeValueObjectidentifier must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueObjectidentifier{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		ObjectidentifierValue:               objectidentifierValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueObjectidentifierBuilder is a builder for BACnetTimerStateChangeValueObjectidentifier
type BACnetTimerStateChangeValueObjectidentifierBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectidentifierValue BACnetApplicationTagObjectIdentifier) BACnetTimerStateChangeValueObjectidentifierBuilder
	// WithObjectidentifierValue adds ObjectidentifierValue (property field)
	WithObjectidentifierValue(BACnetApplicationTagObjectIdentifier) BACnetTimerStateChangeValueObjectidentifierBuilder
	// WithObjectidentifierValueBuilder adds ObjectidentifierValue (property field) which is build by the builder
	WithObjectidentifierValueBuilder(func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetTimerStateChangeValueObjectidentifierBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetTimerStateChangeValueBuilder
	// Build builds the BACnetTimerStateChangeValueObjectidentifier or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValueObjectidentifier, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValueObjectidentifier
}

// NewBACnetTimerStateChangeValueObjectidentifierBuilder() creates a BACnetTimerStateChangeValueObjectidentifierBuilder
func NewBACnetTimerStateChangeValueObjectidentifierBuilder() BACnetTimerStateChangeValueObjectidentifierBuilder {
	return &_BACnetTimerStateChangeValueObjectidentifierBuilder{_BACnetTimerStateChangeValueObjectidentifier: new(_BACnetTimerStateChangeValueObjectidentifier)}
}

type _BACnetTimerStateChangeValueObjectidentifierBuilder struct {
	*_BACnetTimerStateChangeValueObjectidentifier

	parentBuilder *_BACnetTimerStateChangeValueBuilder

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueObjectidentifierBuilder) = (*_BACnetTimerStateChangeValueObjectidentifierBuilder)(nil)

func (b *_BACnetTimerStateChangeValueObjectidentifierBuilder) setParent(contract BACnetTimerStateChangeValueContract) {
	b.BACnetTimerStateChangeValueContract = contract
	contract.(*_BACnetTimerStateChangeValue)._SubType = b._BACnetTimerStateChangeValueObjectidentifier
}

func (b *_BACnetTimerStateChangeValueObjectidentifierBuilder) WithMandatoryFields(objectidentifierValue BACnetApplicationTagObjectIdentifier) BACnetTimerStateChangeValueObjectidentifierBuilder {
	return b.WithObjectidentifierValue(objectidentifierValue)
}

func (b *_BACnetTimerStateChangeValueObjectidentifierBuilder) WithObjectidentifierValue(objectidentifierValue BACnetApplicationTagObjectIdentifier) BACnetTimerStateChangeValueObjectidentifierBuilder {
	b.ObjectidentifierValue = objectidentifierValue
	return b
}

func (b *_BACnetTimerStateChangeValueObjectidentifierBuilder) WithObjectidentifierValueBuilder(builderSupplier func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetTimerStateChangeValueObjectidentifierBuilder {
	builder := builderSupplier(b.ObjectidentifierValue.CreateBACnetApplicationTagObjectIdentifierBuilder())
	var err error
	b.ObjectidentifierValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetTimerStateChangeValueObjectidentifierBuilder) Build() (BACnetTimerStateChangeValueObjectidentifier, error) {
	if b.ObjectidentifierValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectidentifierValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerStateChangeValueObjectidentifier.deepCopy(), nil
}

func (b *_BACnetTimerStateChangeValueObjectidentifierBuilder) MustBuild() BACnetTimerStateChangeValueObjectidentifier {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTimerStateChangeValueObjectidentifierBuilder) Done() BACnetTimerStateChangeValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetTimerStateChangeValueBuilder().(*_BACnetTimerStateChangeValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetTimerStateChangeValueObjectidentifierBuilder) buildForBACnetTimerStateChangeValue() (BACnetTimerStateChangeValue, error) {
	return b.Build()
}

func (b *_BACnetTimerStateChangeValueObjectidentifierBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerStateChangeValueObjectidentifierBuilder().(*_BACnetTimerStateChangeValueObjectidentifierBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerStateChangeValueObjectidentifierBuilder creates a BACnetTimerStateChangeValueObjectidentifierBuilder
func (b *_BACnetTimerStateChangeValueObjectidentifier) CreateBACnetTimerStateChangeValueObjectidentifierBuilder() BACnetTimerStateChangeValueObjectidentifierBuilder {
	if b == nil {
		return NewBACnetTimerStateChangeValueObjectidentifierBuilder()
	}
	return &_BACnetTimerStateChangeValueObjectidentifierBuilder{_BACnetTimerStateChangeValueObjectidentifier: b.deepCopy()}
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

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueObjectidentifier(structType any) BACnetTimerStateChangeValueObjectidentifier {
	if casted, ok := structType.(BACnetTimerStateChangeValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueObjectidentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetTypeName() string {
	return "BACnetTimerStateChangeValueObjectidentifier"
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueObjectidentifier BACnetTimerStateChangeValueObjectidentifier, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueObjectidentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueObjectidentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectidentifierValue, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectidentifierValue", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectidentifierValue' field"))
	}
	m.ObjectidentifierValue = objectidentifierValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueObjectidentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueObjectidentifier")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueObjectidentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueObjectidentifier")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectidentifierValue", m.GetObjectidentifierValue(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueObjectidentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueObjectidentifier")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) IsBACnetTimerStateChangeValueObjectidentifier() {
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) deepCopy() *_BACnetTimerStateChangeValueObjectidentifier {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueObjectidentifierCopy := &_BACnetTimerStateChangeValueObjectidentifier{
		m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagObjectIdentifier](m.ObjectidentifierValue),
	}
	_BACnetTimerStateChangeValueObjectidentifierCopy.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = m
	return _BACnetTimerStateChangeValueObjectidentifierCopy
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) String() string {
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

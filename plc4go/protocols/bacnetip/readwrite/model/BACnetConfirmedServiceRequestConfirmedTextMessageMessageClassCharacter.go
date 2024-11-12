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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetCharacterValue returns CharacterValue (property field)
	GetCharacterValue() BACnetContextTagCharacterString
	// IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter()
	// CreateBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
	CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
}

// _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter struct {
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
	CharacterValue BACnetContextTagCharacterString
}

var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter)(nil)
var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter)(nil)

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter factory function for _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, characterValue BACnetContextTagCharacterString, tagNumber uint8) *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	if characterValue == nil {
		panic("characterValue of type BACnetContextTagCharacterString for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter{
		BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract: NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(openingTag, peekedTagHeader, closingTag, tagNumber),
		CharacterValue: characterValue,
	}
	_result.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder is a builder for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(characterValue BACnetContextTagCharacterString) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
	// WithCharacterValue adds CharacterValue (property field)
	WithCharacterValue(BACnetContextTagCharacterString) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
	// WithCharacterValueBuilder adds CharacterValue (property field) which is build by the builder
	WithCharacterValueBuilder(func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
	// Build builds the BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
}

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder() creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder{_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter: new(_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter)}
}

type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder struct {
	*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter

	parentBuilder *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) setParent(contract BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract) {
	b.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract = contract
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) WithMandatoryFields(characterValue BACnetContextTagCharacterString) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	return b.WithCharacterValue(characterValue)
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) WithCharacterValue(characterValue BACnetContextTagCharacterString) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	b.CharacterValue = characterValue
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) WithCharacterValueBuilder(builderSupplier func(BACnetContextTagCharacterStringBuilder) BACnetContextTagCharacterStringBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	builder := builderSupplier(b.CharacterValue.CreateBACnetContextTagCharacterStringBuilder())
	var err error
	b.CharacterValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) Build() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter, error) {
	if b.CharacterValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'characterValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) MustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) Done() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) buildForBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder().(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder
func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder()
	}
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterBuilder{_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter: b.deepCopy()}
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

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetParent() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract {
	return m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetCharacterValue() BACnetContextTagCharacterString {
	return m.CharacterValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter(structType any) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).getLengthInBits(ctx))

	// Simple field (characterValue)
	lengthInBits += m.CharacterValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, tagNumber uint8) (__bACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter, err error) {
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	characterValue, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "characterValue", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'characterValue' field"))
	}
	m.CharacterValue = characterValue

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
		}

		if err := WriteSimpleField[BACnetContextTagCharacterString](ctx, "characterValue", m.GetCharacterValue(), WriteComplex[BACnetContextTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'characterValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter() {
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) deepCopy() *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterCopy := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter{
		m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).deepCopy(),
		utils.DeepCopy[BACnetContextTagCharacterString](m.CharacterValue),
	}
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)._SubType = m
	return _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterCopy
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) String() string {
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

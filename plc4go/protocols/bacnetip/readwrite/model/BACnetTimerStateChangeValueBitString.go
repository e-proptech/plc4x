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

// BACnetTimerStateChangeValueBitString is the corresponding interface of BACnetTimerStateChangeValueBitString
type BACnetTimerStateChangeValueBitString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimerStateChangeValue
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetApplicationTagBitString
	// IsBACnetTimerStateChangeValueBitString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueBitString()
	// CreateBuilder creates a BACnetTimerStateChangeValueBitStringBuilder
	CreateBACnetTimerStateChangeValueBitStringBuilder() BACnetTimerStateChangeValueBitStringBuilder
}

// _BACnetTimerStateChangeValueBitString is the data-structure of this message
type _BACnetTimerStateChangeValueBitString struct {
	BACnetTimerStateChangeValueContract
	BitStringValue BACnetApplicationTagBitString
}

var _ BACnetTimerStateChangeValueBitString = (*_BACnetTimerStateChangeValueBitString)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueBitString)(nil)

// NewBACnetTimerStateChangeValueBitString factory function for _BACnetTimerStateChangeValueBitString
func NewBACnetTimerStateChangeValueBitString(peekedTagHeader BACnetTagHeader, bitStringValue BACnetApplicationTagBitString, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueBitString {
	if bitStringValue == nil {
		panic("bitStringValue of type BACnetApplicationTagBitString for BACnetTimerStateChangeValueBitString must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueBitString{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		BitStringValue:                      bitStringValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueBitStringBuilder is a builder for BACnetTimerStateChangeValueBitString
type BACnetTimerStateChangeValueBitStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bitStringValue BACnetApplicationTagBitString) BACnetTimerStateChangeValueBitStringBuilder
	// WithBitStringValue adds BitStringValue (property field)
	WithBitStringValue(BACnetApplicationTagBitString) BACnetTimerStateChangeValueBitStringBuilder
	// WithBitStringValueBuilder adds BitStringValue (property field) which is build by the builder
	WithBitStringValueBuilder(func(BACnetApplicationTagBitStringBuilder) BACnetApplicationTagBitStringBuilder) BACnetTimerStateChangeValueBitStringBuilder
	// Build builds the BACnetTimerStateChangeValueBitString or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValueBitString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValueBitString
}

// NewBACnetTimerStateChangeValueBitStringBuilder() creates a BACnetTimerStateChangeValueBitStringBuilder
func NewBACnetTimerStateChangeValueBitStringBuilder() BACnetTimerStateChangeValueBitStringBuilder {
	return &_BACnetTimerStateChangeValueBitStringBuilder{_BACnetTimerStateChangeValueBitString: new(_BACnetTimerStateChangeValueBitString)}
}

type _BACnetTimerStateChangeValueBitStringBuilder struct {
	*_BACnetTimerStateChangeValueBitString

	parentBuilder *_BACnetTimerStateChangeValueBuilder

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueBitStringBuilder) = (*_BACnetTimerStateChangeValueBitStringBuilder)(nil)

func (b *_BACnetTimerStateChangeValueBitStringBuilder) setParent(contract BACnetTimerStateChangeValueContract) {
	b.BACnetTimerStateChangeValueContract = contract
}

func (b *_BACnetTimerStateChangeValueBitStringBuilder) WithMandatoryFields(bitStringValue BACnetApplicationTagBitString) BACnetTimerStateChangeValueBitStringBuilder {
	return b.WithBitStringValue(bitStringValue)
}

func (b *_BACnetTimerStateChangeValueBitStringBuilder) WithBitStringValue(bitStringValue BACnetApplicationTagBitString) BACnetTimerStateChangeValueBitStringBuilder {
	b.BitStringValue = bitStringValue
	return b
}

func (b *_BACnetTimerStateChangeValueBitStringBuilder) WithBitStringValueBuilder(builderSupplier func(BACnetApplicationTagBitStringBuilder) BACnetApplicationTagBitStringBuilder) BACnetTimerStateChangeValueBitStringBuilder {
	builder := builderSupplier(b.BitStringValue.CreateBACnetApplicationTagBitStringBuilder())
	var err error
	b.BitStringValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBitStringBuilder failed"))
	}
	return b
}

func (b *_BACnetTimerStateChangeValueBitStringBuilder) Build() (BACnetTimerStateChangeValueBitString, error) {
	if b.BitStringValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'bitStringValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerStateChangeValueBitString.deepCopy(), nil
}

func (b *_BACnetTimerStateChangeValueBitStringBuilder) MustBuild() BACnetTimerStateChangeValueBitString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetTimerStateChangeValueBitStringBuilder) Done() BACnetTimerStateChangeValueBuilder {
	return b.parentBuilder
}

func (b *_BACnetTimerStateChangeValueBitStringBuilder) buildForBACnetTimerStateChangeValue() (BACnetTimerStateChangeValue, error) {
	return b.Build()
}

func (b *_BACnetTimerStateChangeValueBitStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerStateChangeValueBitStringBuilder().(*_BACnetTimerStateChangeValueBitStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerStateChangeValueBitStringBuilder creates a BACnetTimerStateChangeValueBitStringBuilder
func (b *_BACnetTimerStateChangeValueBitString) CreateBACnetTimerStateChangeValueBitStringBuilder() BACnetTimerStateChangeValueBitStringBuilder {
	if b == nil {
		return NewBACnetTimerStateChangeValueBitStringBuilder()
	}
	return &_BACnetTimerStateChangeValueBitStringBuilder{_BACnetTimerStateChangeValueBitString: b.deepCopy()}
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

func (m *_BACnetTimerStateChangeValueBitString) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueBitString) GetBitStringValue() BACnetApplicationTagBitString {
	return m.BitStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueBitString(structType any) BACnetTimerStateChangeValueBitString {
	if casted, ok := structType.(BACnetTimerStateChangeValueBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueBitString) GetTypeName() string {
	return "BACnetTimerStateChangeValueBitString"
}

func (m *_BACnetTimerStateChangeValueBitString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueBitString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueBitString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueBitString BACnetTimerStateChangeValueBitString, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bitStringValue, err := ReadSimpleField[BACnetApplicationTagBitString](ctx, "bitStringValue", ReadComplex[BACnetApplicationTagBitString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBitString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitStringValue' field"))
	}
	m.BitStringValue = bitStringValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueBitString")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueBitString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueBitString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueBitString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueBitString")
		}

		if err := WriteSimpleField[BACnetApplicationTagBitString](ctx, "bitStringValue", m.GetBitStringValue(), WriteComplex[BACnetApplicationTagBitString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueBitString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueBitString")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueBitString) IsBACnetTimerStateChangeValueBitString() {}

func (m *_BACnetTimerStateChangeValueBitString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValueBitString) deepCopy() *_BACnetTimerStateChangeValueBitString {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueBitStringCopy := &_BACnetTimerStateChangeValueBitString{
		m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBitString](m.BitStringValue),
	}
	m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = m
	return _BACnetTimerStateChangeValueBitStringCopy
}

func (m *_BACnetTimerStateChangeValueBitString) String() string {
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

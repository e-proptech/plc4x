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

// BACnetShedLevelAmount is the corresponding interface of BACnetShedLevelAmount
type BACnetShedLevelAmount interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetShedLevel
	// GetAmount returns Amount (property field)
	GetAmount() BACnetContextTagReal
	// IsBACnetShedLevelAmount is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetShedLevelAmount()
	// CreateBuilder creates a BACnetShedLevelAmountBuilder
	CreateBACnetShedLevelAmountBuilder() BACnetShedLevelAmountBuilder
}

// _BACnetShedLevelAmount is the data-structure of this message
type _BACnetShedLevelAmount struct {
	BACnetShedLevelContract
	Amount BACnetContextTagReal
}

var _ BACnetShedLevelAmount = (*_BACnetShedLevelAmount)(nil)
var _ BACnetShedLevelRequirements = (*_BACnetShedLevelAmount)(nil)

// NewBACnetShedLevelAmount factory function for _BACnetShedLevelAmount
func NewBACnetShedLevelAmount(peekedTagHeader BACnetTagHeader, amount BACnetContextTagReal) *_BACnetShedLevelAmount {
	if amount == nil {
		panic("amount of type BACnetContextTagReal for BACnetShedLevelAmount must not be nil")
	}
	_result := &_BACnetShedLevelAmount{
		BACnetShedLevelContract: NewBACnetShedLevel(peekedTagHeader),
		Amount:                  amount,
	}
	_result.BACnetShedLevelContract.(*_BACnetShedLevel)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetShedLevelAmountBuilder is a builder for BACnetShedLevelAmount
type BACnetShedLevelAmountBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(amount BACnetContextTagReal) BACnetShedLevelAmountBuilder
	// WithAmount adds Amount (property field)
	WithAmount(BACnetContextTagReal) BACnetShedLevelAmountBuilder
	// WithAmountBuilder adds Amount (property field) which is build by the builder
	WithAmountBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetShedLevelAmountBuilder
	// Build builds the BACnetShedLevelAmount or returns an error if something is wrong
	Build() (BACnetShedLevelAmount, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetShedLevelAmount
}

// NewBACnetShedLevelAmountBuilder() creates a BACnetShedLevelAmountBuilder
func NewBACnetShedLevelAmountBuilder() BACnetShedLevelAmountBuilder {
	return &_BACnetShedLevelAmountBuilder{_BACnetShedLevelAmount: new(_BACnetShedLevelAmount)}
}

type _BACnetShedLevelAmountBuilder struct {
	*_BACnetShedLevelAmount

	parentBuilder *_BACnetShedLevelBuilder

	err *utils.MultiError
}

var _ (BACnetShedLevelAmountBuilder) = (*_BACnetShedLevelAmountBuilder)(nil)

func (b *_BACnetShedLevelAmountBuilder) setParent(contract BACnetShedLevelContract) {
	b.BACnetShedLevelContract = contract
}

func (b *_BACnetShedLevelAmountBuilder) WithMandatoryFields(amount BACnetContextTagReal) BACnetShedLevelAmountBuilder {
	return b.WithAmount(amount)
}

func (b *_BACnetShedLevelAmountBuilder) WithAmount(amount BACnetContextTagReal) BACnetShedLevelAmountBuilder {
	b.Amount = amount
	return b
}

func (b *_BACnetShedLevelAmountBuilder) WithAmountBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetShedLevelAmountBuilder {
	builder := builderSupplier(b.Amount.CreateBACnetContextTagRealBuilder())
	var err error
	b.Amount, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetShedLevelAmountBuilder) Build() (BACnetShedLevelAmount, error) {
	if b.Amount == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'amount' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetShedLevelAmount.deepCopy(), nil
}

func (b *_BACnetShedLevelAmountBuilder) MustBuild() BACnetShedLevelAmount {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetShedLevelAmountBuilder) Done() BACnetShedLevelBuilder {
	return b.parentBuilder
}

func (b *_BACnetShedLevelAmountBuilder) buildForBACnetShedLevel() (BACnetShedLevel, error) {
	return b.Build()
}

func (b *_BACnetShedLevelAmountBuilder) DeepCopy() any {
	_copy := b.CreateBACnetShedLevelAmountBuilder().(*_BACnetShedLevelAmountBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetShedLevelAmountBuilder creates a BACnetShedLevelAmountBuilder
func (b *_BACnetShedLevelAmount) CreateBACnetShedLevelAmountBuilder() BACnetShedLevelAmountBuilder {
	if b == nil {
		return NewBACnetShedLevelAmountBuilder()
	}
	return &_BACnetShedLevelAmountBuilder{_BACnetShedLevelAmount: b.deepCopy()}
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

func (m *_BACnetShedLevelAmount) GetParent() BACnetShedLevelContract {
	return m.BACnetShedLevelContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedLevelAmount) GetAmount() BACnetContextTagReal {
	return m.Amount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetShedLevelAmount(structType any) BACnetShedLevelAmount {
	if casted, ok := structType.(BACnetShedLevelAmount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedLevelAmount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedLevelAmount) GetTypeName() string {
	return "BACnetShedLevelAmount"
}

func (m *_BACnetShedLevelAmount) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetShedLevelContract.(*_BACnetShedLevel).getLengthInBits(ctx))

	// Simple field (amount)
	lengthInBits += m.Amount.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetShedLevelAmount) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetShedLevelAmount) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetShedLevel) (__bACnetShedLevelAmount BACnetShedLevelAmount, err error) {
	m.BACnetShedLevelContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetShedLevelAmount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevelAmount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	amount, err := ReadSimpleField[BACnetContextTagReal](ctx, "amount", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'amount' field"))
	}
	m.Amount = amount

	if closeErr := readBuffer.CloseContext("BACnetShedLevelAmount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevelAmount")
	}

	return m, nil
}

func (m *_BACnetShedLevelAmount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetShedLevelAmount) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetShedLevelAmount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetShedLevelAmount")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "amount", m.GetAmount(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'amount' field")
		}

		if popErr := writeBuffer.PopContext("BACnetShedLevelAmount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetShedLevelAmount")
		}
		return nil
	}
	return m.BACnetShedLevelContract.(*_BACnetShedLevel).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetShedLevelAmount) IsBACnetShedLevelAmount() {}

func (m *_BACnetShedLevelAmount) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetShedLevelAmount) deepCopy() *_BACnetShedLevelAmount {
	if m == nil {
		return nil
	}
	_BACnetShedLevelAmountCopy := &_BACnetShedLevelAmount{
		m.BACnetShedLevelContract.(*_BACnetShedLevel).deepCopy(),
		m.Amount.DeepCopy().(BACnetContextTagReal),
	}
	m.BACnetShedLevelContract.(*_BACnetShedLevel)._SubType = m
	return _BACnetShedLevelAmountCopy
}

func (m *_BACnetShedLevelAmount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement is the corresponding interface of BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
type BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameterChangeOfValueCivCriteria
	// GetReferencedPropertyIncrement returns ReferencedPropertyIncrement (property field)
	GetReferencedPropertyIncrement() BACnetContextTagReal
	// IsBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement()
	// CreateBuilder creates a BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder
	CreateBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder() BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder
}

// _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement is the data-structure of this message
type _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement struct {
	BACnetEventParameterChangeOfValueCivCriteriaContract
	ReferencedPropertyIncrement BACnetContextTagReal
}

var _ BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement = (*_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement)(nil)
var _ BACnetEventParameterChangeOfValueCivCriteriaRequirements = (*_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement)(nil)

// NewBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement factory function for _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
func NewBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, referencedPropertyIncrement BACnetContextTagReal, tagNumber uint8) *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement {
	if referencedPropertyIncrement == nil {
		panic("referencedPropertyIncrement of type BACnetContextTagReal for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement must not be nil")
	}
	_result := &_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement{
		BACnetEventParameterChangeOfValueCivCriteriaContract: NewBACnetEventParameterChangeOfValueCivCriteria(openingTag, peekedTagHeader, closingTag, tagNumber),
		ReferencedPropertyIncrement:                          referencedPropertyIncrement,
	}
	_result.BACnetEventParameterChangeOfValueCivCriteriaContract.(*_BACnetEventParameterChangeOfValueCivCriteria)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder is a builder for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
type BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(referencedPropertyIncrement BACnetContextTagReal) BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder
	// WithReferencedPropertyIncrement adds ReferencedPropertyIncrement (property field)
	WithReferencedPropertyIncrement(BACnetContextTagReal) BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder
	// WithReferencedPropertyIncrementBuilder adds ReferencedPropertyIncrement (property field) which is build by the builder
	WithReferencedPropertyIncrementBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder
	// Build builds the BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement or returns an error if something is wrong
	Build() (BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
}

// NewBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder() creates a BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder
func NewBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder() BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder {
	return &_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder{_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement: new(_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement)}
}

type _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder struct {
	*_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement

	parentBuilder *_BACnetEventParameterChangeOfValueCivCriteriaBuilder

	err *utils.MultiError
}

var _ (BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder) = (*_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder)(nil)

func (b *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder) setParent(contract BACnetEventParameterChangeOfValueCivCriteriaContract) {
	b.BACnetEventParameterChangeOfValueCivCriteriaContract = contract
}

func (b *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder) WithMandatoryFields(referencedPropertyIncrement BACnetContextTagReal) BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder {
	return b.WithReferencedPropertyIncrement(referencedPropertyIncrement)
}

func (b *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder) WithReferencedPropertyIncrement(referencedPropertyIncrement BACnetContextTagReal) BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder {
	b.ReferencedPropertyIncrement = referencedPropertyIncrement
	return b
}

func (b *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder) WithReferencedPropertyIncrementBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder {
	builder := builderSupplier(b.ReferencedPropertyIncrement.CreateBACnetContextTagRealBuilder())
	var err error
	b.ReferencedPropertyIncrement, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder) Build() (BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement, error) {
	if b.ReferencedPropertyIncrement == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'referencedPropertyIncrement' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement.deepCopy(), nil
}

func (b *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder) MustBuild() BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder) Done() BACnetEventParameterChangeOfValueCivCriteriaBuilder {
	return b.parentBuilder
}

func (b *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder) buildForBACnetEventParameterChangeOfValueCivCriteria() (BACnetEventParameterChangeOfValueCivCriteria, error) {
	return b.Build()
}

func (b *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder().(*_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder creates a BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder
func (b *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) CreateBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder() BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder {
	if b == nil {
		return NewBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder()
	}
	return &_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilder{_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement: b.deepCopy()}
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

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetParent() BACnetEventParameterChangeOfValueCivCriteriaContract {
	return m.BACnetEventParameterChangeOfValueCivCriteriaContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetReferencedPropertyIncrement() BACnetContextTagReal {
	return m.ReferencedPropertyIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement(structType any) BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement {
	if casted, ok := structType.(BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetTypeName() string {
	return "BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterChangeOfValueCivCriteriaContract.(*_BACnetEventParameterChangeOfValueCivCriteria).getLengthInBits(ctx))

	// Simple field (referencedPropertyIncrement)
	lengthInBits += m.ReferencedPropertyIncrement.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameterChangeOfValueCivCriteria, tagNumber uint8) (__bACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement, err error) {
	m.BACnetEventParameterChangeOfValueCivCriteriaContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referencedPropertyIncrement, err := ReadSimpleField[BACnetContextTagReal](ctx, "referencedPropertyIncrement", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencedPropertyIncrement' field"))
	}
	m.ReferencedPropertyIncrement = referencedPropertyIncrement

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "referencedPropertyIncrement", m.GetReferencedPropertyIncrement(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referencedPropertyIncrement' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
		}
		return nil
	}
	return m.BACnetEventParameterChangeOfValueCivCriteriaContract.(*_BACnetEventParameterChangeOfValueCivCriteria).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) IsBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement() {
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) deepCopy() *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement {
	if m == nil {
		return nil
	}
	_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementCopy := &_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement{
		m.BACnetEventParameterChangeOfValueCivCriteriaContract.(*_BACnetEventParameterChangeOfValueCivCriteria).deepCopy(),
		utils.DeepCopy[BACnetContextTagReal](m.ReferencedPropertyIncrement),
	}
	m.BACnetEventParameterChangeOfValueCivCriteriaContract.(*_BACnetEventParameterChangeOfValueCivCriteria)._SubType = m
	return _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementCopy
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) String() string {
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

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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataBinaryInputAll is the corresponding interface of BACnetConstructedDataBinaryInputAll
type BACnetConstructedDataBinaryInputAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataBinaryInputAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBinaryInputAll()
	// CreateBuilder creates a BACnetConstructedDataBinaryInputAllBuilder
	CreateBACnetConstructedDataBinaryInputAllBuilder() BACnetConstructedDataBinaryInputAllBuilder
}

// _BACnetConstructedDataBinaryInputAll is the data-structure of this message
type _BACnetConstructedDataBinaryInputAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataBinaryInputAll = (*_BACnetConstructedDataBinaryInputAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBinaryInputAll)(nil)

// NewBACnetConstructedDataBinaryInputAll factory function for _BACnetConstructedDataBinaryInputAll
func NewBACnetConstructedDataBinaryInputAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryInputAll {
	_result := &_BACnetConstructedDataBinaryInputAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBinaryInputAllBuilder is a builder for BACnetConstructedDataBinaryInputAll
type BACnetConstructedDataBinaryInputAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataBinaryInputAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBinaryInputAll or returns an error if something is wrong
	Build() (BACnetConstructedDataBinaryInputAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBinaryInputAll
}

// NewBACnetConstructedDataBinaryInputAllBuilder() creates a BACnetConstructedDataBinaryInputAllBuilder
func NewBACnetConstructedDataBinaryInputAllBuilder() BACnetConstructedDataBinaryInputAllBuilder {
	return &_BACnetConstructedDataBinaryInputAllBuilder{_BACnetConstructedDataBinaryInputAll: new(_BACnetConstructedDataBinaryInputAll)}
}

type _BACnetConstructedDataBinaryInputAllBuilder struct {
	*_BACnetConstructedDataBinaryInputAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBinaryInputAllBuilder) = (*_BACnetConstructedDataBinaryInputAllBuilder)(nil)

func (b *_BACnetConstructedDataBinaryInputAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataBinaryInputAll
}

func (b *_BACnetConstructedDataBinaryInputAllBuilder) WithMandatoryFields() BACnetConstructedDataBinaryInputAllBuilder {
	return b
}

func (b *_BACnetConstructedDataBinaryInputAllBuilder) Build() (BACnetConstructedDataBinaryInputAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBinaryInputAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataBinaryInputAllBuilder) MustBuild() BACnetConstructedDataBinaryInputAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBinaryInputAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBinaryInputAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBinaryInputAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBinaryInputAllBuilder().(*_BACnetConstructedDataBinaryInputAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBinaryInputAllBuilder creates a BACnetConstructedDataBinaryInputAllBuilder
func (b *_BACnetConstructedDataBinaryInputAll) CreateBACnetConstructedDataBinaryInputAllBuilder() BACnetConstructedDataBinaryInputAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataBinaryInputAllBuilder()
	}
	return &_BACnetConstructedDataBinaryInputAllBuilder{_BACnetConstructedDataBinaryInputAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryInputAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_INPUT
}

func (m *_BACnetConstructedDataBinaryInputAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryInputAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryInputAll(structType any) BACnetConstructedDataBinaryInputAll {
	if casted, ok := structType.(BACnetConstructedDataBinaryInputAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryInputAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryInputAll) GetTypeName() string {
	return "BACnetConstructedDataBinaryInputAll"
}

func (m *_BACnetConstructedDataBinaryInputAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryInputAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBinaryInputAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBinaryInputAll BACnetConstructedDataBinaryInputAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryInputAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryInputAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryInputAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryInputAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBinaryInputAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryInputAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryInputAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryInputAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryInputAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryInputAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryInputAll) IsBACnetConstructedDataBinaryInputAll() {}

func (m *_BACnetConstructedDataBinaryInputAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBinaryInputAll) deepCopy() *_BACnetConstructedDataBinaryInputAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBinaryInputAllCopy := &_BACnetConstructedDataBinaryInputAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataBinaryInputAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBinaryInputAllCopy
}

func (m *_BACnetConstructedDataBinaryInputAll) String() string {
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

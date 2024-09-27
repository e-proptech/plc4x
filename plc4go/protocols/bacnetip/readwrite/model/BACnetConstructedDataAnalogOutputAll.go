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

// BACnetConstructedDataAnalogOutputAll is the corresponding interface of BACnetConstructedDataAnalogOutputAll
type BACnetConstructedDataAnalogOutputAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataAnalogOutputAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAnalogOutputAll()
	// CreateBuilder creates a BACnetConstructedDataAnalogOutputAllBuilder
	CreateBACnetConstructedDataAnalogOutputAllBuilder() BACnetConstructedDataAnalogOutputAllBuilder
}

// _BACnetConstructedDataAnalogOutputAll is the data-structure of this message
type _BACnetConstructedDataAnalogOutputAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataAnalogOutputAll = (*_BACnetConstructedDataAnalogOutputAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAnalogOutputAll)(nil)

// NewBACnetConstructedDataAnalogOutputAll factory function for _BACnetConstructedDataAnalogOutputAll
func NewBACnetConstructedDataAnalogOutputAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogOutputAll {
	_result := &_BACnetConstructedDataAnalogOutputAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAnalogOutputAllBuilder is a builder for BACnetConstructedDataAnalogOutputAll
type BACnetConstructedDataAnalogOutputAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataAnalogOutputAllBuilder
	// Build builds the BACnetConstructedDataAnalogOutputAll or returns an error if something is wrong
	Build() (BACnetConstructedDataAnalogOutputAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAnalogOutputAll
}

// NewBACnetConstructedDataAnalogOutputAllBuilder() creates a BACnetConstructedDataAnalogOutputAllBuilder
func NewBACnetConstructedDataAnalogOutputAllBuilder() BACnetConstructedDataAnalogOutputAllBuilder {
	return &_BACnetConstructedDataAnalogOutputAllBuilder{_BACnetConstructedDataAnalogOutputAll: new(_BACnetConstructedDataAnalogOutputAll)}
}

type _BACnetConstructedDataAnalogOutputAllBuilder struct {
	*_BACnetConstructedDataAnalogOutputAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAnalogOutputAllBuilder) = (*_BACnetConstructedDataAnalogOutputAllBuilder)(nil)

func (b *_BACnetConstructedDataAnalogOutputAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAnalogOutputAllBuilder) WithMandatoryFields() BACnetConstructedDataAnalogOutputAllBuilder {
	return b
}

func (b *_BACnetConstructedDataAnalogOutputAllBuilder) Build() (BACnetConstructedDataAnalogOutputAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAnalogOutputAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataAnalogOutputAllBuilder) MustBuild() BACnetConstructedDataAnalogOutputAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAnalogOutputAllBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAnalogOutputAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAnalogOutputAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAnalogOutputAllBuilder().(*_BACnetConstructedDataAnalogOutputAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAnalogOutputAllBuilder creates a BACnetConstructedDataAnalogOutputAllBuilder
func (b *_BACnetConstructedDataAnalogOutputAll) CreateBACnetConstructedDataAnalogOutputAllBuilder() BACnetConstructedDataAnalogOutputAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataAnalogOutputAllBuilder()
	}
	return &_BACnetConstructedDataAnalogOutputAllBuilder{_BACnetConstructedDataAnalogOutputAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogOutputAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_OUTPUT
}

func (m *_BACnetConstructedDataAnalogOutputAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogOutputAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogOutputAll(structType any) BACnetConstructedDataAnalogOutputAll {
	if casted, ok := structType.(BACnetConstructedDataAnalogOutputAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogOutputAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogOutputAll) GetTypeName() string {
	return "BACnetConstructedDataAnalogOutputAll"
}

func (m *_BACnetConstructedDataAnalogOutputAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogOutputAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAnalogOutputAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAnalogOutputAll BACnetConstructedDataAnalogOutputAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogOutputAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogOutputAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogOutputAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogOutputAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAnalogOutputAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogOutputAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogOutputAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogOutputAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogOutputAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogOutputAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogOutputAll) IsBACnetConstructedDataAnalogOutputAll() {}

func (m *_BACnetConstructedDataAnalogOutputAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAnalogOutputAll) deepCopy() *_BACnetConstructedDataAnalogOutputAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAnalogOutputAllCopy := &_BACnetConstructedDataAnalogOutputAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAnalogOutputAllCopy
}

func (m *_BACnetConstructedDataAnalogOutputAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// BACnetConstructedDataProgramAll is the corresponding interface of BACnetConstructedDataProgramAll
type BACnetConstructedDataProgramAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataProgramAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataProgramAll()
	// CreateBuilder creates a BACnetConstructedDataProgramAllBuilder
	CreateBACnetConstructedDataProgramAllBuilder() BACnetConstructedDataProgramAllBuilder
}

// _BACnetConstructedDataProgramAll is the data-structure of this message
type _BACnetConstructedDataProgramAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataProgramAll = (*_BACnetConstructedDataProgramAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataProgramAll)(nil)

// NewBACnetConstructedDataProgramAll factory function for _BACnetConstructedDataProgramAll
func NewBACnetConstructedDataProgramAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProgramAll {
	_result := &_BACnetConstructedDataProgramAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataProgramAllBuilder is a builder for BACnetConstructedDataProgramAll
type BACnetConstructedDataProgramAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataProgramAllBuilder
	// Build builds the BACnetConstructedDataProgramAll or returns an error if something is wrong
	Build() (BACnetConstructedDataProgramAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataProgramAll
}

// NewBACnetConstructedDataProgramAllBuilder() creates a BACnetConstructedDataProgramAllBuilder
func NewBACnetConstructedDataProgramAllBuilder() BACnetConstructedDataProgramAllBuilder {
	return &_BACnetConstructedDataProgramAllBuilder{_BACnetConstructedDataProgramAll: new(_BACnetConstructedDataProgramAll)}
}

type _BACnetConstructedDataProgramAllBuilder struct {
	*_BACnetConstructedDataProgramAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataProgramAllBuilder) = (*_BACnetConstructedDataProgramAllBuilder)(nil)

func (b *_BACnetConstructedDataProgramAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataProgramAllBuilder) WithMandatoryFields() BACnetConstructedDataProgramAllBuilder {
	return b
}

func (b *_BACnetConstructedDataProgramAllBuilder) Build() (BACnetConstructedDataProgramAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataProgramAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataProgramAllBuilder) MustBuild() BACnetConstructedDataProgramAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataProgramAllBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataProgramAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataProgramAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataProgramAllBuilder().(*_BACnetConstructedDataProgramAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataProgramAllBuilder creates a BACnetConstructedDataProgramAllBuilder
func (b *_BACnetConstructedDataProgramAll) CreateBACnetConstructedDataProgramAllBuilder() BACnetConstructedDataProgramAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataProgramAllBuilder()
	}
	return &_BACnetConstructedDataProgramAllBuilder{_BACnetConstructedDataProgramAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProgramAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_PROGRAM
}

func (m *_BACnetConstructedDataProgramAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProgramAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProgramAll(structType any) BACnetConstructedDataProgramAll {
	if casted, ok := structType.(BACnetConstructedDataProgramAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProgramAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProgramAll) GetTypeName() string {
	return "BACnetConstructedDataProgramAll"
}

func (m *_BACnetConstructedDataProgramAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataProgramAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataProgramAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataProgramAll BACnetConstructedDataProgramAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProgramAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProgramAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProgramAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProgramAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataProgramAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProgramAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProgramAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProgramAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProgramAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProgramAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProgramAll) IsBACnetConstructedDataProgramAll() {}

func (m *_BACnetConstructedDataProgramAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataProgramAll) deepCopy() *_BACnetConstructedDataProgramAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataProgramAllCopy := &_BACnetConstructedDataProgramAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataProgramAllCopy
}

func (m *_BACnetConstructedDataProgramAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

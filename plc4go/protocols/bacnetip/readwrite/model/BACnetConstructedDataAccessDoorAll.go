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

// BACnetConstructedDataAccessDoorAll is the corresponding interface of BACnetConstructedDataAccessDoorAll
type BACnetConstructedDataAccessDoorAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataAccessDoorAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessDoorAll()
	// CreateBuilder creates a BACnetConstructedDataAccessDoorAllBuilder
	CreateBACnetConstructedDataAccessDoorAllBuilder() BACnetConstructedDataAccessDoorAllBuilder
}

// _BACnetConstructedDataAccessDoorAll is the data-structure of this message
type _BACnetConstructedDataAccessDoorAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataAccessDoorAll = (*_BACnetConstructedDataAccessDoorAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessDoorAll)(nil)

// NewBACnetConstructedDataAccessDoorAll factory function for _BACnetConstructedDataAccessDoorAll
func NewBACnetConstructedDataAccessDoorAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessDoorAll {
	_result := &_BACnetConstructedDataAccessDoorAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccessDoorAllBuilder is a builder for BACnetConstructedDataAccessDoorAll
type BACnetConstructedDataAccessDoorAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataAccessDoorAllBuilder
	// Build builds the BACnetConstructedDataAccessDoorAll or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessDoorAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessDoorAll
}

// NewBACnetConstructedDataAccessDoorAllBuilder() creates a BACnetConstructedDataAccessDoorAllBuilder
func NewBACnetConstructedDataAccessDoorAllBuilder() BACnetConstructedDataAccessDoorAllBuilder {
	return &_BACnetConstructedDataAccessDoorAllBuilder{_BACnetConstructedDataAccessDoorAll: new(_BACnetConstructedDataAccessDoorAll)}
}

type _BACnetConstructedDataAccessDoorAllBuilder struct {
	*_BACnetConstructedDataAccessDoorAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessDoorAllBuilder) = (*_BACnetConstructedDataAccessDoorAllBuilder)(nil)

func (b *_BACnetConstructedDataAccessDoorAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAccessDoorAllBuilder) WithMandatoryFields() BACnetConstructedDataAccessDoorAllBuilder {
	return b
}

func (b *_BACnetConstructedDataAccessDoorAllBuilder) Build() (BACnetConstructedDataAccessDoorAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccessDoorAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccessDoorAllBuilder) MustBuild() BACnetConstructedDataAccessDoorAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAccessDoorAllBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccessDoorAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccessDoorAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccessDoorAllBuilder().(*_BACnetConstructedDataAccessDoorAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccessDoorAllBuilder creates a BACnetConstructedDataAccessDoorAllBuilder
func (b *_BACnetConstructedDataAccessDoorAll) CreateBACnetConstructedDataAccessDoorAllBuilder() BACnetConstructedDataAccessDoorAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccessDoorAllBuilder()
	}
	return &_BACnetConstructedDataAccessDoorAllBuilder{_BACnetConstructedDataAccessDoorAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_DOOR
}

func (m *_BACnetConstructedDataAccessDoorAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessDoorAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessDoorAll(structType any) BACnetConstructedDataAccessDoorAll {
	if casted, ok := structType.(BACnetConstructedDataAccessDoorAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessDoorAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessDoorAll) GetTypeName() string {
	return "BACnetConstructedDataAccessDoorAll"
}

func (m *_BACnetConstructedDataAccessDoorAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessDoorAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessDoorAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessDoorAll BACnetConstructedDataAccessDoorAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessDoorAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessDoorAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessDoorAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessDoorAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessDoorAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessDoorAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessDoorAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessDoorAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessDoorAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessDoorAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessDoorAll) IsBACnetConstructedDataAccessDoorAll() {}

func (m *_BACnetConstructedDataAccessDoorAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessDoorAll) deepCopy() *_BACnetConstructedDataAccessDoorAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessDoorAllCopy := &_BACnetConstructedDataAccessDoorAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessDoorAllCopy
}

func (m *_BACnetConstructedDataAccessDoorAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

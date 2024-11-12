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

// ChangeListRemoveError is the corresponding interface of ChangeListRemoveError
type ChangeListRemoveError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetFirstFailedElementNumber returns FirstFailedElementNumber (property field)
	GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger
	// IsChangeListRemoveError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsChangeListRemoveError()
	// CreateBuilder creates a ChangeListRemoveErrorBuilder
	CreateChangeListRemoveErrorBuilder() ChangeListRemoveErrorBuilder
}

// _ChangeListRemoveError is the data-structure of this message
type _ChangeListRemoveError struct {
	BACnetErrorContract
	ErrorType                ErrorEnclosed
	FirstFailedElementNumber BACnetContextTagUnsignedInteger
}

var _ ChangeListRemoveError = (*_ChangeListRemoveError)(nil)
var _ BACnetErrorRequirements = (*_ChangeListRemoveError)(nil)

// NewChangeListRemoveError factory function for _ChangeListRemoveError
func NewChangeListRemoveError(errorType ErrorEnclosed, firstFailedElementNumber BACnetContextTagUnsignedInteger) *_ChangeListRemoveError {
	if errorType == nil {
		panic("errorType of type ErrorEnclosed for ChangeListRemoveError must not be nil")
	}
	if firstFailedElementNumber == nil {
		panic("firstFailedElementNumber of type BACnetContextTagUnsignedInteger for ChangeListRemoveError must not be nil")
	}
	_result := &_ChangeListRemoveError{
		BACnetErrorContract:      NewBACnetError(),
		ErrorType:                errorType,
		FirstFailedElementNumber: firstFailedElementNumber,
	}
	_result.BACnetErrorContract.(*_BACnetError)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ChangeListRemoveErrorBuilder is a builder for ChangeListRemoveError
type ChangeListRemoveErrorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(errorType ErrorEnclosed, firstFailedElementNumber BACnetContextTagUnsignedInteger) ChangeListRemoveErrorBuilder
	// WithErrorType adds ErrorType (property field)
	WithErrorType(ErrorEnclosed) ChangeListRemoveErrorBuilder
	// WithErrorTypeBuilder adds ErrorType (property field) which is build by the builder
	WithErrorTypeBuilder(func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) ChangeListRemoveErrorBuilder
	// WithFirstFailedElementNumber adds FirstFailedElementNumber (property field)
	WithFirstFailedElementNumber(BACnetContextTagUnsignedInteger) ChangeListRemoveErrorBuilder
	// WithFirstFailedElementNumberBuilder adds FirstFailedElementNumber (property field) which is build by the builder
	WithFirstFailedElementNumberBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) ChangeListRemoveErrorBuilder
	// Build builds the ChangeListRemoveError or returns an error if something is wrong
	Build() (ChangeListRemoveError, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ChangeListRemoveError
}

// NewChangeListRemoveErrorBuilder() creates a ChangeListRemoveErrorBuilder
func NewChangeListRemoveErrorBuilder() ChangeListRemoveErrorBuilder {
	return &_ChangeListRemoveErrorBuilder{_ChangeListRemoveError: new(_ChangeListRemoveError)}
}

type _ChangeListRemoveErrorBuilder struct {
	*_ChangeListRemoveError

	parentBuilder *_BACnetErrorBuilder

	err *utils.MultiError
}

var _ (ChangeListRemoveErrorBuilder) = (*_ChangeListRemoveErrorBuilder)(nil)

func (b *_ChangeListRemoveErrorBuilder) setParent(contract BACnetErrorContract) {
	b.BACnetErrorContract = contract
}

func (b *_ChangeListRemoveErrorBuilder) WithMandatoryFields(errorType ErrorEnclosed, firstFailedElementNumber BACnetContextTagUnsignedInteger) ChangeListRemoveErrorBuilder {
	return b.WithErrorType(errorType).WithFirstFailedElementNumber(firstFailedElementNumber)
}

func (b *_ChangeListRemoveErrorBuilder) WithErrorType(errorType ErrorEnclosed) ChangeListRemoveErrorBuilder {
	b.ErrorType = errorType
	return b
}

func (b *_ChangeListRemoveErrorBuilder) WithErrorTypeBuilder(builderSupplier func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) ChangeListRemoveErrorBuilder {
	builder := builderSupplier(b.ErrorType.CreateErrorEnclosedBuilder())
	var err error
	b.ErrorType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ErrorEnclosedBuilder failed"))
	}
	return b
}

func (b *_ChangeListRemoveErrorBuilder) WithFirstFailedElementNumber(firstFailedElementNumber BACnetContextTagUnsignedInteger) ChangeListRemoveErrorBuilder {
	b.FirstFailedElementNumber = firstFailedElementNumber
	return b
}

func (b *_ChangeListRemoveErrorBuilder) WithFirstFailedElementNumberBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) ChangeListRemoveErrorBuilder {
	builder := builderSupplier(b.FirstFailedElementNumber.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.FirstFailedElementNumber, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_ChangeListRemoveErrorBuilder) Build() (ChangeListRemoveError, error) {
	if b.ErrorType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'errorType' not set"))
	}
	if b.FirstFailedElementNumber == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'firstFailedElementNumber' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ChangeListRemoveError.deepCopy(), nil
}

func (b *_ChangeListRemoveErrorBuilder) MustBuild() ChangeListRemoveError {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ChangeListRemoveErrorBuilder) Done() BACnetErrorBuilder {
	return b.parentBuilder
}

func (b *_ChangeListRemoveErrorBuilder) buildForBACnetError() (BACnetError, error) {
	return b.Build()
}

func (b *_ChangeListRemoveErrorBuilder) DeepCopy() any {
	_copy := b.CreateChangeListRemoveErrorBuilder().(*_ChangeListRemoveErrorBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateChangeListRemoveErrorBuilder creates a ChangeListRemoveErrorBuilder
func (b *_ChangeListRemoveError) CreateChangeListRemoveErrorBuilder() ChangeListRemoveErrorBuilder {
	if b == nil {
		return NewChangeListRemoveErrorBuilder()
	}
	return &_ChangeListRemoveErrorBuilder{_ChangeListRemoveError: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ChangeListRemoveError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ChangeListRemoveError) GetParent() BACnetErrorContract {
	return m.BACnetErrorContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ChangeListRemoveError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_ChangeListRemoveError) GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger {
	return m.FirstFailedElementNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastChangeListRemoveError(structType any) ChangeListRemoveError {
	if casted, ok := structType.(ChangeListRemoveError); ok {
		return casted
	}
	if casted, ok := structType.(*ChangeListRemoveError); ok {
		return *casted
	}
	return nil
}

func (m *_ChangeListRemoveError) GetTypeName() string {
	return "ChangeListRemoveError"
}

func (m *_ChangeListRemoveError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetErrorContract.(*_BACnetError).getLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (firstFailedElementNumber)
	lengthInBits += m.FirstFailedElementNumber.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ChangeListRemoveError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ChangeListRemoveError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetError, errorChoice BACnetConfirmedServiceChoice) (__changeListRemoveError ChangeListRemoveError, err error) {
	m.BACnetErrorContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ChangeListRemoveError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ChangeListRemoveError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}
	m.ErrorType = errorType

	firstFailedElementNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "firstFailedElementNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'firstFailedElementNumber' field"))
	}
	m.FirstFailedElementNumber = firstFailedElementNumber

	if closeErr := readBuffer.CloseContext("ChangeListRemoveError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ChangeListRemoveError")
	}

	return m, nil
}

func (m *_ChangeListRemoveError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ChangeListRemoveError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ChangeListRemoveError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ChangeListRemoveError")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "errorType", m.GetErrorType(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorType' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "firstFailedElementNumber", m.GetFirstFailedElementNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'firstFailedElementNumber' field")
		}

		if popErr := writeBuffer.PopContext("ChangeListRemoveError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ChangeListRemoveError")
		}
		return nil
	}
	return m.BACnetErrorContract.(*_BACnetError).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ChangeListRemoveError) IsChangeListRemoveError() {}

func (m *_ChangeListRemoveError) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ChangeListRemoveError) deepCopy() *_ChangeListRemoveError {
	if m == nil {
		return nil
	}
	_ChangeListRemoveErrorCopy := &_ChangeListRemoveError{
		m.BACnetErrorContract.(*_BACnetError).deepCopy(),
		utils.DeepCopy[ErrorEnclosed](m.ErrorType),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.FirstFailedElementNumber),
	}
	m.BACnetErrorContract.(*_BACnetError)._SubType = m
	return _ChangeListRemoveErrorCopy
}

func (m *_ChangeListRemoveError) String() string {
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

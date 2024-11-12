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

// BACnetErrorGeneral is the corresponding interface of BACnetErrorGeneral
type BACnetErrorGeneral interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetError
	// GetError returns Error (property field)
	GetError() Error
	// IsBACnetErrorGeneral is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetErrorGeneral()
	// CreateBuilder creates a BACnetErrorGeneralBuilder
	CreateBACnetErrorGeneralBuilder() BACnetErrorGeneralBuilder
}

// _BACnetErrorGeneral is the data-structure of this message
type _BACnetErrorGeneral struct {
	BACnetErrorContract
	Error Error
}

var _ BACnetErrorGeneral = (*_BACnetErrorGeneral)(nil)
var _ BACnetErrorRequirements = (*_BACnetErrorGeneral)(nil)

// NewBACnetErrorGeneral factory function for _BACnetErrorGeneral
func NewBACnetErrorGeneral(error Error) *_BACnetErrorGeneral {
	if error == nil {
		panic("error of type Error for BACnetErrorGeneral must not be nil")
	}
	_result := &_BACnetErrorGeneral{
		BACnetErrorContract: NewBACnetError(),
		Error:               error,
	}
	_result.BACnetErrorContract.(*_BACnetError)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetErrorGeneralBuilder is a builder for BACnetErrorGeneral
type BACnetErrorGeneralBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(error Error) BACnetErrorGeneralBuilder
	// WithError adds Error (property field)
	WithError(Error) BACnetErrorGeneralBuilder
	// WithErrorBuilder adds Error (property field) which is build by the builder
	WithErrorBuilder(func(ErrorBuilder) ErrorBuilder) BACnetErrorGeneralBuilder
	// Build builds the BACnetErrorGeneral or returns an error if something is wrong
	Build() (BACnetErrorGeneral, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetErrorGeneral
}

// NewBACnetErrorGeneralBuilder() creates a BACnetErrorGeneralBuilder
func NewBACnetErrorGeneralBuilder() BACnetErrorGeneralBuilder {
	return &_BACnetErrorGeneralBuilder{_BACnetErrorGeneral: new(_BACnetErrorGeneral)}
}

type _BACnetErrorGeneralBuilder struct {
	*_BACnetErrorGeneral

	parentBuilder *_BACnetErrorBuilder

	err *utils.MultiError
}

var _ (BACnetErrorGeneralBuilder) = (*_BACnetErrorGeneralBuilder)(nil)

func (b *_BACnetErrorGeneralBuilder) setParent(contract BACnetErrorContract) {
	b.BACnetErrorContract = contract
}

func (b *_BACnetErrorGeneralBuilder) WithMandatoryFields(error Error) BACnetErrorGeneralBuilder {
	return b.WithError(error)
}

func (b *_BACnetErrorGeneralBuilder) WithError(error Error) BACnetErrorGeneralBuilder {
	b.Error = error
	return b
}

func (b *_BACnetErrorGeneralBuilder) WithErrorBuilder(builderSupplier func(ErrorBuilder) ErrorBuilder) BACnetErrorGeneralBuilder {
	builder := builderSupplier(b.Error.CreateErrorBuilder())
	var err error
	b.Error, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ErrorBuilder failed"))
	}
	return b
}

func (b *_BACnetErrorGeneralBuilder) Build() (BACnetErrorGeneral, error) {
	if b.Error == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'error' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetErrorGeneral.deepCopy(), nil
}

func (b *_BACnetErrorGeneralBuilder) MustBuild() BACnetErrorGeneral {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetErrorGeneralBuilder) Done() BACnetErrorBuilder {
	return b.parentBuilder
}

func (b *_BACnetErrorGeneralBuilder) buildForBACnetError() (BACnetError, error) {
	return b.Build()
}

func (b *_BACnetErrorGeneralBuilder) DeepCopy() any {
	_copy := b.CreateBACnetErrorGeneralBuilder().(*_BACnetErrorGeneralBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetErrorGeneralBuilder creates a BACnetErrorGeneralBuilder
func (b *_BACnetErrorGeneral) CreateBACnetErrorGeneralBuilder() BACnetErrorGeneralBuilder {
	if b == nil {
		return NewBACnetErrorGeneralBuilder()
	}
	return &_BACnetErrorGeneralBuilder{_BACnetErrorGeneral: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetErrorGeneral) GetErrorChoice() BACnetConfirmedServiceChoice {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetErrorGeneral) GetParent() BACnetErrorContract {
	return m.BACnetErrorContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetErrorGeneral) GetError() Error {
	return m.Error
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetErrorGeneral(structType any) BACnetErrorGeneral {
	if casted, ok := structType.(BACnetErrorGeneral); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetErrorGeneral); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetErrorGeneral) GetTypeName() string {
	return "BACnetErrorGeneral"
}

func (m *_BACnetErrorGeneral) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetErrorContract.(*_BACnetError).getLengthInBits(ctx))

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetErrorGeneral) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetErrorGeneral) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetError, errorChoice BACnetConfirmedServiceChoice) (__bACnetErrorGeneral BACnetErrorGeneral, err error) {
	m.BACnetErrorContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetErrorGeneral"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetErrorGeneral")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	error, err := ReadSimpleField[Error](ctx, "error", ReadComplex[Error](ErrorParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'error' field"))
	}
	m.Error = error

	if closeErr := readBuffer.CloseContext("BACnetErrorGeneral"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetErrorGeneral")
	}

	return m, nil
}

func (m *_BACnetErrorGeneral) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetErrorGeneral) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetErrorGeneral"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetErrorGeneral")
		}

		if err := WriteSimpleField[Error](ctx, "error", m.GetError(), WriteComplex[Error](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'error' field")
		}

		if popErr := writeBuffer.PopContext("BACnetErrorGeneral"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetErrorGeneral")
		}
		return nil
	}
	return m.BACnetErrorContract.(*_BACnetError).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetErrorGeneral) IsBACnetErrorGeneral() {}

func (m *_BACnetErrorGeneral) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetErrorGeneral) deepCopy() *_BACnetErrorGeneral {
	if m == nil {
		return nil
	}
	_BACnetErrorGeneralCopy := &_BACnetErrorGeneral{
		m.BACnetErrorContract.(*_BACnetError).deepCopy(),
		utils.DeepCopy[Error](m.Error),
	}
	m.BACnetErrorContract.(*_BACnetError)._SubType = m
	return _BACnetErrorGeneralCopy
}

func (m *_BACnetErrorGeneral) String() string {
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

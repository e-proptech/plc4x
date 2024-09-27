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

// TransferResult is the corresponding interface of TransferResult
type TransferResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfAvailableSequenceNumbers returns NoOfAvailableSequenceNumbers (property field)
	GetNoOfAvailableSequenceNumbers() int32
	// GetAvailableSequenceNumbers returns AvailableSequenceNumbers (property field)
	GetAvailableSequenceNumbers() []uint32
	// IsTransferResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTransferResult()
	// CreateBuilder creates a TransferResultBuilder
	CreateTransferResultBuilder() TransferResultBuilder
}

// _TransferResult is the data-structure of this message
type _TransferResult struct {
	ExtensionObjectDefinitionContract
	StatusCode                   StatusCode
	NoOfAvailableSequenceNumbers int32
	AvailableSequenceNumbers     []uint32
}

var _ TransferResult = (*_TransferResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_TransferResult)(nil)

// NewTransferResult factory function for _TransferResult
func NewTransferResult(statusCode StatusCode, noOfAvailableSequenceNumbers int32, availableSequenceNumbers []uint32) *_TransferResult {
	if statusCode == nil {
		panic("statusCode of type StatusCode for TransferResult must not be nil")
	}
	_result := &_TransferResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		NoOfAvailableSequenceNumbers:      noOfAvailableSequenceNumbers,
		AvailableSequenceNumbers:          availableSequenceNumbers,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TransferResultBuilder is a builder for TransferResult
type TransferResultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusCode StatusCode, noOfAvailableSequenceNumbers int32, availableSequenceNumbers []uint32) TransferResultBuilder
	// WithStatusCode adds StatusCode (property field)
	WithStatusCode(StatusCode) TransferResultBuilder
	// WithStatusCodeBuilder adds StatusCode (property field) which is build by the builder
	WithStatusCodeBuilder(func(StatusCodeBuilder) StatusCodeBuilder) TransferResultBuilder
	// WithNoOfAvailableSequenceNumbers adds NoOfAvailableSequenceNumbers (property field)
	WithNoOfAvailableSequenceNumbers(int32) TransferResultBuilder
	// WithAvailableSequenceNumbers adds AvailableSequenceNumbers (property field)
	WithAvailableSequenceNumbers(...uint32) TransferResultBuilder
	// Build builds the TransferResult or returns an error if something is wrong
	Build() (TransferResult, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TransferResult
}

// NewTransferResultBuilder() creates a TransferResultBuilder
func NewTransferResultBuilder() TransferResultBuilder {
	return &_TransferResultBuilder{_TransferResult: new(_TransferResult)}
}

type _TransferResultBuilder struct {
	*_TransferResult

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (TransferResultBuilder) = (*_TransferResultBuilder)(nil)

func (b *_TransferResultBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_TransferResultBuilder) WithMandatoryFields(statusCode StatusCode, noOfAvailableSequenceNumbers int32, availableSequenceNumbers []uint32) TransferResultBuilder {
	return b.WithStatusCode(statusCode).WithNoOfAvailableSequenceNumbers(noOfAvailableSequenceNumbers).WithAvailableSequenceNumbers(availableSequenceNumbers...)
}

func (b *_TransferResultBuilder) WithStatusCode(statusCode StatusCode) TransferResultBuilder {
	b.StatusCode = statusCode
	return b
}

func (b *_TransferResultBuilder) WithStatusCodeBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) TransferResultBuilder {
	builder := builderSupplier(b.StatusCode.CreateStatusCodeBuilder())
	var err error
	b.StatusCode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StatusCodeBuilder failed"))
	}
	return b
}

func (b *_TransferResultBuilder) WithNoOfAvailableSequenceNumbers(noOfAvailableSequenceNumbers int32) TransferResultBuilder {
	b.NoOfAvailableSequenceNumbers = noOfAvailableSequenceNumbers
	return b
}

func (b *_TransferResultBuilder) WithAvailableSequenceNumbers(availableSequenceNumbers ...uint32) TransferResultBuilder {
	b.AvailableSequenceNumbers = availableSequenceNumbers
	return b
}

func (b *_TransferResultBuilder) Build() (TransferResult, error) {
	if b.StatusCode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusCode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TransferResult.deepCopy(), nil
}

func (b *_TransferResultBuilder) MustBuild() TransferResult {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_TransferResultBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_TransferResultBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_TransferResultBuilder) DeepCopy() any {
	_copy := b.CreateTransferResultBuilder().(*_TransferResultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTransferResultBuilder creates a TransferResultBuilder
func (b *_TransferResult) CreateTransferResultBuilder() TransferResultBuilder {
	if b == nil {
		return NewTransferResultBuilder()
	}
	return &_TransferResultBuilder{_TransferResult: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TransferResult) GetIdentifier() string {
	return "838"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TransferResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TransferResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_TransferResult) GetNoOfAvailableSequenceNumbers() int32 {
	return m.NoOfAvailableSequenceNumbers
}

func (m *_TransferResult) GetAvailableSequenceNumbers() []uint32 {
	return m.AvailableSequenceNumbers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTransferResult(structType any) TransferResult {
	if casted, ok := structType.(TransferResult); ok {
		return casted
	}
	if casted, ok := structType.(*TransferResult); ok {
		return *casted
	}
	return nil
}

func (m *_TransferResult) GetTypeName() string {
	return "TransferResult"
}

func (m *_TransferResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfAvailableSequenceNumbers)
	lengthInBits += 32

	// Array field
	if len(m.AvailableSequenceNumbers) > 0 {
		lengthInBits += 32 * uint16(len(m.AvailableSequenceNumbers))
	}

	return lengthInBits
}

func (m *_TransferResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TransferResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__transferResult TransferResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TransferResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TransferResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	noOfAvailableSequenceNumbers, err := ReadSimpleField(ctx, "noOfAvailableSequenceNumbers", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfAvailableSequenceNumbers' field"))
	}
	m.NoOfAvailableSequenceNumbers = noOfAvailableSequenceNumbers

	availableSequenceNumbers, err := ReadCountArrayField[uint32](ctx, "availableSequenceNumbers", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfAvailableSequenceNumbers))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'availableSequenceNumbers' field"))
	}
	m.AvailableSequenceNumbers = availableSequenceNumbers

	if closeErr := readBuffer.CloseContext("TransferResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TransferResult")
	}

	return m, nil
}

func (m *_TransferResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TransferResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TransferResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TransferResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfAvailableSequenceNumbers", m.GetNoOfAvailableSequenceNumbers(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfAvailableSequenceNumbers' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "availableSequenceNumbers", m.GetAvailableSequenceNumbers(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'availableSequenceNumbers' field")
		}

		if popErr := writeBuffer.PopContext("TransferResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TransferResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TransferResult) IsTransferResult() {}

func (m *_TransferResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TransferResult) deepCopy() *_TransferResult {
	if m == nil {
		return nil
	}
	_TransferResultCopy := &_TransferResult{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.StatusCode.DeepCopy().(StatusCode),
		m.NoOfAvailableSequenceNumbers,
		utils.DeepCopySlice[uint32, uint32](m.AvailableSequenceNumbers),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _TransferResultCopy
}

func (m *_TransferResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

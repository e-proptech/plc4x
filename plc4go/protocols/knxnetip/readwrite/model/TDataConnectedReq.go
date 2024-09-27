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

// TDataConnectedReq is the corresponding interface of TDataConnectedReq
type TDataConnectedReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsTDataConnectedReq is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTDataConnectedReq()
	// CreateBuilder creates a TDataConnectedReqBuilder
	CreateTDataConnectedReqBuilder() TDataConnectedReqBuilder
}

// _TDataConnectedReq is the data-structure of this message
type _TDataConnectedReq struct {
	CEMIContract
}

var _ TDataConnectedReq = (*_TDataConnectedReq)(nil)
var _ CEMIRequirements = (*_TDataConnectedReq)(nil)

// NewTDataConnectedReq factory function for _TDataConnectedReq
func NewTDataConnectedReq(size uint16) *_TDataConnectedReq {
	_result := &_TDataConnectedReq{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TDataConnectedReqBuilder is a builder for TDataConnectedReq
type TDataConnectedReqBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() TDataConnectedReqBuilder
	// Build builds the TDataConnectedReq or returns an error if something is wrong
	Build() (TDataConnectedReq, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TDataConnectedReq
}

// NewTDataConnectedReqBuilder() creates a TDataConnectedReqBuilder
func NewTDataConnectedReqBuilder() TDataConnectedReqBuilder {
	return &_TDataConnectedReqBuilder{_TDataConnectedReq: new(_TDataConnectedReq)}
}

type _TDataConnectedReqBuilder struct {
	*_TDataConnectedReq

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (TDataConnectedReqBuilder) = (*_TDataConnectedReqBuilder)(nil)

func (b *_TDataConnectedReqBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
}

func (b *_TDataConnectedReqBuilder) WithMandatoryFields() TDataConnectedReqBuilder {
	return b
}

func (b *_TDataConnectedReqBuilder) Build() (TDataConnectedReq, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TDataConnectedReq.deepCopy(), nil
}

func (b *_TDataConnectedReqBuilder) MustBuild() TDataConnectedReq {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_TDataConnectedReqBuilder) Done() CEMIBuilder {
	return b.parentBuilder
}

func (b *_TDataConnectedReqBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_TDataConnectedReqBuilder) DeepCopy() any {
	_copy := b.CreateTDataConnectedReqBuilder().(*_TDataConnectedReqBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTDataConnectedReqBuilder creates a TDataConnectedReqBuilder
func (b *_TDataConnectedReq) CreateTDataConnectedReqBuilder() TDataConnectedReqBuilder {
	if b == nil {
		return NewTDataConnectedReqBuilder()
	}
	return &_TDataConnectedReqBuilder{_TDataConnectedReq: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TDataConnectedReq) GetMessageCode() uint8 {
	return 0x41
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TDataConnectedReq) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastTDataConnectedReq(structType any) TDataConnectedReq {
	if casted, ok := structType.(TDataConnectedReq); ok {
		return casted
	}
	if casted, ok := structType.(*TDataConnectedReq); ok {
		return *casted
	}
	return nil
}

func (m *_TDataConnectedReq) GetTypeName() string {
	return "TDataConnectedReq"
}

func (m *_TDataConnectedReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_TDataConnectedReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TDataConnectedReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__tDataConnectedReq TDataConnectedReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TDataConnectedReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TDataConnectedReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TDataConnectedReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TDataConnectedReq")
	}

	return m, nil
}

func (m *_TDataConnectedReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TDataConnectedReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TDataConnectedReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TDataConnectedReq")
		}

		if popErr := writeBuffer.PopContext("TDataConnectedReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TDataConnectedReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TDataConnectedReq) IsTDataConnectedReq() {}

func (m *_TDataConnectedReq) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TDataConnectedReq) deepCopy() *_TDataConnectedReq {
	if m == nil {
		return nil
	}
	_TDataConnectedReqCopy := &_TDataConnectedReq{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _TDataConnectedReqCopy
}

func (m *_TDataConnectedReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

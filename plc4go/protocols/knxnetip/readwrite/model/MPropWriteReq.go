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

// MPropWriteReq is the corresponding interface of MPropWriteReq
type MPropWriteReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsMPropWriteReq is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMPropWriteReq()
	// CreateBuilder creates a MPropWriteReqBuilder
	CreateMPropWriteReqBuilder() MPropWriteReqBuilder
}

// _MPropWriteReq is the data-structure of this message
type _MPropWriteReq struct {
	CEMIContract
}

var _ MPropWriteReq = (*_MPropWriteReq)(nil)
var _ CEMIRequirements = (*_MPropWriteReq)(nil)

// NewMPropWriteReq factory function for _MPropWriteReq
func NewMPropWriteReq(size uint16) *_MPropWriteReq {
	_result := &_MPropWriteReq{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MPropWriteReqBuilder is a builder for MPropWriteReq
type MPropWriteReqBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MPropWriteReqBuilder
	// Build builds the MPropWriteReq or returns an error if something is wrong
	Build() (MPropWriteReq, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MPropWriteReq
}

// NewMPropWriteReqBuilder() creates a MPropWriteReqBuilder
func NewMPropWriteReqBuilder() MPropWriteReqBuilder {
	return &_MPropWriteReqBuilder{_MPropWriteReq: new(_MPropWriteReq)}
}

type _MPropWriteReqBuilder struct {
	*_MPropWriteReq

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (MPropWriteReqBuilder) = (*_MPropWriteReqBuilder)(nil)

func (b *_MPropWriteReqBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
}

func (b *_MPropWriteReqBuilder) WithMandatoryFields() MPropWriteReqBuilder {
	return b
}

func (b *_MPropWriteReqBuilder) Build() (MPropWriteReq, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MPropWriteReq.deepCopy(), nil
}

func (b *_MPropWriteReqBuilder) MustBuild() MPropWriteReq {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MPropWriteReqBuilder) Done() CEMIBuilder {
	return b.parentBuilder
}

func (b *_MPropWriteReqBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_MPropWriteReqBuilder) DeepCopy() any {
	_copy := b.CreateMPropWriteReqBuilder().(*_MPropWriteReqBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMPropWriteReqBuilder creates a MPropWriteReqBuilder
func (b *_MPropWriteReq) CreateMPropWriteReqBuilder() MPropWriteReqBuilder {
	if b == nil {
		return NewMPropWriteReqBuilder()
	}
	return &_MPropWriteReqBuilder{_MPropWriteReq: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropWriteReq) GetMessageCode() uint8 {
	return 0xF6
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropWriteReq) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastMPropWriteReq(structType any) MPropWriteReq {
	if casted, ok := structType.(MPropWriteReq); ok {
		return casted
	}
	if casted, ok := structType.(*MPropWriteReq); ok {
		return *casted
	}
	return nil
}

func (m *_MPropWriteReq) GetTypeName() string {
	return "MPropWriteReq"
}

func (m *_MPropWriteReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MPropWriteReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MPropWriteReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mPropWriteReq MPropWriteReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropWriteReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropWriteReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MPropWriteReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropWriteReq")
	}

	return m, nil
}

func (m *_MPropWriteReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropWriteReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropWriteReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropWriteReq")
		}

		if popErr := writeBuffer.PopContext("MPropWriteReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropWriteReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MPropWriteReq) IsMPropWriteReq() {}

func (m *_MPropWriteReq) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MPropWriteReq) deepCopy() *_MPropWriteReq {
	if m == nil {
		return nil
	}
	_MPropWriteReqCopy := &_MPropWriteReq{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _MPropWriteReqCopy
}

func (m *_MPropWriteReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

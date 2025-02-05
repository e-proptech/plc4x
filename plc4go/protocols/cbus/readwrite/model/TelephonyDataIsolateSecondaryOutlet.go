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

// TelephonyDataIsolateSecondaryOutlet is the corresponding interface of TelephonyDataIsolateSecondaryOutlet
type TelephonyDataIsolateSecondaryOutlet interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	TelephonyData
	// GetIsolateStatus returns IsolateStatus (property field)
	GetIsolateStatus() byte
	// GetIsBehaveNormal returns IsBehaveNormal (virtual field)
	GetIsBehaveNormal() bool
	// GetIsToBeIsolated returns IsToBeIsolated (virtual field)
	GetIsToBeIsolated() bool
	// IsTelephonyDataIsolateSecondaryOutlet is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTelephonyDataIsolateSecondaryOutlet()
	// CreateBuilder creates a TelephonyDataIsolateSecondaryOutletBuilder
	CreateTelephonyDataIsolateSecondaryOutletBuilder() TelephonyDataIsolateSecondaryOutletBuilder
}

// _TelephonyDataIsolateSecondaryOutlet is the data-structure of this message
type _TelephonyDataIsolateSecondaryOutlet struct {
	TelephonyDataContract
	IsolateStatus byte
}

var _ TelephonyDataIsolateSecondaryOutlet = (*_TelephonyDataIsolateSecondaryOutlet)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataIsolateSecondaryOutlet)(nil)

// NewTelephonyDataIsolateSecondaryOutlet factory function for _TelephonyDataIsolateSecondaryOutlet
func NewTelephonyDataIsolateSecondaryOutlet(commandTypeContainer TelephonyCommandTypeContainer, argument byte, isolateStatus byte) *_TelephonyDataIsolateSecondaryOutlet {
	_result := &_TelephonyDataIsolateSecondaryOutlet{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
		IsolateStatus:         isolateStatus,
	}
	_result.TelephonyDataContract.(*_TelephonyData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TelephonyDataIsolateSecondaryOutletBuilder is a builder for TelephonyDataIsolateSecondaryOutlet
type TelephonyDataIsolateSecondaryOutletBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(isolateStatus byte) TelephonyDataIsolateSecondaryOutletBuilder
	// WithIsolateStatus adds IsolateStatus (property field)
	WithIsolateStatus(byte) TelephonyDataIsolateSecondaryOutletBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() TelephonyDataBuilder
	// Build builds the TelephonyDataIsolateSecondaryOutlet or returns an error if something is wrong
	Build() (TelephonyDataIsolateSecondaryOutlet, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TelephonyDataIsolateSecondaryOutlet
}

// NewTelephonyDataIsolateSecondaryOutletBuilder() creates a TelephonyDataIsolateSecondaryOutletBuilder
func NewTelephonyDataIsolateSecondaryOutletBuilder() TelephonyDataIsolateSecondaryOutletBuilder {
	return &_TelephonyDataIsolateSecondaryOutletBuilder{_TelephonyDataIsolateSecondaryOutlet: new(_TelephonyDataIsolateSecondaryOutlet)}
}

type _TelephonyDataIsolateSecondaryOutletBuilder struct {
	*_TelephonyDataIsolateSecondaryOutlet

	parentBuilder *_TelephonyDataBuilder

	err *utils.MultiError
}

var _ (TelephonyDataIsolateSecondaryOutletBuilder) = (*_TelephonyDataIsolateSecondaryOutletBuilder)(nil)

func (b *_TelephonyDataIsolateSecondaryOutletBuilder) setParent(contract TelephonyDataContract) {
	b.TelephonyDataContract = contract
	contract.(*_TelephonyData)._SubType = b._TelephonyDataIsolateSecondaryOutlet
}

func (b *_TelephonyDataIsolateSecondaryOutletBuilder) WithMandatoryFields(isolateStatus byte) TelephonyDataIsolateSecondaryOutletBuilder {
	return b.WithIsolateStatus(isolateStatus)
}

func (b *_TelephonyDataIsolateSecondaryOutletBuilder) WithIsolateStatus(isolateStatus byte) TelephonyDataIsolateSecondaryOutletBuilder {
	b.IsolateStatus = isolateStatus
	return b
}

func (b *_TelephonyDataIsolateSecondaryOutletBuilder) Build() (TelephonyDataIsolateSecondaryOutlet, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TelephonyDataIsolateSecondaryOutlet.deepCopy(), nil
}

func (b *_TelephonyDataIsolateSecondaryOutletBuilder) MustBuild() TelephonyDataIsolateSecondaryOutlet {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_TelephonyDataIsolateSecondaryOutletBuilder) Done() TelephonyDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewTelephonyDataBuilder().(*_TelephonyDataBuilder)
	}
	return b.parentBuilder
}

func (b *_TelephonyDataIsolateSecondaryOutletBuilder) buildForTelephonyData() (TelephonyData, error) {
	return b.Build()
}

func (b *_TelephonyDataIsolateSecondaryOutletBuilder) DeepCopy() any {
	_copy := b.CreateTelephonyDataIsolateSecondaryOutletBuilder().(*_TelephonyDataIsolateSecondaryOutletBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTelephonyDataIsolateSecondaryOutletBuilder creates a TelephonyDataIsolateSecondaryOutletBuilder
func (b *_TelephonyDataIsolateSecondaryOutlet) CreateTelephonyDataIsolateSecondaryOutletBuilder() TelephonyDataIsolateSecondaryOutletBuilder {
	if b == nil {
		return NewTelephonyDataIsolateSecondaryOutletBuilder()
	}
	return &_TelephonyDataIsolateSecondaryOutletBuilder{_TelephonyDataIsolateSecondaryOutlet: b.deepCopy()}
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

func (m *_TelephonyDataIsolateSecondaryOutlet) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataIsolateSecondaryOutlet) GetIsolateStatus() byte {
	return m.IsolateStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_TelephonyDataIsolateSecondaryOutlet) GetIsBehaveNormal() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetIsolateStatus()) == (0x00)))
}

func (m *_TelephonyDataIsolateSecondaryOutlet) GetIsToBeIsolated() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetIsolateStatus()) == (0x01)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTelephonyDataIsolateSecondaryOutlet(structType any) TelephonyDataIsolateSecondaryOutlet {
	if casted, ok := structType.(TelephonyDataIsolateSecondaryOutlet); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataIsolateSecondaryOutlet); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataIsolateSecondaryOutlet) GetTypeName() string {
	return "TelephonyDataIsolateSecondaryOutlet"
}

func (m *_TelephonyDataIsolateSecondaryOutlet) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	// Simple field (isolateStatus)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_TelephonyDataIsolateSecondaryOutlet) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataIsolateSecondaryOutlet) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData) (__telephonyDataIsolateSecondaryOutlet TelephonyDataIsolateSecondaryOutlet, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataIsolateSecondaryOutlet"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataIsolateSecondaryOutlet")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	isolateStatus, err := ReadSimpleField(ctx, "isolateStatus", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isolateStatus' field"))
	}
	m.IsolateStatus = isolateStatus

	isBehaveNormal, err := ReadVirtualField[bool](ctx, "isBehaveNormal", (*bool)(nil), bool((isolateStatus) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isBehaveNormal' field"))
	}
	_ = isBehaveNormal

	isToBeIsolated, err := ReadVirtualField[bool](ctx, "isToBeIsolated", (*bool)(nil), bool((isolateStatus) == (0x01)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isToBeIsolated' field"))
	}
	_ = isToBeIsolated

	if closeErr := readBuffer.CloseContext("TelephonyDataIsolateSecondaryOutlet"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataIsolateSecondaryOutlet")
	}

	return m, nil
}

func (m *_TelephonyDataIsolateSecondaryOutlet) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataIsolateSecondaryOutlet) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataIsolateSecondaryOutlet"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataIsolateSecondaryOutlet")
		}

		if err := WriteSimpleField[byte](ctx, "isolateStatus", m.GetIsolateStatus(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'isolateStatus' field")
		}
		// Virtual field
		isBehaveNormal := m.GetIsBehaveNormal()
		_ = isBehaveNormal
		if _isBehaveNormalErr := writeBuffer.WriteVirtual(ctx, "isBehaveNormal", m.GetIsBehaveNormal()); _isBehaveNormalErr != nil {
			return errors.Wrap(_isBehaveNormalErr, "Error serializing 'isBehaveNormal' field")
		}
		// Virtual field
		isToBeIsolated := m.GetIsToBeIsolated()
		_ = isToBeIsolated
		if _isToBeIsolatedErr := writeBuffer.WriteVirtual(ctx, "isToBeIsolated", m.GetIsToBeIsolated()); _isToBeIsolatedErr != nil {
			return errors.Wrap(_isToBeIsolatedErr, "Error serializing 'isToBeIsolated' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataIsolateSecondaryOutlet"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataIsolateSecondaryOutlet")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataIsolateSecondaryOutlet) IsTelephonyDataIsolateSecondaryOutlet() {}

func (m *_TelephonyDataIsolateSecondaryOutlet) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TelephonyDataIsolateSecondaryOutlet) deepCopy() *_TelephonyDataIsolateSecondaryOutlet {
	if m == nil {
		return nil
	}
	_TelephonyDataIsolateSecondaryOutletCopy := &_TelephonyDataIsolateSecondaryOutlet{
		m.TelephonyDataContract.(*_TelephonyData).deepCopy(),
		m.IsolateStatus,
	}
	_TelephonyDataIsolateSecondaryOutletCopy.TelephonyDataContract.(*_TelephonyData)._SubType = m
	return _TelephonyDataIsolateSecondaryOutletCopy
}

func (m *_TelephonyDataIsolateSecondaryOutlet) String() string {
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

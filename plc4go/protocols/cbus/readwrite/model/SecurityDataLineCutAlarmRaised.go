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

// SecurityDataLineCutAlarmRaised is the corresponding interface of SecurityDataLineCutAlarmRaised
type SecurityDataLineCutAlarmRaised interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataLineCutAlarmRaised is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataLineCutAlarmRaised()
	// CreateBuilder creates a SecurityDataLineCutAlarmRaisedBuilder
	CreateSecurityDataLineCutAlarmRaisedBuilder() SecurityDataLineCutAlarmRaisedBuilder
}

// _SecurityDataLineCutAlarmRaised is the data-structure of this message
type _SecurityDataLineCutAlarmRaised struct {
	SecurityDataContract
}

var _ SecurityDataLineCutAlarmRaised = (*_SecurityDataLineCutAlarmRaised)(nil)
var _ SecurityDataRequirements = (*_SecurityDataLineCutAlarmRaised)(nil)

// NewSecurityDataLineCutAlarmRaised factory function for _SecurityDataLineCutAlarmRaised
func NewSecurityDataLineCutAlarmRaised(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataLineCutAlarmRaised {
	_result := &_SecurityDataLineCutAlarmRaised{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataLineCutAlarmRaisedBuilder is a builder for SecurityDataLineCutAlarmRaised
type SecurityDataLineCutAlarmRaisedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataLineCutAlarmRaisedBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SecurityDataBuilder
	// Build builds the SecurityDataLineCutAlarmRaised or returns an error if something is wrong
	Build() (SecurityDataLineCutAlarmRaised, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataLineCutAlarmRaised
}

// NewSecurityDataLineCutAlarmRaisedBuilder() creates a SecurityDataLineCutAlarmRaisedBuilder
func NewSecurityDataLineCutAlarmRaisedBuilder() SecurityDataLineCutAlarmRaisedBuilder {
	return &_SecurityDataLineCutAlarmRaisedBuilder{_SecurityDataLineCutAlarmRaised: new(_SecurityDataLineCutAlarmRaised)}
}

type _SecurityDataLineCutAlarmRaisedBuilder struct {
	*_SecurityDataLineCutAlarmRaised

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataLineCutAlarmRaisedBuilder) = (*_SecurityDataLineCutAlarmRaisedBuilder)(nil)

func (b *_SecurityDataLineCutAlarmRaisedBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
	contract.(*_SecurityData)._SubType = b._SecurityDataLineCutAlarmRaised
}

func (b *_SecurityDataLineCutAlarmRaisedBuilder) WithMandatoryFields() SecurityDataLineCutAlarmRaisedBuilder {
	return b
}

func (b *_SecurityDataLineCutAlarmRaisedBuilder) Build() (SecurityDataLineCutAlarmRaised, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataLineCutAlarmRaised.deepCopy(), nil
}

func (b *_SecurityDataLineCutAlarmRaisedBuilder) MustBuild() SecurityDataLineCutAlarmRaised {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SecurityDataLineCutAlarmRaisedBuilder) Done() SecurityDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSecurityDataBuilder().(*_SecurityDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SecurityDataLineCutAlarmRaisedBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataLineCutAlarmRaisedBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataLineCutAlarmRaisedBuilder().(*_SecurityDataLineCutAlarmRaisedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataLineCutAlarmRaisedBuilder creates a SecurityDataLineCutAlarmRaisedBuilder
func (b *_SecurityDataLineCutAlarmRaised) CreateSecurityDataLineCutAlarmRaisedBuilder() SecurityDataLineCutAlarmRaisedBuilder {
	if b == nil {
		return NewSecurityDataLineCutAlarmRaisedBuilder()
	}
	return &_SecurityDataLineCutAlarmRaisedBuilder{_SecurityDataLineCutAlarmRaised: b.deepCopy()}
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

func (m *_SecurityDataLineCutAlarmRaised) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataLineCutAlarmRaised(structType any) SecurityDataLineCutAlarmRaised {
	if casted, ok := structType.(SecurityDataLineCutAlarmRaised); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataLineCutAlarmRaised); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataLineCutAlarmRaised) GetTypeName() string {
	return "SecurityDataLineCutAlarmRaised"
}

func (m *_SecurityDataLineCutAlarmRaised) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataLineCutAlarmRaised) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataLineCutAlarmRaised) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataLineCutAlarmRaised SecurityDataLineCutAlarmRaised, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataLineCutAlarmRaised"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataLineCutAlarmRaised")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataLineCutAlarmRaised"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataLineCutAlarmRaised")
	}

	return m, nil
}

func (m *_SecurityDataLineCutAlarmRaised) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataLineCutAlarmRaised) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataLineCutAlarmRaised"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataLineCutAlarmRaised")
		}

		if popErr := writeBuffer.PopContext("SecurityDataLineCutAlarmRaised"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataLineCutAlarmRaised")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataLineCutAlarmRaised) IsSecurityDataLineCutAlarmRaised() {}

func (m *_SecurityDataLineCutAlarmRaised) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataLineCutAlarmRaised) deepCopy() *_SecurityDataLineCutAlarmRaised {
	if m == nil {
		return nil
	}
	_SecurityDataLineCutAlarmRaisedCopy := &_SecurityDataLineCutAlarmRaised{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	_SecurityDataLineCutAlarmRaisedCopy.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataLineCutAlarmRaisedCopy
}

func (m *_SecurityDataLineCutAlarmRaised) String() string {
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

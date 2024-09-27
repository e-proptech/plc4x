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

// SecurityDataEntryDelayStarted is the corresponding interface of SecurityDataEntryDelayStarted
type SecurityDataEntryDelayStarted interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataEntryDelayStarted is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataEntryDelayStarted()
	// CreateBuilder creates a SecurityDataEntryDelayStartedBuilder
	CreateSecurityDataEntryDelayStartedBuilder() SecurityDataEntryDelayStartedBuilder
}

// _SecurityDataEntryDelayStarted is the data-structure of this message
type _SecurityDataEntryDelayStarted struct {
	SecurityDataContract
}

var _ SecurityDataEntryDelayStarted = (*_SecurityDataEntryDelayStarted)(nil)
var _ SecurityDataRequirements = (*_SecurityDataEntryDelayStarted)(nil)

// NewSecurityDataEntryDelayStarted factory function for _SecurityDataEntryDelayStarted
func NewSecurityDataEntryDelayStarted(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataEntryDelayStarted {
	_result := &_SecurityDataEntryDelayStarted{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataEntryDelayStartedBuilder is a builder for SecurityDataEntryDelayStarted
type SecurityDataEntryDelayStartedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataEntryDelayStartedBuilder
	// Build builds the SecurityDataEntryDelayStarted or returns an error if something is wrong
	Build() (SecurityDataEntryDelayStarted, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataEntryDelayStarted
}

// NewSecurityDataEntryDelayStartedBuilder() creates a SecurityDataEntryDelayStartedBuilder
func NewSecurityDataEntryDelayStartedBuilder() SecurityDataEntryDelayStartedBuilder {
	return &_SecurityDataEntryDelayStartedBuilder{_SecurityDataEntryDelayStarted: new(_SecurityDataEntryDelayStarted)}
}

type _SecurityDataEntryDelayStartedBuilder struct {
	*_SecurityDataEntryDelayStarted

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataEntryDelayStartedBuilder) = (*_SecurityDataEntryDelayStartedBuilder)(nil)

func (b *_SecurityDataEntryDelayStartedBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
}

func (b *_SecurityDataEntryDelayStartedBuilder) WithMandatoryFields() SecurityDataEntryDelayStartedBuilder {
	return b
}

func (b *_SecurityDataEntryDelayStartedBuilder) Build() (SecurityDataEntryDelayStarted, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataEntryDelayStarted.deepCopy(), nil
}

func (b *_SecurityDataEntryDelayStartedBuilder) MustBuild() SecurityDataEntryDelayStarted {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SecurityDataEntryDelayStartedBuilder) Done() SecurityDataBuilder {
	return b.parentBuilder
}

func (b *_SecurityDataEntryDelayStartedBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataEntryDelayStartedBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataEntryDelayStartedBuilder().(*_SecurityDataEntryDelayStartedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataEntryDelayStartedBuilder creates a SecurityDataEntryDelayStartedBuilder
func (b *_SecurityDataEntryDelayStarted) CreateSecurityDataEntryDelayStartedBuilder() SecurityDataEntryDelayStartedBuilder {
	if b == nil {
		return NewSecurityDataEntryDelayStartedBuilder()
	}
	return &_SecurityDataEntryDelayStartedBuilder{_SecurityDataEntryDelayStarted: b.deepCopy()}
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

func (m *_SecurityDataEntryDelayStarted) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataEntryDelayStarted(structType any) SecurityDataEntryDelayStarted {
	if casted, ok := structType.(SecurityDataEntryDelayStarted); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataEntryDelayStarted); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataEntryDelayStarted) GetTypeName() string {
	return "SecurityDataEntryDelayStarted"
}

func (m *_SecurityDataEntryDelayStarted) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataEntryDelayStarted) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataEntryDelayStarted) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataEntryDelayStarted SecurityDataEntryDelayStarted, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataEntryDelayStarted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataEntryDelayStarted")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataEntryDelayStarted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataEntryDelayStarted")
	}

	return m, nil
}

func (m *_SecurityDataEntryDelayStarted) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataEntryDelayStarted) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataEntryDelayStarted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataEntryDelayStarted")
		}

		if popErr := writeBuffer.PopContext("SecurityDataEntryDelayStarted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataEntryDelayStarted")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataEntryDelayStarted) IsSecurityDataEntryDelayStarted() {}

func (m *_SecurityDataEntryDelayStarted) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataEntryDelayStarted) deepCopy() *_SecurityDataEntryDelayStarted {
	if m == nil {
		return nil
	}
	_SecurityDataEntryDelayStartedCopy := &_SecurityDataEntryDelayStarted{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataEntryDelayStartedCopy
}

func (m *_SecurityDataEntryDelayStarted) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

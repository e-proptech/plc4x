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

// SecurityDataLowBatteryDetected is the corresponding interface of SecurityDataLowBatteryDetected
type SecurityDataLowBatteryDetected interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataLowBatteryDetected is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataLowBatteryDetected()
	// CreateBuilder creates a SecurityDataLowBatteryDetectedBuilder
	CreateSecurityDataLowBatteryDetectedBuilder() SecurityDataLowBatteryDetectedBuilder
}

// _SecurityDataLowBatteryDetected is the data-structure of this message
type _SecurityDataLowBatteryDetected struct {
	SecurityDataContract
}

var _ SecurityDataLowBatteryDetected = (*_SecurityDataLowBatteryDetected)(nil)
var _ SecurityDataRequirements = (*_SecurityDataLowBatteryDetected)(nil)

// NewSecurityDataLowBatteryDetected factory function for _SecurityDataLowBatteryDetected
func NewSecurityDataLowBatteryDetected(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataLowBatteryDetected {
	_result := &_SecurityDataLowBatteryDetected{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataLowBatteryDetectedBuilder is a builder for SecurityDataLowBatteryDetected
type SecurityDataLowBatteryDetectedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataLowBatteryDetectedBuilder
	// Build builds the SecurityDataLowBatteryDetected or returns an error if something is wrong
	Build() (SecurityDataLowBatteryDetected, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataLowBatteryDetected
}

// NewSecurityDataLowBatteryDetectedBuilder() creates a SecurityDataLowBatteryDetectedBuilder
func NewSecurityDataLowBatteryDetectedBuilder() SecurityDataLowBatteryDetectedBuilder {
	return &_SecurityDataLowBatteryDetectedBuilder{_SecurityDataLowBatteryDetected: new(_SecurityDataLowBatteryDetected)}
}

type _SecurityDataLowBatteryDetectedBuilder struct {
	*_SecurityDataLowBatteryDetected

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataLowBatteryDetectedBuilder) = (*_SecurityDataLowBatteryDetectedBuilder)(nil)

func (b *_SecurityDataLowBatteryDetectedBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
}

func (b *_SecurityDataLowBatteryDetectedBuilder) WithMandatoryFields() SecurityDataLowBatteryDetectedBuilder {
	return b
}

func (b *_SecurityDataLowBatteryDetectedBuilder) Build() (SecurityDataLowBatteryDetected, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataLowBatteryDetected.deepCopy(), nil
}

func (b *_SecurityDataLowBatteryDetectedBuilder) MustBuild() SecurityDataLowBatteryDetected {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SecurityDataLowBatteryDetectedBuilder) Done() SecurityDataBuilder {
	return b.parentBuilder
}

func (b *_SecurityDataLowBatteryDetectedBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataLowBatteryDetectedBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataLowBatteryDetectedBuilder().(*_SecurityDataLowBatteryDetectedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataLowBatteryDetectedBuilder creates a SecurityDataLowBatteryDetectedBuilder
func (b *_SecurityDataLowBatteryDetected) CreateSecurityDataLowBatteryDetectedBuilder() SecurityDataLowBatteryDetectedBuilder {
	if b == nil {
		return NewSecurityDataLowBatteryDetectedBuilder()
	}
	return &_SecurityDataLowBatteryDetectedBuilder{_SecurityDataLowBatteryDetected: b.deepCopy()}
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

func (m *_SecurityDataLowBatteryDetected) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataLowBatteryDetected(structType any) SecurityDataLowBatteryDetected {
	if casted, ok := structType.(SecurityDataLowBatteryDetected); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataLowBatteryDetected); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataLowBatteryDetected) GetTypeName() string {
	return "SecurityDataLowBatteryDetected"
}

func (m *_SecurityDataLowBatteryDetected) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataLowBatteryDetected) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataLowBatteryDetected) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataLowBatteryDetected SecurityDataLowBatteryDetected, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataLowBatteryDetected"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataLowBatteryDetected")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataLowBatteryDetected"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataLowBatteryDetected")
	}

	return m, nil
}

func (m *_SecurityDataLowBatteryDetected) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataLowBatteryDetected) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataLowBatteryDetected"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataLowBatteryDetected")
		}

		if popErr := writeBuffer.PopContext("SecurityDataLowBatteryDetected"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataLowBatteryDetected")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataLowBatteryDetected) IsSecurityDataLowBatteryDetected() {}

func (m *_SecurityDataLowBatteryDetected) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataLowBatteryDetected) deepCopy() *_SecurityDataLowBatteryDetected {
	if m == nil {
		return nil
	}
	_SecurityDataLowBatteryDetectedCopy := &_SecurityDataLowBatteryDetected{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataLowBatteryDetectedCopy
}

func (m *_SecurityDataLowBatteryDetected) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

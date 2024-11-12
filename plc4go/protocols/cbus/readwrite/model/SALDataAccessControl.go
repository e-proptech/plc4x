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

// SALDataAccessControl is the corresponding interface of SALDataAccessControl
type SALDataAccessControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetAccessControlData returns AccessControlData (property field)
	GetAccessControlData() AccessControlData
	// IsSALDataAccessControl is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataAccessControl()
	// CreateBuilder creates a SALDataAccessControlBuilder
	CreateSALDataAccessControlBuilder() SALDataAccessControlBuilder
}

// _SALDataAccessControl is the data-structure of this message
type _SALDataAccessControl struct {
	SALDataContract
	AccessControlData AccessControlData
}

var _ SALDataAccessControl = (*_SALDataAccessControl)(nil)
var _ SALDataRequirements = (*_SALDataAccessControl)(nil)

// NewSALDataAccessControl factory function for _SALDataAccessControl
func NewSALDataAccessControl(salData SALData, accessControlData AccessControlData) *_SALDataAccessControl {
	if accessControlData == nil {
		panic("accessControlData of type AccessControlData for SALDataAccessControl must not be nil")
	}
	_result := &_SALDataAccessControl{
		SALDataContract:   NewSALData(salData),
		AccessControlData: accessControlData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataAccessControlBuilder is a builder for SALDataAccessControl
type SALDataAccessControlBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(accessControlData AccessControlData) SALDataAccessControlBuilder
	// WithAccessControlData adds AccessControlData (property field)
	WithAccessControlData(AccessControlData) SALDataAccessControlBuilder
	// WithAccessControlDataBuilder adds AccessControlData (property field) which is build by the builder
	WithAccessControlDataBuilder(func(AccessControlDataBuilder) AccessControlDataBuilder) SALDataAccessControlBuilder
	// Build builds the SALDataAccessControl or returns an error if something is wrong
	Build() (SALDataAccessControl, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataAccessControl
}

// NewSALDataAccessControlBuilder() creates a SALDataAccessControlBuilder
func NewSALDataAccessControlBuilder() SALDataAccessControlBuilder {
	return &_SALDataAccessControlBuilder{_SALDataAccessControl: new(_SALDataAccessControl)}
}

type _SALDataAccessControlBuilder struct {
	*_SALDataAccessControl

	parentBuilder *_SALDataBuilder

	err *utils.MultiError
}

var _ (SALDataAccessControlBuilder) = (*_SALDataAccessControlBuilder)(nil)

func (b *_SALDataAccessControlBuilder) setParent(contract SALDataContract) {
	b.SALDataContract = contract
}

func (b *_SALDataAccessControlBuilder) WithMandatoryFields(accessControlData AccessControlData) SALDataAccessControlBuilder {
	return b.WithAccessControlData(accessControlData)
}

func (b *_SALDataAccessControlBuilder) WithAccessControlData(accessControlData AccessControlData) SALDataAccessControlBuilder {
	b.AccessControlData = accessControlData
	return b
}

func (b *_SALDataAccessControlBuilder) WithAccessControlDataBuilder(builderSupplier func(AccessControlDataBuilder) AccessControlDataBuilder) SALDataAccessControlBuilder {
	builder := builderSupplier(b.AccessControlData.CreateAccessControlDataBuilder())
	var err error
	b.AccessControlData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "AccessControlDataBuilder failed"))
	}
	return b
}

func (b *_SALDataAccessControlBuilder) Build() (SALDataAccessControl, error) {
	if b.AccessControlData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'accessControlData' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SALDataAccessControl.deepCopy(), nil
}

func (b *_SALDataAccessControlBuilder) MustBuild() SALDataAccessControl {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SALDataAccessControlBuilder) Done() SALDataBuilder {
	return b.parentBuilder
}

func (b *_SALDataAccessControlBuilder) buildForSALData() (SALData, error) {
	return b.Build()
}

func (b *_SALDataAccessControlBuilder) DeepCopy() any {
	_copy := b.CreateSALDataAccessControlBuilder().(*_SALDataAccessControlBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSALDataAccessControlBuilder creates a SALDataAccessControlBuilder
func (b *_SALDataAccessControl) CreateSALDataAccessControlBuilder() SALDataAccessControlBuilder {
	if b == nil {
		return NewSALDataAccessControlBuilder()
	}
	return &_SALDataAccessControlBuilder{_SALDataAccessControl: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataAccessControl) GetApplicationId() ApplicationId {
	return ApplicationId_ACCESS_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataAccessControl) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataAccessControl) GetAccessControlData() AccessControlData {
	return m.AccessControlData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataAccessControl(structType any) SALDataAccessControl {
	if casted, ok := structType.(SALDataAccessControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataAccessControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataAccessControl) GetTypeName() string {
	return "SALDataAccessControl"
}

func (m *_SALDataAccessControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (accessControlData)
	lengthInBits += m.AccessControlData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataAccessControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataAccessControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataAccessControl SALDataAccessControl, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataAccessControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataAccessControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessControlData, err := ReadSimpleField[AccessControlData](ctx, "accessControlData", ReadComplex[AccessControlData](AccessControlDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessControlData' field"))
	}
	m.AccessControlData = accessControlData

	if closeErr := readBuffer.CloseContext("SALDataAccessControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataAccessControl")
	}

	return m, nil
}

func (m *_SALDataAccessControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataAccessControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataAccessControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataAccessControl")
		}

		if err := WriteSimpleField[AccessControlData](ctx, "accessControlData", m.GetAccessControlData(), WriteComplex[AccessControlData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessControlData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataAccessControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataAccessControl")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataAccessControl) IsSALDataAccessControl() {}

func (m *_SALDataAccessControl) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataAccessControl) deepCopy() *_SALDataAccessControl {
	if m == nil {
		return nil
	}
	_SALDataAccessControlCopy := &_SALDataAccessControl{
		m.SALDataContract.(*_SALData).deepCopy(),
		utils.DeepCopy[AccessControlData](m.AccessControlData),
	}
	m.SALDataContract.(*_SALData)._SubType = m
	return _SALDataAccessControlCopy
}

func (m *_SALDataAccessControl) String() string {
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

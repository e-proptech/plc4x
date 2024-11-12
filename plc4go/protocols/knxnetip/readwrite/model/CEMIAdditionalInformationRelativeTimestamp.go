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

// Constant values.
const CEMIAdditionalInformationRelativeTimestamp_LEN uint8 = uint8(2)

// CEMIAdditionalInformationRelativeTimestamp is the corresponding interface of CEMIAdditionalInformationRelativeTimestamp
type CEMIAdditionalInformationRelativeTimestamp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMIAdditionalInformation
	// GetRelativeTimestamp returns RelativeTimestamp (property field)
	GetRelativeTimestamp() RelativeTimestamp
	// IsCEMIAdditionalInformationRelativeTimestamp is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCEMIAdditionalInformationRelativeTimestamp()
	// CreateBuilder creates a CEMIAdditionalInformationRelativeTimestampBuilder
	CreateCEMIAdditionalInformationRelativeTimestampBuilder() CEMIAdditionalInformationRelativeTimestampBuilder
}

// _CEMIAdditionalInformationRelativeTimestamp is the data-structure of this message
type _CEMIAdditionalInformationRelativeTimestamp struct {
	CEMIAdditionalInformationContract
	RelativeTimestamp RelativeTimestamp
}

var _ CEMIAdditionalInformationRelativeTimestamp = (*_CEMIAdditionalInformationRelativeTimestamp)(nil)
var _ CEMIAdditionalInformationRequirements = (*_CEMIAdditionalInformationRelativeTimestamp)(nil)

// NewCEMIAdditionalInformationRelativeTimestamp factory function for _CEMIAdditionalInformationRelativeTimestamp
func NewCEMIAdditionalInformationRelativeTimestamp(relativeTimestamp RelativeTimestamp) *_CEMIAdditionalInformationRelativeTimestamp {
	if relativeTimestamp == nil {
		panic("relativeTimestamp of type RelativeTimestamp for CEMIAdditionalInformationRelativeTimestamp must not be nil")
	}
	_result := &_CEMIAdditionalInformationRelativeTimestamp{
		CEMIAdditionalInformationContract: NewCEMIAdditionalInformation(),
		RelativeTimestamp:                 relativeTimestamp,
	}
	_result.CEMIAdditionalInformationContract.(*_CEMIAdditionalInformation)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CEMIAdditionalInformationRelativeTimestampBuilder is a builder for CEMIAdditionalInformationRelativeTimestamp
type CEMIAdditionalInformationRelativeTimestampBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(relativeTimestamp RelativeTimestamp) CEMIAdditionalInformationRelativeTimestampBuilder
	// WithRelativeTimestamp adds RelativeTimestamp (property field)
	WithRelativeTimestamp(RelativeTimestamp) CEMIAdditionalInformationRelativeTimestampBuilder
	// WithRelativeTimestampBuilder adds RelativeTimestamp (property field) which is build by the builder
	WithRelativeTimestampBuilder(func(RelativeTimestampBuilder) RelativeTimestampBuilder) CEMIAdditionalInformationRelativeTimestampBuilder
	// Build builds the CEMIAdditionalInformationRelativeTimestamp or returns an error if something is wrong
	Build() (CEMIAdditionalInformationRelativeTimestamp, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CEMIAdditionalInformationRelativeTimestamp
}

// NewCEMIAdditionalInformationRelativeTimestampBuilder() creates a CEMIAdditionalInformationRelativeTimestampBuilder
func NewCEMIAdditionalInformationRelativeTimestampBuilder() CEMIAdditionalInformationRelativeTimestampBuilder {
	return &_CEMIAdditionalInformationRelativeTimestampBuilder{_CEMIAdditionalInformationRelativeTimestamp: new(_CEMIAdditionalInformationRelativeTimestamp)}
}

type _CEMIAdditionalInformationRelativeTimestampBuilder struct {
	*_CEMIAdditionalInformationRelativeTimestamp

	parentBuilder *_CEMIAdditionalInformationBuilder

	err *utils.MultiError
}

var _ (CEMIAdditionalInformationRelativeTimestampBuilder) = (*_CEMIAdditionalInformationRelativeTimestampBuilder)(nil)

func (b *_CEMIAdditionalInformationRelativeTimestampBuilder) setParent(contract CEMIAdditionalInformationContract) {
	b.CEMIAdditionalInformationContract = contract
}

func (b *_CEMIAdditionalInformationRelativeTimestampBuilder) WithMandatoryFields(relativeTimestamp RelativeTimestamp) CEMIAdditionalInformationRelativeTimestampBuilder {
	return b.WithRelativeTimestamp(relativeTimestamp)
}

func (b *_CEMIAdditionalInformationRelativeTimestampBuilder) WithRelativeTimestamp(relativeTimestamp RelativeTimestamp) CEMIAdditionalInformationRelativeTimestampBuilder {
	b.RelativeTimestamp = relativeTimestamp
	return b
}

func (b *_CEMIAdditionalInformationRelativeTimestampBuilder) WithRelativeTimestampBuilder(builderSupplier func(RelativeTimestampBuilder) RelativeTimestampBuilder) CEMIAdditionalInformationRelativeTimestampBuilder {
	builder := builderSupplier(b.RelativeTimestamp.CreateRelativeTimestampBuilder())
	var err error
	b.RelativeTimestamp, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RelativeTimestampBuilder failed"))
	}
	return b
}

func (b *_CEMIAdditionalInformationRelativeTimestampBuilder) Build() (CEMIAdditionalInformationRelativeTimestamp, error) {
	if b.RelativeTimestamp == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'relativeTimestamp' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CEMIAdditionalInformationRelativeTimestamp.deepCopy(), nil
}

func (b *_CEMIAdditionalInformationRelativeTimestampBuilder) MustBuild() CEMIAdditionalInformationRelativeTimestamp {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CEMIAdditionalInformationRelativeTimestampBuilder) Done() CEMIAdditionalInformationBuilder {
	return b.parentBuilder
}

func (b *_CEMIAdditionalInformationRelativeTimestampBuilder) buildForCEMIAdditionalInformation() (CEMIAdditionalInformation, error) {
	return b.Build()
}

func (b *_CEMIAdditionalInformationRelativeTimestampBuilder) DeepCopy() any {
	_copy := b.CreateCEMIAdditionalInformationRelativeTimestampBuilder().(*_CEMIAdditionalInformationRelativeTimestampBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCEMIAdditionalInformationRelativeTimestampBuilder creates a CEMIAdditionalInformationRelativeTimestampBuilder
func (b *_CEMIAdditionalInformationRelativeTimestamp) CreateCEMIAdditionalInformationRelativeTimestampBuilder() CEMIAdditionalInformationRelativeTimestampBuilder {
	if b == nil {
		return NewCEMIAdditionalInformationRelativeTimestampBuilder()
	}
	return &_CEMIAdditionalInformationRelativeTimestampBuilder{_CEMIAdditionalInformationRelativeTimestamp: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetAdditionalInformationType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetParent() CEMIAdditionalInformationContract {
	return m.CEMIAdditionalInformationContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetRelativeTimestamp() RelativeTimestamp {
	return m.RelativeTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetLen() uint8 {
	return CEMIAdditionalInformationRelativeTimestamp_LEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCEMIAdditionalInformationRelativeTimestamp(structType any) CEMIAdditionalInformationRelativeTimestamp {
	if casted, ok := structType.(CEMIAdditionalInformationRelativeTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(*CEMIAdditionalInformationRelativeTimestamp); ok {
		return *casted
	}
	return nil
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetTypeName() string {
	return "CEMIAdditionalInformationRelativeTimestamp"
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIAdditionalInformationContract.(*_CEMIAdditionalInformation).getLengthInBits(ctx))

	// Const Field (len)
	lengthInBits += 8

	// Simple field (relativeTimestamp)
	lengthInBits += m.RelativeTimestamp.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMIAdditionalInformation) (__cEMIAdditionalInformationRelativeTimestamp CEMIAdditionalInformationRelativeTimestamp, err error) {
	m.CEMIAdditionalInformationContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CEMIAdditionalInformationRelativeTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CEMIAdditionalInformationRelativeTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	len, err := ReadConstField[uint8](ctx, "len", ReadUnsignedByte(readBuffer, uint8(8)), CEMIAdditionalInformationRelativeTimestamp_LEN)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'len' field"))
	}
	_ = len

	relativeTimestamp, err := ReadSimpleField[RelativeTimestamp](ctx, "relativeTimestamp", ReadComplex[RelativeTimestamp](RelativeTimestampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relativeTimestamp' field"))
	}
	m.RelativeTimestamp = relativeTimestamp

	if closeErr := readBuffer.CloseContext("CEMIAdditionalInformationRelativeTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CEMIAdditionalInformationRelativeTimestamp")
	}

	return m, nil
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CEMIAdditionalInformationRelativeTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CEMIAdditionalInformationRelativeTimestamp")
		}

		if err := WriteConstField(ctx, "len", CEMIAdditionalInformationRelativeTimestamp_LEN, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'len' field")
		}

		if err := WriteSimpleField[RelativeTimestamp](ctx, "relativeTimestamp", m.GetRelativeTimestamp(), WriteComplex[RelativeTimestamp](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relativeTimestamp' field")
		}

		if popErr := writeBuffer.PopContext("CEMIAdditionalInformationRelativeTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CEMIAdditionalInformationRelativeTimestamp")
		}
		return nil
	}
	return m.CEMIAdditionalInformationContract.(*_CEMIAdditionalInformation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) IsCEMIAdditionalInformationRelativeTimestamp() {
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) deepCopy() *_CEMIAdditionalInformationRelativeTimestamp {
	if m == nil {
		return nil
	}
	_CEMIAdditionalInformationRelativeTimestampCopy := &_CEMIAdditionalInformationRelativeTimestamp{
		m.CEMIAdditionalInformationContract.(*_CEMIAdditionalInformation).deepCopy(),
		utils.DeepCopy[RelativeTimestamp](m.RelativeTimestamp),
	}
	m.CEMIAdditionalInformationContract.(*_CEMIAdditionalInformation)._SubType = m
	return _CEMIAdditionalInformationRelativeTimestampCopy
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) String() string {
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

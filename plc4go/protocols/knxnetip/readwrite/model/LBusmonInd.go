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

// LBusmonInd is the corresponding interface of LBusmonInd
type LBusmonInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// GetAdditionalInformationLength returns AdditionalInformationLength (property field)
	GetAdditionalInformationLength() uint8
	// GetAdditionalInformation returns AdditionalInformation (property field)
	GetAdditionalInformation() []CEMIAdditionalInformation
	// GetDataFrame returns DataFrame (property field)
	GetDataFrame() LDataFrame
	// GetCrc returns Crc (property field)
	GetCrc() *uint8
	// IsLBusmonInd is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLBusmonInd()
	// CreateBuilder creates a LBusmonIndBuilder
	CreateLBusmonIndBuilder() LBusmonIndBuilder
}

// _LBusmonInd is the data-structure of this message
type _LBusmonInd struct {
	CEMIContract
	AdditionalInformationLength uint8
	AdditionalInformation       []CEMIAdditionalInformation
	DataFrame                   LDataFrame
	Crc                         *uint8
}

var _ LBusmonInd = (*_LBusmonInd)(nil)
var _ CEMIRequirements = (*_LBusmonInd)(nil)

// NewLBusmonInd factory function for _LBusmonInd
func NewLBusmonInd(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame, crc *uint8, size uint16) *_LBusmonInd {
	if dataFrame == nil {
		panic("dataFrame of type LDataFrame for LBusmonInd must not be nil")
	}
	_result := &_LBusmonInd{
		CEMIContract:                NewCEMI(size),
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Crc:                         crc,
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LBusmonIndBuilder is a builder for LBusmonInd
type LBusmonIndBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame) LBusmonIndBuilder
	// WithAdditionalInformationLength adds AdditionalInformationLength (property field)
	WithAdditionalInformationLength(uint8) LBusmonIndBuilder
	// WithAdditionalInformation adds AdditionalInformation (property field)
	WithAdditionalInformation(...CEMIAdditionalInformation) LBusmonIndBuilder
	// WithDataFrame adds DataFrame (property field)
	WithDataFrame(LDataFrame) LBusmonIndBuilder
	// WithDataFrameBuilder adds DataFrame (property field) which is build by the builder
	WithDataFrameBuilder(func(LDataFrameBuilder) LDataFrameBuilder) LBusmonIndBuilder
	// WithCrc adds Crc (property field)
	WithOptionalCrc(uint8) LBusmonIndBuilder
	// Build builds the LBusmonInd or returns an error if something is wrong
	Build() (LBusmonInd, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LBusmonInd
}

// NewLBusmonIndBuilder() creates a LBusmonIndBuilder
func NewLBusmonIndBuilder() LBusmonIndBuilder {
	return &_LBusmonIndBuilder{_LBusmonInd: new(_LBusmonInd)}
}

type _LBusmonIndBuilder struct {
	*_LBusmonInd

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (LBusmonIndBuilder) = (*_LBusmonIndBuilder)(nil)

func (b *_LBusmonIndBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
}

func (b *_LBusmonIndBuilder) WithMandatoryFields(additionalInformationLength uint8, additionalInformation []CEMIAdditionalInformation, dataFrame LDataFrame) LBusmonIndBuilder {
	return b.WithAdditionalInformationLength(additionalInformationLength).WithAdditionalInformation(additionalInformation...).WithDataFrame(dataFrame)
}

func (b *_LBusmonIndBuilder) WithAdditionalInformationLength(additionalInformationLength uint8) LBusmonIndBuilder {
	b.AdditionalInformationLength = additionalInformationLength
	return b
}

func (b *_LBusmonIndBuilder) WithAdditionalInformation(additionalInformation ...CEMIAdditionalInformation) LBusmonIndBuilder {
	b.AdditionalInformation = additionalInformation
	return b
}

func (b *_LBusmonIndBuilder) WithDataFrame(dataFrame LDataFrame) LBusmonIndBuilder {
	b.DataFrame = dataFrame
	return b
}

func (b *_LBusmonIndBuilder) WithDataFrameBuilder(builderSupplier func(LDataFrameBuilder) LDataFrameBuilder) LBusmonIndBuilder {
	builder := builderSupplier(b.DataFrame.CreateLDataFrameBuilder())
	var err error
	b.DataFrame, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LDataFrameBuilder failed"))
	}
	return b
}

func (b *_LBusmonIndBuilder) WithOptionalCrc(crc uint8) LBusmonIndBuilder {
	b.Crc = &crc
	return b
}

func (b *_LBusmonIndBuilder) Build() (LBusmonInd, error) {
	if b.DataFrame == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dataFrame' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LBusmonInd.deepCopy(), nil
}

func (b *_LBusmonIndBuilder) MustBuild() LBusmonInd {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_LBusmonIndBuilder) Done() CEMIBuilder {
	return b.parentBuilder
}

func (b *_LBusmonIndBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_LBusmonIndBuilder) DeepCopy() any {
	_copy := b.CreateLBusmonIndBuilder().(*_LBusmonIndBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLBusmonIndBuilder creates a LBusmonIndBuilder
func (b *_LBusmonInd) CreateLBusmonIndBuilder() LBusmonIndBuilder {
	if b == nil {
		return NewLBusmonIndBuilder()
	}
	return &_LBusmonIndBuilder{_LBusmonInd: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LBusmonInd) GetMessageCode() uint8 {
	return 0x2B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LBusmonInd) GetParent() CEMIContract {
	return m.CEMIContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LBusmonInd) GetAdditionalInformationLength() uint8 {
	return m.AdditionalInformationLength
}

func (m *_LBusmonInd) GetAdditionalInformation() []CEMIAdditionalInformation {
	return m.AdditionalInformation
}

func (m *_LBusmonInd) GetDataFrame() LDataFrame {
	return m.DataFrame
}

func (m *_LBusmonInd) GetCrc() *uint8 {
	return m.Crc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLBusmonInd(structType any) LBusmonInd {
	if casted, ok := structType.(LBusmonInd); ok {
		return casted
	}
	if casted, ok := structType.(*LBusmonInd); ok {
		return *casted
	}
	return nil
}

func (m *_LBusmonInd) GetTypeName() string {
	return "LBusmonInd"
}

func (m *_LBusmonInd) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.GetLengthInBits(ctx)

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_LBusmonInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LBusmonInd) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__lBusmonInd LBusmonInd, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LBusmonInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LBusmonInd")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	additionalInformationLength, err := ReadSimpleField(ctx, "additionalInformationLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInformationLength' field"))
	}
	m.AdditionalInformationLength = additionalInformationLength

	additionalInformation, err := ReadLengthArrayField[CEMIAdditionalInformation](ctx, "additionalInformation", ReadComplex[CEMIAdditionalInformation](CEMIAdditionalInformationParseWithBuffer, readBuffer), int(additionalInformationLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalInformation' field"))
	}
	m.AdditionalInformation = additionalInformation

	dataFrame, err := ReadSimpleField[LDataFrame](ctx, "dataFrame", ReadComplex[LDataFrame](LDataFrameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataFrame' field"))
	}
	m.DataFrame = dataFrame

	var crc *uint8
	crc, err = ReadOptionalField[uint8](ctx, "crc", ReadUnsignedByte(readBuffer, uint8(8)), CastLDataFrame(dataFrame).GetNotAckFrame())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'crc' field"))
	}
	m.Crc = crc

	if closeErr := readBuffer.CloseContext("LBusmonInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LBusmonInd")
	}

	return m, nil
}

func (m *_LBusmonInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LBusmonInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LBusmonInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LBusmonInd")
		}

		if err := WriteSimpleField[uint8](ctx, "additionalInformationLength", m.GetAdditionalInformationLength(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalInformationLength' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "additionalInformation", m.GetAdditionalInformation(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalInformation' field")
		}

		if err := WriteSimpleField[LDataFrame](ctx, "dataFrame", m.GetDataFrame(), WriteComplex[LDataFrame](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataFrame' field")
		}

		if err := WriteOptionalField[uint8](ctx, "crc", m.GetCrc(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("LBusmonInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LBusmonInd")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LBusmonInd) IsLBusmonInd() {}

func (m *_LBusmonInd) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LBusmonInd) deepCopy() *_LBusmonInd {
	if m == nil {
		return nil
	}
	_LBusmonIndCopy := &_LBusmonInd{
		m.CEMIContract.(*_CEMI).deepCopy(),
		m.AdditionalInformationLength,
		utils.DeepCopySlice[CEMIAdditionalInformation, CEMIAdditionalInformation](m.AdditionalInformation),
		m.DataFrame.DeepCopy().(LDataFrame),
		utils.CopyPtr[uint8](m.Crc),
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _LBusmonIndCopy
}

func (m *_LBusmonInd) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

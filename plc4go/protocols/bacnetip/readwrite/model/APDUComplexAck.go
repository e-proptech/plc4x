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

// APDUComplexAck is the corresponding interface of APDUComplexAck
type APDUComplexAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	APDU
	// GetSegmentedMessage returns SegmentedMessage (property field)
	GetSegmentedMessage() bool
	// GetMoreFollows returns MoreFollows (property field)
	GetMoreFollows() bool
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() *uint8
	// GetProposedWindowSize returns ProposedWindowSize (property field)
	GetProposedWindowSize() *uint8
	// GetServiceAck returns ServiceAck (property field)
	GetServiceAck() BACnetServiceAck
	// GetSegmentServiceChoice returns SegmentServiceChoice (property field)
	GetSegmentServiceChoice() *BACnetConfirmedServiceChoice
	// GetSegment returns Segment (property field)
	GetSegment() []byte
	// GetApduHeaderReduction returns ApduHeaderReduction (virtual field)
	GetApduHeaderReduction() uint16
	// GetSegmentReduction returns SegmentReduction (virtual field)
	GetSegmentReduction() uint16
	// IsAPDUComplexAck is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUComplexAck()
	// CreateBuilder creates a APDUComplexAckBuilder
	CreateAPDUComplexAckBuilder() APDUComplexAckBuilder
}

// _APDUComplexAck is the data-structure of this message
type _APDUComplexAck struct {
	APDUContract
	SegmentedMessage     bool
	MoreFollows          bool
	OriginalInvokeId     uint8
	SequenceNumber       *uint8
	ProposedWindowSize   *uint8
	ServiceAck           BACnetServiceAck
	SegmentServiceChoice *BACnetConfirmedServiceChoice
	Segment              []byte
	// Reserved Fields
	reservedField0 *uint8
}

var _ APDUComplexAck = (*_APDUComplexAck)(nil)
var _ APDURequirements = (*_APDUComplexAck)(nil)

// NewAPDUComplexAck factory function for _APDUComplexAck
func NewAPDUComplexAck(segmentedMessage bool, moreFollows bool, originalInvokeId uint8, sequenceNumber *uint8, proposedWindowSize *uint8, serviceAck BACnetServiceAck, segmentServiceChoice *BACnetConfirmedServiceChoice, segment []byte, apduLength uint16) *_APDUComplexAck {
	_result := &_APDUComplexAck{
		APDUContract:         NewAPDU(apduLength),
		SegmentedMessage:     segmentedMessage,
		MoreFollows:          moreFollows,
		OriginalInvokeId:     originalInvokeId,
		SequenceNumber:       sequenceNumber,
		ProposedWindowSize:   proposedWindowSize,
		ServiceAck:           serviceAck,
		SegmentServiceChoice: segmentServiceChoice,
		Segment:              segment,
	}
	_result.APDUContract.(*_APDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// APDUComplexAckBuilder is a builder for APDUComplexAck
type APDUComplexAckBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(segmentedMessage bool, moreFollows bool, originalInvokeId uint8, segment []byte) APDUComplexAckBuilder
	// WithSegmentedMessage adds SegmentedMessage (property field)
	WithSegmentedMessage(bool) APDUComplexAckBuilder
	// WithMoreFollows adds MoreFollows (property field)
	WithMoreFollows(bool) APDUComplexAckBuilder
	// WithOriginalInvokeId adds OriginalInvokeId (property field)
	WithOriginalInvokeId(uint8) APDUComplexAckBuilder
	// WithSequenceNumber adds SequenceNumber (property field)
	WithOptionalSequenceNumber(uint8) APDUComplexAckBuilder
	// WithProposedWindowSize adds ProposedWindowSize (property field)
	WithOptionalProposedWindowSize(uint8) APDUComplexAckBuilder
	// WithServiceAck adds ServiceAck (property field)
	WithOptionalServiceAck(BACnetServiceAck) APDUComplexAckBuilder
	// WithOptionalServiceAckBuilder adds ServiceAck (property field) which is build by the builder
	WithOptionalServiceAckBuilder(func(BACnetServiceAckBuilder) BACnetServiceAckBuilder) APDUComplexAckBuilder
	// WithSegmentServiceChoice adds SegmentServiceChoice (property field)
	WithOptionalSegmentServiceChoice(BACnetConfirmedServiceChoice) APDUComplexAckBuilder
	// WithSegment adds Segment (property field)
	WithSegment(...byte) APDUComplexAckBuilder
	// Build builds the APDUComplexAck or returns an error if something is wrong
	Build() (APDUComplexAck, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() APDUComplexAck
}

// NewAPDUComplexAckBuilder() creates a APDUComplexAckBuilder
func NewAPDUComplexAckBuilder() APDUComplexAckBuilder {
	return &_APDUComplexAckBuilder{_APDUComplexAck: new(_APDUComplexAck)}
}

type _APDUComplexAckBuilder struct {
	*_APDUComplexAck

	parentBuilder *_APDUBuilder

	err *utils.MultiError
}

var _ (APDUComplexAckBuilder) = (*_APDUComplexAckBuilder)(nil)

func (b *_APDUComplexAckBuilder) setParent(contract APDUContract) {
	b.APDUContract = contract
}

func (b *_APDUComplexAckBuilder) WithMandatoryFields(segmentedMessage bool, moreFollows bool, originalInvokeId uint8, segment []byte) APDUComplexAckBuilder {
	return b.WithSegmentedMessage(segmentedMessage).WithMoreFollows(moreFollows).WithOriginalInvokeId(originalInvokeId).WithSegment(segment...)
}

func (b *_APDUComplexAckBuilder) WithSegmentedMessage(segmentedMessage bool) APDUComplexAckBuilder {
	b.SegmentedMessage = segmentedMessage
	return b
}

func (b *_APDUComplexAckBuilder) WithMoreFollows(moreFollows bool) APDUComplexAckBuilder {
	b.MoreFollows = moreFollows
	return b
}

func (b *_APDUComplexAckBuilder) WithOriginalInvokeId(originalInvokeId uint8) APDUComplexAckBuilder {
	b.OriginalInvokeId = originalInvokeId
	return b
}

func (b *_APDUComplexAckBuilder) WithOptionalSequenceNumber(sequenceNumber uint8) APDUComplexAckBuilder {
	b.SequenceNumber = &sequenceNumber
	return b
}

func (b *_APDUComplexAckBuilder) WithOptionalProposedWindowSize(proposedWindowSize uint8) APDUComplexAckBuilder {
	b.ProposedWindowSize = &proposedWindowSize
	return b
}

func (b *_APDUComplexAckBuilder) WithOptionalServiceAck(serviceAck BACnetServiceAck) APDUComplexAckBuilder {
	b.ServiceAck = serviceAck
	return b
}

func (b *_APDUComplexAckBuilder) WithOptionalServiceAckBuilder(builderSupplier func(BACnetServiceAckBuilder) BACnetServiceAckBuilder) APDUComplexAckBuilder {
	builder := builderSupplier(b.ServiceAck.CreateBACnetServiceAckBuilder())
	var err error
	b.ServiceAck, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetServiceAckBuilder failed"))
	}
	return b
}

func (b *_APDUComplexAckBuilder) WithOptionalSegmentServiceChoice(segmentServiceChoice BACnetConfirmedServiceChoice) APDUComplexAckBuilder {
	b.SegmentServiceChoice = &segmentServiceChoice
	return b
}

func (b *_APDUComplexAckBuilder) WithSegment(segment ...byte) APDUComplexAckBuilder {
	b.Segment = segment
	return b
}

func (b *_APDUComplexAckBuilder) Build() (APDUComplexAck, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._APDUComplexAck.deepCopy(), nil
}

func (b *_APDUComplexAckBuilder) MustBuild() APDUComplexAck {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_APDUComplexAckBuilder) Done() APDUBuilder {
	return b.parentBuilder
}

func (b *_APDUComplexAckBuilder) buildForAPDU() (APDU, error) {
	return b.Build()
}

func (b *_APDUComplexAckBuilder) DeepCopy() any {
	_copy := b.CreateAPDUComplexAckBuilder().(*_APDUComplexAckBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAPDUComplexAckBuilder creates a APDUComplexAckBuilder
func (b *_APDUComplexAck) CreateAPDUComplexAckBuilder() APDUComplexAckBuilder {
	if b == nil {
		return NewAPDUComplexAckBuilder()
	}
	return &_APDUComplexAckBuilder{_APDUComplexAck: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUComplexAck) GetApduType() ApduType {
	return ApduType_COMPLEX_ACK_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUComplexAck) GetParent() APDUContract {
	return m.APDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUComplexAck) GetSegmentedMessage() bool {
	return m.SegmentedMessage
}

func (m *_APDUComplexAck) GetMoreFollows() bool {
	return m.MoreFollows
}

func (m *_APDUComplexAck) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUComplexAck) GetSequenceNumber() *uint8 {
	return m.SequenceNumber
}

func (m *_APDUComplexAck) GetProposedWindowSize() *uint8 {
	return m.ProposedWindowSize
}

func (m *_APDUComplexAck) GetServiceAck() BACnetServiceAck {
	return m.ServiceAck
}

func (m *_APDUComplexAck) GetSegmentServiceChoice() *BACnetConfirmedServiceChoice {
	return m.SegmentServiceChoice
}

func (m *_APDUComplexAck) GetSegment() []byte {
	return m.Segment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_APDUComplexAck) GetApduHeaderReduction() uint16 {
	ctx := context.Background()
	_ = ctx
	sequenceNumber := m.GetSequenceNumber()
	_ = sequenceNumber
	proposedWindowSize := m.GetProposedWindowSize()
	_ = proposedWindowSize
	serviceAck := m.GetServiceAck()
	_ = serviceAck
	segmentServiceChoice := m.GetSegmentServiceChoice()
	_ = segmentServiceChoice
	return uint16(uint16(uint16(2)) + uint16((utils.InlineIf(m.GetSegmentedMessage(), func() any { return uint16(uint16(2)) }, func() any { return uint16(uint16(0)) }).(uint16))))
}

func (m *_APDUComplexAck) GetSegmentReduction() uint16 {
	ctx := context.Background()
	_ = ctx
	sequenceNumber := m.GetSequenceNumber()
	_ = sequenceNumber
	proposedWindowSize := m.GetProposedWindowSize()
	_ = proposedWindowSize
	serviceAck := m.GetServiceAck()
	_ = serviceAck
	segmentServiceChoice := m.GetSegmentServiceChoice()
	_ = segmentServiceChoice
	return uint16(utils.InlineIf((bool((m.GetSegmentServiceChoice()) != (nil))), func() any { return uint16((uint16(m.GetApduHeaderReduction()) + uint16(uint16(1)))) }, func() any { return uint16(m.GetApduHeaderReduction()) }).(uint16))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAPDUComplexAck(structType any) APDUComplexAck {
	if casted, ok := structType.(APDUComplexAck); ok {
		return casted
	}
	if casted, ok := structType.(*APDUComplexAck); ok {
		return *casted
	}
	return nil
}

func (m *_APDUComplexAck) GetTypeName() string {
	return "APDUComplexAck"
}

func (m *_APDUComplexAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.APDUContract.(*_APDU).getLengthInBits(ctx))

	// Simple field (segmentedMessage)
	lengthInBits += 1

	// Simple field (moreFollows)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Optional Field (sequenceNumber)
	if m.SequenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (proposedWindowSize)
	if m.ProposedWindowSize != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (serviceAck)
	if m.ServiceAck != nil {
		lengthInBits += m.ServiceAck.GetLengthInBits(ctx)
	}

	// Optional Field (segmentServiceChoice)
	if m.SegmentServiceChoice != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Array field
	if len(m.Segment) > 0 {
		lengthInBits += 8 * uint16(len(m.Segment))
	}

	return lengthInBits
}

func (m *_APDUComplexAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_APDUComplexAck) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_APDU, apduLength uint16) (__aPDUComplexAck APDUComplexAck, err error) {
	m.APDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUComplexAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUComplexAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentedMessage, err := ReadSimpleField(ctx, "segmentedMessage", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentedMessage' field"))
	}
	m.SegmentedMessage = segmentedMessage

	moreFollows, err := ReadSimpleField(ctx, "moreFollows", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moreFollows' field"))
	}
	m.MoreFollows = moreFollows

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(2)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	originalInvokeId, err := ReadSimpleField(ctx, "originalInvokeId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originalInvokeId' field"))
	}
	m.OriginalInvokeId = originalInvokeId

	var sequenceNumber *uint8
	sequenceNumber, err = ReadOptionalField[uint8](ctx, "sequenceNumber", ReadUnsignedByte(readBuffer, uint8(8)), segmentedMessage)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	var proposedWindowSize *uint8
	proposedWindowSize, err = ReadOptionalField[uint8](ctx, "proposedWindowSize", ReadUnsignedByte(readBuffer, uint8(8)), segmentedMessage)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proposedWindowSize' field"))
	}
	m.ProposedWindowSize = proposedWindowSize

	apduHeaderReduction, err := ReadVirtualField[uint16](ctx, "apduHeaderReduction", (*uint16)(nil), uint16(uint16(2))+uint16((utils.InlineIf(segmentedMessage, func() any { return uint16(uint16(2)) }, func() any { return uint16(uint16(0)) }).(uint16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apduHeaderReduction' field"))
	}
	_ = apduHeaderReduction

	var serviceAck BACnetServiceAck
	_serviceAck, err := ReadOptionalField[BACnetServiceAck](ctx, "serviceAck", ReadComplex[BACnetServiceAck](BACnetServiceAckParseWithBufferProducer[BACnetServiceAck]((uint32)(uint32(apduLength)-uint32(apduHeaderReduction))), readBuffer), !(segmentedMessage))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceAck' field"))
	}
	if _serviceAck != nil {
		serviceAck = *_serviceAck
		m.ServiceAck = serviceAck
	}

	// Validation
	if !(bool((bool(!(segmentedMessage)) && bool(bool((serviceAck) != (nil))))) || bool(segmentedMessage)) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "service ack should be set"})
	}

	var segmentServiceChoice *BACnetConfirmedServiceChoice
	segmentServiceChoice, err = ReadOptionalField[BACnetConfirmedServiceChoice](ctx, "segmentServiceChoice", ReadEnum(BACnetConfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, uint8(8))), bool(segmentedMessage) && bool(bool((*sequenceNumber) != (0))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentServiceChoice' field"))
	}
	m.SegmentServiceChoice = segmentServiceChoice

	segmentReduction, err := ReadVirtualField[uint16](ctx, "segmentReduction", (*uint16)(nil), utils.InlineIf((bool((segmentServiceChoice) != (nil))), func() any { return uint16((uint16(apduHeaderReduction) + uint16(uint16(1)))) }, func() any { return uint16(apduHeaderReduction) }).(uint16))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentReduction' field"))
	}
	_ = segmentReduction

	segment, err := readBuffer.ReadByteArray("segment", int(utils.InlineIf(segmentedMessage, func() any {
		return int32((utils.InlineIf((bool((apduLength) > (0))), func() any { return int32((int32(apduLength) - int32(segmentReduction))) }, func() any { return int32(int32(0)) }).(int32)))
	}, func() any { return int32(int32(0)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segment' field"))
	}
	m.Segment = segment

	if closeErr := readBuffer.CloseContext("APDUComplexAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUComplexAck")
	}

	return m, nil
}

func (m *_APDUComplexAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUComplexAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUComplexAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUComplexAck")
		}

		if err := WriteSimpleField[bool](ctx, "segmentedMessage", m.GetSegmentedMessage(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentedMessage' field")
		}

		if err := WriteSimpleField[bool](ctx, "moreFollows", m.GetMoreFollows(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'moreFollows' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0), WriteUnsignedByte(writeBuffer, 2)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "originalInvokeId", m.GetOriginalInvokeId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'originalInvokeId' field")
		}

		if err := WriteOptionalField[uint8](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if err := WriteOptionalField[uint8](ctx, "proposedWindowSize", m.GetProposedWindowSize(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'proposedWindowSize' field")
		}
		// Virtual field
		apduHeaderReduction := m.GetApduHeaderReduction()
		_ = apduHeaderReduction
		if _apduHeaderReductionErr := writeBuffer.WriteVirtual(ctx, "apduHeaderReduction", m.GetApduHeaderReduction()); _apduHeaderReductionErr != nil {
			return errors.Wrap(_apduHeaderReductionErr, "Error serializing 'apduHeaderReduction' field")
		}

		if err := WriteOptionalField[BACnetServiceAck](ctx, "serviceAck", GetRef(m.GetServiceAck()), WriteComplex[BACnetServiceAck](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceAck' field")
		}

		if err := WriteOptionalEnumField[BACnetConfirmedServiceChoice](ctx, "segmentServiceChoice", "BACnetConfirmedServiceChoice", m.GetSegmentServiceChoice(), WriteEnum[BACnetConfirmedServiceChoice, uint8](BACnetConfirmedServiceChoice.GetValue, BACnetConfirmedServiceChoice.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8)), bool(m.GetSegmentedMessage()) && bool(bool((*m.GetSequenceNumber()) != (0)))); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentServiceChoice' field")
		}
		// Virtual field
		segmentReduction := m.GetSegmentReduction()
		_ = segmentReduction
		if _segmentReductionErr := writeBuffer.WriteVirtual(ctx, "segmentReduction", m.GetSegmentReduction()); _segmentReductionErr != nil {
			return errors.Wrap(_segmentReductionErr, "Error serializing 'segmentReduction' field")
		}

		if err := WriteByteArrayField(ctx, "segment", m.GetSegment(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'segment' field")
		}

		if popErr := writeBuffer.PopContext("APDUComplexAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUComplexAck")
		}
		return nil
	}
	return m.APDUContract.(*_APDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUComplexAck) IsAPDUComplexAck() {}

func (m *_APDUComplexAck) DeepCopy() any {
	return m.deepCopy()
}

func (m *_APDUComplexAck) deepCopy() *_APDUComplexAck {
	if m == nil {
		return nil
	}
	_APDUComplexAckCopy := &_APDUComplexAck{
		m.APDUContract.(*_APDU).deepCopy(),
		m.SegmentedMessage,
		m.MoreFollows,
		m.OriginalInvokeId,
		utils.CopyPtr[uint8](m.SequenceNumber),
		utils.CopyPtr[uint8](m.ProposedWindowSize),
		utils.DeepCopy[BACnetServiceAck](m.ServiceAck),
		utils.CopyPtr[BACnetConfirmedServiceChoice](m.SegmentServiceChoice),
		utils.DeepCopySlice[byte, byte](m.Segment),
		m.reservedField0,
	}
	m.APDUContract.(*_APDU)._SubType = m
	return _APDUComplexAckCopy
}

func (m *_APDUComplexAck) String() string {
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

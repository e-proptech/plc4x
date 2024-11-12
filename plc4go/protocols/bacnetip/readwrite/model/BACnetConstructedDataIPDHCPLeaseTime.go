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

// BACnetConstructedDataIPDHCPLeaseTime is the corresponding interface of BACnetConstructedDataIPDHCPLeaseTime
type BACnetConstructedDataIPDHCPLeaseTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpDhcpLeaseTime returns IpDhcpLeaseTime (property field)
	GetIpDhcpLeaseTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataIPDHCPLeaseTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPDHCPLeaseTime()
	// CreateBuilder creates a BACnetConstructedDataIPDHCPLeaseTimeBuilder
	CreateBACnetConstructedDataIPDHCPLeaseTimeBuilder() BACnetConstructedDataIPDHCPLeaseTimeBuilder
}

// _BACnetConstructedDataIPDHCPLeaseTime is the data-structure of this message
type _BACnetConstructedDataIPDHCPLeaseTime struct {
	BACnetConstructedDataContract
	IpDhcpLeaseTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataIPDHCPLeaseTime = (*_BACnetConstructedDataIPDHCPLeaseTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPDHCPLeaseTime)(nil)

// NewBACnetConstructedDataIPDHCPLeaseTime factory function for _BACnetConstructedDataIPDHCPLeaseTime
func NewBACnetConstructedDataIPDHCPLeaseTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipDhcpLeaseTime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPDHCPLeaseTime {
	if ipDhcpLeaseTime == nil {
		panic("ipDhcpLeaseTime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataIPDHCPLeaseTime must not be nil")
	}
	_result := &_BACnetConstructedDataIPDHCPLeaseTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		IpDhcpLeaseTime:               ipDhcpLeaseTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPDHCPLeaseTimeBuilder is a builder for BACnetConstructedDataIPDHCPLeaseTime
type BACnetConstructedDataIPDHCPLeaseTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipDhcpLeaseTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPDHCPLeaseTimeBuilder
	// WithIpDhcpLeaseTime adds IpDhcpLeaseTime (property field)
	WithIpDhcpLeaseTime(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPDHCPLeaseTimeBuilder
	// WithIpDhcpLeaseTimeBuilder adds IpDhcpLeaseTime (property field) which is build by the builder
	WithIpDhcpLeaseTimeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPDHCPLeaseTimeBuilder
	// Build builds the BACnetConstructedDataIPDHCPLeaseTime or returns an error if something is wrong
	Build() (BACnetConstructedDataIPDHCPLeaseTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPDHCPLeaseTime
}

// NewBACnetConstructedDataIPDHCPLeaseTimeBuilder() creates a BACnetConstructedDataIPDHCPLeaseTimeBuilder
func NewBACnetConstructedDataIPDHCPLeaseTimeBuilder() BACnetConstructedDataIPDHCPLeaseTimeBuilder {
	return &_BACnetConstructedDataIPDHCPLeaseTimeBuilder{_BACnetConstructedDataIPDHCPLeaseTime: new(_BACnetConstructedDataIPDHCPLeaseTime)}
}

type _BACnetConstructedDataIPDHCPLeaseTimeBuilder struct {
	*_BACnetConstructedDataIPDHCPLeaseTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPDHCPLeaseTimeBuilder) = (*_BACnetConstructedDataIPDHCPLeaseTimeBuilder)(nil)

func (b *_BACnetConstructedDataIPDHCPLeaseTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIPDHCPLeaseTimeBuilder) WithMandatoryFields(ipDhcpLeaseTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPDHCPLeaseTimeBuilder {
	return b.WithIpDhcpLeaseTime(ipDhcpLeaseTime)
}

func (b *_BACnetConstructedDataIPDHCPLeaseTimeBuilder) WithIpDhcpLeaseTime(ipDhcpLeaseTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPDHCPLeaseTimeBuilder {
	b.IpDhcpLeaseTime = ipDhcpLeaseTime
	return b
}

func (b *_BACnetConstructedDataIPDHCPLeaseTimeBuilder) WithIpDhcpLeaseTimeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPDHCPLeaseTimeBuilder {
	builder := builderSupplier(b.IpDhcpLeaseTime.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.IpDhcpLeaseTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPDHCPLeaseTimeBuilder) Build() (BACnetConstructedDataIPDHCPLeaseTime, error) {
	if b.IpDhcpLeaseTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ipDhcpLeaseTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPDHCPLeaseTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPDHCPLeaseTimeBuilder) MustBuild() BACnetConstructedDataIPDHCPLeaseTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataIPDHCPLeaseTimeBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPDHCPLeaseTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPDHCPLeaseTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPDHCPLeaseTimeBuilder().(*_BACnetConstructedDataIPDHCPLeaseTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPDHCPLeaseTimeBuilder creates a BACnetConstructedDataIPDHCPLeaseTimeBuilder
func (b *_BACnetConstructedDataIPDHCPLeaseTime) CreateBACnetConstructedDataIPDHCPLeaseTimeBuilder() BACnetConstructedDataIPDHCPLeaseTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPDHCPLeaseTimeBuilder()
	}
	return &_BACnetConstructedDataIPDHCPLeaseTimeBuilder{_BACnetConstructedDataIPDHCPLeaseTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DHCP_LEASE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetIpDhcpLeaseTime() BACnetApplicationTagUnsignedInteger {
	return m.IpDhcpLeaseTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpDhcpLeaseTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPDHCPLeaseTime(structType any) BACnetConstructedDataIPDHCPLeaseTime {
	if casted, ok := structType.(BACnetConstructedDataIPDHCPLeaseTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDHCPLeaseTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetTypeName() string {
	return "BACnetConstructedDataIPDHCPLeaseTime"
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipDhcpLeaseTime)
	lengthInBits += m.IpDhcpLeaseTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPDHCPLeaseTime BACnetConstructedDataIPDHCPLeaseTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDHCPLeaseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPDHCPLeaseTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipDhcpLeaseTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipDhcpLeaseTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipDhcpLeaseTime' field"))
	}
	m.IpDhcpLeaseTime = ipDhcpLeaseTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), ipDhcpLeaseTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDHCPLeaseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPDHCPLeaseTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDHCPLeaseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPDHCPLeaseTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipDhcpLeaseTime", m.GetIpDhcpLeaseTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipDhcpLeaseTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDHCPLeaseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPDHCPLeaseTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) IsBACnetConstructedDataIPDHCPLeaseTime() {}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) deepCopy() *_BACnetConstructedDataIPDHCPLeaseTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPDHCPLeaseTimeCopy := &_BACnetConstructedDataIPDHCPLeaseTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.IpDhcpLeaseTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPDHCPLeaseTimeCopy
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) String() string {
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

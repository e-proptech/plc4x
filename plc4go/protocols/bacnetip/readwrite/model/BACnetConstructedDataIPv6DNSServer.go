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

// BACnetConstructedDataIPv6DNSServer is the corresponding interface of BACnetConstructedDataIPv6DNSServer
type BACnetConstructedDataIPv6DNSServer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetIpv6DnsServer returns Ipv6DnsServer (property field)
	GetIpv6DnsServer() []BACnetApplicationTagOctetString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataIPv6DNSServer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPv6DNSServer()
	// CreateBuilder creates a BACnetConstructedDataIPv6DNSServerBuilder
	CreateBACnetConstructedDataIPv6DNSServerBuilder() BACnetConstructedDataIPv6DNSServerBuilder
}

// _BACnetConstructedDataIPv6DNSServer is the data-structure of this message
type _BACnetConstructedDataIPv6DNSServer struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	Ipv6DnsServer        []BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataIPv6DNSServer = (*_BACnetConstructedDataIPv6DNSServer)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPv6DNSServer)(nil)

// NewBACnetConstructedDataIPv6DNSServer factory function for _BACnetConstructedDataIPv6DNSServer
func NewBACnetConstructedDataIPv6DNSServer(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, ipv6DnsServer []BACnetApplicationTagOctetString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6DNSServer {
	_result := &_BACnetConstructedDataIPv6DNSServer{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		Ipv6DnsServer:                 ipv6DnsServer,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPv6DNSServerBuilder is a builder for BACnetConstructedDataIPv6DNSServer
type BACnetConstructedDataIPv6DNSServerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipv6DnsServer []BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DNSServerBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6DNSServerBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPv6DNSServerBuilder
	// WithIpv6DnsServer adds Ipv6DnsServer (property field)
	WithIpv6DnsServer(...BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DNSServerBuilder
	// Build builds the BACnetConstructedDataIPv6DNSServer or returns an error if something is wrong
	Build() (BACnetConstructedDataIPv6DNSServer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPv6DNSServer
}

// NewBACnetConstructedDataIPv6DNSServerBuilder() creates a BACnetConstructedDataIPv6DNSServerBuilder
func NewBACnetConstructedDataIPv6DNSServerBuilder() BACnetConstructedDataIPv6DNSServerBuilder {
	return &_BACnetConstructedDataIPv6DNSServerBuilder{_BACnetConstructedDataIPv6DNSServer: new(_BACnetConstructedDataIPv6DNSServer)}
}

type _BACnetConstructedDataIPv6DNSServerBuilder struct {
	*_BACnetConstructedDataIPv6DNSServer

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPv6DNSServerBuilder) = (*_BACnetConstructedDataIPv6DNSServerBuilder)(nil)

func (b *_BACnetConstructedDataIPv6DNSServerBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIPv6DNSServerBuilder) WithMandatoryFields(ipv6DnsServer []BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DNSServerBuilder {
	return b.WithIpv6DnsServer(ipv6DnsServer...)
}

func (b *_BACnetConstructedDataIPv6DNSServerBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6DNSServerBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataIPv6DNSServerBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPv6DNSServerBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPv6DNSServerBuilder) WithIpv6DnsServer(ipv6DnsServer ...BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DNSServerBuilder {
	b.Ipv6DnsServer = ipv6DnsServer
	return b
}

func (b *_BACnetConstructedDataIPv6DNSServerBuilder) Build() (BACnetConstructedDataIPv6DNSServer, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPv6DNSServer.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPv6DNSServerBuilder) MustBuild() BACnetConstructedDataIPv6DNSServer {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataIPv6DNSServerBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPv6DNSServerBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPv6DNSServerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPv6DNSServerBuilder().(*_BACnetConstructedDataIPv6DNSServerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPv6DNSServerBuilder creates a BACnetConstructedDataIPv6DNSServerBuilder
func (b *_BACnetConstructedDataIPv6DNSServer) CreateBACnetConstructedDataIPv6DNSServerBuilder() BACnetConstructedDataIPv6DNSServerBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPv6DNSServerBuilder()
	}
	return &_BACnetConstructedDataIPv6DNSServerBuilder{_BACnetConstructedDataIPv6DNSServer: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6DNSServer) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6DNSServer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DNS_SERVER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6DNSServer) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DNSServer) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataIPv6DNSServer) GetIpv6DnsServer() []BACnetApplicationTagOctetString {
	return m.Ipv6DnsServer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DNSServer) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6DNSServer(structType any) BACnetConstructedDataIPv6DNSServer {
	if casted, ok := structType.(BACnetConstructedDataIPv6DNSServer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DNSServer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6DNSServer) GetTypeName() string {
	return "BACnetConstructedDataIPv6DNSServer"
}

func (m *_BACnetConstructedDataIPv6DNSServer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.Ipv6DnsServer) > 0 {
		for _, element := range m.Ipv6DnsServer {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6DNSServer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPv6DNSServer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPv6DNSServer BACnetConstructedDataIPv6DNSServer, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DNSServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6DNSServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	ipv6DnsServer, err := ReadTerminatedArrayField[BACnetApplicationTagOctetString](ctx, "ipv6DnsServer", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6DnsServer' field"))
	}
	m.Ipv6DnsServer = ipv6DnsServer

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DNSServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6DNSServer")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPv6DNSServer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6DNSServer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DNSServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6DNSServer")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "ipv6DnsServer", m.GetIpv6DnsServer(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6DnsServer' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DNSServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6DNSServer")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6DNSServer) IsBACnetConstructedDataIPv6DNSServer() {}

func (m *_BACnetConstructedDataIPv6DNSServer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPv6DNSServer) deepCopy() *_BACnetConstructedDataIPv6DNSServer {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPv6DNSServerCopy := &_BACnetConstructedDataIPv6DNSServer{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfDataElements),
		utils.DeepCopySlice[BACnetApplicationTagOctetString, BACnetApplicationTagOctetString](m.Ipv6DnsServer),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPv6DNSServerCopy
}

func (m *_BACnetConstructedDataIPv6DNSServer) String() string {
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

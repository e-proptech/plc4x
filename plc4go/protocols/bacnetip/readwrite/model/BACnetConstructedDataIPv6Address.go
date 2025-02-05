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

// BACnetConstructedDataIPv6Address is the corresponding interface of BACnetConstructedDataIPv6Address
type BACnetConstructedDataIPv6Address interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpv6Address returns Ipv6Address (property field)
	GetIpv6Address() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
	// IsBACnetConstructedDataIPv6Address is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPv6Address()
	// CreateBuilder creates a BACnetConstructedDataIPv6AddressBuilder
	CreateBACnetConstructedDataIPv6AddressBuilder() BACnetConstructedDataIPv6AddressBuilder
}

// _BACnetConstructedDataIPv6Address is the data-structure of this message
type _BACnetConstructedDataIPv6Address struct {
	BACnetConstructedDataContract
	Ipv6Address BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataIPv6Address = (*_BACnetConstructedDataIPv6Address)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPv6Address)(nil)

// NewBACnetConstructedDataIPv6Address factory function for _BACnetConstructedDataIPv6Address
func NewBACnetConstructedDataIPv6Address(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipv6Address BACnetApplicationTagOctetString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6Address {
	if ipv6Address == nil {
		panic("ipv6Address of type BACnetApplicationTagOctetString for BACnetConstructedDataIPv6Address must not be nil")
	}
	_result := &_BACnetConstructedDataIPv6Address{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Ipv6Address:                   ipv6Address,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPv6AddressBuilder is a builder for BACnetConstructedDataIPv6Address
type BACnetConstructedDataIPv6AddressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipv6Address BACnetApplicationTagOctetString) BACnetConstructedDataIPv6AddressBuilder
	// WithIpv6Address adds Ipv6Address (property field)
	WithIpv6Address(BACnetApplicationTagOctetString) BACnetConstructedDataIPv6AddressBuilder
	// WithIpv6AddressBuilder adds Ipv6Address (property field) which is build by the builder
	WithIpv6AddressBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataIPv6AddressBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIPv6Address or returns an error if something is wrong
	Build() (BACnetConstructedDataIPv6Address, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPv6Address
}

// NewBACnetConstructedDataIPv6AddressBuilder() creates a BACnetConstructedDataIPv6AddressBuilder
func NewBACnetConstructedDataIPv6AddressBuilder() BACnetConstructedDataIPv6AddressBuilder {
	return &_BACnetConstructedDataIPv6AddressBuilder{_BACnetConstructedDataIPv6Address: new(_BACnetConstructedDataIPv6Address)}
}

type _BACnetConstructedDataIPv6AddressBuilder struct {
	*_BACnetConstructedDataIPv6Address

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPv6AddressBuilder) = (*_BACnetConstructedDataIPv6AddressBuilder)(nil)

func (b *_BACnetConstructedDataIPv6AddressBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataIPv6Address
}

func (b *_BACnetConstructedDataIPv6AddressBuilder) WithMandatoryFields(ipv6Address BACnetApplicationTagOctetString) BACnetConstructedDataIPv6AddressBuilder {
	return b.WithIpv6Address(ipv6Address)
}

func (b *_BACnetConstructedDataIPv6AddressBuilder) WithIpv6Address(ipv6Address BACnetApplicationTagOctetString) BACnetConstructedDataIPv6AddressBuilder {
	b.Ipv6Address = ipv6Address
	return b
}

func (b *_BACnetConstructedDataIPv6AddressBuilder) WithIpv6AddressBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataIPv6AddressBuilder {
	builder := builderSupplier(b.Ipv6Address.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	b.Ipv6Address, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPv6AddressBuilder) Build() (BACnetConstructedDataIPv6Address, error) {
	if b.Ipv6Address == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ipv6Address' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPv6Address.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPv6AddressBuilder) MustBuild() BACnetConstructedDataIPv6Address {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIPv6AddressBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPv6AddressBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPv6AddressBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPv6AddressBuilder().(*_BACnetConstructedDataIPv6AddressBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPv6AddressBuilder creates a BACnetConstructedDataIPv6AddressBuilder
func (b *_BACnetConstructedDataIPv6Address) CreateBACnetConstructedDataIPv6AddressBuilder() BACnetConstructedDataIPv6AddressBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPv6AddressBuilder()
	}
	return &_BACnetConstructedDataIPv6AddressBuilder{_BACnetConstructedDataIPv6Address: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6Address) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6Address) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6Address) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6Address) GetIpv6Address() BACnetApplicationTagOctetString {
	return m.Ipv6Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6Address) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetIpv6Address())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6Address(structType any) BACnetConstructedDataIPv6Address {
	if casted, ok := structType.(BACnetConstructedDataIPv6Address); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6Address); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6Address) GetTypeName() string {
	return "BACnetConstructedDataIPv6Address"
}

func (m *_BACnetConstructedDataIPv6Address) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipv6Address)
	lengthInBits += m.Ipv6Address.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6Address) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPv6Address) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPv6Address BACnetConstructedDataIPv6Address, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6Address"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6Address")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipv6Address, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "ipv6Address", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6Address' field"))
	}
	m.Ipv6Address = ipv6Address

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), ipv6Address)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6Address"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6Address")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPv6Address) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6Address) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6Address"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6Address")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "ipv6Address", m.GetIpv6Address(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6Address' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6Address"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6Address")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6Address) IsBACnetConstructedDataIPv6Address() {}

func (m *_BACnetConstructedDataIPv6Address) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPv6Address) deepCopy() *_BACnetConstructedDataIPv6Address {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPv6AddressCopy := &_BACnetConstructedDataIPv6Address{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagOctetString](m.Ipv6Address),
	}
	_BACnetConstructedDataIPv6AddressCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPv6AddressCopy
}

func (m *_BACnetConstructedDataIPv6Address) String() string {
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

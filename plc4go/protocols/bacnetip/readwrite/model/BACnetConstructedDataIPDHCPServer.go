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

// BACnetConstructedDataIPDHCPServer is the corresponding interface of BACnetConstructedDataIPDHCPServer
type BACnetConstructedDataIPDHCPServer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDhcpServer returns DhcpServer (property field)
	GetDhcpServer() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
	// IsBACnetConstructedDataIPDHCPServer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPDHCPServer()
	// CreateBuilder creates a BACnetConstructedDataIPDHCPServerBuilder
	CreateBACnetConstructedDataIPDHCPServerBuilder() BACnetConstructedDataIPDHCPServerBuilder
}

// _BACnetConstructedDataIPDHCPServer is the data-structure of this message
type _BACnetConstructedDataIPDHCPServer struct {
	BACnetConstructedDataContract
	DhcpServer BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataIPDHCPServer = (*_BACnetConstructedDataIPDHCPServer)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPDHCPServer)(nil)

// NewBACnetConstructedDataIPDHCPServer factory function for _BACnetConstructedDataIPDHCPServer
func NewBACnetConstructedDataIPDHCPServer(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, dhcpServer BACnetApplicationTagOctetString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPDHCPServer {
	if dhcpServer == nil {
		panic("dhcpServer of type BACnetApplicationTagOctetString for BACnetConstructedDataIPDHCPServer must not be nil")
	}
	_result := &_BACnetConstructedDataIPDHCPServer{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DhcpServer:                    dhcpServer,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPDHCPServerBuilder is a builder for BACnetConstructedDataIPDHCPServer
type BACnetConstructedDataIPDHCPServerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dhcpServer BACnetApplicationTagOctetString) BACnetConstructedDataIPDHCPServerBuilder
	// WithDhcpServer adds DhcpServer (property field)
	WithDhcpServer(BACnetApplicationTagOctetString) BACnetConstructedDataIPDHCPServerBuilder
	// WithDhcpServerBuilder adds DhcpServer (property field) which is build by the builder
	WithDhcpServerBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataIPDHCPServerBuilder
	// Build builds the BACnetConstructedDataIPDHCPServer or returns an error if something is wrong
	Build() (BACnetConstructedDataIPDHCPServer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPDHCPServer
}

// NewBACnetConstructedDataIPDHCPServerBuilder() creates a BACnetConstructedDataIPDHCPServerBuilder
func NewBACnetConstructedDataIPDHCPServerBuilder() BACnetConstructedDataIPDHCPServerBuilder {
	return &_BACnetConstructedDataIPDHCPServerBuilder{_BACnetConstructedDataIPDHCPServer: new(_BACnetConstructedDataIPDHCPServer)}
}

type _BACnetConstructedDataIPDHCPServerBuilder struct {
	*_BACnetConstructedDataIPDHCPServer

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPDHCPServerBuilder) = (*_BACnetConstructedDataIPDHCPServerBuilder)(nil)

func (b *_BACnetConstructedDataIPDHCPServerBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIPDHCPServerBuilder) WithMandatoryFields(dhcpServer BACnetApplicationTagOctetString) BACnetConstructedDataIPDHCPServerBuilder {
	return b.WithDhcpServer(dhcpServer)
}

func (b *_BACnetConstructedDataIPDHCPServerBuilder) WithDhcpServer(dhcpServer BACnetApplicationTagOctetString) BACnetConstructedDataIPDHCPServerBuilder {
	b.DhcpServer = dhcpServer
	return b
}

func (b *_BACnetConstructedDataIPDHCPServerBuilder) WithDhcpServerBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataIPDHCPServerBuilder {
	builder := builderSupplier(b.DhcpServer.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	b.DhcpServer, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPDHCPServerBuilder) Build() (BACnetConstructedDataIPDHCPServer, error) {
	if b.DhcpServer == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dhcpServer' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPDHCPServer.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPDHCPServerBuilder) MustBuild() BACnetConstructedDataIPDHCPServer {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataIPDHCPServerBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPDHCPServerBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPDHCPServerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPDHCPServerBuilder().(*_BACnetConstructedDataIPDHCPServerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPDHCPServerBuilder creates a BACnetConstructedDataIPDHCPServerBuilder
func (b *_BACnetConstructedDataIPDHCPServer) CreateBACnetConstructedDataIPDHCPServerBuilder() BACnetConstructedDataIPDHCPServerBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPDHCPServerBuilder()
	}
	return &_BACnetConstructedDataIPDHCPServerBuilder{_BACnetConstructedDataIPDHCPServer: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPServer) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPDHCPServer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DHCP_SERVER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPDHCPServer) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPServer) GetDhcpServer() BACnetApplicationTagOctetString {
	return m.DhcpServer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPServer) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetDhcpServer())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPDHCPServer(structType any) BACnetConstructedDataIPDHCPServer {
	if casted, ok := structType.(BACnetConstructedDataIPDHCPServer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDHCPServer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPDHCPServer) GetTypeName() string {
	return "BACnetConstructedDataIPDHCPServer"
}

func (m *_BACnetConstructedDataIPDHCPServer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (dhcpServer)
	lengthInBits += m.DhcpServer.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPDHCPServer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPDHCPServer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPDHCPServer BACnetConstructedDataIPDHCPServer, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDHCPServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPDHCPServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dhcpServer, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "dhcpServer", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dhcpServer' field"))
	}
	m.DhcpServer = dhcpServer

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), dhcpServer)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDHCPServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPDHCPServer")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPDHCPServer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPDHCPServer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDHCPServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPDHCPServer")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "dhcpServer", m.GetDhcpServer(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dhcpServer' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDHCPServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPDHCPServer")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPDHCPServer) IsBACnetConstructedDataIPDHCPServer() {}

func (m *_BACnetConstructedDataIPDHCPServer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPDHCPServer) deepCopy() *_BACnetConstructedDataIPDHCPServer {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPDHCPServerCopy := &_BACnetConstructedDataIPDHCPServer{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagOctetString](m.DhcpServer),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPDHCPServerCopy
}

func (m *_BACnetConstructedDataIPDHCPServer) String() string {
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

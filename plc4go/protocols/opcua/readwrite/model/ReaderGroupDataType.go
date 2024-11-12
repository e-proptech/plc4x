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

// ReaderGroupDataType is the corresponding interface of ReaderGroupDataType
type ReaderGroupDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetEnabled returns Enabled (property field)
	GetEnabled() bool
	// GetSecurityMode returns SecurityMode (property field)
	GetSecurityMode() MessageSecurityMode
	// GetSecurityGroupId returns SecurityGroupId (property field)
	GetSecurityGroupId() PascalString
	// GetSecurityKeyServices returns SecurityKeyServices (property field)
	GetSecurityKeyServices() []EndpointDescription
	// GetMaxNetworkMessageSize returns MaxNetworkMessageSize (property field)
	GetMaxNetworkMessageSize() uint32
	// GetGroupProperties returns GroupProperties (property field)
	GetGroupProperties() []KeyValuePair
	// GetTransportSettings returns TransportSettings (property field)
	GetTransportSettings() ExtensionObject
	// GetMessageSettings returns MessageSettings (property field)
	GetMessageSettings() ExtensionObject
	// GetDataSetReaders returns DataSetReaders (property field)
	GetDataSetReaders() []DataSetReaderDataType
	// IsReaderGroupDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReaderGroupDataType()
	// CreateBuilder creates a ReaderGroupDataTypeBuilder
	CreateReaderGroupDataTypeBuilder() ReaderGroupDataTypeBuilder
}

// _ReaderGroupDataType is the data-structure of this message
type _ReaderGroupDataType struct {
	ExtensionObjectDefinitionContract
	Name                  PascalString
	Enabled               bool
	SecurityMode          MessageSecurityMode
	SecurityGroupId       PascalString
	SecurityKeyServices   []EndpointDescription
	MaxNetworkMessageSize uint32
	GroupProperties       []KeyValuePair
	TransportSettings     ExtensionObject
	MessageSettings       ExtensionObject
	DataSetReaders        []DataSetReaderDataType
	// Reserved Fields
	reservedField0 *uint8
}

var _ ReaderGroupDataType = (*_ReaderGroupDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReaderGroupDataType)(nil)

// NewReaderGroupDataType factory function for _ReaderGroupDataType
func NewReaderGroupDataType(name PascalString, enabled bool, securityMode MessageSecurityMode, securityGroupId PascalString, securityKeyServices []EndpointDescription, maxNetworkMessageSize uint32, groupProperties []KeyValuePair, transportSettings ExtensionObject, messageSettings ExtensionObject, dataSetReaders []DataSetReaderDataType) *_ReaderGroupDataType {
	if name == nil {
		panic("name of type PascalString for ReaderGroupDataType must not be nil")
	}
	if securityGroupId == nil {
		panic("securityGroupId of type PascalString for ReaderGroupDataType must not be nil")
	}
	if transportSettings == nil {
		panic("transportSettings of type ExtensionObject for ReaderGroupDataType must not be nil")
	}
	if messageSettings == nil {
		panic("messageSettings of type ExtensionObject for ReaderGroupDataType must not be nil")
	}
	_result := &_ReaderGroupDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Name:                              name,
		Enabled:                           enabled,
		SecurityMode:                      securityMode,
		SecurityGroupId:                   securityGroupId,
		SecurityKeyServices:               securityKeyServices,
		MaxNetworkMessageSize:             maxNetworkMessageSize,
		GroupProperties:                   groupProperties,
		TransportSettings:                 transportSettings,
		MessageSettings:                   messageSettings,
		DataSetReaders:                    dataSetReaders,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReaderGroupDataTypeBuilder is a builder for ReaderGroupDataType
type ReaderGroupDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(name PascalString, enabled bool, securityMode MessageSecurityMode, securityGroupId PascalString, securityKeyServices []EndpointDescription, maxNetworkMessageSize uint32, groupProperties []KeyValuePair, transportSettings ExtensionObject, messageSettings ExtensionObject, dataSetReaders []DataSetReaderDataType) ReaderGroupDataTypeBuilder
	// WithName adds Name (property field)
	WithName(PascalString) ReaderGroupDataTypeBuilder
	// WithNameBuilder adds Name (property field) which is build by the builder
	WithNameBuilder(func(PascalStringBuilder) PascalStringBuilder) ReaderGroupDataTypeBuilder
	// WithEnabled adds Enabled (property field)
	WithEnabled(bool) ReaderGroupDataTypeBuilder
	// WithSecurityMode adds SecurityMode (property field)
	WithSecurityMode(MessageSecurityMode) ReaderGroupDataTypeBuilder
	// WithSecurityGroupId adds SecurityGroupId (property field)
	WithSecurityGroupId(PascalString) ReaderGroupDataTypeBuilder
	// WithSecurityGroupIdBuilder adds SecurityGroupId (property field) which is build by the builder
	WithSecurityGroupIdBuilder(func(PascalStringBuilder) PascalStringBuilder) ReaderGroupDataTypeBuilder
	// WithSecurityKeyServices adds SecurityKeyServices (property field)
	WithSecurityKeyServices(...EndpointDescription) ReaderGroupDataTypeBuilder
	// WithMaxNetworkMessageSize adds MaxNetworkMessageSize (property field)
	WithMaxNetworkMessageSize(uint32) ReaderGroupDataTypeBuilder
	// WithGroupProperties adds GroupProperties (property field)
	WithGroupProperties(...KeyValuePair) ReaderGroupDataTypeBuilder
	// WithTransportSettings adds TransportSettings (property field)
	WithTransportSettings(ExtensionObject) ReaderGroupDataTypeBuilder
	// WithTransportSettingsBuilder adds TransportSettings (property field) which is build by the builder
	WithTransportSettingsBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) ReaderGroupDataTypeBuilder
	// WithMessageSettings adds MessageSettings (property field)
	WithMessageSettings(ExtensionObject) ReaderGroupDataTypeBuilder
	// WithMessageSettingsBuilder adds MessageSettings (property field) which is build by the builder
	WithMessageSettingsBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) ReaderGroupDataTypeBuilder
	// WithDataSetReaders adds DataSetReaders (property field)
	WithDataSetReaders(...DataSetReaderDataType) ReaderGroupDataTypeBuilder
	// Build builds the ReaderGroupDataType or returns an error if something is wrong
	Build() (ReaderGroupDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReaderGroupDataType
}

// NewReaderGroupDataTypeBuilder() creates a ReaderGroupDataTypeBuilder
func NewReaderGroupDataTypeBuilder() ReaderGroupDataTypeBuilder {
	return &_ReaderGroupDataTypeBuilder{_ReaderGroupDataType: new(_ReaderGroupDataType)}
}

type _ReaderGroupDataTypeBuilder struct {
	*_ReaderGroupDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ReaderGroupDataTypeBuilder) = (*_ReaderGroupDataTypeBuilder)(nil)

func (b *_ReaderGroupDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ReaderGroupDataTypeBuilder) WithMandatoryFields(name PascalString, enabled bool, securityMode MessageSecurityMode, securityGroupId PascalString, securityKeyServices []EndpointDescription, maxNetworkMessageSize uint32, groupProperties []KeyValuePair, transportSettings ExtensionObject, messageSettings ExtensionObject, dataSetReaders []DataSetReaderDataType) ReaderGroupDataTypeBuilder {
	return b.WithName(name).WithEnabled(enabled).WithSecurityMode(securityMode).WithSecurityGroupId(securityGroupId).WithSecurityKeyServices(securityKeyServices...).WithMaxNetworkMessageSize(maxNetworkMessageSize).WithGroupProperties(groupProperties...).WithTransportSettings(transportSettings).WithMessageSettings(messageSettings).WithDataSetReaders(dataSetReaders...)
}

func (b *_ReaderGroupDataTypeBuilder) WithName(name PascalString) ReaderGroupDataTypeBuilder {
	b.Name = name
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) ReaderGroupDataTypeBuilder {
	builder := builderSupplier(b.Name.CreatePascalStringBuilder())
	var err error
	b.Name, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithEnabled(enabled bool) ReaderGroupDataTypeBuilder {
	b.Enabled = enabled
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithSecurityMode(securityMode MessageSecurityMode) ReaderGroupDataTypeBuilder {
	b.SecurityMode = securityMode
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithSecurityGroupId(securityGroupId PascalString) ReaderGroupDataTypeBuilder {
	b.SecurityGroupId = securityGroupId
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithSecurityGroupIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) ReaderGroupDataTypeBuilder {
	builder := builderSupplier(b.SecurityGroupId.CreatePascalStringBuilder())
	var err error
	b.SecurityGroupId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithSecurityKeyServices(securityKeyServices ...EndpointDescription) ReaderGroupDataTypeBuilder {
	b.SecurityKeyServices = securityKeyServices
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithMaxNetworkMessageSize(maxNetworkMessageSize uint32) ReaderGroupDataTypeBuilder {
	b.MaxNetworkMessageSize = maxNetworkMessageSize
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithGroupProperties(groupProperties ...KeyValuePair) ReaderGroupDataTypeBuilder {
	b.GroupProperties = groupProperties
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithTransportSettings(transportSettings ExtensionObject) ReaderGroupDataTypeBuilder {
	b.TransportSettings = transportSettings
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithTransportSettingsBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) ReaderGroupDataTypeBuilder {
	builder := builderSupplier(b.TransportSettings.CreateExtensionObjectBuilder())
	var err error
	b.TransportSettings, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithMessageSettings(messageSettings ExtensionObject) ReaderGroupDataTypeBuilder {
	b.MessageSettings = messageSettings
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithMessageSettingsBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) ReaderGroupDataTypeBuilder {
	builder := builderSupplier(b.MessageSettings.CreateExtensionObjectBuilder())
	var err error
	b.MessageSettings, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_ReaderGroupDataTypeBuilder) WithDataSetReaders(dataSetReaders ...DataSetReaderDataType) ReaderGroupDataTypeBuilder {
	b.DataSetReaders = dataSetReaders
	return b
}

func (b *_ReaderGroupDataTypeBuilder) Build() (ReaderGroupDataType, error) {
	if b.Name == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'name' not set"))
	}
	if b.SecurityGroupId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'securityGroupId' not set"))
	}
	if b.TransportSettings == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'transportSettings' not set"))
	}
	if b.MessageSettings == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'messageSettings' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ReaderGroupDataType.deepCopy(), nil
}

func (b *_ReaderGroupDataTypeBuilder) MustBuild() ReaderGroupDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ReaderGroupDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ReaderGroupDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ReaderGroupDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateReaderGroupDataTypeBuilder().(*_ReaderGroupDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateReaderGroupDataTypeBuilder creates a ReaderGroupDataTypeBuilder
func (b *_ReaderGroupDataType) CreateReaderGroupDataTypeBuilder() ReaderGroupDataTypeBuilder {
	if b == nil {
		return NewReaderGroupDataTypeBuilder()
	}
	return &_ReaderGroupDataTypeBuilder{_ReaderGroupDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReaderGroupDataType) GetExtensionId() int32 {
	return int32(15522)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReaderGroupDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReaderGroupDataType) GetName() PascalString {
	return m.Name
}

func (m *_ReaderGroupDataType) GetEnabled() bool {
	return m.Enabled
}

func (m *_ReaderGroupDataType) GetSecurityMode() MessageSecurityMode {
	return m.SecurityMode
}

func (m *_ReaderGroupDataType) GetSecurityGroupId() PascalString {
	return m.SecurityGroupId
}

func (m *_ReaderGroupDataType) GetSecurityKeyServices() []EndpointDescription {
	return m.SecurityKeyServices
}

func (m *_ReaderGroupDataType) GetMaxNetworkMessageSize() uint32 {
	return m.MaxNetworkMessageSize
}

func (m *_ReaderGroupDataType) GetGroupProperties() []KeyValuePair {
	return m.GroupProperties
}

func (m *_ReaderGroupDataType) GetTransportSettings() ExtensionObject {
	return m.TransportSettings
}

func (m *_ReaderGroupDataType) GetMessageSettings() ExtensionObject {
	return m.MessageSettings
}

func (m *_ReaderGroupDataType) GetDataSetReaders() []DataSetReaderDataType {
	return m.DataSetReaders
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastReaderGroupDataType(structType any) ReaderGroupDataType {
	if casted, ok := structType.(ReaderGroupDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ReaderGroupDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ReaderGroupDataType) GetTypeName() string {
	return "ReaderGroupDataType"
}

func (m *_ReaderGroupDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enabled)
	lengthInBits += 1

	// Simple field (securityMode)
	lengthInBits += 32

	// Simple field (securityGroupId)
	lengthInBits += m.SecurityGroupId.GetLengthInBits(ctx)

	// Implicit Field (noOfSecurityKeyServices)
	lengthInBits += 32

	// Array field
	if len(m.SecurityKeyServices) > 0 {
		for _curItem, element := range m.SecurityKeyServices {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SecurityKeyServices), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (maxNetworkMessageSize)
	lengthInBits += 32

	// Implicit Field (noOfGroupProperties)
	lengthInBits += 32

	// Array field
	if len(m.GroupProperties) > 0 {
		for _curItem, element := range m.GroupProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GroupProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (transportSettings)
	lengthInBits += m.TransportSettings.GetLengthInBits(ctx)

	// Simple field (messageSettings)
	lengthInBits += m.MessageSettings.GetLengthInBits(ctx)

	// Implicit Field (noOfDataSetReaders)
	lengthInBits += 32

	// Array field
	if len(m.DataSetReaders) > 0 {
		for _curItem, element := range m.DataSetReaders {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataSetReaders), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ReaderGroupDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReaderGroupDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__readerGroupDataType ReaderGroupDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReaderGroupDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReaderGroupDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	enabled, err := ReadSimpleField(ctx, "enabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enabled' field"))
	}
	m.Enabled = enabled

	securityMode, err := ReadEnumField[MessageSecurityMode](ctx, "securityMode", "MessageSecurityMode", ReadEnum(MessageSecurityModeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityMode' field"))
	}
	m.SecurityMode = securityMode

	securityGroupId, err := ReadSimpleField[PascalString](ctx, "securityGroupId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityGroupId' field"))
	}
	m.SecurityGroupId = securityGroupId

	noOfSecurityKeyServices, err := ReadImplicitField[int32](ctx, "noOfSecurityKeyServices", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSecurityKeyServices' field"))
	}
	_ = noOfSecurityKeyServices

	securityKeyServices, err := ReadCountArrayField[EndpointDescription](ctx, "securityKeyServices", ReadComplex[EndpointDescription](ExtensionObjectDefinitionParseWithBufferProducer[EndpointDescription]((int32)(int32(314))), readBuffer), uint64(noOfSecurityKeyServices))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityKeyServices' field"))
	}
	m.SecurityKeyServices = securityKeyServices

	maxNetworkMessageSize, err := ReadSimpleField(ctx, "maxNetworkMessageSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNetworkMessageSize' field"))
	}
	m.MaxNetworkMessageSize = maxNetworkMessageSize

	noOfGroupProperties, err := ReadImplicitField[int32](ctx, "noOfGroupProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfGroupProperties' field"))
	}
	_ = noOfGroupProperties

	groupProperties, err := ReadCountArrayField[KeyValuePair](ctx, "groupProperties", ReadComplex[KeyValuePair](ExtensionObjectDefinitionParseWithBufferProducer[KeyValuePair]((int32)(int32(14535))), readBuffer), uint64(noOfGroupProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'groupProperties' field"))
	}
	m.GroupProperties = groupProperties

	transportSettings, err := ReadSimpleField[ExtensionObject](ctx, "transportSettings", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSettings' field"))
	}
	m.TransportSettings = transportSettings

	messageSettings, err := ReadSimpleField[ExtensionObject](ctx, "messageSettings", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageSettings' field"))
	}
	m.MessageSettings = messageSettings

	noOfDataSetReaders, err := ReadImplicitField[int32](ctx, "noOfDataSetReaders", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDataSetReaders' field"))
	}
	_ = noOfDataSetReaders

	dataSetReaders, err := ReadCountArrayField[DataSetReaderDataType](ctx, "dataSetReaders", ReadComplex[DataSetReaderDataType](ExtensionObjectDefinitionParseWithBufferProducer[DataSetReaderDataType]((int32)(int32(15625))), readBuffer), uint64(noOfDataSetReaders))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetReaders' field"))
	}
	m.DataSetReaders = dataSetReaders

	if closeErr := readBuffer.CloseContext("ReaderGroupDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReaderGroupDataType")
	}

	return m, nil
}

func (m *_ReaderGroupDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReaderGroupDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReaderGroupDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReaderGroupDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "enabled", m.GetEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enabled' field")
		}

		if err := WriteSimpleEnumField[MessageSecurityMode](ctx, "securityMode", "MessageSecurityMode", m.GetSecurityMode(), WriteEnum[MessageSecurityMode, uint32](MessageSecurityMode.GetValue, MessageSecurityMode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'securityMode' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "securityGroupId", m.GetSecurityGroupId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityGroupId' field")
		}
		noOfSecurityKeyServices := int32(utils.InlineIf(bool((m.GetSecurityKeyServices()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetSecurityKeyServices()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfSecurityKeyServices", noOfSecurityKeyServices, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSecurityKeyServices' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "securityKeyServices", m.GetSecurityKeyServices(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'securityKeyServices' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxNetworkMessageSize", m.GetMaxNetworkMessageSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNetworkMessageSize' field")
		}
		noOfGroupProperties := int32(utils.InlineIf(bool((m.GetGroupProperties()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetGroupProperties()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfGroupProperties", noOfGroupProperties, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfGroupProperties' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "groupProperties", m.GetGroupProperties(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'groupProperties' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "transportSettings", m.GetTransportSettings(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'transportSettings' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "messageSettings", m.GetMessageSettings(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'messageSettings' field")
		}
		noOfDataSetReaders := int32(utils.InlineIf(bool((m.GetDataSetReaders()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDataSetReaders()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDataSetReaders", noOfDataSetReaders, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDataSetReaders' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "dataSetReaders", m.GetDataSetReaders(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetReaders' field")
		}

		if popErr := writeBuffer.PopContext("ReaderGroupDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReaderGroupDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReaderGroupDataType) IsReaderGroupDataType() {}

func (m *_ReaderGroupDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReaderGroupDataType) deepCopy() *_ReaderGroupDataType {
	if m == nil {
		return nil
	}
	_ReaderGroupDataTypeCopy := &_ReaderGroupDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.Name),
		m.Enabled,
		m.SecurityMode,
		utils.DeepCopy[PascalString](m.SecurityGroupId),
		utils.DeepCopySlice[EndpointDescription, EndpointDescription](m.SecurityKeyServices),
		m.MaxNetworkMessageSize,
		utils.DeepCopySlice[KeyValuePair, KeyValuePair](m.GroupProperties),
		utils.DeepCopy[ExtensionObject](m.TransportSettings),
		utils.DeepCopy[ExtensionObject](m.MessageSettings),
		utils.DeepCopySlice[DataSetReaderDataType, DataSetReaderDataType](m.DataSetReaders),
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ReaderGroupDataTypeCopy
}

func (m *_ReaderGroupDataType) String() string {
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

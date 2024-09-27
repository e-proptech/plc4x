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

// ServerStatusDataType is the corresponding interface of ServerStatusDataType
type ServerStatusDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStartTime returns StartTime (property field)
	GetStartTime() int64
	// GetCurrentTime returns CurrentTime (property field)
	GetCurrentTime() int64
	// GetState returns State (property field)
	GetState() ServerState
	// GetBuildInfo returns BuildInfo (property field)
	GetBuildInfo() ExtensionObjectDefinition
	// GetSecondsTillShutdown returns SecondsTillShutdown (property field)
	GetSecondsTillShutdown() uint32
	// GetShutdownReason returns ShutdownReason (property field)
	GetShutdownReason() LocalizedText
	// IsServerStatusDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsServerStatusDataType()
	// CreateBuilder creates a ServerStatusDataTypeBuilder
	CreateServerStatusDataTypeBuilder() ServerStatusDataTypeBuilder
}

// _ServerStatusDataType is the data-structure of this message
type _ServerStatusDataType struct {
	ExtensionObjectDefinitionContract
	StartTime           int64
	CurrentTime         int64
	State               ServerState
	BuildInfo           ExtensionObjectDefinition
	SecondsTillShutdown uint32
	ShutdownReason      LocalizedText
}

var _ ServerStatusDataType = (*_ServerStatusDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ServerStatusDataType)(nil)

// NewServerStatusDataType factory function for _ServerStatusDataType
func NewServerStatusDataType(startTime int64, currentTime int64, state ServerState, buildInfo ExtensionObjectDefinition, secondsTillShutdown uint32, shutdownReason LocalizedText) *_ServerStatusDataType {
	if buildInfo == nil {
		panic("buildInfo of type ExtensionObjectDefinition for ServerStatusDataType must not be nil")
	}
	if shutdownReason == nil {
		panic("shutdownReason of type LocalizedText for ServerStatusDataType must not be nil")
	}
	_result := &_ServerStatusDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StartTime:                         startTime,
		CurrentTime:                       currentTime,
		State:                             state,
		BuildInfo:                         buildInfo,
		SecondsTillShutdown:               secondsTillShutdown,
		ShutdownReason:                    shutdownReason,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ServerStatusDataTypeBuilder is a builder for ServerStatusDataType
type ServerStatusDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(startTime int64, currentTime int64, state ServerState, buildInfo ExtensionObjectDefinition, secondsTillShutdown uint32, shutdownReason LocalizedText) ServerStatusDataTypeBuilder
	// WithStartTime adds StartTime (property field)
	WithStartTime(int64) ServerStatusDataTypeBuilder
	// WithCurrentTime adds CurrentTime (property field)
	WithCurrentTime(int64) ServerStatusDataTypeBuilder
	// WithState adds State (property field)
	WithState(ServerState) ServerStatusDataTypeBuilder
	// WithBuildInfo adds BuildInfo (property field)
	WithBuildInfo(ExtensionObjectDefinition) ServerStatusDataTypeBuilder
	// WithBuildInfoBuilder adds BuildInfo (property field) which is build by the builder
	WithBuildInfoBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) ServerStatusDataTypeBuilder
	// WithSecondsTillShutdown adds SecondsTillShutdown (property field)
	WithSecondsTillShutdown(uint32) ServerStatusDataTypeBuilder
	// WithShutdownReason adds ShutdownReason (property field)
	WithShutdownReason(LocalizedText) ServerStatusDataTypeBuilder
	// WithShutdownReasonBuilder adds ShutdownReason (property field) which is build by the builder
	WithShutdownReasonBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) ServerStatusDataTypeBuilder
	// Build builds the ServerStatusDataType or returns an error if something is wrong
	Build() (ServerStatusDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ServerStatusDataType
}

// NewServerStatusDataTypeBuilder() creates a ServerStatusDataTypeBuilder
func NewServerStatusDataTypeBuilder() ServerStatusDataTypeBuilder {
	return &_ServerStatusDataTypeBuilder{_ServerStatusDataType: new(_ServerStatusDataType)}
}

type _ServerStatusDataTypeBuilder struct {
	*_ServerStatusDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ServerStatusDataTypeBuilder) = (*_ServerStatusDataTypeBuilder)(nil)

func (b *_ServerStatusDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ServerStatusDataTypeBuilder) WithMandatoryFields(startTime int64, currentTime int64, state ServerState, buildInfo ExtensionObjectDefinition, secondsTillShutdown uint32, shutdownReason LocalizedText) ServerStatusDataTypeBuilder {
	return b.WithStartTime(startTime).WithCurrentTime(currentTime).WithState(state).WithBuildInfo(buildInfo).WithSecondsTillShutdown(secondsTillShutdown).WithShutdownReason(shutdownReason)
}

func (b *_ServerStatusDataTypeBuilder) WithStartTime(startTime int64) ServerStatusDataTypeBuilder {
	b.StartTime = startTime
	return b
}

func (b *_ServerStatusDataTypeBuilder) WithCurrentTime(currentTime int64) ServerStatusDataTypeBuilder {
	b.CurrentTime = currentTime
	return b
}

func (b *_ServerStatusDataTypeBuilder) WithState(state ServerState) ServerStatusDataTypeBuilder {
	b.State = state
	return b
}

func (b *_ServerStatusDataTypeBuilder) WithBuildInfo(buildInfo ExtensionObjectDefinition) ServerStatusDataTypeBuilder {
	b.BuildInfo = buildInfo
	return b
}

func (b *_ServerStatusDataTypeBuilder) WithBuildInfoBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) ServerStatusDataTypeBuilder {
	builder := builderSupplier(b.BuildInfo.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.BuildInfo, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_ServerStatusDataTypeBuilder) WithSecondsTillShutdown(secondsTillShutdown uint32) ServerStatusDataTypeBuilder {
	b.SecondsTillShutdown = secondsTillShutdown
	return b
}

func (b *_ServerStatusDataTypeBuilder) WithShutdownReason(shutdownReason LocalizedText) ServerStatusDataTypeBuilder {
	b.ShutdownReason = shutdownReason
	return b
}

func (b *_ServerStatusDataTypeBuilder) WithShutdownReasonBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) ServerStatusDataTypeBuilder {
	builder := builderSupplier(b.ShutdownReason.CreateLocalizedTextBuilder())
	var err error
	b.ShutdownReason, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_ServerStatusDataTypeBuilder) Build() (ServerStatusDataType, error) {
	if b.BuildInfo == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'buildInfo' not set"))
	}
	if b.ShutdownReason == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'shutdownReason' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ServerStatusDataType.deepCopy(), nil
}

func (b *_ServerStatusDataTypeBuilder) MustBuild() ServerStatusDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ServerStatusDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ServerStatusDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ServerStatusDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateServerStatusDataTypeBuilder().(*_ServerStatusDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateServerStatusDataTypeBuilder creates a ServerStatusDataTypeBuilder
func (b *_ServerStatusDataType) CreateServerStatusDataTypeBuilder() ServerStatusDataTypeBuilder {
	if b == nil {
		return NewServerStatusDataTypeBuilder()
	}
	return &_ServerStatusDataTypeBuilder{_ServerStatusDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ServerStatusDataType) GetIdentifier() string {
	return "864"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServerStatusDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ServerStatusDataType) GetStartTime() int64 {
	return m.StartTime
}

func (m *_ServerStatusDataType) GetCurrentTime() int64 {
	return m.CurrentTime
}

func (m *_ServerStatusDataType) GetState() ServerState {
	return m.State
}

func (m *_ServerStatusDataType) GetBuildInfo() ExtensionObjectDefinition {
	return m.BuildInfo
}

func (m *_ServerStatusDataType) GetSecondsTillShutdown() uint32 {
	return m.SecondsTillShutdown
}

func (m *_ServerStatusDataType) GetShutdownReason() LocalizedText {
	return m.ShutdownReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastServerStatusDataType(structType any) ServerStatusDataType {
	if casted, ok := structType.(ServerStatusDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ServerStatusDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ServerStatusDataType) GetTypeName() string {
	return "ServerStatusDataType"
}

func (m *_ServerStatusDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (startTime)
	lengthInBits += 64

	// Simple field (currentTime)
	lengthInBits += 64

	// Simple field (state)
	lengthInBits += 32

	// Simple field (buildInfo)
	lengthInBits += m.BuildInfo.GetLengthInBits(ctx)

	// Simple field (secondsTillShutdown)
	lengthInBits += 32

	// Simple field (shutdownReason)
	lengthInBits += m.ShutdownReason.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ServerStatusDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ServerStatusDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__serverStatusDataType ServerStatusDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ServerStatusDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServerStatusDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startTime, err := ReadSimpleField(ctx, "startTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startTime' field"))
	}
	m.StartTime = startTime

	currentTime, err := ReadSimpleField(ctx, "currentTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentTime' field"))
	}
	m.CurrentTime = currentTime

	state, err := ReadEnumField[ServerState](ctx, "state", "ServerState", ReadEnum(ServerStateByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'state' field"))
	}
	m.State = state

	buildInfo, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "buildInfo", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("340")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'buildInfo' field"))
	}
	m.BuildInfo = buildInfo

	secondsTillShutdown, err := ReadSimpleField(ctx, "secondsTillShutdown", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secondsTillShutdown' field"))
	}
	m.SecondsTillShutdown = secondsTillShutdown

	shutdownReason, err := ReadSimpleField[LocalizedText](ctx, "shutdownReason", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'shutdownReason' field"))
	}
	m.ShutdownReason = shutdownReason

	if closeErr := readBuffer.CloseContext("ServerStatusDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServerStatusDataType")
	}

	return m, nil
}

func (m *_ServerStatusDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServerStatusDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServerStatusDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServerStatusDataType")
		}

		if err := WriteSimpleField[int64](ctx, "startTime", m.GetStartTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'startTime' field")
		}

		if err := WriteSimpleField[int64](ctx, "currentTime", m.GetCurrentTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentTime' field")
		}

		if err := WriteSimpleEnumField[ServerState](ctx, "state", "ServerState", m.GetState(), WriteEnum[ServerState, uint32](ServerState.GetValue, ServerState.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'state' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "buildInfo", m.GetBuildInfo(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'buildInfo' field")
		}

		if err := WriteSimpleField[uint32](ctx, "secondsTillShutdown", m.GetSecondsTillShutdown(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'secondsTillShutdown' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "shutdownReason", m.GetShutdownReason(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'shutdownReason' field")
		}

		if popErr := writeBuffer.PopContext("ServerStatusDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServerStatusDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServerStatusDataType) IsServerStatusDataType() {}

func (m *_ServerStatusDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ServerStatusDataType) deepCopy() *_ServerStatusDataType {
	if m == nil {
		return nil
	}
	_ServerStatusDataTypeCopy := &_ServerStatusDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.StartTime,
		m.CurrentTime,
		m.State,
		m.BuildInfo.DeepCopy().(ExtensionObjectDefinition),
		m.SecondsTillShutdown,
		m.ShutdownReason.DeepCopy().(LocalizedText),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ServerStatusDataTypeCopy
}

func (m *_ServerStatusDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

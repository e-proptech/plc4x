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

// AggregateFilter is the corresponding interface of AggregateFilter
type AggregateFilter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStartTime returns StartTime (property field)
	GetStartTime() int64
	// GetAggregateType returns AggregateType (property field)
	GetAggregateType() NodeId
	// GetProcessingInterval returns ProcessingInterval (property field)
	GetProcessingInterval() float64
	// GetAggregateConfiguration returns AggregateConfiguration (property field)
	GetAggregateConfiguration() AggregateConfiguration
	// IsAggregateFilter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAggregateFilter()
	// CreateBuilder creates a AggregateFilterBuilder
	CreateAggregateFilterBuilder() AggregateFilterBuilder
}

// _AggregateFilter is the data-structure of this message
type _AggregateFilter struct {
	ExtensionObjectDefinitionContract
	StartTime              int64
	AggregateType          NodeId
	ProcessingInterval     float64
	AggregateConfiguration AggregateConfiguration
}

var _ AggregateFilter = (*_AggregateFilter)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AggregateFilter)(nil)

// NewAggregateFilter factory function for _AggregateFilter
func NewAggregateFilter(startTime int64, aggregateType NodeId, processingInterval float64, aggregateConfiguration AggregateConfiguration) *_AggregateFilter {
	if aggregateType == nil {
		panic("aggregateType of type NodeId for AggregateFilter must not be nil")
	}
	if aggregateConfiguration == nil {
		panic("aggregateConfiguration of type AggregateConfiguration for AggregateFilter must not be nil")
	}
	_result := &_AggregateFilter{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StartTime:                         startTime,
		AggregateType:                     aggregateType,
		ProcessingInterval:                processingInterval,
		AggregateConfiguration:            aggregateConfiguration,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AggregateFilterBuilder is a builder for AggregateFilter
type AggregateFilterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(startTime int64, aggregateType NodeId, processingInterval float64, aggregateConfiguration AggregateConfiguration) AggregateFilterBuilder
	// WithStartTime adds StartTime (property field)
	WithStartTime(int64) AggregateFilterBuilder
	// WithAggregateType adds AggregateType (property field)
	WithAggregateType(NodeId) AggregateFilterBuilder
	// WithAggregateTypeBuilder adds AggregateType (property field) which is build by the builder
	WithAggregateTypeBuilder(func(NodeIdBuilder) NodeIdBuilder) AggregateFilterBuilder
	// WithProcessingInterval adds ProcessingInterval (property field)
	WithProcessingInterval(float64) AggregateFilterBuilder
	// WithAggregateConfiguration adds AggregateConfiguration (property field)
	WithAggregateConfiguration(AggregateConfiguration) AggregateFilterBuilder
	// WithAggregateConfigurationBuilder adds AggregateConfiguration (property field) which is build by the builder
	WithAggregateConfigurationBuilder(func(AggregateConfigurationBuilder) AggregateConfigurationBuilder) AggregateFilterBuilder
	// Build builds the AggregateFilter or returns an error if something is wrong
	Build() (AggregateFilter, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AggregateFilter
}

// NewAggregateFilterBuilder() creates a AggregateFilterBuilder
func NewAggregateFilterBuilder() AggregateFilterBuilder {
	return &_AggregateFilterBuilder{_AggregateFilter: new(_AggregateFilter)}
}

type _AggregateFilterBuilder struct {
	*_AggregateFilter

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (AggregateFilterBuilder) = (*_AggregateFilterBuilder)(nil)

func (b *_AggregateFilterBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_AggregateFilterBuilder) WithMandatoryFields(startTime int64, aggregateType NodeId, processingInterval float64, aggregateConfiguration AggregateConfiguration) AggregateFilterBuilder {
	return b.WithStartTime(startTime).WithAggregateType(aggregateType).WithProcessingInterval(processingInterval).WithAggregateConfiguration(aggregateConfiguration)
}

func (b *_AggregateFilterBuilder) WithStartTime(startTime int64) AggregateFilterBuilder {
	b.StartTime = startTime
	return b
}

func (b *_AggregateFilterBuilder) WithAggregateType(aggregateType NodeId) AggregateFilterBuilder {
	b.AggregateType = aggregateType
	return b
}

func (b *_AggregateFilterBuilder) WithAggregateTypeBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) AggregateFilterBuilder {
	builder := builderSupplier(b.AggregateType.CreateNodeIdBuilder())
	var err error
	b.AggregateType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_AggregateFilterBuilder) WithProcessingInterval(processingInterval float64) AggregateFilterBuilder {
	b.ProcessingInterval = processingInterval
	return b
}

func (b *_AggregateFilterBuilder) WithAggregateConfiguration(aggregateConfiguration AggregateConfiguration) AggregateFilterBuilder {
	b.AggregateConfiguration = aggregateConfiguration
	return b
}

func (b *_AggregateFilterBuilder) WithAggregateConfigurationBuilder(builderSupplier func(AggregateConfigurationBuilder) AggregateConfigurationBuilder) AggregateFilterBuilder {
	builder := builderSupplier(b.AggregateConfiguration.CreateAggregateConfigurationBuilder())
	var err error
	b.AggregateConfiguration, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "AggregateConfigurationBuilder failed"))
	}
	return b
}

func (b *_AggregateFilterBuilder) Build() (AggregateFilter, error) {
	if b.AggregateType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'aggregateType' not set"))
	}
	if b.AggregateConfiguration == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'aggregateConfiguration' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AggregateFilter.deepCopy(), nil
}

func (b *_AggregateFilterBuilder) MustBuild() AggregateFilter {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AggregateFilterBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_AggregateFilterBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_AggregateFilterBuilder) DeepCopy() any {
	_copy := b.CreateAggregateFilterBuilder().(*_AggregateFilterBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAggregateFilterBuilder creates a AggregateFilterBuilder
func (b *_AggregateFilter) CreateAggregateFilterBuilder() AggregateFilterBuilder {
	if b == nil {
		return NewAggregateFilterBuilder()
	}
	return &_AggregateFilterBuilder{_AggregateFilter: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AggregateFilter) GetExtensionId() int32 {
	return int32(730)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AggregateFilter) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AggregateFilter) GetStartTime() int64 {
	return m.StartTime
}

func (m *_AggregateFilter) GetAggregateType() NodeId {
	return m.AggregateType
}

func (m *_AggregateFilter) GetProcessingInterval() float64 {
	return m.ProcessingInterval
}

func (m *_AggregateFilter) GetAggregateConfiguration() AggregateConfiguration {
	return m.AggregateConfiguration
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAggregateFilter(structType any) AggregateFilter {
	if casted, ok := structType.(AggregateFilter); ok {
		return casted
	}
	if casted, ok := structType.(*AggregateFilter); ok {
		return *casted
	}
	return nil
}

func (m *_AggregateFilter) GetTypeName() string {
	return "AggregateFilter"
}

func (m *_AggregateFilter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (startTime)
	lengthInBits += 64

	// Simple field (aggregateType)
	lengthInBits += m.AggregateType.GetLengthInBits(ctx)

	// Simple field (processingInterval)
	lengthInBits += 64

	// Simple field (aggregateConfiguration)
	lengthInBits += m.AggregateConfiguration.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AggregateFilter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AggregateFilter) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__aggregateFilter AggregateFilter, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AggregateFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AggregateFilter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startTime, err := ReadSimpleField(ctx, "startTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startTime' field"))
	}
	m.StartTime = startTime

	aggregateType, err := ReadSimpleField[NodeId](ctx, "aggregateType", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'aggregateType' field"))
	}
	m.AggregateType = aggregateType

	processingInterval, err := ReadSimpleField(ctx, "processingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processingInterval' field"))
	}
	m.ProcessingInterval = processingInterval

	aggregateConfiguration, err := ReadSimpleField[AggregateConfiguration](ctx, "aggregateConfiguration", ReadComplex[AggregateConfiguration](ExtensionObjectDefinitionParseWithBufferProducer[AggregateConfiguration]((int32)(int32(950))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'aggregateConfiguration' field"))
	}
	m.AggregateConfiguration = aggregateConfiguration

	if closeErr := readBuffer.CloseContext("AggregateFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AggregateFilter")
	}

	return m, nil
}

func (m *_AggregateFilter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AggregateFilter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AggregateFilter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AggregateFilter")
		}

		if err := WriteSimpleField[int64](ctx, "startTime", m.GetStartTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'startTime' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "aggregateType", m.GetAggregateType(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'aggregateType' field")
		}

		if err := WriteSimpleField[float64](ctx, "processingInterval", m.GetProcessingInterval(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'processingInterval' field")
		}

		if err := WriteSimpleField[AggregateConfiguration](ctx, "aggregateConfiguration", m.GetAggregateConfiguration(), WriteComplex[AggregateConfiguration](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'aggregateConfiguration' field")
		}

		if popErr := writeBuffer.PopContext("AggregateFilter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AggregateFilter")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AggregateFilter) IsAggregateFilter() {}

func (m *_AggregateFilter) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AggregateFilter) deepCopy() *_AggregateFilter {
	if m == nil {
		return nil
	}
	_AggregateFilterCopy := &_AggregateFilter{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.StartTime,
		utils.DeepCopy[NodeId](m.AggregateType),
		m.ProcessingInterval,
		utils.DeepCopy[AggregateConfiguration](m.AggregateConfiguration),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _AggregateFilterCopy
}

func (m *_AggregateFilter) String() string {
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

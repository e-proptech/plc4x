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

// LogicalSegmentType is the corresponding interface of LogicalSegmentType
type LogicalSegmentType interface {
	LogicalSegmentTypeContract
	LogicalSegmentTypeRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsLogicalSegmentType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLogicalSegmentType()
	// CreateBuilder creates a LogicalSegmentTypeBuilder
	CreateLogicalSegmentTypeBuilder() LogicalSegmentTypeBuilder
}

// LogicalSegmentTypeContract provides a set of functions which can be overwritten by a sub struct
type LogicalSegmentTypeContract interface {
	// IsLogicalSegmentType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLogicalSegmentType()
	// CreateBuilder creates a LogicalSegmentTypeBuilder
	CreateLogicalSegmentTypeBuilder() LogicalSegmentTypeBuilder
}

// LogicalSegmentTypeRequirements provides a set of functions which need to be implemented by a sub struct
type LogicalSegmentTypeRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetLogicalSegmentType returns LogicalSegmentType (discriminator field)
	GetLogicalSegmentType() uint8
}

// _LogicalSegmentType is the data-structure of this message
type _LogicalSegmentType struct {
	_SubType LogicalSegmentType
}

var _ LogicalSegmentTypeContract = (*_LogicalSegmentType)(nil)

// NewLogicalSegmentType factory function for _LogicalSegmentType
func NewLogicalSegmentType() *_LogicalSegmentType {
	return &_LogicalSegmentType{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LogicalSegmentTypeBuilder is a builder for LogicalSegmentType
type LogicalSegmentTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() LogicalSegmentTypeBuilder
	// AsClassID converts this build to a subType of LogicalSegmentType. It is always possible to return to current builder using Done()
	AsClassID() interface {
		ClassIDBuilder
		Done() LogicalSegmentTypeBuilder
	}
	// AsInstanceID converts this build to a subType of LogicalSegmentType. It is always possible to return to current builder using Done()
	AsInstanceID() interface {
		InstanceIDBuilder
		Done() LogicalSegmentTypeBuilder
	}
	// AsMemberID converts this build to a subType of LogicalSegmentType. It is always possible to return to current builder using Done()
	AsMemberID() interface {
		MemberIDBuilder
		Done() LogicalSegmentTypeBuilder
	}
	// Build builds the LogicalSegmentType or returns an error if something is wrong
	PartialBuild() (LogicalSegmentTypeContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() LogicalSegmentTypeContract
	// Build builds the LogicalSegmentType or returns an error if something is wrong
	Build() (LogicalSegmentType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LogicalSegmentType
}

// NewLogicalSegmentTypeBuilder() creates a LogicalSegmentTypeBuilder
func NewLogicalSegmentTypeBuilder() LogicalSegmentTypeBuilder {
	return &_LogicalSegmentTypeBuilder{_LogicalSegmentType: new(_LogicalSegmentType)}
}

type _LogicalSegmentTypeChildBuilder interface {
	utils.Copyable
	setParent(LogicalSegmentTypeContract)
	buildForLogicalSegmentType() (LogicalSegmentType, error)
}

type _LogicalSegmentTypeBuilder struct {
	*_LogicalSegmentType

	childBuilder _LogicalSegmentTypeChildBuilder

	err *utils.MultiError
}

var _ (LogicalSegmentTypeBuilder) = (*_LogicalSegmentTypeBuilder)(nil)

func (b *_LogicalSegmentTypeBuilder) WithMandatoryFields() LogicalSegmentTypeBuilder {
	return b
}

func (b *_LogicalSegmentTypeBuilder) PartialBuild() (LogicalSegmentTypeContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LogicalSegmentType.deepCopy(), nil
}

func (b *_LogicalSegmentTypeBuilder) PartialMustBuild() LogicalSegmentTypeContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LogicalSegmentTypeBuilder) AsClassID() interface {
	ClassIDBuilder
	Done() LogicalSegmentTypeBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		ClassIDBuilder
		Done() LogicalSegmentTypeBuilder
	}); ok {
		return cb
	}
	cb := NewClassIDBuilder().(*_ClassIDBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_LogicalSegmentTypeBuilder) AsInstanceID() interface {
	InstanceIDBuilder
	Done() LogicalSegmentTypeBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		InstanceIDBuilder
		Done() LogicalSegmentTypeBuilder
	}); ok {
		return cb
	}
	cb := NewInstanceIDBuilder().(*_InstanceIDBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_LogicalSegmentTypeBuilder) AsMemberID() interface {
	MemberIDBuilder
	Done() LogicalSegmentTypeBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		MemberIDBuilder
		Done() LogicalSegmentTypeBuilder
	}); ok {
		return cb
	}
	cb := NewMemberIDBuilder().(*_MemberIDBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_LogicalSegmentTypeBuilder) Build() (LogicalSegmentType, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForLogicalSegmentType()
}

func (b *_LogicalSegmentTypeBuilder) MustBuild() LogicalSegmentType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LogicalSegmentTypeBuilder) DeepCopy() any {
	_copy := b.CreateLogicalSegmentTypeBuilder().(*_LogicalSegmentTypeBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_LogicalSegmentTypeChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLogicalSegmentTypeBuilder creates a LogicalSegmentTypeBuilder
func (b *_LogicalSegmentType) CreateLogicalSegmentTypeBuilder() LogicalSegmentTypeBuilder {
	if b == nil {
		return NewLogicalSegmentTypeBuilder()
	}
	return &_LogicalSegmentTypeBuilder{_LogicalSegmentType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLogicalSegmentType(structType any) LogicalSegmentType {
	if casted, ok := structType.(LogicalSegmentType); ok {
		return casted
	}
	if casted, ok := structType.(*LogicalSegmentType); ok {
		return *casted
	}
	return nil
}

func (m *_LogicalSegmentType) GetTypeName() string {
	return "LogicalSegmentType"
}

func (m *_LogicalSegmentType) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (logicalSegmentType)
	lengthInBits += 3

	return lengthInBits
}

func (m *_LogicalSegmentType) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func LogicalSegmentTypeParse[T LogicalSegmentType](ctx context.Context, theBytes []byte) (T, error) {
	return LogicalSegmentTypeParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func LogicalSegmentTypeParseWithBufferProducer[T LogicalSegmentType]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := LogicalSegmentTypeParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func LogicalSegmentTypeParseWithBuffer[T LogicalSegmentType](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_LogicalSegmentType{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_LogicalSegmentType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__logicalSegmentType LogicalSegmentType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LogicalSegmentType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LogicalSegmentType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	logicalSegmentType, err := ReadDiscriminatorField[uint8](ctx, "logicalSegmentType", ReadUnsignedByte(readBuffer, uint8(3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logicalSegmentType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child LogicalSegmentType
	switch {
	case logicalSegmentType == 0x00: // ClassID
		if _child, err = new(_ClassID).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ClassID for type-switch of LogicalSegmentType")
		}
	case logicalSegmentType == 0x01: // InstanceID
		if _child, err = new(_InstanceID).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type InstanceID for type-switch of LogicalSegmentType")
		}
	case logicalSegmentType == 0x02: // MemberID
		if _child, err = new(_MemberID).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MemberID for type-switch of LogicalSegmentType")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [logicalSegmentType=%v]", logicalSegmentType)
	}

	if closeErr := readBuffer.CloseContext("LogicalSegmentType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LogicalSegmentType")
	}

	return _child, nil
}

func (pm *_LogicalSegmentType) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child LogicalSegmentType, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("LogicalSegmentType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LogicalSegmentType")
	}

	if err := WriteDiscriminatorField(ctx, "logicalSegmentType", m.GetLogicalSegmentType(), WriteUnsignedByte(writeBuffer, 3)); err != nil {
		return errors.Wrap(err, "Error serializing 'logicalSegmentType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("LogicalSegmentType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LogicalSegmentType")
	}
	return nil
}

func (m *_LogicalSegmentType) IsLogicalSegmentType() {}

func (m *_LogicalSegmentType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LogicalSegmentType) deepCopy() *_LogicalSegmentType {
	if m == nil {
		return nil
	}
	_LogicalSegmentTypeCopy := &_LogicalSegmentType{
		nil, // will be set by child
	}
	return _LogicalSegmentTypeCopy
}

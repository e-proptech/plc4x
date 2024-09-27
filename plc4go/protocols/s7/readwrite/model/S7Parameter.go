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

// S7Parameter is the corresponding interface of S7Parameter
type S7Parameter interface {
	S7ParameterContract
	S7ParameterRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsS7Parameter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7Parameter()
	// CreateBuilder creates a S7ParameterBuilder
	CreateS7ParameterBuilder() S7ParameterBuilder
}

// S7ParameterContract provides a set of functions which can be overwritten by a sub struct
type S7ParameterContract interface {
	// IsS7Parameter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7Parameter()
	// CreateBuilder creates a S7ParameterBuilder
	CreateS7ParameterBuilder() S7ParameterBuilder
}

// S7ParameterRequirements provides a set of functions which need to be implemented by a sub struct
type S7ParameterRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetParameterType returns ParameterType (discriminator field)
	GetParameterType() uint8
}

// _S7Parameter is the data-structure of this message
type _S7Parameter struct {
	_SubType S7Parameter
}

var _ S7ParameterContract = (*_S7Parameter)(nil)

// NewS7Parameter factory function for _S7Parameter
func NewS7Parameter() *_S7Parameter {
	return &_S7Parameter{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7ParameterBuilder is a builder for S7Parameter
type S7ParameterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() S7ParameterBuilder
	// AsS7ParameterSetupCommunication converts this build to a subType of S7Parameter. It is always possible to return to current builder using Done()
	AsS7ParameterSetupCommunication() interface {
		S7ParameterSetupCommunicationBuilder
		Done() S7ParameterBuilder
	}
	// AsS7ParameterReadVarRequest converts this build to a subType of S7Parameter. It is always possible to return to current builder using Done()
	AsS7ParameterReadVarRequest() interface {
		S7ParameterReadVarRequestBuilder
		Done() S7ParameterBuilder
	}
	// AsS7ParameterReadVarResponse converts this build to a subType of S7Parameter. It is always possible to return to current builder using Done()
	AsS7ParameterReadVarResponse() interface {
		S7ParameterReadVarResponseBuilder
		Done() S7ParameterBuilder
	}
	// AsS7ParameterWriteVarRequest converts this build to a subType of S7Parameter. It is always possible to return to current builder using Done()
	AsS7ParameterWriteVarRequest() interface {
		S7ParameterWriteVarRequestBuilder
		Done() S7ParameterBuilder
	}
	// AsS7ParameterWriteVarResponse converts this build to a subType of S7Parameter. It is always possible to return to current builder using Done()
	AsS7ParameterWriteVarResponse() interface {
		S7ParameterWriteVarResponseBuilder
		Done() S7ParameterBuilder
	}
	// AsS7ParameterUserData converts this build to a subType of S7Parameter. It is always possible to return to current builder using Done()
	AsS7ParameterUserData() interface {
		S7ParameterUserDataBuilder
		Done() S7ParameterBuilder
	}
	// AsS7ParameterModeTransition converts this build to a subType of S7Parameter. It is always possible to return to current builder using Done()
	AsS7ParameterModeTransition() interface {
		S7ParameterModeTransitionBuilder
		Done() S7ParameterBuilder
	}
	// Build builds the S7Parameter or returns an error if something is wrong
	PartialBuild() (S7ParameterContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() S7ParameterContract
	// Build builds the S7Parameter or returns an error if something is wrong
	Build() (S7Parameter, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7Parameter
}

// NewS7ParameterBuilder() creates a S7ParameterBuilder
func NewS7ParameterBuilder() S7ParameterBuilder {
	return &_S7ParameterBuilder{_S7Parameter: new(_S7Parameter)}
}

type _S7ParameterChildBuilder interface {
	utils.Copyable
	setParent(S7ParameterContract)
	buildForS7Parameter() (S7Parameter, error)
}

type _S7ParameterBuilder struct {
	*_S7Parameter

	childBuilder _S7ParameterChildBuilder

	err *utils.MultiError
}

var _ (S7ParameterBuilder) = (*_S7ParameterBuilder)(nil)

func (b *_S7ParameterBuilder) WithMandatoryFields() S7ParameterBuilder {
	return b
}

func (b *_S7ParameterBuilder) PartialBuild() (S7ParameterContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7Parameter.deepCopy(), nil
}

func (b *_S7ParameterBuilder) PartialMustBuild() S7ParameterContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7ParameterBuilder) AsS7ParameterSetupCommunication() interface {
	S7ParameterSetupCommunicationBuilder
	Done() S7ParameterBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7ParameterSetupCommunicationBuilder
		Done() S7ParameterBuilder
	}); ok {
		return cb
	}
	cb := NewS7ParameterSetupCommunicationBuilder().(*_S7ParameterSetupCommunicationBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7ParameterBuilder) AsS7ParameterReadVarRequest() interface {
	S7ParameterReadVarRequestBuilder
	Done() S7ParameterBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7ParameterReadVarRequestBuilder
		Done() S7ParameterBuilder
	}); ok {
		return cb
	}
	cb := NewS7ParameterReadVarRequestBuilder().(*_S7ParameterReadVarRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7ParameterBuilder) AsS7ParameterReadVarResponse() interface {
	S7ParameterReadVarResponseBuilder
	Done() S7ParameterBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7ParameterReadVarResponseBuilder
		Done() S7ParameterBuilder
	}); ok {
		return cb
	}
	cb := NewS7ParameterReadVarResponseBuilder().(*_S7ParameterReadVarResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7ParameterBuilder) AsS7ParameterWriteVarRequest() interface {
	S7ParameterWriteVarRequestBuilder
	Done() S7ParameterBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7ParameterWriteVarRequestBuilder
		Done() S7ParameterBuilder
	}); ok {
		return cb
	}
	cb := NewS7ParameterWriteVarRequestBuilder().(*_S7ParameterWriteVarRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7ParameterBuilder) AsS7ParameterWriteVarResponse() interface {
	S7ParameterWriteVarResponseBuilder
	Done() S7ParameterBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7ParameterWriteVarResponseBuilder
		Done() S7ParameterBuilder
	}); ok {
		return cb
	}
	cb := NewS7ParameterWriteVarResponseBuilder().(*_S7ParameterWriteVarResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7ParameterBuilder) AsS7ParameterUserData() interface {
	S7ParameterUserDataBuilder
	Done() S7ParameterBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7ParameterUserDataBuilder
		Done() S7ParameterBuilder
	}); ok {
		return cb
	}
	cb := NewS7ParameterUserDataBuilder().(*_S7ParameterUserDataBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7ParameterBuilder) AsS7ParameterModeTransition() interface {
	S7ParameterModeTransitionBuilder
	Done() S7ParameterBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7ParameterModeTransitionBuilder
		Done() S7ParameterBuilder
	}); ok {
		return cb
	}
	cb := NewS7ParameterModeTransitionBuilder().(*_S7ParameterModeTransitionBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7ParameterBuilder) Build() (S7Parameter, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForS7Parameter()
}

func (b *_S7ParameterBuilder) MustBuild() S7Parameter {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7ParameterBuilder) DeepCopy() any {
	_copy := b.CreateS7ParameterBuilder().(*_S7ParameterBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_S7ParameterChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7ParameterBuilder creates a S7ParameterBuilder
func (b *_S7Parameter) CreateS7ParameterBuilder() S7ParameterBuilder {
	if b == nil {
		return NewS7ParameterBuilder()
	}
	return &_S7ParameterBuilder{_S7Parameter: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7Parameter(structType any) S7Parameter {
	if casted, ok := structType.(S7Parameter); ok {
		return casted
	}
	if casted, ok := structType.(*S7Parameter); ok {
		return *casted
	}
	return nil
}

func (m *_S7Parameter) GetTypeName() string {
	return "S7Parameter"
}

func (m *_S7Parameter) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (parameterType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7Parameter) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func S7ParameterParse[T S7Parameter](ctx context.Context, theBytes []byte, messageType uint8) (T, error) {
	return S7ParameterParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), messageType)
}

func S7ParameterParseWithBufferProducer[T S7Parameter](messageType uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := S7ParameterParseWithBuffer[T](ctx, readBuffer, messageType)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func S7ParameterParseWithBuffer[T S7Parameter](ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8) (T, error) {
	v, err := (&_S7Parameter{}).parse(ctx, readBuffer, messageType)
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

func (m *_S7Parameter) parse(ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8) (__s7Parameter S7Parameter, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7Parameter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7Parameter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	parameterType, err := ReadDiscriminatorField[uint8](ctx, "parameterType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameterType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child S7Parameter
	switch {
	case parameterType == 0xF0: // S7ParameterSetupCommunication
		if _child, err = new(_S7ParameterSetupCommunication).parse(ctx, readBuffer, m, messageType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7ParameterSetupCommunication for type-switch of S7Parameter")
		}
	case parameterType == 0x04 && messageType == 0x01: // S7ParameterReadVarRequest
		if _child, err = new(_S7ParameterReadVarRequest).parse(ctx, readBuffer, m, messageType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7ParameterReadVarRequest for type-switch of S7Parameter")
		}
	case parameterType == 0x04 && messageType == 0x03: // S7ParameterReadVarResponse
		if _child, err = new(_S7ParameterReadVarResponse).parse(ctx, readBuffer, m, messageType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7ParameterReadVarResponse for type-switch of S7Parameter")
		}
	case parameterType == 0x05 && messageType == 0x01: // S7ParameterWriteVarRequest
		if _child, err = new(_S7ParameterWriteVarRequest).parse(ctx, readBuffer, m, messageType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7ParameterWriteVarRequest for type-switch of S7Parameter")
		}
	case parameterType == 0x05 && messageType == 0x03: // S7ParameterWriteVarResponse
		if _child, err = new(_S7ParameterWriteVarResponse).parse(ctx, readBuffer, m, messageType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7ParameterWriteVarResponse for type-switch of S7Parameter")
		}
	case parameterType == 0x00 && messageType == 0x07: // S7ParameterUserData
		if _child, err = new(_S7ParameterUserData).parse(ctx, readBuffer, m, messageType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7ParameterUserData for type-switch of S7Parameter")
		}
	case parameterType == 0x01 && messageType == 0x07: // S7ParameterModeTransition
		if _child, err = new(_S7ParameterModeTransition).parse(ctx, readBuffer, m, messageType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7ParameterModeTransition for type-switch of S7Parameter")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [parameterType=%v, messageType=%v]", parameterType, messageType)
	}

	if closeErr := readBuffer.CloseContext("S7Parameter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7Parameter")
	}

	return _child, nil
}

func (pm *_S7Parameter) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7Parameter, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("S7Parameter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7Parameter")
	}

	if err := WriteDiscriminatorField(ctx, "parameterType", m.GetParameterType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'parameterType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Parameter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7Parameter")
	}
	return nil
}

func (m *_S7Parameter) IsS7Parameter() {}

func (m *_S7Parameter) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7Parameter) deepCopy() *_S7Parameter {
	if m == nil {
		return nil
	}
	_S7ParameterCopy := &_S7Parameter{
		nil, // will be set by child
	}
	return _S7ParameterCopy
}

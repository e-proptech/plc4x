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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SessionAuthenticationToken is the corresponding interface of SessionAuthenticationToken
type SessionAuthenticationToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsSessionAuthenticationToken is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSessionAuthenticationToken()
	// CreateBuilder creates a SessionAuthenticationTokenBuilder
	CreateSessionAuthenticationTokenBuilder() SessionAuthenticationTokenBuilder
}

// _SessionAuthenticationToken is the data-structure of this message
type _SessionAuthenticationToken struct {
}

var _ SessionAuthenticationToken = (*_SessionAuthenticationToken)(nil)

// NewSessionAuthenticationToken factory function for _SessionAuthenticationToken
func NewSessionAuthenticationToken() *_SessionAuthenticationToken {
	return &_SessionAuthenticationToken{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SessionAuthenticationTokenBuilder is a builder for SessionAuthenticationToken
type SessionAuthenticationTokenBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SessionAuthenticationTokenBuilder
	// Build builds the SessionAuthenticationToken or returns an error if something is wrong
	Build() (SessionAuthenticationToken, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SessionAuthenticationToken
}

// NewSessionAuthenticationTokenBuilder() creates a SessionAuthenticationTokenBuilder
func NewSessionAuthenticationTokenBuilder() SessionAuthenticationTokenBuilder {
	return &_SessionAuthenticationTokenBuilder{_SessionAuthenticationToken: new(_SessionAuthenticationToken)}
}

type _SessionAuthenticationTokenBuilder struct {
	*_SessionAuthenticationToken

	err *utils.MultiError
}

var _ (SessionAuthenticationTokenBuilder) = (*_SessionAuthenticationTokenBuilder)(nil)

func (b *_SessionAuthenticationTokenBuilder) WithMandatoryFields() SessionAuthenticationTokenBuilder {
	return b
}

func (b *_SessionAuthenticationTokenBuilder) Build() (SessionAuthenticationToken, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SessionAuthenticationToken.deepCopy(), nil
}

func (b *_SessionAuthenticationTokenBuilder) MustBuild() SessionAuthenticationToken {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SessionAuthenticationTokenBuilder) DeepCopy() any {
	_copy := b.CreateSessionAuthenticationTokenBuilder().(*_SessionAuthenticationTokenBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSessionAuthenticationTokenBuilder creates a SessionAuthenticationTokenBuilder
func (b *_SessionAuthenticationToken) CreateSessionAuthenticationTokenBuilder() SessionAuthenticationTokenBuilder {
	if b == nil {
		return NewSessionAuthenticationTokenBuilder()
	}
	return &_SessionAuthenticationTokenBuilder{_SessionAuthenticationToken: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSessionAuthenticationToken(structType any) SessionAuthenticationToken {
	if casted, ok := structType.(SessionAuthenticationToken); ok {
		return casted
	}
	if casted, ok := structType.(*SessionAuthenticationToken); ok {
		return *casted
	}
	return nil
}

func (m *_SessionAuthenticationToken) GetTypeName() string {
	return "SessionAuthenticationToken"
}

func (m *_SessionAuthenticationToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_SessionAuthenticationToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SessionAuthenticationTokenParse(ctx context.Context, theBytes []byte) (SessionAuthenticationToken, error) {
	return SessionAuthenticationTokenParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SessionAuthenticationTokenParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SessionAuthenticationToken, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SessionAuthenticationToken, error) {
		return SessionAuthenticationTokenParseWithBuffer(ctx, readBuffer)
	}
}

func SessionAuthenticationTokenParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SessionAuthenticationToken, error) {
	v, err := (&_SessionAuthenticationToken{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_SessionAuthenticationToken) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__sessionAuthenticationToken SessionAuthenticationToken, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SessionAuthenticationToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SessionAuthenticationToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SessionAuthenticationToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SessionAuthenticationToken")
	}

	return m, nil
}

func (m *_SessionAuthenticationToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SessionAuthenticationToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SessionAuthenticationToken"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SessionAuthenticationToken")
	}

	if popErr := writeBuffer.PopContext("SessionAuthenticationToken"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SessionAuthenticationToken")
	}
	return nil
}

func (m *_SessionAuthenticationToken) IsSessionAuthenticationToken() {}

func (m *_SessionAuthenticationToken) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SessionAuthenticationToken) deepCopy() *_SessionAuthenticationToken {
	if m == nil {
		return nil
	}
	_SessionAuthenticationTokenCopy := &_SessionAuthenticationToken{}
	return _SessionAuthenticationTokenCopy
}

func (m *_SessionAuthenticationToken) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

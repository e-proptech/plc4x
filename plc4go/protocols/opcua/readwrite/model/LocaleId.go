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

// LocaleId is the corresponding interface of LocaleId
type LocaleId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsLocaleId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLocaleId()
	// CreateBuilder creates a LocaleIdBuilder
	CreateLocaleIdBuilder() LocaleIdBuilder
}

// _LocaleId is the data-structure of this message
type _LocaleId struct {
}

var _ LocaleId = (*_LocaleId)(nil)

// NewLocaleId factory function for _LocaleId
func NewLocaleId() *_LocaleId {
	return &_LocaleId{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LocaleIdBuilder is a builder for LocaleId
type LocaleIdBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() LocaleIdBuilder
	// Build builds the LocaleId or returns an error if something is wrong
	Build() (LocaleId, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LocaleId
}

// NewLocaleIdBuilder() creates a LocaleIdBuilder
func NewLocaleIdBuilder() LocaleIdBuilder {
	return &_LocaleIdBuilder{_LocaleId: new(_LocaleId)}
}

type _LocaleIdBuilder struct {
	*_LocaleId

	err *utils.MultiError
}

var _ (LocaleIdBuilder) = (*_LocaleIdBuilder)(nil)

func (b *_LocaleIdBuilder) WithMandatoryFields() LocaleIdBuilder {
	return b
}

func (b *_LocaleIdBuilder) Build() (LocaleId, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LocaleId.deepCopy(), nil
}

func (b *_LocaleIdBuilder) MustBuild() LocaleId {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LocaleIdBuilder) DeepCopy() any {
	_copy := b.CreateLocaleIdBuilder().(*_LocaleIdBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLocaleIdBuilder creates a LocaleIdBuilder
func (b *_LocaleId) CreateLocaleIdBuilder() LocaleIdBuilder {
	if b == nil {
		return NewLocaleIdBuilder()
	}
	return &_LocaleIdBuilder{_LocaleId: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLocaleId(structType any) LocaleId {
	if casted, ok := structType.(LocaleId); ok {
		return casted
	}
	if casted, ok := structType.(*LocaleId); ok {
		return *casted
	}
	return nil
}

func (m *_LocaleId) GetTypeName() string {
	return "LocaleId"
}

func (m *_LocaleId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_LocaleId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LocaleIdParse(ctx context.Context, theBytes []byte) (LocaleId, error) {
	return LocaleIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LocaleIdParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (LocaleId, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (LocaleId, error) {
		return LocaleIdParseWithBuffer(ctx, readBuffer)
	}
}

func LocaleIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LocaleId, error) {
	v, err := (&_LocaleId{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_LocaleId) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__localeId LocaleId, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LocaleId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LocaleId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LocaleId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LocaleId")
	}

	return m, nil
}

func (m *_LocaleId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LocaleId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("LocaleId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LocaleId")
	}

	if popErr := writeBuffer.PopContext("LocaleId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LocaleId")
	}
	return nil
}

func (m *_LocaleId) IsLocaleId() {}

func (m *_LocaleId) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LocaleId) deepCopy() *_LocaleId {
	if m == nil {
		return nil
	}
	_LocaleIdCopy := &_LocaleId{}
	return _LocaleIdCopy
}

func (m *_LocaleId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

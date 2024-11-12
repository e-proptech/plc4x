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

// BACnetAccessThreatLevel is the corresponding interface of BACnetAccessThreatLevel
type BACnetAccessThreatLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetThreatLevel returns ThreatLevel (property field)
	GetThreatLevel() BACnetApplicationTagUnsignedInteger
	// IsBACnetAccessThreatLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccessThreatLevel()
	// CreateBuilder creates a BACnetAccessThreatLevelBuilder
	CreateBACnetAccessThreatLevelBuilder() BACnetAccessThreatLevelBuilder
}

// _BACnetAccessThreatLevel is the data-structure of this message
type _BACnetAccessThreatLevel struct {
	ThreatLevel BACnetApplicationTagUnsignedInteger
}

var _ BACnetAccessThreatLevel = (*_BACnetAccessThreatLevel)(nil)

// NewBACnetAccessThreatLevel factory function for _BACnetAccessThreatLevel
func NewBACnetAccessThreatLevel(threatLevel BACnetApplicationTagUnsignedInteger) *_BACnetAccessThreatLevel {
	if threatLevel == nil {
		panic("threatLevel of type BACnetApplicationTagUnsignedInteger for BACnetAccessThreatLevel must not be nil")
	}
	return &_BACnetAccessThreatLevel{ThreatLevel: threatLevel}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAccessThreatLevelBuilder is a builder for BACnetAccessThreatLevel
type BACnetAccessThreatLevelBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(threatLevel BACnetApplicationTagUnsignedInteger) BACnetAccessThreatLevelBuilder
	// WithThreatLevel adds ThreatLevel (property field)
	WithThreatLevel(BACnetApplicationTagUnsignedInteger) BACnetAccessThreatLevelBuilder
	// WithThreatLevelBuilder adds ThreatLevel (property field) which is build by the builder
	WithThreatLevelBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetAccessThreatLevelBuilder
	// Build builds the BACnetAccessThreatLevel or returns an error if something is wrong
	Build() (BACnetAccessThreatLevel, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAccessThreatLevel
}

// NewBACnetAccessThreatLevelBuilder() creates a BACnetAccessThreatLevelBuilder
func NewBACnetAccessThreatLevelBuilder() BACnetAccessThreatLevelBuilder {
	return &_BACnetAccessThreatLevelBuilder{_BACnetAccessThreatLevel: new(_BACnetAccessThreatLevel)}
}

type _BACnetAccessThreatLevelBuilder struct {
	*_BACnetAccessThreatLevel

	err *utils.MultiError
}

var _ (BACnetAccessThreatLevelBuilder) = (*_BACnetAccessThreatLevelBuilder)(nil)

func (b *_BACnetAccessThreatLevelBuilder) WithMandatoryFields(threatLevel BACnetApplicationTagUnsignedInteger) BACnetAccessThreatLevelBuilder {
	return b.WithThreatLevel(threatLevel)
}

func (b *_BACnetAccessThreatLevelBuilder) WithThreatLevel(threatLevel BACnetApplicationTagUnsignedInteger) BACnetAccessThreatLevelBuilder {
	b.ThreatLevel = threatLevel
	return b
}

func (b *_BACnetAccessThreatLevelBuilder) WithThreatLevelBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetAccessThreatLevelBuilder {
	builder := builderSupplier(b.ThreatLevel.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.ThreatLevel, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetAccessThreatLevelBuilder) Build() (BACnetAccessThreatLevel, error) {
	if b.ThreatLevel == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'threatLevel' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAccessThreatLevel.deepCopy(), nil
}

func (b *_BACnetAccessThreatLevelBuilder) MustBuild() BACnetAccessThreatLevel {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAccessThreatLevelBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAccessThreatLevelBuilder().(*_BACnetAccessThreatLevelBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAccessThreatLevelBuilder creates a BACnetAccessThreatLevelBuilder
func (b *_BACnetAccessThreatLevel) CreateBACnetAccessThreatLevelBuilder() BACnetAccessThreatLevelBuilder {
	if b == nil {
		return NewBACnetAccessThreatLevelBuilder()
	}
	return &_BACnetAccessThreatLevelBuilder{_BACnetAccessThreatLevel: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessThreatLevel) GetThreatLevel() BACnetApplicationTagUnsignedInteger {
	return m.ThreatLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAccessThreatLevel(structType any) BACnetAccessThreatLevel {
	if casted, ok := structType.(BACnetAccessThreatLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessThreatLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessThreatLevel) GetTypeName() string {
	return "BACnetAccessThreatLevel"
}

func (m *_BACnetAccessThreatLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (threatLevel)
	lengthInBits += m.ThreatLevel.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAccessThreatLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessThreatLevelParse(ctx context.Context, theBytes []byte) (BACnetAccessThreatLevel, error) {
	return BACnetAccessThreatLevelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccessThreatLevelParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessThreatLevel, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessThreatLevel, error) {
		return BACnetAccessThreatLevelParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAccessThreatLevelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessThreatLevel, error) {
	v, err := (&_BACnetAccessThreatLevel{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccessThreatLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAccessThreatLevel BACnetAccessThreatLevel, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessThreatLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessThreatLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	threatLevel, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "threatLevel", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'threatLevel' field"))
	}
	m.ThreatLevel = threatLevel

	if closeErr := readBuffer.CloseContext("BACnetAccessThreatLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessThreatLevel")
	}

	return m, nil
}

func (m *_BACnetAccessThreatLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessThreatLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessThreatLevel"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessThreatLevel")
	}

	if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "threatLevel", m.GetThreatLevel(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'threatLevel' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessThreatLevel"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessThreatLevel")
	}
	return nil
}

func (m *_BACnetAccessThreatLevel) IsBACnetAccessThreatLevel() {}

func (m *_BACnetAccessThreatLevel) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAccessThreatLevel) deepCopy() *_BACnetAccessThreatLevel {
	if m == nil {
		return nil
	}
	_BACnetAccessThreatLevelCopy := &_BACnetAccessThreatLevel{
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.ThreatLevel),
	}
	return _BACnetAccessThreatLevelCopy
}

func (m *_BACnetAccessThreatLevel) String() string {
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

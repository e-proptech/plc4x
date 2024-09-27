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

// BACnetKeyIdentifier is the corresponding interface of BACnetKeyIdentifier
type BACnetKeyIdentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetAlgorithm returns Algorithm (property field)
	GetAlgorithm() BACnetContextTagUnsignedInteger
	// GetKeyId returns KeyId (property field)
	GetKeyId() BACnetContextTagUnsignedInteger
	// IsBACnetKeyIdentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetKeyIdentifier()
	// CreateBuilder creates a BACnetKeyIdentifierBuilder
	CreateBACnetKeyIdentifierBuilder() BACnetKeyIdentifierBuilder
}

// _BACnetKeyIdentifier is the data-structure of this message
type _BACnetKeyIdentifier struct {
	Algorithm BACnetContextTagUnsignedInteger
	KeyId     BACnetContextTagUnsignedInteger
}

var _ BACnetKeyIdentifier = (*_BACnetKeyIdentifier)(nil)

// NewBACnetKeyIdentifier factory function for _BACnetKeyIdentifier
func NewBACnetKeyIdentifier(algorithm BACnetContextTagUnsignedInteger, keyId BACnetContextTagUnsignedInteger) *_BACnetKeyIdentifier {
	if algorithm == nil {
		panic("algorithm of type BACnetContextTagUnsignedInteger for BACnetKeyIdentifier must not be nil")
	}
	if keyId == nil {
		panic("keyId of type BACnetContextTagUnsignedInteger for BACnetKeyIdentifier must not be nil")
	}
	return &_BACnetKeyIdentifier{Algorithm: algorithm, KeyId: keyId}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetKeyIdentifierBuilder is a builder for BACnetKeyIdentifier
type BACnetKeyIdentifierBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(algorithm BACnetContextTagUnsignedInteger, keyId BACnetContextTagUnsignedInteger) BACnetKeyIdentifierBuilder
	// WithAlgorithm adds Algorithm (property field)
	WithAlgorithm(BACnetContextTagUnsignedInteger) BACnetKeyIdentifierBuilder
	// WithAlgorithmBuilder adds Algorithm (property field) which is build by the builder
	WithAlgorithmBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetKeyIdentifierBuilder
	// WithKeyId adds KeyId (property field)
	WithKeyId(BACnetContextTagUnsignedInteger) BACnetKeyIdentifierBuilder
	// WithKeyIdBuilder adds KeyId (property field) which is build by the builder
	WithKeyIdBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetKeyIdentifierBuilder
	// Build builds the BACnetKeyIdentifier or returns an error if something is wrong
	Build() (BACnetKeyIdentifier, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetKeyIdentifier
}

// NewBACnetKeyIdentifierBuilder() creates a BACnetKeyIdentifierBuilder
func NewBACnetKeyIdentifierBuilder() BACnetKeyIdentifierBuilder {
	return &_BACnetKeyIdentifierBuilder{_BACnetKeyIdentifier: new(_BACnetKeyIdentifier)}
}

type _BACnetKeyIdentifierBuilder struct {
	*_BACnetKeyIdentifier

	err *utils.MultiError
}

var _ (BACnetKeyIdentifierBuilder) = (*_BACnetKeyIdentifierBuilder)(nil)

func (b *_BACnetKeyIdentifierBuilder) WithMandatoryFields(algorithm BACnetContextTagUnsignedInteger, keyId BACnetContextTagUnsignedInteger) BACnetKeyIdentifierBuilder {
	return b.WithAlgorithm(algorithm).WithKeyId(keyId)
}

func (b *_BACnetKeyIdentifierBuilder) WithAlgorithm(algorithm BACnetContextTagUnsignedInteger) BACnetKeyIdentifierBuilder {
	b.Algorithm = algorithm
	return b
}

func (b *_BACnetKeyIdentifierBuilder) WithAlgorithmBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetKeyIdentifierBuilder {
	builder := builderSupplier(b.Algorithm.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.Algorithm, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetKeyIdentifierBuilder) WithKeyId(keyId BACnetContextTagUnsignedInteger) BACnetKeyIdentifierBuilder {
	b.KeyId = keyId
	return b
}

func (b *_BACnetKeyIdentifierBuilder) WithKeyIdBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetKeyIdentifierBuilder {
	builder := builderSupplier(b.KeyId.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.KeyId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetKeyIdentifierBuilder) Build() (BACnetKeyIdentifier, error) {
	if b.Algorithm == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'algorithm' not set"))
	}
	if b.KeyId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'keyId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetKeyIdentifier.deepCopy(), nil
}

func (b *_BACnetKeyIdentifierBuilder) MustBuild() BACnetKeyIdentifier {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetKeyIdentifierBuilder) DeepCopy() any {
	_copy := b.CreateBACnetKeyIdentifierBuilder().(*_BACnetKeyIdentifierBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetKeyIdentifierBuilder creates a BACnetKeyIdentifierBuilder
func (b *_BACnetKeyIdentifier) CreateBACnetKeyIdentifierBuilder() BACnetKeyIdentifierBuilder {
	if b == nil {
		return NewBACnetKeyIdentifierBuilder()
	}
	return &_BACnetKeyIdentifierBuilder{_BACnetKeyIdentifier: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetKeyIdentifier) GetAlgorithm() BACnetContextTagUnsignedInteger {
	return m.Algorithm
}

func (m *_BACnetKeyIdentifier) GetKeyId() BACnetContextTagUnsignedInteger {
	return m.KeyId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetKeyIdentifier(structType any) BACnetKeyIdentifier {
	if casted, ok := structType.(BACnetKeyIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetKeyIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetKeyIdentifier) GetTypeName() string {
	return "BACnetKeyIdentifier"
}

func (m *_BACnetKeyIdentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (algorithm)
	lengthInBits += m.Algorithm.GetLengthInBits(ctx)

	// Simple field (keyId)
	lengthInBits += m.KeyId.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetKeyIdentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetKeyIdentifierParse(ctx context.Context, theBytes []byte) (BACnetKeyIdentifier, error) {
	return BACnetKeyIdentifierParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetKeyIdentifierParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetKeyIdentifier, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetKeyIdentifier, error) {
		return BACnetKeyIdentifierParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetKeyIdentifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetKeyIdentifier, error) {
	v, err := (&_BACnetKeyIdentifier{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetKeyIdentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetKeyIdentifier BACnetKeyIdentifier, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetKeyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetKeyIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	algorithm, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "algorithm", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'algorithm' field"))
	}
	m.Algorithm = algorithm

	keyId, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "keyId", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keyId' field"))
	}
	m.KeyId = keyId

	if closeErr := readBuffer.CloseContext("BACnetKeyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetKeyIdentifier")
	}

	return m, nil
}

func (m *_BACnetKeyIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetKeyIdentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetKeyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetKeyIdentifier")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "algorithm", m.GetAlgorithm(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'algorithm' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "keyId", m.GetKeyId(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'keyId' field")
	}

	if popErr := writeBuffer.PopContext("BACnetKeyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetKeyIdentifier")
	}
	return nil
}

func (m *_BACnetKeyIdentifier) IsBACnetKeyIdentifier() {}

func (m *_BACnetKeyIdentifier) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetKeyIdentifier) deepCopy() *_BACnetKeyIdentifier {
	if m == nil {
		return nil
	}
	_BACnetKeyIdentifierCopy := &_BACnetKeyIdentifier{
		m.Algorithm.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.KeyId.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	return _BACnetKeyIdentifierCopy
}

func (m *_BACnetKeyIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

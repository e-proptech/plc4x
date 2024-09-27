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

// AdsDiscoveryBlockUserName is the corresponding interface of AdsDiscoveryBlockUserName
type AdsDiscoveryBlockUserName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AdsDiscoveryBlock
	// GetUserName returns UserName (property field)
	GetUserName() AmsString
	// IsAdsDiscoveryBlockUserName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsDiscoveryBlockUserName()
	// CreateBuilder creates a AdsDiscoveryBlockUserNameBuilder
	CreateAdsDiscoveryBlockUserNameBuilder() AdsDiscoveryBlockUserNameBuilder
}

// _AdsDiscoveryBlockUserName is the data-structure of this message
type _AdsDiscoveryBlockUserName struct {
	AdsDiscoveryBlockContract
	UserName AmsString
}

var _ AdsDiscoveryBlockUserName = (*_AdsDiscoveryBlockUserName)(nil)
var _ AdsDiscoveryBlockRequirements = (*_AdsDiscoveryBlockUserName)(nil)

// NewAdsDiscoveryBlockUserName factory function for _AdsDiscoveryBlockUserName
func NewAdsDiscoveryBlockUserName(userName AmsString) *_AdsDiscoveryBlockUserName {
	if userName == nil {
		panic("userName of type AmsString for AdsDiscoveryBlockUserName must not be nil")
	}
	_result := &_AdsDiscoveryBlockUserName{
		AdsDiscoveryBlockContract: NewAdsDiscoveryBlock(),
		UserName:                  userName,
	}
	_result.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsDiscoveryBlockUserNameBuilder is a builder for AdsDiscoveryBlockUserName
type AdsDiscoveryBlockUserNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(userName AmsString) AdsDiscoveryBlockUserNameBuilder
	// WithUserName adds UserName (property field)
	WithUserName(AmsString) AdsDiscoveryBlockUserNameBuilder
	// WithUserNameBuilder adds UserName (property field) which is build by the builder
	WithUserNameBuilder(func(AmsStringBuilder) AmsStringBuilder) AdsDiscoveryBlockUserNameBuilder
	// Build builds the AdsDiscoveryBlockUserName or returns an error if something is wrong
	Build() (AdsDiscoveryBlockUserName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsDiscoveryBlockUserName
}

// NewAdsDiscoveryBlockUserNameBuilder() creates a AdsDiscoveryBlockUserNameBuilder
func NewAdsDiscoveryBlockUserNameBuilder() AdsDiscoveryBlockUserNameBuilder {
	return &_AdsDiscoveryBlockUserNameBuilder{_AdsDiscoveryBlockUserName: new(_AdsDiscoveryBlockUserName)}
}

type _AdsDiscoveryBlockUserNameBuilder struct {
	*_AdsDiscoveryBlockUserName

	parentBuilder *_AdsDiscoveryBlockBuilder

	err *utils.MultiError
}

var _ (AdsDiscoveryBlockUserNameBuilder) = (*_AdsDiscoveryBlockUserNameBuilder)(nil)

func (b *_AdsDiscoveryBlockUserNameBuilder) setParent(contract AdsDiscoveryBlockContract) {
	b.AdsDiscoveryBlockContract = contract
}

func (b *_AdsDiscoveryBlockUserNameBuilder) WithMandatoryFields(userName AmsString) AdsDiscoveryBlockUserNameBuilder {
	return b.WithUserName(userName)
}

func (b *_AdsDiscoveryBlockUserNameBuilder) WithUserName(userName AmsString) AdsDiscoveryBlockUserNameBuilder {
	b.UserName = userName
	return b
}

func (b *_AdsDiscoveryBlockUserNameBuilder) WithUserNameBuilder(builderSupplier func(AmsStringBuilder) AmsStringBuilder) AdsDiscoveryBlockUserNameBuilder {
	builder := builderSupplier(b.UserName.CreateAmsStringBuilder())
	var err error
	b.UserName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "AmsStringBuilder failed"))
	}
	return b
}

func (b *_AdsDiscoveryBlockUserNameBuilder) Build() (AdsDiscoveryBlockUserName, error) {
	if b.UserName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'userName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsDiscoveryBlockUserName.deepCopy(), nil
}

func (b *_AdsDiscoveryBlockUserNameBuilder) MustBuild() AdsDiscoveryBlockUserName {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AdsDiscoveryBlockUserNameBuilder) Done() AdsDiscoveryBlockBuilder {
	return b.parentBuilder
}

func (b *_AdsDiscoveryBlockUserNameBuilder) buildForAdsDiscoveryBlock() (AdsDiscoveryBlock, error) {
	return b.Build()
}

func (b *_AdsDiscoveryBlockUserNameBuilder) DeepCopy() any {
	_copy := b.CreateAdsDiscoveryBlockUserNameBuilder().(*_AdsDiscoveryBlockUserNameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsDiscoveryBlockUserNameBuilder creates a AdsDiscoveryBlockUserNameBuilder
func (b *_AdsDiscoveryBlockUserName) CreateAdsDiscoveryBlockUserNameBuilder() AdsDiscoveryBlockUserNameBuilder {
	if b == nil {
		return NewAdsDiscoveryBlockUserNameBuilder()
	}
	return &_AdsDiscoveryBlockUserNameBuilder{_AdsDiscoveryBlockUserName: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockUserName) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_USER_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockUserName) GetParent() AdsDiscoveryBlockContract {
	return m.AdsDiscoveryBlockContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockUserName) GetUserName() AmsString {
	return m.UserName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockUserName(structType any) AdsDiscoveryBlockUserName {
	if casted, ok := structType.(AdsDiscoveryBlockUserName); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockUserName); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockUserName) GetTypeName() string {
	return "AdsDiscoveryBlockUserName"
}

func (m *_AdsDiscoveryBlockUserName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).getLengthInBits(ctx))

	// Simple field (userName)
	lengthInBits += m.UserName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AdsDiscoveryBlockUserName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsDiscoveryBlockUserName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AdsDiscoveryBlock) (__adsDiscoveryBlockUserName AdsDiscoveryBlockUserName, err error) {
	m.AdsDiscoveryBlockContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockUserName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockUserName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	userName, err := ReadSimpleField[AmsString](ctx, "userName", ReadComplex[AmsString](AmsStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userName' field"))
	}
	m.UserName = userName

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockUserName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockUserName")
	}

	return m, nil
}

func (m *_AdsDiscoveryBlockUserName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockUserName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockUserName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockUserName")
		}

		if err := WriteSimpleField[AmsString](ctx, "userName", m.GetUserName(), WriteComplex[AmsString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userName' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockUserName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockUserName")
		}
		return nil
	}
	return m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockUserName) IsAdsDiscoveryBlockUserName() {}

func (m *_AdsDiscoveryBlockUserName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsDiscoveryBlockUserName) deepCopy() *_AdsDiscoveryBlockUserName {
	if m == nil {
		return nil
	}
	_AdsDiscoveryBlockUserNameCopy := &_AdsDiscoveryBlockUserName{
		m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock).deepCopy(),
		m.UserName.DeepCopy().(AmsString),
	}
	m.AdsDiscoveryBlockContract.(*_AdsDiscoveryBlock)._SubType = m
	return _AdsDiscoveryBlockUserNameCopy
}

func (m *_AdsDiscoveryBlockUserName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

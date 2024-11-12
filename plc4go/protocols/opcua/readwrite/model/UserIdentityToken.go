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

// UserIdentityToken is the corresponding interface of UserIdentityToken
type UserIdentityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetPolicyId returns PolicyId (property field)
	GetPolicyId() PascalString
	// IsUserIdentityToken is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUserIdentityToken()
	// CreateBuilder creates a UserIdentityTokenBuilder
	CreateUserIdentityTokenBuilder() UserIdentityTokenBuilder
}

// _UserIdentityToken is the data-structure of this message
type _UserIdentityToken struct {
	ExtensionObjectDefinitionContract
	PolicyId PascalString
}

var _ UserIdentityToken = (*_UserIdentityToken)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_UserIdentityToken)(nil)

// NewUserIdentityToken factory function for _UserIdentityToken
func NewUserIdentityToken(policyId PascalString) *_UserIdentityToken {
	if policyId == nil {
		panic("policyId of type PascalString for UserIdentityToken must not be nil")
	}
	_result := &_UserIdentityToken{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		PolicyId:                          policyId,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// UserIdentityTokenBuilder is a builder for UserIdentityToken
type UserIdentityTokenBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(policyId PascalString) UserIdentityTokenBuilder
	// WithPolicyId adds PolicyId (property field)
	WithPolicyId(PascalString) UserIdentityTokenBuilder
	// WithPolicyIdBuilder adds PolicyId (property field) which is build by the builder
	WithPolicyIdBuilder(func(PascalStringBuilder) PascalStringBuilder) UserIdentityTokenBuilder
	// Build builds the UserIdentityToken or returns an error if something is wrong
	Build() (UserIdentityToken, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() UserIdentityToken
}

// NewUserIdentityTokenBuilder() creates a UserIdentityTokenBuilder
func NewUserIdentityTokenBuilder() UserIdentityTokenBuilder {
	return &_UserIdentityTokenBuilder{_UserIdentityToken: new(_UserIdentityToken)}
}

type _UserIdentityTokenBuilder struct {
	*_UserIdentityToken

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (UserIdentityTokenBuilder) = (*_UserIdentityTokenBuilder)(nil)

func (b *_UserIdentityTokenBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_UserIdentityTokenBuilder) WithMandatoryFields(policyId PascalString) UserIdentityTokenBuilder {
	return b.WithPolicyId(policyId)
}

func (b *_UserIdentityTokenBuilder) WithPolicyId(policyId PascalString) UserIdentityTokenBuilder {
	b.PolicyId = policyId
	return b
}

func (b *_UserIdentityTokenBuilder) WithPolicyIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) UserIdentityTokenBuilder {
	builder := builderSupplier(b.PolicyId.CreatePascalStringBuilder())
	var err error
	b.PolicyId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_UserIdentityTokenBuilder) Build() (UserIdentityToken, error) {
	if b.PolicyId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'policyId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._UserIdentityToken.deepCopy(), nil
}

func (b *_UserIdentityTokenBuilder) MustBuild() UserIdentityToken {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_UserIdentityTokenBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_UserIdentityTokenBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_UserIdentityTokenBuilder) DeepCopy() any {
	_copy := b.CreateUserIdentityTokenBuilder().(*_UserIdentityTokenBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateUserIdentityTokenBuilder creates a UserIdentityTokenBuilder
func (b *_UserIdentityToken) CreateUserIdentityTokenBuilder() UserIdentityTokenBuilder {
	if b == nil {
		return NewUserIdentityTokenBuilder()
	}
	return &_UserIdentityTokenBuilder{_UserIdentityToken: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UserIdentityToken) GetExtensionId() int32 {
	return int32(318)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UserIdentityToken) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UserIdentityToken) GetPolicyId() PascalString {
	return m.PolicyId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastUserIdentityToken(structType any) UserIdentityToken {
	if casted, ok := structType.(UserIdentityToken); ok {
		return casted
	}
	if casted, ok := structType.(*UserIdentityToken); ok {
		return *casted
	}
	return nil
}

func (m *_UserIdentityToken) GetTypeName() string {
	return "UserIdentityToken"
}

func (m *_UserIdentityToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (policyId)
	lengthInBits += m.PolicyId.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_UserIdentityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_UserIdentityToken) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__userIdentityToken UserIdentityToken, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UserIdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UserIdentityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	policyId, err := ReadSimpleField[PascalString](ctx, "policyId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'policyId' field"))
	}
	m.PolicyId = policyId

	if closeErr := readBuffer.CloseContext("UserIdentityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UserIdentityToken")
	}

	return m, nil
}

func (m *_UserIdentityToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UserIdentityToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UserIdentityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UserIdentityToken")
		}

		if err := WriteSimpleField[PascalString](ctx, "policyId", m.GetPolicyId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'policyId' field")
		}

		if popErr := writeBuffer.PopContext("UserIdentityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UserIdentityToken")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UserIdentityToken) IsUserIdentityToken() {}

func (m *_UserIdentityToken) DeepCopy() any {
	return m.deepCopy()
}

func (m *_UserIdentityToken) deepCopy() *_UserIdentityToken {
	if m == nil {
		return nil
	}
	_UserIdentityTokenCopy := &_UserIdentityToken{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.PolicyId),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _UserIdentityTokenCopy
}

func (m *_UserIdentityToken) String() string {
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

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

// X509IdentityToken is the corresponding interface of X509IdentityToken
type X509IdentityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetPolicyId returns PolicyId (property field)
	GetPolicyId() PascalString
	// GetCertificateData returns CertificateData (property field)
	GetCertificateData() PascalByteString
	// IsX509IdentityToken is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsX509IdentityToken()
	// CreateBuilder creates a X509IdentityTokenBuilder
	CreateX509IdentityTokenBuilder() X509IdentityTokenBuilder
}

// _X509IdentityToken is the data-structure of this message
type _X509IdentityToken struct {
	ExtensionObjectDefinitionContract
	PolicyId        PascalString
	CertificateData PascalByteString
}

var _ X509IdentityToken = (*_X509IdentityToken)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_X509IdentityToken)(nil)

// NewX509IdentityToken factory function for _X509IdentityToken
func NewX509IdentityToken(policyId PascalString, certificateData PascalByteString) *_X509IdentityToken {
	if policyId == nil {
		panic("policyId of type PascalString for X509IdentityToken must not be nil")
	}
	if certificateData == nil {
		panic("certificateData of type PascalByteString for X509IdentityToken must not be nil")
	}
	_result := &_X509IdentityToken{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		PolicyId:                          policyId,
		CertificateData:                   certificateData,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// X509IdentityTokenBuilder is a builder for X509IdentityToken
type X509IdentityTokenBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(policyId PascalString, certificateData PascalByteString) X509IdentityTokenBuilder
	// WithPolicyId adds PolicyId (property field)
	WithPolicyId(PascalString) X509IdentityTokenBuilder
	// WithPolicyIdBuilder adds PolicyId (property field) which is build by the builder
	WithPolicyIdBuilder(func(PascalStringBuilder) PascalStringBuilder) X509IdentityTokenBuilder
	// WithCertificateData adds CertificateData (property field)
	WithCertificateData(PascalByteString) X509IdentityTokenBuilder
	// WithCertificateDataBuilder adds CertificateData (property field) which is build by the builder
	WithCertificateDataBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) X509IdentityTokenBuilder
	// Build builds the X509IdentityToken or returns an error if something is wrong
	Build() (X509IdentityToken, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() X509IdentityToken
}

// NewX509IdentityTokenBuilder() creates a X509IdentityTokenBuilder
func NewX509IdentityTokenBuilder() X509IdentityTokenBuilder {
	return &_X509IdentityTokenBuilder{_X509IdentityToken: new(_X509IdentityToken)}
}

type _X509IdentityTokenBuilder struct {
	*_X509IdentityToken

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (X509IdentityTokenBuilder) = (*_X509IdentityTokenBuilder)(nil)

func (b *_X509IdentityTokenBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_X509IdentityTokenBuilder) WithMandatoryFields(policyId PascalString, certificateData PascalByteString) X509IdentityTokenBuilder {
	return b.WithPolicyId(policyId).WithCertificateData(certificateData)
}

func (b *_X509IdentityTokenBuilder) WithPolicyId(policyId PascalString) X509IdentityTokenBuilder {
	b.PolicyId = policyId
	return b
}

func (b *_X509IdentityTokenBuilder) WithPolicyIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) X509IdentityTokenBuilder {
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

func (b *_X509IdentityTokenBuilder) WithCertificateData(certificateData PascalByteString) X509IdentityTokenBuilder {
	b.CertificateData = certificateData
	return b
}

func (b *_X509IdentityTokenBuilder) WithCertificateDataBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) X509IdentityTokenBuilder {
	builder := builderSupplier(b.CertificateData.CreatePascalByteStringBuilder())
	var err error
	b.CertificateData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return b
}

func (b *_X509IdentityTokenBuilder) Build() (X509IdentityToken, error) {
	if b.PolicyId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'policyId' not set"))
	}
	if b.CertificateData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'certificateData' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._X509IdentityToken.deepCopy(), nil
}

func (b *_X509IdentityTokenBuilder) MustBuild() X509IdentityToken {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_X509IdentityTokenBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_X509IdentityTokenBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_X509IdentityTokenBuilder) DeepCopy() any {
	_copy := b.CreateX509IdentityTokenBuilder().(*_X509IdentityTokenBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateX509IdentityTokenBuilder creates a X509IdentityTokenBuilder
func (b *_X509IdentityToken) CreateX509IdentityTokenBuilder() X509IdentityTokenBuilder {
	if b == nil {
		return NewX509IdentityTokenBuilder()
	}
	return &_X509IdentityTokenBuilder{_X509IdentityToken: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_X509IdentityToken) GetExtensionId() int32 {
	return int32(327)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_X509IdentityToken) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_X509IdentityToken) GetPolicyId() PascalString {
	return m.PolicyId
}

func (m *_X509IdentityToken) GetCertificateData() PascalByteString {
	return m.CertificateData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastX509IdentityToken(structType any) X509IdentityToken {
	if casted, ok := structType.(X509IdentityToken); ok {
		return casted
	}
	if casted, ok := structType.(*X509IdentityToken); ok {
		return *casted
	}
	return nil
}

func (m *_X509IdentityToken) GetTypeName() string {
	return "X509IdentityToken"
}

func (m *_X509IdentityToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (policyId)
	lengthInBits += m.PolicyId.GetLengthInBits(ctx)

	// Simple field (certificateData)
	lengthInBits += m.CertificateData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_X509IdentityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_X509IdentityToken) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__x509IdentityToken X509IdentityToken, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("X509IdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for X509IdentityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	policyId, err := ReadSimpleField[PascalString](ctx, "policyId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'policyId' field"))
	}
	m.PolicyId = policyId

	certificateData, err := ReadSimpleField[PascalByteString](ctx, "certificateData", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'certificateData' field"))
	}
	m.CertificateData = certificateData

	if closeErr := readBuffer.CloseContext("X509IdentityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for X509IdentityToken")
	}

	return m, nil
}

func (m *_X509IdentityToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_X509IdentityToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("X509IdentityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for X509IdentityToken")
		}

		if err := WriteSimpleField[PascalString](ctx, "policyId", m.GetPolicyId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'policyId' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "certificateData", m.GetCertificateData(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'certificateData' field")
		}

		if popErr := writeBuffer.PopContext("X509IdentityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for X509IdentityToken")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_X509IdentityToken) IsX509IdentityToken() {}

func (m *_X509IdentityToken) DeepCopy() any {
	return m.deepCopy()
}

func (m *_X509IdentityToken) deepCopy() *_X509IdentityToken {
	if m == nil {
		return nil
	}
	_X509IdentityTokenCopy := &_X509IdentityToken{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.PolicyId),
		utils.DeepCopy[PascalByteString](m.CertificateData),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _X509IdentityTokenCopy
}

func (m *_X509IdentityToken) String() string {
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

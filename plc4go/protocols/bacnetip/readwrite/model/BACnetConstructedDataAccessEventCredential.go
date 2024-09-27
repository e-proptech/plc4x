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

// BACnetConstructedDataAccessEventCredential is the corresponding interface of BACnetConstructedDataAccessEventCredential
type BACnetConstructedDataAccessEventCredential interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAccessEventCredential returns AccessEventCredential (property field)
	GetAccessEventCredential() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
	// IsBACnetConstructedDataAccessEventCredential is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessEventCredential()
	// CreateBuilder creates a BACnetConstructedDataAccessEventCredentialBuilder
	CreateBACnetConstructedDataAccessEventCredentialBuilder() BACnetConstructedDataAccessEventCredentialBuilder
}

// _BACnetConstructedDataAccessEventCredential is the data-structure of this message
type _BACnetConstructedDataAccessEventCredential struct {
	BACnetConstructedDataContract
	AccessEventCredential BACnetDeviceObjectReference
}

var _ BACnetConstructedDataAccessEventCredential = (*_BACnetConstructedDataAccessEventCredential)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessEventCredential)(nil)

// NewBACnetConstructedDataAccessEventCredential factory function for _BACnetConstructedDataAccessEventCredential
func NewBACnetConstructedDataAccessEventCredential(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, accessEventCredential BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessEventCredential {
	if accessEventCredential == nil {
		panic("accessEventCredential of type BACnetDeviceObjectReference for BACnetConstructedDataAccessEventCredential must not be nil")
	}
	_result := &_BACnetConstructedDataAccessEventCredential{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AccessEventCredential:         accessEventCredential,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccessEventCredentialBuilder is a builder for BACnetConstructedDataAccessEventCredential
type BACnetConstructedDataAccessEventCredentialBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(accessEventCredential BACnetDeviceObjectReference) BACnetConstructedDataAccessEventCredentialBuilder
	// WithAccessEventCredential adds AccessEventCredential (property field)
	WithAccessEventCredential(BACnetDeviceObjectReference) BACnetConstructedDataAccessEventCredentialBuilder
	// WithAccessEventCredentialBuilder adds AccessEventCredential (property field) which is build by the builder
	WithAccessEventCredentialBuilder(func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataAccessEventCredentialBuilder
	// Build builds the BACnetConstructedDataAccessEventCredential or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessEventCredential, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessEventCredential
}

// NewBACnetConstructedDataAccessEventCredentialBuilder() creates a BACnetConstructedDataAccessEventCredentialBuilder
func NewBACnetConstructedDataAccessEventCredentialBuilder() BACnetConstructedDataAccessEventCredentialBuilder {
	return &_BACnetConstructedDataAccessEventCredentialBuilder{_BACnetConstructedDataAccessEventCredential: new(_BACnetConstructedDataAccessEventCredential)}
}

type _BACnetConstructedDataAccessEventCredentialBuilder struct {
	*_BACnetConstructedDataAccessEventCredential

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessEventCredentialBuilder) = (*_BACnetConstructedDataAccessEventCredentialBuilder)(nil)

func (b *_BACnetConstructedDataAccessEventCredentialBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAccessEventCredentialBuilder) WithMandatoryFields(accessEventCredential BACnetDeviceObjectReference) BACnetConstructedDataAccessEventCredentialBuilder {
	return b.WithAccessEventCredential(accessEventCredential)
}

func (b *_BACnetConstructedDataAccessEventCredentialBuilder) WithAccessEventCredential(accessEventCredential BACnetDeviceObjectReference) BACnetConstructedDataAccessEventCredentialBuilder {
	b.AccessEventCredential = accessEventCredential
	return b
}

func (b *_BACnetConstructedDataAccessEventCredentialBuilder) WithAccessEventCredentialBuilder(builderSupplier func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataAccessEventCredentialBuilder {
	builder := builderSupplier(b.AccessEventCredential.CreateBACnetDeviceObjectReferenceBuilder())
	var err error
	b.AccessEventCredential, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAccessEventCredentialBuilder) Build() (BACnetConstructedDataAccessEventCredential, error) {
	if b.AccessEventCredential == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'accessEventCredential' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccessEventCredential.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccessEventCredentialBuilder) MustBuild() BACnetConstructedDataAccessEventCredential {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAccessEventCredentialBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccessEventCredentialBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccessEventCredentialBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccessEventCredentialBuilder().(*_BACnetConstructedDataAccessEventCredentialBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccessEventCredentialBuilder creates a BACnetConstructedDataAccessEventCredentialBuilder
func (b *_BACnetConstructedDataAccessEventCredential) CreateBACnetConstructedDataAccessEventCredentialBuilder() BACnetConstructedDataAccessEventCredentialBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccessEventCredentialBuilder()
	}
	return &_BACnetConstructedDataAccessEventCredentialBuilder{_BACnetConstructedDataAccessEventCredential: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessEventCredential) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessEventCredential) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT_CREDENTIAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessEventCredential) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventCredential) GetAccessEventCredential() BACnetDeviceObjectReference {
	return m.AccessEventCredential
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventCredential) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetAccessEventCredential())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessEventCredential(structType any) BACnetConstructedDataAccessEventCredential {
	if casted, ok := structType.(BACnetConstructedDataAccessEventCredential); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEventCredential); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessEventCredential) GetTypeName() string {
	return "BACnetConstructedDataAccessEventCredential"
}

func (m *_BACnetConstructedDataAccessEventCredential) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (accessEventCredential)
	lengthInBits += m.AccessEventCredential.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessEventCredential) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessEventCredential) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessEventCredential BACnetConstructedDataAccessEventCredential, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEventCredential"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessEventCredential")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessEventCredential, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "accessEventCredential", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessEventCredential' field"))
	}
	m.AccessEventCredential = accessEventCredential

	actualValue, err := ReadVirtualField[BACnetDeviceObjectReference](ctx, "actualValue", (*BACnetDeviceObjectReference)(nil), accessEventCredential)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEventCredential"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessEventCredential")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessEventCredential) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessEventCredential) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEventCredential"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessEventCredential")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReference](ctx, "accessEventCredential", m.GetAccessEventCredential(), WriteComplex[BACnetDeviceObjectReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessEventCredential' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEventCredential"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessEventCredential")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessEventCredential) IsBACnetConstructedDataAccessEventCredential() {
}

func (m *_BACnetConstructedDataAccessEventCredential) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessEventCredential) deepCopy() *_BACnetConstructedDataAccessEventCredential {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessEventCredentialCopy := &_BACnetConstructedDataAccessEventCredential{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.AccessEventCredential.DeepCopy().(BACnetDeviceObjectReference),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessEventCredentialCopy
}

func (m *_BACnetConstructedDataAccessEventCredential) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

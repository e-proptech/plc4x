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

// BACnetConstructedDataEventAlgorithmInhibitRef is the corresponding interface of BACnetConstructedDataEventAlgorithmInhibitRef
type BACnetConstructedDataEventAlgorithmInhibitRef interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEventAlgorithmInhibitRef returns EventAlgorithmInhibitRef (property field)
	GetEventAlgorithmInhibitRef() BACnetObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectPropertyReference
	// IsBACnetConstructedDataEventAlgorithmInhibitRef is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEventAlgorithmInhibitRef()
	// CreateBuilder creates a BACnetConstructedDataEventAlgorithmInhibitRefBuilder
	CreateBACnetConstructedDataEventAlgorithmInhibitRefBuilder() BACnetConstructedDataEventAlgorithmInhibitRefBuilder
}

// _BACnetConstructedDataEventAlgorithmInhibitRef is the data-structure of this message
type _BACnetConstructedDataEventAlgorithmInhibitRef struct {
	BACnetConstructedDataContract
	EventAlgorithmInhibitRef BACnetObjectPropertyReference
}

var _ BACnetConstructedDataEventAlgorithmInhibitRef = (*_BACnetConstructedDataEventAlgorithmInhibitRef)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEventAlgorithmInhibitRef)(nil)

// NewBACnetConstructedDataEventAlgorithmInhibitRef factory function for _BACnetConstructedDataEventAlgorithmInhibitRef
func NewBACnetConstructedDataEventAlgorithmInhibitRef(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, eventAlgorithmInhibitRef BACnetObjectPropertyReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventAlgorithmInhibitRef {
	if eventAlgorithmInhibitRef == nil {
		panic("eventAlgorithmInhibitRef of type BACnetObjectPropertyReference for BACnetConstructedDataEventAlgorithmInhibitRef must not be nil")
	}
	_result := &_BACnetConstructedDataEventAlgorithmInhibitRef{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EventAlgorithmInhibitRef:      eventAlgorithmInhibitRef,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEventAlgorithmInhibitRefBuilder is a builder for BACnetConstructedDataEventAlgorithmInhibitRef
type BACnetConstructedDataEventAlgorithmInhibitRefBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(eventAlgorithmInhibitRef BACnetObjectPropertyReference) BACnetConstructedDataEventAlgorithmInhibitRefBuilder
	// WithEventAlgorithmInhibitRef adds EventAlgorithmInhibitRef (property field)
	WithEventAlgorithmInhibitRef(BACnetObjectPropertyReference) BACnetConstructedDataEventAlgorithmInhibitRefBuilder
	// WithEventAlgorithmInhibitRefBuilder adds EventAlgorithmInhibitRef (property field) which is build by the builder
	WithEventAlgorithmInhibitRefBuilder(func(BACnetObjectPropertyReferenceBuilder) BACnetObjectPropertyReferenceBuilder) BACnetConstructedDataEventAlgorithmInhibitRefBuilder
	// Build builds the BACnetConstructedDataEventAlgorithmInhibitRef or returns an error if something is wrong
	Build() (BACnetConstructedDataEventAlgorithmInhibitRef, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEventAlgorithmInhibitRef
}

// NewBACnetConstructedDataEventAlgorithmInhibitRefBuilder() creates a BACnetConstructedDataEventAlgorithmInhibitRefBuilder
func NewBACnetConstructedDataEventAlgorithmInhibitRefBuilder() BACnetConstructedDataEventAlgorithmInhibitRefBuilder {
	return &_BACnetConstructedDataEventAlgorithmInhibitRefBuilder{_BACnetConstructedDataEventAlgorithmInhibitRef: new(_BACnetConstructedDataEventAlgorithmInhibitRef)}
}

type _BACnetConstructedDataEventAlgorithmInhibitRefBuilder struct {
	*_BACnetConstructedDataEventAlgorithmInhibitRef

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEventAlgorithmInhibitRefBuilder) = (*_BACnetConstructedDataEventAlgorithmInhibitRefBuilder)(nil)

func (b *_BACnetConstructedDataEventAlgorithmInhibitRefBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataEventAlgorithmInhibitRefBuilder) WithMandatoryFields(eventAlgorithmInhibitRef BACnetObjectPropertyReference) BACnetConstructedDataEventAlgorithmInhibitRefBuilder {
	return b.WithEventAlgorithmInhibitRef(eventAlgorithmInhibitRef)
}

func (b *_BACnetConstructedDataEventAlgorithmInhibitRefBuilder) WithEventAlgorithmInhibitRef(eventAlgorithmInhibitRef BACnetObjectPropertyReference) BACnetConstructedDataEventAlgorithmInhibitRefBuilder {
	b.EventAlgorithmInhibitRef = eventAlgorithmInhibitRef
	return b
}

func (b *_BACnetConstructedDataEventAlgorithmInhibitRefBuilder) WithEventAlgorithmInhibitRefBuilder(builderSupplier func(BACnetObjectPropertyReferenceBuilder) BACnetObjectPropertyReferenceBuilder) BACnetConstructedDataEventAlgorithmInhibitRefBuilder {
	builder := builderSupplier(b.EventAlgorithmInhibitRef.CreateBACnetObjectPropertyReferenceBuilder())
	var err error
	b.EventAlgorithmInhibitRef, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetObjectPropertyReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataEventAlgorithmInhibitRefBuilder) Build() (BACnetConstructedDataEventAlgorithmInhibitRef, error) {
	if b.EventAlgorithmInhibitRef == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'eventAlgorithmInhibitRef' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEventAlgorithmInhibitRef.deepCopy(), nil
}

func (b *_BACnetConstructedDataEventAlgorithmInhibitRefBuilder) MustBuild() BACnetConstructedDataEventAlgorithmInhibitRef {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataEventAlgorithmInhibitRefBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEventAlgorithmInhibitRefBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEventAlgorithmInhibitRefBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEventAlgorithmInhibitRefBuilder().(*_BACnetConstructedDataEventAlgorithmInhibitRefBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEventAlgorithmInhibitRefBuilder creates a BACnetConstructedDataEventAlgorithmInhibitRefBuilder
func (b *_BACnetConstructedDataEventAlgorithmInhibitRef) CreateBACnetConstructedDataEventAlgorithmInhibitRefBuilder() BACnetConstructedDataEventAlgorithmInhibitRefBuilder {
	if b == nil {
		return NewBACnetConstructedDataEventAlgorithmInhibitRefBuilder()
	}
	return &_BACnetConstructedDataEventAlgorithmInhibitRefBuilder{_BACnetConstructedDataEventAlgorithmInhibitRef: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_ALGORITHM_INHIBIT_REF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetEventAlgorithmInhibitRef() BACnetObjectPropertyReference {
	return m.EventAlgorithmInhibitRef
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetActualValue() BACnetObjectPropertyReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetObjectPropertyReference(m.GetEventAlgorithmInhibitRef())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventAlgorithmInhibitRef(structType any) BACnetConstructedDataEventAlgorithmInhibitRef {
	if casted, ok := structType.(BACnetConstructedDataEventAlgorithmInhibitRef); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventAlgorithmInhibitRef); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetTypeName() string {
	return "BACnetConstructedDataEventAlgorithmInhibitRef"
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (eventAlgorithmInhibitRef)
	lengthInBits += m.EventAlgorithmInhibitRef.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventAlgorithmInhibitRef BACnetConstructedDataEventAlgorithmInhibitRef, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventAlgorithmInhibitRef"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventAlgorithmInhibitRef")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	eventAlgorithmInhibitRef, err := ReadSimpleField[BACnetObjectPropertyReference](ctx, "eventAlgorithmInhibitRef", ReadComplex[BACnetObjectPropertyReference](BACnetObjectPropertyReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventAlgorithmInhibitRef' field"))
	}
	m.EventAlgorithmInhibitRef = eventAlgorithmInhibitRef

	actualValue, err := ReadVirtualField[BACnetObjectPropertyReference](ctx, "actualValue", (*BACnetObjectPropertyReference)(nil), eventAlgorithmInhibitRef)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventAlgorithmInhibitRef"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventAlgorithmInhibitRef")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventAlgorithmInhibitRef"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventAlgorithmInhibitRef")
		}

		if err := WriteSimpleField[BACnetObjectPropertyReference](ctx, "eventAlgorithmInhibitRef", m.GetEventAlgorithmInhibitRef(), WriteComplex[BACnetObjectPropertyReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventAlgorithmInhibitRef' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventAlgorithmInhibitRef"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventAlgorithmInhibitRef")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) IsBACnetConstructedDataEventAlgorithmInhibitRef() {
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) deepCopy() *_BACnetConstructedDataEventAlgorithmInhibitRef {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEventAlgorithmInhibitRefCopy := &_BACnetConstructedDataEventAlgorithmInhibitRef{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetObjectPropertyReference](m.EventAlgorithmInhibitRef),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEventAlgorithmInhibitRefCopy
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) String() string {
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

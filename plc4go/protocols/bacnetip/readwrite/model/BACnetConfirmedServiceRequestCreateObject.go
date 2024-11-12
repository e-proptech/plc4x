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

// BACnetConfirmedServiceRequestCreateObject is the corresponding interface of BACnetConfirmedServiceRequestCreateObject
type BACnetConfirmedServiceRequestCreateObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetObjectSpecifier returns ObjectSpecifier (property field)
	GetObjectSpecifier() BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() BACnetPropertyValues
	// IsBACnetConfirmedServiceRequestCreateObject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestCreateObject()
	// CreateBuilder creates a BACnetConfirmedServiceRequestCreateObjectBuilder
	CreateBACnetConfirmedServiceRequestCreateObjectBuilder() BACnetConfirmedServiceRequestCreateObjectBuilder
}

// _BACnetConfirmedServiceRequestCreateObject is the data-structure of this message
type _BACnetConfirmedServiceRequestCreateObject struct {
	BACnetConfirmedServiceRequestContract
	ObjectSpecifier BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
	ListOfValues    BACnetPropertyValues
}

var _ BACnetConfirmedServiceRequestCreateObject = (*_BACnetConfirmedServiceRequestCreateObject)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestCreateObject)(nil)

// NewBACnetConfirmedServiceRequestCreateObject factory function for _BACnetConfirmedServiceRequestCreateObject
func NewBACnetConfirmedServiceRequestCreateObject(objectSpecifier BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, listOfValues BACnetPropertyValues, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestCreateObject {
	if objectSpecifier == nil {
		panic("objectSpecifier of type BACnetConfirmedServiceRequestCreateObjectObjectSpecifier for BACnetConfirmedServiceRequestCreateObject must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestCreateObject{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		ObjectSpecifier:                       objectSpecifier,
		ListOfValues:                          listOfValues,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestCreateObjectBuilder is a builder for BACnetConfirmedServiceRequestCreateObject
type BACnetConfirmedServiceRequestCreateObjectBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectSpecifier BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) BACnetConfirmedServiceRequestCreateObjectBuilder
	// WithObjectSpecifier adds ObjectSpecifier (property field)
	WithObjectSpecifier(BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) BACnetConfirmedServiceRequestCreateObjectBuilder
	// WithObjectSpecifierBuilder adds ObjectSpecifier (property field) which is build by the builder
	WithObjectSpecifierBuilder(func(BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) BACnetConfirmedServiceRequestCreateObjectBuilder
	// WithListOfValues adds ListOfValues (property field)
	WithOptionalListOfValues(BACnetPropertyValues) BACnetConfirmedServiceRequestCreateObjectBuilder
	// WithOptionalListOfValuesBuilder adds ListOfValues (property field) which is build by the builder
	WithOptionalListOfValuesBuilder(func(BACnetPropertyValuesBuilder) BACnetPropertyValuesBuilder) BACnetConfirmedServiceRequestCreateObjectBuilder
	// Build builds the BACnetConfirmedServiceRequestCreateObject or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestCreateObject, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestCreateObject
}

// NewBACnetConfirmedServiceRequestCreateObjectBuilder() creates a BACnetConfirmedServiceRequestCreateObjectBuilder
func NewBACnetConfirmedServiceRequestCreateObjectBuilder() BACnetConfirmedServiceRequestCreateObjectBuilder {
	return &_BACnetConfirmedServiceRequestCreateObjectBuilder{_BACnetConfirmedServiceRequestCreateObject: new(_BACnetConfirmedServiceRequestCreateObject)}
}

type _BACnetConfirmedServiceRequestCreateObjectBuilder struct {
	*_BACnetConfirmedServiceRequestCreateObject

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestCreateObjectBuilder) = (*_BACnetConfirmedServiceRequestCreateObjectBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
}

func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) WithMandatoryFields(objectSpecifier BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) BACnetConfirmedServiceRequestCreateObjectBuilder {
	return b.WithObjectSpecifier(objectSpecifier)
}

func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) WithObjectSpecifier(objectSpecifier BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) BACnetConfirmedServiceRequestCreateObjectBuilder {
	b.ObjectSpecifier = objectSpecifier
	return b
}

func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) WithObjectSpecifierBuilder(builderSupplier func(BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder) BACnetConfirmedServiceRequestCreateObjectBuilder {
	builder := builderSupplier(b.ObjectSpecifier.CreateBACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder())
	var err error
	b.ObjectSpecifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConfirmedServiceRequestCreateObjectObjectSpecifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) WithOptionalListOfValues(listOfValues BACnetPropertyValues) BACnetConfirmedServiceRequestCreateObjectBuilder {
	b.ListOfValues = listOfValues
	return b
}

func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) WithOptionalListOfValuesBuilder(builderSupplier func(BACnetPropertyValuesBuilder) BACnetPropertyValuesBuilder) BACnetConfirmedServiceRequestCreateObjectBuilder {
	builder := builderSupplier(b.ListOfValues.CreateBACnetPropertyValuesBuilder())
	var err error
	b.ListOfValues, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPropertyValuesBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) Build() (BACnetConfirmedServiceRequestCreateObject, error) {
	if b.ObjectSpecifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectSpecifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestCreateObject.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) MustBuild() BACnetConfirmedServiceRequestCreateObject {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestCreateObjectBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestCreateObjectBuilder().(*_BACnetConfirmedServiceRequestCreateObjectBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestCreateObjectBuilder creates a BACnetConfirmedServiceRequestCreateObjectBuilder
func (b *_BACnetConfirmedServiceRequestCreateObject) CreateBACnetConfirmedServiceRequestCreateObjectBuilder() BACnetConfirmedServiceRequestCreateObjectBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestCreateObjectBuilder()
	}
	return &_BACnetConfirmedServiceRequestCreateObjectBuilder{_BACnetConfirmedServiceRequestCreateObject: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestCreateObject) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CREATE_OBJECT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestCreateObject) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestCreateObject) GetObjectSpecifier() BACnetConfirmedServiceRequestCreateObjectObjectSpecifier {
	return m.ObjectSpecifier
}

func (m *_BACnetConfirmedServiceRequestCreateObject) GetListOfValues() BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestCreateObject(structType any) BACnetConfirmedServiceRequestCreateObject {
	if casted, ok := structType.(BACnetConfirmedServiceRequestCreateObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestCreateObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestCreateObject) GetTypeName() string {
	return "BACnetConfirmedServiceRequestCreateObject"
}

func (m *_BACnetConfirmedServiceRequestCreateObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (objectSpecifier)
	lengthInBits += m.ObjectSpecifier.GetLengthInBits(ctx)

	// Optional Field (listOfValues)
	if m.ListOfValues != nil {
		lengthInBits += m.ListOfValues.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestCreateObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestCreateObject) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestCreateObject BACnetConfirmedServiceRequestCreateObject, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestCreateObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestCreateObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectSpecifier, err := ReadSimpleField[BACnetConfirmedServiceRequestCreateObjectObjectSpecifier](ctx, "objectSpecifier", ReadComplex[BACnetConfirmedServiceRequestCreateObjectObjectSpecifier](BACnetConfirmedServiceRequestCreateObjectObjectSpecifierParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectSpecifier' field"))
	}
	m.ObjectSpecifier = objectSpecifier

	var listOfValues BACnetPropertyValues
	_listOfValues, err := ReadOptionalField[BACnetPropertyValues](ctx, "listOfValues", ReadComplex[BACnetPropertyValues](BACnetPropertyValuesParseWithBufferProducer((uint8)(uint8(1)), (BACnetObjectType)(CastBACnetObjectType(utils.InlineIf(objectSpecifier.GetIsObjectType(), func() any { return CastBACnetObjectType(objectSpecifier.GetObjectType()) }, func() any { return CastBACnetObjectType(objectSpecifier.GetObjectIdentifier().GetObjectType()) })))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfValues' field"))
	}
	if _listOfValues != nil {
		listOfValues = *_listOfValues
		m.ListOfValues = listOfValues
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestCreateObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestCreateObject")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestCreateObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestCreateObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestCreateObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestCreateObject")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestCreateObjectObjectSpecifier](ctx, "objectSpecifier", m.GetObjectSpecifier(), WriteComplex[BACnetConfirmedServiceRequestCreateObjectObjectSpecifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectSpecifier' field")
		}

		if err := WriteOptionalField[BACnetPropertyValues](ctx, "listOfValues", GetRef(m.GetListOfValues()), WriteComplex[BACnetPropertyValues](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestCreateObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestCreateObject")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestCreateObject) IsBACnetConfirmedServiceRequestCreateObject() {}

func (m *_BACnetConfirmedServiceRequestCreateObject) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestCreateObject) deepCopy() *_BACnetConfirmedServiceRequestCreateObject {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestCreateObjectCopy := &_BACnetConfirmedServiceRequestCreateObject{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetConfirmedServiceRequestCreateObjectObjectSpecifier](m.ObjectSpecifier),
		utils.DeepCopy[BACnetPropertyValues](m.ListOfValues),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestCreateObjectCopy
}

func (m *_BACnetConfirmedServiceRequestCreateObject) String() string {
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

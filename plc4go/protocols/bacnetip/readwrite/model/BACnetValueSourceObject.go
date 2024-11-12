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

// BACnetValueSourceObject is the corresponding interface of BACnetValueSourceObject
type BACnetValueSourceObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetValueSource
	// GetObject returns Object (property field)
	GetObject() BACnetDeviceObjectReferenceEnclosed
	// IsBACnetValueSourceObject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetValueSourceObject()
	// CreateBuilder creates a BACnetValueSourceObjectBuilder
	CreateBACnetValueSourceObjectBuilder() BACnetValueSourceObjectBuilder
}

// _BACnetValueSourceObject is the data-structure of this message
type _BACnetValueSourceObject struct {
	BACnetValueSourceContract
	Object BACnetDeviceObjectReferenceEnclosed
}

var _ BACnetValueSourceObject = (*_BACnetValueSourceObject)(nil)
var _ BACnetValueSourceRequirements = (*_BACnetValueSourceObject)(nil)

// NewBACnetValueSourceObject factory function for _BACnetValueSourceObject
func NewBACnetValueSourceObject(peekedTagHeader BACnetTagHeader, object BACnetDeviceObjectReferenceEnclosed) *_BACnetValueSourceObject {
	if object == nil {
		panic("object of type BACnetDeviceObjectReferenceEnclosed for BACnetValueSourceObject must not be nil")
	}
	_result := &_BACnetValueSourceObject{
		BACnetValueSourceContract: NewBACnetValueSource(peekedTagHeader),
		Object:                    object,
	}
	_result.BACnetValueSourceContract.(*_BACnetValueSource)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetValueSourceObjectBuilder is a builder for BACnetValueSourceObject
type BACnetValueSourceObjectBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(object BACnetDeviceObjectReferenceEnclosed) BACnetValueSourceObjectBuilder
	// WithObject adds Object (property field)
	WithObject(BACnetDeviceObjectReferenceEnclosed) BACnetValueSourceObjectBuilder
	// WithObjectBuilder adds Object (property field) which is build by the builder
	WithObjectBuilder(func(BACnetDeviceObjectReferenceEnclosedBuilder) BACnetDeviceObjectReferenceEnclosedBuilder) BACnetValueSourceObjectBuilder
	// Build builds the BACnetValueSourceObject or returns an error if something is wrong
	Build() (BACnetValueSourceObject, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetValueSourceObject
}

// NewBACnetValueSourceObjectBuilder() creates a BACnetValueSourceObjectBuilder
func NewBACnetValueSourceObjectBuilder() BACnetValueSourceObjectBuilder {
	return &_BACnetValueSourceObjectBuilder{_BACnetValueSourceObject: new(_BACnetValueSourceObject)}
}

type _BACnetValueSourceObjectBuilder struct {
	*_BACnetValueSourceObject

	parentBuilder *_BACnetValueSourceBuilder

	err *utils.MultiError
}

var _ (BACnetValueSourceObjectBuilder) = (*_BACnetValueSourceObjectBuilder)(nil)

func (b *_BACnetValueSourceObjectBuilder) setParent(contract BACnetValueSourceContract) {
	b.BACnetValueSourceContract = contract
}

func (b *_BACnetValueSourceObjectBuilder) WithMandatoryFields(object BACnetDeviceObjectReferenceEnclosed) BACnetValueSourceObjectBuilder {
	return b.WithObject(object)
}

func (b *_BACnetValueSourceObjectBuilder) WithObject(object BACnetDeviceObjectReferenceEnclosed) BACnetValueSourceObjectBuilder {
	b.Object = object
	return b
}

func (b *_BACnetValueSourceObjectBuilder) WithObjectBuilder(builderSupplier func(BACnetDeviceObjectReferenceEnclosedBuilder) BACnetDeviceObjectReferenceEnclosedBuilder) BACnetValueSourceObjectBuilder {
	builder := builderSupplier(b.Object.CreateBACnetDeviceObjectReferenceEnclosedBuilder())
	var err error
	b.Object, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectReferenceEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetValueSourceObjectBuilder) Build() (BACnetValueSourceObject, error) {
	if b.Object == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'object' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetValueSourceObject.deepCopy(), nil
}

func (b *_BACnetValueSourceObjectBuilder) MustBuild() BACnetValueSourceObject {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetValueSourceObjectBuilder) Done() BACnetValueSourceBuilder {
	return b.parentBuilder
}

func (b *_BACnetValueSourceObjectBuilder) buildForBACnetValueSource() (BACnetValueSource, error) {
	return b.Build()
}

func (b *_BACnetValueSourceObjectBuilder) DeepCopy() any {
	_copy := b.CreateBACnetValueSourceObjectBuilder().(*_BACnetValueSourceObjectBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetValueSourceObjectBuilder creates a BACnetValueSourceObjectBuilder
func (b *_BACnetValueSourceObject) CreateBACnetValueSourceObjectBuilder() BACnetValueSourceObjectBuilder {
	if b == nil {
		return NewBACnetValueSourceObjectBuilder()
	}
	return &_BACnetValueSourceObjectBuilder{_BACnetValueSourceObject: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetValueSourceObject) GetParent() BACnetValueSourceContract {
	return m.BACnetValueSourceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetValueSourceObject) GetObject() BACnetDeviceObjectReferenceEnclosed {
	return m.Object
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetValueSourceObject(structType any) BACnetValueSourceObject {
	if casted, ok := structType.(BACnetValueSourceObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetValueSourceObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetValueSourceObject) GetTypeName() string {
	return "BACnetValueSourceObject"
}

func (m *_BACnetValueSourceObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetValueSourceContract.(*_BACnetValueSource).getLengthInBits(ctx))

	// Simple field (object)
	lengthInBits += m.Object.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetValueSourceObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetValueSourceObject) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetValueSource) (__bACnetValueSourceObject BACnetValueSourceObject, err error) {
	m.BACnetValueSourceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetValueSourceObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetValueSourceObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	object, err := ReadSimpleField[BACnetDeviceObjectReferenceEnclosed](ctx, "object", ReadComplex[BACnetDeviceObjectReferenceEnclosed](BACnetDeviceObjectReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'object' field"))
	}
	m.Object = object

	if closeErr := readBuffer.CloseContext("BACnetValueSourceObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetValueSourceObject")
	}

	return m, nil
}

func (m *_BACnetValueSourceObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetValueSourceObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetValueSourceObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetValueSourceObject")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReferenceEnclosed](ctx, "object", m.GetObject(), WriteComplex[BACnetDeviceObjectReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'object' field")
		}

		if popErr := writeBuffer.PopContext("BACnetValueSourceObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetValueSourceObject")
		}
		return nil
	}
	return m.BACnetValueSourceContract.(*_BACnetValueSource).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetValueSourceObject) IsBACnetValueSourceObject() {}

func (m *_BACnetValueSourceObject) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetValueSourceObject) deepCopy() *_BACnetValueSourceObject {
	if m == nil {
		return nil
	}
	_BACnetValueSourceObjectCopy := &_BACnetValueSourceObject{
		m.BACnetValueSourceContract.(*_BACnetValueSource).deepCopy(),
		utils.DeepCopy[BACnetDeviceObjectReferenceEnclosed](m.Object),
	}
	m.BACnetValueSourceContract.(*_BACnetValueSource)._SubType = m
	return _BACnetValueSourceObjectCopy
}

func (m *_BACnetValueSourceObject) String() string {
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

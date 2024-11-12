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

// BACnetConstructedDataLoopAction is the corresponding interface of BACnetConstructedDataLoopAction
type BACnetConstructedDataLoopAction interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAction returns Action (property field)
	GetAction() BACnetActionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetActionTagged
	// IsBACnetConstructedDataLoopAction is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLoopAction()
	// CreateBuilder creates a BACnetConstructedDataLoopActionBuilder
	CreateBACnetConstructedDataLoopActionBuilder() BACnetConstructedDataLoopActionBuilder
}

// _BACnetConstructedDataLoopAction is the data-structure of this message
type _BACnetConstructedDataLoopAction struct {
	BACnetConstructedDataContract
	Action BACnetActionTagged
}

var _ BACnetConstructedDataLoopAction = (*_BACnetConstructedDataLoopAction)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLoopAction)(nil)

// NewBACnetConstructedDataLoopAction factory function for _BACnetConstructedDataLoopAction
func NewBACnetConstructedDataLoopAction(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, action BACnetActionTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLoopAction {
	if action == nil {
		panic("action of type BACnetActionTagged for BACnetConstructedDataLoopAction must not be nil")
	}
	_result := &_BACnetConstructedDataLoopAction{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Action:                        action,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLoopActionBuilder is a builder for BACnetConstructedDataLoopAction
type BACnetConstructedDataLoopActionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(action BACnetActionTagged) BACnetConstructedDataLoopActionBuilder
	// WithAction adds Action (property field)
	WithAction(BACnetActionTagged) BACnetConstructedDataLoopActionBuilder
	// WithActionBuilder adds Action (property field) which is build by the builder
	WithActionBuilder(func(BACnetActionTaggedBuilder) BACnetActionTaggedBuilder) BACnetConstructedDataLoopActionBuilder
	// Build builds the BACnetConstructedDataLoopAction or returns an error if something is wrong
	Build() (BACnetConstructedDataLoopAction, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLoopAction
}

// NewBACnetConstructedDataLoopActionBuilder() creates a BACnetConstructedDataLoopActionBuilder
func NewBACnetConstructedDataLoopActionBuilder() BACnetConstructedDataLoopActionBuilder {
	return &_BACnetConstructedDataLoopActionBuilder{_BACnetConstructedDataLoopAction: new(_BACnetConstructedDataLoopAction)}
}

type _BACnetConstructedDataLoopActionBuilder struct {
	*_BACnetConstructedDataLoopAction

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLoopActionBuilder) = (*_BACnetConstructedDataLoopActionBuilder)(nil)

func (b *_BACnetConstructedDataLoopActionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLoopActionBuilder) WithMandatoryFields(action BACnetActionTagged) BACnetConstructedDataLoopActionBuilder {
	return b.WithAction(action)
}

func (b *_BACnetConstructedDataLoopActionBuilder) WithAction(action BACnetActionTagged) BACnetConstructedDataLoopActionBuilder {
	b.Action = action
	return b
}

func (b *_BACnetConstructedDataLoopActionBuilder) WithActionBuilder(builderSupplier func(BACnetActionTaggedBuilder) BACnetActionTaggedBuilder) BACnetConstructedDataLoopActionBuilder {
	builder := builderSupplier(b.Action.CreateBACnetActionTaggedBuilder())
	var err error
	b.Action, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetActionTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLoopActionBuilder) Build() (BACnetConstructedDataLoopAction, error) {
	if b.Action == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'action' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLoopAction.deepCopy(), nil
}

func (b *_BACnetConstructedDataLoopActionBuilder) MustBuild() BACnetConstructedDataLoopAction {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLoopActionBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLoopActionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLoopActionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLoopActionBuilder().(*_BACnetConstructedDataLoopActionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLoopActionBuilder creates a BACnetConstructedDataLoopActionBuilder
func (b *_BACnetConstructedDataLoopAction) CreateBACnetConstructedDataLoopActionBuilder() BACnetConstructedDataLoopActionBuilder {
	if b == nil {
		return NewBACnetConstructedDataLoopActionBuilder()
	}
	return &_BACnetConstructedDataLoopActionBuilder{_BACnetConstructedDataLoopAction: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLoopAction) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LOOP
}

func (m *_BACnetConstructedDataLoopAction) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLoopAction) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLoopAction) GetAction() BACnetActionTagged {
	return m.Action
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLoopAction) GetActualValue() BACnetActionTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetActionTagged(m.GetAction())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLoopAction(structType any) BACnetConstructedDataLoopAction {
	if casted, ok := structType.(BACnetConstructedDataLoopAction); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoopAction); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLoopAction) GetTypeName() string {
	return "BACnetConstructedDataLoopAction"
}

func (m *_BACnetConstructedDataLoopAction) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (action)
	lengthInBits += m.Action.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLoopAction) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLoopAction) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLoopAction BACnetConstructedDataLoopAction, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoopAction"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLoopAction")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	action, err := ReadSimpleField[BACnetActionTagged](ctx, "action", ReadComplex[BACnetActionTagged](BACnetActionTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'action' field"))
	}
	m.Action = action

	actualValue, err := ReadVirtualField[BACnetActionTagged](ctx, "actualValue", (*BACnetActionTagged)(nil), action)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoopAction"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLoopAction")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLoopAction) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLoopAction) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoopAction"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLoopAction")
		}

		if err := WriteSimpleField[BACnetActionTagged](ctx, "action", m.GetAction(), WriteComplex[BACnetActionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'action' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoopAction"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLoopAction")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLoopAction) IsBACnetConstructedDataLoopAction() {}

func (m *_BACnetConstructedDataLoopAction) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLoopAction) deepCopy() *_BACnetConstructedDataLoopAction {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLoopActionCopy := &_BACnetConstructedDataLoopAction{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetActionTagged](m.Action),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLoopActionCopy
}

func (m *_BACnetConstructedDataLoopAction) String() string {
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

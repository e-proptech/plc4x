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

// BACnetConstructedDataProcessIdentifierFilter is the corresponding interface of BACnetConstructedDataProcessIdentifierFilter
type BACnetConstructedDataProcessIdentifierFilter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetProcessIdentifierFilter returns ProcessIdentifierFilter (property field)
	GetProcessIdentifierFilter() BACnetProcessIdSelection
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetProcessIdSelection
	// IsBACnetConstructedDataProcessIdentifierFilter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataProcessIdentifierFilter()
	// CreateBuilder creates a BACnetConstructedDataProcessIdentifierFilterBuilder
	CreateBACnetConstructedDataProcessIdentifierFilterBuilder() BACnetConstructedDataProcessIdentifierFilterBuilder
}

// _BACnetConstructedDataProcessIdentifierFilter is the data-structure of this message
type _BACnetConstructedDataProcessIdentifierFilter struct {
	BACnetConstructedDataContract
	ProcessIdentifierFilter BACnetProcessIdSelection
}

var _ BACnetConstructedDataProcessIdentifierFilter = (*_BACnetConstructedDataProcessIdentifierFilter)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataProcessIdentifierFilter)(nil)

// NewBACnetConstructedDataProcessIdentifierFilter factory function for _BACnetConstructedDataProcessIdentifierFilter
func NewBACnetConstructedDataProcessIdentifierFilter(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, processIdentifierFilter BACnetProcessIdSelection, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProcessIdentifierFilter {
	if processIdentifierFilter == nil {
		panic("processIdentifierFilter of type BACnetProcessIdSelection for BACnetConstructedDataProcessIdentifierFilter must not be nil")
	}
	_result := &_BACnetConstructedDataProcessIdentifierFilter{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ProcessIdentifierFilter:       processIdentifierFilter,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataProcessIdentifierFilterBuilder is a builder for BACnetConstructedDataProcessIdentifierFilter
type BACnetConstructedDataProcessIdentifierFilterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(processIdentifierFilter BACnetProcessIdSelection) BACnetConstructedDataProcessIdentifierFilterBuilder
	// WithProcessIdentifierFilter adds ProcessIdentifierFilter (property field)
	WithProcessIdentifierFilter(BACnetProcessIdSelection) BACnetConstructedDataProcessIdentifierFilterBuilder
	// WithProcessIdentifierFilterBuilder adds ProcessIdentifierFilter (property field) which is build by the builder
	WithProcessIdentifierFilterBuilder(func(BACnetProcessIdSelectionBuilder) BACnetProcessIdSelectionBuilder) BACnetConstructedDataProcessIdentifierFilterBuilder
	// Build builds the BACnetConstructedDataProcessIdentifierFilter or returns an error if something is wrong
	Build() (BACnetConstructedDataProcessIdentifierFilter, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataProcessIdentifierFilter
}

// NewBACnetConstructedDataProcessIdentifierFilterBuilder() creates a BACnetConstructedDataProcessIdentifierFilterBuilder
func NewBACnetConstructedDataProcessIdentifierFilterBuilder() BACnetConstructedDataProcessIdentifierFilterBuilder {
	return &_BACnetConstructedDataProcessIdentifierFilterBuilder{_BACnetConstructedDataProcessIdentifierFilter: new(_BACnetConstructedDataProcessIdentifierFilter)}
}

type _BACnetConstructedDataProcessIdentifierFilterBuilder struct {
	*_BACnetConstructedDataProcessIdentifierFilter

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataProcessIdentifierFilterBuilder) = (*_BACnetConstructedDataProcessIdentifierFilterBuilder)(nil)

func (b *_BACnetConstructedDataProcessIdentifierFilterBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataProcessIdentifierFilterBuilder) WithMandatoryFields(processIdentifierFilter BACnetProcessIdSelection) BACnetConstructedDataProcessIdentifierFilterBuilder {
	return b.WithProcessIdentifierFilter(processIdentifierFilter)
}

func (b *_BACnetConstructedDataProcessIdentifierFilterBuilder) WithProcessIdentifierFilter(processIdentifierFilter BACnetProcessIdSelection) BACnetConstructedDataProcessIdentifierFilterBuilder {
	b.ProcessIdentifierFilter = processIdentifierFilter
	return b
}

func (b *_BACnetConstructedDataProcessIdentifierFilterBuilder) WithProcessIdentifierFilterBuilder(builderSupplier func(BACnetProcessIdSelectionBuilder) BACnetProcessIdSelectionBuilder) BACnetConstructedDataProcessIdentifierFilterBuilder {
	builder := builderSupplier(b.ProcessIdentifierFilter.CreateBACnetProcessIdSelectionBuilder())
	var err error
	b.ProcessIdentifierFilter, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetProcessIdSelectionBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataProcessIdentifierFilterBuilder) Build() (BACnetConstructedDataProcessIdentifierFilter, error) {
	if b.ProcessIdentifierFilter == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'processIdentifierFilter' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataProcessIdentifierFilter.deepCopy(), nil
}

func (b *_BACnetConstructedDataProcessIdentifierFilterBuilder) MustBuild() BACnetConstructedDataProcessIdentifierFilter {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataProcessIdentifierFilterBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataProcessIdentifierFilterBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataProcessIdentifierFilterBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataProcessIdentifierFilterBuilder().(*_BACnetConstructedDataProcessIdentifierFilterBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataProcessIdentifierFilterBuilder creates a BACnetConstructedDataProcessIdentifierFilterBuilder
func (b *_BACnetConstructedDataProcessIdentifierFilter) CreateBACnetConstructedDataProcessIdentifierFilterBuilder() BACnetConstructedDataProcessIdentifierFilterBuilder {
	if b == nil {
		return NewBACnetConstructedDataProcessIdentifierFilterBuilder()
	}
	return &_BACnetConstructedDataProcessIdentifierFilterBuilder{_BACnetConstructedDataProcessIdentifierFilter: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROCESS_IDENTIFIER_FILTER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetProcessIdentifierFilter() BACnetProcessIdSelection {
	return m.ProcessIdentifierFilter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetActualValue() BACnetProcessIdSelection {
	ctx := context.Background()
	_ = ctx
	return CastBACnetProcessIdSelection(m.GetProcessIdentifierFilter())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProcessIdentifierFilter(structType any) BACnetConstructedDataProcessIdentifierFilter {
	if casted, ok := structType.(BACnetConstructedDataProcessIdentifierFilter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProcessIdentifierFilter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetTypeName() string {
	return "BACnetConstructedDataProcessIdentifierFilter"
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (processIdentifierFilter)
	lengthInBits += m.ProcessIdentifierFilter.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataProcessIdentifierFilter BACnetConstructedDataProcessIdentifierFilter, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProcessIdentifierFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProcessIdentifierFilter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	processIdentifierFilter, err := ReadSimpleField[BACnetProcessIdSelection](ctx, "processIdentifierFilter", ReadComplex[BACnetProcessIdSelection](BACnetProcessIdSelectionParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifierFilter' field"))
	}
	m.ProcessIdentifierFilter = processIdentifierFilter

	actualValue, err := ReadVirtualField[BACnetProcessIdSelection](ctx, "actualValue", (*BACnetProcessIdSelection)(nil), processIdentifierFilter)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProcessIdentifierFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProcessIdentifierFilter")
	}

	return m, nil
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProcessIdentifierFilter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProcessIdentifierFilter")
		}

		if err := WriteSimpleField[BACnetProcessIdSelection](ctx, "processIdentifierFilter", m.GetProcessIdentifierFilter(), WriteComplex[BACnetProcessIdSelection](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'processIdentifierFilter' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProcessIdentifierFilter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProcessIdentifierFilter")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) IsBACnetConstructedDataProcessIdentifierFilter() {
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) deepCopy() *_BACnetConstructedDataProcessIdentifierFilter {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataProcessIdentifierFilterCopy := &_BACnetConstructedDataProcessIdentifierFilter{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetProcessIdSelection](m.ProcessIdentifierFilter),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataProcessIdentifierFilterCopy
}

func (m *_BACnetConstructedDataProcessIdentifierFilter) String() string {
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

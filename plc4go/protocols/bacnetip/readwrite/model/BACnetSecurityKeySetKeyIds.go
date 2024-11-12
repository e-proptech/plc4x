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

// BACnetSecurityKeySetKeyIds is the corresponding interface of BACnetSecurityKeySetKeyIds
type BACnetSecurityKeySetKeyIds interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetKeyIds returns KeyIds (property field)
	GetKeyIds() []BACnetKeyIdentifier
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetSecurityKeySetKeyIds is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetSecurityKeySetKeyIds()
	// CreateBuilder creates a BACnetSecurityKeySetKeyIdsBuilder
	CreateBACnetSecurityKeySetKeyIdsBuilder() BACnetSecurityKeySetKeyIdsBuilder
}

// _BACnetSecurityKeySetKeyIds is the data-structure of this message
type _BACnetSecurityKeySetKeyIds struct {
	OpeningTag BACnetOpeningTag
	KeyIds     []BACnetKeyIdentifier
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetSecurityKeySetKeyIds = (*_BACnetSecurityKeySetKeyIds)(nil)

// NewBACnetSecurityKeySetKeyIds factory function for _BACnetSecurityKeySetKeyIds
func NewBACnetSecurityKeySetKeyIds(openingTag BACnetOpeningTag, keyIds []BACnetKeyIdentifier, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetSecurityKeySetKeyIds {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetSecurityKeySetKeyIds must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetSecurityKeySetKeyIds must not be nil")
	}
	return &_BACnetSecurityKeySetKeyIds{OpeningTag: openingTag, KeyIds: keyIds, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetSecurityKeySetKeyIdsBuilder is a builder for BACnetSecurityKeySetKeyIds
type BACnetSecurityKeySetKeyIdsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, keyIds []BACnetKeyIdentifier, closingTag BACnetClosingTag) BACnetSecurityKeySetKeyIdsBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetSecurityKeySetKeyIdsBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetSecurityKeySetKeyIdsBuilder
	// WithKeyIds adds KeyIds (property field)
	WithKeyIds(...BACnetKeyIdentifier) BACnetSecurityKeySetKeyIdsBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetSecurityKeySetKeyIdsBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetSecurityKeySetKeyIdsBuilder
	// Build builds the BACnetSecurityKeySetKeyIds or returns an error if something is wrong
	Build() (BACnetSecurityKeySetKeyIds, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetSecurityKeySetKeyIds
}

// NewBACnetSecurityKeySetKeyIdsBuilder() creates a BACnetSecurityKeySetKeyIdsBuilder
func NewBACnetSecurityKeySetKeyIdsBuilder() BACnetSecurityKeySetKeyIdsBuilder {
	return &_BACnetSecurityKeySetKeyIdsBuilder{_BACnetSecurityKeySetKeyIds: new(_BACnetSecurityKeySetKeyIds)}
}

type _BACnetSecurityKeySetKeyIdsBuilder struct {
	*_BACnetSecurityKeySetKeyIds

	err *utils.MultiError
}

var _ (BACnetSecurityKeySetKeyIdsBuilder) = (*_BACnetSecurityKeySetKeyIdsBuilder)(nil)

func (b *_BACnetSecurityKeySetKeyIdsBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, keyIds []BACnetKeyIdentifier, closingTag BACnetClosingTag) BACnetSecurityKeySetKeyIdsBuilder {
	return b.WithOpeningTag(openingTag).WithKeyIds(keyIds...).WithClosingTag(closingTag)
}

func (b *_BACnetSecurityKeySetKeyIdsBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetSecurityKeySetKeyIdsBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetSecurityKeySetKeyIdsBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetSecurityKeySetKeyIdsBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetSecurityKeySetKeyIdsBuilder) WithKeyIds(keyIds ...BACnetKeyIdentifier) BACnetSecurityKeySetKeyIdsBuilder {
	b.KeyIds = keyIds
	return b
}

func (b *_BACnetSecurityKeySetKeyIdsBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetSecurityKeySetKeyIdsBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetSecurityKeySetKeyIdsBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetSecurityKeySetKeyIdsBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetSecurityKeySetKeyIdsBuilder) Build() (BACnetSecurityKeySetKeyIds, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetSecurityKeySetKeyIds.deepCopy(), nil
}

func (b *_BACnetSecurityKeySetKeyIdsBuilder) MustBuild() BACnetSecurityKeySetKeyIds {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetSecurityKeySetKeyIdsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetSecurityKeySetKeyIdsBuilder().(*_BACnetSecurityKeySetKeyIdsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetSecurityKeySetKeyIdsBuilder creates a BACnetSecurityKeySetKeyIdsBuilder
func (b *_BACnetSecurityKeySetKeyIds) CreateBACnetSecurityKeySetKeyIdsBuilder() BACnetSecurityKeySetKeyIdsBuilder {
	if b == nil {
		return NewBACnetSecurityKeySetKeyIdsBuilder()
	}
	return &_BACnetSecurityKeySetKeyIdsBuilder{_BACnetSecurityKeySetKeyIds: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSecurityKeySetKeyIds) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetSecurityKeySetKeyIds) GetKeyIds() []BACnetKeyIdentifier {
	return m.KeyIds
}

func (m *_BACnetSecurityKeySetKeyIds) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetSecurityKeySetKeyIds(structType any) BACnetSecurityKeySetKeyIds {
	if casted, ok := structType.(BACnetSecurityKeySetKeyIds); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSecurityKeySetKeyIds); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSecurityKeySetKeyIds) GetTypeName() string {
	return "BACnetSecurityKeySetKeyIds"
}

func (m *_BACnetSecurityKeySetKeyIds) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.KeyIds) > 0 {
		for _, element := range m.KeyIds {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSecurityKeySetKeyIds) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSecurityKeySetKeyIdsParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetSecurityKeySetKeyIds, error) {
	return BACnetSecurityKeySetKeyIdsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetSecurityKeySetKeyIdsParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSecurityKeySetKeyIds, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSecurityKeySetKeyIds, error) {
		return BACnetSecurityKeySetKeyIdsParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetSecurityKeySetKeyIdsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetSecurityKeySetKeyIds, error) {
	v, err := (&_BACnetSecurityKeySetKeyIds{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetSecurityKeySetKeyIds) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetSecurityKeySetKeyIds BACnetSecurityKeySetKeyIds, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSecurityKeySetKeyIds"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSecurityKeySetKeyIds")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	keyIds, err := ReadTerminatedArrayField[BACnetKeyIdentifier](ctx, "keyIds", ReadComplex[BACnetKeyIdentifier](BACnetKeyIdentifierParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keyIds' field"))
	}
	m.KeyIds = keyIds

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetSecurityKeySetKeyIds"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSecurityKeySetKeyIds")
	}

	return m, nil
}

func (m *_BACnetSecurityKeySetKeyIds) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSecurityKeySetKeyIds) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetSecurityKeySetKeyIds"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSecurityKeySetKeyIds")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "keyIds", m.GetKeyIds(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'keyIds' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSecurityKeySetKeyIds"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSecurityKeySetKeyIds")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetSecurityKeySetKeyIds) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetSecurityKeySetKeyIds) IsBACnetSecurityKeySetKeyIds() {}

func (m *_BACnetSecurityKeySetKeyIds) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetSecurityKeySetKeyIds) deepCopy() *_BACnetSecurityKeySetKeyIds {
	if m == nil {
		return nil
	}
	_BACnetSecurityKeySetKeyIdsCopy := &_BACnetSecurityKeySetKeyIds{
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopySlice[BACnetKeyIdentifier, BACnetKeyIdentifier](m.KeyIds),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
		m.TagNumber,
	}
	return _BACnetSecurityKeySetKeyIdsCopy
}

func (m *_BACnetSecurityKeySetKeyIds) String() string {
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

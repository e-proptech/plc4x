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

// BACnetApplicationTag is the corresponding interface of BACnetApplicationTag
type BACnetApplicationTag interface {
	BACnetApplicationTagContract
	BACnetApplicationTagRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetApplicationTag is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTag()
	// CreateBuilder creates a BACnetApplicationTagBuilder
	CreateBACnetApplicationTagBuilder() BACnetApplicationTagBuilder
}

// BACnetApplicationTagContract provides a set of functions which can be overwritten by a sub struct
type BACnetApplicationTagContract interface {
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetActualTagNumber returns ActualTagNumber (virtual field)
	GetActualTagNumber() uint8
	// GetActualLength returns ActualLength (virtual field)
	GetActualLength() uint32
	// IsBACnetApplicationTag is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTag()
	// CreateBuilder creates a BACnetApplicationTagBuilder
	CreateBACnetApplicationTagBuilder() BACnetApplicationTagBuilder
}

// BACnetApplicationTagRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetApplicationTagRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetActualTagNumber returns ActualTagNumber (discriminator field)
	GetActualTagNumber() uint8
}

// _BACnetApplicationTag is the data-structure of this message
type _BACnetApplicationTag struct {
	_SubType interface {
		BACnetApplicationTagContract
		BACnetApplicationTagRequirements
	}
	Header BACnetTagHeader
}

var _ BACnetApplicationTagContract = (*_BACnetApplicationTag)(nil)

// NewBACnetApplicationTag factory function for _BACnetApplicationTag
func NewBACnetApplicationTag(header BACnetTagHeader) *_BACnetApplicationTag {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetApplicationTag must not be nil")
	}
	return &_BACnetApplicationTag{Header: header}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetApplicationTagBuilder is a builder for BACnetApplicationTag
type BACnetApplicationTagBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader) BACnetApplicationTagBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetApplicationTagBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetApplicationTagBuilder
	// AsBACnetApplicationTagNull converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagNull() interface {
		BACnetApplicationTagNullBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagBoolean converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagBoolean() interface {
		BACnetApplicationTagBooleanBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagUnsignedInteger converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagUnsignedInteger() interface {
		BACnetApplicationTagUnsignedIntegerBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagSignedInteger converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagSignedInteger() interface {
		BACnetApplicationTagSignedIntegerBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagReal converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagReal() interface {
		BACnetApplicationTagRealBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagDouble converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagDouble() interface {
		BACnetApplicationTagDoubleBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagOctetString converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagOctetString() interface {
		BACnetApplicationTagOctetStringBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagCharacterString converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagCharacterString() interface {
		BACnetApplicationTagCharacterStringBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagBitString converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagBitString() interface {
		BACnetApplicationTagBitStringBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagEnumerated converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagEnumerated() interface {
		BACnetApplicationTagEnumeratedBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagDate converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagDate() interface {
		BACnetApplicationTagDateBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagTime converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagTime() interface {
		BACnetApplicationTagTimeBuilder
		Done() BACnetApplicationTagBuilder
	}
	// AsBACnetApplicationTagObjectIdentifier converts this build to a subType of BACnetApplicationTag. It is always possible to return to current builder using Done()
	AsBACnetApplicationTagObjectIdentifier() interface {
		BACnetApplicationTagObjectIdentifierBuilder
		Done() BACnetApplicationTagBuilder
	}
	// Build builds the BACnetApplicationTag or returns an error if something is wrong
	PartialBuild() (BACnetApplicationTagContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetApplicationTagContract
	// Build builds the BACnetApplicationTag or returns an error if something is wrong
	Build() (BACnetApplicationTag, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetApplicationTag
}

// NewBACnetApplicationTagBuilder() creates a BACnetApplicationTagBuilder
func NewBACnetApplicationTagBuilder() BACnetApplicationTagBuilder {
	return &_BACnetApplicationTagBuilder{_BACnetApplicationTag: new(_BACnetApplicationTag)}
}

type _BACnetApplicationTagChildBuilder interface {
	utils.Copyable
	setParent(BACnetApplicationTagContract)
	buildForBACnetApplicationTag() (BACnetApplicationTag, error)
}

type _BACnetApplicationTagBuilder struct {
	*_BACnetApplicationTag

	childBuilder _BACnetApplicationTagChildBuilder

	err *utils.MultiError
}

var _ (BACnetApplicationTagBuilder) = (*_BACnetApplicationTagBuilder)(nil)

func (b *_BACnetApplicationTagBuilder) WithMandatoryFields(header BACnetTagHeader) BACnetApplicationTagBuilder {
	return b.WithHeader(header)
}

func (b *_BACnetApplicationTagBuilder) WithHeader(header BACnetTagHeader) BACnetApplicationTagBuilder {
	b.Header = header
	return b
}

func (b *_BACnetApplicationTagBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetApplicationTagBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetApplicationTagBuilder) PartialBuild() (BACnetApplicationTagContract, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetApplicationTag.deepCopy(), nil
}

func (b *_BACnetApplicationTagBuilder) PartialMustBuild() BACnetApplicationTagContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagNull() interface {
	BACnetApplicationTagNullBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagNullBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagNullBuilder().(*_BACnetApplicationTagNullBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagBoolean() interface {
	BACnetApplicationTagBooleanBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagBooleanBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagBooleanBuilder().(*_BACnetApplicationTagBooleanBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagUnsignedInteger() interface {
	BACnetApplicationTagUnsignedIntegerBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagUnsignedIntegerBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagUnsignedIntegerBuilder().(*_BACnetApplicationTagUnsignedIntegerBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagSignedInteger() interface {
	BACnetApplicationTagSignedIntegerBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagSignedIntegerBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagSignedIntegerBuilder().(*_BACnetApplicationTagSignedIntegerBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagReal() interface {
	BACnetApplicationTagRealBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagRealBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagRealBuilder().(*_BACnetApplicationTagRealBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagDouble() interface {
	BACnetApplicationTagDoubleBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagDoubleBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagDoubleBuilder().(*_BACnetApplicationTagDoubleBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagOctetString() interface {
	BACnetApplicationTagOctetStringBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagOctetStringBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagOctetStringBuilder().(*_BACnetApplicationTagOctetStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagCharacterString() interface {
	BACnetApplicationTagCharacterStringBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagCharacterStringBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagCharacterStringBuilder().(*_BACnetApplicationTagCharacterStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagBitString() interface {
	BACnetApplicationTagBitStringBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagBitStringBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagBitStringBuilder().(*_BACnetApplicationTagBitStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagEnumerated() interface {
	BACnetApplicationTagEnumeratedBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagEnumeratedBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagEnumeratedBuilder().(*_BACnetApplicationTagEnumeratedBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagDate() interface {
	BACnetApplicationTagDateBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagDateBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagDateBuilder().(*_BACnetApplicationTagDateBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagTime() interface {
	BACnetApplicationTagTimeBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagTimeBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagTimeBuilder().(*_BACnetApplicationTagTimeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) AsBACnetApplicationTagObjectIdentifier() interface {
	BACnetApplicationTagObjectIdentifierBuilder
	Done() BACnetApplicationTagBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetApplicationTagObjectIdentifierBuilder
		Done() BACnetApplicationTagBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetApplicationTagObjectIdentifierBuilder().(*_BACnetApplicationTagObjectIdentifierBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetApplicationTagBuilder) Build() (BACnetApplicationTag, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetApplicationTag()
}

func (b *_BACnetApplicationTagBuilder) MustBuild() BACnetApplicationTag {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetApplicationTagBuilder) DeepCopy() any {
	_copy := b.CreateBACnetApplicationTagBuilder().(*_BACnetApplicationTagBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetApplicationTagChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetApplicationTagBuilder creates a BACnetApplicationTagBuilder
func (b *_BACnetApplicationTag) CreateBACnetApplicationTagBuilder() BACnetApplicationTagBuilder {
	if b == nil {
		return NewBACnetApplicationTagBuilder()
	}
	return &_BACnetApplicationTagBuilder{_BACnetApplicationTag: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTag) GetHeader() BACnetTagHeader {
	return m.Header
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetApplicationTag) GetActualTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetHeader().GetActualTagNumber())
}

func (pm *_BACnetApplicationTag) GetActualLength() uint32 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint32(m.GetHeader().GetActualLength())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTag(structType any) BACnetApplicationTag {
	if casted, ok := structType.(BACnetApplicationTag); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTag); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTag) GetTypeName() string {
	return "BACnetApplicationTag"
}

func (m *_BACnetApplicationTag) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTag) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_BACnetApplicationTag) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetApplicationTagParse[T BACnetApplicationTag](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetApplicationTagParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetApplicationTagParseWithBufferProducer[T BACnetApplicationTag]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetApplicationTagParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetApplicationTagParseWithBuffer[T BACnetApplicationTag](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetApplicationTag{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetApplicationTag) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetApplicationTag BACnetApplicationTag, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTag")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "should be a application tag"})
	}

	actualTagNumber, err := ReadVirtualField[uint8](ctx, "actualTagNumber", (*uint8)(nil), header.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualTagNumber' field"))
	}
	_ = actualTagNumber

	actualLength, err := ReadVirtualField[uint32](ctx, "actualLength", (*uint32)(nil), header.GetActualLength())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualLength' field"))
	}
	_ = actualLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetApplicationTag
	switch {
	case actualTagNumber == 0x0: // BACnetApplicationTagNull
		if _child, err = new(_BACnetApplicationTagNull).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagNull for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0x1: // BACnetApplicationTagBoolean
		if _child, err = new(_BACnetApplicationTagBoolean).parse(ctx, readBuffer, m, header); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagBoolean for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0x2: // BACnetApplicationTagUnsignedInteger
		if _child, err = new(_BACnetApplicationTagUnsignedInteger).parse(ctx, readBuffer, m, header); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagUnsignedInteger for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0x3: // BACnetApplicationTagSignedInteger
		if _child, err = new(_BACnetApplicationTagSignedInteger).parse(ctx, readBuffer, m, header); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagSignedInteger for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0x4: // BACnetApplicationTagReal
		if _child, err = new(_BACnetApplicationTagReal).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagReal for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0x5: // BACnetApplicationTagDouble
		if _child, err = new(_BACnetApplicationTagDouble).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagDouble for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0x6: // BACnetApplicationTagOctetString
		if _child, err = new(_BACnetApplicationTagOctetString).parse(ctx, readBuffer, m, header); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagOctetString for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0x7: // BACnetApplicationTagCharacterString
		if _child, err = new(_BACnetApplicationTagCharacterString).parse(ctx, readBuffer, m, header); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagCharacterString for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0x8: // BACnetApplicationTagBitString
		if _child, err = new(_BACnetApplicationTagBitString).parse(ctx, readBuffer, m, header); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagBitString for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0x9: // BACnetApplicationTagEnumerated
		if _child, err = new(_BACnetApplicationTagEnumerated).parse(ctx, readBuffer, m, header); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagEnumerated for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0xA: // BACnetApplicationTagDate
		if _child, err = new(_BACnetApplicationTagDate).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagDate for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0xB: // BACnetApplicationTagTime
		if _child, err = new(_BACnetApplicationTagTime).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagTime for type-switch of BACnetApplicationTag")
		}
	case actualTagNumber == 0xC: // BACnetApplicationTagObjectIdentifier
		if _child, err = new(_BACnetApplicationTagObjectIdentifier).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetApplicationTagObjectIdentifier for type-switch of BACnetApplicationTag")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [actualTagNumber=%v]", actualTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetApplicationTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTag")
	}

	return _child, nil
}

func (pm *_BACnetApplicationTag) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetApplicationTag, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetApplicationTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTag")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}
	// Virtual field
	actualTagNumber := m.GetActualTagNumber()
	_ = actualTagNumber
	if _actualTagNumberErr := writeBuffer.WriteVirtual(ctx, "actualTagNumber", m.GetActualTagNumber()); _actualTagNumberErr != nil {
		return errors.Wrap(_actualTagNumberErr, "Error serializing 'actualTagNumber' field")
	}
	// Virtual field
	actualLength := m.GetActualLength()
	_ = actualLength
	if _actualLengthErr := writeBuffer.WriteVirtual(ctx, "actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetApplicationTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetApplicationTag")
	}
	return nil
}

func (m *_BACnetApplicationTag) IsBACnetApplicationTag() {}

func (m *_BACnetApplicationTag) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetApplicationTag) deepCopy() *_BACnetApplicationTag {
	if m == nil {
		return nil
	}
	_BACnetApplicationTagCopy := &_BACnetApplicationTag{
		nil, // will be set by child
		utils.DeepCopy[BACnetTagHeader](m.Header),
	}
	return _BACnetApplicationTagCopy
}

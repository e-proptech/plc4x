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

// BACnetLogRecordLogDatum is the corresponding interface of BACnetLogRecordLogDatum
type BACnetLogRecordLogDatum interface {
	BACnetLogRecordLogDatumContract
	BACnetLogRecordLogDatumRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetLogRecordLogDatum is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogRecordLogDatum()
	// CreateBuilder creates a BACnetLogRecordLogDatumBuilder
	CreateBACnetLogRecordLogDatumBuilder() BACnetLogRecordLogDatumBuilder
}

// BACnetLogRecordLogDatumContract provides a set of functions which can be overwritten by a sub struct
type BACnetLogRecordLogDatumContract interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetTagNumber() returns a parser argument
	GetTagNumber() uint8
	// IsBACnetLogRecordLogDatum is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogRecordLogDatum()
	// CreateBuilder creates a BACnetLogRecordLogDatumBuilder
	CreateBACnetLogRecordLogDatumBuilder() BACnetLogRecordLogDatumBuilder
}

// BACnetLogRecordLogDatumRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetLogRecordLogDatumRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetLogRecordLogDatum is the data-structure of this message
type _BACnetLogRecordLogDatum struct {
	_SubType interface {
		BACnetLogRecordLogDatumContract
		BACnetLogRecordLogDatumRequirements
	}
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetLogRecordLogDatumContract = (*_BACnetLogRecordLogDatum)(nil)

// NewBACnetLogRecordLogDatum factory function for _BACnetLogRecordLogDatum
func NewBACnetLogRecordLogDatum(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatum {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetLogRecordLogDatum must not be nil")
	}
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetLogRecordLogDatum must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetLogRecordLogDatum must not be nil")
	}
	return &_BACnetLogRecordLogDatum{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogRecordLogDatumBuilder is a builder for BACnetLogRecordLogDatum
type BACnetLogRecordLogDatumBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) BACnetLogRecordLogDatumBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetLogRecordLogDatumBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetLogRecordLogDatumBuilder
	// WithPeekedTagHeader adds PeekedTagHeader (property field)
	WithPeekedTagHeader(BACnetTagHeader) BACnetLogRecordLogDatumBuilder
	// WithPeekedTagHeaderBuilder adds PeekedTagHeader (property field) which is build by the builder
	WithPeekedTagHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLogRecordLogDatumBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetLogRecordLogDatumBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetLogRecordLogDatumBuilder
	// AsBACnetLogRecordLogDatumLogStatus converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumLogStatus() interface {
		BACnetLogRecordLogDatumLogStatusBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// AsBACnetLogRecordLogDatumBooleanValue converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumBooleanValue() interface {
		BACnetLogRecordLogDatumBooleanValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// AsBACnetLogRecordLogDatumRealValue converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumRealValue() interface {
		BACnetLogRecordLogDatumRealValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// AsBACnetLogRecordLogDatumEnumeratedValue converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumEnumeratedValue() interface {
		BACnetLogRecordLogDatumEnumeratedValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// AsBACnetLogRecordLogDatumUnsignedValue converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumUnsignedValue() interface {
		BACnetLogRecordLogDatumUnsignedValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// AsBACnetLogRecordLogDatumIntegerValue converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumIntegerValue() interface {
		BACnetLogRecordLogDatumIntegerValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// AsBACnetLogRecordLogDatumBitStringValue converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumBitStringValue() interface {
		BACnetLogRecordLogDatumBitStringValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// AsBACnetLogRecordLogDatumNullValue converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumNullValue() interface {
		BACnetLogRecordLogDatumNullValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// AsBACnetLogRecordLogDatumFailure converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumFailure() interface {
		BACnetLogRecordLogDatumFailureBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// AsBACnetLogRecordLogDatumTimeChange converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumTimeChange() interface {
		BACnetLogRecordLogDatumTimeChangeBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// AsBACnetLogRecordLogDatumAnyValue converts this build to a subType of BACnetLogRecordLogDatum. It is always possible to return to current builder using Done()
	AsBACnetLogRecordLogDatumAnyValue() interface {
		BACnetLogRecordLogDatumAnyValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}
	// Build builds the BACnetLogRecordLogDatum or returns an error if something is wrong
	PartialBuild() (BACnetLogRecordLogDatumContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() BACnetLogRecordLogDatumContract
	// Build builds the BACnetLogRecordLogDatum or returns an error if something is wrong
	Build() (BACnetLogRecordLogDatum, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogRecordLogDatum
}

// NewBACnetLogRecordLogDatumBuilder() creates a BACnetLogRecordLogDatumBuilder
func NewBACnetLogRecordLogDatumBuilder() BACnetLogRecordLogDatumBuilder {
	return &_BACnetLogRecordLogDatumBuilder{_BACnetLogRecordLogDatum: new(_BACnetLogRecordLogDatum)}
}

type _BACnetLogRecordLogDatumChildBuilder interface {
	utils.Copyable
	setParent(BACnetLogRecordLogDatumContract)
	buildForBACnetLogRecordLogDatum() (BACnetLogRecordLogDatum, error)
}

type _BACnetLogRecordLogDatumBuilder struct {
	*_BACnetLogRecordLogDatum

	childBuilder _BACnetLogRecordLogDatumChildBuilder

	err *utils.MultiError
}

var _ (BACnetLogRecordLogDatumBuilder) = (*_BACnetLogRecordLogDatumBuilder)(nil)

func (b *_BACnetLogRecordLogDatumBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) BACnetLogRecordLogDatumBuilder {
	return b.WithOpeningTag(openingTag).WithPeekedTagHeader(peekedTagHeader).WithClosingTag(closingTag)
}

func (b *_BACnetLogRecordLogDatumBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetLogRecordLogDatumBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetLogRecordLogDatumBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetLogRecordLogDatumBuilder {
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

func (b *_BACnetLogRecordLogDatumBuilder) WithPeekedTagHeader(peekedTagHeader BACnetTagHeader) BACnetLogRecordLogDatumBuilder {
	b.PeekedTagHeader = peekedTagHeader
	return b
}

func (b *_BACnetLogRecordLogDatumBuilder) WithPeekedTagHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLogRecordLogDatumBuilder {
	builder := builderSupplier(b.PeekedTagHeader.CreateBACnetTagHeaderBuilder())
	var err error
	b.PeekedTagHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetLogRecordLogDatumBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetLogRecordLogDatumBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetLogRecordLogDatumBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetLogRecordLogDatumBuilder {
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

func (b *_BACnetLogRecordLogDatumBuilder) PartialBuild() (BACnetLogRecordLogDatumContract, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.PeekedTagHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'peekedTagHeader' not set"))
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
	return b._BACnetLogRecordLogDatum.deepCopy(), nil
}

func (b *_BACnetLogRecordLogDatumBuilder) PartialMustBuild() BACnetLogRecordLogDatumContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumLogStatus() interface {
	BACnetLogRecordLogDatumLogStatusBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumLogStatusBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumLogStatusBuilder().(*_BACnetLogRecordLogDatumLogStatusBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumBooleanValue() interface {
	BACnetLogRecordLogDatumBooleanValueBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumBooleanValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumBooleanValueBuilder().(*_BACnetLogRecordLogDatumBooleanValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumRealValue() interface {
	BACnetLogRecordLogDatumRealValueBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumRealValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumRealValueBuilder().(*_BACnetLogRecordLogDatumRealValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumEnumeratedValue() interface {
	BACnetLogRecordLogDatumEnumeratedValueBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumEnumeratedValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumEnumeratedValueBuilder().(*_BACnetLogRecordLogDatumEnumeratedValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumUnsignedValue() interface {
	BACnetLogRecordLogDatumUnsignedValueBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumUnsignedValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumUnsignedValueBuilder().(*_BACnetLogRecordLogDatumUnsignedValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumIntegerValue() interface {
	BACnetLogRecordLogDatumIntegerValueBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumIntegerValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumIntegerValueBuilder().(*_BACnetLogRecordLogDatumIntegerValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumBitStringValue() interface {
	BACnetLogRecordLogDatumBitStringValueBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumBitStringValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumBitStringValueBuilder().(*_BACnetLogRecordLogDatumBitStringValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumNullValue() interface {
	BACnetLogRecordLogDatumNullValueBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumNullValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumNullValueBuilder().(*_BACnetLogRecordLogDatumNullValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumFailure() interface {
	BACnetLogRecordLogDatumFailureBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumFailureBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumFailureBuilder().(*_BACnetLogRecordLogDatumFailureBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumTimeChange() interface {
	BACnetLogRecordLogDatumTimeChangeBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumTimeChangeBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumTimeChangeBuilder().(*_BACnetLogRecordLogDatumTimeChangeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) AsBACnetLogRecordLogDatumAnyValue() interface {
	BACnetLogRecordLogDatumAnyValueBuilder
	Done() BACnetLogRecordLogDatumBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		BACnetLogRecordLogDatumAnyValueBuilder
		Done() BACnetLogRecordLogDatumBuilder
	}); ok {
		return cb
	}
	cb := NewBACnetLogRecordLogDatumAnyValueBuilder().(*_BACnetLogRecordLogDatumAnyValueBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_BACnetLogRecordLogDatumBuilder) Build() (BACnetLogRecordLogDatum, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForBACnetLogRecordLogDatum()
}

func (b *_BACnetLogRecordLogDatumBuilder) MustBuild() BACnetLogRecordLogDatum {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLogRecordLogDatumBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLogRecordLogDatumBuilder().(*_BACnetLogRecordLogDatumBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_BACnetLogRecordLogDatumChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLogRecordLogDatumBuilder creates a BACnetLogRecordLogDatumBuilder
func (b *_BACnetLogRecordLogDatum) CreateBACnetLogRecordLogDatumBuilder() BACnetLogRecordLogDatumBuilder {
	if b == nil {
		return NewBACnetLogRecordLogDatumBuilder()
	}
	return &_BACnetLogRecordLogDatumBuilder{_BACnetLogRecordLogDatum: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatum) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetLogRecordLogDatum) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetLogRecordLogDatum) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetLogRecordLogDatum) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatum(structType any) BACnetLogRecordLogDatum {
	if casted, ok := structType.(BACnetLogRecordLogDatum); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatum); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatum) GetTypeName() string {
	return "BACnetLogRecordLogDatum"
}

func (m *_BACnetLogRecordLogDatum) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatum) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_BACnetLogRecordLogDatum) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetLogRecordLogDatumParse[T BACnetLogRecordLogDatum](ctx context.Context, theBytes []byte, tagNumber uint8) (T, error) {
	return BACnetLogRecordLogDatumParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLogRecordLogDatumParseWithBufferProducer[T BACnetLogRecordLogDatum](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetLogRecordLogDatumParseWithBuffer[T](ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetLogRecordLogDatumParseWithBuffer[T BACnetLogRecordLogDatum](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (T, error) {
	v, err := (&_BACnetLogRecordLogDatum{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
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

func (m *_BACnetLogRecordLogDatum) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetLogRecordLogDatum BACnetLogRecordLogDatum, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatum")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetLogRecordLogDatum
	switch {
	case peekedTagNumber == uint8(0): // BACnetLogRecordLogDatumLogStatus
		if _child, err = new(_BACnetLogRecordLogDatumLogStatus).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumLogStatus for type-switch of BACnetLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(1): // BACnetLogRecordLogDatumBooleanValue
		if _child, err = new(_BACnetLogRecordLogDatumBooleanValue).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumBooleanValue for type-switch of BACnetLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(2): // BACnetLogRecordLogDatumRealValue
		if _child, err = new(_BACnetLogRecordLogDatumRealValue).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumRealValue for type-switch of BACnetLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(3): // BACnetLogRecordLogDatumEnumeratedValue
		if _child, err = new(_BACnetLogRecordLogDatumEnumeratedValue).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumEnumeratedValue for type-switch of BACnetLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(4): // BACnetLogRecordLogDatumUnsignedValue
		if _child, err = new(_BACnetLogRecordLogDatumUnsignedValue).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumUnsignedValue for type-switch of BACnetLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(5): // BACnetLogRecordLogDatumIntegerValue
		if _child, err = new(_BACnetLogRecordLogDatumIntegerValue).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumIntegerValue for type-switch of BACnetLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(6): // BACnetLogRecordLogDatumBitStringValue
		if _child, err = new(_BACnetLogRecordLogDatumBitStringValue).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumBitStringValue for type-switch of BACnetLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(7): // BACnetLogRecordLogDatumNullValue
		if _child, err = new(_BACnetLogRecordLogDatumNullValue).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumNullValue for type-switch of BACnetLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(8): // BACnetLogRecordLogDatumFailure
		if _child, err = new(_BACnetLogRecordLogDatumFailure).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumFailure for type-switch of BACnetLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(9): // BACnetLogRecordLogDatumTimeChange
		if _child, err = new(_BACnetLogRecordLogDatumTimeChange).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumTimeChange for type-switch of BACnetLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(10): // BACnetLogRecordLogDatumAnyValue
		if _child, err = new(_BACnetLogRecordLogDatumAnyValue).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogRecordLogDatumAnyValue for type-switch of BACnetLogRecordLogDatum")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatum")
	}

	return _child, nil
}

func (pm *_BACnetLogRecordLogDatum) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetLogRecordLogDatum, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatum"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatum")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatum"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatum")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLogRecordLogDatum) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetLogRecordLogDatum) IsBACnetLogRecordLogDatum() {}

func (m *_BACnetLogRecordLogDatum) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogRecordLogDatum) deepCopy() *_BACnetLogRecordLogDatum {
	if m == nil {
		return nil
	}
	_BACnetLogRecordLogDatumCopy := &_BACnetLogRecordLogDatum{
		nil, // will be set by child
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetTagHeader](m.PeekedTagHeader),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
		m.TagNumber,
	}
	return _BACnetLogRecordLogDatumCopy
}

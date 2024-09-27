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

// CommandSpecificDataItem is the corresponding interface of CommandSpecificDataItem
type CommandSpecificDataItem interface {
	CommandSpecificDataItemContract
	CommandSpecificDataItemRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsCommandSpecificDataItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCommandSpecificDataItem()
	// CreateBuilder creates a CommandSpecificDataItemBuilder
	CreateCommandSpecificDataItemBuilder() CommandSpecificDataItemBuilder
}

// CommandSpecificDataItemContract provides a set of functions which can be overwritten by a sub struct
type CommandSpecificDataItemContract interface {
	// IsCommandSpecificDataItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCommandSpecificDataItem()
	// CreateBuilder creates a CommandSpecificDataItemBuilder
	CreateCommandSpecificDataItemBuilder() CommandSpecificDataItemBuilder
}

// CommandSpecificDataItemRequirements provides a set of functions which need to be implemented by a sub struct
type CommandSpecificDataItemRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetItemType returns ItemType (discriminator field)
	GetItemType() uint16
}

// _CommandSpecificDataItem is the data-structure of this message
type _CommandSpecificDataItem struct {
	_SubType CommandSpecificDataItem
}

var _ CommandSpecificDataItemContract = (*_CommandSpecificDataItem)(nil)

// NewCommandSpecificDataItem factory function for _CommandSpecificDataItem
func NewCommandSpecificDataItem() *_CommandSpecificDataItem {
	return &_CommandSpecificDataItem{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CommandSpecificDataItemBuilder is a builder for CommandSpecificDataItem
type CommandSpecificDataItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() CommandSpecificDataItemBuilder
	// AsCipIdentity converts this build to a subType of CommandSpecificDataItem. It is always possible to return to current builder using Done()
	AsCipIdentity() interface {
		CipIdentityBuilder
		Done() CommandSpecificDataItemBuilder
	}
	// AsCipSecurityInformation converts this build to a subType of CommandSpecificDataItem. It is always possible to return to current builder using Done()
	AsCipSecurityInformation() interface {
		CipSecurityInformationBuilder
		Done() CommandSpecificDataItemBuilder
	}
	// Build builds the CommandSpecificDataItem or returns an error if something is wrong
	PartialBuild() (CommandSpecificDataItemContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() CommandSpecificDataItemContract
	// Build builds the CommandSpecificDataItem or returns an error if something is wrong
	Build() (CommandSpecificDataItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CommandSpecificDataItem
}

// NewCommandSpecificDataItemBuilder() creates a CommandSpecificDataItemBuilder
func NewCommandSpecificDataItemBuilder() CommandSpecificDataItemBuilder {
	return &_CommandSpecificDataItemBuilder{_CommandSpecificDataItem: new(_CommandSpecificDataItem)}
}

type _CommandSpecificDataItemChildBuilder interface {
	utils.Copyable
	setParent(CommandSpecificDataItemContract)
	buildForCommandSpecificDataItem() (CommandSpecificDataItem, error)
}

type _CommandSpecificDataItemBuilder struct {
	*_CommandSpecificDataItem

	childBuilder _CommandSpecificDataItemChildBuilder

	err *utils.MultiError
}

var _ (CommandSpecificDataItemBuilder) = (*_CommandSpecificDataItemBuilder)(nil)

func (b *_CommandSpecificDataItemBuilder) WithMandatoryFields() CommandSpecificDataItemBuilder {
	return b
}

func (b *_CommandSpecificDataItemBuilder) PartialBuild() (CommandSpecificDataItemContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CommandSpecificDataItem.deepCopy(), nil
}

func (b *_CommandSpecificDataItemBuilder) PartialMustBuild() CommandSpecificDataItemContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CommandSpecificDataItemBuilder) AsCipIdentity() interface {
	CipIdentityBuilder
	Done() CommandSpecificDataItemBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipIdentityBuilder
		Done() CommandSpecificDataItemBuilder
	}); ok {
		return cb
	}
	cb := NewCipIdentityBuilder().(*_CipIdentityBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CommandSpecificDataItemBuilder) AsCipSecurityInformation() interface {
	CipSecurityInformationBuilder
	Done() CommandSpecificDataItemBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CipSecurityInformationBuilder
		Done() CommandSpecificDataItemBuilder
	}); ok {
		return cb
	}
	cb := NewCipSecurityInformationBuilder().(*_CipSecurityInformationBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CommandSpecificDataItemBuilder) Build() (CommandSpecificDataItem, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForCommandSpecificDataItem()
}

func (b *_CommandSpecificDataItemBuilder) MustBuild() CommandSpecificDataItem {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CommandSpecificDataItemBuilder) DeepCopy() any {
	_copy := b.CreateCommandSpecificDataItemBuilder().(*_CommandSpecificDataItemBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_CommandSpecificDataItemChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCommandSpecificDataItemBuilder creates a CommandSpecificDataItemBuilder
func (b *_CommandSpecificDataItem) CreateCommandSpecificDataItemBuilder() CommandSpecificDataItemBuilder {
	if b == nil {
		return NewCommandSpecificDataItemBuilder()
	}
	return &_CommandSpecificDataItemBuilder{_CommandSpecificDataItem: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCommandSpecificDataItem(structType any) CommandSpecificDataItem {
	if casted, ok := structType.(CommandSpecificDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*CommandSpecificDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_CommandSpecificDataItem) GetTypeName() string {
	return "CommandSpecificDataItem"
}

func (m *_CommandSpecificDataItem) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (itemType)
	lengthInBits += 16

	return lengthInBits
}

func (m *_CommandSpecificDataItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CommandSpecificDataItemParse[T CommandSpecificDataItem](ctx context.Context, theBytes []byte) (T, error) {
	return CommandSpecificDataItemParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func CommandSpecificDataItemParseWithBufferProducer[T CommandSpecificDataItem]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CommandSpecificDataItemParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func CommandSpecificDataItemParseWithBuffer[T CommandSpecificDataItem](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_CommandSpecificDataItem{}).parse(ctx, readBuffer)
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

func (m *_CommandSpecificDataItem) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__commandSpecificDataItem CommandSpecificDataItem, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CommandSpecificDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CommandSpecificDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemType, err := ReadDiscriminatorField[uint16](ctx, "itemType", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CommandSpecificDataItem
	switch {
	case itemType == 0x000C: // CipIdentity
		if _child, err = new(_CipIdentity).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipIdentity for type-switch of CommandSpecificDataItem")
		}
	case itemType == 0x0086: // CipSecurityInformation
		if _child, err = new(_CipSecurityInformation).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CipSecurityInformation for type-switch of CommandSpecificDataItem")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [itemType=%v]", itemType)
	}

	if closeErr := readBuffer.CloseContext("CommandSpecificDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CommandSpecificDataItem")
	}

	return _child, nil
}

func (pm *_CommandSpecificDataItem) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CommandSpecificDataItem, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CommandSpecificDataItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CommandSpecificDataItem")
	}

	if err := WriteDiscriminatorField(ctx, "itemType", m.GetItemType(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'itemType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CommandSpecificDataItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CommandSpecificDataItem")
	}
	return nil
}

func (m *_CommandSpecificDataItem) IsCommandSpecificDataItem() {}

func (m *_CommandSpecificDataItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CommandSpecificDataItem) deepCopy() *_CommandSpecificDataItem {
	if m == nil {
		return nil
	}
	_CommandSpecificDataItemCopy := &_CommandSpecificDataItem{
		nil, // will be set by child
	}
	return _CommandSpecificDataItemCopy
}

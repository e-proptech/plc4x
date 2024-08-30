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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetNotificationParametersChangeOfReliability extends BACnetNotificationParameters
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag innerOpeningTag;
  protected final BACnetReliabilityTagged reliability;
  protected final BACnetStatusFlagsTagged statusFlags;
  protected final BACnetPropertyValues propertyValues;
  protected final BACnetClosingTag innerClosingTag;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetObjectType objectTypeArgument;

  public BACnetNotificationParametersChangeOfReliability(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetOpeningTag innerOpeningTag,
      BACnetReliabilityTagged reliability,
      BACnetStatusFlagsTagged statusFlags,
      BACnetPropertyValues propertyValues,
      BACnetClosingTag innerClosingTag,
      Short tagNumber,
      BACnetObjectType objectTypeArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument);
    this.innerOpeningTag = innerOpeningTag;
    this.reliability = reliability;
    this.statusFlags = statusFlags;
    this.propertyValues = propertyValues;
    this.innerClosingTag = innerClosingTag;
    this.tagNumber = tagNumber;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetOpeningTag getInnerOpeningTag() {
    return innerOpeningTag;
  }

  public BACnetReliabilityTagged getReliability() {
    return reliability;
  }

  public BACnetStatusFlagsTagged getStatusFlags() {
    return statusFlags;
  }

  public BACnetPropertyValues getPropertyValues() {
    return propertyValues;
  }

  public BACnetClosingTag getInnerClosingTag() {
    return innerClosingTag;
  }

  @Override
  protected void serializeBACnetNotificationParametersChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetNotificationParametersChangeOfReliability");

    // Simple Field (innerOpeningTag)
    writeSimpleField("innerOpeningTag", innerOpeningTag, writeComplex(writeBuffer));

    // Simple Field (reliability)
    writeSimpleField("reliability", reliability, writeComplex(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, writeComplex(writeBuffer));

    // Simple Field (propertyValues)
    writeSimpleField("propertyValues", propertyValues, writeComplex(writeBuffer));

    // Simple Field (innerClosingTag)
    writeSimpleField("innerClosingTag", innerClosingTag, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersChangeOfReliability");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersChangeOfReliability _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (innerOpeningTag)
    lengthInBits += innerOpeningTag.getLengthInBits();

    // Simple field (reliability)
    lengthInBits += reliability.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (propertyValues)
    lengthInBits += propertyValues.getLengthInBits();

    // Simple field (innerClosingTag)
    lengthInBits += innerClosingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersBuilder staticParseBACnetNotificationParametersBuilder(
      ReadBuffer readBuffer,
      Short peekedTagNumber,
      Short tagNumber,
      BACnetObjectType objectTypeArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersChangeOfReliability");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag innerOpeningTag =
        readSimpleField(
            "innerOpeningTag",
            readComplex(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    BACnetReliabilityTagged reliability =
        readSimpleField(
            "reliability",
            readComplex(
                () ->
                    BACnetReliabilityTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetStatusFlagsTagged statusFlags =
        readSimpleField(
            "statusFlags",
            readComplex(
                () ->
                    BACnetStatusFlagsTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetPropertyValues propertyValues =
        readSimpleField(
            "propertyValues",
            readComplex(
                () ->
                    BACnetPropertyValues.staticParse(
                        readBuffer, (short) (2), (BACnetObjectType) (objectTypeArgument)),
                readBuffer));

    BACnetClosingTag innerClosingTag =
        readSimpleField(
            "innerClosingTag",
            readComplex(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (peekedTagNumber)),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersChangeOfReliability");
    // Create the instance
    return new BACnetNotificationParametersChangeOfReliabilityBuilderImpl(
        innerOpeningTag,
        reliability,
        statusFlags,
        propertyValues,
        innerClosingTag,
        tagNumber,
        objectTypeArgument);
  }

  public static class BACnetNotificationParametersChangeOfReliabilityBuilderImpl
      implements BACnetNotificationParameters.BACnetNotificationParametersBuilder {
    private final BACnetOpeningTag innerOpeningTag;
    private final BACnetReliabilityTagged reliability;
    private final BACnetStatusFlagsTagged statusFlags;
    private final BACnetPropertyValues propertyValues;
    private final BACnetClosingTag innerClosingTag;
    private final Short tagNumber;
    private final BACnetObjectType objectTypeArgument;

    public BACnetNotificationParametersChangeOfReliabilityBuilderImpl(
        BACnetOpeningTag innerOpeningTag,
        BACnetReliabilityTagged reliability,
        BACnetStatusFlagsTagged statusFlags,
        BACnetPropertyValues propertyValues,
        BACnetClosingTag innerClosingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      this.innerOpeningTag = innerOpeningTag;
      this.reliability = reliability;
      this.statusFlags = statusFlags;
      this.propertyValues = propertyValues;
      this.innerClosingTag = innerClosingTag;
      this.tagNumber = tagNumber;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetNotificationParametersChangeOfReliability build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetObjectType objectTypeArgument) {
      BACnetNotificationParametersChangeOfReliability
          bACnetNotificationParametersChangeOfReliability =
              new BACnetNotificationParametersChangeOfReliability(
                  openingTag,
                  peekedTagHeader,
                  closingTag,
                  innerOpeningTag,
                  reliability,
                  statusFlags,
                  propertyValues,
                  innerClosingTag,
                  tagNumber,
                  objectTypeArgument);
      return bACnetNotificationParametersChangeOfReliability;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersChangeOfReliability)) {
      return false;
    }
    BACnetNotificationParametersChangeOfReliability that =
        (BACnetNotificationParametersChangeOfReliability) o;
    return (getInnerOpeningTag() == that.getInnerOpeningTag())
        && (getReliability() == that.getReliability())
        && (getStatusFlags() == that.getStatusFlags())
        && (getPropertyValues() == that.getPropertyValues())
        && (getInnerClosingTag() == that.getInnerClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getInnerOpeningTag(),
        getReliability(),
        getStatusFlags(),
        getPropertyValues(),
        getInnerClosingTag());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}

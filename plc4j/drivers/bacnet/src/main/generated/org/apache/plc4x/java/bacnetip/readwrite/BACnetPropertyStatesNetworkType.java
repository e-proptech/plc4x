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

public class BACnetPropertyStatesNetworkType extends BACnetPropertyStates implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetNetworkTypeTagged networkType;

  public BACnetPropertyStatesNetworkType(
      BACnetTagHeader peekedTagHeader, BACnetNetworkTypeTagged networkType) {
    super(peekedTagHeader);
    this.networkType = networkType;
  }

  public BACnetNetworkTypeTagged getNetworkType() {
    return networkType;
  }

  @Override
  protected void serializeBACnetPropertyStatesChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetPropertyStatesNetworkType");

    // Simple Field (networkType)
    writeSimpleField("networkType", networkType, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetPropertyStatesNetworkType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPropertyStatesNetworkType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (networkType)
    lengthInBits += networkType.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyStatesBuilder staticParseBACnetPropertyStatesBuilder(
      ReadBuffer readBuffer, Short peekedTagNumber) throws ParseException {
    readBuffer.pullContext("BACnetPropertyStatesNetworkType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetNetworkTypeTagged networkType =
        readSimpleField(
            "networkType",
            readComplex(
                () ->
                    BACnetNetworkTypeTagged.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyStatesNetworkType");
    // Create the instance
    return new BACnetPropertyStatesNetworkTypeBuilderImpl(networkType);
  }

  public static class BACnetPropertyStatesNetworkTypeBuilderImpl
      implements BACnetPropertyStates.BACnetPropertyStatesBuilder {
    private final BACnetNetworkTypeTagged networkType;

    public BACnetPropertyStatesNetworkTypeBuilderImpl(BACnetNetworkTypeTagged networkType) {
      this.networkType = networkType;
    }

    public BACnetPropertyStatesNetworkType build(BACnetTagHeader peekedTagHeader) {
      BACnetPropertyStatesNetworkType bACnetPropertyStatesNetworkType =
          new BACnetPropertyStatesNetworkType(peekedTagHeader, networkType);
      return bACnetPropertyStatesNetworkType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyStatesNetworkType)) {
      return false;
    }
    BACnetPropertyStatesNetworkType that = (BACnetPropertyStatesNetworkType) o;
    return (getNetworkType() == that.getNetworkType()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNetworkType());
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

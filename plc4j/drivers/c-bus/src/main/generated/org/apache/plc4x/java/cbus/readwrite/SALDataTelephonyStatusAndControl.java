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
package org.apache.plc4x.java.cbus.readwrite;

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

public class SALDataTelephonyStatusAndControl extends SALData implements Message {

  // Accessors for discriminator values.
  public ApplicationId getApplicationId() {
    return ApplicationId.TELEPHONY_STATUS_AND_CONTROL;
  }

  // Properties.
  protected final TelephonyData telephonyData;

  public SALDataTelephonyStatusAndControl(SALData salData, TelephonyData telephonyData) {
    super(salData);
    this.telephonyData = telephonyData;
  }

  public TelephonyData getTelephonyData() {
    return telephonyData;
  }

  @Override
  protected void serializeSALDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SALDataTelephonyStatusAndControl");

    // Simple Field (telephonyData)
    writeSimpleField("telephonyData", telephonyData, writeComplex(writeBuffer));

    writeBuffer.popContext("SALDataTelephonyStatusAndControl");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SALDataTelephonyStatusAndControl _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (telephonyData)
    lengthInBits += telephonyData.getLengthInBits();

    return lengthInBits;
  }

  public static SALDataBuilder staticParseSALDataBuilder(
      ReadBuffer readBuffer, ApplicationId applicationId) throws ParseException {
    readBuffer.pullContext("SALDataTelephonyStatusAndControl");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    TelephonyData telephonyData =
        readSimpleField(
            "telephonyData", readComplex(() -> TelephonyData.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("SALDataTelephonyStatusAndControl");
    // Create the instance
    return new SALDataTelephonyStatusAndControlBuilderImpl(telephonyData);
  }

  public static class SALDataTelephonyStatusAndControlBuilderImpl
      implements SALData.SALDataBuilder {
    private final TelephonyData telephonyData;

    public SALDataTelephonyStatusAndControlBuilderImpl(TelephonyData telephonyData) {
      this.telephonyData = telephonyData;
    }

    public SALDataTelephonyStatusAndControl build(SALData salData) {
      SALDataTelephonyStatusAndControl sALDataTelephonyStatusAndControl =
          new SALDataTelephonyStatusAndControl(salData, telephonyData);
      return sALDataTelephonyStatusAndControl;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SALDataTelephonyStatusAndControl)) {
      return false;
    }
    SALDataTelephonyStatusAndControl that = (SALDataTelephonyStatusAndControl) o;
    return (getTelephonyData() == that.getTelephonyData()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getTelephonyData());
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

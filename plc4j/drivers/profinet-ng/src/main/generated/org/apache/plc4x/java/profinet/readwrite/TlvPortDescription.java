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
package org.apache.plc4x.java.profinet.readwrite;

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

public class TlvPortDescription extends LldpUnit implements Message {

  // Accessors for discriminator values.
  public TlvType getTlvId() {
    return TlvType.PORT_DESCRIPTION;
  }

  // Properties.
  protected final String chassisId;

  public TlvPortDescription(short tlvIdLength, String chassisId) {
    super(tlvIdLength);
    this.chassisId = chassisId;
  }

  public String getChassisId() {
    return chassisId;
  }

  @Override
  protected void serializeLldpUnitChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TlvPortDescription");

    // Simple Field (chassisId)
    writeSimpleField("chassisId", chassisId, writeString(writeBuffer, ((tlvIdLength)) * (8)));

    writeBuffer.popContext("TlvPortDescription");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TlvPortDescription _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (chassisId)
    lengthInBits += ((tlvIdLength)) * (8);

    return lengthInBits;
  }

  public static LldpUnitBuilder staticParseLldpUnitBuilder(ReadBuffer readBuffer, Short tlvIdLength)
      throws ParseException {
    readBuffer.pullContext("TlvPortDescription");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    String chassisId = readSimpleField("chassisId", readString(readBuffer, ((tlvIdLength)) * (8)));

    readBuffer.closeContext("TlvPortDescription");
    // Create the instance
    return new TlvPortDescriptionBuilderImpl(chassisId);
  }

  public static class TlvPortDescriptionBuilderImpl implements LldpUnit.LldpUnitBuilder {
    private final String chassisId;

    public TlvPortDescriptionBuilderImpl(String chassisId) {
      this.chassisId = chassisId;
    }

    public TlvPortDescription build(short tlvIdLength) {
      TlvPortDescription tlvPortDescription = new TlvPortDescription(tlvIdLength, chassisId);
      return tlvPortDescription;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TlvPortDescription)) {
      return false;
    }
    TlvPortDescription that = (TlvPortDescription) o;
    return (getChassisId() == that.getChassisId()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getChassisId());
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

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

public class TlvSystemCapabilities extends LldpUnit implements Message {

  // Accessors for discriminator values.
  public TlvType getTlvId() {
    return TlvType.SYSTEM_CAPABILITIES;
  }

  // Properties.
  protected final boolean stationOnlyCapable;
  protected final boolean docsisCableDeviceCapable;
  protected final boolean telephoneCapable;
  protected final boolean routerCapable;
  protected final boolean wlanAccessPointCapable;
  protected final boolean bridgeCapable;
  protected final boolean repeaterCapable;
  protected final boolean otherCapable;
  protected final boolean stationOnlyEnabled;
  protected final boolean docsisCableDeviceEnabled;
  protected final boolean telephoneEnabled;
  protected final boolean routerEnabled;
  protected final boolean wlanAccessPointEnabled;
  protected final boolean bridgeEnabled;
  protected final boolean repeaterEnabled;
  protected final boolean otherEnabled;

  // Reserved Fields
  private Short reservedField0;
  private Short reservedField1;

  public TlvSystemCapabilities(
      short tlvIdLength,
      boolean stationOnlyCapable,
      boolean docsisCableDeviceCapable,
      boolean telephoneCapable,
      boolean routerCapable,
      boolean wlanAccessPointCapable,
      boolean bridgeCapable,
      boolean repeaterCapable,
      boolean otherCapable,
      boolean stationOnlyEnabled,
      boolean docsisCableDeviceEnabled,
      boolean telephoneEnabled,
      boolean routerEnabled,
      boolean wlanAccessPointEnabled,
      boolean bridgeEnabled,
      boolean repeaterEnabled,
      boolean otherEnabled) {
    super(tlvIdLength);
    this.stationOnlyCapable = stationOnlyCapable;
    this.docsisCableDeviceCapable = docsisCableDeviceCapable;
    this.telephoneCapable = telephoneCapable;
    this.routerCapable = routerCapable;
    this.wlanAccessPointCapable = wlanAccessPointCapable;
    this.bridgeCapable = bridgeCapable;
    this.repeaterCapable = repeaterCapable;
    this.otherCapable = otherCapable;
    this.stationOnlyEnabled = stationOnlyEnabled;
    this.docsisCableDeviceEnabled = docsisCableDeviceEnabled;
    this.telephoneEnabled = telephoneEnabled;
    this.routerEnabled = routerEnabled;
    this.wlanAccessPointEnabled = wlanAccessPointEnabled;
    this.bridgeEnabled = bridgeEnabled;
    this.repeaterEnabled = repeaterEnabled;
    this.otherEnabled = otherEnabled;
  }

  public boolean getStationOnlyCapable() {
    return stationOnlyCapable;
  }

  public boolean getDocsisCableDeviceCapable() {
    return docsisCableDeviceCapable;
  }

  public boolean getTelephoneCapable() {
    return telephoneCapable;
  }

  public boolean getRouterCapable() {
    return routerCapable;
  }

  public boolean getWlanAccessPointCapable() {
    return wlanAccessPointCapable;
  }

  public boolean getBridgeCapable() {
    return bridgeCapable;
  }

  public boolean getRepeaterCapable() {
    return repeaterCapable;
  }

  public boolean getOtherCapable() {
    return otherCapable;
  }

  public boolean getStationOnlyEnabled() {
    return stationOnlyEnabled;
  }

  public boolean getDocsisCableDeviceEnabled() {
    return docsisCableDeviceEnabled;
  }

  public boolean getTelephoneEnabled() {
    return telephoneEnabled;
  }

  public boolean getRouterEnabled() {
    return routerEnabled;
  }

  public boolean getWlanAccessPointEnabled() {
    return wlanAccessPointEnabled;
  }

  public boolean getBridgeEnabled() {
    return bridgeEnabled;
  }

  public boolean getRepeaterEnabled() {
    return repeaterEnabled;
  }

  public boolean getOtherEnabled() {
    return otherEnabled;
  }

  @Override
  protected void serializeLldpUnitChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TlvSystemCapabilities");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 8));

    // Simple Field (stationOnlyCapable)
    writeSimpleField("stationOnlyCapable", stationOnlyCapable, writeBoolean(writeBuffer));

    // Simple Field (docsisCableDeviceCapable)
    writeSimpleField(
        "docsisCableDeviceCapable", docsisCableDeviceCapable, writeBoolean(writeBuffer));

    // Simple Field (telephoneCapable)
    writeSimpleField("telephoneCapable", telephoneCapable, writeBoolean(writeBuffer));

    // Simple Field (routerCapable)
    writeSimpleField("routerCapable", routerCapable, writeBoolean(writeBuffer));

    // Simple Field (wlanAccessPointCapable)
    writeSimpleField("wlanAccessPointCapable", wlanAccessPointCapable, writeBoolean(writeBuffer));

    // Simple Field (bridgeCapable)
    writeSimpleField("bridgeCapable", bridgeCapable, writeBoolean(writeBuffer));

    // Simple Field (repeaterCapable)
    writeSimpleField("repeaterCapable", repeaterCapable, writeBoolean(writeBuffer));

    // Simple Field (otherCapable)
    writeSimpleField("otherCapable", otherCapable, writeBoolean(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField1 != null ? reservedField1 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 8));

    // Simple Field (stationOnlyEnabled)
    writeSimpleField("stationOnlyEnabled", stationOnlyEnabled, writeBoolean(writeBuffer));

    // Simple Field (docsisCableDeviceEnabled)
    writeSimpleField(
        "docsisCableDeviceEnabled", docsisCableDeviceEnabled, writeBoolean(writeBuffer));

    // Simple Field (telephoneEnabled)
    writeSimpleField("telephoneEnabled", telephoneEnabled, writeBoolean(writeBuffer));

    // Simple Field (routerEnabled)
    writeSimpleField("routerEnabled", routerEnabled, writeBoolean(writeBuffer));

    // Simple Field (wlanAccessPointEnabled)
    writeSimpleField("wlanAccessPointEnabled", wlanAccessPointEnabled, writeBoolean(writeBuffer));

    // Simple Field (bridgeEnabled)
    writeSimpleField("bridgeEnabled", bridgeEnabled, writeBoolean(writeBuffer));

    // Simple Field (repeaterEnabled)
    writeSimpleField("repeaterEnabled", repeaterEnabled, writeBoolean(writeBuffer));

    // Simple Field (otherEnabled)
    writeSimpleField("otherEnabled", otherEnabled, writeBoolean(writeBuffer));

    writeBuffer.popContext("TlvSystemCapabilities");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TlvSystemCapabilities _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (stationOnlyCapable)
    lengthInBits += 1;

    // Simple field (docsisCableDeviceCapable)
    lengthInBits += 1;

    // Simple field (telephoneCapable)
    lengthInBits += 1;

    // Simple field (routerCapable)
    lengthInBits += 1;

    // Simple field (wlanAccessPointCapable)
    lengthInBits += 1;

    // Simple field (bridgeCapable)
    lengthInBits += 1;

    // Simple field (repeaterCapable)
    lengthInBits += 1;

    // Simple field (otherCapable)
    lengthInBits += 1;

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (stationOnlyEnabled)
    lengthInBits += 1;

    // Simple field (docsisCableDeviceEnabled)
    lengthInBits += 1;

    // Simple field (telephoneEnabled)
    lengthInBits += 1;

    // Simple field (routerEnabled)
    lengthInBits += 1;

    // Simple field (wlanAccessPointEnabled)
    lengthInBits += 1;

    // Simple field (bridgeEnabled)
    lengthInBits += 1;

    // Simple field (repeaterEnabled)
    lengthInBits += 1;

    // Simple field (otherEnabled)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static LldpUnitBuilder staticParseLldpUnitBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("TlvSystemCapabilities");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    boolean stationOnlyCapable = readSimpleField("stationOnlyCapable", readBoolean(readBuffer));

    boolean docsisCableDeviceCapable =
        readSimpleField("docsisCableDeviceCapable", readBoolean(readBuffer));

    boolean telephoneCapable = readSimpleField("telephoneCapable", readBoolean(readBuffer));

    boolean routerCapable = readSimpleField("routerCapable", readBoolean(readBuffer));

    boolean wlanAccessPointCapable =
        readSimpleField("wlanAccessPointCapable", readBoolean(readBuffer));

    boolean bridgeCapable = readSimpleField("bridgeCapable", readBoolean(readBuffer));

    boolean repeaterCapable = readSimpleField("repeaterCapable", readBoolean(readBuffer));

    boolean otherCapable = readSimpleField("otherCapable", readBoolean(readBuffer));

    Short reservedField1 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    boolean stationOnlyEnabled = readSimpleField("stationOnlyEnabled", readBoolean(readBuffer));

    boolean docsisCableDeviceEnabled =
        readSimpleField("docsisCableDeviceEnabled", readBoolean(readBuffer));

    boolean telephoneEnabled = readSimpleField("telephoneEnabled", readBoolean(readBuffer));

    boolean routerEnabled = readSimpleField("routerEnabled", readBoolean(readBuffer));

    boolean wlanAccessPointEnabled =
        readSimpleField("wlanAccessPointEnabled", readBoolean(readBuffer));

    boolean bridgeEnabled = readSimpleField("bridgeEnabled", readBoolean(readBuffer));

    boolean repeaterEnabled = readSimpleField("repeaterEnabled", readBoolean(readBuffer));

    boolean otherEnabled = readSimpleField("otherEnabled", readBoolean(readBuffer));

    readBuffer.closeContext("TlvSystemCapabilities");
    // Create the instance
    return new TlvSystemCapabilitiesBuilderImpl(
        stationOnlyCapable,
        docsisCableDeviceCapable,
        telephoneCapable,
        routerCapable,
        wlanAccessPointCapable,
        bridgeCapable,
        repeaterCapable,
        otherCapable,
        stationOnlyEnabled,
        docsisCableDeviceEnabled,
        telephoneEnabled,
        routerEnabled,
        wlanAccessPointEnabled,
        bridgeEnabled,
        repeaterEnabled,
        otherEnabled,
        reservedField0,
        reservedField1);
  }

  public static class TlvSystemCapabilitiesBuilderImpl implements LldpUnit.LldpUnitBuilder {
    private final boolean stationOnlyCapable;
    private final boolean docsisCableDeviceCapable;
    private final boolean telephoneCapable;
    private final boolean routerCapable;
    private final boolean wlanAccessPointCapable;
    private final boolean bridgeCapable;
    private final boolean repeaterCapable;
    private final boolean otherCapable;
    private final boolean stationOnlyEnabled;
    private final boolean docsisCableDeviceEnabled;
    private final boolean telephoneEnabled;
    private final boolean routerEnabled;
    private final boolean wlanAccessPointEnabled;
    private final boolean bridgeEnabled;
    private final boolean repeaterEnabled;
    private final boolean otherEnabled;
    private final Short reservedField0;
    private final Short reservedField1;

    public TlvSystemCapabilitiesBuilderImpl(
        boolean stationOnlyCapable,
        boolean docsisCableDeviceCapable,
        boolean telephoneCapable,
        boolean routerCapable,
        boolean wlanAccessPointCapable,
        boolean bridgeCapable,
        boolean repeaterCapable,
        boolean otherCapable,
        boolean stationOnlyEnabled,
        boolean docsisCableDeviceEnabled,
        boolean telephoneEnabled,
        boolean routerEnabled,
        boolean wlanAccessPointEnabled,
        boolean bridgeEnabled,
        boolean repeaterEnabled,
        boolean otherEnabled,
        Short reservedField0,
        Short reservedField1) {
      this.stationOnlyCapable = stationOnlyCapable;
      this.docsisCableDeviceCapable = docsisCableDeviceCapable;
      this.telephoneCapable = telephoneCapable;
      this.routerCapable = routerCapable;
      this.wlanAccessPointCapable = wlanAccessPointCapable;
      this.bridgeCapable = bridgeCapable;
      this.repeaterCapable = repeaterCapable;
      this.otherCapable = otherCapable;
      this.stationOnlyEnabled = stationOnlyEnabled;
      this.docsisCableDeviceEnabled = docsisCableDeviceEnabled;
      this.telephoneEnabled = telephoneEnabled;
      this.routerEnabled = routerEnabled;
      this.wlanAccessPointEnabled = wlanAccessPointEnabled;
      this.bridgeEnabled = bridgeEnabled;
      this.repeaterEnabled = repeaterEnabled;
      this.otherEnabled = otherEnabled;
      this.reservedField0 = reservedField0;
      this.reservedField1 = reservedField1;
    }

    public TlvSystemCapabilities build(short tlvIdLength) {
      TlvSystemCapabilities tlvSystemCapabilities =
          new TlvSystemCapabilities(
              tlvIdLength,
              stationOnlyCapable,
              docsisCableDeviceCapable,
              telephoneCapable,
              routerCapable,
              wlanAccessPointCapable,
              bridgeCapable,
              repeaterCapable,
              otherCapable,
              stationOnlyEnabled,
              docsisCableDeviceEnabled,
              telephoneEnabled,
              routerEnabled,
              wlanAccessPointEnabled,
              bridgeEnabled,
              repeaterEnabled,
              otherEnabled);
      tlvSystemCapabilities.reservedField0 = reservedField0;
      tlvSystemCapabilities.reservedField1 = reservedField1;
      return tlvSystemCapabilities;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TlvSystemCapabilities)) {
      return false;
    }
    TlvSystemCapabilities that = (TlvSystemCapabilities) o;
    return (getStationOnlyCapable() == that.getStationOnlyCapable())
        && (getDocsisCableDeviceCapable() == that.getDocsisCableDeviceCapable())
        && (getTelephoneCapable() == that.getTelephoneCapable())
        && (getRouterCapable() == that.getRouterCapable())
        && (getWlanAccessPointCapable() == that.getWlanAccessPointCapable())
        && (getBridgeCapable() == that.getBridgeCapable())
        && (getRepeaterCapable() == that.getRepeaterCapable())
        && (getOtherCapable() == that.getOtherCapable())
        && (getStationOnlyEnabled() == that.getStationOnlyEnabled())
        && (getDocsisCableDeviceEnabled() == that.getDocsisCableDeviceEnabled())
        && (getTelephoneEnabled() == that.getTelephoneEnabled())
        && (getRouterEnabled() == that.getRouterEnabled())
        && (getWlanAccessPointEnabled() == that.getWlanAccessPointEnabled())
        && (getBridgeEnabled() == that.getBridgeEnabled())
        && (getRepeaterEnabled() == that.getRepeaterEnabled())
        && (getOtherEnabled() == that.getOtherEnabled())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getStationOnlyCapable(),
        getDocsisCableDeviceCapable(),
        getTelephoneCapable(),
        getRouterCapable(),
        getWlanAccessPointCapable(),
        getBridgeCapable(),
        getRepeaterCapable(),
        getOtherCapable(),
        getStationOnlyEnabled(),
        getDocsisCableDeviceEnabled(),
        getTelephoneEnabled(),
        getRouterEnabled(),
        getWlanAccessPointEnabled(),
        getBridgeEnabled(),
        getRepeaterEnabled(),
        getOtherEnabled());
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

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
package org.apache.plc4x.java.opcua.readwrite;

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

public class PortableQualifiedName extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 24107;
  }

  // Properties.
  protected final PascalString namespaceUri;
  protected final PascalString name;

  public PortableQualifiedName(PascalString namespaceUri, PascalString name) {
    super();
    this.namespaceUri = namespaceUri;
    this.name = name;
  }

  public PascalString getNamespaceUri() {
    return namespaceUri;
  }

  public PascalString getName() {
    return name;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PortableQualifiedName");

    // Simple Field (namespaceUri)
    writeSimpleField("namespaceUri", namespaceUri, writeComplex(writeBuffer));

    // Simple Field (name)
    writeSimpleField("name", name, writeComplex(writeBuffer));

    writeBuffer.popContext("PortableQualifiedName");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PortableQualifiedName _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (namespaceUri)
    lengthInBits += namespaceUri.getLengthInBits();

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("PortableQualifiedName");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString namespaceUri =
        readSimpleField(
            "namespaceUri", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString name =
        readSimpleField(
            "name", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("PortableQualifiedName");
    // Create the instance
    return new PortableQualifiedNameBuilderImpl(namespaceUri, name);
  }

  public static class PortableQualifiedNameBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString namespaceUri;
    private final PascalString name;

    public PortableQualifiedNameBuilderImpl(PascalString namespaceUri, PascalString name) {
      this.namespaceUri = namespaceUri;
      this.name = name;
    }

    public PortableQualifiedName build() {
      PortableQualifiedName portableQualifiedName = new PortableQualifiedName(namespaceUri, name);
      return portableQualifiedName;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PortableQualifiedName)) {
      return false;
    }
    PortableQualifiedName that = (PortableQualifiedName) o;
    return (getNamespaceUri() == that.getNamespaceUri())
        && (getName() == that.getName())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNamespaceUri(), getName());
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

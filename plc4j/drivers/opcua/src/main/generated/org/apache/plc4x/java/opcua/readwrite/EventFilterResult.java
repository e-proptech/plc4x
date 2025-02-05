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

public class EventFilterResult extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 736;
  }

  // Properties.
  protected final List<StatusCode> selectClauseResults;
  protected final List<DiagnosticInfo> selectClauseDiagnosticInfos;
  protected final ContentFilterResult whereClauseResult;

  public EventFilterResult(
      List<StatusCode> selectClauseResults,
      List<DiagnosticInfo> selectClauseDiagnosticInfos,
      ContentFilterResult whereClauseResult) {
    super();
    this.selectClauseResults = selectClauseResults;
    this.selectClauseDiagnosticInfos = selectClauseDiagnosticInfos;
    this.whereClauseResult = whereClauseResult;
  }

  public List<StatusCode> getSelectClauseResults() {
    return selectClauseResults;
  }

  public List<DiagnosticInfo> getSelectClauseDiagnosticInfos() {
    return selectClauseDiagnosticInfos;
  }

  public ContentFilterResult getWhereClauseResult() {
    return whereClauseResult;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("EventFilterResult");

    // Implicit Field (noOfSelectClauseResults) (Used for parsing, but its value is not stored as
    // it's implicitly given by the objects content)
    int noOfSelectClauseResults =
        (int) ((((getSelectClauseResults()) == (null)) ? -(1) : COUNT(getSelectClauseResults())));
    writeImplicitField(
        "noOfSelectClauseResults", noOfSelectClauseResults, writeSignedInt(writeBuffer, 32));

    // Array Field (selectClauseResults)
    writeComplexTypeArrayField("selectClauseResults", selectClauseResults, writeBuffer);

    // Implicit Field (noOfSelectClauseDiagnosticInfos) (Used for parsing, but its value is not
    // stored as it's implicitly given by the objects content)
    int noOfSelectClauseDiagnosticInfos =
        (int)
            ((((getSelectClauseDiagnosticInfos()) == (null))
                ? -(1)
                : COUNT(getSelectClauseDiagnosticInfos())));
    writeImplicitField(
        "noOfSelectClauseDiagnosticInfos",
        noOfSelectClauseDiagnosticInfos,
        writeSignedInt(writeBuffer, 32));

    // Array Field (selectClauseDiagnosticInfos)
    writeComplexTypeArrayField(
        "selectClauseDiagnosticInfos", selectClauseDiagnosticInfos, writeBuffer);

    // Simple Field (whereClauseResult)
    writeSimpleField("whereClauseResult", whereClauseResult, writeComplex(writeBuffer));

    writeBuffer.popContext("EventFilterResult");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    EventFilterResult _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (noOfSelectClauseResults)
    lengthInBits += 32;

    // Array field
    if (selectClauseResults != null) {
      int i = 0;
      for (StatusCode element : selectClauseResults) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= selectClauseResults.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfSelectClauseDiagnosticInfos)
    lengthInBits += 32;

    // Array field
    if (selectClauseDiagnosticInfos != null) {
      int i = 0;
      for (DiagnosticInfo element : selectClauseDiagnosticInfos) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= selectClauseDiagnosticInfos.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (whereClauseResult)
    lengthInBits += whereClauseResult.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("EventFilterResult");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfSelectClauseResults =
        readImplicitField("noOfSelectClauseResults", readSignedInt(readBuffer, 32));

    List<StatusCode> selectClauseResults =
        readCountArrayField(
            "selectClauseResults",
            readComplex(() -> StatusCode.staticParse(readBuffer), readBuffer),
            noOfSelectClauseResults);

    int noOfSelectClauseDiagnosticInfos =
        readImplicitField("noOfSelectClauseDiagnosticInfos", readSignedInt(readBuffer, 32));

    List<DiagnosticInfo> selectClauseDiagnosticInfos =
        readCountArrayField(
            "selectClauseDiagnosticInfos",
            readComplex(() -> DiagnosticInfo.staticParse(readBuffer), readBuffer),
            noOfSelectClauseDiagnosticInfos);

    ContentFilterResult whereClauseResult =
        readSimpleField(
            "whereClauseResult",
            readComplex(
                () ->
                    (ContentFilterResult)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (609)),
                readBuffer));

    readBuffer.closeContext("EventFilterResult");
    // Create the instance
    return new EventFilterResultBuilderImpl(
        selectClauseResults, selectClauseDiagnosticInfos, whereClauseResult);
  }

  public static class EventFilterResultBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final List<StatusCode> selectClauseResults;
    private final List<DiagnosticInfo> selectClauseDiagnosticInfos;
    private final ContentFilterResult whereClauseResult;

    public EventFilterResultBuilderImpl(
        List<StatusCode> selectClauseResults,
        List<DiagnosticInfo> selectClauseDiagnosticInfos,
        ContentFilterResult whereClauseResult) {
      this.selectClauseResults = selectClauseResults;
      this.selectClauseDiagnosticInfos = selectClauseDiagnosticInfos;
      this.whereClauseResult = whereClauseResult;
    }

    public EventFilterResult build() {
      EventFilterResult eventFilterResult =
          new EventFilterResult(
              selectClauseResults, selectClauseDiagnosticInfos, whereClauseResult);
      return eventFilterResult;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EventFilterResult)) {
      return false;
    }
    EventFilterResult that = (EventFilterResult) o;
    return (getSelectClauseResults() == that.getSelectClauseResults())
        && (getSelectClauseDiagnosticInfos() == that.getSelectClauseDiagnosticInfos())
        && (getWhereClauseResult() == that.getWhereClauseResult())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSelectClauseResults(),
        getSelectClauseDiagnosticInfos(),
        getWhereClauseResult());
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

/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 */

def checkBison() {
    print "bash git --version".execute().text
}

def checkDotnet() {
    print "bash dotnet --version".execute().text
}

def checkFlex() {
    print "bash flex --version".execute().text
}

def checkGcc() {
    print "bash gcc --version".execute().text
}

def checkGit() {
    print "bash git --version".execute().text
}

def checkGpp() {
    print "bash g++ --version".execute().text
}

def checkPython() {
    print "bash python --version".execute().text
}

// Additional Checks we should add:
// - Windows:
//     - Check the length of the path of the base dir as we're having issues with the length of paths being too long.

checkBison()

checkDotnet()

checkFlex()

checkGcc()

checkGit()

checkGpp()

checkPython()
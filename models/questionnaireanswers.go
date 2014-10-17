// Copyright (c) 2011-2014, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import "time"

type QuestionnaireAnswers struct {
	Id            string         `json:"-" bson:"_id"`
	Identifier    Identifier     `bson:"identifier"`
	Questionnaire Reference      `bson:"questionnaire"`
	Status        string         `bson:"status"`
	Subject       Reference      `bson:"subject"`
	Author        Reference      `bson:"author"`
	Authored      time.Time      `bson:"authored"`
	Source        Reference      `bson:"source"`
	Encounter     Reference      `bson:"encounter"`
	Group         GroupComponent `bson:"group"`
}

// This is an ugly hack to deal with embedded structures in the spec answer
type QuestionAnswerComponent struct {
	ValueBoolean    bool       `bson:"valueBoolean"`
	ValueDecimal    float64    `bson:"valueDecimal"`
	ValueInteger    float64    `bson:"valueInteger"`
	ValueDate       time.Time  `bson:"valueDate"`
	ValueDateTime   time.Time  `bson:"valueDateTime"`
	ValueInstant    time.Time  `bson:"valueInstant"`
	ValueTime       time.Time  `bson:"valueTime"`
	ValueString     string     `bson:"valueString"`
	ValueAttachment Attachment `bson:"valueAttachment"`
	ValueCoding     Coding     `bson:"valueCoding"`
	ValueQuantity   Quantity   `bson:"valueQuantity"`
	ValueReference  Reference  `bson:"valueReference"`
}

// This is an ugly hack to deal with embedded structures in the spec question
type QuestionComponent struct {
	LinkId string                    `bson:"linkId"`
	Text   string                    `bson:"text"`
	Answer []QuestionAnswerComponent `bson:"answer"`
	Group  []GroupComponent          `bson:"group"`
}

// This is an ugly hack to deal with embedded structures in the spec group
type GroupComponent struct {
	LinkId   string              `bson:"linkId"`
	Title    string              `bson:"title"`
	Text     string              `bson:"text"`
	Subject  Reference           `bson:"subject"`
	Group    []GroupComponent    `bson:"group"`
	Question []QuestionComponent `bson:"question"`
}
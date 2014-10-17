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

type DiagnosticReport struct {
	Id                 string                           `json:"-" bson:"_id"`
	Name               CodeableConcept                  `bson:"name"`
	Status             string                           `bson:"status"`
	Issued             time.Time                        `bson:"issued"`
	Subject            Reference                        `bson:"subject"`
	Performer          Reference                        `bson:"performer"`
	Identifier         Identifier                       `bson:"identifier"`
	RequestDetail      []Reference                      `bson:"requestDetail"`
	ServiceCategory    CodeableConcept                  `bson:"serviceCategory"`
	DiagnosticDateTime time.Time                        `bson:"diagnosticDateTime"`
	DiagnosticPeriod   Period                           `bson:"diagnosticPeriod"`
	Specimen           []Reference                      `bson:"specimen"`
	Result             []Reference                      `bson:"result"`
	ImagingStudy       []Reference                      `bson:"imagingStudy"`
	Image              []DiagnosticReportImageComponent `bson:"image"`
	Conclusion         string                           `bson:"conclusion"`
	CodedDiagnosis     []CodeableConcept                `bson:"codedDiagnosis"`
	PresentedForm      []Attachment                     `bson:"presentedForm"`
}

// This is an ugly hack to deal with embedded structures in the spec image
type DiagnosticReportImageComponent struct {
	Comment string    `bson:"comment"`
	Link    Reference `bson:"link"`
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FalconxSandboxReportV1 falconx sandbox report v1
//
// swagger:model falconx.SandboxReportV1
type FalconxSandboxReportV1 struct {

	// architecture
	Architecture string `json:"architecture,omitempty"`

	// certificates
	Certificates []*FalconxCertificate `json:"certificates"`

	// certificates validation message
	CertificatesValidationMessage string `json:"certificates_validation_message,omitempty"`

	// classification
	Classification []string `json:"classification"`

	// classification tags
	ClassificationTags []string `json:"classification_tags"`

	// contacted hosts
	ContactedHosts []*FalconxContactedHost `json:"contacted_hosts"`

	// dll characteristics
	DllCharacteristics []string `json:"dll_characteristics"`

	// dns requests
	DNSRequests []*FalconxDNSRequest `json:"dns_requests"`

	// entrypoint
	Entrypoint string `json:"entrypoint,omitempty"`

	// entrypoint preview count
	EntrypointPreviewCount int32 `json:"entrypoint_preview_count,omitempty"`

	// entrypoint preview instructions
	EntrypointPreviewInstructions []string `json:"entrypoint_preview_instructions"`

	// entrypoint section
	EntrypointSection string `json:"entrypoint_section,omitempty"`

	// environment description
	EnvironmentDescription string `json:"environment_description,omitempty"`

	// environment id
	EnvironmentID int32 `json:"environment_id,omitempty"`

	// error message
	ErrorMessage string `json:"error_message,omitempty"`

	// error origin
	ErrorOrigin string `json:"error_origin,omitempty"`

	// error type
	ErrorType string `json:"error_type,omitempty"`

	// exact deep hash
	ExactDeepHash string `json:"exact_deep_hash,omitempty"`

	// extracted files
	ExtractedFiles []*FalconxExtractedFile `json:"extracted_files"`

	// extracted interesting strings
	ExtractedInterestingStrings []*FalconxExtractedInterestingString `json:"extracted_interesting_strings"`

	// file data directories
	FileDataDirectories []*FalconxFileDataDirectory `json:"file_data_directories"`

	// file imports
	FileImports []*FalconxFileImport `json:"file_imports"`

	// file metadata
	FileMetadata *FalconxFileMetadata `json:"file_metadata,omitempty"`

	// file resources
	FileResources []*FalconxFileResource `json:"file_resources"`

	// file sections
	FileSections []*FalconxFileSection `json:"file_sections"`

	// file size
	FileSize int32 `json:"file_size,omitempty"`

	// file type
	FileType string `json:"file_type,omitempty"`

	// file type short
	FileTypeShort []string `json:"file_type_short"`

	// http requests
	HTTPRequests []*FalconxHTTPRequest `json:"http_requests"`

	// icon
	Icon string `json:"icon,omitempty"`

	// image base
	ImageBase string `json:"image_base,omitempty"`

	// image file characteristics
	ImageFileCharacteristics []string `json:"image_file_characteristics"`

	// incidents
	Incidents []*FalconxIncident `json:"incidents"`

	// intelligence mitre attacks
	IntelligenceMitreAttacks []*FalconxMITREAttack `json:"intelligence_mitre_attacks"`

	// ioc report broad artifact id
	IocReportBroadArtifactID string `json:"ioc_report_broad_artifact_id,omitempty"`

	// ioc report strict artifact id
	IocReportStrictArtifactID string `json:"ioc_report_strict_artifact_id,omitempty"`

	// is certificates valid
	// Required: true
	IsCertificatesValid *bool `json:"is_certificates_valid"`

	// language
	Language string `json:"language,omitempty"`

	// major os version
	MajorOsVersion int32 `json:"major_os_version,omitempty"`

	// memory dumps
	MemoryDumps []*FalconxMemoryDumpData `json:"memory_dumps"`

	// memory dumps artifact id
	MemoryDumpsArtifactID string `json:"memory_dumps_artifact_id,omitempty"`

	// memory forensics
	MemoryForensics []*FalconxMemoryForensic `json:"memory_forensics"`

	// memory strings artifact id
	MemoryStringsArtifactID string `json:"memory_strings_artifact_id,omitempty"`

	// minor os version
	MinorOsVersion int32 `json:"minor_os_version,omitempty"`

	// mitre attacks
	MitreAttacks []*FalconxMITREAttack `json:"mitre_attacks"`

	// network settings
	NetworkSettings string `json:"network_settings,omitempty"`

	// packer
	Packer string `json:"packer,omitempty"`

	// pcap report artifact id
	PcapReportArtifactID string `json:"pcap_report_artifact_id,omitempty"`

	// processes
	Processes []*FalconxProcess `json:"processes"`

	// sample flags
	SampleFlags []string `json:"sample_flags"`

	// screenshots artifact ids
	ScreenshotsArtifactIds []string `json:"screenshots_artifact_ids"`

	// sha256
	Sha256 string `json:"sha256,omitempty"`

	// signatures
	Signatures []*FalconxSignature `json:"signatures"`

	// submission type
	SubmissionType string `json:"submission_type,omitempty"`

	// submit name
	SubmitName string `json:"submit_name,omitempty"`

	// submit url
	SubmitURL string `json:"submit_url,omitempty"`

	// subsystem
	Subsystem string `json:"subsystem,omitempty"`

	// suricata alerts
	SuricataAlerts []*FalconxSuricataAlert `json:"suricata_alerts"`

	// target url
	TargetURL string `json:"target_url,omitempty"`

	// threat score
	ThreatScore int32 `json:"threat_score,omitempty"`

	// urls
	Urls []*FalconxURLData `json:"urls"`

	// verdict
	Verdict string `json:"verdict,omitempty"`

	// version info
	VersionInfo []*FalconxVersionInfo `json:"version_info"`

	// visualization
	Visualization string `json:"visualization,omitempty"`

	// windows version bitness
	WindowsVersionBitness int32 `json:"windows_version_bitness,omitempty"`

	// windows version edition
	WindowsVersionEdition string `json:"windows_version_edition,omitempty"`

	// windows version name
	WindowsVersionName string `json:"windows_version_name,omitempty"`

	// windows version service pack
	WindowsVersionServicePack string `json:"windows_version_service_pack,omitempty"`

	// windows version version
	WindowsVersionVersion string `json:"windows_version_version,omitempty"`
}

// Validate validates this falconx sandbox report v1
func (m *FalconxSandboxReportV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactedHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtractedFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtractedInterestingStrings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileDataDirectories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileImports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileSections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncidents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntelligenceMitreAttacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsCertificatesValid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryDumps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryForensics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMitreAttacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcesses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuricataAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUrls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxSandboxReportV1) validateCertificates(formats strfmt.Registry) error {
	if swag.IsZero(m.Certificates) { // not required
		return nil
	}

	for i := 0; i < len(m.Certificates); i++ {
		if swag.IsZero(m.Certificates[i]) { // not required
			continue
		}

		if m.Certificates[i] != nil {
			if err := m.Certificates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("certificates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateContactedHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.ContactedHosts) { // not required
		return nil
	}

	for i := 0; i < len(m.ContactedHosts); i++ {
		if swag.IsZero(m.ContactedHosts[i]) { // not required
			continue
		}

		if m.ContactedHosts[i] != nil {
			if err := m.ContactedHosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contacted_hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contacted_hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateDNSRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSRequests); i++ {
		if swag.IsZero(m.DNSRequests[i]) { // not required
			continue
		}

		if m.DNSRequests[i] != nil {
			if err := m.DNSRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns_requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dns_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateExtractedFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtractedFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtractedFiles); i++ {
		if swag.IsZero(m.ExtractedFiles[i]) { // not required
			continue
		}

		if m.ExtractedFiles[i] != nil {
			if err := m.ExtractedFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extracted_files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extracted_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateExtractedInterestingStrings(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtractedInterestingStrings) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtractedInterestingStrings); i++ {
		if swag.IsZero(m.ExtractedInterestingStrings[i]) { // not required
			continue
		}

		if m.ExtractedInterestingStrings[i] != nil {
			if err := m.ExtractedInterestingStrings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extracted_interesting_strings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extracted_interesting_strings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateFileDataDirectories(formats strfmt.Registry) error {
	if swag.IsZero(m.FileDataDirectories) { // not required
		return nil
	}

	for i := 0; i < len(m.FileDataDirectories); i++ {
		if swag.IsZero(m.FileDataDirectories[i]) { // not required
			continue
		}

		if m.FileDataDirectories[i] != nil {
			if err := m.FileDataDirectories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("file_data_directories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("file_data_directories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateFileImports(formats strfmt.Registry) error {
	if swag.IsZero(m.FileImports) { // not required
		return nil
	}

	for i := 0; i < len(m.FileImports); i++ {
		if swag.IsZero(m.FileImports[i]) { // not required
			continue
		}

		if m.FileImports[i] != nil {
			if err := m.FileImports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("file_imports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("file_imports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateFileMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.FileMetadata) { // not required
		return nil
	}

	if m.FileMetadata != nil {
		if err := m.FileMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file_metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("file_metadata")
			}
			return err
		}
	}

	return nil
}

func (m *FalconxSandboxReportV1) validateFileResources(formats strfmt.Registry) error {
	if swag.IsZero(m.FileResources) { // not required
		return nil
	}

	for i := 0; i < len(m.FileResources); i++ {
		if swag.IsZero(m.FileResources[i]) { // not required
			continue
		}

		if m.FileResources[i] != nil {
			if err := m.FileResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("file_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("file_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateFileSections(formats strfmt.Registry) error {
	if swag.IsZero(m.FileSections) { // not required
		return nil
	}

	for i := 0; i < len(m.FileSections); i++ {
		if swag.IsZero(m.FileSections[i]) { // not required
			continue
		}

		if m.FileSections[i] != nil {
			if err := m.FileSections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("file_sections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("file_sections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateHTTPRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.HTTPRequests); i++ {
		if swag.IsZero(m.HTTPRequests[i]) { // not required
			continue
		}

		if m.HTTPRequests[i] != nil {
			if err := m.HTTPRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("http_requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("http_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateIncidents(formats strfmt.Registry) error {
	if swag.IsZero(m.Incidents) { // not required
		return nil
	}

	for i := 0; i < len(m.Incidents); i++ {
		if swag.IsZero(m.Incidents[i]) { // not required
			continue
		}

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateIntelligenceMitreAttacks(formats strfmt.Registry) error {
	if swag.IsZero(m.IntelligenceMitreAttacks) { // not required
		return nil
	}

	for i := 0; i < len(m.IntelligenceMitreAttacks); i++ {
		if swag.IsZero(m.IntelligenceMitreAttacks[i]) { // not required
			continue
		}

		if m.IntelligenceMitreAttacks[i] != nil {
			if err := m.IntelligenceMitreAttacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intelligence_mitre_attacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intelligence_mitre_attacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateIsCertificatesValid(formats strfmt.Registry) error {

	if err := validate.Required("is_certificates_valid", "body", m.IsCertificatesValid); err != nil {
		return err
	}

	return nil
}

func (m *FalconxSandboxReportV1) validateMemoryDumps(formats strfmt.Registry) error {
	if swag.IsZero(m.MemoryDumps) { // not required
		return nil
	}

	for i := 0; i < len(m.MemoryDumps); i++ {
		if swag.IsZero(m.MemoryDumps[i]) { // not required
			continue
		}

		if m.MemoryDumps[i] != nil {
			if err := m.MemoryDumps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memory_dumps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("memory_dumps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateMemoryForensics(formats strfmt.Registry) error {
	if swag.IsZero(m.MemoryForensics) { // not required
		return nil
	}

	for i := 0; i < len(m.MemoryForensics); i++ {
		if swag.IsZero(m.MemoryForensics[i]) { // not required
			continue
		}

		if m.MemoryForensics[i] != nil {
			if err := m.MemoryForensics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memory_forensics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("memory_forensics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateMitreAttacks(formats strfmt.Registry) error {
	if swag.IsZero(m.MitreAttacks) { // not required
		return nil
	}

	for i := 0; i < len(m.MitreAttacks); i++ {
		if swag.IsZero(m.MitreAttacks[i]) { // not required
			continue
		}

		if m.MitreAttacks[i] != nil {
			if err := m.MitreAttacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mitre_attacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mitre_attacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateProcesses(formats strfmt.Registry) error {
	if swag.IsZero(m.Processes) { // not required
		return nil
	}

	for i := 0; i < len(m.Processes); i++ {
		if swag.IsZero(m.Processes[i]) { // not required
			continue
		}

		if m.Processes[i] != nil {
			if err := m.Processes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("processes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateSignatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Signatures) { // not required
		return nil
	}

	for i := 0; i < len(m.Signatures); i++ {
		if swag.IsZero(m.Signatures[i]) { // not required
			continue
		}

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signatures" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateSuricataAlerts(formats strfmt.Registry) error {
	if swag.IsZero(m.SuricataAlerts) { // not required
		return nil
	}

	for i := 0; i < len(m.SuricataAlerts); i++ {
		if swag.IsZero(m.SuricataAlerts[i]) { // not required
			continue
		}

		if m.SuricataAlerts[i] != nil {
			if err := m.SuricataAlerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suricata_alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("suricata_alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateUrls(formats strfmt.Registry) error {
	if swag.IsZero(m.Urls) { // not required
		return nil
	}

	for i := 0; i < len(m.Urls); i++ {
		if swag.IsZero(m.Urls[i]) { // not required
			continue
		}

		if m.Urls[i] != nil {
			if err := m.Urls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("urls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("urls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) validateVersionInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.VersionInfo); i++ {
		if swag.IsZero(m.VersionInfo[i]) { // not required
			continue
		}

		if m.VersionInfo[i] != nil {
			if err := m.VersionInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("version_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("version_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this falconx sandbox report v1 based on the context it is used
func (m *FalconxSandboxReportV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContactedHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNSRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtractedFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtractedInterestingStrings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileDataDirectories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileImports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileSections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncidents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIntelligenceMitreAttacks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemoryDumps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemoryForensics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMitreAttacks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcesses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuricataAlerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUrls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FalconxSandboxReportV1) contextValidateCertificates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Certificates); i++ {

		if m.Certificates[i] != nil {

			if swag.IsZero(m.Certificates[i]) { // not required
				return nil
			}

			if err := m.Certificates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("certificates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateContactedHosts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContactedHosts); i++ {

		if m.ContactedHosts[i] != nil {

			if swag.IsZero(m.ContactedHosts[i]) { // not required
				return nil
			}

			if err := m.ContactedHosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contacted_hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contacted_hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateDNSRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSRequests); i++ {

		if m.DNSRequests[i] != nil {

			if swag.IsZero(m.DNSRequests[i]) { // not required
				return nil
			}

			if err := m.DNSRequests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns_requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dns_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateExtractedFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExtractedFiles); i++ {

		if m.ExtractedFiles[i] != nil {

			if swag.IsZero(m.ExtractedFiles[i]) { // not required
				return nil
			}

			if err := m.ExtractedFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extracted_files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extracted_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateExtractedInterestingStrings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExtractedInterestingStrings); i++ {

		if m.ExtractedInterestingStrings[i] != nil {

			if swag.IsZero(m.ExtractedInterestingStrings[i]) { // not required
				return nil
			}

			if err := m.ExtractedInterestingStrings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extracted_interesting_strings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extracted_interesting_strings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateFileDataDirectories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileDataDirectories); i++ {

		if m.FileDataDirectories[i] != nil {

			if swag.IsZero(m.FileDataDirectories[i]) { // not required
				return nil
			}

			if err := m.FileDataDirectories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("file_data_directories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("file_data_directories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateFileImports(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileImports); i++ {

		if m.FileImports[i] != nil {

			if swag.IsZero(m.FileImports[i]) { // not required
				return nil
			}

			if err := m.FileImports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("file_imports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("file_imports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateFileMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.FileMetadata != nil {

		if swag.IsZero(m.FileMetadata) { // not required
			return nil
		}

		if err := m.FileMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file_metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("file_metadata")
			}
			return err
		}
	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateFileResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileResources); i++ {

		if m.FileResources[i] != nil {

			if swag.IsZero(m.FileResources[i]) { // not required
				return nil
			}

			if err := m.FileResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("file_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("file_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateFileSections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileSections); i++ {

		if m.FileSections[i] != nil {

			if swag.IsZero(m.FileSections[i]) { // not required
				return nil
			}

			if err := m.FileSections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("file_sections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("file_sections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateHTTPRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HTTPRequests); i++ {

		if m.HTTPRequests[i] != nil {

			if swag.IsZero(m.HTTPRequests[i]) { // not required
				return nil
			}

			if err := m.HTTPRequests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("http_requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("http_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateIncidents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Incidents); i++ {

		if m.Incidents[i] != nil {

			if swag.IsZero(m.Incidents[i]) { // not required
				return nil
			}

			if err := m.Incidents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateIntelligenceMitreAttacks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IntelligenceMitreAttacks); i++ {

		if m.IntelligenceMitreAttacks[i] != nil {

			if swag.IsZero(m.IntelligenceMitreAttacks[i]) { // not required
				return nil
			}

			if err := m.IntelligenceMitreAttacks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intelligence_mitre_attacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intelligence_mitre_attacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateMemoryDumps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MemoryDumps); i++ {

		if m.MemoryDumps[i] != nil {

			if swag.IsZero(m.MemoryDumps[i]) { // not required
				return nil
			}

			if err := m.MemoryDumps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memory_dumps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("memory_dumps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateMemoryForensics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MemoryForensics); i++ {

		if m.MemoryForensics[i] != nil {

			if swag.IsZero(m.MemoryForensics[i]) { // not required
				return nil
			}

			if err := m.MemoryForensics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memory_forensics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("memory_forensics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateMitreAttacks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MitreAttacks); i++ {

		if m.MitreAttacks[i] != nil {

			if swag.IsZero(m.MitreAttacks[i]) { // not required
				return nil
			}

			if err := m.MitreAttacks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mitre_attacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mitre_attacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateProcesses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Processes); i++ {

		if m.Processes[i] != nil {

			if swag.IsZero(m.Processes[i]) { // not required
				return nil
			}

			if err := m.Processes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("processes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateSignatures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Signatures); i++ {

		if m.Signatures[i] != nil {

			if swag.IsZero(m.Signatures[i]) { // not required
				return nil
			}

			if err := m.Signatures[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signatures" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateSuricataAlerts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SuricataAlerts); i++ {

		if m.SuricataAlerts[i] != nil {

			if swag.IsZero(m.SuricataAlerts[i]) { // not required
				return nil
			}

			if err := m.SuricataAlerts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suricata_alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("suricata_alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateUrls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Urls); i++ {

		if m.Urls[i] != nil {

			if swag.IsZero(m.Urls[i]) { // not required
				return nil
			}

			if err := m.Urls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("urls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("urls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FalconxSandboxReportV1) contextValidateVersionInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VersionInfo); i++ {

		if m.VersionInfo[i] != nil {

			if swag.IsZero(m.VersionInfo[i]) { // not required
				return nil
			}

			if err := m.VersionInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("version_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("version_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FalconxSandboxReportV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxSandboxReportV1) UnmarshalBinary(b []byte) error {
	var res FalconxSandboxReportV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

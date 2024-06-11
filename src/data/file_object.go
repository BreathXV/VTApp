package data

var (
	EngineNameType  = map[string]string{}
	SandboxNameType = map[string]string{}
	RulesetNameType = map[string]string{}
	ValueType       = map[string]string{}
)

// TODO: Fix JSON keys that are undefined w/h maps
type FileReportStructure struct {
	Data struct {
		Attributes struct {
			CapabilitiesTags       []string `json:"capabilities_tags"`
			CreationDate           int      `json:"creation_date"`
			CrowdsourcedIdsResults []struct {
				AlertContext []struct {
					Proto   string `json:"proto"`
					SrcIP   string `json:"src_ip"`
					SrcPort int    `json:"src_port"`
				} `json:"alert_context"`
				AlertSeverity string `json:"alert_severity"`
				RuleCategory  string `json:"rule_category"`
				RuleID        string `json:"rule_id"`
				RuleMsg       string `json:"rule_msg"`
				RuleSource    string `json:"rule_source"`
			} `json:"crowdsourced_ids_results"`
			CrowdsourcedIdsStats struct {
				High   int `json:"high"`
				Info   int `json:"info"`
				Low    int `json:"low"`
				Medium int `json:"medium"`
			} `json:"crowdsourced_ids_stats"`
			CrowdsourcedYaraResults []struct {
				Description    string `json:"description"`
				MatchInSubfile bool   `json:"match_in_subfile"`
				RuleName       string `json:"rule_name"`
				RulesetID      string `json:"ruleset_id"`
				RulesetName    string `json:"ruleset_name"`
				Source         string `json:"source"`
			} `json:"crowdsourced_yara_results"`
			Downloadable        bool `json:"downloadable"`
			FirstSubmissionDate int  `json:"first_submission_date"`
			LastAnalysisDate    int  `json:"last_analysis_date"`
<<<<<<< HEAD
			LastAnalysisResults []struct {
				EngineNameType struct {
=======
			LastAnalysisResults map[struct]interface {
				EngineNameType []string {
>>>>>>> 98a400bed8698fbf2bf13de095957382c9265f5a
					Category      string `json:"category"`
					EngineName    string `json:"engine_name"`
					EngineUpdate  string `json:"engine_update"`
					EngineVersion string `json:"engine_version"`
					Method        string `json:"method"`
					Result        string `json:"result"`
<<<<<<< HEAD
				} `json:""` // Unknown key name
			} `json:""` // Unknown key name
			LastAnalysisStats    struct{} `json:"last_analysis_stats"`
=======
				} `json:unknown_field`
			} `json:"last_analysis_results"`
			LastAnalysisStats struct {
				ConfirmedTimeout int `json:"confirmed-timeout"`
				Failure          int `json:"failure"`
				Harmless         int `json:"harmless"`
				Malicious        int `json:"malicious"`
				Suspicious       int `json:"suspicious"`
				Timeout          int `json:"timeout"`
				TypeUnsupported  int `json:"type-unsupported"`
				Undetected       int `json:"undetected"`
			} `json:"last_analysis_stats"`
>>>>>>> 98a400bed8698fbf2bf13de095957382c9265f5a
			LastModificationDate int      `json:"last_modification_date"`
			LastSubmissionDate   int      `json:"last_submission_date"`
			Md5                  string   `json:"md5"`
			MeaningfulName       string   `json:"meaningful_name"`
			Names                []string `json:"names"`
			Reputation           int      `json:"reputation"`
			SandboxVerdicts      struct {
				VirusTotalJujubox struct {
					Category              string   `json:"category"`
					Confidence            int      `json:"confidence"`
					MalwareClassification []string `json:"malware_classification"`
					MalwareNames          []string `json:"malware_names"`
					SandboxName           string   `json:"sandbox_name"`
				} `json:"VirusTotal Jujubox"`
			} `json:"sandbox_verdicts"`
			Sha1                 string `json:"sha1"`
			Sha256               string `json:"sha256"`
			SigmaAnalysisSummary struct {
				SigmaIntegratedRuleSetGitHub struct {
					High     int `json:"high"`
					Medium   int `json:"medium"`
					Critical int `json:"critical"`
					Low      int `json:"low"`
				} `json:"Sigma Integrated Rule Set (GitHub)"`
				SOCPrimeThreatDetectionMarketplace struct {
					High     int `json:"high"`
					Medium   int `json:"medium"`
					Critical int `json:"critical"`
					Low      int `json:"low"`
				} `json:"SOC Prime Threat Detection Marketplace"`
			} `json:"sigma_analysis_summary"`
			SigmaAnalysisStats struct {
				High     int `json:"high"`
				Medium   int `json:"medium"`
				Critical int `json:"critical"`
				Low      int `json:"low"`
			} `json:"sigma_analysis_stats"`
			SigmaAnalysisResults []struct {
				RuleTitle    string `json:"rule_title"`
				RuleSource   string `json:"rule_source"`
				MatchContext []struct {
					Values struct {
						TerminalSessionID string `json:"TerminalSessionId"`
						ProcessGUID       string `json:"ProcessGuid"`
						ProcessID         string `json:"ProcessId"`
						Product           string `json:"Product"`
						Description       string `json:"Description"`
						Company           string `json:"Company"`
						ParentProcessGUID string `json:"ParentProcessGuid"`
						User              string `json:"User"`
						Hashes            string `json:"Hashes"`
						OriginalFileName  string `json:"OriginalFileName"`
						ParentImage       string `json:"ParentImage"`
						FileVersion       string `json:"FileVersion"`
						ParentProcessID   string `json:"ParentProcessId"`
						CurrentDirectory  string `json:"CurrentDirectory"`
						CommandLine       string `json:"CommandLine"`
						EventID           string `json:"EventID"`
						LogonGUID         string `json:"LogonGuid"`
						LogonID           string `json:"LogonId"`
						Image             string `json:"Image"`
						IntegrityLevel    string `json:"IntegrityLevel"`
						ParentCommandLine string `json:"ParentCommandLine"`
						UtcTime           string `json:"UtcTime"`
						RuleName          string `json:"RuleName"`
					} `json:"values"`
				} `json:"match_context"`
				RuleLevel       string `json:"rule_level"`
				RuleDescription string `json:"rule_description"`
				RuleAuthor      string `json:"rule_author"`
				RuleID          string `json:"rule_id"`
			} `json:"sigma_analysis_results"`
			Size           int      `json:"size"`
			Tags           []string `json:"tags"`
			TimesSubmitted int      `json:"times_submitted"`
			TotalVotes     struct {
				Harmless  int `json:"harmless"`
				Malicious int `json:"malicious"`
			} `json:"total_votes"`
			TypeDescription string `json:"type_description"`
			TypeTag         string `json:"type_tag"`
			UniqueSources   int    `json:"unique_sources"`
			Vhash           string `json:"vhash"`
		} `json:"attributes"`
		ID    string `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Type string `json:"type"`
	} `json:"data"`
}

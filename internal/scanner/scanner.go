package scanner

type Resource struct {
	Name   string `json:"name"`
	Public bool   `json:"public"`
}

type Finding struct {
	Resource string `json:"resource"`
	Rule     string `json:"rule"`
	Severity string `json:"severity"`
}

func Scan(resources []Resource) []Finding {
	var findings []Finding
	for _, r := range resources {
		if r.Public {
			findings = append(findings, Finding{Resource: r.Name, Rule: "no-public-resource", Severity: "high"})
		}
	}
	return findings
}

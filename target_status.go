package main

// map const status Operational, Degraded Performance, Partial Outage, Major Outage, Under Maintenance
type TargetStatus int

const (
	Operational TargetStatus = iota
	DegradedPerformance
	PartialOutage
	MajorOutage
	UnderMaintenance
)

func (s TargetStatus) String() string {
	return [...]string{"operational", "degraded_performance", "partial_outage", "major_outage", "under_maintenance"}[s]
}

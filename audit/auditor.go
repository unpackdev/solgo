package audit

import (
	"context"

	"github.com/txpull/solgo"
)

// Auditor represents a structure that manages the auditing process
// of smart contracts using the Slither tool.
type Auditor struct {
	ctx     context.Context // Context for the auditor operations.
	config  *Config         // Configuration for the Slither tool.
	sources *solgo.Sources  // Sources of the smart contracts to be audited.
	slither *Slither        // Instance of the Slither tool.
}

// NewAuditor initializes a new Auditor instance with the provided context,
// configuration, and sources. It ensures that the Slither tool is properly
// initialized and that the sources are prepared for analysis.
func NewAuditor(ctx context.Context, config *Config, sources *solgo.Sources) (*Auditor, error) {
	slither, err := NewSlither(ctx, config)
	if err != nil {
		return nil, err
	}

	// Ensure that the sources are prepared for future consumption.
	if !sources.ArePrepared() {
		if err := sources.Prepare(); err != nil {
			return nil, err
		}
	}

	return &Auditor{
		ctx:     ctx,
		config:  config,
		sources: sources,
		slither: slither,
	}, nil
}

// IsReady checks if the Auditor is ready to perform an analysis.
// It ensures that the Slither tool is installed and that the sources are prepared.
func (a *Auditor) IsReady() bool {
	return a.slither.IsInstalled() && a.sources.ArePrepared()
}

// GetConfig returns the configuration used by the Auditor.
func (a *Auditor) GetConfig() *Config {
	return a.config
}

// GetSources returns the smart contract sources managed by the Auditor.
func (a *Auditor) GetSources() *solgo.Sources {
	return a.sources
}

// GetSlither returns the instance of the Slither tool used by the Auditor.
func (a *Auditor) GetSlither() *Slither {
	return a.slither
}

// Analyze performs an analysis of the smart contracts using the Slither tool.
// It returns the analysis response or an error if the analysis fails.
func (a *Auditor) Analyze() (*Report, error) {
	response, _, err := a.slither.Analyze(a.sources)
	if err != nil {
		return nil, err
	}

	return response, nil
}

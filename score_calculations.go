package gosecurityscorecard

type ScoreCalculator struct {
  scorecard *SecurityScorecard
}

func NewScoreCalculator(scorecard *SecurityScorecard) *ScoreCalculator {
  return &ScoreCalculator{
    scorecard: scorecard,
  }
}

func (c *ScoreCalculator) CalculateSSLScore() {
  c.scorecard.SSLConfiguration = 0 // calculate score
}

func (c *ScoreCalculator) CalculateAuthScore() {
  c.scorecard.Authentication = 0 // calculate score 
}

func (c *ScoreCalculator) CalculateDBEncryptionScore() {
  c.scorecard.DatabaseEncryption = 0 // calculate score
}

func (c *ScoreCalculator) CalculateOverallScore() {
  // calculate overall score
  c.scorecard.OverallScore = 0 
}
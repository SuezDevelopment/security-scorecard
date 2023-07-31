package gosecurityscorecard

type SecurityAssessmentLogic struct {
  scorecard *SecurityScorecard
}

func NewSecurityAssessmentLogic(scorecard *SecurityScorecard) *SecurityAssessmentLogic {
  return &SecurityAssessmentLogic{
    scorecard: scorecard,
  }
}


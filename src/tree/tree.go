package tree

type Question struct {
	Question string `json:"question"`  // question text
	YesScore int    `json:"yes_score"` // score for yes answer
	NoScore  int    `json:"no_score"`  // score for no answer

	YesLink *Question `json:"yes"` // next question for yes answer
	NoLink  *Question `json:"no"`  // next question for no answer
}

type DecisionTree struct {
	Root *Question `json:"root"` // root question

	AllQuestions []*Question `json:"all_questions"` // all questions in the tree
}

package agent


type MainOrchestratorService struct {
	AgentInp *AgentServiceInput
	AgentOut *AgentServiceOutput
	database  map[int]*UserAnswer
}


func NewMainOrchestratorService(agentInp *AgentServiceInput, agentOut *AgentServiceOutput) *MainOrchestratorService {
	return &MainOrchestratorService{
		AgentInp: agentInp,
		AgentOut: agentOut,
		database: make(map[int]*UserAnswer),
	}
}

func (s *MainOrchestratorService) GetAnswerData(id int) *UserAnswer { 
	if vl, ok := s.database[id]; ok {
		delete(s.database, id)
		return vl
	} else {
		return nil
	}
}

package services

import (
	"project/application/core/entities"
	"project/application/core/ports"
	"strconv"
)

type GPTFillService struct {
	userRepo      ports.UserRepository
	challengeRepo ports.ChallengeRepository
	videoRepo     ports.VideoRepository
}

func NewGPTFillService(userRepo ports.UserRepository, challengeRepo ports.ChallengeRepository, videoRepo ports.VideoRepository) *GPTFillService {
	return &GPTFillService{userRepo, challengeRepo, videoRepo}
}

func (s *GPTFillService) FillTables() error {
	// Llenar tabla de usuarios
	for i := 0; i < 30; i++ {
		user := &entities.User{
			Name:  "User " + strconv.Itoa(i),
			Email: "user" + strconv.Itoa(i) + "@example.com",
		}
		if err := s.userRepo.CreateUser(user); err != nil {
			return err
		}
	}

	// Llenar tabla de desafíos
	for i := 0; i < 30; i++ {
		// prompt := "Generate a detailed description for a challenge titled 'Challenge " + strconv.Itoa(i) + "' with difficulty level " + strconv.Itoa(i%5+1)
		//description, err := pkg.GenerateDescription(prompt)
		//if err != nil {
		//	fmt.Printf(err.Error())
		//}
		challenge := &entities.Challenge{
			Title: "Challenge " + strconv.Itoa(i),
			//Description: description,
			Description: "Description " + strconv.Itoa(i),
			Difficulty:  i%5 + 1,
			UserID:      i%10 + 1,
		}
		if err := s.challengeRepo.CreateChallenge(challenge); err != nil {
			return err
		}
	}

	// Llenar tabla de videos
	for i := 0; i < 30; i++ {
		//prompt := "Generate a detailed description for a video titled 'Video " + strconv.Itoa(i) + "'"
		//description, err := pkg.GenerateDescription(prompt)
		//if err != nil {
		//	return err
		//}
		video := &entities.Video{
			Title: "Video " + strconv.Itoa(i),
			// Description: description,
			Description: "Description " + strconv.Itoa(i),
			URL:         "http://example.com/video" + strconv.Itoa(i),
			UserID:      i%10 + 1,
		}
		if err := s.videoRepo.CreateVideo(video); err != nil {
			return err
		}
	}

	return nil
}

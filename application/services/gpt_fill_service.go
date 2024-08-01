package services

import (
	"fmt"
	"project/application/core/entities"
	"project/application/core/ports"
	"project/pkg"
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
	if err := s.fillUserTable(); err != nil {
		return err
	}

	if err := s.fillChallengeTable(); err != nil {
		return err
	}

	if err := s.fillVideoTable(); err != nil {
		return err
	}
	return nil
}

// Llenar tabla de usuarios
func (s *GPTFillService) fillUserTable() error {
	for i := 0; i < 30; i++ {
		user := &entities.User{
			Name:  "User " + strconv.Itoa(i),
			Email: "user" + strconv.Itoa(i) + "@example.com",
		}
		if err := s.userRepo.CreateUser(user); err != nil {
			return err
		}
	}
	return nil
}

// Llenar tabla de desafíos
func (s *GPTFillService) fillChallengeTable() error {
	for i := 0; i < 30; i++ {
		title, err := pkg.GenerateDescription("Generate a short title regarding this topic: art challenges to discover talents. This title will fill the 'Title' column in a Challenge table")
		if err != nil {
			fmt.Printf(err.Error())
			title = "Random title"
		}
		description, err := pkg.GenerateDescription("Generate a description based on this challenge title '" + title + "'")
		if err != nil {
			fmt.Printf(err.Error())
		}
		challenge := &entities.Challenge{
			Title:       title,
			Description: description,
			Difficulty:  i%5 + 1,
			UserID:      i%10 + 1,
		}
		if err := s.challengeRepo.CreateChallenge(challenge); err != nil {
			return err
		}
	}
	return nil
}

// Llenar tabla de videos
func (s *GPTFillService) fillVideoTable() error {
	for i := 0; i < 30; i++ {
		title, err := pkg.GenerateDescription("Generate a random short title regarding this topic: naming a video to demo a talent. This title will fill the 'Title' column in a Video table")
		if err != nil {
			fmt.Printf(err.Error())
			title = "Random title"
		}
		description, err := pkg.GenerateDescription("Generate a description based on this video title '" + title + "'")
		if err != nil {
			fmt.Printf(err.Error())
		}
		video := &entities.Video{
			Title:       title,
			Description: description,
			URL:         "http://example.com/video" + strconv.Itoa(i),
			UserID:      i%10 + 1,
		}
		if err := s.videoRepo.CreateVideo(video); err != nil {
			return err
		}
	}

	return nil
}

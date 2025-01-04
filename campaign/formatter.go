package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatted := CampaignFormatter{}

	formatted.ID = campaign.ID
	formatted.UserID = campaign.UserID
	formatted.Name = campaign.Name
	formatted.ShortDescription = campaign.ShortDescription

	if len(campaign.CampaignImages) > 0 {
		formatted.ImageURL = campaign.CampaignImages[0].FileName
	}

	formatted.GoalAmount = campaign.GoalAmount
	formatted.CurrentAmount = campaign.CurrentAmount
	formatted.Slug = campaign.Slug

	return formatted
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	formatted := []CampaignFormatter{}

	for _, campaign := range campaigns {
		formatted = append(formatted, FormatCampaign(campaign))
	}

	return formatted
}

type CampaignDetailFormatter struct {
	ID               int                   `json:"id"`
	UserID           int                   `json:"user_id"`
	Name             string                `json:"name"`
	ShortDescription string                `json:"short_description"`
	Description      string                `json:"description"`
	Perks            []string              `json:"perks"`
	ImageURL         string                `json:"image_url"`
	GoalAmount       int                   `json:"goal_amount"`
	CurrentAmount    int                   `json:"current_amount"`
	Slug             string                `json:"slug"`
	User             CampaignUserFormatter `json:"user"`
	CampaignImages   []CampaignImagesFormatter
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	formatted := CampaignDetailFormatter{}

	formatted.ID = campaign.ID
	formatted.UserID = campaign.UserID
	formatted.Name = campaign.Name
	formatted.ShortDescription = campaign.ShortDescription
	formatted.Description = campaign.Description

	perks := []string{}
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	formatted.Perks = perks

	if len(campaign.CampaignImages) > 0 {
		formatted.ImageURL = campaign.CampaignImages[0].FileName
	}

	formatted.GoalAmount = campaign.GoalAmount
	formatted.CurrentAmount = campaign.CurrentAmount
	formatted.Slug = campaign.Slug

	user := campaign.User
	formattedUser := CampaignUserFormatter{}
	formattedUser.Name = user.Name
	formattedUser.ImageURL = user.AvatarFileName
	formatted.User = formattedUser

	images := []CampaignImagesFormatter{}
	for _, image := range campaign.CampaignImages {
		images = append(images, CampaignImagesFormatter{
			ImageURL:  image.FileName,
			IsPrimary: image.IsPrimary == 1,
		})
	}
	formatted.CampaignImages = images

	return formatted
}

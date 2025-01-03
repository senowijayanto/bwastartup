package campaign

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

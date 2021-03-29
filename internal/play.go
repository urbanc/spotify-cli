package internal

import (
	"spotify/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewPlayCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "play",
		Short:   "Play music.",
		PreRunE: checkLogin,
		RunE: func(cmd *cobra.Command, _ []string) error {
			token := viper.GetString("token")
			api := pkg.NewSpotifyAPI(token)

			return api.Play()
		},
	}
}

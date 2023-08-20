package cmd

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var retryTimes int
func init() {
	retryCmd.Flags().IntVar(&retryTimes, "times", 0, "retry times")
	rootCmd.AddCommand(retryCmd)
}

var retryCmd = &cobra.Command{
	Use:   "retry",
	Short: "retry strategy",
	Long:  `It uses exponential back-off strategy to calculate the retry delay.`,
	Run: func(cmd *cobra.Command, args []string) {
		t := DefaultRetryDelayFunc(retryTimes)
		fmt.Printf("default retry delay - n: [%d], delay: [%v]\n", retryTimes, t.Seconds())
	},
}

// DefaultRetryDelayFunc is the default RetryDelayFunc used if one is not specified in Config.
// It uses exponential back-off strategy to calculate the retry delay.
func DefaultRetryDelayFunc(n int) time.Duration {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := int(math.Pow(float64(n), 4)) + 15 + (r.Intn(30) * (n + 1))
	return time.Duration(s) * time.Second
}

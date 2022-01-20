/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// getallCmd represents the getall command
var getallCmd = &cobra.Command{
	Use:   "getall",
	Short: "get all users ",
	Long:  `Get Infomation of all users in details`,
	Run: func(cmd *cobra.Command, args []string) {
		getAllUsers()
	},
}

func init() {
	rootCmd.AddCommand(getallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getAllUsers() {
	res, err := http.Get("http://localhost:9010/users/")

	if err != nil {
		fmt.Println("Error While Fetching Data")
		return
	}

	var users []User

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&users)

	fmt.Println("List of all user: \n---")

	for _, usr := range users {

		fmt.Printf("id : %v\n", usr.Id)
		fmt.Printf("name : %v\n", usr.Name)
		fmt.Printf("email : %v\n", usr.Email)
		fmt.Printf("age : %v\n", usr.Age)
		fmt.Printf("gender : %v\n", usr.Gender)
		fmt.Printf("country : %v\n", usr.Country)
		fmt.Printf("status : %v\n", usr.Status)
		fmt.Println("---")

	}
}

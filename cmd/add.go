/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new user data",
	Long:  `Add new user data by adding flag --name {NAME} --email {EMAIL} --age {AGE} --gender {GENDER} --country {COUNTRY} --status {STATUS}`,
	Run: func(cmd *cobra.Command, args []string) {
		addUser()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&user_model.Name, "name", "n", "", "specify name")
	addCmd.MarkFlagRequired("name")
	addCmd.Flags().StringVarP(&user_model.Email, "email", "e", "", "specify email")
	addCmd.MarkFlagRequired("email")
	addCmd.Flags().IntVarP(&user_model.Age, "age", "a", -1, "specify age")
	addCmd.MarkFlagRequired("age")
	addCmd.Flags().StringVarP(&user_model.Gender, "gender", "g", "", "specify gender")
	addCmd.MarkFlagRequired("gender")
	addCmd.Flags().StringVarP(&user_model.Country, "country", "c", "", "specify country")
	addCmd.MarkFlagRequired("country")
	addCmd.Flags().StringVarP(&user_model.Status, "status", "s", "", "specify status")
	addCmd.MarkFlagRequired("status")

}

func addUser() {
	user_model.Id = 0
	jsonStr, _ := json.Marshal(user_model)
	// fmt.Printf((string(jsonStr)))

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if status := res.StatusCode; status != http.StatusOK {
		json.NewDecoder(res.Body).Decode(&custom_error)
		pj, _ := json.MarshalIndent(custom_error, "", "    ")
		fmt.Println(string(pj))
	} else {
		json.NewDecoder(res.Body).Decode(&user_model)
		fmt.Println("User Added: \n---")
		pj, _ := json.MarshalIndent(user_model, "", "    ")
		fmt.Println(string(pj))
		fmt.Println("---")
	}

}

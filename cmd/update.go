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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update new user data",
	Long:  "Update new user data by adding flag --name {NAME} --email {EMAIL} --age {AGE} --gender {GENDER} --country {COUNTRY} --status {STATUS}",
	Run: func(cmd *cobra.Command, args []string) {
		updateUser()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntVarP(&user_model.Id, "id", "i", -1, "specify id")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringVarP(&user_model.Name, "name", "n", "", "specify name")
	// updateCmd.MarkFlagRequired("name")
	updateCmd.Flags().StringVarP(&user_model.Email, "email", "e", "", "specify email")
	// updateCmd.MarkFlagRequired("email")
	updateCmd.Flags().IntVarP(&user_model.Age, "age", "a", -1, "specify age")
	// updateCmd.MarkFlagRequired("age")
	updateCmd.Flags().StringVarP(&user_model.Gender, "gender", "g", "", "specify gender")
	// updateCmd.MarkFlagRequired("gender")
	updateCmd.Flags().StringVarP(&user_model.Country, "country", "c", "", "specify country")
	// updateCmd.MarkFlagRequired("country")
	updateCmd.Flags().StringVarP(&user_model.Status, "status", "s", "", "specify status")
	// updateCmd.MarkFlagRequired("status")

}

func updateUser() {

	jsonStr, _ := json.Marshal(user_model)
	// fmt.Printf((string(jsonStr)))

	// fmt.Println(URL + strconv.Itoa(id))
	req, err := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonStr))
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
		fmt.Println("User Updated: \n---")
		pj, _ := json.MarshalIndent(user_model, "", "    ")
		fmt.Println(string(pj))
		fmt.Println("---")
	}

}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var Name string
var Email string
var Age int
var Gender string
var Country string
var Status string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new user data",
	Long:  `Add new user data by adding flag --name {NAME} --email {EMAIL} --age {AGE} --`,
	Run: func(cmd *cobra.Command, args []string) {
		addUser()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&Name, "name", "j", "", "specify name")
	addCmd.MarkFlagRequired("name")
	addCmd.Flags().StringVarP(&Email, "email", "e", "", "specify email")
	addCmd.MarkFlagRequired("email")
	addCmd.Flags().IntVarP(&Age, "age", "a", -1, "specify age")
	addCmd.MarkFlagRequired("age")
	addCmd.Flags().StringVarP(&Gender, "gender", "g", "", "specify gender")
	addCmd.MarkFlagRequired("gender")
	addCmd.Flags().StringVarP(&Country, "country", "w", "", "specify country")
	addCmd.MarkFlagRequired("country")
	addCmd.Flags().StringVarP(&Status, "status", "q", "", "specify status")
	addCmd.MarkFlagRequired("status")

}

func addUser() {
	// fmt.Printf("here")
	s := fmt.Sprintf(`
	{
		"name":"%v",
		"email":"%v",
		"age":%v,
		"gender":"%v",	
		"country":"%v",
		"status":"%v"

	}`, Name, Email, Age, Gender, Country, Status) + "\n"
	fmt.Printf(s)
	var jsonStr = []byte(s)
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

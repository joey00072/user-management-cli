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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateUser()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntVarP(&id, "id", "i", -1, "specify id")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringVarP(&Name, "name", "j", "", "specify name")
	// updateCmd.MarkFlagRequired("name")
	updateCmd.Flags().StringVarP(&Email, "email", "e", "", "specify email")
	// updateCmd.MarkFlagRequired("email")
	updateCmd.Flags().IntVarP(&Age, "age", "a", -1, "specify age")
	// updateCmd.MarkFlagRequired("age")
	updateCmd.Flags().StringVarP(&Gender, "gender", "g", "", "specify gender")
	// updateCmd.MarkFlagRequired("gender")
	updateCmd.Flags().StringVarP(&Country, "country", "w", "", "specify country")
	// updateCmd.MarkFlagRequired("country")
	updateCmd.Flags().StringVarP(&Status, "status", "q", "", "specify status")
	// updateCmd.MarkFlagRequired("status")

}

func updateUser() {

	if id == -1 {
		fmt.Println("id is required")
		return
	}
	s := "{\n"
	if id != -1 {
		s += fmt.Sprintf("\"id\": %d", id)
	}
	if Name != "" {
		s += fmt.Sprintf(",\n\"name\": \"%s\"", Name)
	}
	if Email != "" {
		s += fmt.Sprintf(",\n\"email\": \"%s\"", Email)
	}
	if Age != -1 {
		s += fmt.Sprintf(",\n\"age\": %d", Age)
	}
	if Gender != "" {
		s += fmt.Sprintf(",\n\"gender\": \"%s\"", Gender)
	}
	if Country != "" {
		s += fmt.Sprintf(",\n\"country\": \"%s\"", Country)
	}
	if Status != "" {
		s += fmt.Sprintf(",\n\"status\": \"%s\"", Status)
	}

	s += "\n}"

	fmt.Printf(s)
	var jsonStr = []byte(s)
	// fmt.Println(URL + strconv.Itoa(id))
	req, err := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonStr))
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

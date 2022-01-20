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

var id int

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get user data with id",
	Long:  `Get spectfic using by adding flag --id {ID} `,
	Run: func(cmd *cobra.Command, args []string) {
		getUser(id)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().IntVarP(&id, "id", "f", -1, "specify id")
	getCmd.MarkFlagRequired("id")

}

func getUser(id int) {

	fmt.Println(id)

	res, err := http.Get(fmt.Sprintf("http://localhost:9010/users/%d", id))

	if err != nil {
		fmt.Println("Error While Fetching Data")
		return
	}

	var user User

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&user)

	fmt.Println("User: \n---")

	fmt.Printf("id : %v\n", user.Id)
	fmt.Printf("name : %v\n", user.Name)
	fmt.Printf("email : %v\n", user.Email)
	fmt.Printf("age : %v\n", user.Age)
	fmt.Printf("gender : %v\n", user.Gender)
	fmt.Printf("country : %v\n", user.Country)
	fmt.Printf("status : %v\n", user.Status)
	fmt.Println("---")
}

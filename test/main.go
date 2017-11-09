package main

import (
	"log"
	"os"
	"time"

	"github.com/ustream/terraform-provider-composeio/composeapi"
)

func main() {

	apiToken := os.Getenv("BM_API_KEY")
	areneMySQLDeployment := "bmix-eude-yp-dacd993c-8989-47c8-96a5-01a8ea4a99f4"

	client, err := composeapi.NewClient(apiToken, composeapi.BxEuDeApiBase)

	client.SetLogger(true, os.Stdout)

	if err != nil {
		log.Fatal(err)
	}

	recipe, errs := client.AddWhitelistForDeployment(areneMySQLDeployment, composeapi.Whitelist{IP: "1.2.3.4/32", Description: "terraform teszt"})

	if errs != nil {
		log.Fatal(errs)
	}

	log.Println(recipe)

	whitelist, errs := client.GetWhitelistForDeployment(areneMySQLDeployment)

	if errs != nil {
		log.Fatal(errs)
	}

	log.Println(whitelist.Embedded.Whitelist)

	recipe, errs = client.DeleteWhitelistForDeployment(areneMySQLDeployment, whitelist.Embedded.Whitelist[0].ID)

	if errs != nil {
		log.Fatal(errs)
	}

	log.Println(recipe)

	time.Sleep(time.Second * 30)

	recipe, errs = client.GetRecipe(areneMySQLDeployment, recipe.ID)

	if errs != nil {
		log.Fatal(errs)
	}

	log.Println(recipe)
}
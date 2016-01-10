package commands

import (
    "fmt"
    "github.com/JFrogDev/bintray-cli-go/utils"
)

func CreatePackage(packageDetails *utils.VersionDetails, flags *utils.PackageFlags) {
    data := utils.CreatePackageJson(packageDetails.Package, flags)
    url := flags.BintrayDetails.ApiUrl + "packages/" + packageDetails.Subject + "/" +
        packageDetails.Repo

    fmt.Println("Creating package: " + packageDetails.Package)
    resp, body := utils.SendPost(url, []byte(data), flags.BintrayDetails.User, flags.BintrayDetails.Key)
    if resp.StatusCode != 201 {
        utils.Exit(resp.Status + ". " + utils.ReadBintrayMessage(body))
    }
    fmt.Println("Bintray response: " + resp.Status)
    fmt.Println(utils.IndentJson(body))
}
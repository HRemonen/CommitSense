/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import "commitsense/cmd"

var (
	version	 string
	date    string
)

func main() {
	cmd.SetVersion(version, date)
	cmd.Execute()
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"imuslab.com/arozos/mod/utils"
)

// Handle console command from the console module
func consoleCommandHandler(input string) string {
	//chunk := strings.Split(input, " ");
	chunk, err := parseCommandLine(input)
	if err != nil {
		return err.Error()
	}
	if len(chunk) > 0 && chunk[0] == "auth" {
		if matchSubfix(chunk, []string{"auth", "new"}, 4, "auth new {username} {password}") {
			return "Creating a new user " + chunk[2] + " with password " + chunk[3]
		} else if matchSubfix(chunk, []string{"auth", "dump"}, 4, "auth dump {filename}.csv") {
			filename := chunk[2]
			fmt.Println("Dumping user list to " + filename + " csv file")
			csv := authAgent.ExportUserListAsCSV()
			err := os.WriteFile(filename, []byte(csv), 0755)
			if err != nil {
				return err.Error()
			}
			return "OK"
		}
	} else if len(chunk) > 0 && chunk[0] == "permission" {
		if matchSubfix(chunk, []string{"permission", "list"}, 2, "") {
			fmt.Println("> ", permissionHandler.PermissionGroups)
			return "OK"
		} else if matchSubfix(chunk, []string{"permission", "user"}, 3, "permission user {username}") {
			username := chunk[2]
			group, _ := permissionHandler.GetUsersPermissionGroup(username)
			for _, thisGroup := range group {
				fmt.Println(thisGroup)
			}
			return "OK"
		} else if matchSubfix(chunk, []string{"permission", "group"}, 3, "permission group {groupname}") {
			groupname := chunk[2]
			groups := permissionHandler.PermissionGroups
			for _, thisGroup := range groups {
				if thisGroup.Name == groupname {
					fmt.Println(thisGroup)
				}
			}
			return "OK"
		} else if matchSubfix(chunk, []string{"permission", "getinterface"}, 3, "permission getinterface {username}") {
			//Get the list of interface module for this user
			userinfo, err := userHandler.GetUserInfoFromUsername(chunk[2])
			if err != nil {
				return err.Error()
			}
			return strings.Join(userinfo.GetInterfaceModules(), ",")
		}
	} else if len(chunk) > 0 && chunk[0] == "quota" {
		if matchSubfix(chunk, []string{"quota", "user"}, 3, "quota user {username}") {
			userinfo, err := userHandler.GetUserInfoFromUsername(chunk[2])
			if err != nil {
				return err.Error()
			}

			fmt.Println("> "+"User Quota: ", userinfo.StorageQuota.UsedStorageQuota, "/", userinfo.StorageQuota.GetUserStorageQuota(), "bytes")
			return "OK"
		}
	} else if len(chunk) > 0 && chunk[0] == "database" {
		if matchSubfix(chunk, []string{"database", "dump"}, 3, "database dump {filename}") {
			//Dump the database to file

			return "WIP"
		} else if matchSubfix(chunk, []string{"database", "list", "tables"}, 3, "") {
			//List all opened tables
			sysdb.Tables.Range(func(k, v interface{}) bool {
				fmt.Println(k.(string))
				return true
			})
			return "OK"
		} else if matchSubfix(chunk, []string{"database", "view"}, 3, "database list {tablename}") {
			//List everything in this table
			tableList := []string{}

			sysdb.Tables.Range(func(k, v interface{}) bool {
				tableList = append(tableList, k.(string))
				return true
			})
			if !utils.StringInArray(tableList, chunk[2]) {
				return "Table not exists"
			} else if chunk[2] == "auth" {
				return "You cannot view this database table"
			}
			entries, err := sysdb.ListTable(chunk[2])
			if err != nil {
				return err.Error()
			}

			for _, keypairs := range entries {
				fmt.Println("> " + string(keypairs[0]) + ":" + string(keypairs[1]))
			}

			fmt.Println("Total Entry Count: ", len(entries))
			return "OK"
		}
	} else if len(chunk) > 0 && chunk[0] == "user" {
		if matchSubfix(chunk, []string{"user", "object", "dump"}, 4, "user object dump {username}") {
			//Dump the given user object as json
			userinfo, err := userHandler.GetUserInfoFromUsername(chunk[3])
			if err != nil {
				return err.Error()
			}

			jsonString, _ := json.Marshal(userinfo)
			return string(jsonString)
		} else if matchSubfix(chunk, []string{"user", "quota"}, 3, "user quota {username}") {
			//List user quota of the given username
			userinfo, err := userHandler.GetUserInfoFromUsername(chunk[2])
			if err != nil {
				return err.Error()
			}

			fmt.Println(userinfo.StorageQuota.UsedStorageQuota, "/", userinfo.StorageQuota.TotalStorageQuota)
			return "OK"
		}
	} else if len(chunk) > 0 && chunk[0] == "storage" {
		if matchSubfix(chunk, []string{"storage", "list", "basepool"}, 3, "") {
			//Dump the base storage pool
			jsonString, _ := json.Marshal(userHandler.GetStoragePool())
			return string(jsonString)
		}
	} else if len(chunk) > 0 && chunk[0] == "scan" {
		if matchSubfix(chunk, []string{"scan", "all"}, 2, "") {
			//scan all nearby wdos units
			fmt.Println("Scanning (Should take around 10s)")
			hosts := MDNS.Scan(10, "")
			for _, host := range hosts {
				fmt.Println(host)
			}
			return "OK"
		} else if matchSubfix(chunk, []string{"scan", "wdos"}, 2, "") || matchSubfix(chunk, []string{"scan", "wdos"}, 2, "") {
			//scan all nearby wdos units
			fmt.Println("Scanning nearybe WDOS Hosts (Should take around 10s)")
			hosts := MDNS.Scan(10, "wdos.com")
			for _, host := range hosts {
				fmt.Println(host)
			}
			return "OK"
		}
	} else if len(chunk) > 0 && chunk[0] == "find" {
		if matchSubfix(chunk, []string{"find", "module"}, 3, "list module {modulename}") {
			//Display all loaded modules
			for _, module := range moduleHandler.LoadedModule {
				if strings.ToLower(module.Name) == strings.ToLower(chunk[2]) {
					jsonString, _ := json.Marshal(module)
					return string(jsonString)
				}
			}
			return string("Module not found")

		} else if matchSubfix(chunk, []string{"find", "modules"}, 2, "") {
			//Display all loaded modules
			jsonString, _ := json.Marshal(moduleHandler.LoadedModule)
			return string(jsonString)
		} else if matchSubfix(chunk, []string{"find", "subservices"}, 2, "") {
			//Display all loaded subservices
			fmt.Println(ssRouter.RunningSubService)
			return "OK"
		}
	} else if len(chunk) > 0 && chunk[0] == "access" {
		//Handle emergency situation where the user is blocked by himself
		if matchSubfix(chunk, []string{"access", "whitelist", "disable"}, 3, "") {
			//Disable whitelist
			authAgent.WhitelistManager.SetWhitelistEnabled(false)
			return "Whitelist Disabled"
		} else if matchSubfix(chunk, []string{"access", "whitelist", "enable"}, 3, "") {
			//Enable whitelist
			authAgent.WhitelistManager.SetWhitelistEnabled(true)
			return "Whitelist Enabled"
		} else if matchSubfix(chunk, []string{"access", "whitelist", "add"}, 4, "access whitelist add {ip_range}") {
			err = authAgent.WhitelistManager.SetWhitelist(chunk[3])
			if err != nil {
				return err.Error()
			}
			return "OK"
		} else if matchSubfix(chunk, []string{"access", "whitelist", "del"}, 4, "access whitelist del {ip_range}") {
			err = authAgent.WhitelistManager.UnsetWhitelist(chunk[3])
			if err != nil {
				return err.Error()
			}
			return "OK"
		} else if matchSubfix(chunk, []string{"access", "blacklist", "enable"}, 3, "") {
			//Enable blacklist
			authAgent.WhitelistManager.SetWhitelistEnabled(true)
			return "Blacklist Enabled"
		} else if matchSubfix(chunk, []string{"access", "blacklist", "disable"}, 3, "") {
			//Disable blacklist
			authAgent.BlacklistManager.SetBlacklistEnabled(false)
			return "Blacklist Disabled"
		} else {
			return "[Whitelist / Blacklist Console Control API] \nUsage: access {whitelist/blacklist} {action} {data}"
		}
	} else if len(chunk) == 1 && chunk[0] == "stop" {
		//Stopping the server
		fmt.Println("Shutting down wdos online system by terminal request")
		executeShutdownSequence()
	}

	return "Invalid Command. Given: '" + strings.Join(chunk, " ") + "'"
}

// Check if the given line input match the requirement
func matchSubfix(chunk []string, match []string, minlength int, usageExample string) bool {
	matching := true
	//Check if the chunk contains minmium length of the command request
	if len(chunk) >= len(match) {
		for i, cchunk := range match {
			if chunk[i] != cchunk {
				matching = false
			}
		}
	} else {
		matching = false
	}

	if len(chunk)-minlength == -1 && chunk[len(chunk)-1] == match[len(match)-1] {
		fmt.Println("Paramter missing. Usage: " + usageExample)
		return false
	}

	return matching
}

func parseCommandLine(command string) ([]string, error) {
	var args []string
	state := "start"
	current := ""
	quote := "\""
	escapeNext := true
	for i := 0; i < len(command); i++ {
		c := command[i]

		if state == "quotes" {
			if string(c) != quote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = "start"
			}
			continue
		}

		if escapeNext {
			current += string(c)
			escapeNext = false
			continue
		}

		if c == '\\' {
			escapeNext = true
			continue
		}

		if c == '"' || c == '\'' {
			state = "quotes"
			quote = string(c)
			continue
		}

		if state == "arg" {
			if c == ' ' || c == '\t' {
				args = append(args, current)
				current = ""
				state = "start"
			} else {
				current += string(c)
			}
			continue
		}

		if c != ' ' && c != '\t' {
			state = "arg"
			current += string(c)
		}
	}

	if state == "quotes" {
		return []string{}, errors.New(fmt.Sprintf("Unclosed quote in command line: %s", command))
	}

	if current != "" {
		args = append(args, current)
	}

	return args, nil
}

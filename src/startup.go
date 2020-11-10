package main

/*
	System Startup Script for ArOZ Online System
	author: tobychui
*/

import (
	"log"

	//ArOZ Online Core Modules
	db "imuslab.com/aroz_online/mod/database"
	permission "imuslab.com/aroz_online/mod/permission"
	
)



func RunStartup(){
	//1. Initiate the main system database
	if !fileExists("system/"){
		panic("▒▒ ERROR: SYSTEM FOLDER NOT FOUND ▒▒")
	}

	if !fileExists("web/"){
		panic("▒▒ ERROR: WEB FOLDER NOT FOUND ▒▒")
	}

	dbconn, err := db.NewDatabase("system/ao.db", false)
	if err != nil{
		panic(err)
	}
	sysdb = dbconn;

	//2. Initiate the auth Agent
	AuthInit();	//See auth.go

	//3. Start Permission Management Module
	ph, err := permission.NewPermissionHandler(sysdb)
	if err != nil{
		log.Println("Permission Handler creation failed.")
		panic(err)
	}
	permissionHandler = ph
	permissionHandler.LoadPermissionGroupsFromDatabase();



	//4. Mount and create the storage system base
	StorageInit(); 					//See storage.go

	//5. Startup user and permission sytem
	UserSystemInit(); 				//See user.go
	permissionInit(); 				//Register permission interface after user
	RegisterSystemInit(); 			//See register.go
	
	//6.Start Modules, Subservice and AGI loaders
	ModuleServiceInit();			//Module Handler
	PackagManagerInit();			//Start APT service agent
	AGIInit();						//ArOZ Javascript Gateway Interface
	SubserviceInit();				//Subservice Handler

	//7. Kickstart the File System and Desktop
	FileSystemInit();				//Start FileSystem
	DesktopInit();					//Start Desktop
	HardwarePowerInit();			//Start host power manager

	//8. Initiate System Settings Handlers
	SystemSettingInit();			//Start System Setting Core 
	DiskQuotaInit();				//Disk Quota Management
	DiskServiceInit();				//Start Disk Services
	DeviceServiceInit();			//Client Device Management
	SystemInfoInit();				//System Information UI
	SystemIDInit();					//System UUID Manager
	AdvanceSettingInit();			//System Advance Settings

	//9. Startup network services
	NetworkServiceInit();
	WiFiInit();


	//10. Other legacy stuffs
	system_time_init();
	util_init();
	system_resetpw_init();
	mediaServer_init();

	//Finally
	moduleHandler.ModuleSortList();				//Sort the system module list
	
}
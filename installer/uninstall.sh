#!/bin/bash
cat << "EOF"
    _               ___  ___   ___   __  
   /_\  _ _ ___ ___/ _ \/ __| |_  ) /  \ 
  / _ \| '_/ _ \_ / (_) \__ \  / / | () |
 /_/ \_\_| \___/__|\___/|___/ /___(_)__/ 
                                         
	----- WDOS 2.0 Uninstall -----	
	
EOF

# Ask user to confirm uninstall
read -p "Are you sure you want to uninstall WDOS? This will delete all data in the wdos directory. (y/n) " choice
case "$choice" in
  y|Y )
	# Stop the WDOS service if it is running
	if [[ $(uname) == "Linux" ]]; then
		if systemctl status wdos >/dev/null 2>&1; then
			sudo systemctl stop wdos
			echo "Stopped WDOS service."
		fi
	fi

	# Remove the WDOS folder
	cd ~/ || exit
	if [[ -d "wdos" ]]; then
		sudo rm -rf wdos
		echo "Removed WDOS folder."
	fi

	# Remove the WDOS service file
	if [[ $(uname) == "Linux" ]]; then
		if [[ -f "/etc/systemd/system/wdos.service" ]]; then
			sudo rm /etc/systemd/system/wdos.service
			echo "Removed WDOS systemd service file."
		fi
	fi
	sudo systemctl daemon-reload
	echo "WDOS has been uninstalled successfully!"
	;;
  n|N ) 
    echo "Uninstall cancelled"
    ;;
  * ) 
    echo "Invalid input, uninstall cancelled"
    ;;
esac

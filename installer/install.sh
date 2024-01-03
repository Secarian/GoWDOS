#!/bin/bash
cat << "EOF"
    _               ___  ___   ___   __  
   /_\  _ _ ___ ___/ _ \/ __| |_  ) /  \ 
  / _ \| '_/ _ \_ / (_) \__ \  / / | () |
 /_/ \_\_| \___/__|\___/|___/ /___(_)__/ 
                                         
	----- WDOS 2.0 Installer -----	
	
EOF

echo ""

# Prompt the user to agree to the GPLv3 license
read -p "Do you agree to the terms of the GPLv3 license? (y/n) " agree

if [[ $agree != "y" ]]; then
  echo "You must agree to the GPLv3 license to use WDOS."
  exit 1
fi

# Create the required folder structure to hold the installation
cd ~/ || exit
mkdir wdos
cd wdos || exit

# Run apt-updates
sudo apt-get update
sudo apt-get install ffmpeg net-tools -y

# Determine the CPU architecture of the host
if [[ $(uname -m) == "x86_64" ]]; then
  arch="amd64"
elif [[ $(uname -m) == "aarch64" ]]; then
  arch="arm64"
elif [[ $(uname -m) == "armv"* ]]; then
  arch="arm"
else
  read -p "Enter the target architecture (e.g. darwin_amd64, windows_amd64): " arch
fi

# Download the corresponding executable from Github
if [[ $arch == "amd64" ]]; then
  download_url="https://github.com/Secarian/GoWDOS/releases/latest/download/wdos_linux_amd64"
elif [[ $arch == "arm64" ]]; then
  download_url="https://github.com/Secarian/GoWDOS/releases/latest/download/wdos_linux_arm64"
elif [[ $arch == "arm" ]]; then
  download_url="https://github.com/Secarian/GoWDOS/releases/latest/download/wdos_linux_arm"
elif [[ $arch == "windows_amd64" ]]; then
  download_url="https://github.com/Secarian/GoWDOS/releases/latest/download/wdos_windows_amd64.exe"
elif [[ $arch == "windows_arm64" ]]; then
  download_url="https://github.com/Secarian/GoWDOS/releases/latest/download/wdos_windows_arm64.exe"
else
  download_url="https://github.com/Secarian/GoWDOS/releases/latest/download/wdos_${arch}"
fi

# Download the wdos binary
echo "Downloading WDOS from ${download_url} ..."
wget -O wdos "${download_url}"
chmod +x wdos

# Download the webpack
wget -O web.tar.gz "https://github.com/Secarian/GoWDOS/releases/latest/download/web.tar.gz"

# Check if the platform is supported for the launcher
if [[ "$arch" == "amd64" || "$arch" == "arm" || "$arch" == "arm64" ]]; then
  # Ask if the user wants to install the launcher
  read -p "Do you want to install the WDOS launcher for OTA updates? [Y/n] " answer
  case ${answer:0:1} in
      y|Y )
          # Download the appropriate binary
          echo "Downloading WDOS launcher from https://github.com/wdos-online/launcher/releases/latest/ ..."
          case "$arch" in
              amd64)
                  launcher_url="https://github.com/wdos-online/launcher/releases/latest/download/launcher_linux_amd64"
                  ;;
              arm)
                  launcher_url="https://github.com/wdos-online/launcher/releases/latest/download/launcher_linux_arm"
                  ;;
              arm64)
                  launcher_url="https://github.com/wdos-online/launcher/releases/latest/download/launcher_linux_arm64"
                  ;;
              *)
                  echo "Unsupported architecture for WDOS launcher"
                  ;;
          esac
          if [[ -n "$launcher_url" ]]; then
              wget -O launcher "${launcher_url}"
              chmod +x launcher
              echo "WDOS launcher has been installed successfully!"
          fi
          ;;
      * )
          echo "WDOS launcher installation skipped"
          ;;
  esac
fi

# Ask for setup name
read -p "Enter setup name (default: wdos): " wdosname
wdosname=${wdosname:-wdos}

# Ask for preferred listening port
read -p "Enter preferred listening port (default: 8080): " wdosport
wdosport=${wdosport:-8080}

# Check if launcher exists
if [[ -f "./launcher" ]]; then
  # Create start.sh with launcher command
  echo "#!/bin/bash" > start.sh
  echo "sudo ./launcher -port=$wdosport -hostname=\"$wdosname\"" >> start.sh
else
  # Create start.sh with wdos command
  echo "#!/bin/bash" > start.sh
  echo "sudo wdos -port=$wdosport -hostname=\"$wdosname\"" >> start.sh
fi

# Make start.sh executable
chmod +x start.sh

echo "Setup name: $wdosname"
echo "Preferred listening port: $wdosport"
echo "start.sh created successfully!"

# Ask if user wants to install WDOS to systemd
if [[ $(uname) == "Linux" ]]; then
    read -p "Do you want to install WDOS to systemd service? (y/n)" -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        # Get current user
        CURRENT_USER=$(whoami)
		
		sudo touch /etc/systemd/system/wdos.service
		sudo chmod 777 /etc/systemd/system/wdos.service
        # Create systemd service file
        cat <<EOF > /etc/systemd/system/wdos.service
[Unit]
Description=WDOS Cloud Service
After=systemd-networkd-wait-online.service
Wants=systemd-networkd-wait-online.service

[Service]
Type=simple
ExecStartPre=/bin/sleep 10
WorkingDirectory=/home/${CURRENT_USER}/wdos/
ExecStart=/bin/bash /home/${CURRENT_USER}/wdos/start.sh

Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
EOF
		sudo chmod 644 /etc/systemd/system/wdos.service
		
        # Reload systemd daemon and enable service
        sudo systemctl daemon-reload
        sudo systemctl enable wdos.service
		sudo systemctl start wdos.service
        echo "WDOS installation completed!"
		ip_address=$(hostname -I | awk '{print $1}')
		echo "Please continue the system setup at http://$ip_address:$wdosport/"
    fi
else
	echo "WDOS installation completed! Execute start.sh to startup your WDOS system."
fi



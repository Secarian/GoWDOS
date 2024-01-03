![Image](img/banner.png?raw=true)

<img src="https://img.shields.io/badge/License-GPLv3-blue"> <img src="https://img.shields.io/badge/Device-Raspberry%20Pi%203B%2B%20%2F%204B-red"> <img src="https://img.shields.io/badge/Made%20In%20Hong%20Kong-香港開發-blueviolet">

## WICHTIGE HINWEISE

Der aktuelle Arozos befindet sich noch in intensiver Entwicklung. Die Systemstruktur kann jederzeit geändert werden. Bitte entwickeln Sie nur auf der aktuellen vorhandenen ArOZ Gateway Interface (AGI) JavaScript Interface oder Standard HTML Webapps mit ao_module.js Endpunkten.

## Features

### Benutzeroberfläche

- Web-Desktop-Schnittstelle (besser als Synology DSM)
- Ubuntu-Remix Windows-Stil Startmenü und Taskleisten
- Sauber und einfach zu bedienender Dateimanager (Unterstützung Drag & Drop, Upload usw.)
- Einfaches Systemeinstellungsmenü
- Kein-Bullshit-Modulnamenschema

### Netzwerk

- FTP-Server
- Statischer Webserver
- WebDAV-Server
- UPnP-Portweiterleitung
- Samba (über Drittanbieter-Subservices unterstützt)
- WLAN-Verwaltung (Unterstützung von wpa_supplicant für Rpi oder nmcli für Armbian)

### Datei-/Festplattenverwaltung

- Laufwerks-/Formatierungs-Dienstprogramme (unterstützt NTFS, EXT4 und mehr!)
- Virtuelle Dateisystemarchitektur
- Dateifreigabe (ähnlich wie Google Drive)
- Grundlegende Dateioperationen mit Echtzeit-Fortschritt (Kopieren / Ausschneiden / Einfügen / Neue Datei oder Ordner usw.)

### Sicherheit

- oAuth
- LDAP
- IP-Weiß-/Schwarzliste
- Exponentielles Anmeldetimeout

### Erweiterbarkeit

- ECMA5 (JavaScript ähnliche) Skript-Schnittstelle
- Entwicklung von 3rd-Party-Go / Python-Modulen mit Sub-Service-Reverse-Proxy

### Andere

- Erfordert so wenig wie 512 MB Arbeitsspeicher und 8 GB Systemspeicher
- Basiert auf einer der stabilsten Linux-Distributionen - Debian
- Unterstützung für Desktop-, Laptop- (Touchpad-) und Mobilbildschirmgrößen

## Installation

Erfordert GO 1.14 oder höher (siehe [Installationsanleitung](https://dev.to/tobychui/install-go-on-raspberry-pi-os-shortest-tutorial-4pb))

Führen Sie den folgenden Befehl aus, um das System zu erstellen

```
git clone <https://github.com/Secarian/GoWDOS>
cd ./wdos/src/
Build starten
./wdos
#sudo ./wdos zur Aktivierung der Hardware- und WiFi-Verwaltungsfunktionen
```

Ja, es ist so einfach.

## Bereitstellen

### Für Raspberry Pi (Für Raspberry Pi 4B+)

Wenn Sie Raspberry Pi als Ihren Host verwenden, können Sie eines der Images herunterladen und das Image auf Ihre SD-Karte flashen. Sie finden ein neues Netzwerkgerät mit dem Namen "WDOS (ARxxx)" in Ihrem "Netzwerkumfeld". Doppelklicken Sie auf das Symbol und Sie werden auf die System-Websetup-Schnittstelle umgeleitet. Wenn Sie das neue Gerät nicht in Ihrem Netzwerk finden können, können Sie sich auch direkt über `http://{raspberry_pi_ip_address}:8080/` mit WDOS verbinden.

| Version | Download | Kopie | Kommentar |
| --- | --- | --- | --- |
| wdos-v1.120 | https://www.mediafire.com/file/rg3gkdt4asgk789/wdos_v1.120.7z/file | https://drive.google.com/file/d/1neTDxFpXxQQzsHvyqmCJOfAzjCPi8RzY/view?usp=sharing | |
| wdos v1.119 | https://www.mediafire.com/file/4vx4f5boj8pfeu1/wdos_v1.119.7z/file | https://drive.google.com/file/d/1Gl_wYCvbio2lmW6YiFObIJHlejLzFrRu/view?usp=sharing | Aktualisiert auf Raspberry Pi OS 64-bit. Siehe kompatible Liste https://www.raspberrypi.com/news/raspberry-pi-os-64-bit/ |
| wdos v1.118 (v2) | https://www.mediafire.com/file/f1i4xsp4rplwbko/wdos_v1.118_v2.7z/file | https://drive.google.com/file/d/1sgG-QOlaUmXhSiUJIB3DpnejElud1yvn/view?usp=sharing | Unterstützt Pi zero 2w |
| wdos v1.115 (Stable) | https://www.mediafire.com/file/zbhieo59fq2sw80/wdos_v1.115.7z/file | | Bauen in https://github.com/aroz-online/WsTTY |
| wdos v1.114 | EOL | | Unstable, aktualisieren Sie auf 1.115, wenn Sie diese Version noch benutzen |
| wdos v1.113 | https://www.mediafire.com/file/u42ha6ljfq6q0g9/wdos_v1.113.7z/file | |
| wdos v1.112 (Stable) | https://www.mediafire.com/file/eonn1weu8jvfz29/wdos_v1.112.7z/file | | Fehlerbehebung und Patches für v1.111 |
| wdos v1.111 (Stable) | https://www.mediafire.com/file/cusm5jwsuey6b4k/wdos_v1.111.7z/file | | IoT Hub hinzugefügt |
| wdos v1.110 (Stable) | http://www.mediafire.com/file/r7l40jv727covej/wdos_v1.110.7z/file | | wdos v1.110 (Stable) | http://www.mediafire.com/file/r7l40jv727covej/wdos_v1.110.7z/file
| wdos v1.109 | https://www.mediafire.com/file/mmjyv77ei9fwab5/wdos_v1.109.7z/file | |
| wdos v1.108 | https://www.mediafire.com/file/aa8176setz3ljtv/wdos_v1.108.7z/file | | WebDAV-Unterstützung hinzugefügt |
| wdos v1.107 | | https://drive.google.com/file/d/1klI6fVaSLHFr213kI35W6a6hYyBUSIrI/view?usp=sharing | Samba-Unterstützung hinzugefügt |
| wdos v1.106 | | https://drive.google.com/file/d/1ysZxeIQ5dBu7x5eEyCDwHtMJmJoMp9El/view?usp=sharing | |

Alle oben aufgelisteten Imges erfordern eine microSD-Karte mit 8 GB oder mehr.

Um die .img-Datei zu erhalten, können Sie das komprimierte Bild mit 7zip entpacken. Wenn Sie es nicht haben, können Sie es [hier](https://www.7-zip.org/download.html) herunterladen.

### Für alle Pi-Modelle

### Erstellen Sie mit dem Installationsskript aus dem Quellcode (Empfohlen)

Seit v1.119 wurde das vorinstallierte wdos-Images von Raspberry Pi OS 32-Bit auf 64-Bit verschoben, um die Systemressourcen besser zu nutzen. Für ältere Versionen von Pis können Sie wdos mit dem unten stehenden Befehl mit einer frischen Installation von Raspberry Pi OS installieren.

```
curl -L <https://raw.githubusercontent.com/tobychui/wdos/master/installer/install_for_pi.sh> | bash

```

ohne curl

```
cd ~/
wget <https://raw.githubusercontent.com/tobychui/wdos/master/installer/install_for_pi.sh>
sudo chmod 775 ./install_for_pi.sh
./install_for_pi.sh

```

Der Installer wird alle erforderlichen Abhängigkeiten einschließlich ffmpeg und go-Compiler installieren. Um die Installation erfolgreich zu bestätigen, überprüfen Sie den Ausführungsstatus von Arozos mit dem folgenden Befehl.

```
sudo systemctl status wdos

```

### Verwenden eines vorgefertigten Binärpakets

Siehe Installationsschritte für andere ARM SBC (verwenden Sie jedoch das Binärpaket `wdos_linux_arm` anstelle von `wdos_linux_arm64`).

### Für andere ARM SBC (z.B. Orange Pi / Banana Pi / Friendly ARM's Pis)

Laden Sie das korrekte Architekturbinärpaket aus dem Tab "release" herunter und laden Sie das Binärpaket mit dem Ordner "web" und "system" in "/src" hoch. Nach dem Hochladen sollten Sie die folgende Dateistruktur haben.

```
$ ls
wdos_linux_arm64  web  system

```

Starten Sie das Binärprogramm, indem Sie `sudo ./wdos_linux_arm64` aufrufen (oder ohne sudo, wenn Sie keine Hardwareverwaltung bevorzugen).

### Windows

Wenn Sie auf Windows bereitstellen, müssen Sie ffmpeg zur %PATH%-Umgebungsvariable hinzufügen.

Dieses System kann mit den folgenden Build-Anweisungen auf Windows-Hosts erstellt und ausgeführt werden.

```
# Laden Sie das gesamte Repo als zip herunter und installieren Sie es mit cd
cd .\\wdos\\src\\
Build starten
wdos.exe
```

**Allerdings sind nicht alle Funktionen für Windows verfügbar**.

## Docker
Danke an [Saren](https://github.com/Saren-Arterius) für das erstellen dieses großartigen DockerFile

Sehen Sie sich sein Repo an [hier] (https://github.com/Saren-Arterius/aroz-dockerize)

## Screenshots
![Image](img/screenshots/1.png?raw=true)
![Image](img/screenshots/2.png?raw=true)
![Image](img/screenshots/3.png?raw=true)
![Image](img/screenshots/4.png?raw=true)
![Image](img/screenshots/5.png?raw=true)
![Image](img/screenshots/6.png?raw=true)

## Starten Sie die WDOS-Plattform

### Unterstützte Startparameter

Die folgenden Startparameter werden unterstützt (v1.113)

```
-allow_autologin
    	Erlaubt eine RESTFUL Login-Umleitung, die es Maschinen wie Billboards erlaubt, sich beim Systemstart anzumelden (Standardwert true)
  -allow_cluster
    	Erlaubt Cluster-Operationen im LAN. Erfordert das allow_mdns=true Flag (Voreinstellung true)
  -allow_iot
    	Ermöglicht IoT-bezogene APIs und Scanner. Erfordert die Aktivierung von MDNS (Standard true)
  -allow_mdns
    	Aktiviert den MDNS-Dienst. Erlaubt, dass das Gerät von nahegelegenen ArOZ-Hosts gescannt wird (standardmäßig true)
  -allow_pkg_install
    	Erlaubt dem System, Pakete mit dem Advanced Package Tool (auch bekannt als apt oder apt-get) zu installieren (Standardwert true)
  -allow_ssdp
    	Aktivieren Sie den SSDP-Dienst. Deaktivieren Sie diese Option, wenn Sie nicht möchten, dass Ihr Gerät von der Windows-Netzwerkumgebung gescannt wird (Standardwert: true)
  -allow_upnp
    	Aktiviert den uPNP-Dienst, empfohlen für Hosts unter NAT-Routern
  -beta_scan
    	Erlaubt Kompatibilität zu ArOZ Online Beta Clusters
  -cert Zeichenfolge
    	TLS-Zertifikatsdatei (.crt) (Standard "localhost.crt")
  -console
    	Aktiviert die Debugging-Konsole.
  -demo_Modus
    	Führt das System im Demomodus aus. Alle Verzeichnisse und die Datenbank sind schreibgeschützt.
  -dir_list
    	Aktiviert die Verzeichnisauflistung (Voreinstellung true)
  -disable_http
    	HTTP-Server deaktivieren, tls=true erforderlich
  -disable_ip_resolver
    	Deaktiviert die IP-Auflösung, wenn das System unter einer Reverse-Proxy-Umgebung läuft
  -disable_subservice
    	Vollständige Deaktivierung von Unterdiensten
  -enable_hwman
    	Aktiviert die Hardware-Verwaltungsfunktionen im System (Standardwert true)
  -gzip
    	Aktiviert die gzip-Komprimierung auf dem Dateiserver (Voreinstellung true)
  -homepage
    	Aktiviert die Benutzer-Homepage. Erreichbar über /www/{Benutzername}/ (Voreinstellung true)
  -hostname Zeichenfolge
    	Standardname für diesen Host (standardmäßig "Mein ArOZ")
  -iobuf int
    	Menge des Pufferspeichers für IO-Operationen (Standardwert 1024)
  -key string
    	TLS-Schlüsseldatei (.key) (Voreinstellung "localhost.key")
  -max_upload_size int
    	Maximale Upload-Größe in MB. Darf den verfügbaren Speicherplatz auf Ihrem System nicht überschreiten (Standardwert 8192)
  -ntt int
    	Ausführungszeit der nächtlichen Aufgaben. Standardwert 3 = 3 Uhr morgens (Standardwert 3)
  -port int
    	Lauschender Port für HTTP-Server (Standardwert 8080)
  -public_reg
    	Öffentliche Registerschnittstelle für die Kontoerstellung aktivieren
  -root string
    	Benutzer-Root-Verzeichnisse (standardmäßig "./files/")
  -session_key Zeichenfolge
    	Sitzungsschlüssel, muss 16, 24 oder 32 Bytes lang sein (AES-128, AES-192 oder AES-256). Leer lassen für automatische Generierung.
  -storage_config Zeichenfolge
    	Dateispeicherort der Speicherkonfigurationsdatei (Standard "./system/storage.json")
  -tls
    	Aktiviert TLS bei der HTTP-Übertragung (HTTPS-Modus)
  -tls_port int
    	Lauschender Port für HTTPS-Server (Standardwert 8443)
  -tmp string
    	Temporärer Speicher, kann über tmp:/ erreicht werden. Ein tmp/ Ordner wird in diesem Pfad erstellt. Empfohlen werden schnelle Speichergeräte wie SSD (Standard "./")
  -tmp_zeit int
    	Zeit bevor die tmp Datei gelöscht wird in Sekunden. Standardwert 86400 Sekunden = 24 Stunden (Standardwert 86400)
  -upload_async
    	Aktiviert die Pufferung des Datei-Uploads, um im asynchronen Modus zu arbeiten (schnellerer Upload, erfordert RAM >= 8GB)
  -upload_buf int
    	Pufferspeicher für den Upload in MB. Jede Datei, die größer als diese Größe ist, wird auf der Festplatte gepuffert (langsamer). (Voreinstellung 25)
  -uuid string
    	System UUID für Clustering und verteiltes Rechnen. Muss nur einmal beim ersten Start konfiguriert werden. Leer lassen für automatische Generierung.
  -version
    	System-Build-Version anzeigen
  -wlan_schnittstelle_name string
    	Die standardmäßige drahtlose Schnittstelle für die Verbindung mit einem AP (standardmäßig "wlan0")
  -wpa_supplicant_config string
    	Pfad für die wpa_supplicant-Konfiguration (standardmäßig "/etc/wpa_supplicant/wpa_supplicant.conf")
```

Beispiel:

```
//Start von aroz online mit Standard-Webport
./wdos -port 80

//Start von aroz online im Demo-Modus
./wdos -demo_mode=true

//https anstelle von http verwenden
./wdos -tls=true -tls_port 443 -key mykey.key -cert mycert.crt -disable_http=true

//Starten Sie sowohl den HTTPS- als auch den HTTP-Server auf zwei verschiedenen Ports
./wdos -port 80 -tls=true -key mykey.key -cert mycert.crt -tls_port 443

//Ändern Sie die maximale Upload-Größe auf 25MB
./wdos -max_upload_size 25
```

Siehe Dokumentation für weitere Beispiele.

### WDOS Launcher (erforderlich für OTA-Update-Unterstützung)

Siehe [https://github.com/aroz-online/launcher](https://github.com/aroz-online/launcher)

### Speicherkonfiguration

### Bereitstellung einer einzelnen Maschine

Wenn Sie eine einzelne Maschine bereitstellen, können Sie im Systemeinstellungen > Disk & Storage > Storage Pools die Storage-Pool "System" bearbeiten, um die globalen Storage-Pools für alle Benutzer im System einzurichten.

![](img/screenshots/sp.png)

### Bereitstellen auf mehreren Maschinen

Wenn Sie auf mehreren Maschinen bereitstellen, können Sie sich die Konfigurationsdatei für den Speicher ansehen, die sich unter folgendem Pfad befindet:

```
src/system/storage.json.example

```

Benennen Sie die storage.json.example in storage.json um und starten Sie wdos. Die erforderlichen virtuellen Speicherlaufwerke werden entsprechend eingehängt.

## ArOZ JavaScript Gateway Interface (AGI) / Plugin Loader

Das ArOZ AJGI / AGI-Interface bietet eine javascript-programmierbare Schnittstelle für ArOZ Online-Benutzer, um Plugins für das System zu erstellen. Um das Modul zu initiieren, können Sie eine "init.agi"-Datei im Webverzeichnis des Moduls (auch als Modulwurzel bezeichnet) ablegen. Weitere Details finden Sie in der![AJGI Dokumenation](https://github.com/Secarian/GoWDOS/blob/master/src/AGI%20Documentation.md).

## WDOS OTA Update Launcher
Seit Version 1.119 kann WDOS mit Hilfe des [WDOS Launchers](https://github.com/aroz-online/launcher) ein OTA-Update durchführen. Weitere Informationen zur Installation finden Sie im Github-Repository des Launchers.

## Weitere Ressourcen
Wenn Sie nach anderen WebApps (über die Systemeinstellungen installiert) oder Subservices (erfordern eine SSH-Anmeldung zur Installation, nur für OEM) suchen, schauen Sie sich bitte unsere Sammlung hier an:
https://github.com/aroz-online/WebApp-and-Subservice-Index

### 💬 Direkter Kontakt

Sie können die Autoren über [Telegram](https://t.me/ArOZBeta) erreichen! Wir freuen uns über alle Arten von Feedback und Fragen.

### 🖥️ Hardware Kompatibilität

Verwenden Sie WDOS auf etwas anderem als Raspberry Pis? Zeigen Sie uns Ihre Serverkonfiguration und Hardware-Spezifikationen!

https://github.com/Secarian/GoWDOS/issues/50

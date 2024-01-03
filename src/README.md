# WDOS 2.0

Dies ist die Go-Implementierung von WDOS (auch bekannt als WDOS Online) Web-Desktop-Umgebung, die für Linux entwickelt wurde, aber irgendwie auch auf Windows und macOS funktioniert.

Diese README-Datei ist nur für Entwickler gedacht. Wenn Sie ein normaler Benutzer sind, konsultieren Sie bitte die README-Datei außerhalb des /src-Ordners.

## Entwicklerhinweise

- Beginnen Sie jedes Modul mit der Funktion {ModuleName}Init(), z. B. ```WiFiInit()```
- Platzieren Sie Ihre Funktion im Modul (falls möglich) und rufen Sie sie im Hauptprogramm auf
- Ändern Sie die Reihenfolge in der startup() Funktion nicht, es sei denn, es ist notwendig
- Bei Unsicherheit fügen Sie Startflags hinzu (und verwenden Sie Startflags, um experimentelle Funktionen beim Start zu deaktivieren)

## Überschreiben von Vendor-Ressourcen

Wenn Sie Vendor-bezogene Ressourcen in WDOS 2.012 oder höher überschreiben möchten, erstellen Sie einen Ordner im Systemstammverzeichnis mit dem Namen ```vendor-res``` und legen Sie die Ersatzdateien hier ab. Hier ist eine Liste der unterstützten Ersatzressourcendateien:

| Dateiname       | Empfohlenes Format | Verwendung                 |
| --------------- | ------------------ | -------------------------- |
| auth_bg.jpg     | 2938 x 1653 px     | Login-Hintergrund          |
| auth_icon.png   | 5900 x 1180 px     | Logo auf der Authentifizierungsseite |
| vendor_icon.png | 1560 x 600 px      | Marken-Symbol des Anbieters |

(Zum Ausbau)

## Dateisystem-Virtualisierung und Abstraktionsebenen

Das WDOS-System enthält sowohl die Virtualisierungsebene als auch die Abstraktionsebene. Der einfachste Weg, um zu überprüfen, ob Ihr Pfad unter welcher Ebene liegt, besteht darin, den Startverzeichnisnamen zu betrachten.

| Pfadstruktur                                  | Beispiel-Pfad                                     | Ebene                                            |
| --------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| {vroot_id}:/{subpath}                         | user:/Desktop/meinedatei.txt                       | Dateisystem-Virtualisierungsebene (höchste Ebene) |
| fsh (*File System Handler) + subpath (string)  | fsh (localfs) + /dateien/benutzer/alan/Desktop/meinedatei.txt | Dateisystem-Abstraktionsebene                     |
| {physical_location}/{subpath}                 | /home/wdos/wdos/dateien/benutzer/Desktop/meinedatei.txt | Physische (Festplatten-)Ebene                     |

Seit WDOS v2.000 haben wir der (bereits komplexen) File System Handler (fsh) Infrastruktur eine Dateisystem-Abstraktion (fsa oder manchmal als fshAbs, Abkürzung für "File System Handler underlying File System Abstraction") hinzugefügt. Es gibt zwei Arten von fsh, die derzeit von der WDOS Dateisystem-Abstraktionsebene unterstützt werden.

## WDOS JavaScript Gateway Interface / Plugin Loader

Die WDOS AJGI / AGI-Schnittstelle bietet eine JavaScript-programmierbare Schnittstelle für WDOS-Benutzer, um ein Plugin für das System zu erstellen. Um das Modul zu initialisieren, können Sie eine "init.agi"-Datei im Webverzeichnis des Moduls (auch Modulwurzel genannt) platzieren. Weitere Details finden Sie in der [AJGI-Dokumentation](AJGI Dokumentation.md).

AGI-Skripte können mit verschiedenen Berechtigungen ausgeführt werden.

| Bereich                                      | Verwendbare Funktionen                                       |
| -------------------------------------------- | ------------------------------------------------------------ |
| WebApp-Startskript (init.agi)                | Systemfunktionen und Registrierungen                          |
| In WebApp enthaltene Skripte                 | Systemfunktionen und Benutzerfunktionen                       |
| Andere (Web-Root / Serverless / Scheduler)    | Systemfunktionen, Benutzerfunktionen (mit Script-Registrierungseigentümerbereich) und serverlos |

## Unterlogiken und Konfiguration von Subdiensten

Um andere binär basierte Webserver in die Subdienstschnittstelle zu integrieren, können Sie einen Ordner im "./subservice/your_service" erstellen, in dem Ihre binäre Ausführung den gleichen Namen wie das enthaltende Verzeichnis haben sollte. Wenn Sie z.B. ein Modul haben, das eine Web-Benutzeroberfläche namens "demo.exe" bereitstellt, sollten Sie die demo.exe in "./subservice/demo/demo.exe" platzieren.

Im Fall einer Linux-Umgebung wird die Subdienstroutine zuerst überprüfen, ob das Modul über apt-get installiert ist, indem sie das Programm "which" verwendet. (Wenn Sie busybox haben, sollte es integriert sein.) Wenn das Paket nicht in der apt-Liste gefunden wird, wird die binäre Datei des Programms im Subservice-Verzeichnis gesucht.

Bitte beachten Sie die Namenskonvention in der build.sh-Vorlage. Zum Beispiel wird die entsprechende Plattform nach dem entsprechenden ausführbaren Binärdateinamen suchen:

```
demo_linux_amd64    => Linux AMD64
demo_linux_arm      => Linux ARMv6l / v7l
demo_linux_arm64    => Linux ARM64
demo_macOS_amd64    => MacOS AMD64 
```

### Startflags

Während des Starts des Subdienstes werden zwei Arten von Parametern übergeben. Hier sind Beispiele:

```
demo.exe -info
demo.exe -port 12810 -rpt "http://localhost:8080/api/ajgi/interface"
```

Im Falle des Empfangs des "info"-Flags sollte das Programm die JSON-Zeichenfolge mit den korrekten Modulinformationen ausgeben, wie unten beschrieben.

```
//Struktur zur Speicherung von Modulinformationen
type serviecInfo struct{
    Name string                //Name dieses Moduls, z. B. "Audio"
    Desc string                //Beschreibung für dieses Modul
    Group string            //Gruppe des Moduls, z. B. "system" / "media" etc.
    IconPath string            //Pfad zum Modul-Symbolbild, z. B. "Audio/img/function_icon.png"
    Version string            //Version des Moduls. Format: [0-9]*.[0-9][0-9].[0-9]
   

 StartDir string         //Standardstartverzeichnis, z. B. "Audio/index.html"
    SupportFW bool             //Unterstützung von FloatWindow. Wenn ja, wird das FloatWindow-Verzeichnis geladen
    LaunchFWDir string         //Dieser Link wird anstelle von 'StartDir' im fw-Modus gestartet
    SupportEmb bool            //Eingebetteter Modus unterstützen
    LaunchEmb string         //Dieser Link wird gestartet, wenn eine Datei mit diesem Modul geöffnet wird
    InitFWSize []int         //Floatwindow-Startgröße. [0] => Breite, [1] => Höhe
    InitEmbSize []int        //Größe für den eingebetteten Modus. [0] => Breite, [1] => Höhe
    SupportedExt []string     //Unterstützte Dateierweiterungen, z. B. ".mp3", ".flac", ".wav"
}

//Beispiel für die Verwendung beim Empfangen des -info-Flags
infoObject := serviecInfo{
        Name: "Demo Subservice",
        Desc: "Ein einfacher Subdienst-Code, um zu zeigen, wie Subdienste in WDOS Online funktionieren",            
        Group: "Entwicklung",
        IconPath: "demo/icon.png",
        Version: "0.0.1",
        StartDir: "demo/home.html",            
        SupportFW: true, 
        LaunchFWDir: "demo/home.html",
        SupportEmb: true,
        LaunchEmb: "demo/embedded.html",
        InitFWSize: []int{720, 480},
        InitEmbSize: []int{720, 480},
        SupportedExt: []string{".txt",".md"},
    }

jsonString, _ := json.Marshal(info);
fmt.Println(string(infoObject))
os.Exit(0);
```

Wenn das Port-Flag empfangen wird, sollte das Programm die Web-Benutzeroberfläche am angegebenen Port starten. Im Folgenden finden Sie ein Beispiel für die Implementierung einer solchen Funktionalität.

```
var port = flag.String("port", ":80", "Der Standardendpunkt zum Zuhören für diesen Subdienst")
flag.Parse()
err := http.ListenAndServe(*port, nil)
if err != nil {
    log.Fatal(err)
}
```

### Subdienst-Ausführungseinstellungen

Standardmäßig wird die Subdienstroutine eine Reverse-Proxy-Verbindung mit integriertem URL-Umschreiben erstellen, die Ihre Web-Benutzeroberfläche bedient, die von der ausführbaren Binärdatei gestartet wurde. Wenn Sie keine Reverse-Proxy-Verbindung benötigen, ein benutzerdefiniertes Startskript oder etwas anderes wünschen, können Sie die folgenden Einstellungsdateien verwenden.

```
.noproxy        => Kein Proxy zum angegebenen Port starten
.startscript    => Die Startparameter an die Datei "start.bat" oder "start.sh" senden, anstatt die ausführbare Binärdatei zu verwenden
.disabled        => Diesen Subdienst beim Start nicht laden. Der Benutzer kann ihn jedoch über die Einstellungsschnittstelle aktivieren
```

Hier ist ein Beispiel für eine "start.bat"-Datei, die in die Integration von Syncthing in das WDOS Online System mit einer ".startscript"-Datei integriert wird, die neben der syncthing.exe-Datei platziert wird.

```
if not exist ".\config" mkdir ".\config"
syncthing.exe -home=".\config" -no-browser -gui-address=127.0.0.1%2
```

## Systemd-Unterstützung

Um systemd in Ihrem Host zu aktivieren, der das WDOS Online System unterstützt, erstellen Sie ein Bash-Skript in Ihrem WDOS Online-Stammverzeichnis namens "start.sh" und füllen Sie es mit Ihren bevorzugten Startparametern. Das einfachste ist wie folgt:

```
#/bin/bash
sudo ./wdos_online_linux_amd64
```

Dann können Sie eine neue Datei namens "wdos.service" in /etc/systemd/system mit folgendem Inhalt erstellen (nehmen Sie an, dass Ihr WDOS Online-Stammverzeichnis bei /home/pi/wdos liegt):

```
[Unit]
Description=WDOS Cloud Desktop Service.

[Service]
Type=simple
WorkingDirectory=/home/pi/wdos/
ExecStart=/bin/bash /home/pi/wdos/start.sh

Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

Schließlich verwenden Sie die folgenden systemd-Befehle, um den Dienst zu aktivieren, zu starten, den Status anzuzeigen und den Dienst zu deaktivieren:

```
# Aktivieren Sie das Skript während des Startvorgangs
sudo systemctl enable wdos.service

# Starten Sie den Dienst jetzt
sudo systemctl start wdos.service

# Den Status des Dienstes anzeigen
systemctl status wdos.service

# Deaktivieren Sie den Dienst, wenn Sie ihn beim Start nicht mehr ausführen möchten
sudo systemctl disable wdos-online.service
```

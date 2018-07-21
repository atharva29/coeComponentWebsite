#include <ESP8266WiFi.h>
int count = 0 ;
int counter = 0 ;
String ssid = "iball-baton" ; // wifi ssid
String password = "ATad2913"; // wifi password
String ip = "192.168.1.104" ; // ip address of server
int port = 8000 ; // port of server
String val = "" ;
WiFiClient client1; //client object

char* string2char(String command) {
  if (command.length() != 0)
  {
    char*p = const_cast<char*>(command.c_str());
    return p;
  }
}




void setup() {
  // put your setup code here, to run once:
  Serial.begin(19200);
}

void loop() {
  // put your main code here, to run repeatedly:
  clientMode();
}

String readLaser() {
  Serial.print("D");
  while (!Serial.available());
  val = Serial.readString();
  return val ;
}




void clientMode() {
  if (count == 0 ) {
    WiFi.mode(WIFI_STA);
    WiFi.begin(string2char(ssid), string2char(password) );
    while (WiFi.status() != WL_CONNECTED)
    {
      delay(250);
    }
    delay(100);
    while (!client1.connect(string2char(ip), port))
    {
      Serial.println("Connecting to server");
      delay(500);
    }
    count = 2 ; // go to count = 1
  }

  else if (count == 2 ) { // connection to server is successful
    if (client1.connected()) {
      if (client1.available() > 0 ) { // server is sending some data
        String Command = client1.readStringUntil('\n'); // read data
        if (Command == "D") {
          client1.println(readLaser());
        }
    } else { // if connection is lost then go to count = 0
      count = 0 ;
    }
  }
}

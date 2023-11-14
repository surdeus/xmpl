void setup() {
  // put your setup code here, to run once:
  pinMode(12, OUTPUT);
}

const int Delay = 500 ;

void loop() {
  digitalWrite(12, HIGH);  // turn the LED on (HIGH is the voltage level)
  delay(Delay);                      // wait for a second
  digitalWrite(12, LOW);   // turn the LED off by making the voltage LOW
  delay(Delay);  
}

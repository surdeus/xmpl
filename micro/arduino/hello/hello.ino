const int Out = 12 ;
const int Delay = 500 ;

void setup() {
	pinMode(Out, OUTPUT);
}

void loop() {
	delay(Delay);
	digitalWrite(Out, HIGH);
	delay(Delay);
	digitalWrite(Out, LOW);
}


"use strict";

fetch("x-counter.htm")
	.then(stream => stream.text())
	.then(text => define(text))

function define(htm){
customElements.define("x-counter", class extends HTMLElement{
	constructor() {
		super();
		var shadow = this.attachShadow({mode: "open"}) ;
		shadow.innerHTML = htm ;
		this.valueElement = shadow.getElementById("value") ;
		var incrementButton = shadow.getElementById("increment-button") ;
		var decrementButton = shadow.getElementById("decrement-button") ;
		console.log(this.valueElement);
		incrementButton.onclick = () => {this.increment()} ;
		decrementButton.onclick = () => {this.decrement()} ;

		if(this.hasAttribute("value"))
			this.value = parseInt(this.getAttribute("value")) ;
		else
			this.value = 1 ;

		if(this.hasAttribute("step"))
			this.step = parseInt(this.getAttribute("step")) ;
		else
			this.step = 1 ;
	}

	increment() {
		this.value += this.step ;
	}

	decrement() {
		this.value -= this.step ;
	}

	set value(value) {
		this._value = value ;
		this.valueElement.innerHTML = value ;
	}

	get value() {
		return this._value
	}

	attributeChangedCallback(name, oldval, newval){
		switch(name){
		case "value" :
			this.value = int(newval) ;
		break;
		case "step" :
			this.step = int(newval) ;
		}
	}
})}

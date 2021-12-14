"use strict";

fetch("body.html")
	.then(stream => stream.text())
	.then(text => define(text))

customElements.define("show-hello", class extends HTMLElement {
	// Most of the callbacks are described.

	constructor() {
		super();
	}

	// Element is added.
	connectedCallback() { setTimeout(() =>{
		const shadow = this.attachShadow({mode: "open"}) ;
		shadow.innerHTML = `<p>
			Hello, ${this.getAttribute("name")}.
			${this.innerHTML}
		</p>` ;
	}); }

	// Element is removed.
	disconnectedCallback() {
	}

	// Array of attributes to callback on their change.
	static get observedAttributes() {
		return ["name"] ;
	}
	// Callback itself.
	attributeChangedCallback(name, oldval, newval) {
		console.log(`${name} was ${oldval} became ${newval}`)
	}
});


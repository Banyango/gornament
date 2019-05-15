(function() {
    class StaffCard extends HTMLElement {
        constructor() {
            super();
        }
        connectedCallback() {
            const shadowRoot = this.attachShadow({mode:'open'});

            const template = document.querySelector('#staff-card-template');
            const instance = template.content.cloneNode(true);

            this.shadowRoot.appendChild(instance);

            const name = this.getAttribute('person');
            const type = this.getAttribute('type');
            const number = this.getAttribute('number');

            this.shadowRoot.getElementById("person").innerHTML = name;
            this.shadowRoot.getElementById("type").innerHTML = type;
            this.shadowRoot.getElementById("number").innerHTML = number;
        }
    }

    window.customElements.define('staff-card', StaffCard);
})();


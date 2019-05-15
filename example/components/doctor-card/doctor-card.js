(function () {
    class DoctorCard extends HTMLElement {
        constructor() {
            super();
        }
        connectedCallback() {
            const shadowRoot = this.attachShadow({mode:'open'});

            const template = document.querySelector('#doctor-card-template');
            const instance = template.content.cloneNode(true);

            shadowRoot.appendChild(instance);
            shadowRoot.getElementById("doc_name").innerHTML = this.getAttribute('doc-name');
        }
    }

    window.customElements.define('doctor-card', DoctorCard);
})();

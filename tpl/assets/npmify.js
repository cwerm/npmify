class PxMarquee extends HTMLElement {
    constructor() {
        // Always call super first in constructor
        super();

        // Create a shadow root
        const shadow = this.attachShadow({mode: 'open'});

        // Create spans
        const marquee = document.createElement('h1');
        marquee.setAttribute('class', 'marquee');

        const info = document.createElement('span');
        info.setAttribute('class', 'info');

        const infoContainer = document.createElement('span');
        infoContainer.setAttribute('class', 'scrolling');

        const outdatedEl = document.createElement('span');
        outdatedEl.setAttribute('class', 'blink');
        outdatedEl.textContent = this.getAttribute('data-outdated');

        // Take attribute content and put it inside the info span
        const text = document.createElement('span');
        text.setAttribute('class', 'rainbow');
        text.textContent = ` out of ${this.getAttribute('data-total')} dependencies are outdated.`;
        // info.textContent = text;

        // Create some CSS to apply to the shadow dom
        const style = document.createElement('style');

        style.textContent = `
        /*! minireset.css v0.0.5 | MIT License | github.com/jgthms/minireset.css */html,body,p,ol,ul,li,dl,dt,dd,blockquote,figure,fieldset,legend,textarea,pre,iframe,hr,h1,h2,h3,h4,h5,h6{margin:0;padding:0}h1,h2,h3,h4,h5,h6{font-size:100%;font-weight:normal}ul{list-style:none}button,input,select,textarea{margin:0}html{box-sizing:border-box}*,*:before,*:after{box-sizing:inherit}img,video{height:auto;max-width:100%}iframe{border:0}table{border-collapse:collapse;border-spacing:0}td,th{padding:0;text-align:left}
        h1, h2, h3 {
            font-weight: 700;
        }
        
        h1 {
            font-size: 2rem;
            padding: 2rem;
        }
        
        h2 {
            font-size: 1.5rem;
        }
        .marquee {
          width: 100vw;
          white-space: nowrap;
          overflow: hidden;
          box-sizing: border-box;
        }
        
        .marquee .scrolling {
          display: inline-block;
          padding-left: 100%;
          /* show the marquee just outside the paragraph */
          animation: marquee 15s linear infinite;
        }
        
        .marquee .scrolling:hover {
          animation-play-state: paused
        }
        
        .marquee .blink {
            color: #db5147;
            animation: 1s linear infinite blinky;
        }
        
        .rainbow {
            animation: 2s linear infinite colorchange;
        }

        /* Make it move */
        
        @keyframes marquee {
          0% {
            transform: translate(0, 0);
          }
          100% {
            transform: translate(-100%, 0);
          }
        }
        
        @keyframes blinky {
            0% {
                visibility: hidden;
            }
            50% {
                visibility: hidden;
            }
            100% {
                visibility: visible;
            }
        }
        
        @-webkit-keyframes colorchange {
            0% {
                color: red;
            }
            16.6666666667% {
                color: orange;
            }
            33.3333333333% {
                color: yellow;
            }
            50% {
                color: green;
            }
            66.6666666667% {
                color: blue;
            }
            83.3333333333% {
                color: indigo;
            }
            100% {
                color: purple;
            }
        }
    `;

        // Attach the created elements to the shadow dom
        shadow.appendChild(style);
        shadow.appendChild(marquee);
        info.appendChild(outdatedEl);
        info.appendChild(text);
        infoContainer.appendChild(info)
        marquee.appendChild(infoContainer);
    }
}

const $ = (selector) => {
    return document.querySelector(selector);
};

const $$ = (selector) => {
    return document.querySelectorAll(selector);
};

customElements.define('px-marquee', PxMarquee);

const fancyBtn = $('.makeFancy');

const toggleHeader = () => {
    const headerEls = $$('.header');

    headerEls.forEach(h => {
        if (h.classList.contains('hidden')) {
            h.classList.remove('hidden');
        } else {
            h.classList.add('hidden');
        }
    })
};

fancyBtn.addEventListener('click', toggleHeader, null);


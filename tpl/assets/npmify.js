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

// Remember jQuery?!?!
const $ = (selector) => {
    return document.querySelector(selector);
};

const $$ = (selector) => {
    return document.querySelectorAll(selector);
};

// Go template to compile the dependencies to js.
let deps = [
    {{ range .Bower }}
    {
        name: "{{ .Name }}",
        version: "{{ .Version }}",
        npmVersion: "{{ .NpmVersion }}",
        license: "{{ .License }}",
        outdated: {{ .Outdated }},
        bowerType: "{{ .Type }}",
    },
{{ end }}
];

// Append the list of dependencies to the #app div.
// const buildList = (dependencies) => {
//     return $('#app').innerHTML = `
//     ${dependencies.map(dep => {
//         let isChecked = localStorage.getItem(dep.name) === 'true';
//         // console.log(dep.name, isChecked)
//         return `${dep.outdated ? `<div class="dependency outdated">` : `<div class="dependency">`}
//             ${isChecked ? `<h2 class="title done">${dep.name}</h2>` : `<h2 class="title">${dep.name}</h2>`}
//             <div class="grid">
//                 <strong>Bower Version:</strong> <span>${dep.version}</span>
//                 <strong>NPM Version:</strong> <span>${dep.npmVersion}</span>
//                 <strong>Bower Type:</strong> <span>${dep.bowerType}</span>
//                 <a class="btn" href="https://www.npmjs.com/package/${dep.name}" target="_blank">NPM Link</a>
//                 <span class="grid-spacer"></span>
//                 <label for="${dep.name}-done">Done?</label>
//                 ${isChecked ? `<input id="${dep.name}-done" class="isDone" type="checkbox" value="${dep.name}" checked />` :
//             `<input id="${dep.name}-done" class="isDone" type="checkbox" value="${dep.name}" />`
//             }
//
//             </div>
//         </div>`
//     }).join('')}
// `
// }

const buildList = (dependencies) => {
    return $('#list').innerHTML = `
    ${dependencies.map(dep => {
        let isChecked = localStorage.getItem(dep.name) === 'true';
        console.log(typeof dep.license)
        // console.log(dep.name, isChecked)
            return `<div class="cursor-pointer my-1 px-2 hover:bg-blue-300 rounded border-b">
                <div class="flex">
                    <div class="w-8 text-center pt-1">
                      ${dep.outdated ? `<p class="text-3xl text-red-500">&bull;</p>` : `<p class="text-3xl p-0 text-green-500">&bull;</p>`} 
                    </div>
                    <div class="w-3/5 py-3 px-1">
                      <p class="hover:text-blue-700 font-bold">${dep.name}</p>
                    </div>
                    <div class="w-2/5 text-right p-3">
                      <p class="text-sm flex text-grey-500 font-light truncate align-center justify-end">
                        <img class="w-8 mr-4" src="assets/bower-logo.svg" alt="bower logo" /> v${dep.version}
                      </p>
                    </div>
                </div>
                <div class="flex">
                    <div class="w-3/5 py-3 px-1"> 
                        ${dep.license ? 
                            `<span class="text-sm flex text-grey-500 font-light truncate align-center justify-start">
                                <svg width="16" height="16" class="mr-4" viewBox="0 0 14 16" version="1.1" aria-hidden="true">
                                    <path fill-rule="evenodd"
                                        d="M7 4c-.83 0-1.5-.67-1.5-1.5S6.17 1 7 1s1.5.67 1.5 1.5S7.83 4 7 4zm7 6c0 1.11-.89 2-2 2h-1c-1.11 0-2-.89-2-2l2-4h-1c-.55 0-1-.45-1-1H8v8c.42 0 1 .45 1 1h1c.42 0 1 .45 1 1H3c0-.55.58-1 1-1h1c0-.55.58-1 1-1h.03L6 5H5c0 .55-.45 1-1 1H3l2 4c0 1.11-.89 2-2 2H2c-1.11 0-2-.89-2-2l2-4H1V5h3c0-.55.45-1 1-1h4c.55 0 1 .45 1 1h3v1h-1l2 4zM2.5 7L1 10h3L2.5 7zM13 10l-1.5-3-1.5 3h3z"></path>
                                </svg> 
                                ${dep.license.includes('GPL') ? `<span class="text-red-600 font-bold"> ${dep.license}</span>` : `<span>${dep.license}</span>` }` : `<span class="text-red-600 font-bold">NO LICENSE</span>`
                         }
                        </p>
                    </div>
                    <div class="w-2/5 text-right p-3">
                      <p class="text-sm flex text-grey-500 font-light truncate align-center justify-end">
                        <img class="w-8 mr-4" src="assets/npm-logo.svg" alt="npm logo" /> v${dep.npmVersion}
                      </p>
                    </div>
                </div>
              </div>`
    }).join('')}
`
};

customElements.define('px-marquee', PxMarquee);

const fancyBtn = $('.makeFancy');
const filterBtns = $$('button[data-filter]');
const searchFilter = $('#searchFilter');

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

const filterOutdated = (prop) => {
    if (prop === "reset") {
        return buildList(deps);
    }
    let filtered = deps.filter(dep => {
        if (prop === "outdated") {
            return !dep.outdated;
        } else {
            return dep.outdated;
        }
    });

    return buildList(filtered);
};

const filterByName = (name) => {
    let filtered = deps.filter(dep => {
        return dep.name.indexOf(name) > -1;
    });

    return buildList(filtered);
}

const store = {
    set: (key, val) => {
        return new Promise(resolve => resolve(localStorage.setItem(key, val)))
    },
    get: (key) => {
        return new Promise((resolve, reject) => {
            if (localStorage.getItem(key)) {
                resolve(localStorage.getItem(key));
            }
            reject('No key/val pair found.')
        })
    }
};

const markDone = (dep, done) => {
    store.set(dep, done)
        .then(() => {
            buildList(deps)
        })
};

fancyBtn.addEventListener('click', toggleHeader, null);

filterBtns.forEach(btn => {
    btn.addEventListener('click', () => { filterOutdated(btn.dataset.filter); });
});

window.addEventListener('DOMContentLoaded', () => {
    buildList(deps);

    const isDoneCheckboxes = $$('.isDone');

    isDoneCheckboxes.forEach(check => {
        check.addEventListener('change', (e) => {
            markDone(e.target.value, e.target.checked);
        });
    });

    searchFilter.addEventListener('keyup', e => {
        filterByName(e.target.value);
    })

});

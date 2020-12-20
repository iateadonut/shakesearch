const Controller = {
  search: (ev, page=0) => {
    if(ev) ev.preventDefault();
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    const response = fetch(`/search?q=${data.query}&page=${page}`).then((response) => {
      response.json().then((results) => {
        Controller.updateTable(results);
      });
    });
  },

  pageNumber: null,

  goForward: (ev, page) => {
    console.log('running');
    Controller.search(event, Controller.pageNumber);
  },

  goBackward: (ev, page) => {
    console.log('running');
    Controller.search(event, Controller.pageNumber - 2);
  },

  bindPages: () => {
    const pageBack = document.getElementById("page-back");
    pageBack.onclick = Controller.goBackward;

    const pageForward = document.getElementById("page-forward");
    pageForward.onclick = Controller.goForward;
  },

  updateTable: (results) => {
    
    const numFound = results['response']['numFound'];
    document.getElementById('number-results-found').innerHTML = "Results found: " + numFound;

    const pageBack = document.getElementById('page-back');
    const pageForward = document.getElementById('page-forward');
    const pageNumber = document.getElementById('page-number');
    const startFrom = results['response']['start'];

    const page = startFrom/10 + 1;
    Controller.pageNumber = page;

    pageBack.innerHTML = null;
    pageForward.innerHTML = null;

    if( 1 != page)
      pageBack.innerHTML = '< BACK | ';
    //console.log( Math.ceil(numFound/10) + ' ' + page );
    if( Math.ceil(numFound/10) > page )
      pageForward.innerHTML = ' | NEXT >';
    pageNumber.innerHTML = ' Page: '+page;

    Controller.bindPages();

    const table = document.getElementById("table-body");
    table.innerHTML = "";
    const rows = [];

    for (let result of results['response']['docs']) {
      rows.push(`${result.line}`);
      if ( result.poem_title ) {
        rows.push('- Poem: ' + result.poem_title);
      } else {
        rows.push(`- Play: ${result.play_title}: ${result.acttitle}, ${result.scenetitle}`);
      }
      rows.push('');
      
      
    }
    console.log(rows);
    for (row of rows) {
      let newRow = table.insertRow(-1);
      let newCell = newRow.insertCell(0);
      newCell.appendChild(document.createTextNode(row));
      //table.insertRow(document.createTextNode('hi'));
    }
    
  },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);

// const pageBack = document.getElementById("page-back");

// const pageForward = document.getElementById("page-forward");
// pageForward.addEventListener("onclick", Controller.goForward(2));

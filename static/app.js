const Controller = {
  search: (ev) => {
    ev.preventDefault();
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    const response = fetch(`/search?q=${data.query}`).then((response) => {
      response.json().then((results) => {
        Controller.updateTable(results);
      });
    });
  },

  updateTable: (results) => {
    const table = document.getElementById("table-body");
    const rows = [];
    console.log(results['response']['docs']);
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

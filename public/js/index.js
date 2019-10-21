$(document).ready(function () {
  let options = {
    processing: true,
    serverSide: true,
    ajax: {
      url: "/teams",
      dataSrc: 'data'
    },
    columns: [{
        data: 'Rank'
      },
      {
        data: 'TeamID'
      },
      {
        data: 'TeamName'
      },
      {
        data: 'TotalScore'
      }
    ],
    error: function (xhr, textStatus, error) {
      console.log("error: ", textStatus);
    }
    // "deferRender": true
  };

  $('#teams').DataTable(options);
});

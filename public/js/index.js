$(document).ready(function () {
  let options = {
    processing: true,
    serverSide: true,
    pagingType: "full_numbers",
    ajax: {
      url: "/teams",
      dataSrc: 'data'
    },
    columns: [{
        data: 'Rank'
      },
      {
        data: 'TeamID',
        orderable: false
      },
      {
        data: 'TeamName'
      },
      {
        data: 'TotalScore',
        orderable: false
      }
    ],
    error: function (xhr, textStatus, error) {
      console.log("error: ", textStatus);
    }
    // "deferRender": true
  };

  $('#teams').DataTable(options);
});

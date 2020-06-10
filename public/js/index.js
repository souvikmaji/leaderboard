$(document).ready(function () {
  // Hide dataTable error alert
  $.fn.dataTable.ext.errMode = "throw";

  let options = {
    processing: true,
    serverSide: true,
    pagingType: "simple_numbers",
    ordering: false,
    searching: false,
    ajax: {
      url: "/game/leaderboard",
      dataSrc: "data"
    },
    columns: [{
        data: "Rank"
      },
      {
        data: "GameID",
        // orderable: false
      },
      {
        data: "Game.Name"
      },
      {
        data: "Score",
        // orderable: false
      }
    ],
    error: function (xhr, textStatus, error) {
      console.log("error: ", textStatus);
    }
    // "deferRender": true
  };

  $("#games").DataTable(options);
});

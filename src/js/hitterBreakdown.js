
function submitPlayerBreakdownRequest(playerName, gameName) {
  var xhttp = new XMLHttpRequest();
  xhttp.open("POST", "http://localhost:9090/getStatisticsByPostionForPlayer?playerName=" + playerName, false);
  xhttp.setRequestHeader("Content-type", "application/json");
  xhttp.send();
  var response = JSON.parse(xhttp.responseText);
  return response;
}

function seeBreakdown() {
  var playerName = document.getElementById("playerName").value;
  var displayArea = document.getElementById("displayArea");
  while (displayArea.firstChild) {
    displayArea.removeChild(displayArea.firstChild);
  }
  if (playerName != "") {
    // Forms have been successfully filled out
    var response = submitPlayerBreakdownRequest(playerName)
    if (response.Success) {
      // Create element for each position breakdown
      for (var x in response.PlayerStatsByPosition) {
        var hitStatsElement = document.createElement('ul');
        hitStatsElement.id = "hit_stats_element_" + x
        hitStatsElement.className = "hit_stats_element column";
        hitStatsElement.innerHTML = response.PlayerStatsByPosition[x].PositionName
        displayArea.appendChild(hitStatsElement);
        Object.keys(response.PlayerStatsByPosition[x]).forEach(function(key,index) {
          if (!key.includes("Position")) {
            var hitStatsElementItem = document.createElement('li');
            hitStatsElementItem.id = "hit_stats_element_" + x + "_" + index
            hitStatsElementItem.className = "hit_stats_element_item"
            hitStatsElementItem.innerHTML = key + ": " + response.PlayerStatsByPosition[x][key]
            document.getElementById("hit_stats_element_" + x).appendChild(hitStatsElementItem);
          }
        });
      }
    } else {
      var errorElement = document.createElement('div');
      errorElement.id = "get_player_breakdown_failed";
      errorElement.className = "error";
      errorElement.innerHTML = "There was an issue joining the game"
      document.getElementById("errorArea").appendChild(errorElement);
    }
  } else {
    var errorElement = document.createElement('div');
    errorElement.id = "name_or_game_missing";
    errorElement.className = "error";
    errorElement.innerHTML = "Please fill out both the name and game fields"
    document.getElementById("signInBlock").appendChild(errorElement);
  }
}

function setup() {
  document.getElementById("seeBreakdown").addEventListener("click", seeBreakdown, false);
}

window.onload = setup;

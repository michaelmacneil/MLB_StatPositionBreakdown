
function submitPlayerBreakdownRequest(playerName, gameName) {
  var xhttp = new XMLHttpRequest();
  xhttp.open("POST", "http://ec2-54-165-12-169.compute-1.amazonaws.com:9090/getStatisticsByPostionForPlayer?playerName=" + playerName, false);
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
      var playerMessage = document.createElement('div');
      playerMessage.id = "playerMessage";
      playerMessage.className = "bold text-center center";
      playerMessage.innerHTML = "Hitting stats for " + playerName;
      displayArea.appendChild(playerMessage);

      // Create element for each position breakdown
      for (var x in response.PlayerStatsByPosition) {
        var hitStatsElement = document.createElement('ul');
        hitStatsElement.id = "hit_stats_element_" + x;
        hitStatsElement.className = "hit_stats_element column";
        displayArea.appendChild(hitStatsElement);

        var hitStatsElementPosition = document.createElement('li');
        hitStatsElementPosition.id = "hit_stats_element_" + x + "_position";
        hitStatsElementPosition.className = "hit_stats_element_item bold";
        hitStatsElementPosition.innerHTML = response.PlayerStatsByPosition[x].PositionName;
        document.getElementById("hit_stats_element_" + x).appendChild(hitStatsElementPosition);

        Object.keys(response.PlayerStatsByPosition[x]).forEach(function(key,index) {
          if (!key.includes("Position")) {
            var hitStatsElementItem = document.createElement('li');
            hitStatsElementItem.id = "hit_stats_element_" + x + "_" + index
            hitStatsElementItem.className = "hit_stats_element_item"
            if (response.PlayerStatsByPosition[x][key] % 1 != 0) {
              hitStatsElementItem.innerHTML = key + ": " + response.PlayerStatsByPosition[x][key].toFixed(3)
            } else {
              hitStatsElementItem.innerHTML = key + ": " + response.PlayerStatsByPosition[x][key]
            }
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

//var host = "http://ec2-54-165-12-169.compute-1.amazonaws.com";
var host = "http://localhost";


function submitPlayerBreakdownRequest(playerName, gameName) {
  var xhttp = new XMLHttpRequest();
  xhttp.open("POST", host + ":9090/getStatisticsByPositionForPlayer?playerName=" + playerName, false);
  xhttp.setRequestHeader("Content-type", "application/json");
  xhttp.send();
  var response = JSON.parse(xhttp.responseText);
  return response;
}

function getMLBTeams() {
  var xhttp = new XMLHttpRequest();
  xhttp.open("POST", host + ":9090/getMLBTeams", false);
  xhttp.setRequestHeader("Content-type", "application/json");
  xhttp.send();
  var response = JSON.parse(xhttp.responseText);
  return response;
}

function getPlayersFromTeam(teamAbbrev) {
  var xhttp = new XMLHttpRequest();
  xhttp.open("POST", host + ":9090/getPlayersFromTeam?teamAbbrev=" + teamAbbrev, false);
  xhttp.setRequestHeader("Content-type", "application/json");
  xhttp.send();
  var response = JSON.parse(xhttp.responseText);
  return response;
}

function searchSeeBreakdown() {
  seeBreakdown(document.getElementById("playerName").value);
}

function selectSeeBreakdown() {
  seeBreakdown(document.getElementById("player_list").value);
}

function seeBreakdown(playerName) {
  //var playerName = document.getElementById("playerName").value;
  if (playerName != "" && playerName != "N/A") {
    var displayArea = document.getElementById("displayArea");
    // Forms have been successfully filled out
    var response = submitPlayerBreakdownRequest(playerName)
    if (response.Success) {
      while (displayArea.firstChild) {
        displayArea.removeChild(displayArea.firstChild);
      }

      var playerMessage = document.getElementById('playerMessage');
      playerMessage.innerHTML = "Hitting stats for " + playerName;
      var br = document.createElement('br');
      displayArea.appendChild(br);

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
    document.getElementById("errorArea").appendChild(errorElement);
  }
}

function updatePlayers() {
  var teamValue = document.getElementById("team_list").value;
  if (teamValue != "N/A") {
    var playersResponse = getPlayersFromTeam(teamValue);
    if (playersResponse.Success) {
      var playerList = document.getElementById('player_list');
      while (playerList.firstChild) {
        playerList.removeChild(playerList.firstChild);
      }
      var emptyElement = document.createElement('option');
      emptyElement.id = "player_element_empty";
      emptyElement.className = "player_element";
      emptyElement.value = "N/A";
      emptyElement.innerHTML = "Select A Player";
      playerList.appendChild(emptyElement);
      for (var x in playersResponse.Players) {
        var playerElement = document.createElement('option');
        playerElement.id = "player_element_" + x;
        playerElement.className = "player_element";
        playerElement.value = playersResponse.Players[x].PlayerName;
        playerElement.innerHTML = playersResponse.Players[x].PlayerName;
        playerList.appendChild(playerElement);
      }
    } else {
      var errorElement = document.createElement('div');
      errorElement.id = "error_reaching_backend";
      errorElement.className = "error";
      errorElement.innerHTML = "Error - Server is down";
      document.getElementById("playerSelect").appendChild(errorElement);
    }
    playerList.addEventListener("change", selectSeeBreakdown, false)
  }
}

function setup() {
  var teamsResponse = getMLBTeams();
  if (teamsResponse.Success) {
    var teamList = document.getElementById('team_list');
    for (var x in teamsResponse.Teams) {
      var teamElement = document.createElement('option');
      teamElement.id = "team_element_" + x;
      teamElement.className = "team_element";
      teamElement.value = teamsResponse.Teams[x].Abbrev;
      teamElement.innerHTML = teamsResponse.Teams[x].CityName;
      teamList.appendChild(teamElement);
    }
    teamList.addEventListener("change", updatePlayers, false)
  } else {
    var errorElement = document.createElement('div');
    errorElement.id = "error_reaching_backend";
    errorElement.className = "error";
    errorElement.innerHTML = "Error - Server is down";
    document.getElementById("teamSelect").appendChild(errorElement);
  }
  document.getElementById("seeBreakdown").addEventListener("click", searchSeeBreakdown, false);
  document.getElementById("playerName").addEventListener("keyup", function(event) {
    if (event.keyCode === 13) {
      searchSeeBreakdown();
    }
  }, false);
}

window.onload = setup;

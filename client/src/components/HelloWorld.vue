<template>
  <div class="hello">
    <div class="container-fluid">
      <div>
        <div v-if="state.mutableViewServer" class="row">
          <div class="col"></div>
          <div class="col-4 jumbotron animated jello">
            <h2>Battle Royale Chess</h2>
            <img class="animated shake" width="200px" alt="Logo" src="../assets/logo.png">

            <h3>Connect</h3>
            <br />
            <br />


            <div class="form-group">
              <label for="nickname">Nickname</label>
              <input
                type="text"
                class="form-control"
                id="nicname"
                v-model="state.nickname"
                aria-describedby="nicknameHelp"
                placeholder="Nickname for the game"
              >
              <small
                id="nicknameHelp"
                class="form-text text-muted"
              >Your name will be shown to other players.</small>
            </div>

            <div class="form-group">
              <label for="url">URL</label>
              <input
                type="text"
                class="form-control"
                id="url"
                v-model="state.url"
                aria-describedby="urlHelp"
                placeholder="URL of server"
              >
              <small
                id="urlHelp"
                class="form-text text-muted"
              >Game server url.</small>
            </div>

            <br>
            <button class="btn btn-success" v-on:click="connect">Connect</button>
          </div>
          <div class="col"></div>
        </div>
          
        <!-- /row -->
        <div v-if="state.mutableViewGames" class="row">
          <div class="col"></div>
          <div class="col-4 jumbotron">
            <h2>Battle Royale Chess</h2>
            <img class="animated bounce" alt="Logo" src="../assets/logo.png">

            <h3>Lobby</h3>
            <br />

            <button class="btn btn-success" v-on:click="createGame">New game</button>&nbsp;
            <button class="btn btn-success" v-on:click="searchAgain">Refresh</button>

            <div class="md-layout">
              <div class="md-layout-item"></div>
              <div class="md-layout-item">
                <br>
                <table class="table table-hover">
                  <tr>
                    <th md-numeric>Name</th>
                    <th>Players</th>
                    <th>Actions</th>
                  </tr>

                  <tr v-for="game in state.listOfGames" :key="game['id']">
                    <td>{{ game["title"] }}</td>
                    <td>{{ game["players"] }}</td>
                    <td>
                      <button class="btn btn-primary btn-sm" v-on:click="joinGame(game['id'])">Join</button>
                    </td>
                  </tr>
                </table>
              </div>
              <div class="md-layout-item"></div>
            </div>
          </div>
          <div class="col"></div>
        </div>

        <div v-if="state.mutableViewGame">
          <button v-if="amIAdmin() && !state.gameState.started" class="btn btn-success" v-on:click="startGame">Start game</button>&nbsp;
          <span v-for="i in getPlayersArr()" :key="1 + i.id" class="badge badge-pill badge-success" style="margin-right: 10px" v-if="i.myTurn">{{i.nick}}</span>
          <span v-for="i in getPlayersArr()" :key="2 + i.id" class="badge badge-pill badge-warning" style="margin-right: 10px" v-if="!i.myTurn">{{i.nick}}</span>
          <canvas style="outline: none;" id="renderCanvas"></canvas>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
const networkR = require("./network.js");
const stateR = require("./state.js");
const B = require("./engine.js").B;

export default {
  name: "HelloWorld",
  props: {
    msg: String
  },
  data: function() {
    return {
      NetworkManager: networkR.NetworkManager,
      state: new stateR.State()
    };
  },
  methods: {
    connect() {
      if (!this.state.nickname) {
        alert("You need a nickname");
        return;
      }
      this.NetworkManager.connect(this.state, B);
      B.NetworkManager = this.NetworkManager;
    },

    searchAgain() {
      this.NetworkManager.searchAgain();
    },

    createGame() {
      this.NetworkManager.createGame();
    },

    joinGame(id) {
      this.NetworkManager.joinGame(id);
    },

    startGame() {
      this.NetworkManager.startGame();
    },

    getPlayersArr() {
      let arr = [];
      for (let seat in this.state.gameState.players) {
        let item = {
          id:  this.state.gameState.players[seat].profile.id,
          nick: this.state.gameState.players[seat].profile.nick,
          myTurn: this.state.gameState.players[seat].myTurn
        }
        arr.push(item);
      }
      return arr;
    },

    amIAdmin() {
      if (this.state && this.state.gameState && this.state.gameState.owner) {
        return this.state.gameState.owner.profile.id === this.state.id;
      }
      return false
    },

    getBoards() {
      let boards = [];
      let i = 1;
      for (i = 1; i <= this.state.gameState.boards; i++) {
        let board = {};
        board.id = i;
        boards.push(board);
      }

      return boards;
    },

    myTurn() {
      const id = this.state.id;
      for (let seat in this.state.gameState.players) {
        if (this.state.gameState.players[seat].profile.id == id) {
          return this.state.gameState.players[seat].myTurn;
        }
      }

      return true;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
#renderCanvas {
  width: 100%;
  height: 100%;
  touch-action: none;
}
</style>



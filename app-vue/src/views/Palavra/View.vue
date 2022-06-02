<script>
import axios from "../../axiosBase";

export default {
  data() {
    return {
      palavras: [],
    };
  },

  methods: {
    loadPalavras() {
      let vue = this;
      axios
        .get("palavra")
        .then(function (response) {
          let data = response.data;
          vue.palavras = data;
        })
        .catch(function (error) {
          alert(error);
        });
    },
  },

  mounted() {
    this.loadPalavras();
  },
};
</script>

<template>
    <div class="columns is-mobile is-centered mt-5">
      <div class="card">
        <div class="field">
          <p class="control has-icons-left has-icons-right">
            <input class="input" type="text" placeholder="Pesquisar..." />
            <span class="icon is-small is-left">
              <i class="fas fa-search"></i>
            </span>
          </p>
        </div>
      </div>
    </div>

    <div class="columns is-mobile is-centered mt-5">
      <div class="card">
        <router-link to="/palavra/add" class="button is-success is-fullwidth"
          >Adicionar Palavra</router-link
        >
      </div>
    </div>

    <div
      class="columns is-mobile is-centered mt-5"
      v-for="palavra in palavras"
      :key="palavra.ID"
    >
      <div class="card">
        <div class="card-header">
          <p class="card-header-title">{{ palavra.Nome }}</p>
        </div>
        <div class="card-content">
          <div class="content">
            {{ palavra.Descricao }}
          </div>
        </div>
      </div>
    </div>
</template>

<style>
.card {
  width: 30vh;
}
</style>
<script>
import axios from "../../axiosBase";

export default {
  data() {
    return {
      id: parseInt(this.$route.params.id),
      method: this.$route.params.method,
      palavra: {
        ID: null,
        Nome: null,
        Descricao: null,
      },
    };
  },
  methods: {
    salvar() {
      let vue = this;
      axios
        .post("palavra", vue.palavra)
        .then(function (response) {
					console.log(response);
					if(response.status === 201) {
						vue.$toast.open({
							message: "Salvo!",
							type: "success",
						});
						vue.$router.push('/palavra')
					} else {
						vue.$toast.open({
							message: "Não foi possível Salvar!",
							type: "error",
						});
					}
        })
        .catch(function (error) {
          alert(error);
        });
    },
  },
};
</script>

<template>
  <template v-if="method == 'add' || method == 'edit'">
    <div class="columns is-mobile is-centered">
      <div class="column is-one-third">
        <p class="title is-3 is-spaced mt-3">
          {{ id > 0 ? "Editar" : "Cadastrar" }}
        </p>
      </div>
    </div>

    <div class="columns is-mobile is-centered">
      <div class="column is-one-third">
        <div class="field">
          <label class="label">Palavra</label>
          <div class="control has-icons-left has-icons-right">
            <input
              class="input"
              v-model="palavra.Nome"
              type="text"
              placeholder="Abacate"
            />
            <span class="icon is-small is-left">
              <i class="fa-solid fa-book-bookmark"></i>
            </span>
          </div>
        </div>
        <div class="field">
          <label class="label">Descrição</label>
          <div class="control">
            <textarea
              v-model="palavra.Descricao"
              class="textarea"
              placeholder="Uma fruta muito saborosa!!"
            ></textarea>
          </div>
        </div>
        <div class="has-text-right">
          <button @click="$router.back()" class="button is-light mr-3">
            Voltar
          </button>
          <button class="button is-success" @click="salvar()">Salvar</button>
        </div>
      </div>
    </div>
  </template>
</template>

<style>
.card {
  width: 30vh;
}
</style>
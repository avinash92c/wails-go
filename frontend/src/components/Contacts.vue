<template>
  <div id="contactsview">
    <div id="actions">
      <el-button
        type="primary"
        icon="el-icon-edit"
        circle
        @click="toggledialog"
      ></el-button>
    </div>
    <div>
      <el-table :data="allcontacts" style="width: 100%">
        <el-table-column prop="Firstname" label="Firstname" width="180">
        </el-table-column>
        <el-table-column prop="Lastname" label="Lastname" width="180">
        </el-table-column>
        <el-table-column prop="Mobile" label="Mobile"></el-table-column>
        <el-table-column prop="Office" label="Office"></el-table-column>
        <el-table-column prop="Home" label="Home"></el-table-column>
      </el-table>
      <el-dialog
        title="Modify Contact Details"
        :visible.sync="dialogVisible"
        width="30%"
      >
        <form>
          <el-input
            placeholder="Firstname"
            v-model="newcontact.Firstname"
            clearable
          ></el-input>
          <el-input
            placeholder="Lastname"
            v-model="newcontact.Lastname"
            clearable
          ></el-input>
          <el-input
            type="tel"
            placeholder="Mobile"
            v-model="newcontact.Mobile"
            clearable
          ></el-input>
          <el-input
            type="tel"
            placeholder="Home"
            v-model="newcontact.Home"
            clearable
          ></el-input>
          <el-input
            type="tel"
            placeholder="Office"
            v-model="newcontact.Office"
            clearable
          ></el-input>
        </form>
        <span slot="footer" class="dialog-footer">
          <el-button
            @click="
              toggledialog();
              clearnew();
            "
            >Cancel</el-button
          >
          <el-button
            type="primary"
            @click="
              savecontact();
              toggledialog();
            "
            >Confirm</el-button
          >
        </span>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
export default {
  name: "Contacts",
  data() {
    return {
      dialogVisible: false,
      allcontacts: [],
      newcontact: {
        Firstname: "",
        Lastname: "",
        Mobile: 0,
        Home: 0,
        Office: 0
      }
    };
  },
  mounted() {
    this.contacts();
  },
  methods: {
    toggledialog() {
      this.dialogVisible = !this.dialogVisible;
    },
    clearnew() {
      Vue.set(this.newcontact, "Firstname", "");
      Vue.set(this.newcontact, "Lastname", "");
      Vue.set(this.newcontact, "Mobile", 0);
      Vue.set(this.newcontact, "Home", 0);
      Vue.set(this.newcontact, "Office", 0);
    },
    getContact: function() {
      var self = this;
      window.backend.readContact().then(result => {});
    },
    savecontact() {
      var self = this;
      window.backend.saveContact(this.newcontact).then(result => {
        this.contacts();
      });
      this.clearnew();
    },
    contacts() {
      var self = this;
      window.backend.allContacts().then(result => {
        self.allcontacts = result;
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}

a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}

a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>

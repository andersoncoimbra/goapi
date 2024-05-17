<template>
  <v-card>
   
    <v-tabs
      v-model="tab"
      bg-color="primary"
    >
      <v-tab value="users">Produtos</v-tab>
    </v-tabs>
    

    <v-card-text>
      
      <v-tabs-window v-model="tab">
        <v-tabs-window-item value="users">         
          <TableProducts/>
        </v-tabs-window-item>

        
      </v-tabs-window>
    </v-card-text>
  </v-card>
</template>
<script>

import TableProducts  from "./components/TableProducts.vue";

  export default {
    components: {
      TableProducts,
      
    },
    data: () => ({
      tab: null,
      users: [],
      message: '',
      typemessage: ''
    })
  }
</script>
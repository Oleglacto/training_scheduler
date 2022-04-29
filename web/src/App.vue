<template>
  <BaseHeader />
  <div style="display: flex">
    <BaseSide />
    <div>
      <img alt="Vue logo" class="element-plus-logo" src="./assets/logo.png" />
      <HelloWorld :msg="msg" />
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      msg: ''
    }
  },
  created() {
    fetch('/api/v1/test').then((response) => {
      return response.json();
    })
    .then((data) => {
      this.msg = data.text
    });
  }
}
</script>

<style>
#app {
  text-align: center;
  color: var(--el-text-color-primary);
}

.element-plus-logo {
  width: 50%;
}
</style>

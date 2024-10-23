<template>
  <q-page padding>
    <div>
      <h1>用户管理</h1>
      <q-input v-model="fullname" label="输入全名" outlined />
      <q-input v-model="username" label="输入用户名" outlined />
      <q-input v-model="password" type="password" label="输入密码" outlined />
      <q-input v-model="email" label="输入邮箱" outlined />

      <q-checkbox v-model="selectedOptions.freenas" label="Freenas" />
      <q-checkbox v-model="selectedOptions.gitlab" label="GitLab" />
      <q-checkbox v-model="selectedOptions.gpu" label="GPU算力池" />

      <q-btn @click="addUser" label="添加用户" color="primary" />

      <q-input
        v-model="searchUsername"
        label="输入查询用户名"
        @keyup.enter="fetchUserInfo"
        outlined
      />
      <q-btn @click="fetchUserInfo" label="查询用户" color="secondary" />

      <div v-if="fetchedUser">
        <q-card>
          <q-card-section>
            <h2>查询结果:</h2>
            <p>全名: {{ fetchedUser.name }}</p>
            <p>用户名: {{ fetchedUser.username }}</p>
            <p>邮箱: {{ fetchedUser.email }}</p>
          </q-card-section>
        </q-card>
      </div>

      <q-list bordered>
        <q-item v-for="(user, index) in users" :key="index">
          <q-item-section
            >{{ user.fullname }} ({{ user.username }})</q-item-section
          >
          <q-item-section side>
            <q-btn
              flat
              color="red"
              icon="delete"
              @click.stop="removeUser(index)"
            />
          </q-item-section>
        </q-item>
      </q-list>

      <q-separator />

      <div v-if="!users.length">暂无用户</div>
    </div>
  </q-page>
</template>

<script>
import { ref } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const fullname = ref('');
    const username = ref('');
    const password = ref('');
    const email = ref('');
    const selectedOptions = ref({
      freenas: false,
      gitlab: false,
      gpu: false,
    });
    const users = ref([]);
    const searchUsername = ref('');
    const fetchedUser = ref('');

    const addUser = async () => {
      if (
        fullname.value.trim() &&
        username.value.trim() &&
        password.value.trim() &&
        email.value.trim()
      ) {
        const userData = {
          fullname: fullname.value.trim(),
          username: username.value.trim(),
          password: password.value.trim(),
          email: email.value.trim(),
          services: {
            freenas: selectedOptions.value.freenas,
            gitlab: selectedOptions.value.gitlab,
            gpu: selectedOptions.value.gpu,
          },
        };

        // 调用 API 创建用户
        try {
          await axios.post('YOUR_CREATE_API_ENDPOINT', userData);
          users.value.push(userData);
          clearFields();
        } catch (error) {
          console.error('添加用户失败：', error);
        }
      }
    };

    const removeUser = async (index) => {
      const user = users.value[index];
      try {
        // 调用 API 删除用户
        await axios.delete(`YOUR_DELETE_API_ENDPOINT/${user.username}`);
        users.value.splice(index, 1);
      } catch (error) {
        console.error('删除用户失败：', error);
      }
    };

    const fetchUserInfo = async () => {
      if (searchUsername.value.trim()) {
        try {
          const response = await axios.get(
            `http://localhost:8080/gitlab/users`,
            { params: { username: searchUsername.value } } // 通过 query 参数传递数据
          );
          fetchedUser.value = response.data; // 假设返回的数据格式正确
          //   console.log(response.data);
          console.log(fetchedUser.value);
        } catch (error) {
          console.error('查询用户失败：', error);
          fetchedUser.value = null; // 清空结果
        }
      }
    };

    const clearFields = () => {
      fullname.value = '';
      username.value = '';
      password.value = '';
      email.value = '';
      selectedOptions.value = {
        freenas: false,
        gitlab: false,
        gpu: false,
      };
      searchUsername.value = '';
      fetchedUser.value = null; // 清空查询结果
    };

    return {
      fullname,
      username,
      password,
      email,
      selectedOptions,
      users,
      searchUsername,
      fetchedUser,
      addUser,
      removeUser,
      fetchUserInfo,
    };
  },
};
</script>

<style scoped>
/* 您可以在此处添加自定义样式 */
</style>

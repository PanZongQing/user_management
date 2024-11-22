<template>
  <q-page padding>
    <div>
      <h1>用户管理</h1>
      <q-input v-model="full_name" label="输入全名" outlined />
      <q-input v-model="username" label="输入用户名" outlined />
      <q-input v-model="password" type="password" label="输入密码" outlined />
      <q-input v-model="email" label="输入邮箱" outlined />
      <q-input v-model="phonenumber" label="输入手机号" outlined />
      <q-input v-model="group_name" label="输入部门" outlined />

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

      <div v-if="fetchUserInfo">
        <q-card>
          <q-card-section>
            <h4>查询结果:</h4>
            <p>全名: {{ searchUsername }}</p>
            <ul>
              <li>
                GPU算力池权限：
                <q-icon :name="fetchedgitlabsuccess ? 'check' : 'cancel'" />
              </li>
              <li>
                FreeNas权限：
                <q-icon :name="fetchedfreensuccess ? 'check' : 'cancel'" />
              </li>
              <li>
                GitLab权限：
                <q-icon :name="fetchedgpusuccess ? 'check' : 'cancel'" />
              </li>
            </ul>
          </q-card-section>
        </q-card>
      </div>
      <!-- <div v-else>
        <q-card>
          <q-card-section>
            <h2>查询失败</h2>
            <p>无法获取用户信息，请检查输入是否正确</p>
          </q-card-section>
        </q-card>
      </div> -->

      <q-list bordered>
        <q-item v-for="(user, index) in users" :key="index">
          <q-item-section
            >{{ user.full_name }} ({{ user.username }})</q-item-section
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
    const full_name = ref('');
    const username = ref('');
    const password = ref('');
    const email = ref('');
    const phonenumber = ref('');
    const group_name = ref('');
    const selectedOptions = ref({
      freenas: false,
      gitlab: false,
      gpu: false,
    });
    const users = ref([]);
    const searchUsername = ref('');
    const fetchedgitlabUser = ref({});
    const fetchedfreenasUser = ref({});
    const fetchedgpuUser = ref({});
    const fetchedgitlabsuccess = ref({});
    const fetchedfreensuccess = ref({});
    const fetchedgpusuccess = ref({});

    const addUser = async () => {
      if (
        full_name.value.trim() &&
        username.value.trim() &&
        password.value.trim() &&
        email.value.trim() &&
        phonenumber.value.trim() &&
        group_name.value.trim()
      ) {
        const freenasuserData = {
          full_name: full_name.value.trim(),
          username: username.value.trim(),
          password: password.value.trim(),
          email: email.value.trim(),
          group_name: group_name.value.trim(),
        };
        const gitlabuserData = {
          full_name: full_name.value.trim(),
          username: username.value.trim(),
          password: password.value.trim(),
          email: email.value.trim(),
        };
        const gpuuserData = {
          username: username.value.trim(),
          phonenumber: phonenumber.value.trim(),
        };

        // 遍历选择的服务，将其对应的 API 地址加入列表
        if (selectedOptions.value.freenas) {
          try {
            await axios.post(
              'http://localhost:8080/freenas/users',
              freenasuserData
            );
            users.value.push(freenasuserData);
            clearFields();
          } catch (error) {
            console.error('添加用户失败：', error);
          }
        }
        if (selectedOptions.value.gitlab) {
          try {
            await axios.post(
              'http://localhost:8080/gitlab/users',
              gitlabuserData
            );
            users.value.push(gitlabuserData);
            clearFields();
          } catch (error) {
            console.error('添加用户失败：', error);
          }
        }
        if (selectedOptions.value.gpu) {
          // 调用 GPU 算力池 API 创建用户
          try {
            await axios.post('http://localhost:8080/gpu/users', gitlabuserData);
            users.value.push(gpuuserData);
            clearFields();
          } catch (error) {
            console.error('添加用户失败：', error);
          }
        }

        // 调用 API 创建用户
        // try {
        //   await axios.post('YOUR_CREATE_API_ENDPOINT', userData);
        //   users.value.push(userData);
        //   clearFields();
        // } catch (error) {
        //   console.error('添加用户失败：', error);
        // }
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
        const usernames = searchUsername.value.trim(); // 避免多次调用 trim
        const successStatuses = [];
        const endpoints = [
          {
            url: 'http://localhost:8080/gitlab/users',
            result: fetchedgitlabUser,
          },
          {
            url: 'http://localhost:8080/freenas/users',
            result: fetchedfreenasUser,
          },
          {
            url: 'http://localhost:8080/gpu/users',
            result: fetchedgpuUser,
          },
        ];

        for (const endpoint of endpoints) {
          try {
            const response = await axios.get(endpoint.url, {
              params: { username: usernames },
            });
            endpoint.result.value = response.data.data; // 假设返回的数据格式正确
            successStatuses.push(true);
            endpoint.success = true;
            console.log(endpoint.result.value);
          } catch (error) {
            console.error(`查询 ${endpoint.url} 用户失败：`, error);
            endpoint.result.value = null; // 清空结果
            successStatuses.push(false);
            endpoint.success = false;
          }
        }
        updateFrontendStatus(endpoints); // 更新前端状态
        // 更新权限状态
        console.log(successStatuses);
        fetchedgitlabsuccess.value = successStatuses[0];
        fetchedfreensuccess.value = successStatuses[1];
        fetchedgpusuccess.value = successStatuses[2];
        console.log(fetchedgitlabsuccess.value);
        console.log(fetchedfreensuccess.value);
        console.log(fetchedgpusuccess.value);
      }
    };

    //更新前端状态函数
    const updateFrontendStatus = (endpoints) => {
      endpoints.forEach((endpoint) => {
        if (endpoint.success) {
          // 在前端展示勾选状态，例如渲染一个勾
          console.log(`${endpoint.url} 查询成功，显示勾选`);
        } else {
          // 查询失败的状态可以根据需求处理
          console.log(`${endpoint.url} 查询失败，隐藏勾选`);
        }
      });
    };

    const clearFields = () => {
      full_name.value = '';
      username.value = '';
      password.value = '';
      email.value = '';
      selectedOptions.value = {
        freenas: false,
        gitlab: false,
        gpu: false,
      };
      searchUsername.value = '';
      // 清空查询结果
    };

    return {
      full_name,
      username,
      password,
      email,
      selectedOptions,
      users,
      searchUsername,
      addUser,
      removeUser,
      fetchUserInfo,
      group_name,
      fetchedgitlabsuccess,
      fetchedfreensuccess,
      fetchedgpusuccess,
      phonenumber,
    };
  },
};
</script>

<style scoped>
/* 您可以在此处添加自定义样式 */
</style>

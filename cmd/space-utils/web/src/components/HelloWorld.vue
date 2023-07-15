<template>
  <el-table :data="machinesInfo" style="width: 100%">
    <el-table-column prop="machine" label="机器" width="180" />
    <el-table-column prop="state" label="状态" width="180" />
    <el-table-column prop="completedSize" label="完成" />
    <el-table-column prop="commitmentSize" label="目标" />
    <el-table-column prop="percent" label="完成率" />
  </el-table>
</template>

<script lang="ts">
import { getMachineConnection, getMachineInfo, Status } from '~/api/api';

import { defineComponent, getCurrentInstance, onMounted, ref } from 'vue';

export default defineComponent({
  setup() {
    let machinesInfo = ref<Status[]>([]);

    async function fetMachineInfo() {
      let urlMap = await getMachineConnection();
      for (const [key, value] of urlMap.entries()) {
        const machineInfo = await getMachineInfo(key);
        machinesInfo.value.push(machineInfo)
      }
    }

    fetMachineInfo();
    return {
      machinesInfo
    };
  }
});
</script>
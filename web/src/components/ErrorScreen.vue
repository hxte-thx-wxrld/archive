<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter()

const goBack = () => {
    router.back()
}
</script>

<template>
    <div>

        <a href="#" @click="goBack">Go back</a>
        <h1>Oh! The specified item doesn't exist</h1>
    </div>
</template>
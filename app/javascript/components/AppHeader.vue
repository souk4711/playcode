<script setup lang="ts">
import { LanguageIcon } from '@heroicons/vue/24/solid'
import { GitHubIcon } from '@/components/icons'
</script>

<template>
  <div class="navbar bg-base-100 shadow">
    <div class="flex-none px-2">
      <a class="text-lg font-bold uppercase">Playcode</a>
    </div>
    <div class="flex-1 flex justify-end px-2">
      <div class="dropdown dropdown-end">
        <div class="tooltip tooltip-bottom" data-tip="Change Language">
          <label tabindex="0" class="btn btn-ghost">
            <LanguageIcon class="h-5 w-5 md:h-6 md:w-6" />
          </label>
        </div>
        <ul
          tabindex="0"
          class="dropdown-content menu w-52 p-2 mt-3 bg-base-100 shadow"
        >
          <li><a>English</a></li>
          <li><a>简体中文</a></li>
        </ul>
      </div>
      <div class="tooltip tooltip-bottom" data-tip="View on GitHub">
        <a
          tabindex="0"
          class="btn btn-ghost"
          href="https://github.com/souk4711/playcode"
          target="_blank"
          rel="noopener noreferrer"
        >
          <GitHubIcon class="h-5 w-5 md:h-6 md:w-6" />
        </a>
      </div>
    </div>
  </div>
</template>

// uno.config.ts
import presetAttributify from '@unocss/preset-attributify'
import presetUno from '@unocss/preset-uno'
import transformerDirectives from '@unocss/transformer-directives'
import { defineConfig } from 'unocss'


export default defineConfig({
  // ...UnoCSS options
  rules: [
    [/^bg-color-#(.*)$/, ([, color]) => ({ 'background-color': `#${color}` })],
    ['m-100', { margin: '1.25rem' }],
    ['/^w-(.*)px$/', ([, num]) => ({'width': `${num}px`})],

  ],
  presets: [
    presetUno(),
    presetAttributify({
      /* preset options */
    })
  ],

  shortcuts: {
    'act': '',
    'btn': 'py-2 px-4 font-semibold rounded-lg shadow-md',
    'btn-green': 'text-white bg-green-500 hover:bg-green-700',
  },
  transformers: [
    transformerDirectives(),
  ],
})
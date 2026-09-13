import path from "path"
import react, { reactCompilerPreset } from '@vitejs/plugin-react'
import babel from '@rolldown/plugin-babel'
import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'
import dotenv from 'dotenv'
import fs from "fs"

// https://vite.dev/config/
export default defineConfig(({mode})=>{
	const envDir = process.env.CONFIG_PATH || "."
	const targetEnvFile = path.resolve(import.meta.dirname, envDir, process.env.ENV_FILE || '.env')

	let parsedEnv:any = {}
	if (fs.existsSync(targetEnvFile)) {
		const envConfig = fs.readFileSync(targetEnvFile)
		parsedEnv = dotenv.parse(envConfig)
	}

	const envInjection = Object.keys(parsedEnv).reduce((acc:any, key) => {
		acc[`import.meta.env.${key}`] = JSON.stringify(parsedEnv[key])
		return acc
	}, {})

	return {
		plugins: [
			react(),
			babel({ presets: [reactCompilerPreset()] }),
			tailwindcss(),
		],
		resolve: {
			alias: {
				"@": path.resolve(import.meta.dirname, "./src"),
			},
		},
		envDir,
		mode: process.env.MODE || mode,
		server: {
			port: parsedEnv.PORT ? Number(parsedEnv.PORT) : 3000,
		},
		define: envInjection,
	}
})

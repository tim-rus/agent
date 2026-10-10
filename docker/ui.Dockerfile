# builder
FROM oven/bun:1.4-alpine AS builder
WORKDIR /app
ARG CONFIG_PATH
ARG MODE
COPY ./ui/package.json ./ui/bun.lockb* ./
RUN --mount=type=cache,target=/root/.bun/install/cache \
	bun install --frozen-lockfile
COPY ./ui .
RUN --mount=type=bind,source=${CONFIG_PATH},target=/tmp/config cp /tmp/config/.env.${MODE} ./.env.${MODE}
RUN CONFIG_PATH=. MODE=${MODE} bun run build-only

# runner
FROM sebp/lighttpd:latest
COPY --from=builder /app/dist /app/www
COPY ./ui/lighttpd.conf /etc/lighttpd/lighttpd.conf
EXPOSE 80
CMD ["start.sh"]

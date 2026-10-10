# builder
FROM oven/bun:1.4-alpine AS builder
WORKDIR /app
ARG CONFIG_PATH
ARG MODE
COPY ./admin/package.json ./admin/bun.lockb* ./
RUN --mount=type=cache,target=/root/.bun/install/cache \
	bun install --frozen-lockfile
COPY ./admin .
RUN --mount=type=bind,source=${CONFIG_PATH},target=/tmp/config cp /tmp/config/.env.${MODE} ./.env.${MODE}
RUN CONFIG_PATH=. MODE=${MODE} bun run build-only

# runner
FROM sebp/lighttpd:latest
COPY --from=builder /app/dist /app/www
COPY ./admin/lighttpd.conf /etc/lighttpd/lighttpd.conf
EXPOSE 80
CMD ["start.sh"]

export function joinUrl(base: string, ...segments: string[]): string {
    const url = new URL(base);
    const cleanPathname = [url.pathname, ...segments]
        .flatMap(part => part.split('/'))
        .filter(Boolean)
        .join('/');
    url.pathname = `/${cleanPathname}`;
    return url.toString();
}
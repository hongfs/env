async function fetch(r) {
    var h;

    let headers = {
        'X-Ratelimit-Host': r.headersIn.host,
        'X-Ratelimit-Uri': r.uri,
        'X-Ratelimit-IP': r.remoteAddress,
    };

    for (h in r.headersIn) {
        headers['X-Header-' + h] = r.headersIn[h];
    }

    let result = await ngx.fetch('https://ratelimit.hongfs.cn/', {
        method: r.method,
        headers: headers,
        verify: false,
    });

    // 200: 可以访问
    // 429: 限流
    // 500：服务器错误

    const status = result.status;

    if(status === 429) {
        return r.return(status);
    }

    const backend_reply = await  r.subrequest('/backend' + r.uri, {
        method: r.method,
        body: r.requestBody,
    });

    return r.return(backend_reply.status, backend_reply.responseBody);
}

export default {
    fetch,
};

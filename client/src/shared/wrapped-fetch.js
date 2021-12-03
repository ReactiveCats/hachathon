/**
 * Обертка над fetch
 * @description Добавляет заголовок `Content-Type` к запросу
 * и выбрасывает ошибку если сервер вернул статус вне диапазона от 200 до 300
 * @param {RequestInfo} input адрес запроса в формате `string | URL` или объект `Request`
 * @param {RequestInit} init дополнительные параметры запроса
 * @returns JSON полученный с сервера в виде объекта
 */
export async function wrappedFetch(input, init) {
  const { body = null, headers = {} } = init;

  if (headers instanceof Headers) {
    headers.append('Content-Type', 'application/json');
  } else {
    headers['Content-Type'] = 'application/json';
  }

  const response = await fetch(input, {
    ...init,
    body: body !== null ? JSON.stringify(body) : undefined,
    headers,
  });

  const json = await response.json();

  if (!response.ok) {
    if ('error' in body) {
      throw json.error;
    }

    throw json;
  }

  return json;
}

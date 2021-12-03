/**
 * Обертка над fetch
 * @description Добавляет заголовок `Content-Type` к запросу
 * и выбрасывает ошибку если сервер вернул статус вне диапазона от 200 до 300
 * @param {RequestInfo} input адрес запроса в формате `string | URL` или объект `Request`
 * @param {RequestInit} init дополнительные параметры запроса
 * @returns JSON полученный с сервера в виде объекта
 */
export async function wrappedFetch(input, init) {
  const { headers = {} } = init;

  if (headers instanceof Headers) {
    headers.append('Content-Type', 'application/json');
  } else {
    headers['Content-Type'] = 'application/json';
  }

  const response = await fetch(input, {
    ...init,
    headers,
  });

  if (!response.ok) {
    if ('error' in response) {
      throw response.error;
    }

    throw response;
  }

  return response.json();
}

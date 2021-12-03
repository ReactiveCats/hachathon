/**
 * Вызов эффектов после вызова `dispatch` события
 * @param {Function} dispatch `dispatch` который получаем из `useReducer`
 * @param {Object} effects Эффекты
 * @returns `dispatch` который при вызове вызывает эффекты привязанные к событию
 */
export function makeDispatchWithEffects(dispatch, effects) {
  return (action) => {
    dispatch(action);

    if (typeof effects[action.type] === 'function') {
      effects[action.type](action.data, dispatch);
    }
  };
}

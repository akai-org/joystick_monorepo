export const onOpenHandler = data => {
  const globalId = JSON.parse(localStorage.getItem('global_id'))
  data.target.send(JSON.stringify({ global_id: globalId }))
}

import { onMounted, onUnmounted } from 'vue'
import { events } from '../interfaces/InterfacesEvents'

export function useInterfaceLifecycleEssentials (context, wantedOrientation) {
  onMounted(() => context.emit(events.onSetOrientation, wantedOrientation))
  onUnmounted(() => context.emit(events.onSetOrientation, null))
}

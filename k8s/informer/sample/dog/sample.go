package dog

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

type WatchDog struct {
	lw      *cache.ListWatch
	objType runtime.Object
	h       cache.ResourceEventHandler

	reflector *cache.Reflector
	fifo      *cache.DeltaFIFO
	store     cache.Store
}

func NewWatchDog(lw *cache.ListWatch, objType runtime.Object, h cache.ResourceEventHandler) *WatchDog {

	store := cache.NewStore(cache.MetaNamespaceKeyFunc)

	fifo := cache.NewDeltaFIFOWithOptions(cache.DeltaFIFOOptions{
		KeyFunction:  cache.MetaNamespaceKeyFunc,
		KnownObjects: store,
	})

	reflector := cache.NewReflector(lw, &v1.Pod{}, fifo, 0)

	return &WatchDog{lw: lw, objType: objType, h: h, reflector: reflector, fifo: fifo, store: store}
}

func (w *WatchDog) Run() {
	ch := make(chan struct{})
	go func() {
		w.reflector.Run(ch)
	}()

	for {
		w.fifo.Pop(func(obj interface{}) error {
			for _, delta := range obj.(cache.Deltas) {

				switch delta.Type {
				case cache.Added, cache.Sync:
					w.store.Add(delta.Object)
					w.h.OnAdd(delta.Object)
				case cache.Updated:
					if old, exist, err := w.store.Get(delta.Object); err == nil && exist {
						w.store.Update(delta.Object)
						w.h.OnUpdate(old, delta.Object)
					}

				case cache.Deleted:
					w.store.Delete(delta.Object)
					w.h.OnDelete(delta.Object) // 相当于回调
				}
			}

			return nil
		})
	}
}

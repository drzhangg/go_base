package main

import (
	"crypto/rand"
	"fmt"
	cnins "github.com/containernetworking/plugins/pkg/ns"
	"golang.org/x/sys/unix"
	"os"
	"path"
	"runtime"
	"sync"
)

func main() {
	p,err := newNS("/",0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}

func newNS(baseDir string, pid uint32) (nsPath string, err error) {
	b := make([]byte, 16)

	_, err = rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random netns name: %w", err)
	}

	// Create the directory for mounting network namespaces
	// This needs to be a shared mountpoint in case it is mounted in to
	// other namespaces (containers)
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", err
	}

	// create an empty file at the mount point and fail if it already exists
	nsName := fmt.Sprintf("cni-%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	nsPath = path.Join(baseDir, nsName)
	mountPointFd, err := os.OpenFile(nsPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return "", err
	}
	mountPointFd.Close()

	defer func() {
		// Ensure the mount point is cleaned up on errors
		if err != nil {
			os.RemoveAll(nsPath)
		}
	}()

	if pid != 0 {
		procNsPath := getNetNSPathFromPID(pid)
		// bind mount the netns onto the mount point. This causes the namespace
		// to persist, even when there are no threads in the ns.
		if err = unix.Mount(procNsPath, nsPath, "none", unix.MS_BIND, ""); err != nil {
			return "", fmt.Errorf("failed to bind mount ns src: %v at %s: %w", procNsPath, nsPath, err)
		}
		return nsPath, nil
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// do namespace work in a dedicated goroutine, so that we can safely
	// Lock/Unlock OSThread without upsetting the lock/unlock state of
	// the caller of this function
	go (func() {
		defer wg.Done()
		runtime.LockOSThread()
		// Don't unlock. By not unlocking, golang will kill the OS thread when the
		// goroutine is done (for go1.10+)


		var origNS cnins.NetNS
		origNS, err = cnins.GetNS(getCurrentThreadNetNSPath())
		if err != nil {
			return
		}
		defer origNS.Close()

		// create a new netns on the current thread
		err = unix.Unshare(unix.CLONE_NEWNET)
		if err != nil {
			return
		}

		// Put this thread back to the orig ns, since it might get reused (pre go1.10)
		defer origNS.Set()

		// bind mount the netns from the current thread (from /proc) onto the
		// mount point. This causes the namespace to persist, even when there
		// are no threads in the ns.
		err = unix.Mount(getCurrentThreadNetNSPath(), nsPath, "none", unix.MS_BIND, "")
		if err != nil {
			err = fmt.Errorf("failed to bind mount ns at %s: %w", nsPath, err)
		}
	})()
	wg.Wait()

	if err != nil {
		return "", fmt.Errorf("failed to create namespace: %w", err)
	}

	return nsPath, nil
}

func getNetNSPathFromPID(pid uint32) string {
	return fmt.Sprintf("/proc/%d/ns/net", pid)
}

func getCurrentThreadNetNSPath() string {
	// /proc/self/ns/net returns the namespace of the main thread, not
	// of whatever thread this goroutine is running on.  Make sure we
	// use the thread's net namespace since the thread is switching around
	return fmt.Sprintf("/proc/%d/task/%d/ns/net", os.Getpid(), unix.Gettid())
}


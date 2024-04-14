// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package fileproperties

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/foundation/collections"
)

const GUIDIStorageItemExtraProperties string = "c54361b2-54cd-432b-bdbc-4b19c4b470d7"
const SignatureIStorageItemExtraProperties string = "{c54361b2-54cd-432b-bdbc-4b19c4b470d7}"

type IStorageItemExtraProperties struct {
	ole.IInspectable
}

type IStorageItemExtraPropertiesVtbl struct {
	ole.IInspectableVtbl

	RetrievePropertiesAsync            uintptr
	SavePropertiesAsync                uintptr
	SavePropertiesAsyncOverloadDefault uintptr
}

func (v *IStorageItemExtraProperties) VTable() *IStorageItemExtraPropertiesVtbl {
	return (*IStorageItemExtraPropertiesVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IStorageItemExtraProperties) RetrievePropertiesAsync(propertiesToRetrieve *collections.IIterable) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().RetrievePropertiesAsync,
		uintptr(unsafe.Pointer(v)),                    // this
		uintptr(unsafe.Pointer(propertiesToRetrieve)), // in collections.IIterable
		uintptr(unsafe.Pointer(&out)),                 // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageItemExtraProperties) SavePropertiesAsync(propertiesToSave *collections.IIterable) (*foundation.IAsyncAction, error) {
	var out *foundation.IAsyncAction
	hr, _, _ := syscall.SyscallN(
		v.VTable().SavePropertiesAsync,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(unsafe.Pointer(propertiesToSave)), // in collections.IIterable
		uintptr(unsafe.Pointer(&out)),             // out foundation.IAsyncAction
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageItemExtraProperties) SavePropertiesAsyncOverloadDefault() (*foundation.IAsyncAction, error) {
	var out *foundation.IAsyncAction
	hr, _, _ := syscall.SyscallN(
		v.VTable().SavePropertiesAsyncOverloadDefault,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncAction
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
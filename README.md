```shell

2021-06-22 17:24:54.239 507-507/? I/VClientImpl: Xposed is enabled.
2021-06-22 17:24:54.349 507-507/? D/ExposedHelper: IXposedMod's classLoader : dalvik.system.PathClassLoader[DexPathList[[zip file "/data/app/io.va.exposed-n6SZOhQ23UPE0Q2YgPqJPw==/base.apk"],nativeLibraryDirectories=[/data/app/io.va.exposed-n6SZOhQ23UPE0Q2YgPqJPw==/lib/arm, /system/fake-libs, /data/app/io.va.exposed-n6SZOhQ23UPE0Q2YgPqJPw==/base.apk!/lib/armeabi-v7a, /system/lib, /system/product/lib]]]
2021-06-22 17:24:58.155 718-718/? I/VClientImpl: Xposed is enabled.
2021-06-22 17:24:59.997 718-718/? I/ExposedBridge: io.liuliu.music is not a Xposed module, do not init epic.force
2021-06-22 17:25:00.004 718-718/? D/ExposedHelper: IXposedMod's classLoader : dalvik.system.PathClassLoader[DexPathList[[zip file "/data/app/io.va.exposed-n6SZOhQ23UPE0Q2YgPqJPw==/base.apk"],nativeLibraryDirectories=[/data/app/io.va.exposed-n6SZOhQ23UPE0Q2YgPqJPw==/lib/arm, /system/fake-libs, /data/app/io.va.exposed-n6SZOhQ23UPE0Q2YgPqJPw==/base.apk!/lib/armeabi-v7a, /system/lib, /system/product/lib]]]
2021-06-22 17:25:00.583 816-816/? I/VClientImpl: Xposed is enabled.
2021-06-22 17:25:00.923 718-1010/? I/Xposed: 调用getSubscriberId获取了imsi
2021-06-22 17:25:00.926 718-1010/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$2.afterHookedMethod(SherLockMonitor.java:75)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    android.telephony.TelephonyManager.getSubscriberId(TelephonyManager.java:3695)
    com.alibaba.one.android.inner.DeviceInfoCapturer.doCommandForString(Unknown Source:847)
2021-06-22 17:25:00.941 718-1010/? I/Xposed: 调用getMacAddress()获取了mac地址
2021-06-22 17:25:00.941 718-1010/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$3.afterHookedMethod(SherLockMonitor.java:93)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    com.alibaba.one.android.inner.DeviceInfoCapturer.c(Unknown Source:28)
    com.alibaba.one.android.inner.DeviceInfoCapturer.doCommandForString(Unknown Source:497)
2021-06-22 17:25:00.956 718-1010/? I/Xposed: 调用getHardwareAddress()获取了mac地址
2021-06-22 17:25:00.959 718-1010/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$4.afterHookedMethod(SherLockMonitor.java:111)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    com.alibaba.one.android.inner.DeviceInfoCapturer.b(Unknown Source:11)
    com.alibaba.one.android.inner.DeviceInfoCapturer.doCommandForString(Unknown Source:368)
2021-06-22 17:25:01.008 718-1010/? I/Xposed: 调用getSubscriberId获取了imsi
2021-06-22 17:25:01.012 718-1010/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$2.afterHookedMethod(SherLockMonitor.java:75)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    android.telephony.TelephonyManager.getSubscriberId(TelephonyManager.java:3695)
    com.ta.utdid2.a.a.d.getImsi(SourceFile:87)
    com.ta.utdid2.device.b.a(SourceFile:50)
    com.ta.utdid2.device.b.b(SourceFile:72)
    com.ta.utdid2.device.UTDevice.a(SourceFile:50)
    com.ta.utdid2.device.UTDevice.getUtdid(SourceFile:14)
    com.ut.device.UTDevice.getUtdid(SourceFile:19)
    java.lang.reflect.Method.invoke(Native Method)
    com.alibaba.one.android.inner.DeviceInfoCapturer.d(Unknown Source:29)
    com.alibaba.one.android.inner.DeviceInfoCapturer.doCommandForString(Unknown Source:194)
2021-06-22 17:25:01.706 718-1010/? I/Xposed: 调用getSubscriberId获取了imsi
2021-06-22 17:25:01.710 718-1010/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$2.afterHookedMethod(SherLockMonitor.java:75)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    android.telephony.TelephonyManager.getSubscriberId(TelephonyManager.java:3695)
    com.alibaba.one.android.inner.DeviceInfoCapturer.doCommandForString(Unknown Source:847)
2021-06-22 17:25:01.722 718-1010/? I/Xposed: 调用getMacAddress()获取了mac地址
2021-06-22 17:25:01.722 718-1010/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$3.afterHookedMethod(SherLockMonitor.java:93)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    com.alibaba.one.android.inner.DeviceInfoCapturer.c(Unknown Source:28)
    com.alibaba.one.android.inner.DeviceInfoCapturer.doCommandForString(Unknown Source:497)
2021-06-22 17:25:01.731 718-1010/? I/Xposed: 调用getHardwareAddress()获取了mac地址
2021-06-22 17:25:01.735 718-1010/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$4.afterHookedMethod(SherLockMonitor.java:111)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    com.alibaba.one.android.inner.DeviceInfoCapturer.b(Unknown Source:11)
    com.alibaba.one.android.inner.DeviceInfoCapturer.doCommandForString(Unknown Source:368)
2021-06-22 17:25:01.938 718-1010/? I/Xposed: 调用getLastKnownLocation获取了GPS地址
2021-06-22 17:25:01.942 718-1010/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$5.afterHookedMethod(SherLockMonitor.java:131)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    com.alibaba.one.sdk.j.a(Unknown Source:98)
    com.alibaba.one.android.inner.DeviceInfoCapturerFull.doCommandForString(Unknown Source:680)
2021-06-22 17:25:01.944 718-1010/? I/Xposed: 调用getLastKnownLocation获取了GPS地址
2021-06-22 17:25:01.950 718-1010/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$5.afterHookedMethod(SherLockMonitor.java:131)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    com.alibaba.one.sdk.j.a(Unknown Source:98)
    com.alibaba.one.android.inner.DeviceInfoCapturerFull.doCommandForString(Unknown Source:656)
2021-06-22 17:25:03.220 816-816/? I/ExposedBridge: io.liuliu.music is not a Xposed module, do not init epic.force
2021-06-22 17:25:03.229 816-816/? D/ExposedHelper: IXposedMod's classLoader : dalvik.system.PathClassLoader[DexPathList[[zip file "/data/app/io.va.exposed-n6SZOhQ23UPE0Q2YgPqJPw==/base.apk"],nativeLibraryDirectories=[/data/app/io.va.exposed-n6SZOhQ23UPE0Q2YgPqJPw==/lib/arm, /system/fake-libs, /data/app/io.va.exposed-n6SZOhQ23UPE0Q2YgPqJPw==/base.apk!/lib/armeabi-v7a, /system/lib, /system/product/lib]]]
2021-06-22 17:25:06.129 718-999/? I/Xposed: 调用getHardwareAddress()获取了mac地址
2021-06-22 17:25:06.132 718-999/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$4.afterHookedMethod(SherLockMonitor.java:111)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
2021-06-22 17:25:06.136 718-999/? I/Xposed: 调用getHardwareAddress()获取了mac地址
2021-06-22 17:25:06.139 718-999/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$4.afterHookedMethod(SherLockMonitor.java:111)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
2021-06-22 17:25:06.453 718-1262/? I/Xposed: 调用getHardwareAddress()获取了mac地址
2021-06-22 17:25:06.460 718-1262/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$4.afterHookedMethod(SherLockMonitor.java:111)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    com.taobao.wireless.security.adapter.datacollection.f.b(Unknown Source:17)
    com.taobao.wireless.security.adapter.datacollection.f.b(Unknown Source:23)
    com.taobao.wireless.security.adapter.datacollection.DeviceInfoCapturer.doCommandForString(Unknown Source:124)
2021-06-22 17:25:12.252 718-1291/? I/Xposed: 调用getSubscriberId获取了imsi
2021-06-22 17:25:12.256 718-1291/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$2.afterHookedMethod(SherLockMonitor.java:75)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    android.telephony.TelephonyManager.getSubscriberId(TelephonyManager.java:3695)
    com.alipay.security.mobile.module.b.b.b(Unknown Source:23)
    com.alipay.apmobilesecuritysdk.d.c.a(Unknown Source:17)
    com.alipay.apmobilesecuritysdk.d.e.c(Unknown Source:28)
    com.alipay.apmobilesecuritysdk.d.e.a(Unknown Source:7)
    com.alipay.apmobilesecuritysdk.d.e.b(Unknown Source:3)
    com.alipay.apmobilesecuritysdk.a.a.a(Unknown Source:71)
    com.alipay.apmobilesecuritysdk.face.APSecuritySdk$1.run(Unknown Source:13)
    com.alipay.apmobilesecuritysdk.f.c.run(Unknown Source:39)
    java.lang.Thread.run(Thread.java:919)
2021-06-22 17:25:12.261 718-1291/? I/Xposed: 调用getMacAddress()获取了mac地址
2021-06-22 17:25:12.262 718-1291/? I/Xposed: dalvik.system.VMStack.getThreadStackTrace(Native Method)
    java.lang.Thread.getStackTrace(Thread.java:1720)
    com.android.sherlock.SherLockMonitor.getMethodStack(SherLockMonitor.java:139)
    com.android.sherlock.SherLockMonitor.access$000(SherLockMonitor.java:33)
    com.android.sherlock.SherLockMonitor$3.afterHookedMethod(SherLockMonitor.java:93)
    de.robv.android.xposed.DexposedBridge.handleHookedArtMethod(DexposedBridge.java:265)
    me.weishu.epic.art.entry.Entry.onHookObject(Entry.java:69)
    me.weishu.epic.art.entry.Entry.referenceBridge(Entry.java:186)
    com.alipay.security.mobile.module.b.b.k(Unknown Source:24)
    com.alipay.apmobilesecuritysdk.d.c.a(Unknown Source:21)
    com.alipay.apmobilesecuritysdk.d.e.c(Unknown Source:28)
    com.alipay.apmobilesecuritysdk.d.e.a(Unknown Source:7)
    com.alipay.apmobilesecuritysdk.d.e.b(Unknown Source:3)
    com.alipay.apmobilesecuritysdk.a.a.a(Unknown Source:71)
    com.alipay.apmobilesecuritysdk.face.APSecuritySdk$1.run(Unknown Source:13)
    com.alipay.apmobilesecuritysdk.f.c.run(Unknown Source:39)
    java.lang.Thread.run(Thread.java:919)



```









```
直接调用
HttpApplicationDelegate
ImageLoadApplication
CrashApplicationDelegate
SocketApplicationDelegate
DataApplicationDelegate
UpdateApplicationDelegate
WidgetApplicationDelegate
InviteApplicationDelegate

WebApplicationDelegate

需要隐私协议
LoginApplicationDelegate
IDataTrackDelegate
PushApplicationDelegate
AliApplicationDelegate
WebApplicationDelegate
```










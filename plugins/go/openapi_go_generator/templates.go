// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

func templates() map[string]string {
	return map[string]string{ 
        "app.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQoKcGFja2FnZSB7ey5SZW5kZXJlci5QYWNrYWdlfX0KCmltcG9ydCAoCgkiZW5jb2RpbmcvanNvbiIKCSJmbXQiCgkibG9nIgoJIm5ldC9odHRwIgoJInN0cmNvbnYiCgoJImdpdGh1Yi5jb20vZ29yaWxsYS9tdXgiCikKCmZ1bmMgaW50VmFsdWUocyBzdHJpbmcpICh2IGludDY0KSB7Cgl2LCBfID0gc3RyY29udi5QYXJzZUludChzLCAxMCwgNjQpCglyZXR1cm4gdgp9CgovLyBoYW5kbGVycwp7e3JhbmdlIC5SZW5kZXJlci5NZXRob2RzfX0KCi8vIHt7LkRlc2NyaXB0aW9ufX0KZnVuYyB7ey5IYW5kbGVyTmFtZX19KHcgaHR0cC5SZXNwb25zZVdyaXRlciwgciAqaHR0cC5SZXF1ZXN0KSB7Cglsb2cuUHJpbnRmKCJ7ey5OYW1lfX0iKTsKCXZhciBlcnIgZXJyb3IKCXt7aWYgaGFzUGFyYW1ldGVycyAufX0KCS8vIGluc3RhbnRpYXRlIHRoZSBwYXJhbWV0ZXJzIHN0cnVjdHVyZQoJdmFyIHBhcmFtZXRlcnMge3suUGFyYW1ldGVyc1R5cGVOYW1lfX0KCQoJe3tpZiBlcSAuTWV0aG9kICJQT1NUIn19Ly8gZGVzZXJpYWxpemUgcmVxdWVzdCBmcm9tIHBvc3QgZGF0YQoJZGVjb2RlciA6PSBqc29uLk5ld0RlY29kZXIoci5Cb2R5KQoJZXJyID0gZGVjb2Rlci5EZWNvZGUoJnBhcmFtZXRlcnMue3tib2R5UGFyYW1ldGVyRmllbGROYW1lIC59fSkKCWlmIGVyciAhPSBuaWwgewoJCWZtdC5GcHJpbnRmKHcsICJFUlJPUjogJXYiLCBlcnIpCgkJcmV0dXJuCgl9Cglsb2cuUHJpbnRmKCJQQVJBTUVURVJTICUrdiIsIHBhcmFtZXRlcnMpCgl7e2VuZH19CgkKCS8vIGdldCByZXF1ZXN0IGZpZWxkcyBpbiBwYXRoIGFuZCBxdWVyeSBwYXJhbWV0ZXJzCgl7e2lmIGhhc1BhdGhQYXJhbWV0ZXJzIC59fQoJdmFycyA6PSBtdXguVmFycyhyKQoJe3tlbmR9fQoJe3tpZiBoYXNGb3JtUGFyYW1ldGVycyAufX0KCXIuUGFyc2VGb3JtKCkKCXt7ZW5kfX0KCQoJe3tyYW5nZSAuUGFyYW1ldGVyc1R5cGUuRmllbGRzfX0JCgkKCXt7aWYgZXEgLlBvc2l0aW9uICJwYXRoIn19CglpZiB2YWx1ZSwgb2sgOj0gdmFyc1sie3suSlNPTk5hbWV9fSJdOyBvayB7CgkJcGFyYW1ldGVycy57ey5OYW1lfX0gPSBpbnRWYWx1ZSh2YWx1ZSkKCX0KCXt7ZW5kfX0KCQoJe3tpZiBlcSAuUG9zaXRpb24gImZvcm1kYXRhIn19CglpZiBsZW4oci5Gb3JtWyJ7ey5KU09OTmFtZX19Il0pID4gMCB7CgkJcGFyYW1ldGVycy57ey5OYW1lfX0gPSBpbnRWYWx1ZShyLkZvcm1bInt7LkpTT05OYW1lfX0iXVswXSkKCX0KCXt7ZW5kfX0KCQoJe3tlbmR9fQoJe3tlbmR9fQoJCgl7e2lmIGhhc1Jlc3BvbnNlcyAufX0JCgkvLyBpbnN0YW50aWF0ZSB0aGUgcmVzcG9uc2VzIHN0cnVjdHVyZQoJdmFyIHJlc3BvbnNlcyB7ey5SZXNwb25zZXNUeXBlTmFtZX19Cgl7e2VuZH19CgoJLy8gY2FsbCB0aGUgcHJvY2Vzc29yCQoJe3tpZiBoYXNQYXJhbWV0ZXJzIC59fQoJe3tpZiBoYXNSZXNwb25zZXMgLn19CgllcnIgPSB7ey5Qcm9jZXNzb3JOYW1lfX0oJnBhcmFtZXRlcnMsICZyZXNwb25zZXMpCgl7e2Vsc2V9fQoJZXJyID0ge3suUHJvY2Vzc29yTmFtZX19KCZwYXJhbWV0ZXJzKQoJe3tlbmR9fQoJe3tlbHNlfX0KCXt7aWYgaGFzUmVzcG9uc2VzIC59fQoJZXJyID0ge3suUHJvY2Vzc29yTmFtZX19KCZyZXNwb25zZXMpCgl7e2Vsc2V9fQoJZXJyID0ge3suUHJvY2Vzc29yTmFtZX19KCkKCXt7ZW5kfX0KCXt7ZW5kfX0KCQoJaWYgZXJyID09IG5pbCB7Cgl7eyBpZiBoYXNSZXNwb25zZXMgLn19CgkJe3sgaWYgLlJlc3BvbnNlc1R5cGUgfCBoYXNGaWVsZE5hbWVkT0sgfX0KCQkvLyB3cml0ZSB0aGUgbm9ybWFsIHJlc3BvbnNlCQkKCQllbmNvZGVyIDo9IGpzb24uTmV3RW5jb2Rlcih3KQoJCWVuY29kZXIuRW5jb2RlKHJlc3BvbnNlcy5PSykKCQl7eyBlbmQgfX0KCXt7ZW5kfX0KCX0gZWxzZSB7CgkJLy8gd3JpdGUgYW4gZXJyb3IgcmVzcG9uc2UuLi4KCQlmbXQuRnByaW50Zih3LCAiRVJST1I6ICV2IiwgZXJyKQoJfQp9Cnt7ZW5kfX0KCmZ1bmMgaW5pdCAoKSB7Cgl2YXIgcm91dGVyID0gbXV4Lk5ld1JvdXRlcigpe3tyYW5nZSAuUmVuZGVyZXIuTWV0aG9kc319Cglyb3V0ZXIuSGFuZGxlRnVuYygie3suUGF0aH19Iiwge3suSGFuZGxlck5hbWV9fSkuTWV0aG9kcygie3suTWV0aG9kfX0iKXt7ZW5kfX0KICAgIGh0dHAuSGFuZGxlKCIvIiwgcm91dGVyKQogICAgaW5pdGlhbGl6ZV9zZXJ2aWNlKCkKfQoKZnVuYyBtYWluICgpIHsKCWh0dHAuTGlzdGVuQW5kU2VydmUoIjo4MDgwIiwgbmlsKQp9Cg==",
        "app.yaml": "YXBwbGljYXRpb246IHt7LlJlbmRlcmVyLk5hbWV9fQp2ZXJzaW9uOiAxCnJ1bnRpbWU6IGdvCmFwaV92ZXJzaW9uOiBnbzEKaGFuZGxlcnM6Ci0gdXJsOiAvLioKICBzY3JpcHQ6IF9nb19hcHAKLSB1cmw6IC8KICBzdGF0aWNfZGlyOiBzdGF0aWM=",
        "client.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQoKcGFja2FnZSB7ey5SZW5kZXJlci5QYWNrYWdlfX0KCmltcG9ydCAoCiAgImJ5dGVzIgogICJlbmNvZGluZy9qc29uIgogICJmbXQiCiAgIm5ldC9odHRwIgogICJzdHJpbmdzIgopCiAgCmNvbnN0IHNlcnZpY2UgPSAiaHR0cDovL2xvY2FsaG9zdDo4MDgwIgoKdHlwZSBDbGllbnQgc3RydWN0IHsKCVNlcnZpY2Ugc3RyaW5nCn0gCgpmdW5jIE5ld0NsaWVudChzZXJ2aWNlIHN0cmluZykgKkNsaWVudCB7CgljbGllbnQgOj0gJkNsaWVudHt9CgljbGllbnQuU2VydmljZSA9IHNlcnZpY2UKCXJldHVybiBjbGllbnQKfQoKe3tyYW5nZSAuUmVuZGVyZXIuTWV0aG9kc319Cnt7aWYgZXEgLlJlc3VsdFR5cGVOYW1lICIifX0KZnVuYyAoY2xpZW50ICpDbGllbnQpIHt7LkNsaWVudE5hbWV9fSh7e3BhcmFtZXRlckxpc3QgLn19KSAoZXJyIGVycm9yKSB7Cnt7ZWxzZX19CmZ1bmMgKGNsaWVudCAqQ2xpZW50KSB7ey5DbGllbnROYW1lfX0oe3twYXJhbWV0ZXJMaXN0IC59fSkgKHJlc3VsdCAqe3suUmVzdWx0VHlwZU5hbWV9fSwgZXJyIGVycm9yKSB7Cnt7ZW5kfX0KCXBhdGggOj0gY2xpZW50LlNlcnZpY2UgKyAie3suUGF0aH19IgoJCgl7e2lmIGhhc1BhcmFtZXRlcnMgLn19Cgl7e3JhbmdlIC5QYXJhbWV0ZXJzVHlwZS5GaWVsZHN9fQkKCXt7aWYgZXEgLlBvc2l0aW9uICJwYXRoIn19CglwYXRoID0gc3RyaW5ncy5SZXBsYWNlKHBhdGgsICJ7IiArICJ7ey5KU09OTmFtZX19IiArICJ9IiwgZm10LlNwcmludGYoIiV2Iiwge3suSlNPTk5hbWV9fSksIDEpCgl7e2VuZH19Cgl7e2VuZH19Cgl7e2VuZH19CgkKCXt7aWYgZXEgLk1ldGhvZCAiUE9TVCJ9fQoJYm9keSA6PSBuZXcoYnl0ZXMuQnVmZmVyKQoJanNvbi5OZXdFbmNvZGVyKGJvZHkpLkVuY29kZSh7e2JvZHlQYXJhbWV0ZXJOYW1lIC59fSkKCXJlcSwgZXJyIDo9IGh0dHAuTmV3UmVxdWVzdCgie3suTWV0aG9kfX0iLCBwYXRoLCBib2R5KQoJe3tlbHNlfX0KCXJlcSwgZXJyIDo9IGh0dHAuTmV3UmVxdWVzdCgie3suTWV0aG9kfX0iLCBwYXRoLCBuaWwpCgl7e2VuZH19CgkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybgoJfQoJcmVzcCwgZXJyIDo9IGh0dHAuRGVmYXVsdENsaWVudC5EbyhyZXEpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gCgl9CglkZWZlciByZXNwLkJvZHkuQ2xvc2UoKQoJe3tpZiBuZSAuUmVzdWx0VHlwZU5hbWUgIiJ9fQoJZGVjb2RlciA6PSBqc29uLk5ld0RlY29kZXIocmVzcC5Cb2R5KQoJcmVzdWx0ID0gJnt7LlJlc3VsdFR5cGVOYW1lfX17fQoJZGVjb2Rlci5EZWNvZGUocmVzdWx0KQoJe3tlbmR9fQoJcmV0dXJuICAKfQp7e2VuZH19Cg==",
        "service.go": "Ly8gUExFQVNFIENPTVBMRVRFIFRIRSBGVU5DVElPTlMgSU4gVEhJUyBGSUxFIFRPIElNUExFTUVOVCBZT1VSIFNFUlZJQ0UKCnBhY2thZ2Uge3suUmVuZGVyZXIuUGFja2FnZX19CgppbXBvcnQgKAoJImxvZyIKKQoKZnVuYyBpbml0aWFsaXplX3NlcnZpY2UoKSB7Cn0KCnt7cmFuZ2UgLlJlbmRlcmVyLk1ldGhvZHN9fQp7e2lmIGhhc1BhcmFtZXRlcnMgLn19Cnt7aWYgaGFzUmVzcG9uc2VzIC59fQpmdW5jIHt7LlByb2Nlc3Nvck5hbWV9fShwYXJhbWV0ZXJzICp7ey5QYXJhbWV0ZXJzVHlwZU5hbWV9fSwgcmVzcG9uc2VzICp7ey5SZXNwb25zZXNUeXBlTmFtZX19KSAoZXJyIGVycm9yKSB7Cnt7ZWxzZX19CmZ1bmMge3suUHJvY2Vzc29yTmFtZX19KHBhcmFtZXRlcnMgKnt7LlBhcmFtZXRlcnNUeXBlTmFtZX19KSAoZXJyIGVycm9yKSB7Cnt7ZW5kfX0Ke3tlbHNlfX0Ke3tpZiBoYXNSZXNwb25zZXMgLn19CmZ1bmMge3suUHJvY2Vzc29yTmFtZX19KHJlc3BvbnNlcyAqe3suUmVzcG9uc2VzVHlwZU5hbWV9fSkgKGVyciBlcnJvcikgewp7e2Vsc2V9fQpmdW5jIHt7LlByb2Nlc3Nvck5hbWV9fSgpIChlcnIgZXJyb3IpIHsKe3tlbmR9fQp7e2VuZH19CQoJe3tpZiBoYXNQYXJhbWV0ZXJzIC59fQoJbG9nLlByaW50Zigie3suTmFtZX19ICUrdiIsIHBhcmFtZXRlcnMpCgl7e2Vsc2V9fQoJbG9nLlByaW50Zigie3suTmFtZX19IikKCXt7ZW5kfX0KCS8vIElNUExFTUVOVCBZT1VSIFNFUlZJQ0UgSEVSRQoJcmV0dXJuIG5pbAp9Cnt7ZW5kfX0=",
        "types.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQoKcGFja2FnZSB7ey5SZW5kZXJlci5QYWNrYWdlfX0KCi8vIHR5cGVzCnt7cmFuZ2UgLlJlbmRlcmVyLlR5cGVzfX0KdHlwZSB7ey5OYW1lfX0gc3RydWN0IHsge3tyYW5nZSAuRmllbGRzfX0KICB7ey5OYW1lfX0ge3suVHlwZX19e3tpZiBuZSAuSlNPTk5hbWUgIiJ9fSBganNvbjoie3suSlNPTk5hbWV9fSJge3tlbmR9fXt7ZW5kfX0KfQp7e2VuZH19",
    }
}
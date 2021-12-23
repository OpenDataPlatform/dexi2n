# dexi2n

A small tool to generate encoded clientID for Dex `oauthclient` resource creation

To dynamically generate a client definition for DEX server, the official method is to use the gRPC API.

One undocumented alternate method would be to directly define the corresponding resources in the API server, using some manifest like the following:

```yaml
apiVersion: dex.coreos.com/v1
kind: OAuth2Client
metadata:
  name: mv4gc3lqnrss2ylqodf7fhheqqrcgji
  namespace: <TheNamespaceOfDexServer>
id: example-app
name: 'Example App'
public: false
redirectURIs:
- 'http://127.0.0.1:5555/callback'
secret: ZXhhbXBsZS1hcHAtc2VjcmV0
```

This approach has several risks and limitations:

- The Dex server must be configured to use 'kubernetes' as his storage layer.
- As this is undocumented, this resource format may change with Dex version.
- Currently, Dex is accessing this storage on each request. So, there is no need to reload or perform any action for this new resource to be taken in account. Maybe this behavior may change in the future.

So, if you can accept these constraints, you may want to use this method. 

But, there is another trick: The resource name must be the clientID encoded in a specific way. This simple tool is provided to help for this: 

```bash
$ dexi2n example-app
mv4gc3lqnrss2ylqodf7fhheqqrcgji
```

Of course, there is no guaranty this encoding will remain the same in the future. Use this undocumented trick at your own risk!.


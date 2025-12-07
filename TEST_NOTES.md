# L2VPN EVPN Profile Test Notes

## Test Results
- **17.15 Router (10.81.239.57)**: ✅ PASS - Feature works correctly
- **17.15 Switch (10.81.239.60)**: ❌ FAIL - TLS/RESTCONF connection issue

## Switch Test Environment Issue
The test Cat9K switch has a RESTCONF/TLS configuration issue that affects ALL L2VPN/EVPN features:
- SSL alert number 80 (internal error) when connecting via HTTPS
- This is NOT a feature limitation - the feature works when configured via CLI
- Manual configuration confirmed working on switch

## Troubleshooting Attempted
1. Proxy bypass (added devices to NO_PROXY in .env)
2. Different TLS versions (1.2, 1.3)
3. Direct curl and openssl testing
4. All resulted in: error:0A000438:SSL routines::tlsv1 alert internal error

## Conclusion
- Feature supports both routers and switches (IOS-XE 17.15.1+)
- Test failure is environmental, not a platform limitation
- Switch needs RESTCONF/HTTPS configuration investigation

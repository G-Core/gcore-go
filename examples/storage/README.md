# Gcore Storage SDK Examples

This directory contains comprehensive examples demonstrating how to use the Gcore Storage SDK for Go. The examples cover both S3-compatible object storage and SFTP file transfer storage types.

## Directory Structure

The examples are organized into separate directories, each containing self-contained code:

```
examples/storage/
├── basic/        # Basic storage operations example
├── s3buckets/    # S3 bucket management example
├── credentials/  # Credential management example
└── README.md
```

## Available Examples

### 1. Basic Operations (`basic/`)

Demonstrates fundamental storage operations:
- Creating S3 and SFTP storage instances
- Listing existing storages
- Retrieving storage details
- Updating storage configuration (expiration, server alias)
- Deleting storage instances

**Run the example:**
```bash
cd examples/storage/basic
go run main.go
```

**Key Features Demonstrated:**
- Storage creation with different types and locations
- Storage lifecycle management
- Error handling

### 2. S3 Bucket Management (`s3buckets/`)

Comprehensive S3 bucket operations:
- Creating S3 storage instances
- Creating and managing S3 buckets
- Setting bucket lifecycle policies
- Configuring CORS (Cross-Origin Resource Sharing)
- Managing bucket access policies
- File upload, download, and deletion operations
- Deleting buckets and cleanup

**Run the example:**
```bash
cd examples/storage/s3buckets
go run main.go
```

**Key Features Demonstrated:**
- Multiple bucket creation and management
- Lifecycle policy configuration (object expiration)
- CORS policy setup for web applications
- Bucket access policy management (public/private access)
- Proper resource cleanup

### 3. Credentials Management (`credentials/`)

Advanced credential management operations:
- S3 access key regeneration
- SFTP password management (generate, set custom, remove)
- SSH key reset operations
- Credential security best practices

**Run the example:**
```bash
cd examples/storage/credentials
go run main.go
```

**Key Features Demonstrated:**
- S3 access and secret key regeneration
- SFTP password lifecycle management
- Disabling/enabling password authentication
- SSH key management operations

## Storage Types

### S3-Compatible Storage

Gcore's S3-compatible storage provides:
- **Object Storage**: Store and retrieve files via S3 API
- **Bucket Management**: Create, configure, and manage buckets
- **Access Control**: Set bucket policies and CORS configurations
- **Lifecycle Policies**: Automatic object expiration and cleanup
- **API Compatibility**: Works with existing S3 SDKs and tools

**Use Cases:**
- Web application file storage
- Backup and archival
- Static website hosting
- Content distribution

### SFTP Storage

Gcore's SFTP storage provides:
- **File Transfer**: Standard SFTP protocol for file operations
- **Authentication**: Password and SSH key-based authentication
- **Directory Structure**: Traditional filesystem hierarchy
- **Secure Transfer**: Encrypted file transfers

**Use Cases:**
- Legacy application integration
- Secure file transfers
- Automated backup systems
- FTP replacement

## Geographic Locations

The examples demonstrate storage creation across different geographic locations. Choose the location closest to your users for optimal performance. Available locations can be retrieved using the Storage API.

## Code Structure

Each example follows a consistent pattern with self-contained functions:

- **Main Function**: Orchestrates the example flow by calling individual operation functions
- **Operation Functions**: Each function handles a complete operation (create, list, update, delete, etc.)
- **Inline Implementation**: All logic is contained within each function, making examples easy to follow
- **Context Usage**: Context is initialized inline for single-use cases, or as a variable for multiple uses
- **Error Handling**: Each function includes appropriate error handling and user feedback

## Troubleshooting

### Common Issues

1. **Storage Creation Failures**
   - Check if location is valid and available
   - Verify storage name is unique and follows naming rules

2. **Bucket Operation Errors**
   - Ensure storage is S3-compatible (not SFTP)
   - Wait for storage to be fully provisioned before bucket operations

3. **SSH Key Management**
   - SSH keys must be created/uploaded before linking to SFTP storage
   - Ensure proper SSH key format

### Getting Help

- Check the [Gcore API Documentation](https://apidocs.gcore.com/)
- Review error messages for specific guidance
- Ensure you're using the latest SDK version

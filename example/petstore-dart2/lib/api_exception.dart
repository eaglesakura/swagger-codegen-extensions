part of client_petstore.api;

class ApiException implements Exception {
  int code = 0;
  String message = null;
  Exception innerException = null;
  StackTrace stackTrace = null;
  Response response = null;

  ApiException(this.code, this.message);

  ApiException.fromResponse(this.response);

  ApiException.withInner(this.code, this.message, this.innerException, this.stackTrace);

  String toString() {
    if (message == null) return "ApiException";

    if (innerException == null) {
      return "ApiException $code: $message";
    }

    return "ApiException $code: $message (Inner exception: ${innerException})\n\n" +
        stackTrace.toString();
  }
}

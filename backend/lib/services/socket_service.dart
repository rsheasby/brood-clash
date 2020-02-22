import 'dart:io';

class SocketService {
  final List<WebSocket> _sockets = [];

  /// Remove all closed sockets from the list of registered sockets.
  void clean() {
    var i = 0;
    while (i < _sockets.length) {
      final socket = _sockets[i];
      if (socket.readyState == WebSocket.closed) {
        _sockets.removeAt(i);
        print("Socket removed");
        socket.close();
      } else {
        i += 1;
      }
    }
  }

  void register(WebSocket socket) {
    _sockets.add(socket);
  }

  // TODO: Once the question model exists we shouldn't use dynamic anymore.
  void broadcast(dynamic questions) {
    clean();

    for (var socket in _sockets) {
      socket.add(questions);
    }
  }

  void close() {
    for (var socket in _sockets) {
      socket.close();
    }
  }
}

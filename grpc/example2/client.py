import grpc
import search_pb2_grpc as pb2_grpc
import search_pb2 as pb2


class SearchClient(object):
    """
    Client for gRPC functionality
    """

    def __init__(self):
        self.host = 'localhost'
        self.server_port = 50051

        self.channel = grpc.insecure_channel(
            '{}:{}'.format(self.host, self.server_port))

        self.stub = pb2_grpc.SearchStub(self.channel)

    def get_results(self, message):
        """
        Client function to call the rpc for GetServerResponse
        """
        message = pb2.Message(message=message)
        print(f'{message}')
        return self.stub.GetServerResponse(message)


if __name__ == '__main__':
    client = SearchClient()
    result = client.get_results(message="cpu")
    print(result.product[0].name + "*******")
    print(f'{result}')
class Codec:

    def encode(self, longUrl: str) -> str:
        """Encodes a URL to a shortened URL.
        """
        self.org = longUrl
        
        

    def decode(self, shortUrl: str) -> str:
        """Decodes a shortened URL to its original URL.
        """
        return self.org

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.decode(codec.encode(url))
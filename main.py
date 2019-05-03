# import image handle from pillow
from PIL import Image
# import collections package
import collections
# import json package
import json
# import os package
import os


DEFAULT_ARGS = (
    "config.json", os.path.dirname(os.path.abspath(__file__)))

DEFAULT_KEYS = (
    "CONFIG","DIRECTORY",)

DEFAULT = collections.namedtuple(
    *("DEFAULT", ' '.join(DEFAULT_KEYS)))(*DEFAULT_ARGS)


class Cartography:

    def __init__(self, config=None):            
        with open(config) as f:
            print(f)


if __name__ == "__main__":

    configuration = os.path.join(DEFAULT.DIRECTORY, DEFAULT.CONFIG)

    cartography = Cartography(
            config=configuration)

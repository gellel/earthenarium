# import image handle from pillow
from PIL import Image
# import collections package
import collections
# import json package
import json
# import os package
import os
# import regexp package
import re
# import system package
import sys

DIR = os.path.dirname(
    os.path.abspath(__file__))

DEFAULT_KEYS = (
    "DIRECTORY", "PROJECTION",)

DEFAULT_PARAMETERS = (
    os.path.join(DIR, "image"), "equirectangular")

DEFAULT = collections.namedtuple(
    *("DEFAULT", ' '.join(DEFAULT_KEYS)))(*DEFAULT_PARAMETERS)


class Cartography(object):
    """Cartography holds projection information
    read from a cartographic file"""

    @property
    def DATUM(self):
        return self.config.datum

    @property
    def DIMENSION(self):
        return tuple(self.config.dimensions)

    @property
    def DPI(self):
        return self.config.dpi

    @property
    def EQUATOR(self):
        return (self.MAX_NORTH - self.MAX_SOUTH)

    @property
    def HEIGHT(self):
        return self.config.height

    @property
    def KEYWORDS(self):
        return self.config.keywords

    @property
    def KILOMETERS_PER_PIXEL(self):
        return self.config.scale

    @property
    def MILES_PER_PIXEL(self):
        return (self.config.scale / 1.609)
    
    @property
    def MAX_EAST(self):
        return self.config.east

    @property
    def MAX_NORTH(self):
        return self.config.north

    @property
    def MAX_SOUTH(self):
        return self.config.south

    @property
    def MAX_WEST(self):
        return self.config.west

    @property
    def MERIDIAN(self):
        return (self.MAX_WEST - self.MAX_EAST)

    @property
    def WIDTH(self):
        return self.config.width

    def latitude(self, y):
        return (y * self.config.scale)

    def longitude(self, x):
        return (x * self.config.scale)


    def __init__(self, config=None): 
        try: 
            with open(config) as f:
                config = json.load(f)
        except:
            sys.exit(1)

        self.config = collections.namedtuple(
            *("config", ' '.join(config.keys())))(*config.values())


if __name__ == "__main__":

    configuration = os.path.join(
        DEFAULT.DIRECTORY, DEFAULT.PROJECTION, "config.json")

    cartography = Cartography(
            config=configuration)

    print(cartography.longitude(10))
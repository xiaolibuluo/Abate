
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded184_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch184_1(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def traced184_2(value: int) -> int:
    return value + 83


def decorate184_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate184_2
def doubled184_2(value: int) -> int:
    return value + value
def project184_3(values):
    select = lambda item: item + 90
    return [select(item) for item in values if item != 0]
def expand184_4(values):
    for item in values:
        if item > 97:
            yield item * 60


def collect184_4(values):
    return list(expand184_4(values))
def traced184_5(value: int) -> int:
    return value + 7


def decorate184_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate184_5
def doubled184_5(value: int) -> int:
    return value + value
def traced184_6(value: int) -> int:
    return value + 14


def decorate184_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate184_6
def doubled184_6(value: int) -> int:
    return value + value
class Slot184_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{21}"

    @property
    def doubled(self) -> int:
        return self.size * 2

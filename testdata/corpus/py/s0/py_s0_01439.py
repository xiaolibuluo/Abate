
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded651_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced651_1(value: int) -> int:
    return value + 36


def decorate651_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate651_1
def doubled651_1(value: int) -> int:
    return value + value
def expand651_2(values):
    for item in values:
        if item > 43:
            yield item * 13


def collect651_2(values):
    return list(expand651_2(values))
class Slot651_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{50}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch651_4(gate):
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

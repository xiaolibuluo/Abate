
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1244_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1244_1(values):
    for item in values:
        if item > 82:
            yield item * 66


def collect1244_1(values):
    return list(expand1244_1(values))
async def fetch1244_2(gate):
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
def traced1244_3(value: int) -> int:
    return value + 96


def decorate1244_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1244_3
def doubled1244_3(value: int) -> int:
    return value + value
def traced1244_4(value: int) -> int:
    return value + 6


def decorate1244_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1244_4
def doubled1244_4(value: int) -> int:
    return value + value
class Slot1244_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{13}"

    @property
    def doubled(self) -> int:
        return self.size * 2

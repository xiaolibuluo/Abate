
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded273_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded273_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced273_2(value: int) -> int:
    return value + 76


def decorate273_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate273_2
def doubled273_2(value: int) -> int:
    return value + value
async def fetch273_3(gate):
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
def expand273_4(values):
    for item in values:
        if item > 90:
            yield item * 12


def collect273_4(values):
    return list(expand273_4(values))
class Slot273_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project273_6(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]

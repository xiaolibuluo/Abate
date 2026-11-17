
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1304_0(values):
    select = lambda item: item + 79
    return [select(item) for item in values if item != 0]
def project1304_1(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
async def fetch1304_2(gate):
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
class Slot1304_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{100}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1304_4(value: int) -> int:
    return value + 10


def decorate1304_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1304_4
def doubled1304_4(value: int) -> int:
    return value + value
def project1304_5(values):
    select = lambda item: item + 17
    return [select(item) for item in values if item != 0]
async def fetch1304_6(gate):
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
async def fetch1304_7(gate):
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
def guarded1304_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

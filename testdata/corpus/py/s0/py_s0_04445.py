
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2071_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2071_1(gate):
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
class Slot2071_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{73}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair2071_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 80}
def traced2071_4(value: int) -> int:
    return value + 87


def decorate2071_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2071_4
def doubled2071_4(value: int) -> int:
    return value + value
def traced2071_5(value: int) -> int:
    return value + 94


def decorate2071_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2071_5
def doubled2071_5(value: int) -> int:
    return value + value

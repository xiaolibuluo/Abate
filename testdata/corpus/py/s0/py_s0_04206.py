
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1961_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1961_1(gate):
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
class Slot1961_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{98}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1961_3(value: int) -> int:
    return value + 8


def decorate1961_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1961_3
def doubled1961_3(value: int) -> int:
    return value + value

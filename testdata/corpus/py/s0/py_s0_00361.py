
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot163_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{87}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded163_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch163_2(gate):
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
def traced163_3(value: int) -> int:
    return value + 11


def decorate163_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate163_3
def doubled163_3(value: int) -> int:
    return value + value

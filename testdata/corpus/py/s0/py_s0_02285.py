
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1042_0(value: int) -> int:
    return value + 68


def decorate1042_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1042_0
def doubled1042_0(value: int) -> int:
    return value + value
class Slot1042_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{75}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1042_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1042_3(gate):
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
def traced1042_4(value: int) -> int:
    return value + 96


def decorate1042_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1042_4
def doubled1042_4(value: int) -> int:
    return value + value

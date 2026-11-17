
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1959_0(value: int) -> int:
    return value + 58


def decorate1959_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1959_0
def doubled1959_0(value: int) -> int:
    return value + value
class Slot1959_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1959_2(gate):
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


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1721_0(gate):
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
def traced1721_1(value: int) -> int:
    return value + 75


def decorate1721_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1721_1
def doubled1721_1(value: int) -> int:
    return value + value
class Slot1721_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{82}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1721_3(value: int) -> int:
    return value + 89


def decorate1721_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1721_3
def doubled1721_3(value: int) -> int:
    return value + value

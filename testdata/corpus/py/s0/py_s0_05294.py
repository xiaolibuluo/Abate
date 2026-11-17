
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2452_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot2452_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{72}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2452_2(value: int) -> int:
    return value + 79


def decorate2452_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2452_2
def doubled2452_2(value: int) -> int:
    return value + value
def project2452_3(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
async def fetch2452_4(gate):
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

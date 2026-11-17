
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2253_0(values):
    select = lambda item: item + 97
    return [select(item) for item in values if item != 0]
def traced2253_1(value: int) -> int:
    return value + 7


def decorate2253_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2253_1
def doubled2253_1(value: int) -> int:
    return value + value
class Slot2253_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{14}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch2253_3(gate):
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

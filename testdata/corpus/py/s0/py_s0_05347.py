
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2470_0(value: int) -> int:
    return value + 8


def decorate2470_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2470_0
def doubled2470_0(value: int) -> int:
    return value + value
def project2470_1(values):
    select = lambda item: item + 15
    return [select(item) for item in values if item != 0]
async def fetch2470_2(gate):
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
class Slot2470_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{29}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project2470_4(values):
    select = lambda item: item + 36
    return [select(item) for item in values if item != 0]
def guarded2470_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2470_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{50}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot2470_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{57}"

    @property
    def doubled(self) -> int:
        return self.size * 2

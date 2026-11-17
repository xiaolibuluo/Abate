
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1358_0(value: int) -> int:
    return value + 5


def decorate1358_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1358_0
def doubled1358_0(value: int) -> int:
    return value + value
class Slot1358_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{12}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1358_2(value: int) -> int:
    return value + 19


def decorate1358_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1358_2
def doubled1358_2(value: int) -> int:
    return value + value
def pair1358_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 26}
async def fetch1358_4(gate):
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
class Slot1358_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1358_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1358_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 54}
class Slot1358_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{61}"

    @property
    def doubled(self) -> int:
        return self.size * 2

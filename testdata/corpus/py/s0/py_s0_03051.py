
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1402_0(gate):
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
def project1402_1(values):
    select = lambda item: item + 99
    return [select(item) for item in values if item != 0]
def guarded1402_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1402_3(value: int) -> int:
    return value + 16


def decorate1402_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1402_3
def doubled1402_3(value: int) -> int:
    return value + value
class Slot1402_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{23}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair1402_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 30}
def traced1402_6(value: int) -> int:
    return value + 37


def decorate1402_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1402_6
def doubled1402_6(value: int) -> int:
    return value + value
class Slot1402_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{44}"

    @property
    def doubled(self) -> int:
        return self.size * 2

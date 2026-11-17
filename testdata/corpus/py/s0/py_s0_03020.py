
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1386_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{78}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1386_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1386_2(value: int) -> int:
    return value + 92


def decorate1386_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1386_2
def doubled1386_2(value: int) -> int:
    return value + value
async def fetch1386_3(gate):
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
class Slot1386_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{9}"

    @property
    def doubled(self) -> int:
        return self.size * 2

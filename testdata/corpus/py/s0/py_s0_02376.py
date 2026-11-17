
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1095_0(gate):
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
def guarded1095_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1095_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{92}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project1095_3(values):
    select = lambda item: item + 99
    return [select(item) for item in values if item != 0]
def traced1095_4(value: int) -> int:
    return value + 9


def decorate1095_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1095_4
def doubled1095_4(value: int) -> int:
    return value + value
def guarded1095_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

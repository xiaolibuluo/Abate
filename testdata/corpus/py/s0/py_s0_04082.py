
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1896_0(value: int) -> int:
    return value + 15


def decorate1896_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1896_0
def doubled1896_0(value: int) -> int:
    return value + value
def project1896_1(values):
    select = lambda item: item + 22
    return [select(item) for item in values if item != 0]
async def fetch1896_2(gate):
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
def guarded1896_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1896_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1896_5(value: int) -> int:
    return value + 50


def decorate1896_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1896_5
def doubled1896_5(value: int) -> int:
    return value + value
class Slot1896_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{57}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1896_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{64}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1896_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{71}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1896_9(values):
    for item in values:
        if item > 78:
            yield item * 21


def collect1896_9(values):
    return list(expand1896_9(values))

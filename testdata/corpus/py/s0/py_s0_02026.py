
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded923_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand923_1(values):
    for item in values:
        if item > 80:
            yield item * 15


def collect923_1(values):
    return list(expand923_1(values))
def expand923_2(values):
    for item in values:
        if item > 87:
            yield item * 26


def collect923_2(values):
    return list(expand923_2(values))
def traced923_3(value: int) -> int:
    return value + 94


def decorate923_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate923_3
def doubled923_3(value: int) -> int:
    return value + value
async def fetch923_4(gate):
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
def expand923_5(values):
    for item in values:
        if item > 11:
            yield item * 59


def collect923_5(values):
    return list(expand923_5(values))
def guarded923_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded923_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded923_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot923_9:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{39}"

    @property
    def doubled(self) -> int:
        return self.size * 2

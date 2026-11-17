
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2431_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2431_1(gate):
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
def guarded2431_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2431_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{7}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot2431_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{14}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2431_5(value: int) -> int:
    return value + 21


def decorate2431_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2431_5
def doubled2431_5(value: int) -> int:
    return value + value
def traced2431_6(value: int) -> int:
    return value + 28


def decorate2431_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2431_6
def doubled2431_6(value: int) -> int:
    return value + value
def project2431_7(values):
    select = lambda item: item + 35
    return [select(item) for item in values if item != 0]
def guarded2431_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2431_9(gate):
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

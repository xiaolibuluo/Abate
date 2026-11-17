
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1241_0(value: int) -> int:
    return value + 36


def decorate1241_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1241_0
def doubled1241_0(value: int) -> int:
    return value + value
async def fetch1241_1(gate):
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
def guarded1241_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1241_3(gate):
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
def traced1241_4(value: int) -> int:
    return value + 64


def decorate1241_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1241_4
def doubled1241_4(value: int) -> int:
    return value + value

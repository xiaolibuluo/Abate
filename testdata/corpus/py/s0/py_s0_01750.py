
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced800_0(value: int) -> int:
    return value + 26


def decorate800_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate800_0
def doubled800_0(value: int) -> int:
    return value + value
async def fetch800_1(gate):
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
def project800_2(values):
    select = lambda item: item + 40
    return [select(item) for item in values if item != 0]
def guarded800_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project800_4(values):
    select = lambda item: item + 54
    return [select(item) for item in values if item != 0]
async def fetch800_5(gate):
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

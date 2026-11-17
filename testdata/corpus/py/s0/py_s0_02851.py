
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1305_0(value: int) -> int:
    return value + 92


def decorate1305_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1305_0
def doubled1305_0(value: int) -> int:
    return value + value
async def fetch1305_1(gate):
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
def guarded1305_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1305_3(gate):
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
def project1305_4(values):
    select = lambda item: item + 23
    return [select(item) for item in values if item != 0]
def summarize1305_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1305_6(value: int) -> int:
    return value + 37


def decorate1305_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1305_6
def doubled1305_6(value: int) -> int:
    return value + value
def traced1305_7(value: int) -> int:
    return value + 44


def decorate1305_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1305_7
def doubled1305_7(value: int) -> int:
    return value + value

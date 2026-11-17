
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1532_0(gate):
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
async def fetch1532_1(gate):
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
async def fetch1532_2(gate):
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
def traced1532_3(value: int) -> int:
    return value + 57


def decorate1532_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1532_3
def doubled1532_3(value: int) -> int:
    return value + value
def traced1532_4(value: int) -> int:
    return value + 64


def decorate1532_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1532_4
def doubled1532_4(value: int) -> int:
    return value + value

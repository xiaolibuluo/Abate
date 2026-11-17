
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch949_0(gate):
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
def traced949_1(value: int) -> int:
    return value + 30


def decorate949_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate949_1
def doubled949_1(value: int) -> int:
    return value + value
def traced949_2(value: int) -> int:
    return value + 37


def decorate949_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate949_2
def doubled949_2(value: int) -> int:
    return value + value
async def fetch949_3(gate):
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
async def fetch949_4(gate):
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

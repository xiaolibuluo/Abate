
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2513_0(gate):
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
def project2513_1(values):
    select = lambda item: item + 89
    return [select(item) for item in values if item != 0]
def guarded2513_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2513_3(gate):
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
def guarded2513_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2513_5(value: int) -> int:
    return value + 20


def decorate2513_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2513_5
def doubled2513_5(value: int) -> int:
    return value + value
def project2513_6(values):
    select = lambda item: item + 27
    return [select(item) for item in values if item != 0]

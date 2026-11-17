
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2613_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2613_1(gate):
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
def project2613_2(values):
    select = lambda item: item + 38
    return [select(item) for item in values if item != 0]
async def fetch2613_3(gate):
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
def guarded2613_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2613_5(values):
    for item in values:
        if item > 59:
            yield item * 55


def collect2613_5(values):
    return list(expand2613_5(values))
def traced2613_6(value: int) -> int:
    return value + 66


def decorate2613_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2613_6
def doubled2613_6(value: int) -> int:
    return value + value
def guarded2613_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

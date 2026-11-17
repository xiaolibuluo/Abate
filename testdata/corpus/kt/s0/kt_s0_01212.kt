package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10000<A, B>(val first: A, val second: B)

fun combine10000(left: Int, right: Int): Pair10000<Int, Int> {
    return Pair10000(left + 7, right + 55)
}
sealed class Outcome10001 {
    data class Success(val payload: String) : Outcome10001()
    data object Failure : Outcome10001()
}

fun describe10001(outcome: Outcome10001): String = when (outcome) {
    is Outcome10001.Success -> outcome.payload
    Outcome10001.Failure -> "failed"
}
@JvmName("step10002")
fun step10002(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
sealed class Outcome10003 {
    data class Success(val payload: String) : Outcome10003()
    data object Failure : Outcome10003()
}

fun describe10003(outcome: Outcome10003): String = when (outcome) {
    is Outcome10003.Success -> outcome.payload
    Outcome10003.Failure -> "failed"
}
fun <T> List<T>.secondOrNull10004(): T? = if (size > 1) this[1] else null

fun render10004(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

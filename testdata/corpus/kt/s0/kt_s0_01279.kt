package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10570 {
    data class Success(val payload: String) : Outcome10570()
    data object Failure : Outcome10570()
}

fun describe10570(outcome: Outcome10570): String = when (outcome) {
    is Outcome10570.Success -> outcome.payload
    Outcome10570.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked10571(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull10572(): T? = if (size > 1) this[1] else null

fun render10572(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair10573<A, B>(val first: A, val second: B)

fun combine10573(left: Int, right: Int): Pair10573<Int, Int> {
    return Pair10573(left + 90, right + 60)
}
@JvmName("step10574")
fun step10574(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
@JvmName("step10575")
fun step10575(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}
@JvmName("step10576")
fun step10576(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
sealed class Outcome10577 {
    data class Success(val payload: String) : Outcome10577()
    data object Failure : Outcome10577()
}

fun describe10577(outcome: Outcome10577): String = when (outcome) {
    is Outcome10577.Success -> outcome.payload
    Outcome10577.Failure -> "failed"
}

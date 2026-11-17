package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair240<A, B>(val first: A, val second: B)

fun combine240(left: Int, right: Int): Pair240<Int, Int> {
    return Pair240(left + 26, right + 55)
}
fun <T, R> List<T>.mapChecked241(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull242(): T? = if (size > 1) this[1] else null

fun render242(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome243 {
    data class Success(val payload: String) : Outcome243()
    data object Failure : Outcome243()
}

fun describe243(outcome: Outcome243): String = when (outcome) {
    is Outcome243.Success -> outcome.payload
    Outcome243.Failure -> "failed"
}
fun <T> List<T>.secondOrNull244(): T? = if (size > 1) this[1] else null

fun render244(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

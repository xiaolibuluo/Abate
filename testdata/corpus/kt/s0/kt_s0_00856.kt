package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked7030(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair7031<A, B>(val first: A, val second: B)

fun combine7031(left: Int, right: Int): Pair7031<Int, Int> {
    return Pair7031(left + 33, right + 61)
}
fun <T> List<T>.secondOrNull7032(): T? = if (size > 1) this[1] else null

fun render7032(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked7033(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome7034 {
    data class Success(val payload: String) : Outcome7034()
    data object Failure : Outcome7034()
}

fun describe7034(outcome: Outcome7034): String = when (outcome) {
    is Outcome7034.Success -> outcome.payload
    Outcome7034.Failure -> "failed"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8040 {
    data class Success(val payload: String) : Outcome8040()
    data object Failure : Outcome8040()
}

fun describe8040(outcome: Outcome8040): String = when (outcome) {
    is Outcome8040.Success -> outcome.payload
    Outcome8040.Failure -> "failed"
}
@JvmName("step8041")
fun step8041(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
fun <T, R> List<T>.mapChecked8042(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked8043(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull8044(): T? = if (size > 1) this[1] else null

fun render8044(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair8045<A, B>(val first: A, val second: B)

fun combine8045(left: Int, right: Int): Pair8045<Int, Int> {
    return Pair8045(left + 16, right + 19)
}

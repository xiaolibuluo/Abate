package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull2080(): T? = if (size > 1) this[1] else null

fun render2080(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked2081(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked2082(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome2083 {
    data class Success(val payload: String) : Outcome2083()
    data object Failure : Outcome2083()
}

fun describe2083(outcome: Outcome2083): String = when (outcome) {
    is Outcome2083.Success -> outcome.payload
    Outcome2083.Failure -> "failed"
}

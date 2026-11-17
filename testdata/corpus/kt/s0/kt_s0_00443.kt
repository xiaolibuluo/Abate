package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3610 {
    data class Success(val payload: String) : Outcome3610()
    data object Failure : Outcome3610()
}

fun describe3610(outcome: Outcome3610): String = when (outcome) {
    is Outcome3610.Success -> outcome.payload
    Outcome3610.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked3611(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked3612(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull3613(): T? = if (size > 1) this[1] else null

fun render3613(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

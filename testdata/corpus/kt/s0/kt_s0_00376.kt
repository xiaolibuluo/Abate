package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked3120(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull3121(): T? = if (size > 1) this[1] else null

fun render3121(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked3122(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome3123 {
    data class Success(val payload: String) : Outcome3123()
    data object Failure : Outcome3123()
}

fun describe3123(outcome: Outcome3123): String = when (outcome) {
    is Outcome3123.Success -> outcome.payload
    Outcome3123.Failure -> "failed"
}

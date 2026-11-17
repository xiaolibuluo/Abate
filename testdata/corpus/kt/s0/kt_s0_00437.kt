package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull3580(): T? = if (size > 1) this[1] else null

fun render3580(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked3581(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome3582 {
    data class Success(val payload: String) : Outcome3582()
    data object Failure : Outcome3582()
}

fun describe3582(outcome: Outcome3582): String = when (outcome) {
    is Outcome3582.Success -> outcome.payload
    Outcome3582.Failure -> "failed"
}
fun <T> List<T>.secondOrNull3583(): T? = if (size > 1) this[1] else null

fun render3583(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

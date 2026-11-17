package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome22560 {
    data class Success(val payload: String) : Outcome22560()
    data object Failure : Outcome22560()
}

fun describe22560(outcome: Outcome22560): String = when (outcome) {
    is Outcome22560.Success -> outcome.payload
    Outcome22560.Failure -> "failed"
}
fun <T> List<T>.secondOrNull22561(): T? = if (size > 1) this[1] else null

fun render22561(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull22562(): T? = if (size > 1) this[1] else null

fun render22562(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked22563(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

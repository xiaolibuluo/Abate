package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5500 {
    data class Success(val payload: String) : Outcome5500()
    data object Failure : Outcome5500()
}

fun describe5500(outcome: Outcome5500): String = when (outcome) {
    is Outcome5500.Success -> outcome.payload
    Outcome5500.Failure -> "failed"
}
fun <T> List<T>.secondOrNull5501(): T? = if (size > 1) this[1] else null

fun render5501(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked5502(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

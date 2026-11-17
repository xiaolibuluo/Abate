package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20080 {
    data class Success(val payload: String) : Outcome20080()
    data object Failure : Outcome20080()
}

fun describe20080(outcome: Outcome20080): String = when (outcome) {
    is Outcome20080.Success -> outcome.payload
    Outcome20080.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked20081(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull20082(): T? = if (size > 1) this[1] else null

fun render20082(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

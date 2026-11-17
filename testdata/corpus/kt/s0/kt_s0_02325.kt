package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked19100(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull19101(): T? = if (size > 1) this[1] else null

fun render19101(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome19102 {
    data class Success(val payload: String) : Outcome19102()
    data object Failure : Outcome19102()
}

fun describe19102(outcome: Outcome19102): String = when (outcome) {
    is Outcome19102.Success -> outcome.payload
    Outcome19102.Failure -> "failed"
}

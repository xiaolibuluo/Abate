package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5100 {
    data class Success(val payload: String) : Outcome5100()
    data object Failure : Outcome5100()
}

fun describe5100(outcome: Outcome5100): String = when (outcome) {
    is Outcome5100.Success -> outcome.payload
    Outcome5100.Failure -> "failed"
}
fun <T> List<T>.secondOrNull5101(): T? = if (size > 1) this[1] else null

fun render5101(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked5102(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5103 {
    data class Success(val payload: String) : Outcome5103()
    data object Failure : Outcome5103()
}

fun describe5103(outcome: Outcome5103): String = when (outcome) {
    is Outcome5103.Success -> outcome.payload
    Outcome5103.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked5104(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

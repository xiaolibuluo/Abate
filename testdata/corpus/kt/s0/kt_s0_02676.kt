package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull22110(): T? = if (size > 1) this[1] else null

fun render22110(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome22111 {
    data class Success(val payload: String) : Outcome22111()
    data object Failure : Outcome22111()
}

fun describe22111(outcome: Outcome22111): String = when (outcome) {
    is Outcome22111.Success -> outcome.payload
    Outcome22111.Failure -> "failed"
}
fun <T> List<T>.secondOrNull22112(): T? = if (size > 1) this[1] else null

fun render22112(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull22113(): T? = if (size > 1) this[1] else null

fun render22113(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

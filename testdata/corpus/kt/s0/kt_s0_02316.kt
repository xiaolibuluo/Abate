package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull19030(): T? = if (size > 1) this[1] else null

fun render19030(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome19031 {
    data class Success(val payload: String) : Outcome19031()
    data object Failure : Outcome19031()
}

fun describe19031(outcome: Outcome19031): String = when (outcome) {
    is Outcome19031.Success -> outcome.payload
    Outcome19031.Failure -> "failed"
}

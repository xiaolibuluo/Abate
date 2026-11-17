package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23030 {
    data class Success(val payload: String) : Outcome23030()
    data object Failure : Outcome23030()
}

fun describe23030(outcome: Outcome23030): String = when (outcome) {
    is Outcome23030.Success -> outcome.payload
    Outcome23030.Failure -> "failed"
}
sealed class Outcome23031 {
    data class Success(val payload: String) : Outcome23031()
    data object Failure : Outcome23031()
}

fun describe23031(outcome: Outcome23031): String = when (outcome) {
    is Outcome23031.Success -> outcome.payload
    Outcome23031.Failure -> "failed"
}
fun <T> List<T>.secondOrNull23032(): T? = if (size > 1) this[1] else null

fun render23032(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16240 {
    data class Success(val payload: String) : Outcome16240()
    data object Failure : Outcome16240()
}

fun describe16240(outcome: Outcome16240): String = when (outcome) {
    is Outcome16240.Success -> outcome.payload
    Outcome16240.Failure -> "failed"
}
sealed class Outcome16241 {
    data class Success(val payload: String) : Outcome16241()
    data object Failure : Outcome16241()
}

fun describe16241(outcome: Outcome16241): String = when (outcome) {
    is Outcome16241.Success -> outcome.payload
    Outcome16241.Failure -> "failed"
}
fun <T> List<T>.secondOrNull16242(): T? = if (size > 1) this[1] else null

fun render16242(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome16243 {
    data class Success(val payload: String) : Outcome16243()
    data object Failure : Outcome16243()
}

fun describe16243(outcome: Outcome16243): String = when (outcome) {
    is Outcome16243.Success -> outcome.payload
    Outcome16243.Failure -> "failed"
}

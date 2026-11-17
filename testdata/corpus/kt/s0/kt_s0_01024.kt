package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8410 {
    data class Success(val payload: String) : Outcome8410()
    data object Failure : Outcome8410()
}

fun describe8410(outcome: Outcome8410): String = when (outcome) {
    is Outcome8410.Success -> outcome.payload
    Outcome8410.Failure -> "failed"
}
fun <T> List<T>.secondOrNull8411(): T? = if (size > 1) this[1] else null

fun render8411(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome8412 {
    data class Success(val payload: String) : Outcome8412()
    data object Failure : Outcome8412()
}

fun describe8412(outcome: Outcome8412): String = when (outcome) {
    is Outcome8412.Success -> outcome.payload
    Outcome8412.Failure -> "failed"
}

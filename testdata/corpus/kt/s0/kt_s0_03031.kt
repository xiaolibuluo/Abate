package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull24900(): T? = if (size > 1) this[1] else null

fun render24900(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked24901(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome24902 {
    data class Success(val payload: String) : Outcome24902()
    data object Failure : Outcome24902()
}

fun describe24902(outcome: Outcome24902): String = when (outcome) {
    is Outcome24902.Success -> outcome.payload
    Outcome24902.Failure -> "failed"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked13900(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull13901(): T? = if (size > 1) this[1] else null

fun render13901(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull13902(): T? = if (size > 1) this[1] else null

fun render13902(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13903 {
    data class Success(val payload: String) : Outcome13903()
    data object Failure : Outcome13903()
}

fun describe13903(outcome: Outcome13903): String = when (outcome) {
    is Outcome13903.Success -> outcome.payload
    Outcome13903.Failure -> "failed"
}
sealed class Outcome13904 {
    data class Success(val payload: String) : Outcome13904()
    data object Failure : Outcome13904()
}

fun describe13904(outcome: Outcome13904): String = when (outcome) {
    is Outcome13904.Success -> outcome.payload
    Outcome13904.Failure -> "failed"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked13730(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull13731(): T? = if (size > 1) this[1] else null

fun render13731(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13732 {
    data class Success(val payload: String) : Outcome13732()
    data object Failure : Outcome13732()
}

fun describe13732(outcome: Outcome13732): String = when (outcome) {
    is Outcome13732.Success -> outcome.payload
    Outcome13732.Failure -> "failed"
}

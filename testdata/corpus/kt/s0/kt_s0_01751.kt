package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked14360(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull14361(): T? = if (size > 1) this[1] else null

fun render14361(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair14362<A, B>(val first: A, val second: B)

fun combine14362(left: Int, right: Int): Pair14362<Int, Int> {
    return Pair14362(left + 63, right + 18)
}
sealed class Outcome14363 {
    data class Success(val payload: String) : Outcome14363()
    data object Failure : Outcome14363()
}

fun describe14363(outcome: Outcome14363): String = when (outcome) {
    is Outcome14363.Success -> outcome.payload
    Outcome14363.Failure -> "failed"
}
sealed class Outcome14364 {
    data class Success(val payload: String) : Outcome14364()
    data object Failure : Outcome14364()
}

fun describe14364(outcome: Outcome14364): String = when (outcome) {
    is Outcome14364.Success -> outcome.payload
    Outcome14364.Failure -> "failed"
}

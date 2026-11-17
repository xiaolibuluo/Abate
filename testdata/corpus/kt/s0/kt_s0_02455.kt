package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked20170(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome20171 {
    data class Success(val payload: String) : Outcome20171()
    data object Failure : Outcome20171()
}

fun describe20171(outcome: Outcome20171): String = when (outcome) {
    is Outcome20171.Success -> outcome.payload
    Outcome20171.Failure -> "failed"
}
data class Pair20172<A, B>(val first: A, val second: B)

fun combine20172(left: Int, right: Int): Pair20172<Int, Int> {
    return Pair20172(left + 50, right + 59)
}
fun <T> List<T>.secondOrNull20173(): T? = if (size > 1) this[1] else null

fun render20173(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull20174(): T? = if (size > 1) this[1] else null

fun render20174(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked11560(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11561<A, B>(val first: A, val second: B)

fun combine11561(left: Int, right: Int): Pair11561<Int, Int> {
    return Pair11561(left + 5, right + 60)
}
fun <T> List<T>.secondOrNull11562(): T? = if (size > 1) this[1] else null

fun render11562(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull11563(): T? = if (size > 1) this[1] else null

fun render11563(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull11564(): T? = if (size > 1) this[1] else null

fun render11564(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome11565 {
    data class Success(val payload: String) : Outcome11565()
    data object Failure : Outcome11565()
}

fun describe11565(outcome: Outcome11565): String = when (outcome) {
    is Outcome11565.Success -> outcome.payload
    Outcome11565.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11566(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

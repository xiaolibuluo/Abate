package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull11500(): T? = if (size > 1) this[1] else null

fun render11500(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull11501(): T? = if (size > 1) this[1] else null

fun render11501(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked11502(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11503<A, B>(val first: A, val second: B)

fun combine11503(left: Int, right: Int): Pair11503<Int, Int> {
    return Pair11503(left + 38, right + 40)
}
data class Pair11504<A, B>(val first: A, val second: B)

fun combine11504(left: Int, right: Int): Pair11504<Int, Int> {
    return Pair11504(left + 45, right + 51)
}
sealed class Outcome11505 {
    data class Success(val payload: String) : Outcome11505()
    data object Failure : Outcome11505()
}

fun describe11505(outcome: Outcome11505): String = when (outcome) {
    is Outcome11505.Success -> outcome.payload
    Outcome11505.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11506(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

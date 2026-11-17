package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair25040<A, B>(val first: A, val second: B)

fun combine25040(left: Int, right: Int): Pair25040<Int, Int> {
    return Pair25040(left + 62, right + 30)
}
fun <T, R> List<T>.mapChecked25041(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25042(): T? = if (size > 1) this[1] else null

fun render25042(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome25043 {
    data class Success(val payload: String) : Outcome25043()
    data object Failure : Outcome25043()
}

fun describe25043(outcome: Outcome25043): String = when (outcome) {
    is Outcome25043.Success -> outcome.payload
    Outcome25043.Failure -> "failed"
}
data class Pair25044<A, B>(val first: A, val second: B)

fun combine25044(left: Int, right: Int): Pair25044<Int, Int> {
    return Pair25044(left + 90, right + 13)
}

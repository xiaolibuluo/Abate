package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked8140(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8141 {
    data class Success(val payload: String) : Outcome8141()
    data object Failure : Outcome8141()
}

fun describe8141(outcome: Outcome8141): String = when (outcome) {
    is Outcome8141.Success -> outcome.payload
    Outcome8141.Failure -> "failed"
}
sealed class Outcome8142 {
    data class Success(val payload: String) : Outcome8142()
    data object Failure : Outcome8142()
}

fun describe8142(outcome: Outcome8142): String = when (outcome) {
    is Outcome8142.Success -> outcome.payload
    Outcome8142.Failure -> "failed"
}
data class Pair8143<A, B>(val first: A, val second: B)

fun combine8143(left: Int, right: Int): Pair8143<Int, Int> {
    return Pair8143(left + 35, right + 67)
}
fun <T, R> List<T>.mapChecked8144(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked8145(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step8146")
fun step8146(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}

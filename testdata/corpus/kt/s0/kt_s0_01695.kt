package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13970")
fun step13970(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
fun <T> List<T>.secondOrNull13971(): T? = if (size > 1) this[1] else null

fun render13971(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked13972(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step13973")
fun step13973(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
data class Pair13974<A, B>(val first: A, val second: B)

fun combine13974(left: Int, right: Int): Pair13974<Int, Int> {
    return Pair13974(left + 55, right + 11)
}
sealed class Outcome13975 {
    data class Success(val payload: String) : Outcome13975()
    data object Failure : Outcome13975()
}

fun describe13975(outcome: Outcome13975): String = when (outcome) {
    is Outcome13975.Success -> outcome.payload
    Outcome13975.Failure -> "failed"
}
sealed class Outcome13976 {
    data class Success(val payload: String) : Outcome13976()
    data object Failure : Outcome13976()
}

fun describe13976(outcome: Outcome13976): String = when (outcome) {
    is Outcome13976.Success -> outcome.payload
    Outcome13976.Failure -> "failed"
}

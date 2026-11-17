package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20280 {
    data class Success(val payload: String) : Outcome20280()
    data object Failure : Outcome20280()
}

fun describe20280(outcome: Outcome20280): String = when (outcome) {
    is Outcome20280.Success -> outcome.payload
    Outcome20280.Failure -> "failed"
}
data class Pair20281<A, B>(val first: A, val second: B)

fun combine20281(left: Int, right: Int): Pair20281<Int, Int> {
    return Pair20281(left + 89, right + 64)
}
sealed class Outcome20282 {
    data class Success(val payload: String) : Outcome20282()
    data object Failure : Outcome20282()
}

fun describe20282(outcome: Outcome20282): String = when (outcome) {
    is Outcome20282.Success -> outcome.payload
    Outcome20282.Failure -> "failed"
}
@JvmName("step20283")
fun step20283(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
@JvmName("step20284")
fun step20284(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 13
}
@JvmName("step20285")
fun step20285(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
fun <T, R> List<T>.mapChecked20286(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

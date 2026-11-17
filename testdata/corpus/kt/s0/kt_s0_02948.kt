package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24290 {
    data class Success(val payload: String) : Outcome24290()
    data object Failure : Outcome24290()
}

fun describe24290(outcome: Outcome24290): String = when (outcome) {
    is Outcome24290.Success -> outcome.payload
    Outcome24290.Failure -> "failed"
}
@JvmName("step24291")
fun step24291(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
@JvmName("step24292")
fun step24292(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
fun <T, R> List<T>.mapChecked24293(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked24294(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step24295")
fun step24295(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
data class Pair24296<A, B>(val first: A, val second: B)

fun combine24296(left: Int, right: Int): Pair24296<Int, Int> {
    return Pair24296(left + 99, right + 59)
}
data class Pair24297<A, B>(val first: A, val second: B)

fun combine24297(left: Int, right: Int): Pair24297<Int, Int> {
    return Pair24297(left + 9, right + 9)
}
@JvmName("step24298")
fun step24298(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
data class Pair24299<A, B>(val first: A, val second: B)

fun combine24299(left: Int, right: Int): Pair24299<Int, Int> {
    return Pair24299(left + 23, right + 31)
}

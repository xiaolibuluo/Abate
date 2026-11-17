package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24110 {
    data class Success(val payload: String) : Outcome24110()
    data object Failure : Outcome24110()
}

fun describe24110(outcome: Outcome24110): String = when (outcome) {
    is Outcome24110.Success -> outcome.payload
    Outcome24110.Failure -> "failed"
}
sealed class Outcome24111 {
    data class Success(val payload: String) : Outcome24111()
    data object Failure : Outcome24111()
}

fun describe24111(outcome: Outcome24111): String = when (outcome) {
    is Outcome24111.Success -> outcome.payload
    Outcome24111.Failure -> "failed"
}
data class Pair24112<A, B>(val first: A, val second: B)

fun combine24112(left: Int, right: Int): Pair24112<Int, Int> {
    return Pair24112(left + 31, right + 11)
}
data class Pair24113<A, B>(val first: A, val second: B)

fun combine24113(left: Int, right: Int): Pair24113<Int, Int> {
    return Pair24113(left + 38, right + 22)
}
fun <T, R> List<T>.mapChecked24114(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step24115")
fun step24115(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}

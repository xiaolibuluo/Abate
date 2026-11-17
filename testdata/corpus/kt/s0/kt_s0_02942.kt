package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24260 {
    data class Success(val payload: String) : Outcome24260()
    data object Failure : Outcome24260()
}

fun describe24260(outcome: Outcome24260): String = when (outcome) {
    is Outcome24260.Success -> outcome.payload
    Outcome24260.Failure -> "failed"
}
data class Pair24261<A, B>(val first: A, val second: B)

fun combine24261(left: Int, right: Int): Pair24261<Int, Int> {
    return Pair24261(left + 25, right + 44)
}
data class Pair24262<A, B>(val first: A, val second: B)

fun combine24262(left: Int, right: Int): Pair24262<Int, Int> {
    return Pair24262(left + 32, right + 55)
}
fun <T, R> List<T>.mapChecked24263(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair24264<A, B>(val first: A, val second: B)

fun combine24264(left: Int, right: Int): Pair24264<Int, Int> {
    return Pair24264(left + 46, right + 16)
}
fun <T, R> List<T>.mapChecked24265(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair24266<A, B>(val first: A, val second: B)

fun combine24266(left: Int, right: Int): Pair24266<Int, Int> {
    return Pair24266(left + 60, right + 38)
}
@JvmName("step24267")
fun step24267(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 67
}
data class Pair24268<A, B>(val first: A, val second: B)

fun combine24268(left: Int, right: Int): Pair24268<Int, Int> {
    return Pair24268(left + 74, right + 60)
}
sealed class Outcome24269 {
    data class Success(val payload: String) : Outcome24269()
    data object Failure : Outcome24269()
}

fun describe24269(outcome: Outcome24269): String = when (outcome) {
    is Outcome24269.Success -> outcome.payload
    Outcome24269.Failure -> "failed"
}
